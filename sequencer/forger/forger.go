package forger

import (
	"fmt"
	"log"
	"math/big"
	"sort"

	"github.com/tokamak-network/syb-sequencer/sequencer/common"
	"github.com/tokamak-network/syb-sequencer/sequencer/db/historydb"
	"github.com/tokamak-network/syb-sequencer/sequencer/db/statedb"
)

// Forger is responsible for creating batches from transactions
type Forger struct {
	historydb *historydb.HistoryDB
	statedb   *statedb.LocalStateDB
	logger    *log.Logger
}

var lastForgedBatch uint32

// NewForger creates a new Forger instance
func NewForger(historydb *historydb.HistoryDB, statedb *statedb.LocalStateDB, logger *log.Logger) *Forger {
	return &Forger{
		historydb: historydb,
		statedb:   statedb,
		logger:    logger,
	}
}

func (f *Forger) GetLastForgedBatchNum() uint32 {
	return lastForgedBatch
}

// ProcessBatch processes transactions for a specific batch number
func (f *Forger) ForgeBatch(batchNum uint32) error {
	f.logger.Printf("Processing batch %d", batchNum)

	// Get all transactions for the batch
	txs, err := f.historydb.GetTxsByBatchNum(int64(batchNum))
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
		f.logger.Printf("  [%d] Type: %s, FromIdx: %v, FromAddr: %s, ToIdx: %d, ToAddr: %s, Amount: %s",
			i, tx.Type, tx.FromIdx, tx.FromEthAddr, tx.ToIdx, tx.ToEthAddr, tx.Amount.String())
	}

	//TODO: With the new roots saved in the statedb call forge function in the smart contract.
	batch := &common.Batch{
		ItemID:      common.BatchNum(batchNum),
		AccountRoot: new(big.Int).SetInt64(1),
		VouchRoot:   new(big.Int).SetInt64(1),
		ScoreRoot:   new(big.Int).SetInt64(1),
	}

	err = f.historydb.AddBatch(batch)

	if err != nil {
		return err
	}
	lastForgedBatch = batchNum
	return nil
}
