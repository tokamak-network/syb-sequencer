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
		FromIdx:  nil,           // Will be set based on event type
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
		txType, fromEthAddr, toEthAddr, amount, fromIdx, toIdx, err := s.ParseTxData(eventData)
		if err != nil {
			s.logger.Printf("Error parsing transaction data: %v", err)
		}
		tx.FromIdx = &fromIdx
		tx.ToIdx = toIdx
		tx.FromEthAddr = fromEthAddr.Bytes()
		tx.ToEthAddr = toEthAddr.Bytes()
		tx.Amount = amount
		tx.Type = txType

	default:
		eventDetails = "Unknown event data"
	}

	s.logger.Printf("Event identified: %s, Data: %s", eventType, eventDetails)

	err = s.AddTransactionToHistoryDB(tx)
	if err != nil {
		s.logger.Fatalf("Error adding transaction to history DB: %v", err)
		return
	}

	// Check for current batch to be synced and update it's number
	if lastSyncBatch == 0 || eventData.QueueIndex > lastSyncBatch {
		lastSyncBatch = eventData.QueueIndex
	}

	if lastForgedBatch < (lastSyncBatch + 2) {
		err = s.forger.ForgeBatch(lastForgedBatch + 1)
		if err != nil {
			s.logger.Fatalf("Error forging batch: %v", err)
		}
	}

	s.logger.Printf("Saved transaction: %v, event: %v, data: %v",
		vLog.TxHash.Hex(), eventType, eventData)
}

// // // Placeholder for L1UserTx parsing logic
// // // You'll need to implement this based on your L1UserTx byte structure
// // type ParsedL1TxData struct {
// // 	TxTypeByte  byte
// // 	FromEthAddr ethCommon.Address
// // 	ToEthAddr   ethCommon.Address // May be zero address for some types
// // 	Amount      *big.Int
// // 	// ... other relevant fields from L1UserTx ...
// // }

// // // func parseL1UserTx(txBytes []byte) (*ParsedL1TxData, error) {
// // // 	if len(txBytes) < 32 { // Minimum length to get type byte
// // // 		return nil, fmt.Errorf("L1UserTx too short")
// // // 	}
// // // 	// THIS IS A SIMPLIFIED PARSING LOGIC - REPLACE WITH YOUR ACTUAL PARSING
// // // 	data := &ParsedL1TxData{}
// // // 	data.TxTypeByte = txBytes[31] // As per your SetType usage

// // // 	// Example: Extract FromEthAddr (bytes 0-19)
// // // 	// if len(txBytes) >= 20 {
// // // 	// 	data.FromEthAddr = ethCommon.BytesToAddress(txBytes[0:20])
// // // 	// }
// // // 	// Example: Extract ToEthAddr (bytes 20-39, if applicable)
// // // 	// if len(txBytes) >= 40 {
// // // 	//  data.ToEthAddr = ethCommon.BytesToAddress(txBytes[20:40])
// // // 	// }
// // // 	// Example: Extract Amount (e.g., bytes 40-71 for a 32-byte big.Int)
// // // 	// if len(txBytes) >= 72 {
// // // 	// 	data.Amount = new(big.Int).SetBytes(txBytes[40:72])
// // // 	// } else {
// // // 	//  data.Amount = big.NewInt(0)
// // // 	// }
// // // 	// Replace above with your actual parsing logic for FromEthAddr, ToEthAddr, Amount
// // // 	// For now, let's assume they are populated somehow for the example to proceed.
// // // 	// For testing, you might hardcode or use a more detailed mock:
// // // 	if len(txBytes) > 20 { // Placeholder
// // // 		data.FromEthAddr = ethCommon.BytesToAddress(txBytes[0:20]) // Placeholder
// // // 	}
// // // 	if len(txBytes) > 40 { // Placeholder
// // // 		data.ToEthAddr = ethCommon.BytesToAddress(txBytes[20:40]) // Placeholder
// // // 	}
// // // 	data.Amount = big.NewInt(100) // Placeholder amount

// // // 	return data, nil
// // // }

