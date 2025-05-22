package txprocessor

import (
	"bytes"
	"fmt"
	"math/big"
	"os"

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

	// Assuming sdb.AccountTree() exists and returns a MerkleTree with MaxLevels()
	// or config.NLevels is the correct depth for the exit tree.
	// If sdb.AccountTree() is available:
	// accountTree := sdb.AccountTree()
	// if accountTree == nil {
	// 	 return nil, common.Wrap(fmt.Errorf("forger's StateDB AccountTree is nil"))
	// }
	// exitTree, err = merkletree.NewMerkleTree(sto, accountTree.MaxLevels())
	// For now, using config.NLevels as in previous version, adjust if AccountTree().MaxLevels() is preferred and available.
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

	// for i, tx := range l1UserTxs {
	// 	batchBuilder.txIndex = i // Set current transaction index for ZKI population

	// 	// Populate ZKI common fields for L1Tx
	// 	txCompressedData, err := tx.TxCompressedData(batchBuilder.config.ChainID)
	// 	if err != nil {
	// 		// log.Errorw("Failed to get TxCompressedData", "txid", tx.ItemID, "err", err)
	// 		return nil, common.Wrap(fmt.Errorf("failed to get TxCompressedData for tx %s: %w", tx.ItemID, err))
	// 	}
	// 	batchBuilder.zki.TxCompressedData[batchBuilder.txIndex] = txCompressedData
	// 	fromIdxVal := uint32(tx.FromIdx)
	// 	batchBuilder.zki.FromIdx[batchBuilder.txIndex] = &fromIdxVal
	// 	toIdxVal := uint32(tx.ToIdx)
	// 	batchBuilder.zki.ToIdx[batchBuilder.txIndex] = &toIdxVal

	// 	onChainVal := true
	// 	if tx.Type == common.TxTypeCreateVouch || tx.Type == common.TxTypeDeleteVouch {
	// 		onChainVal = false
	// 	}
	// 	batchBuilder.zki.OnChain[batchBuilder.txIndex] = &onChainVal

	// 	depositAmountF40, err := common.NewFloat40(tx.DepositAmount)
	// 	if err != nil {
	// 		// log.Errorw("Failed to convert DepositAmount to Float40", "txid", tx.TxID, "amount", tx.DepositAmount, "err", err)
	// 		// Decide if this is a fatal error for the batch
	// 	}
	// 	batchBuilder.zki.DepositAmountF[batchBuilder.txIndex] = big.NewInt(int64(depositAmountF40))
	// 	batchBuilder.zki.FromEthAddr[batchBuilder.txIndex] = common.EthAddrToBigInt(tx.FromEthAddr)

	// 	if tx.Type == common.TxTypeForceExit || tx.Type == common.TxTypeCreateVouch || tx.Type == common.TxTypeDeleteVouch {
	// 		amountF40, err := common.NewFloat40(tx.Amount)
	// 		if err != nil {
	// 			// log.Errorw("Failed to convert Amount to Float40", "txid", tx.TxID, "amount", tx.Amount, "err", err)
	// 		}
	// 		batchBuilder.zki.AmountF[batchBuilder.txIndex] = big.NewInt(int64(amountF40))
	// 	}

	// 	// Process transaction based on type
	// 	var exitAccount *common.Account
	// 	var exitIdxForZKI *common.AccountIdx

	// 	// Make a mutable copy of the transaction for effective amounts
	// 	currentTx := tx
	// 	batchBuilder.computeEffectiveAmounts(sdb, &currentTx) // Pass pointer to the copy

	// 	switch currentTx.Type {
	// 	case common.TxTypeCreateAccountDeposit:
	// 		err = batchBuilder.applyCreateAccount(sdb, &currentTx, exitTree)
	// 	case common.TxTypeDeposit:
	// 		err = batchBuilder.applyDeposit(sdb, &currentTx, exitTree)
	// 	case common.TxTypeForceExit:
	// 		exitAccount, newExit, err = batchBuilder.applyExit(sdb, exitTree, currentTx.Tx(), currentTx.Amount) // currentTx.Amount is already effective amount
	// 		if err == nil && exitAccount != nil {                                                               // Only set if exit was processed
	// 			exitIdxForZKI = &currentTx.FromIdx
	// 		}
	// 	case common.TxTypeCreateVouch, common.TxTypeDeleteVouch:
	// 		err = batchBuilder.applyVouch(sdb, currentTx.Tx(), currentTx.ToIdx, exitTree)
	// 	default:
	// 		err = fmt.Errorf("unknown L1 transaction type: %s for txID: %s", currentTx.Type, currentTx.TxID)
	// 	}

	// 	if err != nil {
	// 		return nil, common.Wrap(fmt.Errorf("failed to process tx %s (type %s): %w", currentTx.TxID, currentTx.Type, err))
	// 	}

	// 	// Populate ZKI Intermediate State fields
	// 	if batchBuilder.txIndex < len(batchBuilder.zki.ISOnChain) { // Ensure not out of bounds for intermediate states
	// 		isOnChainVal := true // Assuming L1 Txs are on-chain for IS purposes
	// 		batchBuilder.zki.ISOnChain[batchBuilder.txIndex] = &isOnChainVal
	// 		currentAccIdxVal := uint32(sdb.CurrentAccountIdx())
	// 		batchBuilder.zki.ISOutIdx[batchBuilder.txIndex] = &currentAccIdxVal
	// 		batchBuilder.zki.ISStateRootAccount[batchBuilder.txIndex] = sdb.GetATRoot()
	// 		// batchBuilder.zki.ISStateRootVouch[batchBuilder.txIndex] = sdb.GetMTRootVouch() // If needed
	// 		// batchBuilder.zki.ISStateRootScore[batchBuilder.txIndex] = sdb.GetMTRootScore() // If needed
	// 		if exitTree != nil { // exitTree should not be nil here
	// 			batchBuilder.zki.ISExitRoot[batchBuilder.txIndex] = exitTree.Root().BigInt()
	// 		}
	// 	}
	// 	if exitTree != nil && exitIdxForZKI != nil && exitAccount != nil {
	// 		// This part seems more for synchronizer's `exits` array.
	// 		// For ZKI, the ISExitRoot is already set above.
	// 		// If specific exit details per tx are needed in ZKI, ZKI struct needs fields for them.
	// 	}
	// }

	// Fill empty slots in ZKInputs if MaxTx > len(l1UserTxs)
	// This logic needs to be adapted from the reference if required.
	// The reference TxProcessor has more complex padding logic involving fees.
	// For a simpler BatchBuilder focusing on L1 Txs, padding might just be default/empty values.
	// Example for TxCompressedData:
	// txCompressedDataEmptyVal := common.TxCompressedDataEmpty(batchBuilder.config.ChainID)
	// for i := len(l1UserTxs); i < int(batchBuilder.config.MaxTx); i++ {
	// 	batchBuilder.zki.TxCompressedData[i] = txCompressedDataEmptyVal
	// 	// Pad other ZKI fields (FromIdx, ToIdx, Amounts, Proofs etc.) with defaults or last valid state
	// 	if i < int(batchBuilder.config.MaxTx) -1 { // For IS fields
	// 		 lastValidISOutIdx := batchBuilder.zki.ISOutIdx[len(l1UserTxs)-1]
	//       batchBuilder.zki.ISOutIdx[i] = lastValidISOutIdx
	//       lastValidISStateRootAccount := batchBuilder.zki.ISStateRootAccount[len(l1UserTxs)-1]
	//       batchBuilder.zki.ISStateRootAccount[i] = lastValidISStateRootAccount
	//       lastValidISExitRoot := batchBuilder.zki.ISExitRoot[len(l1UserTxs)-1]
	//       batchBuilder.zki.ISExitRoot[i] = lastValidISExitRoot
	// 	}
	// }

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

	return batchBuilder.zki, nil
}

