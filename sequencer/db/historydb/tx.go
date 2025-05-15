package historydb

import (
	"fmt"
	"math/big"

	"github.com/tokamak-network/syb-sequencer/sequencer/common"
)

// SaveTx saves a transaction to the database
func (db *HistoryDB) SaveTx(tx *common.Tx) error {
	// Convert big.Int to string for database storage
	amountStr := "0"
	if tx.Amount != nil {
		amountStr = tx.Amount.String()
	}

	var err error // Declare error variable once

	// Insert the transaction into the database, wrapping bytes with pq.Bytea
	_, err = db.dbWrite.Exec(`
		INSERT INTO tx (
			batch_num, position, type, from_idx, from_eth_addr, 
			to_idx, to_eth_addr, amount
		) VALUES (
			$1, $2, $3, $4, $5, $6, $7, $8
		)`,
		tx.BatchNum,
		tx.Position,
		tx.Type,
		tx.FromIdx,
		tx.FromEthAddr,
		tx.ToIdx,
		tx.ToEthAddr,
		amountStr,
	)

	// Return the error from the Exec call (if any)
	if err != nil {
		// Add more context to the error if it's from Exec
		return fmt.Errorf("failed to execute insert transaction statement: %w", err)
	}

	return nil
}

// GetAllTxs retrieves all transactions from the database
func (db *HistoryDB) GetAllTxs() ([]*common.Tx, error) {
	rows, err := db.dbWrite.Query(`
		SELECT 
			item_id, batch_num, position, type, from_idx, from_eth_addr,
			to_idx, to_eth_addr, amount
		FROM tx
		ORDER BY item_id DESC
	`)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var txs []*common.Tx
	for rows.Next() {
		var tx common.Tx
		var amountStr string

		err := rows.Scan(
			&tx.ItemID, &tx.BatchNum, &tx.Position, &tx.Type, &tx.FromIdx, &tx.FromEthAddr,
			&tx.ToIdx, &tx.ToEthAddr, &amountStr,
		)

		if err != nil {
			return nil, err
		}

		// Convert amount string back to big.Int
		tx.Amount = new(big.Int)
		tx.Amount.SetString(amountStr, 10)

		txs = append(txs, &tx)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return txs, nil
}

// GetTxsByBatchNum retrieves all transactions for a specific batch
func (db *HistoryDB) GetTxsByBatchNum(batchNum int64) ([]*common.Tx, error) {
	rows, err := db.dbWrite.Query(`
		SELECT 
			item_id, batch_num, position, type, from_idx, from_eth_addr,
			to_idx, to_eth_addr, amount
		FROM tx
		WHERE batch_num = $1
		ORDER BY position
	`, batchNum)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var txs []*common.Tx
	for rows.Next() {
		var tx common.Tx
		var amountStr string

		err := rows.Scan(
			&tx.ItemID, &tx.BatchNum, &tx.Position, &tx.Type, &tx.FromIdx, &tx.FromEthAddr,
			&tx.ToIdx, &tx.ToEthAddr, &amountStr,
		)

		if err != nil {
			return nil, err
		}

		// Convert amount string back to big.Int
		tx.Amount = new(big.Int)
		tx.Amount.SetString(amountStr, 10)

		txs = append(txs, &tx)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return txs, nil
}
