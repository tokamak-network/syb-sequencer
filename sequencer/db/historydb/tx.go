package historydb

import (
	"database/sql"
	"fmt"
	"math/big"
	"strings"

	ethCommon "github.com/ethereum/go-ethereum/common"
	"github.com/tokamak-network/syb-sequencer/sequencer/common"
)

func scanTxs(rows *sql.Rows) ([]*common.Tx, error) {
	defer rows.Close() // Ensure rows are closed

	var txs []*common.Tx
	for rows.Next() {
		var tx common.Tx
		var amountText sql.NullString
		var gasFeeText sql.NullString
		var positionText sql.NullString
		var fromIdxVal sql.NullInt64
		var toIdxVal sql.NullInt64

		// Ensure the order of scanned fields matches the SELECT statements in calling functions
		err := rows.Scan(
			&tx.ItemID, &tx.BatchNum, &positionText, &tx.Type, &fromIdxVal, &tx.FromEthAddr,
			&toIdxVal, &tx.ToEthAddr, &amountText,
			&tx.BlockNumber, &tx.Timestamp, &gasFeeText, &tx.TxHash, &tx.IsTxForged,
		)
		if err != nil {
			return nil, fmt.Errorf("failed to scan transaction row: %w", err)
		}

		// Nullable from_idx / to_idx
		if fromIdxVal.Valid {
			tx.FromIdx = common.AccountIdx(uint32(fromIdxVal.Int64))
		}
		if toIdxVal.Valid {
			tx.ToIdx = common.AccountIdx(uint32(toIdxVal.Int64))
		}

		// Nullable numeric big.Ints: amount, gas_fee, position
		if amountText.Valid && amountText.String != "" {
			amount := new(big.Int)
			if _, ok := amount.SetString(amountText.String, 10); ok {
				tx.Amount = amount
			} else {
				return nil, fmt.Errorf("failed to parse amount string '%s' for transaction item_id %d", amountText.String, tx.ItemID)
			}
		}

		if gasFeeText.Valid && gasFeeText.String != "" {
			gas := new(big.Int)
			if _, ok := gas.SetString(gasFeeText.String, 10); ok {
				tx.GasFee = gas
			} else {
				return nil, fmt.Errorf("failed to parse gas_fee string '%s' for transaction item_id %d", gasFeeText.String, tx.ItemID)
			}
		}

		if positionText.Valid && positionText.String != "" {
			pos := new(big.Int)
			if _, ok := pos.SetString(positionText.String, 10); ok {
				tx.Position = pos
			} else {
				return nil, fmt.Errorf("failed to parse position string '%s' for transaction item_id %d", positionText.String, tx.ItemID)
			}
		}

		txs = append(txs, &tx)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error iterating transaction rows: %w", err)
	}

	return txs, nil
}

// SaveTx saves a transaction to the database
func (db *HistoryDB) SaveTx(tx *common.Tx) error {
	// Prepare nullable numeric fields as proper SQL NULLs when absent
	var positionVal interface{}
	if tx.Position != nil {
		positionVal = tx.Position.String()
	} else {
		positionVal = nil
	}

	var amountVal interface{}
	if tx.Amount != nil {
		amountVal = tx.Amount.String()
	} else {
		amountVal = nil
	}

	var gasFeeVal interface{}
	if tx.GasFee != nil {
		gasFeeVal = tx.GasFee.String()
	} else {
		gasFeeVal = nil
	}

	var err error // Declare error variable once

	// Insert the transaction into the database, wrapping bytes with pq.Bytea
	_, err = db.dbWrite.Exec(`
		INSERT INTO tx (
			batch_num, position, type, from_idx, from_eth_addr, 
			to_idx, to_eth_addr, amount,
			block_number, tx_timestamp, gas_fee, tx_hash
		) VALUES (
			$1, $2, $3, $4, $5, $6, $7, $8,
			$9, $10, $11, $12
		)`,
		tx.BatchNum,
		positionVal,
		tx.Type,
		tx.FromIdx,
		tx.FromEthAddr,
		tx.ToIdx,
		tx.ToEthAddr,
		amountVal,
		tx.BlockNumber,
		tx.Timestamp,
		gasFeeVal,
		tx.TxHash,
	)

	if err != nil {
		return fmt.Errorf("failed to execute insert transaction statement: %w", err)
	}

	return nil
}

// GetAllTxs retrieves all transactions from the database
func (db *HistoryDB) GetAllTxs() ([]*common.Tx, error) {
	rows, err := db.dbRead.Query(`
		SELECT 
			item_id, batch_num, position, type, from_idx, from_eth_addr,
			to_idx, to_eth_addr, amount,
			block_number, tx_timestamp, gas_fee, tx_hash, 
			is_tx_forged
		FROM tx
		ORDER BY item_id DESC
	`)
	if err != nil {
		return nil, fmt.Errorf("failed to query all transactions: %w", err)
	}
	return scanTxs(rows)
}

// GetTxsByBatchNum retrieves all transactions for a specific batch
func (db *HistoryDB) GetTxsByBatchNum(batchNum uint32) ([]*common.Tx, error) {
	rows, err := db.dbRead.Query(`
		SELECT 
			item_id, batch_num, position, type, from_idx, from_eth_addr,
			to_idx, to_eth_addr, amount,
			block_number, tx_timestamp, gas_fee, tx_hash, is_tx_forged
		FROM tx
		WHERE batch_num = $1
		ORDER BY position
	`, batchNum)
	if err != nil {
		return nil, fmt.Errorf("failed to query transactions by batch number %d: %w", batchNum, err)
	}
	txs, err := scanTxs(rows)
	if err != nil {
		return nil, fmt.Errorf("failed to process transactions for batch number %d: %w", batchNum, err)
	}
	return txs, nil
}

