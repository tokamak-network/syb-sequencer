package historydb

import (
	"database/sql"
	"fmt"

	"github.com/tokamak-network/syb-sequencer/sequencer/common"
)

// AddBatch inserts a new batch entry into the database
func (h *HistoryDB) AddBatch(batch *common.Batch) error {

	accountRootStr := batch.AccountRoot.String()
	vouchRootStr := batch.VouchRoot.String()
	scoreRootStr := batch.ScoreRoot.String()

	_, err := h.dbWrite.Exec(`
		INSERT INTO batch (
			item_id, account_root, vouch_root, score_root
		) VALUES (
			$1, $2, $3, $4
		)
	`, batch.ItemID, accountRootStr, vouchRootStr, scoreRootStr)

	return err
}

func (h *HistoryDB) GetLastBatchItemID() (common.BatchNum, error) {
	var lastItemID common.BatchNum

	query := `SELECT item_id FROM batch ORDER BY item_id DESC LIMIT 1`

	// Query for the single value and scan it directly into lastItemID.
	err := h.dbRead.QueryRow(query).Scan(&lastItemID)
	if err != nil {
		// If no rows were found, we return the zero value for BatchNum (0)
		// and the original sql.ErrNoRows, so the caller can check for this case.
		if err == sql.ErrNoRows {
			return 0, err
		}
		// For any other error, wrap it with context.
		return 0, fmt.Errorf("failed to query last batch item_id: %w", err)
	}

	return lastItemID, nil
}
