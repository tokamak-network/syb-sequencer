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
		var amountStr string
		var gasFeeStr string
		var position string

		// Ensure the order of scanned fields matches the SELECT statements in calling functions
		err := rows.Scan(
			&tx.ItemID, &tx.BatchNum, &position, &tx.Type, &tx.FromIdx, &tx.FromEthAddr,
			&tx.ToIdx, &tx.ToEthAddr, &amountStr,
			&tx.BlockNumber, &tx.Timestamp, &gasFeeStr, &tx.TxHash,
		)
		if err != nil {
			return nil, fmt.Errorf("failed to scan transaction row: %w", err)
		}

		tx.Amount = new(big.Int)
		if amountStr != "" {
			_, success := tx.Amount.SetString(amountStr, 10)
			if !success {
				return nil, fmt.Errorf("failed to parse amount string '%s' for transaction item_id %d: %w", amountStr, tx.ItemID, err)
			}
		}

		tx.GasFee = new(big.Int)
		if gasFeeStr != "" {
			_, success := tx.GasFee.SetString(gasFeeStr, 10)
			if !success {
				return nil, fmt.Errorf("failed to parse gas_fee string '%s' for transaction item_id %d: %w", gasFeeStr, tx.ItemID, err)
			}
		}

		tx.Position = new(big.Int)
		if position != "" {
			_, success := tx.Position.SetString(position, 10)
			if !success {
				return nil, fmt.Errorf("failed to parse position string '%s' for transaction item_id %d: %w", position, tx.ItemID, err)
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
	positionStr := tx.Position.String()
	amountStr := "0"
	if tx.Amount != nil {
		amountStr = tx.Amount.String()
	}

	gasFeeStr := "0"
	if tx.GasFee != nil {
		gasFeeStr = tx.GasFee.String()
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
		positionStr,
		tx.Type,
		tx.FromIdx,
		tx.FromEthAddr,
		tx.ToIdx,
		tx.ToEthAddr,
		amountStr,
		tx.BlockNumber,
		tx.Timestamp,
		gasFeeStr,
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
			block_number, tx_timestamp, gas_fee, tx_hash
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
			block_number, tx_timestamp, gas_fee, tx_hash
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
			block_number, tx_timestamp, gas_fee, tx_hash
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
			block_number, tx_timestamp, gas_fee, tx_hash
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
