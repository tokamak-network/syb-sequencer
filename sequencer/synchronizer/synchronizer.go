package synchronizer

import (
	"context"
	"fmt"
	"log"
	"math/big"
	"sort"
	"time"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	ethCommon "github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/tokamak-network/syb-sequencer/sequencer/abis/bindings"
	"github.com/tokamak-network/syb-sequencer/sequencer/common"
	"github.com/tokamak-network/syb-sequencer/sequencer/db/historydb"
	"github.com/tokamak-network/syb-sequencer/sequencer/db/statedb"
	"github.com/tokamak-network/syb-sequencer/sequencer/forger"
)

// Synchronizer listens to contract events and stores them in the database
type Synchronizer struct {
	client          *ethclient.Client
	contractAddress ethCommon.Address
	sybilContract   *bindings.Sybil
	historydb       *historydb.HistoryDB
	statedb         *statedb.StateDB
	logs            chan types.Log
	sub             ethereum.Subscription
	logger          *log.Logger
	forger          *forger.Forger
	lastBlock       int64
	// Add persistent tracking for last processed block
	lastProcessedBlockKey []byte
}

// NewSynchronizer creates a new synchronizer
func NewSynchronizer(ethRPC, contractAddressHex string, historydb *historydb.HistoryDB, statedb *statedb.StateDB, logger *log.Logger, forger *forger.Forger) (*Synchronizer, error) {
	client, err := ethclient.Dial(ethRPC)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to Ethereum client: %v", err)
	}

	contractAddress := ethCommon.HexToAddress(contractAddressHex)
	sybilContract, err := bindings.NewSybil(contractAddress, client)
	if err != nil {
		return nil, fmt.Errorf("failed to instantiate Sybil contract: %v", err)
	}

	logs := make(chan types.Log)

	return &Synchronizer{
		client:                client,
		contractAddress:       contractAddress,
		sybilContract:         sybilContract,
		historydb:             historydb,
		statedb:               statedb,
		logs:                  logs,
		logger:                logger,
		forger:                forger,
		lastBlock:             0,
		lastProcessedBlockKey: []byte("last_processed_block"),
	}, nil
}

// Start begins the synchronization process
func (s *Synchronizer) Start(ctx context.Context) {
	// Load the last processed block from persistent storage
	s.loadLastProcessedBlock()

	// Get the latest block number to start from
	header, err := s.client.HeaderByNumber(ctx, nil)
	if err != nil {
		s.logger.Fatalf("Failed to get latest block header: %v", err)
	}

	latestBlock := header.Number.Int64()

	// If we have a last processed block, use it; otherwise start from current block
	if s.lastBlock > 0 {
		s.logger.Printf("Resuming synchronizer from last processed block %d", s.lastBlock)
		// Check for missed events since last processed block
		s.checkForMissedEvents(ctx)
	} else {
		s.lastBlock = latestBlock
		s.logger.Printf("Starting synchronizer from block %d", s.lastBlock)
	}

	// Create a filter query for the contract events
	query := ethereum.FilterQuery{
		Addresses: []ethCommon.Address{s.contractAddress},
	}

	// Subscribe to logs
	sub, err := s.client.SubscribeFilterLogs(context.Background(), query, s.logs)
	if err != nil {
		s.logger.Fatalf("Failed to subscribe to logs: %v", err)
	}
	s.sub = sub

	// Start listening for events
	go s.watchEvents(ctx)
}

// watchEvents continuously listens for contract events
func (s *Synchronizer) watchEvents(ctx context.Context) {
	// Create a ticker for periodic safety checks
	ticker := time.NewTicker(5 * time.Minute)
	defer ticker.Stop()

	s.logger.Println("Started watching for contract events")

	for {
		select {
		case <-ctx.Done():
			s.logger.Println("Context cancelled, stopping synchronizer")
			return

		case err := <-s.sub.Err():
			s.logger.Printf("Error in subscription: %v", err)
			// Try to resubscribe
			s.resubscribe(ctx)

		case vLog := <-s.logs:
			// Process the log event in real-time
			s.logger.Printf("Received event in block %d, tx: %s", vLog.BlockNumber, vLog.TxHash.Hex())
			s.processLog(vLog)

		case <-ticker.C:
			// Periodic safety check to ensure we haven't missed any events
			// This is just a backup mechanism and not the primary way of getting events
			s.logger.Println("Performing periodic safety check for missed events")
			s.checkForMissedEvents(ctx)
		}
	}
}