// // Inside your event processing loop (e.g., ProcessEvents method)
// // for _, eventData := range eventsData { // eventsData from L1
// // ...
// 		// Assume eventData is your L1UserTxEvent struct
// 		// eventData.L1UserTx []byte
// 		// eventData.Position uint8
// 		// eventData.QueueIndex uint32

// 		parsedL1Data, err := parseL1UserTx(eventData.L1UserTx)
// 		if err != nil {
// 			log.Errorf("Failed to parse L1UserTx: %v for queue index %d", err, eventData.QueueIndex)
// 			continue
// 		}

// 		txTypeString, err := SetType(parsedL1Data.TxTypeByte)
// 		if err != nil {
// 			log.Errorf("Failed to determine transaction type for L1UserTx[31]=%d: %v", parsedL1Data.TxTypeByte, err)
// 			continue
// 		}

// 		// This is the common.Tx object that will be saved by s.historydb.SaveTx()
// 		// Ensure its fields match what historydb.SaveTx expects and your 'tx' table schema.
// 		historyTx := &common.Tx{ // Assuming common.Tx is the type for history DB's tx table
// 			// BatchNum:    currentBatch.Number, // Needs to be set
// 			Position:    int(eventData.Position), // SQL 'position' is INT
// 			Type:        txTypeString,
// 			FromEthAddr: parsedL1Data.FromEthAddr,
// 			ToEthAddr:   parsedL1Data.ToEthAddr,
// 			Amount:      parsedL1Data.Amount,
// 			// BlockNumber, TxTimestamp, GasFee might be set later or from eventData if available
// 		}

// 		var fromAccount *historydb.Account
// 		var toAccount *historydb.Account // For vouch operations

// 		// 1. Determine FromIdx (Account Index)
// 		if txTypeString == TxTypeCreateAccountDeposit {
// 			// For new accounts, Idx is derived from eventData.Position
// 			newAccountIdx := uint32(eventData.Position)
// 			historyTx.FromIdx = &newAccountIdx // common.Tx.FromIdx is *uint32
// 			// ToIdx for CreateAccountDeposit is typically the same as FromIdx
// 			historyTx.ToIdx = newAccountIdx // common.Tx.ToIdx is uint32
// 		} else {
// 			// For existing accounts, fetch Idx using FromEthAddr
// 			if parsedL1Data.FromEthAddr != (ethCommon.Address{}) {
// 				fromAccount, err = s.historydb.GetAccountByEthAddress(parsedL1Data.FromEthAddr)
// 				if err != nil {
// 					log.Errorf("Error fetching from_account by eth_addr %s for tx type %s: %v", parsedL1Data.FromEthAddr.Hex(), txTypeString, err)
// 					continue // Or handle more gracefully
// 				}
// 				if fromAccount == nil && txTypeString != TxTypeCreateAccountDeposit { // Should not happen for non-create types
// 					log.Errorf("From_account not found for eth_addr %s, required for tx type %s", parsedL1Data.FromEthAddr.Hex(), txTypeString)
// 					continue
// 				}
// 				if fromAccount != nil {
// 					historyTx.FromIdx = &fromAccount.Idx
// 				}
// 			} else if txTypeString != TxTypeCreateAccountDeposit { // Some tx types might not have a from_address
// 				log.Warnf("Tx type %s does not have a FromEthAddr, FromIdx will be nil", txTypeString)
// 			}
// 		}

// 		// 1.1 Determine ToIdx (Account Index) for relevant transactions
// 		// This is simplified; you'll need to ensure ToEthAddr is correctly parsed and handled for each tx type.
// 		if txTypeString == TxTypeCreateVouch || txTypeString == TxTypeUnvouch {
// 			if parsedL1Data.ToEthAddr == (ethCommon.Address{}) {
// 				log.Errorf("ToEthAddr is required for %s but is zero", txTypeString)
// 				continue
// 			}
// 			toAccount, err = s.historydb.GetAccountByEthAddress(parsedL1Data.ToEthAddr)
// 			if err != nil {
// 				log.Errorf("Error fetching to_account by eth_addr %s for tx type %s: %v", parsedL1Data.ToEthAddr.Hex(), txTypeString, err)
// 				continue
// 			}
// 			if toAccount == nil {
// 				log.Errorf("To_account not found for eth_addr %s, required for tx type %s", parsedL1Data.ToEthAddr.Hex(), txTypeString)
// 				continue
// 			}
// 			historyTx.ToIdx = toAccount.Idx // common.Tx.ToIdx is uint32
// 		} else if txTypeString == TxTypeDeposit && historyTx.FromIdx != nil {
// 			// For self-deposit, ToIdx can be the same as FromIdx
// 			historyTx.ToIdx = *historyTx.FromIdx
// 		}
// 		// For other types like Withdraw, ToIdx might be 0 or nil if not applicable.
// 		// Ensure common.Tx.ToIdx (and its DB column) handles this (e.g., nullable or default 0).

