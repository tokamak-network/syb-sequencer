package synchronizer

import (
	"fmt"
	"math/big"

	ethCommon "github.com/ethereum/go-ethereum/common"
	"github.com/tokamak-network/syb-sequencer/sequencer/common"
)

func (s *Synchronizer) AddTransactionToHistoryDB(tx *common.Tx) error {
	if tx == nil {
		return fmt.Errorf("AddTransactionToHistoryDB: transaction is nil")
	}
	err := s.historydb.SaveTx(tx)
	if err != nil {
		s.logger.Printf("Error saving transaction: %v", err)
		return err
	}

	switch tx.Type {
	case common.TxTypeCreateAccountDeposit:
		return s.AddAccountTx(tx)
	case common.TxTypeDeposit:
		return s.DepositWithdrawTx(tx)
	case common.TxTypeWithdraw:
		return s.DepositWithdrawTx(tx)
	case common.TxTypeVouch:
		return s.VouchTx(tx)
	case common.TxTypeUnvouch:
		return s.VouchTx(tx)
	case common.TxTypeForgeBatch, common.TxTypeExplode:
		return nil
	default:
		return fmt.Errorf("AddTransactionToHistoryDB: Unhandled transaction type '%s'", tx.Type)
	}
}

func (s *Synchronizer) AddAccountTx(tx *common.Tx) error {
	acc := &common.Account{
		Idx:           tx.FromIdx,
		EthAddr:       ethCommon.BytesToAddress(tx.FromEthAddr),
		Balance:       tx.Amount,     // Initial deposit amount becomes the balance
		Score:         big.NewInt(0), // Initial score
		ScoreSiblings: make([]*big.Int, 0),
	}
	err := s.historydb.AddAccount(acc)
	if err != nil {
		return fmt.Errorf("AddAccountTx: failed to add account: %w", err)
	}
	s.logger.Printf("Account created/added via AddAccountTx: Idx %d, EthAddr %s", acc.Idx, acc.EthAddr.Hex())
	return nil
}

func (s *Synchronizer) DepositWithdrawTx(tx *common.Tx) error {
	var newBalance *big.Int
	fromEthAddr := ethCommon.BytesToAddress(tx.FromEthAddr)
	account, err := s.historydb.GetAccountByEthAddress(fromEthAddr)
	if err != nil {
		return fmt.Errorf("DepositTx: failed to get account by ethAddr %s: %w", fromEthAddr.Hex(), err)
	}
	if account == nil {
		return fmt.Errorf("DepositTx: account not found for ethAddr %s", fromEthAddr.Hex())
	}

	if tx.Type == common.TxTypeDeposit {
		newBalance = new(big.Int).Add(account.Balance, tx.Amount)
		err = s.historydb.UpdateAccountBalance(account.Idx, newBalance)
		if err != nil {
			return fmt.Errorf("DepositTx: failed to update account balance for idx %d: %w", account.Idx, err)
		}
	} else if tx.Type == common.TxTypeWithdraw {
		if account.Balance.Cmp(tx.Amount) < 0 {
			return fmt.Errorf("WithdrawalTx: insufficient balance for account idx %d. Has: %s, Wants: %s",
				account.Idx, account.Balance.String(), tx.Amount.String())
		}
		newBalance := new(big.Int).Sub(account.Balance, tx.Amount)
		err = s.historydb.UpdateAccountBalance(account.Idx, newBalance)
		if err != nil {
			return fmt.Errorf("WithdrawalTx: failed to update account balance for idx %d: %w", account.Idx, err)
		}
	}

	s.logger.Printf("Account balance updated for: Idx %d, New Balance %s", account.Idx, newBalance.String())
	return nil
}

func (s *Synchronizer) VouchTx(tx *common.Tx) error {
	vouchingAccountEthAddr := ethCommon.BytesToAddress(tx.FromEthAddr)
	vouchedAccountEthAddr := ethCommon.BytesToAddress(tx.ToEthAddr)
	vouchingAccount, err := s.historydb.GetAccountByEthAddress(vouchingAccountEthAddr)
	if err != nil {
		return fmt.Errorf("VouchTx: failed to get 'from' account by ethAddr %s: %w", vouchingAccountEthAddr, err)
	}
	if vouchingAccount == nil {
		return fmt.Errorf("VouchTx: 'from' account not found for ethAddr %s", vouchingAccountEthAddr)
	}

	vouchedAccount, err := s.historydb.GetAccountByEthAddress(vouchedAccountEthAddr)
	if err != nil {
		return fmt.Errorf("VouchTx: failed to get 'to' account by ethAddr %s: %w", vouchedAccountEthAddr, err)
	}
	if vouchedAccount == nil {
		return fmt.Errorf("VouchTx: 'to' account not found for ethAddr %s", vouchedAccountEthAddr)
	}

	vouchTableKeyStr := fmt.Sprintf("%d%d", vouchingAccount.Idx, vouchedAccount.Idx)
	vouchTableKeyBigInt, ok := new(big.Int).SetString(vouchTableKeyStr, 10)
	if !ok {
		return fmt.Errorf("VouchTx: failed to create vouch table key from string '%s'", vouchTableKeyStr)
	}
	vouchTableKeyValue := common.VouchIdx(vouchTableKeyBigInt.Uint64())

	if tx.Type == common.TxTypeVouch {
		existingVouch, err := s.historydb.GetVouchByIdx(vouchTableKeyValue)
		if err != nil {
			return fmt.Errorf("VouchTx: failed to check for existing vouch: %w", err)
		}

		if existingVouch != nil {
			s.logger.Printf("Vouch already exists, skipping: Key %d (From Acct %d -> To Acct %d)", existingVouch.Idx, existingVouch.FromIdx, existingVouch.ToIdx)
		} else {
			vouchEntry := &common.Vouch{
				Idx:         vouchTableKeyValue,
				FromIdx:     vouchingAccount.Idx,
				FromEthAddr: vouchingAccount.EthAddr,
				ToIdx:       vouchedAccount.Idx,
				ToEthAddr:   vouchedAccount.EthAddr,
			}
			err = s.historydb.AddVouch(vouchEntry)
			if err != nil {
				return fmt.Errorf("VouchTx: failed to add vouch (key %d, from %d to %d): %w",
					vouchEntry.Idx, vouchEntry.FromIdx, vouchEntry.ToIdx, err)
			}
			s.logger.Printf("Vouch created: Key %d (From Acct %d -> To Acct %d)", vouchEntry.Idx, vouchingAccount.Idx, vouchedAccount.Idx)
		}
	} else if tx.Type == common.TxTypeUnvouch {
		err = s.historydb.DeleteVouchByIdx(vouchTableKeyValue)
		if err != nil {
			return fmt.Errorf("UnVouchTx: failed to delete vouch (key %s, from %d to %d): %w",
				vouchTableKeyStr, vouchingAccount.Idx, vouchedAccount.Idx, err)
		}
		s.logger.Printf("Vouch deleted: Key %s (From Acct %d -> To Acct %d)", vouchTableKeyStr, vouchingAccount.Idx, vouchedAccount.Idx)
	}

	return nil
}