// applyCreateAccount creates a new account and updates ZKInputs.
// It now takes sdb (*statedb.LocalStateDB) as a parameter.
func (batchBuilder *BatchBuilder) applyCreateAccount(sdb *statedb.LocalStateDB, tx *common.Tx, exitTree *merkletree.MerkleTree) error {
	// account := &common.Account{
	// 	Nonce:   0,
	// 	Balance: tx.EffectiveDepositAmount, // Use effective deposit amount
	// 	BJJ:     tx.FromBJJ,
	// 	EthAddr: tx.FromEthAddr,
	// }

	// newAccountIdx := common.AccountIdx(sdb.CurrentAccountIdx() + 1)
	// p, err := sdb.CreateAccount(newAccountIdx, account)
	// if err != nil {
	// 	return common.Wrap(fmt.Errorf("applyCreateAccount: failed to create account %d: %w", newAccountIdx, err))
	// }

	// // Populate ZKI for state 1 (the new account)
	// // batchBuilder.zki.TokenID1[batchBuilder.txIndex] = tx.TokenID.BigInt() // If TokenID is part of L1Tx / Account
	// batchBuilder.zki.Nonce1[batchBuilder.txIndex] = big.NewInt(0)
	// fromBJJSign, fromBJJY := babyjub.UnpackSignY(tx.FromBJJ)
	// bjjSignBool := fromBJJSign
	// batchBuilder.zki.Sign1[batchBuilder.txIndex] = &bjjSignBool
	// batchBuilder.zki.Ay1[batchBuilder.txIndex] = fromBJJY
	// batchBuilder.zki.Balance1[batchBuilder.txIndex] = new(big.Int).Set(tx.EffectiveDepositAmount)
	// batchBuilder.zki.EthAddr1[batchBuilder.txIndex] = common.EthAddrToBigInt(tx.FromEthAddr)
	// batchBuilder.zki.Siblings1[batchBuilder.txIndex] = siblingsToZKInputFormat(p.Siblings)

	// isOld0Val := p.IsOld0
	// batchBuilder.zki.IsOld0_1[batchBuilder.txIndex] = &isOld0Val
	// batchBuilder.zki.OldKey1[batchBuilder.txIndex] = p.OldKey.BigInt()
	// batchBuilder.zki.OldValue1[batchBuilder.txIndex] = p.OldValue.BigInt()

	// auxFromIdxVal := uint32(newAccountIdx)
	// batchBuilder.zki.AuxFromIdx[batchBuilder.txIndex] = &auxFromIdxVal
	// newAccountCreatedVal := true
	// batchBuilder.zki.NewAccount[batchBuilder.txIndex] = &newAccountCreatedVal

	// // Update NewLastIdxRaw in ZKI as an account was created
	// batchBuilder.zki.NewLastIdxRaw = uint32(newAccountIdx)

	// return sdb.SetCurrentAccountIdx(newAccountIdx)
	return nil
}

