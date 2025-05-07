package forger

import (
	"fmt"
	"log"
	"math/big"
	"sort"

	"github.com/tokamak-network/syb-sequencer/sequencer/db/historydb"
)

// Forger is responsible for creating batches from transactions
type Forger struct {
	db     *historydb.HistoryDB
	logger *log.Logger
}

var lastForgedBatch uint32

// NewForger creates a new Forger instance
func NewForger(db *historydb.HistoryDB, logger *log.Logger) *Forger {
	return &Forger{
		db:     db,
		logger: logger,
	}
}

func (f *Forger) GetLastForgedBatchNum() uint32 {
	return lastForgedBatch
}

// ProcessBatch processes transactions for a specific batch number
func (f *Forger) ForgeBatch(batchNum uint32) error {
	lastForgedBatch = batchNum
	f.logger.Printf("Processing batch %d", batchNum)

	// Get all transactions for the batch
	txs, err := f.db.GetTxsByBatchNum(int64(batchNum))
	if err != nil {
		return fmt.Errorf("failed to get transactions for batch %d: %w", batchNum, err)
	}

	f.logger.Printf("Found %d transactions for batch %d", len(txs), batchNum)

	// Sort transactions by position
	sort.Slice(txs, func(i, j int) bool {
		return txs[i].Position < txs[j].Position
	})

	// Print the sorted transactions
	//TODO: Remove this once done

	//TODO: Add logic here to call the txprocessor to update the stateDB, pass the zki to the circuit.
	f.logger.Printf("Sorted transactions for batch %d:", batchNum)
	for i, tx := range txs {
		fromAddr := "nil"
		if tx.FromEthAddr != "" {
			fromAddr = tx.FromEthAddr
		}

		fromIdx := "nil"
		if tx.FromIdx != nil {
			fromIdx = fmt.Sprintf("%d", *tx.FromIdx)
		}

		f.logger.Printf("  [%d] Type: %s, FromIdx: %s, FromAddr: %s, ToIdx: %d, ToAddr: %s, Amount: %s",
			i, tx.Type, fromIdx, fromAddr, tx.ToIdx, tx.ToEthAddr, tx.Amount.String())
	}

	//TODO: With the new roots saved in the statedb call forge function in the smart contract.
	batch := &historydb.Batch{
		ItemID:      int64(batchNum),
		AccountRoot: new(big.Int).SetInt64(1),
		VouchRoot:   new(big.Int).SetInt64(1),
		ScoreRoot:   new(big.Int).SetInt64(1),
	}

	err = f.db.AddBatch(batch)

	if err != nil {
		return err
	}

	return nil
}
