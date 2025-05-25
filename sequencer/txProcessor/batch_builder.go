package txprocessor

import (
	"fmt"
	"math/big"
	"os"

	ethCommon "github.com/ethereum/go-ethereum/common"
	"github.com/iden3/go-merkletree"
	"github.com/iden3/go-merkletree/db/pebble"
	"github.com/tokamak-network/syb-sequencer/sequencer/common"
	"github.com/tokamak-network/syb-sequencer/sequencer/db/statedb"
	"github.com/tokamak-network/syb-sequencer/sequencer/forger"
)

// Config contains the BatchBuilder configuration parameters
type Config struct {
	NLevels uint32
	MaxTx   uint32
	MaxL1Tx uint32
	ChainID uint64
}

type BatchBuilder struct {
	forger  *forger.Forger
	config  Config
	zki     *common.ZKInputs
	txIndex int
}

func NewBatchBuilder(forger *forger.Forger, config Config) *BatchBuilder {
	return &BatchBuilder{
		forger:  forger,
		config:  config,
		zki:     nil,
		txIndex: 0,
	}
}

// Resets the zkInputs and transaction index
func (batchBuilder *BatchBuilder) resetZKInputs() {
	batchBuilder.zki = nil
	batchBuilder.txIndex = 0
}

// forgeTransactions processes L1 user transactions, updates the state via forger.statedb, and generates ZKInputs.
func (batchBuilder *BatchBuilder) forgeTransactions(l1UserTxs []common.Tx) (*common.ZKInputs, error) {
	batchBuilder.resetZKInputs()

	// Access StateDB via the forger
	sdb := batchBuilder.forger.Statedb
	if sdb == nil {
		return nil, fmt.Errorf("BatchBuilder's forger does not have an initialized StateDB")
	}

	currentBatchNum := uint32(sdb.CurrentBatch() + 1)
	batchBuilder.zki = common.NewZKInputs(
		batchBuilder.config.ChainID,
		batchBuilder.config.MaxTx,
		batchBuilder.config.MaxL1Tx,
		batchBuilder.config.NLevels,
		&currentBatchNum,
	)
	oldLastIdx := uint32(sdb.CurrentAccountIdx())
	batchBuilder.zki.OldLastIdx = &oldLastIdx
	batchBuilder.zki.OldAccountRoot = sdb.GetATRoot()
	batchBuilder.zki.NewLastIdxRaw = uint32(sdb.CurrentAccountIdx())
	batchBuilder.zki.OldVouchRoot = sdb.GetVTRoot()
	batchBuilder.zki.OldScoreRoot = sdb.GetSTRoot()

	var exitTree *merkletree.MerkleTree
	tmpDir, err := os.MkdirTemp("", "batchbuilder-exittree-")
	if err != nil {
		return nil, common.Wrap(fmt.Errorf("failed to create temp dir for exit tree: %w", err))
	}
	defer func() {
		if rErr := os.RemoveAll(tmpDir); rErr != nil {
			fmt.Printf("Warning: failed to remove temp exit tree dir %s: %v\n", tmpDir, rErr)
		}
	}()

	sto, err := pebble.NewPebbleStorage(tmpDir, false)
	if err != nil {
		return nil, common.Wrap(fmt.Errorf("failed to create pebble storage for exit tree: %w", err))
	}
	defer sto.Close()

	// exitTree, err = merkletree.NewMerkleTree(sto, accountTree.MaxLevels())

	exitTree, err = merkletree.NewMerkleTree(sto, int(batchBuilder.config.NLevels))
	if err != nil {
		return nil, common.Wrap(fmt.Errorf("failed to create exit merkle tree: %w", err))
	}

	if len(l1UserTxs) > int(batchBuilder.config.MaxL1Tx) {
		return nil, common.Wrap(fmt.Errorf("number of L1UserTxs (%d) exceeds MaxL1Tx (%d)", len(l1UserTxs), batchBuilder.config.MaxL1Tx))
	}
	if len(l1UserTxs) > int(batchBuilder.config.MaxTx) {
		return nil, common.Wrap(fmt.Errorf("number of L1UserTxs (%d) exceeds MaxTx (%d)", len(l1UserTxs), batchBuilder.config.MaxTx))
	}

	for i, tx := range l1UserTxs {
		batchBuilder.txIndex = i // Set current transaction index for ZKI population

		currentTx := tx

		switch currentTx.Type {
		case common.TxTypeCreateAccountDeposit:
			err = batchBuilder.applyCreateAccount(sdb, &currentTx)
		case common.TxTypeDeposit:
			err = batchBuilder.applyDeposit(sdb, &currentTx)
		// case common.TxTypeForceExit:
		// 	exitAccount, newExit, err = batchBuilder.applyExit(sdb, exitTree, currentTx.Tx(), currentTx.Amount) // currentTx.Amount is already effective amount
		// 	if err == nil && exitAccount != nil {                                                               // Only set if exit was processed
		// 		exitIdxForZKI = &currentTx.FromIdx
		// 	}
		case common.TxTypeCreateVouch, common.TxTypeUnVouch:
			err = batchBuilder.applyVouch(sdb, currentTx, common.AccountIdx(currentTx.ToIdx), currentTx.Type)
		default:
			err = fmt.Errorf("unknown L1 transaction type: %s for txID: %s", currentTx.Type, currentTx.ItemID)

			if err != nil {
				return nil, common.Wrap(fmt.Errorf("failed to process tx %s (type %s): %w", currentTx.ItemID, currentTx.Type, err))
			}
		}

		// Final ZKI parameters
		globalChainIDVal := uint16(batchBuilder.config.ChainID)
		batchBuilder.zki.GlobalChainID = &globalChainIDVal
		batchBuilder.zki.NewAccountRootRaw = sdb.GetATRootHash()
		batchBuilder.zki.NewVouchRootRaw = sdb.GetVTRootHash()
		// batchBuilder.zki.NewScoreRootRaw = sdb.GetMTRootScore()
		if exitTree != nil {
			batchBuilder.zki.NewExitRootRaw = exitTree.Root()
		}

		// Make a checkpoint in the StateDB after processing all transactions for this batch
		if err := sdb.MakeCheckpoint(); err != nil {
			return nil, common.Wrap(fmt.Errorf("failed to make checkpoint in StateDB: %w", err))
		}
	}
	return batchBuilder.zki, nil
}