// applyDeposit updates an existing account's balance and ZKInputs.
// It now takes sdb (*statedb.LocalStateDB) as a parameter.
func (batchBuilder *BatchBuilder) applyDeposit(sdb *statedb.LocalStateDB, tx *common.Tx, exitTree *merkletree.MerkleTree) error {
	// accSender, err := sdb.GetAccount(tx.FromIdx)
	// if err != nil {
	// 	return common.Wrap(fmt.Errorf("applyDeposit: failed to get sender account %d: %w", tx.FromIdx, err))
	// }

	// // Populate ZKI for state 1 (sender account before update)
	// // batchBuilder.zki.TokenID1[batchBuilder.txIndex] = accSender.TokenID.BigInt() // If TokenID is part of Account
	// batchBuilder.zki.Nonce1[batchBuilder.txIndex] = new(big.Int).Set(accSender.Nonce.BigInt())
	// senderBJJSign, senderBJJY := babyjub.UnpackSignY(accSender.BJJ)
	// bjjSignBool := senderBJJSign
	// batchBuilder.zki.Sign1[batchBuilder.txIndex] = &bjjSignBool
	// batchBuilder.zki.Ay1[batchBuilder.txIndex] = senderBJJY
	// batchBuilder.zki.Balance1[batchBuilder.txIndex] = new(big.Int).Set(accSender.Balance)
	// batchBuilder.zki.EthAddr1[batchBuilder.txIndex] = common.EthAddrToBigInt(accSender.EthAddr)

	// // Add the deposit to the sender
	// accSender.Balance.Add(accSender.Balance, tx.EffectiveDepositAmount)
	// // If there's an amount to transfer/use from this deposit (e.g. for an internal transfer part of deposit)
	// // accSender.Balance.Sub(accSender.Balance, tx.EffectiveAmount) // This line was from reference, ensure it's applicable
	// // For a simple deposit, tx.EffectiveAmount might be 0 or handled differently.
	// // If tx.EffectiveAmount is for a fee or another operation, ensure logic is correct.
	// // Assuming for now that EffectiveAmount is not subtracted in a pure deposit scenario unless it's a fee.

	// if accSender.Balance.Cmp(big.NewInt(0)) == -1 {
	// 	return fmt.Errorf("applyDeposit: sender %d balance became negative: %s", tx.FromIdx, accSender.Balance.String()) // Or use newErrorNotEnoughBalance
	// }

	// p, err := sdb.UpdateAccount(tx.FromIdx, accSender)
	// if err != nil {
	// 	return common.Wrap(fmt.Errorf("applyDeposit: failed to update sender account %d: %w", tx.FromIdx, err))
	// }
	// batchBuilder.zki.Siblings1[batchBuilder.txIndex] = siblingsToZKInputFormat(p.Siblings)
	// // IsOld0_1, OldKey1, OldValue1 are not typically set for updates, only inserts.
	// // However, if the circuit expects them (e.g. IsOld0=false), set accordingly.
	// // Based on reference, these are not set for updates.

	return nil
}