// 		// 2. Perform DB operations on `account` or `vouch` tables BEFORE SaveTx
// 		switch txTypeString {
// 		case TxTypeCreateAccountDeposit:
// 			if historyTx.FromIdx == nil { // Should be set above
// 				log.Errorf("Cannot create account: FromIdx is nil. L1UserTx: %x", eventData.L1UserTx)
// 				continue
// 			}
// 			accToCreate := &historydb.Account{
// 				Idx:           *historyTx.FromIdx,
// 				EthAddr:       parsedL1Data.FromEthAddr, // This is the new account's EthAddr
// 				Balance:       parsedL1Data.Amount,
// 				Score:         big.NewInt(0), // Initial score
// 				ScoreSiblings: []*big.Int{},  // Initial empty siblings
// 			}
// 			err = s.historydb.AddAccount(accToCreate)
// 			if err != nil {
// 				log.Errorf("Failed to add account (idx %d, eth %s) for CreateAccountDeposit: %v", accToCreate.Idx, accToCreate.EthAddr.Hex(), err)
// 				continue
// 			}
// 			log.Infof("Account created: Idx %d, EthAddr %s", accToCreate.Idx, accToCreate.EthAddr.Hex())

// 		case TxTypeDeposit:
// 			if fromAccount == nil { // Should be fetched above
// 				log.Errorf("Cannot process deposit: from_account not found for eth_addr %s. L1UserTx: %x", parsedL1Data.FromEthAddr.Hex(), eventData.L1UserTx)
// 				continue
// 			}
// 			newBalance := new(big.Int).Add(fromAccount.Balance, parsedL1Data.Amount)
// 			err = s.historydb.UpdateAccountBalance(fromAccount.Idx, newBalance)
// 			if err != nil {
// 				log.Errorf("Failed to update account balance for Deposit (idx %d): %v", fromAccount.Idx, err)
// 				continue
// 			}
// 			log.Infof("Account %d balance updated for deposit. New balance: %s", fromAccount.Idx, newBalance.String())

// 		case TxTypeWithdraw:
// 			if fromAccount == nil {
// 				log.Errorf("Cannot process withdraw: from_account not found for eth_addr %s. L1UserTx: %x", parsedL1Data.FromEthAddr.Hex(), eventData.L1UserTx)
// 				continue
// 			}
// 			newBalance := new(big.Int).Sub(fromAccount.Balance, parsedL1Data.Amount)
// 			if newBalance.Sign() < 0 {
// 				log.Errorf("Insufficient balance for Withdraw tx for account idx %d (eth %s). Current: %s, Amount: %s",
// 					fromAccount.Idx, fromAccount.EthAddr.Hex(), fromAccount.Balance.String(), parsedL1Data.Amount.String())
// 				continue
// 			}
// 			err = s.historydb.UpdateAccountBalance(fromAccount.Idx, newBalance)
// 			if err != nil {
// 				log.Errorf("Failed to update account balance for Withdraw (idx %d): %v", fromAccount.Idx, err)
// 				continue
// 			}
// 			log.Infof("Account %d balance updated for withdraw. New balance: %s", fromAccount.Idx, newBalance.String())