// applyCreateAccount creates a new account and updates ZKInputs.
// It now takes sdb (*statedb.LocalStateDB) as a parameter.
func (batchBuilder *BatchBuilder) applyCreateAccount(sdb *statedb.LocalStateDB, tx *common.Tx) error {
	account := &common.Account{
		// Nonce:   0,
		Balance: tx.Amount, // Use effective deposit amount
		// BJJ:     tx.FromBJJ,
		EthAddr: ethCommon.BytesToAddress(tx.FromEthAddr),
	}

	newAccountIdx := common.AccountIdx(sdb.CurrentAccountIdx() + 1)
	p, err := sdb.CreateAccount(newAccountIdx, account)
	if err != nil {
		return common.Wrap(fmt.Errorf("applyCreateAccount: failed to create account %d: %w", newAccountIdx, err))
	}

	// Populate ZKI for state 1 (the new account)
	// batchBuilder.zki.TokenID1[batchBuilder.txIndex] = tx.TokenID.BigInt() // If TokenID is part of L1Tx / Account
	// batchBuilder.zki.Nonce1[batchBuilder.txIndex] = big.NewInt(0)
	// fromBJJSign, fromBJJY := babyjub.UnpackSignY(tx.FromBJJ)
	// bjjSignBool := fromBJJSign
	// batchBuilder.zki.Sign1[batchBuilder.txIndex] = &bjjSignBool
	// batchBuilder.zki.Ay1[batchBuilder.txIndex] = fromBJJY
	batchBuilder.zki.Balance1[batchBuilder.txIndex] = new(big.Int).Set(tx.Amount)
	batchBuilder.zki.EthAddr1[batchBuilder.txIndex] = common.EthAddrToBigInt(account.EthAddr)
	batchBuilder.zki.Siblings1[batchBuilder.txIndex] = siblingsToZKInputFormat(p.Siblings)

	isOld0Val := p.IsOld0
	batchBuilder.zki.IsOld0_1[batchBuilder.txIndex] = &isOld0Val
	batchBuilder.zki.OldKey1[batchBuilder.txIndex] = p.OldKey.BigInt()
	batchBuilder.zki.OldValue1[batchBuilder.txIndex] = p.OldValue.BigInt()

	auxFromIdxVal := uint32(newAccountIdx)
	batchBuilder.zki.AuxFromIdx[batchBuilder.txIndex] = &auxFromIdxVal
	newAccountCreatedVal := true
	batchBuilder.zki.NewAccount[batchBuilder.txIndex] = &newAccountCreatedVal

	// Update NewLastIdxRaw in ZKI as an account was created
	batchBuilder.zki.NewLastIdxRaw = uint32(newAccountIdx)

	err = sdb.SetCurrentAccountIdx(newAccountIdx)
	if err != nil {
		return common.Wrap(fmt.Errorf("applyCreateAccount: failed to update current account index to %d: %w", newAccountIdx, err))
	}
	return nil
}

