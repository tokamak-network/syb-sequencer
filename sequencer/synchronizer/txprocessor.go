package synchronizer

import (
	"fmt"
	"math/big"

	ethCommon "github.com/ethereum/go-ethereum/common"
	"github.com/tokamak-network/syb-sequencer/sequencer/common"
	"github.com/tokamak-network/syb-sequencer/sequencer/db/statedb"
)

// TxProcessor represents the TxProcessor object
type TxProcessor struct {
	state *statedb.StateDB
	// updatedAccounts stores the last version of the account when it has
	// been created/updated by any of the processed transactions.
	updatedAccounts map[common.AccountIdx]*common.Account
}

// ProcessTxOutput contains the output of the ProcessTxs method
type ProcessTxOutput struct {
	CreatedAccount common.Account
	// UpdatedAccounts returns the current state of each account
	// created/updated by any of the processed transactions.
	UpdatedAccounts map[common.AccountIdx]*common.Account
}

// NewTxProcessor returns a new TxProcessor with the given *StateDB & Config
func NewTxProcessor(state *statedb.StateDB) *TxProcessor {
	return &TxProcessor{
		state: state,
	}
}

func (txProcessor *TxProcessor) ProcessTxs(tx common.Tx) (ptOut *ProcessTxOutput, err error) {
	defer func() {
		if err == nil {
			err = txProcessor.state.MakeCheckpoint()
		}
	}()

	txProcessor.updatedAccounts = make(map[common.AccountIdx]*common.Account)

	switch tx.Type {
	case common.TxTypeCreateAccountDeposit:
		fmt.Printf("Creating account... EthAddr: %s \n", tx.FromEthAddr)
		// Create Account
		account := &common.Account{
			Idx:     tx.FromIdx,
			EthAddr: ethCommon.BytesToAddress(tx.FromEthAddr),
			Balance: tx.Amount,
		}

		txProcessor.updatedAccounts[tx.FromIdx] = account
		_, err := txProcessor.state.CreateAccount(tx.FromIdx, account)
		if err != nil {
			return nil, common.Wrap(err)
		}

		// Create Score for newly created account
		score := &common.Score{
			Idx:     common.ScoreIdx(tx.FromIdx),
			EthAddr: ethCommon.BytesToAddress(tx.FromEthAddr),
			Score:   big.NewInt(0),
		}
		_, err = txProcessor.state.CreateScore(score.Idx, score)
		if err != nil {
			return nil, common.Wrap(err)
		}

		txProcessor.state.SetCurrentAccountIdx(tx.FromIdx)

		fmt.Printf("Created account %s succssfully \n", tx.FromIdx)
	case common.TxTypeDeposit:
		accSender, err := txProcessor.state.GetAccount(tx.FromIdx)
		if err != nil {
			return nil, common.Wrap(err)
		}
		// add the deposit to the sender
		accSender.Balance = new(big.Int).Add(accSender.Balance, tx.Amount)

		_, err = txProcessor.state.UpdateAccount(tx.FromIdx, accSender)
		if err != nil {
			return nil, common.Wrap(err)
		}
	case common.TxTypeVouch:
		// Create vouchIdx from from & to Idx
		vouchIdx, err := common.VouchIdxFromAccountIdxs(tx.FromIdx, tx.ToIdx)
		if err != nil {
			return nil, common.Wrap(err)
		}

		// Create vouch in stateDB
		vouch := common.Vouch{
			Idx:         vouchIdx,
			FromIdx:     tx.FromIdx,
			FromEthAddr: ethCommon.BytesToAddress(tx.FromEthAddr),
			ToIdx:       tx.ToIdx,
			ToEthAddr:   ethCommon.BytesToAddress(tx.ToEthAddr),
		}

		_, err = txProcessor.state.Vouch(vouchIdx, &vouch)
		if err != nil {
			return nil, common.Wrap(err)
		}
	case common.TxTypeUnvouch:
		// Create vouchIdx from from & to Idx
		vouchIdx, err := common.VouchIdxFromAccountIdxs(tx.FromIdx, tx.ToIdx)
		if err != nil {
			return nil, common.Wrap(err)
		}

		_, err = txProcessor.state.UnVouch(vouchIdx)
		if err != nil {
			return nil, common.Wrap(err)
		}
	case common.TxTypeWithdraw:
		accSender, err := txProcessor.state.GetAccount(tx.FromIdx)
		if err != nil {
			return nil, common.Wrap(err)
		}
		// add the deposit to the sender
		accSender.Balance = new(big.Int).Sub(accSender.Balance, tx.Amount)
		if accSender.Balance.Cmp(big.NewInt(0)) == -1 { // balance<0
			return nil, fmt.Errorf("invalid transaction, not enough balance on account. "+
				"TxID: %d, TxType: %s, FromIdx: %d, ToIdx: %d, Amount: %d",
				tx.ItemID, tx.Type, tx.FromIdx, tx.ToIdx, tx.Amount)
		}

		_, err = txProcessor.state.UpdateAccount(tx.FromIdx, accSender)
		if err != nil {
			return nil, common.Wrap(err)
		}
	case common.TxTypeExplode:
		// get sender account from StateDB
		accSender, err := txProcessor.state.GetAccount(tx.FromIdx)
		if err != nil {
			return nil, common.Wrap(err)
		}

		accSender.Balance = new(big.Int).Add(accSender.Balance, tx.Amount)

		// update sender account in StateDB
		_, err = txProcessor.state.UpdateAccount(tx.FromIdx, accSender)
		if err != nil {
			return nil, common.Wrap(err)
		}

		// get sender account from StateDB
		accReceiver, err := txProcessor.state.GetAccount(tx.ToIdx)
		if err != nil {
			return nil, common.Wrap(err)
		}

		accReceiver.Balance = new(big.Int).Add(accReceiver.Balance, tx.Amount)
		if accReceiver.Balance.Cmp(big.NewInt(0)) == -1 { // balance<0
			return nil, fmt.Errorf("invalid transaction, not enough balance on account. "+
				"TxID: %d, TxType: %s, FromIdx: %d, ToIdx: %d, Amount: %d",
				tx.ItemID, tx.Type, tx.FromIdx, tx.ToIdx, tx.Amount)
		}

		// update sender account in StateDB
		_, err = txProcessor.state.UpdateAccount(tx.ToIdx, accReceiver)
		if err != nil {
			return nil, common.Wrap(err)
		}
	}

	createdAccount, err := txProcessor.state.GetAccount(txProcessor.state.CurrentAccountIdx())
	if err != nil {
		return nil, common.Wrap(err)
	}

	return &ProcessTxOutput{
		CreatedAccount:  *createdAccount,
		UpdatedAccounts: txProcessor.updatedAccounts,
	}, nil
}