// resubscribe attempts to reestablish the subscription
func (s *Synchronizer) resubscribe(ctx context.Context) {
	backoff := 1 * time.Second
	maxBackoff := 2 * time.Minute

	for {
		s.logger.Printf("Attempting to resubscribe in %v...", backoff)
		time.Sleep(backoff)

		query := ethereum.FilterQuery{
			Addresses: []ethCommon.Address{s.contractAddress},
		}

		sub, err := s.client.SubscribeFilterLogs(ctx, query, s.logs)
		if err != nil {
			s.logger.Printf("Failed to resubscribe: %v", err)
			// Exponential backoff with a maximum
			backoff *= 2
			if backoff > maxBackoff {
				backoff = maxBackoff
			}
			continue
		}

		s.sub = sub
		s.logger.Println("Successfully resubscribed to contract events")
		return
	}
}

// loadLastProcessedBlock loads the last processed block from persistent storage
func (s *Synchronizer) loadLastProcessedBlock() {
	// Use StateDB to store the last processed block
	err := s.statedb.LastRead(func(last *statedb.Last) error {
		blockBytes, err := last.DB().Get(s.lastProcessedBlockKey)
		if err != nil {
			// If not found, start from 0
			s.lastBlock = 0
			return nil
		}

		// Convert bytes to int64
		blockNum := new(big.Int).SetBytes(blockBytes)
		s.lastBlock = blockNum.Int64()
		return nil
	})

	if err != nil {
		s.logger.Printf("Error loading last processed block: %v, starting from 0", err)
		s.lastBlock = 0
	}
}

// saveLastProcessedBlock saves the last processed block to persistent storage
func (s *Synchronizer) saveLastProcessedBlock(blockNum int64) {
	blockBytes := big.NewInt(blockNum).Bytes()

	// Use StateDB to store the last processed block
	err := s.statedb.LastRead(func(last *statedb.Last) error {
		// Use the underlying storage with transaction
		tx, err := last.DB().NewTx()
		if err != nil {
			return err
		}
		err = tx.Put(s.lastProcessedBlockKey, blockBytes)
		if err != nil {
			return err
		}
		return tx.Commit()
	})

	if err != nil {
		s.logger.Printf("Error saving last processed block: %v", err)
	}
}

// checkForMissedEvents is a comprehensive safety mechanism to check for any events we might have missed
func (s *Synchronizer) checkForMissedEvents(ctx context.Context) {
	// Get the latest block number
	header, err := s.client.HeaderByNumber(ctx, nil)
	if err != nil {
		s.logger.Printf("Error getting latest block header: %v", err)
		return
	}

	latestBlock := header.Number.Int64()
	if latestBlock <= s.lastBlock {
		return // No new blocks
	}

	s.logger.Printf("Checking for missed events from block %d to %d", s.lastBlock+1, latestBlock)

	// Process missed events in chunks to avoid overwhelming the system
	const maxBlocksPerChunk = 1000
	startBlock := s.lastBlock + 1

	for startBlock <= latestBlock {
		endBlock := startBlock + maxBlocksPerChunk - 1
		if endBlock > latestBlock {
			endBlock = latestBlock
		}

		s.logger.Printf("Processing missed events from block %d to %d", startBlock, endBlock)

		// Create a filter query for the contract events in this chunk
		query := ethereum.FilterQuery{
			FromBlock: big.NewInt(startBlock),
			ToBlock:   big.NewInt(endBlock),
			Addresses: []ethCommon.Address{s.contractAddress},
		}

		// Get logs matching the filter
		logs, err := s.client.FilterLogs(ctx, query)
		if err != nil {
			s.logger.Printf("Error filtering logs for blocks %d-%d: %v", startBlock, endBlock, err)
			// Continue with next chunk instead of failing completely
			startBlock = endBlock + 1
			continue
		}

		// Sort logs by block number and transaction index to maintain order
		sort.Slice(logs, func(i, j int) bool {
			if logs[i].BlockNumber != logs[j].BlockNumber {
				return logs[i].BlockNumber < logs[j].BlockNumber
			}
			return logs[i].TxIndex < logs[j].TxIndex
		})

		// Process missed events in order
		for _, vLog := range logs {
			s.logger.Printf("Processing missed event from block %d, tx: %s", vLog.BlockNumber, vLog.TxHash.Hex())
			s.processMissedLog(vLog)
		}

		// Update the last processed block for this chunk
		s.lastBlock = endBlock
		s.saveLastProcessedBlock(endBlock)

		startBlock = endBlock + 1
	}

	s.logger.Printf("Completed processing missed events up to block %d", latestBlock)
}