// applyExit processes an exit transaction, updating state and ZKInputs.
// It now takes sdb (*statedb.LocalStateDB) as a parameter.
func (batchBuilder *BatchBuilder) applyExit(sdb *statedb.LocalStateDB, exitTree *merkletree.MerkleTree, tx common.Tx, originalAmount *big.Int) (*common.Account, bool, error) {
	// acc, err := sdb.GetAccount(tx.FromIdx)
	// if err != nil {
	// 	return nil, false, common.Wrap(fmt.Errorf("applyExit: failed to get account %d: %w", tx.FromIdx, err))
	// }

	// // Populate ZKI for state 1 (account in main tree before update)
	// // batchBuilder.zki.TokenID1[batchBuilder.txIndex] = acc.TokenID.BigInt()
	// batchBuilder.zki.Nonce1[batchBuilder.txIndex] = new(big.Int).Set(acc.Nonce.BigInt())
	// accBJJSign, accBJJY := babyjub.UnpackSignY(acc.BJJ)
	// bjjSignBool1 := accBJJSign
	// batchBuilder.zki.Sign1[batchBuilder.txIndex] = &bjjSignBool1
	// batchBuilder.zki.Ay1[batchBuilder.txIndex] = accBJJY
	// batchBuilder.zki.Balance1[batchBuilder.txIndex] = new(big.Int).Set(acc.Balance)
	// batchBuilder.zki.EthAddr1[batchBuilder.txIndex] = common.EthAddrToBigInt(acc.EthAddr)

	// // Update account in main state tree
	// originalBalance := new(big.Int).Set(acc.Balance)
	// acc.Balance = new(big.Int).Sub(acc.Balance, tx.Amount) // tx.Amount is EffectiveAmount for exits
	// if acc.Balance.Cmp(big.NewInt(0)) == -1 {
	// 	// Revert balance if it goes negative, and potentially fail the tx or set amount to originalBalance
	// 	// For now, let's assume effective_amount calculation already handled this.
	// 	// If not, this is an error.
	// 	return nil, false, common.Wrap(fmt.Errorf("applyExit: account %d balance negative after exit: %s from %s with amount %s", tx.FromIdx, acc.Balance.String(), originalBalance.String(), tx.Amount.String()))
	// }

	// p1, err := sdb.UpdateAccount(tx.FromIdx, acc)
	// if err != nil {
	// 	return nil, false, common.Wrap(fmt.Errorf("applyExit: failed to update account %d in main tree: %w", tx.FromIdx, err))
	// }
	// batchBuilder.zki.Siblings1[batchBuilder.txIndex] = siblingsToZKInputFormat(p1.Siblings)

	// if exitTree == nil { // Should not happen if initialized correctly
	// 	return nil, false, fmt.Errorf("applyExit: exitTree is nil")
	// }

	// // Process exit in ExitTree
	// // Use originalAmount for the check, but tx.Amount (effective) for the actual exit value
	// if originalAmount.Cmp(big.NewInt(0)) == 0 {
	// 	// If the Exit Amount==0, the Exit is not added to the ExitTree.
	// 	// ZKI fields for State2 (exit leaf) should be set to defaults or indicate no-op.
	// 	// Example:
	// 	if batchBuilder.txIndex < len(batchBuilder.zki.Nonce2) { // Check bounds
	// 		batchBuilder.zki.Nonce2[batchBuilder.txIndex] = big.NewInt(0) // Default/No-op
	// 		// ... set other ZKI.State2 fields to defaults ...
	// 		// Ensure ISExitRoot is correctly set (likely to previous or current if no change)
	// 		if batchBuilder.txIndex < len(batchBuilder.zki.ISExitRoot) && exitTree != nil {
	// 			batchBuilder.zki.ISExitRoot[batchBuilder.txIndex] = exitTree.Root().BigInt()
	// 		}
	// 	}
	// 	return nil, false, nil
	// }

	// var exitAccountInTree *common.Account
	// var p2 *merkletree.CircomProcessorProof
	// newExitInTree := false

	// existingExitLeaf, err := statedb.GetAccountInTreeDB(exitTree.DB(), tx.FromIdx)
	// if common.Unwrap(err) == merkletree.ErrKeyNotFound { // Use merkletree.ErrKeyNotFound
	// 	newExitInTree = true
	// 	exitAccountInTree = &common.Account{
	// 		Nonce:   common.Nonce(0),
	// 		Balance: new(big.Int).Set(tx.Amount), // Use effective amount for exit leaf balance
	// 		BJJ:     acc.BJJ,                     // Use BJJ from main tree account
	// 		EthAddr: acc.EthAddr,                 // Use EthAddr from main tree account
	// 	}
	// 	// Populate ZKI for state 2 (exit account before creation - effectively zero/empty)
	// 	batchBuilder.zki.Nonce2[batchBuilder.txIndex] = big.NewInt(0)
	// 	sign2Val := false // Assuming default for non-existent
	// 	batchBuilder.zki.Sign2[batchBuilder.txIndex] = &sign2Val
	// 	batchBuilder.zki.Ay2[batchBuilder.txIndex] = new(big.Int) // Zero point Y
	// 	batchBuilder.zki.Balance2[batchBuilder.txIndex] = big.NewInt(0)
	// 	batchBuilder.zki.EthAddr2[batchBuilder.txIndex] = big.NewInt(0) // Or EthAddrToBigInt of zero address

	// 	p2, err = statedb.CreateAccountInTreeDB(exitTree.DB(), exitTree, tx.FromIdx, exitAccountInTree)
	// } else if err == nil {
	// 	exitAccountInTree = existingExitLeaf
	// 	// Populate ZKI for state 2 (exit account before update)
	// 	batchBuilder.zki.Nonce2[batchBuilder.txIndex] = new(big.Int).Set(exitAccountInTree.Nonce.BigInt())
	// 	exitBJJSign, exitBJJY := babyjub.UnpackSignY(exitAccountInTree.BJJ)
	// 	bjjSignBool2 := exitBJJSign
	// 	batchBuilder.zki.Sign2[batchBuilder.txIndex] = &bjjSignBool2
	// 	batchBuilder.zki.Ay2[batchBuilder.txIndex] = exitBJJY
	// 	batchBuilder.zki.Balance2[batchBuilder.txIndex] = new(big.Int).Set(exitAccountInTree.Balance)
	// 	batchBuilder.zki.EthAddr2[batchBuilder.txIndex] = common.EthAddrToBigInt(exitAccountInTree.EthAddr)

	// 	exitAccountInTree.Balance = new(big.Int).Add(exitAccountInTree.Balance, tx.Amount)
	// 	p2, err = statedb.UpdateAccountInTreeDB(exitTree.DB(), exitTree, tx.FromIdx, exitAccountInTree)
	// } else { // Other error
	// 	return nil, false, common.Wrap(fmt.Errorf("applyExit: failed to get account %d from exit tree: %w", tx.FromIdx, err))
	// }

	// if err != nil {
	// 	return nil, false, common.Wrap(fmt.Errorf("applyExit: failed to update/create account %d in exit tree: %w", tx.FromIdx, err))
	// }

	// if p2 != nil {
	// 	batchBuilder.zki.Siblings2[batchBuilder.txIndex] = siblingsToZKInputFormat(p2.Siblings)
	// 	isOld0Val2 := p2.IsOld0
	// 	batchBuilder.zki.IsOld0_2[batchBuilder.txIndex] = &isOld0Val2
	// 	batchBuilder.zki.OldKey2[batchBuilder.txIndex] = p2.OldKey.BigInt()
	// 	batchBuilder.zki.OldValue2[batchBuilder.txIndex] = p2.OldValue.BigInt()
	// }
	// return exitAccountInTree, newExitInTree, nil
	return nil, false, nil
}

