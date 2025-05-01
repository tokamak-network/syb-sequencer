package historydb

import (
	"math/big"
)

// Tx represents a transaction in the history database
type Tx struct {
	ItemID      int64
	BatchNum    int64
	Position    int
	Type        string
	FromIdx     *int64 // Nullable
	FromEthAddr []byte // Nullable
	ToIdx       int64
	ToEthAddr   []byte // Nullable
	Amount      *big.Int
}

// SaveTx saves a transaction to the database
func (db *HistoryDB) SaveTx(tx *Tx) error {
	// Convert big.Int to string for database storage
	amountStr := "0"
	if tx.Amount != nil {
		amountStr = tx.Amount.String()
	}

	// Insert the transaction into the database
	_, err := db.dbWrite.Exec(`
		INSERT INTO tx (
			batch_num, position, type, from_idx, from_eth_addr, 
			to_idx, to_eth_addr, amount
		) VALUES (
			$1, $2, $3, $4, $5, $6, $7, $8
		)`,
		tx.BatchNum, tx.Position, tx.Type, tx.FromIdx, tx.FromEthAddr,
		tx.ToIdx, tx.ToEthAddr, amountStr,
	)

	return err
}

// GetAllTxs retrieves all transactions from the database
func (db *HistoryDB) GetAllTxs() ([]*Tx, error) {
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

	var txs []*Tx
	for rows.Next() {
		var tx Tx
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
