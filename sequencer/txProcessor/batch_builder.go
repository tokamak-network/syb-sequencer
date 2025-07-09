package txprocessor

import (
	"fmt"
	"math/big"
	"os"

	ethCommon "github.com/ethereum/go-ethereum/common"
	"github.com/iden3/go-merkletree"
	"github.com/iden3/go-merkletree/db/pebble"
	"github.com/tokamak-network/syb-sequencer/sequencer/common"
	"github.com/tokamak-network/syb-sequencer/sequencer/db/historydb"
	"github.com/tokamak-network/syb-sequencer/sequencer/db/statedb"
)

type Config struct {
	NLevels uint32
	MaxTx   uint32
	MaxL1Tx uint32
	ChainID uint64
}
type BatchBuilder struct {
	statedb   *statedb.LocalStateDB
	config    Config
	historydb *historydb.HistoryDB
	zki       *common.ZKInputs
	txIndex   int
}

func NewBatchBuilder(config Config, statedb *statedb.LocalStateDB, historydb *historydb.HistoryDB) *BatchBuilder {
	return &BatchBuilder{
		statedb:   statedb,
		config:    config,
		historydb: historydb,
		zki:       nil,
		txIndex:   0,
	}
}

// Resets the zkInputs and transaction index
func (batchBuilder *BatchBuilder) resetZKInputs() {
	batchBuilder.zki = nil
	batchBuilder.txIndex = 0
}

// forgeTransactions processes L1 user transactions, updates the state via forger.statedb, and generates ZKInputs.
func (batchBuilder *BatchBuilder) ForgeTransactions(l1UserTxs []*common.Tx) (*common.ZKInputs, error) {
	batchBuilder.resetZKInputs()

	// Access StateDB via the forger
	sdb := batchBuilder.statedb
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

	for i, tx := range l1UserTxs {
		batchBuilder.txIndex = i // Set current transaction index for ZKI population

		currentTx := tx

		switch currentTx.Type {
		case common.TxTypeCreateAccountDeposit:
			err = batchBuilder.applyCreateAccount(sdb, currentTx)
		case common.TxTypeDeposit, common.TxTypeWithdraw:
			err = batchBuilder.applyDepositWithdrawal(sdb, currentTx)
		case common.TxTypeVouch, common.TxTypeUnvouch:
			err = batchBuilder.applyVouch(sdb, *currentTx)
		default:
			err = fmt.Errorf("unknown L1 transaction type: %s for txID: %s", currentTx.Type, currentTx.ItemID)

			if err != nil {
				return nil, common.Wrap(fmt.Errorf("failed to process tx %s (type %s): %w", currentTx.ItemID, currentTx.Type, err))
			}
		}
	}
	// Update score after processing all transactions
	batchBuilder.UpdateScore()

	// Final ZKI parameters
	globalChainIDVal := uint16(batchBuilder.config.ChainID)
	batchBuilder.zki.GlobalChainID = &globalChainIDVal
	batchBuilder.zki.NewAccountRootRaw = sdb.GetATRootHash()
	batchBuilder.zki.NewVouchRootRaw = sdb.GetVTRootHash()
	batchBuilder.zki.NewScoreRootRaw = sdb.GetSTRootScore()
	if exitTree != nil {
		batchBuilder.zki.NewExitRootRaw = exitTree.Root()
	}

	// Make a checkpoint in the StateDB after processing all transactions for this batch
	if err := sdb.MakeCheckpoint(); err != nil {
		return nil, common.Wrap(fmt.Errorf("failed to make checkpoint in StateDB: %w", err))
	}
	return batchBuilder.zki, nil
}

// TODO: Update zki inputs as per the requirements in all the apply functions
// applyCreateAccount creates a new account and updates ZKInputs.
// It now takes sdb (*statedb.LocalStateDB) as a parameter.
func (batchBuilder *BatchBuilder) applyCreateAccount(sdb *statedb.LocalStateDB, tx *common.Tx) error {
	account := &common.Account{
		Idx:     tx.FromIdx,
		Balance: tx.Amount,
		EthAddr: ethCommon.BytesToAddress(tx.FromEthAddr),
	}

	_, err := sdb.CreateAccount(tx.FromIdx, account)
	if err != nil {
		return common.Wrap(fmt.Errorf("applyCreateAccount: failed to create account %d: %w", tx.FromIdx, err))
	}

	// Create Score for newly created account
	score := &common.Score{
		Idx:     common.ScoreIdx(tx.FromIdx),
		EthAddr: ethCommon.BytesToAddress(tx.FromEthAddr),
		Score:   big.NewInt(0),
	}
	_, err = sdb.CreateScore(score.Idx, score)
	if err != nil {
		return common.Wrap(err)
	}

	err = sdb.SetCurrentAccountIdx(tx.FromIdx)
	if err != nil {
		return common.Wrap(fmt.Errorf("applyCreateAccount: failed to update current account index to %d: %w", tx.FromIdx, err))
	}
	return nil
}