// processMissedLog processes a single missed log entry with special handling for batch processing
func (s *Synchronizer) processMissedLog(vLog types.Log) {
	ctx := context.Background()
	callOpts := &bind.CallOpts{
		Context: ctx,
	}

	s.logger.Printf("Processing missed log: BlockNumber=%d TxHash=%s", vLog.BlockNumber, vLog.TxHash.Hex())

	// Parse the event
	eventData, eventType, err := ParseEvent(&vLog)
	if err != nil {
		s.logger.Printf("Error parsing missed event: %v", err)
		return
	}

	// Get current contract state
	lastForgedTx, err := s.sybilContract.LastForgedTxn(callOpts)
	if err != nil {
		s.logger.Printf("Error fetching last forged transaction: %v", err)
		return
	}

	batchSize, err := s.sybilContract.BatchSize(callOpts)
	if err != nil {
		s.logger.Printf("Error fetching batch size: %v", err)
		return
	}

	lastForgedBatch, err := s.sybilContract.LastForgedBatch(callOpts)
	if err != nil {
		s.logger.Printf("Error fetching last forged batch: %v", err)
		return
	}

	// Handle different event types
	switch e := eventData.(type) {
	case *bindings.SybilTxEvent:
		// Process transaction event
		s.processMissedTxEvent(e, vLog, lastForgedTx, batchSize, big.NewInt(int64(lastForgedBatch)))

	case *bindings.SybilForgeBatch:
		// For missed ForgeBatch events, we need to update our state
		s.logger.Printf("Processing missed ForgeBatch event for batch %d", e.LastForgedBatch)
		// Update the last forged batch in our tracking TODO

	default:
		s.logger.Printf("Unknown missed event type: %s", eventType)
	}
}

