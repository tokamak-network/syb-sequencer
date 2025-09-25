package forger

import (
	"context"
	"fmt"
	"log"
	"math/big"
	"os"
	"sort"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	ethCommon "github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/tokamak-network/syb-sequencer/sequencer/abis/bindings"
	"github.com/tokamak-network/syb-sequencer/sequencer/common"
	"github.com/tokamak-network/syb-sequencer/sequencer/db/historydb"
	"github.com/tokamak-network/syb-sequencer/sequencer/db/statedb"
	txprocessor "github.com/tokamak-network/syb-sequencer/sequencer/txProcessor"
)

// Forger is responsible for creating batches from transactions
type Forger struct {
	historydb     *historydb.HistoryDB
	Statedb       *statedb.LocalStateDB
	logger        *log.Logger
	sybilContract *bindings.Sybil
	transactOpts  *bind.TransactOpts
	client        *ethclient.Client
}

// NewForger creates a new Forger instance
func NewForger(ethRPC, contractAddressHex string, historydb *historydb.HistoryDB, statedb *statedb.LocalStateDB, logger *log.Logger) (*Forger, error) {
	client, err := ethclient.Dial(ethRPC)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to Ethereum client: %v", err)
	}

	privateKeyHex := os.Getenv("PRIVATE_KEY")
	if privateKeyHex == "" {
		return nil, fmt.Errorf("PRIVATE_KEY environment variable not set")
	}
	privateKey, err := crypto.HexToECDSA(privateKeyHex)
	if err != nil {
		return nil, fmt.Errorf("failed to parse private key: %v", err)
	}

	chainID, err := client.ChainID(context.Background())
	if err != nil {
		return nil, fmt.Errorf("failed to get chain ID: %v", err)
	}

	transactOpts, err := bind.NewKeyedTransactorWithChainID(privateKey, chainID)
	if err != nil {
		return nil, fmt.Errorf("failed to create transactor: %v", err)
	}
	transactOpts.GasLimit = uint64(1000000)

	contractAddress := ethCommon.HexToAddress(contractAddressHex)
	sybilContract, err := bindings.NewSybil(contractAddress, client)
	if err != nil {
		return nil, fmt.Errorf("failed to instantiate Sybil contract: %v", err)
	}
	return &Forger{
		historydb:     historydb,
		Statedb:       statedb,
		logger:        logger,
		sybilContract: sybilContract,
		transactOpts:  transactOpts,
		client:        client,
	}, nil
}

// ProcessBatch processes transactions for a specific batch number
func (f *Forger) ForgeBatch(batchNum uint32) error {
	f.logger.Printf("Processing batch %d", batchNum)

	// Get all transactions for the batch
	txs, err := f.historydb.GetTxsByBatchNum(batchNum)
	if err != nil {
		return fmt.Errorf("failed to get transactions for batch %d: %w", batchNum, err)
	}

	f.logger.Printf("Found %d transactions for batch %d", len(txs), batchNum)

	// Sort transactions by position
	sort.Slice(txs, func(i, j int) bool {
		return txs[i].Position.Cmp(txs[j].Position) == -1
	})

	f.logger.Printf("Sorted transactions for batch %d:", batchNum)

	config := txprocessor.Config{
		NLevels: 24,
		MaxTx:   1,
		MaxL1Tx: 5,
		ChainID: 0,
	}

	newBatchBuilder := txprocessor.NewBatchBuilder(config, f.Statedb)

	zki, err := newBatchBuilder.ForgeTransactions(txs)
	if err != nil {
		return fmt.Errorf("failed to forge transactions for batch %d: %w", batchNum, err)
	}

	//TODO: Call the circuit with the ZKI to generate the proof

	batch := &common.Batch{
		ItemID:      common.BatchNum(batchNum),
		AccountRoot: zki.NewAccountRootRaw.BigInt(),
		VouchRoot:   zki.NewVouchRootRaw.BigInt(),
		ScoreRoot:   zki.NewScoreRootRaw.BigInt(),
	}
	proofA := [2]*big.Int{big.NewInt(0), big.NewInt(0)}
	proofB := [2][2]*big.Int{
		{big.NewInt(0), big.NewInt(0)},
		{big.NewInt(0), big.NewInt(0)},
	}
	proofC := [2]*big.Int{big.NewInt(0), big.NewInt(0)}

	forgeTx, err := f.sybilContract.ForgeBatch(f.transactOpts, batch.AccountRoot, batch.VouchRoot, batch.ScoreRoot, proofA, proofB, proofC)
	if err != nil {
		return fmt.Errorf("failed to call ForgeBatch: %w", err)
	}
	f.logger.Printf("ForgeBatch transaction hash: %s", forgeTx.Hash().Hex())
	receipt, err := bind.WaitMined(context.Background(), f.client, forgeTx)
	if err != nil {
		return fmt.Errorf("failed to wait for transaction to be mined: %w", err)
	}

	// Check if the transaction was successful
	if receipt.Status != 1 {
		return fmt.Errorf("transaction failed: %s", forgeTx.Hash().Hex())
	}
	// Add the batch to the history database
	err = f.historydb.AddBatch(batch)
	if err != nil {
		return err
	}
	f.Statedb.UpdateBatchBumber()
	return nil
}
