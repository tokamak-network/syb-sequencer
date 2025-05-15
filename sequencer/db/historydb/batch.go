package historydb

import "github.com/tokamak-network/syb-sequencer/sequencer/common"

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