// applyDeposit updates an existing account's balance and ZKInputs.
// It now takes sdb (*statedb.LocalStateDB) as a parameter.
func (batchBuilder *BatchBuilder) applyDeposit(sdb *statedb.LocalStateDB, tx *common.Tx) error {
	if tx.FromIdx == nil {
		return fmt.Errorf("applyDeposit: tx.FromIdx is nil for txID %s", tx.ItemID)
	}
	senderAccountIdx := common.AccountIdx(*tx.FromIdx)
	accSender, err := sdb.GetAccount(senderAccountIdx)
	if err != nil {
		return common.Wrap(fmt.Errorf("applyDeposit: failed to get sender account %d: %w", tx.FromIdx, err))
	}

	// Populate ZKI for state 1 (sender account before update)
	// batchBuilder.zki.TokenID1[batchBuilder.txIndex] = accSender.TokenID.BigInt() // If TokenID is part of Account
	// batchBuilder.zki.Nonce1[batchBuilder.txIndex] = new(big.Int).Set(accSender.Nonce.BigInt())
	// senderBJJSign, senderBJJY := babyjub.UnpackSignY(accSender.BJJ)
	// bjjSignBool := senderBJJSign
	// batchBuilder.zki.Sign1[batchBuilder.txIndex] = &bjjSignBool
	// batchBuilder.zki.Ay1[batchBuilder.txIndex] = senderBJJY
	batchBuilder.zki.Balance1[batchBuilder.txIndex] = new(big.Int).Set(accSender.Balance)
	batchBuilder.zki.EthAddr1[batchBuilder.txIndex] = common.EthAddrToBigInt(accSender.EthAddr)

	// Add the deposit to the sender
	accSender.Balance.Add(accSender.Balance, tx.Amount)

	if accSender.Balance.Cmp(big.NewInt(0)) == -1 {
		return fmt.Errorf("applyDeposit: sender %d balance became negative: %s", tx.FromIdx, accSender.Balance.String()) // Or use newErrorNotEnoughBalance
	}

	p, err := sdb.UpdateAccount(senderAccountIdx, accSender)
	if err != nil {
		return common.Wrap(fmt.Errorf("applyDeposit: failed to update sender account %d: %w", tx.FromIdx, err))
	}
	batchBuilder.zki.Siblings1[batchBuilder.txIndex] = siblingsToZKInputFormat(p.Siblings)

	return nil
}

