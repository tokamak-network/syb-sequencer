package synchronizer

import (
	"context"
	"fmt"
	"log"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum"
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
	sybilContract   *bindings.Bindings
	historydb       *historydb.HistoryDB
	statedb         *statedb.StateDB
	logs            chan types.Log
	sub             ethereum.Subscription
	logger          *log.Logger
	forger          *forger.Forger
	lastBlock       int64
}

var lastSyncBatch uint32

// NewSynchronizer creates a new synchronizer
func NewSynchronizer(ethRPC, contractAddressHex string, historydb *historydb.HistoryDB, statedb *statedb.StateDB, logger *log.Logger, forger *forger.Forger) (*Synchronizer, error) {
	client, err := ethclient.Dial(ethRPC)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to Ethereum client: %v", err)
	}

	contractAddress := ethCommon.HexToAddress(contractAddressHex)
	sybilContract, err := bindings.NewBindings(contractAddress, client)
	if err != nil {
		return nil, fmt.Errorf("failed to instantiate Sybil contract: %v", err)
	}

	logs := make(chan types.Log)

	return &Synchronizer{
		client:          client,
		contractAddress: contractAddress,
		sybilContract:   sybilContract,
		historydb:       historydb,
		statedb:         statedb,
		logs:            logs,
		logger:          logger,
		forger:          forger,
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
		Addresses: []ethCommon.Address{s.contractAddress},
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
		// TODO: Logic to only update the missed events
		// s.processLog(vLog)
	}

	s.lastBlock = latestBlock
}

// processLog processes a single log entry
func (s *Synchronizer) processLog(vLog types.Log) {
	lastForgedBatch := s.forger.GetLastForgedBatchNum()
	s.logger.Printf("Processing log: BlockNumber=%d TxHash=%s", vLog.BlockNumber, vLog.TxHash.Hex())

	// // Parse the event
	eventData, eventType, err := ParseEvent(&vLog)
	if err != nil {
		s.logger.Printf("Error parsing event: %v", err)
		return
	}

	fmt.Printf("event: %v, eventType: %s, err: %v \n", eventData, eventType, err)

	tx := &common.Tx{
		BatchNum: int64(eventData.QueueIndex), // Initial batch number is 0
		Position: int(eventData.Position),
		Type:     eventType,
		FromIdx:  0,             // Will be set based on event type
		ToIdx:    0,             // Will be set based on event type
		Amount:   big.NewInt(0), // Will be set based on event type
	}

	// Fetch Block Timestamp
	header, err := s.client.HeaderByNumber(context.Background(), new(big.Int).SetUint64(vLog.BlockNumber))
	if err != nil {
		s.logger.Printf("Error fetching block header for block %d: %v", vLog.BlockNumber, err)
	} else {
		tx.Timestamp = header.Time
	}

	tx.BlockNumber = vLog.BlockNumber

	// Fetch Transaction Receipt for Gas Fee
	receipt, err := s.client.TransactionReceipt(context.Background(), vLog.TxHash)
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

	// Format event data based on event type
	//TODO: Add Different events
	var eventDetails string
	switch eventType {
	case "L1UserTxEvent":
		eventDetails = fmt.Sprintf("QueueIndex: %d, Position: %d",
			eventData.QueueIndex, eventData.Position)
		txType, fromEthAddr, toEthAddr, amount, err := ParseTxData(eventData.L1UserTx)
		if err != nil {
			s.logger.Printf("Error parsing transaction data: %v", err)
		}

		tx.Type = txType

		fromIdx, err := s.statedb.GetAccountIdxByEthAddr(fromEthAddr)
		if err != nil {
			if fromIdx == common.AccountIdx(0) {
				tx.Type = common.TxTypeCreateAccountDeposit
			} else {
				s.logger.Printf("Failed to get idx from ethAddr: %v", err)
			}
		}
		tx.FromIdx = fromIdx

		if tx.Type == common.TxTypeVouch ||
			tx.Type == common.TxTypeUnvouch ||
			tx.Type == common.TxTypeExplode {
			toIdx, err := s.statedb.GetAccountIdxByEthAddr(tx.ToEthAddr)
			if err != nil {
				s.logger.Printf("Failed to get idx from ethAddr: %v", err)
			}
			tx.ToIdx = toIdx
		}

		tx.FromEthAddr = fromEthAddr
		tx.ToEthAddr = toEthAddr
		tx.Amount = amount

		s.logger.Println("Transaction", tx)

	default:
		eventDetails = "Unknown event data"
	}

	s.logger.Printf("Event identified: %s, Data: %s", eventType, eventDetails)

	// Process tx
	tp := NewTxProcessor(s.statedb)

	pOut, err := tp.ProcessTxs(*tx)
	if err != nil {
		s.logger.Printf("Error processing transaction: %v", err)
		return
	}

	s.logger.Printf("Transaction processed: %s", pOut)

	err = s.historydb.SaveTx(tx)
	if err != nil {
		s.logger.Printf("Error saving transaction: %v", err)
		return
	}
	// Check for current batch to be synced and update it's number
	if lastSyncBatch == 0 || eventData.QueueIndex > lastSyncBatch {
		lastSyncBatch = eventData.QueueIndex
	}

	if lastForgedBatch < (lastSyncBatch + 2) {
		err = s.forger.ForgeBatch(lastForgedBatch + 1)
		if err != nil {
			s.logger.Printf("Error forging batch: %v", err)
		}
	}

	s.logger.Printf("Saved transaction: %v, event: %v, data: %v",
		vLog.TxHash.Hex(), eventType, eventData)
}
