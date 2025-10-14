package synchronizer

import (
	"context"
	"fmt"
	"log"
	"math/big"
	"os"
	"time"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	ethCommon "github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/tokamak-network/syb-sequencer/sequencer/abis/bindings"
	"github.com/tokamak-network/syb-sequencer/sequencer/common"
	"github.com/tokamak-network/syb-sequencer/sequencer/config"
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
	txProcessor     *TxProcessor
	logs            chan types.Log
	sub             ethereum.Subscription
	logger          *log.Logger
	forger          *forger.Forger
	lastBlock       int64
	liveSync        bool
	batchTxToSync   int64
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

	batchTxToSync := config.GetEnvInt64("BATCH_TO_SYNC", 100)

	txProcessor := NewTxProcessor(statedb)

	logs := make(chan types.Log)

	return &Synchronizer{
		client:          client,
		contractAddress: contractAddress,
		sybilContract:   sybilContract,
		historydb:       historydb,
		statedb:         statedb,
		txProcessor:     txProcessor,
		logs:            logs,
		logger:          logger,
		forger:          forger,
		lastBlock:       config.GetEnvInt64("LAST_PROCESSED_BLOCK", 0),
		liveSync:        false,
		batchTxToSync:   batchTxToSync,
	}, nil
}

// Start begins the synchronization process
func (s *Synchronizer) Start(ctx context.Context) {
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
	ticker := time.NewTicker(10 * time.Second)
	defer ticker.Stop()

	s.logger.Println("Started watching for contract events")
	fmt.Println(s.liveSync)

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
			if s.liveSync {
				s.processLog(vLog)
			}

		case <-ticker.C:
			// Periodic safety check to ensure we haven't missed any events
			// This is just a backup mechanism and not the primary way of getting events
			if !s.liveSync {
				s.logger.Println("Performing periodic safety check for missed events")
				s.checkForMissedEvents(ctx)
			}
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
	var blockNumberToSync int64
	var startLiveSync bool
	// Get the latest block number
	header, err := s.client.HeaderByNumber(ctx, nil)
	if err != nil {
		s.logger.Printf("Error getting latest block header: %v", err)
		return
	}

	latestBlock := header.Number.Int64()

	if s.lastBlock+s.batchTxToSync > latestBlock {
		blockNumberToSync = latestBlock
		startLiveSync = true
	} else {
		blockNumberToSync = s.lastBlock + s.batchTxToSync
	}

	s.logger.Printf("Checking for missed events from block %d to %d", s.lastBlock, blockNumberToSync)

	// Create a filter query for the contract events
	query := ethereum.FilterQuery{
		FromBlock: big.NewInt(s.lastBlock),
		ToBlock:   big.NewInt(blockNumberToSync),
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
		s.processLog(vLog)
	}

	fmt.Println("Finished processing missed events", latestBlock)

	s.lastBlock = blockNumberToSync + 1

	if latestBlock <= s.lastBlock || startLiveSync {
		s.liveSync = true
		s.logger.Printf("Started syncing live events")
		return // No new blocks
	}
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
	if eventType == "Initialized" {
		return
	}
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
		txType, fromEthAddr, toEthAddr, amount, fromIdx, toIdx, err := s.ParseL1UserTxData(e, sender)
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
		eventDetails = fmt.Sprintf(", Position: %d", e.LastForgedTxn)
		_, _, _, txs, err := s.ParseForgeBatchTxData(e)
		if err != nil {
			s.logger.Printf("Error parsing transaction data: %v", err)
		}

		for _, txn := range txs {
			s.txProcessor.ProcessTx(txn)
		}
		tx.Type = common.TxTypeForgeBatch
		tx.BatchNum = e.LastForgedBatch
		tx.TxHash = vLog.TxHash.Bytes()
		privateKeyHex := os.Getenv("PRIVATE_KEY")
		fromEthAddr, err := common.AddressFromPrivKeyHex(privateKeyHex)
		if err != nil {
			s.logger.Printf("Error getting address from private key: %v", err)
		}
		tx.FromEthAddr = fromEthAddr.Bytes()

	case *bindings.SybilProveScore:
		eventDetails = fmt.Sprintf(", Score: %d, BatchNum: %d", e.Score, e.BatchNum)
		tx.Type = common.TxTypeProveScore
		tx.FromEthAddr = e.User.Bytes()
		tx.BatchNum = e.BatchNum
		tx.TxHash = vLog.TxHash.Bytes()

	default:
		eventDetails = "Unknown event data"
	}

	s.logger.Printf("Event identified: %s, Data: %s", eventType, eventDetails)

	err = s.AddTransactionToHistoryDB(tx)
	if err != nil {
		s.logger.Fatalf("Error adding transaction to history DB: %v", err)
		return
	}

	// Check if eventData is of type SybilTxEvent
	if _, ok := eventData.(*bindings.SybilTxEvent); ok {
		// Check if we need to forge a new batch
		if tx.Position.Cmp(new(big.Int).Add(lastForgedTx, big.NewInt(int64(batchSize.Uint64())))) >= 0 {
			err = s.forger.ForgeBatch(lastForgedBatch + 1)
			if err != nil {
				s.logger.Fatalf("Error forging batch: %v", err)
			}
		}
	}

	s.logger.Printf("Saved transaction: %v, event: %v, data: %v",
		vLog.TxHash.Hex(), eventType, eventData)
}