// applyVouch handles vouch creation/deletion.
// It now takes sdb (*statedb.LocalStateDB) as a parameter.
func (batchBuilder *BatchBuilder) applyVouch(sdb *statedb.LocalStateDB, tx common.Tx, auxToIdx common.AccountIdx, txType string) error {
	if tx.FromIdx == nil {
		return fmt.Errorf("applyVouch: tx.FromIdx is nil")
	}
	fromAccountIdx := common.AccountIdx(*tx.FromIdx)
	toAccountIdx := auxToIdx // This is the vouchee

	// Construct VouchIdx from the two account indices
	// vouchIdx, err := common.NewVouchIdx(fromAccountIdx, toAccountIdx)
	// if err != nil {
	// 	return common.Wrap(fmt.Errorf("applyVouch: failed to create VouchIdx from %d and %d: %w", fromAccountIdx, toAccountIdx, err))
	// }

	var vouchProof *merkletree.CircomProcessorProof
	var err error

	// Perform Vouch/UnVouch operation on the Vouch Tree
	switch txType {
	case common.TxTypeCreateVouch:
		// The *common.Vouch argument to sdb.Vouch might be for additional details,
		// but based on statedb/vouch.go, it might not be strictly used if it only calls CreateVouchInTreeDB.
		// Passing a constructed one for completeness or future use.
		vouchDetails := &common.Vouch{FromIdx: fromAccountIdx, ToIdx: toAccountIdx}
		vouchProof, err = sdb.Vouch(common.VouchIdx(fromAccountIdx), vouchDetails)
		if err != nil {
			return common.Wrap(fmt.Errorf("applyVouch: failed to create vouch for VouchIdx %s: %w", common.VouchIdx(fromAccountIdx).String(), err))
		}
	case common.TxTypeUnVouch:
		vouchProof, err = sdb.UnVouch(common.VouchIdx(fromAccountIdx))
		if err != nil {
			return common.Wrap(fmt.Errorf("applyVouch: failed to delete vouch for VouchIdx %s: %w", common.VouchIdx(fromAccountIdx).String(), err))
		}
	default:
		return fmt.Errorf("applyVouch: unsupported txType for vouch operation: %s", txType)
	}

	// --- Populate ZKI ---

	// 1. Populate ZKI fields for the Vouch Tree Merkle operation (using "State1" fields)
	// These ZKI field names (Key1, OldValue1, NewValue1, IsOld0_1) are assumed.
	// Adjust them if your ZKInputs struct uses different names for generic Merkle tree operations.
	batchBuilder.zki.OldValue1[batchBuilder.txIndex] = vouchProof.OldValue.BigInt() // Previous state of vouch (0 or 1)
	isOld0Bool := vouchProof.IsOld0
	batchBuilder.zki.IsOld0_1[batchBuilder.txIndex] = &isOld0Bool
	batchBuilder.zki.Siblings1[batchBuilder.txIndex] = siblingsToZKInputFormat(vouchProof.Siblings) // Siblings from Vouch Tree

	// 2. Populate ZKI fields with sender (voucher) account information (read-only)
	// These are for context if the circuit needs them, accounts are not updated here.
	accSender, err := sdb.GetAccount(fromAccountIdx)
	if err != nil {
		return common.Wrap(fmt.Errorf("applyVouch: failed to get sender account %d for ZKI: %w", fromAccountIdx, err))
	}
	// batchBuilder.zki.Nonce1[batchBuilder.txIndex] = new(big.Int).Set(accSender.Nonce.BigInt())
	// senderBJJSign, senderBJJY := babyjub.UnpackSignY(accSender.BJJ)
	// bjjSignBool1 := senderBJJSign
	// batchBuilder.zki.Sign1[batchBuilder.txIndex] = &bjjSignBool1
	// batchBuilder.zki.Ay1[batchBuilder.txIndex] = senderBJJY
	batchBuilder.zki.Balance1[batchBuilder.txIndex] = new(big.Int).Set(accSender.Balance)
	batchBuilder.zki.EthAddr1[batchBuilder.txIndex] = common.EthAddrToBigInt(accSender.EthAddr)

	// 3. Populate ZKI fields with receiver (vouchee) account information (read-only)
	accReceiver, err := sdb.GetAccount(toAccountIdx)
	if err != nil {
		return common.Wrap(fmt.Errorf("applyVouch: failed to get receiver account %d for ZKI: %w", toAccountIdx, err))
	}
	// batchBuilder.zki.Nonce2[batchBuilder.txIndex] = new(big.Int).Set(accReceiver.Nonce.BigInt())
	// receiverBJJSign, receiverBJJY := babyjub.UnpackSignY(accReceiver.BJJ)
	// bjjSignBool2 := receiverBJJSign
	// batchBuilder.zki.Sign2[batchBuilder.txIndex] = &bjjSignBool2
	// batchBuilder.zki.Ay2[batchBuilder.txIndex] = receiverBJJY
	batchBuilder.zki.Balance2[batchBuilder.txIndex] = new(big.Int).Set(accReceiver.Balance)
	batchBuilder.zki.EthAddr2[batchBuilder.txIndex] = common.EthAddrToBigInt(accReceiver.EthAddr)

	// 4. Siblings2 would be for a secondary Merkle proof, which is not generated by this Vouch/UnVouch logic.
	// Set to empty/default if the ZKI slot must be filled.
	batchBuilder.zki.Siblings2[batchBuilder.txIndex] = siblingsToZKInputFormat(nil) // Or an empty slice of the correct type

	// batchBuilder.zki.FromIdx[batchBuilder.txIndex] = fromAccountIdx
	// batchBuilder.zki.ToIdx[batchBuilder.txIndex] = toAccountIdx.BigInt()

	return nil
}

// siblingsToZKInputFormat converts Merkle tree siblings to the format expected by ZKInputs.
func siblingsToZKInputFormat(siblings []*merkletree.Hash) []*big.Int {
	zkSiblings := make([]*big.Int, len(siblings))
	for i, s := range siblings {
		if s != nil {
			zkSiblings[i] = s.BigInt()
		} else {
			zkSiblings[i] = big.NewInt(0)
		}
	}
	return zkSiblings
}