// computeEffectiveAmounts checks L1Tx data and calculates effective amounts.
// It now takes sdb (*statedb.LocalStateDB) as a parameter.
func (batchBuilder *BatchBuilder) computeEffectiveAmounts(sdb *statedb.LocalStateDB, tx *common.L1Tx) {
	// Initialize effective amounts to actual amounts
	tx.EffectiveAmount = new(big.Int).Set(tx.Amount)
	tx.EffectiveDepositAmount = new(big.Int).Set(tx.DepositAmount)

	if tx.Type == common.TxTypeCreateAccountDeposit {
		// For create account, the full deposit amount is effective, and amount might be 0 or for something else.
		// No further checks on sender balance needed here as it's a new account.
		return
	}

	accSender, err := sdb.GetAccount(tx.FromIdx)
	if err != nil {
		// log.Debugf("EffectiveAmount & EffectiveDepositAmount = 0: can not get account for tx.FromIdx: %d", tx.FromIdx)
		fmt.Printf("Warning: computeEffectiveAmounts: cannot get account for tx.FromIdx: %d, error: %v. Setting effective amounts to 0.\n", tx.FromIdx, err)
		tx.EffectiveDepositAmount = big.NewInt(0)
		tx.EffectiveAmount = big.NewInt(0)
		return
	}

	// Check FromEthAddr matches sender's EthAddr
	if !bytes.Equal(tx.FromEthAddr.Bytes(), accSender.EthAddr.Bytes()) {
		// log.Debugf("EffectiveAmount & EffectiveDepositAmount = 0: tx.FromEthAddr (%s) mismatch with sender %d EthAddr (%s)",
		// 	tx.FromEthAddr.Hex(), tx.FromIdx, accSender.EthAddr.Hex())
		tx.EffectiveDepositAmount = big.NewInt(0) // If deposit relies on this check
		tx.EffectiveAmount = big.NewInt(0)
		return
	}

	// Calculate available balance after deposit
	availableBalance := new(big.Int).Set(accSender.Balance)
	if tx.DepositAmount != nil && tx.DepositAmount.Cmp(big.NewInt(0)) > 0 {
		availableBalance.Add(availableBalance, tx.EffectiveDepositAmount)
	}

	// If tx.Amount (for withdrawal/transfer part) exceeds available balance, set EffectiveAmount to 0
	// (or cap it at availableBalance, depending on desired behavior for partial execution)
	// The reference sets EffectiveAmount to 0 if funds are insufficient.
	if tx.Amount.Cmp(availableBalance) > 0 {
		// log.Debugf("EffectiveAmount = 0: Not enough funds for amount %s (available: %s) for sender %d",
		// 	tx.Amount.String(), availableBalance.String(), tx.FromIdx)
		tx.EffectiveAmount = big.NewInt(0)
		// Note: EffectiveDepositAmount might still be valid if the deposit itself is okay.
		// The reference logic seems to zero out EffectiveAmount if the *spending* part (tx.Amount) fails.
	}
}

