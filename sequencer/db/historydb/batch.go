package historydb

import (
	"math/big"
)

type Batch struct {
	ItemID      int64    `meddler:"item_id,pk"`
	AccountRoot *big.Int `meddler:"account_root,bigint"`
	VouchRoot   *big.Int `meddler:"vouch_root,bigint"`
	ScoreRoot   *big.Int `meddler:"score_root,bigint"`
}

// AddBatch inserts a new batch entry into the database
func (h *HistoryDB) AddBatch(batch *Batch) error {

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