// 		case TxTypeCreateVouch:
// 			if fromAccount == nil || toAccount == nil {
// 				log.Errorf("Cannot create vouch: from_account (eth %s) or to_account (eth %s) not found. L1UserTx: %x",
// 					parsedL1Data.FromEthAddr.Hex(), parsedL1Data.ToEthAddr.Hex(), eventData.L1UserTx)
// 				continue
// 			}
// 			vouchTableIdxStr := fmt.Sprintf("%d%d", fromAccount.Idx, toAccount.Idx)
// 			vouchTableIdxBigInt, ok := new(big.Int).SetString(vouchTableIdxStr, 10)
// 			if !ok {
// 				log.Errorf("Failed to create vouch table idx from string '%s' for CreateVouch", vouchTableIdxStr)
// 				continue
// 			}
// 			vouchToCreate := &historydb.Vouch{
// 				Idx:         vouchTableIdxBigInt.Uint64(),
// 				FromIdx:     fromAccount.Idx,
// 				FromEthAddr: fromAccount.EthAddr,
// 				ToIdx:       toAccount.Idx,
// 				ToEthAddr:   toAccount.EthAddr,
// 			}
// 			err = s.historydb.AddVouch(vouchToCreate)
// 			if err != nil {
// 				log.Errorf("Failed to add vouch (idx %d, from %d to %d): %v",
// 					vouchToCreate.Idx, vouchToCreate.FromIdx, vouchToCreate.ToIdx, err)
// 				continue
// 			}
// 			log.Infof("Vouch created: Idx %d (From Acct %d -> To Acct %d)", vouchToCreate.Idx, fromAccount.Idx, toAccount.Idx)

// 		case TxTypeUnvouch:
// 			if fromAccount == nil || toAccount == nil {
// 				log.Errorf("Cannot unvouch: from_account (eth %s) or to_account (eth %s) not found. L1UserTx: %x",
// 					parsedL1Data.FromEthAddr.Hex(), parsedL1Data.ToEthAddr.Hex(), eventData.L1UserTx)
// 				continue
// 			}
// 			vouchTableIdxStr := fmt.Sprintf("%d%d", fromAccount.Idx, toAccount.Idx)
// 			vouchTableIdxBigInt, ok := new(big.Int).SetString(vouchTableIdxStr, 10)
// 			if !ok {
// 				log.Errorf("Failed to create vouch table idx from string '%s' for Unvouch", vouchTableIdxStr)
// 				continue
// 			}
// 			err = s.historydb.DeleteVouchByIdx(vouchTableIdxBigInt.Uint64())
// 			if err != nil {
// 				log.Errorf("Failed to delete vouch (table_idx %s, from %d to %d): %v",
// 					vouchTableIdxStr, fromAccount.Idx, toAccount.Idx, err)
// 				continue
// 			}
// 			log.Infof("Vouch deleted: Table Idx %s (From Acct %d -> To Acct %d)", vouchTableIdxStr, fromAccount.Idx, toAccount.Idx)

// 		case TxTypeExplode:
// 			// Implement logic for TxTypeExplode if it affects account/vouch tables
// 			if fromAccount != nil {
// 				log.Infof("Processing Explode transaction for account Idx %d (Eth %s)", fromAccount.Idx, fromAccount.EthAddr.Hex())
// 			} else {
// 				log.Infof("Processing Explode transaction with no specific from_account found based on FromEthAddr %s", parsedL1Data.FromEthAddr.Hex())
// 			}
// 			// Add specific DB interactions here if needed

// 		default:
// 			log.Warnf("Transaction type %s has no specific pre-save logic in synchronizer.", txTypeString)
// 		}

// 		// 3. Finally, save the transaction to the main `tx` history table
// 		// Ensure historyTx has all necessary fields populated (BatchNum, BlockNumber, Timestamp etc.)
// 		// err = s.historydb.SaveTx(historyTx) // Assuming SaveTx exists and takes *common.Tx
// 		// if err != nil {
// 		// 	log.Errorf("Failed to save transaction to history DB (type %s, from_idx %v, to_idx %v): %v",
// 		//		historyTx.Type, historyTx.FromIdx, historyTx.ToIdx, err)
// 		// 	continue
// 		// }
// 		// log.Infof("Transaction %s saved to history DB.", historyTx.Type)
// // ...
// // } // end of event processing loop