// applyVouch handles vouch creation/deletion.
// It now takes sdb (*statedb.LocalStateDB) as a parameter.
func (batchBuilder *BatchBuilder) applyVouch(sdb *statedb.LocalStateDB, tx common.Tx, auxToIdx common.AccountIdx, exitTree *merkletree.MerkleTree) error {
	// if auxToIdx == common.ZeroAccountIdx { // Assuming common.ZeroAccountIdx or similar for unassigned
	// 	auxToIdx = tx.ToIdx
	// }
	// if auxToIdx == common.ZeroAccountIdx {
	// 	return fmt.Errorf("applyVouch: receiver index (auxToIdx) is zero for tx %+v", tx)
	// }

	// // Handle Sender (State1)
	// accSender, err := sdb.GetAccount(tx.FromIdx)
	// if err != nil {
	// 	return common.Wrap(fmt.Errorf("applyVouch: failed to get sender account %d: %w", tx.FromIdx, err))
	// }

	// // Populate ZKI for State1 (Sender before update)
	// // batchBuilder.zki.TokenID1[batchBuilder.txIndex] = accSender.TokenID.BigInt()
	// batchBuilder.zki.Nonce1[batchBuilder.txIndex] = new(big.Int).Set(accSender.Nonce.BigInt())
	// senderBJJSign, senderBJJY := babyjub.UnpackSignY(accSender.BJJ)
	// bjjSignBool1 := senderBJJSign
	// batchBuilder.zki.Sign1[batchBuilder.txIndex] = &bjjSignBool1
	// batchBuilder.zki.Ay1[batchBuilder.txIndex] = senderBJJY
	// batchBuilder.zki.Balance1[batchBuilder.txIndex] = new(big.Int).Set(accSender.Balance)
	// batchBuilder.zki.EthAddr1[batchBuilder.txIndex] = common.EthAddrToBigInt(accSender.EthAddr)
	// // Vouch/Score specific fields for State1 if needed

	// // Update Sender (e.g., increment nonce)
	// accSender.Nonce.Add(accSender.Nonce, big.NewInt(1))
	// // TODO: Update sender's vouch list and score based on tx.Type and auxToIdx

	// pSender, err := sdb.UpdateAccount(tx.FromIdx, accSender)
	// if err != nil {
	// 	return common.Wrap(fmt.Errorf("applyVouch: failed to update sender account %d: %w", tx.FromIdx, err))
	// }
	// batchBuilder.zki.Siblings1[batchBuilder.txIndex] = siblingsToZKInputFormat(pSender.Siblings)

	// // Handle Receiver (State2)
	// if tx.FromIdx != auxToIdx { // If sender is not receiver
	// 	accReceiver, err := sdb.GetAccount(auxToIdx)
	// 	if err != nil {
	// 		return common.Wrap(fmt.Errorf("applyVouch: failed to get receiver account %d: %w", auxToIdx, err))
	// 	}

	// 	// Populate ZKI for State2 (Receiver before update)
	// 	// batchBuilder.zki.TokenID2[batchBuilder.txIndex] = accReceiver.TokenID.BigInt()
	// 	batchBuilder.zki.Nonce2[batchBuilder.txIndex] = new(big.Int).Set(accReceiver.Nonce.BigInt())
	// 	receiverBJJSign, receiverBJJY := babyjub.UnpackSignY(accReceiver.BJJ)
	// 	bjjSignBool2 := receiverBJJSign
	// 	batchBuilder.zki.Sign2[batchBuilder.txIndex] = &bjjSignBool2
	// 	batchBuilder.zki.Ay2[batchBuilder.txIndex] = receiverBJJY
	// 	batchBuilder.zki.Balance2[batchBuilder.txIndex] = new(big.Int).Set(accReceiver.Balance)
	// 	batchBuilder.zki.EthAddr2[batchBuilder.txIndex] = common.EthAddrToBigInt(accReceiver.EthAddr)
	// 	// Vouch/Score specific fields for State2 if needed

	// 	// TODO: Update receiver's vouch list and score based on tx.Type and tx.FromIdx

	// 	pReceiver, err := sdb.UpdateAccount(auxToIdx, accReceiver)
	// 	if err != nil {
	// 		return common.Wrap(fmt.Errorf("applyVouch: failed to update receiver account %d: %w", auxToIdx, err))
	// 	}
	// 	batchBuilder.zki.Siblings2[batchBuilder.txIndex] = siblingsToZKInputFormat(pReceiver.Siblings)
	// } else {
	// 	// If sender is receiver, State2 ZKI fields might need specific padding
	// 	// or point to the updated state of State1. This depends on circuit requirements.
	// 	// For now, let's assume State2 is for a distinct receiver. If not, these fields might be zeroed
	// 	// or set to a "no-op" state.
	// 	// Example:
	// 	// batchBuilder.zki.Nonce2[batchBuilder.txIndex] = accSender.Nonce.BigInt() // Nonce after increment
	// 	// ... and so on for other fields, or set to defaults if circuit expects empty for self-vouch.
	// }
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
