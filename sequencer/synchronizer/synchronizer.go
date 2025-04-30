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
			Addresses: []common.Address{s.contractAddress},
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

// checkForMissedEvents is a safety mechanism to check for any events we might have missed
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

	// Process any missed events
	for _, vLog := range logs {
		s.logger.Printf("Processing missed event from block %d, tx: %s", vLog.BlockNumber, vLog.TxHash.Hex())
		s.processLog(vLog)
	}

	s.lastBlock = latestBlock
}

// processLog processes a single log entry
func (s *Synchronizer) processLog(vLog types.Log) {
	s.logger.Printf("Processing log: BlockNumber=%d TxHash=%s", vLog.BlockNumber, vLog.TxHash.Hex())
}
