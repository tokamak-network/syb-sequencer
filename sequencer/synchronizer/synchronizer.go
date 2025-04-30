package synchronizer

import (
	"context"
	"fmt"
	"log"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/tokamak-network/syb-sequencer/sequencer/abis/bindings"
	"github.com/tokamak-network/syb-sequencer/sequencer/db/historydb"
)

// Synchronizer listens to contract events and stores them in the database
type Synchronizer struct {
	client          *ethclient.Client
	contractAddress common.Address
	sybilContract   *bindings.Bindings
	db              *historydb.HistoryDB
	logs            chan types.Log
	sub             ethereum.Subscription
	logger          *log.Logger
	lastBlock       int64
}

// NewSynchronizer creates a new synchronizer
func NewSynchronizer(ethRPC, contractAddressHex string, db *historydb.HistoryDB, logger *log.Logger) (*Synchronizer, error) {
	client, err := ethclient.Dial(ethRPC)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to Ethereum client: %v", err)
	}

	contractAddress := common.HexToAddress(contractAddressHex)
	sybilContract, err := bindings.NewBindings(contractAddress, client)
	if err != nil {
		return nil, fmt.Errorf("failed to instantiate Sybil contract: %v", err)
	}

	logs := make(chan types.Log)

	return &Synchronizer{
		client:          client,
		contractAddress: contractAddress,
		sybilContract:   sybilContract,
		db:              db,
		logs:            logs,
		logger:          logger,
		lastBlock:       0,
	}, nil
}

// Start begins the synchronization process
func (s *Synchronizer) Start(ctx context.Context) {
	// Get the latest block number to start from
	header, err := s.client.HeaderByNumber(ctx, nil)
	if err != nil {
		s.logger.Fatalf("Failed to get latest block header: %v", err)
	}

	s.lastBlock = header.Number.Int64()
	s.logger.Printf("Starting synchronizer from block %d", s.lastBlock)

	// Create a filter query for the contract events
	query := ethereum.FilterQuery{
		Addresses: []common.Address{s.contractAddress},
	}

	// Subscribe to logs
	sub, err := s.client.SubscribeFilterLogs(ctx, query, s.logs)
	if err != nil {
		s.logger.Fatalf("Failed to subscribe to logs: %v", err)
	}
	s.sub = sub

	// Start listening for events
	go s.watchEvents(ctx)
}

// watchEvents continuously listens for contract events
func (s *Synchronizer) watchEvents(ctx context.Context) {
	// Create a ticker for periodic checks
	ticker := time.NewTicker(15 * time.Second)
	defer ticker.Stop()

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
			// Process the log
			s.processLog(vLog)

		case <-ticker.C:
			// Periodic check for new blocks
			s.checkNewBlocks(ctx)
		}
	}
}

// resubscribe attempts to reestablish the subscription
func (s *Synchronizer) resubscribe(ctx context.Context) {
	for {
		s.logger.Println("Attempting to resubscribe...")
		query := ethereum.FilterQuery{
			Addresses: []common.Address{s.contractAddress},
		}

		sub, err := s.client.SubscribeFilterLogs(ctx, query, s.logs)
		if err != nil {
			s.logger.Printf("Failed to resubscribe: %v, retrying in 10 seconds", err)
			time.Sleep(10 * time.Second)
			continue
		}

		s.sub = sub
		s.logger.Println("Successfully resubscribed")
		return
	}
}

// checkNewBlocks checks for new blocks and processes any missed events
func (s *Synchronizer) checkNewBlocks(ctx context.Context) {
	header, err := s.client.HeaderByNumber(ctx, nil)
	if err != nil {
		s.logger.Printf("Error getting latest block header: %v", err)
		return
	}

	latestBlock := header.Number.Int64()
	if latestBlock <= s.lastBlock {
		return // No new blocks
	}

	s.logger.Printf("Processing blocks from %d to %d", s.lastBlock+1, latestBlock)

	// Create a filter query for the contract events
	query := ethereum.FilterQuery{
		FromBlock: big.NewInt(s.lastBlock + 1),
		ToBlock:   big.NewInt(latestBlock),
		Addresses: []common.Address{s.contractAddress},
	}

	// Get logs matching the filter
	logs, err := s.client.FilterLogs(ctx, query)
	if err != nil {
		s.logger.Printf("Error filtering logs: %v", err)
		return
	}

	// Process each log
	for _, vLog := range logs {
		s.processLog(vLog)
	}

	s.lastBlock = latestBlock
}

// processLog processes a single log entry
func (s *Synchronizer) processLog(vLog types.Log) {
	s.logger.Printf("Processing log: BlockNumber=%d TxHash=%s", vLog.BlockNumber, vLog.TxHash.Hex())

	// Try to identify the event type based on the topics
	eventName := "Unknown"
	var eventData string

	// Check for ForgeBatch event
	if len(vLog.Topics) > 0 && vLog.Topics[0] == common.HexToHash("0x5838115dee16474c1ff6e07c103e24fac5f63a1d9e7c3fd28c3e75c2b42f32a0") {
		eventName = "ForgeBatch"
		forgeBatch, err := s.sybilContract.ParseForgeBatch(vLog)
		if err == nil {
			eventData = fmt.Sprintf("BatchNum: %d, L1UserTxsLen: %d", forgeBatch.BatchNum, forgeBatch.L1UserTxsLen)
		}
	}

	// Check for L1UserTxEvent
	if len(vLog.Topics) > 0 && vLog.Topics[0] == common.HexToHash("0x8b2a1e1a3bda2e8fbd2af44c30f3a9e0e90d7af2a7a7b9c2c5e1c8ce2db047f3") {
		eventName = "L1UserTxEvent"
		l1UserTxEvent, err := s.sybilContract.ParseL1UserTxEvent(vLog)
		if err == nil {
			eventData = fmt.Sprintf("QueueIndex: %d, Position: %d", l1UserTxEvent.QueueIndex, l1UserTxEvent.Position)
		}
	}

	// Check for WithdrawEvent
	if len(vLog.Topics) > 0 && vLog.Topics[0] == common.HexToHash("0x8c1d8f1f64246a3a5f7bd7ddc070d0c8b98a8fec8fd11ee4e2a41bd3f1a5f865") {
		eventName = "WithdrawEvent"
		withdrawEvent, err := s.sybilContract.ParseWithdrawEvent(vLog)
		if err == nil {
			eventData = fmt.Sprintf("Idx: %d, NumExitRoot: %d", withdrawEvent.Idx, withdrawEvent.NumExitRoot)
		}
	}

	// // Save the event to the database
	// tx := &historydb.Transaction{
	// 	TxHash:      vLog.TxHash.Hex(),
	// 	BlockNumber: int64(vLog.BlockNumber),
	// 	EventName:   eventName,
	// 	EventData:   eventData,
	// 	CreatedAt:   time.Now(),
	// }

	// err := s.db.SaveTransaction(tx)
	// if err != nil {
	// 	s.logger.Printf("Error saving transaction: %v", err)
	// 	return
	// }

	s.logger.Printf("Saved transaction: %s, event: %s, data: %s",
		vLog.TxHash.Hex(), eventName, eventData)
}
