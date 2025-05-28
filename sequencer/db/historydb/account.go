package historydb

import (
	"database/sql"
	"fmt"
	"math/big"

	ethCommon "github.com/ethereum/go-ethereum/common"
	"github.com/lib/pq"
	"github.com/tokamak-network/syb-sequencer/sequencer/common"
)

func (hdb *HistoryDB) AddAccount(acc *common.Account) error {
	query := `
		INSERT INTO account (idx, eth_addr, balance, score, score_siblings)
		VALUES ($1, $2, $3, $4, $5)
	`

	scoreSiblingsStr := make([]string, len(acc.ScoreSiblings))
	for i, s := range acc.ScoreSiblings {
		if s == nil {
			scoreSiblingsStr[i] = "0"
		} else {
			scoreSiblingsStr[i] = s.String()
		}
	}

	_, err := hdb.dbWrite.Exec(
		query,
		acc.Idx,
		acc.EthAddr.Bytes(),
		acc.Balance.String(),
		acc.Score.String(),
		pq.Array(scoreSiblingsStr),
	)

	if err != nil {
		return fmt.Errorf("AddAccount: failed to add account with idx %d: %w", acc.Idx, err)
	}
	return nil
}

func (hdb *HistoryDB) GetAccountByIdx(idx uint32) (*common.Account, error) {
	query := `
		SELECT eth_addr, balance, score, score_siblings
		FROM account
		WHERE idx = $1;
	`
	row := hdb.dbRead.QueryRow(query, idx)

	acc := &common.Account{Idx: common.AccountIdx(idx)}
	var ethAddrBytes []byte
	var balanceStr, scoreStr string
	var scoreSiblingsStr pq.StringArray

	err := row.Scan(
		&ethAddrBytes,
		&balanceStr,
		&scoreStr,
		&scoreSiblingsStr,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, fmt.Errorf("GetAccountByIdx: failed to get account with idx %d: %w", idx, err)
	}

	acc.EthAddr = ethCommon.BytesToAddress(ethAddrBytes)

	var ok bool
	acc.Balance, ok = new(big.Int).SetString(balanceStr, 10)
	if !ok {
		return nil, fmt.Errorf("GetAccountByIdx: failed to parse balance string '%s' for account idx %d", balanceStr, idx)
	}
	acc.Score, ok = new(big.Int).SetString(scoreStr, 10)
	if !ok {
		return nil, fmt.Errorf("GetAccountByIdx: failed to parse score string '%s' for account idx %d", scoreStr, idx)
	}

	acc.ScoreSiblings = make([]*big.Int, len(scoreSiblingsStr))
	for i, s := range scoreSiblingsStr {
		val, ok := new(big.Int).SetString(s, 10)
		if !ok {
			acc.ScoreSiblings[i] = big.NewInt(0)
		} else {
			acc.ScoreSiblings[i] = val
		}
	}

	return acc, nil
}

func (hdb *HistoryDB) GetAccountByEthAddress(ethAddr ethCommon.Address) (*common.Account, error) {
	query := `
		SELECT idx, balance, score, score_siblings
		FROM account
		WHERE eth_addr = $1;
	`
	row := hdb.dbRead.QueryRow(query, ethAddr.Bytes())

	acc := &common.Account{EthAddr: ethAddr}
	var balanceStr, scoreStr string
	var scoreSiblingsStr pq.StringArray

	err := row.Scan(
		&acc.Idx,
		&balanceStr,
		&scoreStr,
		&scoreSiblingsStr,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, fmt.Errorf("GetAccountByEthAddress: failed to get account with eth_addr %s: %w", ethAddr.String(), err)
	}

	var ok bool
	acc.Balance, ok = new(big.Int).SetString(balanceStr, 10)
	if !ok {
		return nil, fmt.Errorf("GetAccountByEthAddress: failed to parse balance string '%s' for eth_addr %s", balanceStr, ethAddr.String())
	}
	acc.Score, ok = new(big.Int).SetString(scoreStr, 10)
	if !ok {
		return nil, fmt.Errorf("GetAccountByEthAddress: failed to parse score string '%s' for eth_addr %s", scoreStr, ethAddr.String())
	}

	acc.ScoreSiblings = make([]*big.Int, len(scoreSiblingsStr))
	for i, s := range scoreSiblingsStr {
		val, ok := new(big.Int).SetString(s, 10)
		if !ok {
			acc.ScoreSiblings[i] = big.NewInt(0)
		} else {
			acc.ScoreSiblings[i] = val
		}
	}
	return acc, nil
}

func (hdb *HistoryDB) UpdateAccountBalance(idx common.AccountIdx, newBalance *big.Int) error {
	query := `UPDATE account SET balance = $1 WHERE idx = $2;`
	result, err := hdb.dbWrite.Exec(query, newBalance.String(), idx)
	if err != nil {
		return fmt.Errorf("UpdateAccountBalance: failed to update balance for account idx %d: %w", idx, err)
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("UpdateAccountBalance: failed to get rows affected for account idx %d: %w", idx, err)
	}
	if rowsAffected == 0 {
		return fmt.Errorf("UpdateAccountBalance: no account found with idx %d to update", idx)
	}
	return nil
}