// processMissedTxEvent processes a missed transaction event with proper batch handling
func (s *Synchronizer) processMissedTxEvent(eventData *bindings.SybilTxEvent, vLog types.Log, lastForgedTx, batchSize, lastForgedBatch *big.Int) {
	ctx := context.Background()

	// Create transaction object
	tx := &common.Tx{
		BatchNum: uint32(lastForgedBatch.Uint64() + 1),
		Type:     "Unknown", // Will be set below
		FromIdx:  0,
		ToIdx:    0,
		Amount:   big.NewInt(0),
	}

	// Fetch Block Timestamp
	header, err := s.client.HeaderByNumber(ctx, new(big.Int).SetUint64(vLog.BlockNumber))
	if err != nil {
		s.logger.Printf("Error fetching block header for block %d: %v", vLog.BlockNumber, err)
	} else {
		tx.Timestamp = header.Time
	}

	tx.BlockNumber = vLog.BlockNumber

	// Fetch Transaction Receipt for Gas Fee
	receipt, err := s.client.TransactionReceipt(ctx, vLog.TxHash)
	if err != nil {
		s.logger.Printf("Error fetching transaction receipt for tx %s: %v", vLog.TxHash.Hex(), err)
	} else {
		if receipt.EffectiveGasPrice != nil && receipt.GasUsed > 0 {
			gasFee := new(big.Int).Mul(new(big.Int).SetUint64(receipt.GasUsed), receipt.EffectiveGasPrice)
			tx.GasFee = gasFee
		} else {
			tx.GasFee = big.NewInt(0)
		}
	}

	// Get transaction sender
	var sender ethCommon.Address
	txDetails, isPending, err := s.client.TransactionByHash(ctx, vLog.TxHash)
	if err != nil {
		s.logger.Printf("Error fetching full transaction details for tx %s: %v", vLog.TxHash.Hex(), err)
	} else if !isPending && txDetails != nil {
		chainID, err := s.client.ChainID(ctx)
		if err != nil {
			s.logger.Printf("Error fetching chain ID for tx %s: %v", vLog.TxHash.Hex(), err)
		} else {
			signer := types.LatestSignerForChainID(chainID)
			sender, err = types.Sender(signer, txDetails)
			if err != nil {
				s.logger.Printf("Error deriving sender for transaction %s: %v", vLog.TxHash.Hex(), err)
			}
		}
	}

	// Parse transaction data
	txType, fromEthAddr, toEthAddr, amount, fromIdx, toIdx, err := s.ParseTxData(eventData, sender)
	if err != nil {
		s.logger.Printf("Error parsing transaction data: %v", err)
		return
	}

	tx.FromIdx = fromIdx
	tx.ToIdx = toIdx
	tx.FromEthAddr = fromEthAddr.Bytes()
	tx.ToEthAddr = toEthAddr.Bytes()
	tx.Amount = amount
	tx.Type = txType
	tx.Position = eventData.LastAddedTxn
	tx.TxHash = vLog.TxHash.Bytes()

	s.logger.Printf("Processing missed transaction: %v", tx)

	// Add transaction to history DB
	err = s.AddTransactionToHistoryDB(tx)
	if err != nil {
		s.logger.Printf("Error adding missed transaction to history DB: %v", err)
		return
	}

	// Check if we need to forge a batch for this missed transaction
	// For missed events, we need to be more careful about batch forging
	// since the batch might have already been forged on-chain
	if tx.Position.Cmp(new(big.Int).Add(lastForgedTx, big.NewInt(int64(batchSize.Uint64())))) >= 0 {
		// Check if this batch was already forged
		currentLastForgedBatch, err := s.sybilContract.LastForgedBatch(&bind.CallOpts{Context: ctx})
		if err != nil {
			s.logger.Printf("Error checking current last forged batch: %v", err)
		} else if currentLastForgedBatch > uint32(lastForgedBatch.Uint64()) {
			s.logger.Printf("Batch %d was already forged on-chain, skipping", lastForgedBatch.Uint64()+1)
		} else {
			s.logger.Printf("Forging batch %d for missed transaction", lastForgedBatch.Uint64()+1)
			err = s.forger.ForgeBatch(uint32(lastForgedBatch.Uint64() + 1))
			if err != nil {
				s.logger.Printf("Error forging batch for missed transaction: %v", err)
			}
		}
	}

	s.logger.Printf("Successfully processed missed transaction: %v", vLog.TxHash.Hex())
}