func (db *HistoryDB) GetTxsByAccountAddress(accountAddress string) ([]*common.Tx, error) {
	addrBytes := ethCommon.HexToAddress(accountAddress).Bytes()

	rows, err := db.dbRead.Query(`
		SELECT 
			item_id, batch_num, position, type, from_idx, from_eth_addr,
			to_idx, to_eth_addr, amount,
			block_number, tx_timestamp, gas_fee, tx_hash, is_tx_forged
		FROM tx
		WHERE from_eth_addr = $1
		ORDER BY item_id DESC
	`, addrBytes)
	if err != nil {
		return nil, fmt.Errorf("failed to query transactions by account address %s: %w", accountAddress, err)
	}
	txs, err := scanTxs(rows)
	if err != nil {
		return nil, fmt.Errorf("failed to process transactions for account address %s: %w", accountAddress, err)
	}
	return txs, nil
}

func (db *HistoryDB) GetTxsPaginated(limit, offset int, sortBy, sortOrder string) ([]*common.Tx, int64, error) {
	// Whitelist columns for sorting to prevent SQL injection
	allowedSortColumns := map[string]string{
		"item_id":      "item_id",
		"batch_num":    "batch_num",
		"type":         "type",
		"block_number": "block_number",
		"tx_timestamp": "tx_timestamp",
	}
	dbSortBy, ok := allowedSortColumns[strings.ToLower(sortBy)]
	if !ok {
		dbSortBy = "item_id" // Default sort column
	}

	dbSortOrder := "DESC" // Default sort order
	if strings.ToUpper(sortOrder) == "ASC" {
		dbSortOrder = "ASC"
	}

	query := fmt.Sprintf(`
		SELECT
			item_id, batch_num, position, type, from_idx, from_eth_addr,
			to_idx, to_eth_addr, amount,
			block_number, tx_timestamp, gas_fee, tx_hash, is_tx_forged
		FROM tx
		ORDER BY %s %s
		LIMIT $1 OFFSET $2
	`, dbSortBy, dbSortOrder)

	rows, err := db.dbRead.Query(query, limit, offset)
	if err != nil {
		return nil, 0, fmt.Errorf("failed to query paginated transactions: %w", err)
	}
	txs, err := scanTxs(rows)
	if err != nil {
		return nil, 0, fmt.Errorf("failed to process paginated transactions: %w", err)
	}

	var totalItems int64
	err = db.dbRead.QueryRow("SELECT COUNT(*) FROM tx").Scan(&totalItems)
	if err != nil {
		return nil, 0, fmt.Errorf("failed to query total transaction count: %w", err)
	}

	return txs, totalItems, nil
}

// GetTxByHash retrieves a transaction by its hash
func (db *HistoryDB) GetTxByHash(txHash []byte) (*common.Tx, error) {
	query := `
		SELECT 
			item_id, batch_num, position, type, from_idx, from_eth_addr,
			to_idx, to_eth_addr, amount,
			block_number, tx_timestamp, gas_fee, tx_hash, is_tx_forged
		FROM tx
		WHERE tx_hash = $1
	`
	row := db.dbRead.QueryRow(query, txHash)

	var tx common.Tx
	var amountText sql.NullString
	var gasFeeText sql.NullString
	var positionText sql.NullString
	var fromIdxVal sql.NullInt64
	var toIdxVal sql.NullInt64

	err := row.Scan(
		&tx.ItemID, &tx.BatchNum, &positionText, &tx.Type, &fromIdxVal, &tx.FromEthAddr,
		&toIdxVal, &tx.ToEthAddr, &amountText,
		&tx.BlockNumber, &tx.Timestamp, &gasFeeText, &tx.TxHash, &tx.IsTxForged,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, fmt.Errorf("failed to get transaction by hash: %w", err)
	}

	if fromIdxVal.Valid {
		tx.FromIdx = common.AccountIdx(uint32(fromIdxVal.Int64))
	}
	if toIdxVal.Valid {
		tx.ToIdx = common.AccountIdx(uint32(toIdxVal.Int64))
	}

	if amountText.Valid && amountText.String != "" {
		amount := new(big.Int)
		if _, ok := amount.SetString(amountText.String, 10); ok {
			tx.Amount = amount
		} else {
			return nil, fmt.Errorf("failed to parse amount string '%s' for transaction hash %x", amountText.String, txHash)
		}
	}

	if gasFeeText.Valid && gasFeeText.String != "" {
		gas := new(big.Int)
		if _, ok := gas.SetString(gasFeeText.String, 10); ok {
			tx.GasFee = gas
		} else {
			return nil, fmt.Errorf("failed to parse gas_fee string '%s' for transaction hash %x", gasFeeText.String, txHash)
		}
	}

	if positionText.Valid && positionText.String != "" {
		pos := new(big.Int)
		if _, ok := pos.SetString(positionText.String, 10); ok {
			tx.Position = pos
		} else {
			return nil, fmt.Errorf("failed to parse position string '%s' for transaction hash %x", positionText.String, txHash)
		}
	}

	return &tx, nil
}

// MarkTxsForgedByBatch sets is_tx_forged = TRUE for all transactions in a batch.
// It returns the number of rows updated.
func (db *HistoryDB) MarkTxsForgedByBatch(batchNum uint32) (int64, error) {
	res, err := db.dbWrite.Exec(`
        UPDATE tx
        SET is_tx_forged = TRUE
        WHERE batch_num = $1 AND is_tx_forged = FALSE
    `, batchNum)
	if err != nil {
		return 0, fmt.Errorf("failed to mark batch txs as forged: %w", err)
	}

	affected, err := res.RowsAffected()
	if err != nil {
		return 0, fmt.Errorf("failed to read rows affected: %w", err)
	}
	return affected, nil
}