// applyDeposit updates an existing account's balance and ZKInputs.
// It now takes sdb (*statedb.LocalStateDB) as a parameter.
func (batchBuilder *BatchBuilder) applyDepositWithdrawal(sdb *statedb.LocalStateDB, tx *common.Tx) error {
	account, err := sdb.GetAccount(tx.FromIdx)
	if err != nil {
		return common.Wrap(fmt.Errorf("applyDeposit: failed to get sender account %d: %w", tx.FromIdx, err))
	}

	// Add the deposit to the sender
	if tx.Type == common.TxTypeDeposit {
		account.Balance.Add(account.Balance, tx.Amount)
	} else if tx.Type == common.TxTypeWithdraw {
		if account.Balance.Cmp(tx.Amount) < 0 {
			return fmt.Errorf("WithdrawalTx: insufficient balance for account idx %d. Has: %s, Wants: %s",
				account.Idx, account.Balance.String(), tx.Amount.String())
		}
		account.Balance.Sub(account.Balance, tx.Amount)
	}
	_, err = sdb.UpdateAccount(tx.FromIdx, account)
	if err != nil {
		return common.Wrap(fmt.Errorf("applyDeposit: failed to update sender account %d: %w", tx.FromIdx, err))
	}

	return nil
}

// applyVouch handles vouch creation/deletion.
// It now takes sdb (*statedb.LocalStateDB) as a parameter.
// TODO: Remove the part updating the score in the vouch function and add this at the end of the batch processing when all the transactions are forged via score calculation logic.
func (batchBuilder *BatchBuilder) applyVouch(sdb *statedb.LocalStateDB, tx common.Tx) error {
	fromAccountIdx := tx.FromIdx
	toAccountIdx := tx.ToIdx
	fromEthAddr := ethCommon.Address(tx.FromEthAddr)
	toEthAddr := ethCommon.Address(tx.ToEthAddr)
	vouchTableKeyStr := fmt.Sprintf("%d%d", tx.FromIdx, tx.ToIdx)
	vouchTableKeyBigInt, ok := new(big.Int).SetString(vouchTableKeyStr, 10)
	if !ok {
		return fmt.Errorf("VouchTx: failed to create vouch table key from string '%s'", vouchTableKeyStr)
	}
	var err error

	vouchIdx := common.VouchIdx(vouchTableKeyBigInt.Uint64())
	if err != nil {
		return common.Wrap(fmt.Errorf("applyVouch: failed to get score for account idx %d: %w", toAccountIdx, err))
	}

	switch tx.Type {
	case common.TxTypeVouch:
		vouchDetails := &common.Vouch{Idx: vouchIdx, FromEthAddr: fromEthAddr, ToEthAddr: toEthAddr, FromIdx: fromAccountIdx, ToIdx: toAccountIdx}
		_, err = sdb.Vouch(common.VouchIdx(fromAccountIdx), vouchDetails)
		if err != nil {
			return common.Wrap(fmt.Errorf("applyVouch: failed to create vouch for VouchIdx %s: %w", common.VouchIdx(fromAccountIdx).String(), err))
		}
	case common.TxTypeUnvouch:
		_, err = sdb.UnVouch(vouchIdx)
		if err != nil {
			return common.Wrap(fmt.Errorf("applyVouch: failed to delete vouch for VouchIdx %s: %w", common.VouchIdx(fromAccountIdx).String(), err))
		}
	default:
		return fmt.Errorf("applyVouch: unsupported txType for vouch operation: %s", tx.Type)
	}
	return nil
}

func (bb *BatchBuilder) UpdateScore() {
	accounts, totalAccountNumber, err := bb.historydb.GetAllAccounts()
	if err != nil {
		fmt.Printf("Error fetching accounts: %v\n", err)
		return
	}

	if totalAccountNumber == 0 {
		fmt.Printf("No accounts found to update scores\n")
		return
	}

	balances := make([]*big.Int, totalAccountNumber)
	scores := make([]*big.Int, totalAccountNumber)
	vouches := make([][]int, totalAccountNumber)

	// Initialize vouches matrix with proper dimensions
	for i := range vouches {
		vouches[i] = make([]int, totalAccountNumber)
	}

	// Initialize balances and scores with zero values
	for i := int64(0); i < totalAccountNumber; i++ {
		balances[i] = big.NewInt(0)
		scores[i] = big.NewInt(0)
	}

	// Populate data from accounts
	for _, account := range accounts {
		if int64(account.Idx) >= totalAccountNumber {
			fmt.Printf("Warning: Account index %d exceeds total account number %d\n", account.Idx, totalAccountNumber)
			continue
		}

		balances[account.Idx] = account.Balance

		score, err := bb.statedb.GetScore(common.ScoreIdx(account.Idx))
		if err != nil {
			fmt.Printf("Error fetching score for account %d: %v\n", account.Idx, err)
			continue
		}
		scores[account.Idx] = score.Score

		// Build vouches matrix
		for _, vouchedAccount := range accounts {
			if int64(vouchedAccount.Idx) >= totalAccountNumber {
				continue
			}

			vouchIdx, err := common.VouchIdxFromAccountIdxs(account.Idx, vouchedAccount.Idx)
			if err != nil {
				fmt.Printf("Error creating vouch index for accounts %d->%d: %v\n", account.Idx, vouchedAccount.Idx, err)
				vouches[account.Idx][vouchedAccount.Idx] = 0
				continue
			}

			vouch, err := bb.statedb.GetVouch(vouchIdx)
			if err != nil {
				fmt.Printf("Error fetching vouch for accounts %d->%d: %v\n", account.Idx, vouchedAccount.Idx, err)
				vouches[account.Idx][vouchedAccount.Idx] = 0
				continue
			}

			if vouch != nil {
				vouches[account.Idx][vouchedAccount.Idx] = 1
			} else {
				vouches[account.Idx][vouchedAccount.Idx] = 0
			}
		}
	}

	// Calculate new scores using the scoring algorithm
	newScores := common.CalculateScore(vouches, balances, scores)
	for _, score := range newScores {
		fmt.Printf("New Score: %s\n", score.String())
	}
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