// processLog processes a single log entry
func (s *Synchronizer) processLog(vLog types.Log) {
	ctx := context.Background()
	callOpts := &bind.CallOpts{
		Context: ctx,
	}

	s.logger.Printf("Processing log: BlockNumber=%d TxHash=%s", vLog.BlockNumber, vLog.TxHash.Hex())

	// // Parse the event
	eventData, eventType, err := ParseEvent(&vLog)
	if err != nil {
		s.logger.Printf("Error parsing event: %v", err)
		return
	}

	lastForgedTx, err := s.sybilContract.LastForgedTxn(callOpts)
	if err != nil {
		s.logger.Printf("Error fetching last forged transaction: %v", err)
	}

	batchSize, err := s.sybilContract.BatchSize(callOpts)
	if err != nil {
		s.logger.Printf("Error fetching batch size: %v", err)
	}

	lastForgedBatch, err := s.sybilContract.LastForgedBatch(callOpts)
	if err != nil {
		s.logger.Printf("Error fetching last forged batch: %v", err)
	}

	fmt.Printf("event: %v, eventType: %s, err: %v \n", eventData, eventType, err)

	tx := &common.Tx{
		BatchNum: uint32(lastForgedBatch + 1),
		Type:     eventType,
		FromIdx:  0,
		ToIdx:    0,
		Amount:   big.NewInt(0),
	}

	// Fetch Block Timestamp
	header, err := s.client.HeaderByNumber(ctx, new(big.Int).SetUint64(vLog.BlockNumber))
	if err != nil {
		s.logger.Printf("Error fetching block header for block %d: %v", vLog.BlockNumber, err)
	} else {
		tx.Timestamp = header.Time
	}

	tx.BlockNumber = vLog.BlockNumber

	// Fetch Transaction Receipt for Gas Fee
	receipt, err := s.client.TransactionReceipt(ctx, vLog.TxHash)
	if err != nil {
		s.logger.Printf("Error fetching transaction receipt for tx %s: %v", vLog.TxHash.Hex(), err)
	} else {
		if receipt.EffectiveGasPrice != nil && receipt.GasUsed > 0 {
			gasFee := new(big.Int).Mul(new(big.Int).SetUint64(receipt.GasUsed), receipt.EffectiveGasPrice)
			tx.GasFee = gasFee
		} else {
			s.logger.Printf("Could not calculate gas fee for tx %s: EffectiveGasPrice or GasUsed missing/zero. GasUsed: %d", vLog.TxHash.Hex(), receipt.GasUsed)
			tx.GasFee = big.NewInt(0)
		}
	}

	var sender ethCommon.Address
	txDetails, isPending, err := s.client.TransactionByHash(ctx, vLog.TxHash)
	if err != nil {
		s.logger.Printf("Error fetching full transaction details for tx %s: %v", vLog.TxHash.Hex(), err)
	} else if isPending {
		s.logger.Printf("Transaction %s is still pending, sender might not be final.", vLog.TxHash.Hex())
	} else if txDetails != nil {
		chainID, err := s.client.ChainID(ctx)
		if err != nil {
			s.logger.Printf("Error fetching chain ID for tx %s: %v", vLog.TxHash.Hex(), err)
		} else {
			signer := types.LatestSignerForChainID(chainID)
			sender, err = types.Sender(signer, txDetails)
			if err != nil {
				s.logger.Printf("Error deriving sender for transaction %s: %v", vLog.TxHash.Hex(), err)
			}
		}
	}

	// Format event data based on event type
	//TODO: Add Different events
	var eventDetails string
	switch e := eventData.(type) {
	case *bindings.SybilTxEvent:
		eventDetails = fmt.Sprintf(", Position: %d", e.LastAddedTxn)
		txType, fromEthAddr, toEthAddr, amount, fromIdx, toIdx, err := s.ParseTxData(e, sender)
		if err != nil {
			s.logger.Printf("Error parsing transaction data: %v", err)
		}
		tx.FromIdx = fromIdx
		tx.ToIdx = toIdx
		tx.FromEthAddr = fromEthAddr.Bytes()
		tx.ToEthAddr = toEthAddr.Bytes()
		tx.Amount = amount
		tx.Type = txType
		tx.Position = e.LastAddedTxn
		tx.TxHash = vLog.TxHash.Bytes()

		s.logger.Println("Transaction", tx)

	case *bindings.SybilForgeBatch:
		s.logger.Printf("Received SybilForgeBatch event, skipping transaction processing")
		return

	default:
		eventDetails = "Unknown event data"
	}

	s.logger.Printf("Event identified: %s, Data: %s", eventType, eventDetails)

	err = s.AddTransactionToHistoryDB(tx)
	if err != nil {
		s.logger.Fatalf("Error adding transaction to history DB: %v", err)
		return
	}

	if tx.Position.Cmp(new(big.Int).Add(lastForgedTx, big.NewInt(int64(batchSize.Uint64())))) >= 0 {
		err = s.forger.ForgeBatch(lastForgedBatch + 1)
		if err != nil {
			s.logger.Fatalf("Error forging batch: %v", err)
		}
	}

	s.logger.Printf("Saved transaction: %v, event: %v, data: %v",
		vLog.TxHash.Hex(), eventType, eventData)

	// Update the last processed block
	s.lastBlock = int64(vLog.BlockNumber)
	s.saveLastProcessedBlock(s.lastBlock)
}
