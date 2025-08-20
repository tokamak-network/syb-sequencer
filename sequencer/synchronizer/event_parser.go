package synchronizer

import (
	"fmt"
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	ethCommon "github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/tokamak-network/syb-sequencer/sequencer/abis/bindings"
	"github.com/tokamak-network/syb-sequencer/sequencer/common"
)

var (
	logSybilTxEvent = crypto.Keccak256Hash([]byte(
		"TxEvent(uint256,uint8,uint24,uint24,uint256)"))
	logSybilForgeBatch = crypto.Keccak256Hash([]byte(
		"ForgeBatch(uint32,uint256,uint256,bytes)"))
	logSybilExplodeAmountUpdated = crypto.Keccak256Hash([]byte(
		"ExplodeAmountUpdated(uint256)"))
	logSybilScoringRequiredBalanceUpdated = crypto.Keccak256Hash([]byte(
		"ScoringRequiredBalanceUpdated(uint256)"))
)

// ParseEvent parses data of events found on blockchain.
// It now returns an interface{} for the parsed event data, as it can be one of several types.
// The string return value is the name of the event.
func ParseEvent(log *types.Log) (interface{}, string, error) {
	sybilABI, err := abi.JSON(strings.NewReader(bindings.SybilMetaData.ABI))
	if err != nil {
		return nil, "", fmt.Errorf("ParseEvent: failed to parse ABI: %w", err)
	}

	// Check which event it is and unpack accordingly
	switch log.Topics[0].Hex() {
	case logSybilTxEvent.Hex():
		var parsedEvent bindings.SybilTxEvent
		err := sybilABI.UnpackIntoInterface(&parsedEvent, "TxEvent", log.Data)
		if err != nil {
			return nil, "", fmt.Errorf("ParseEvent: failed to unpack SybilTxEvent: %w", err)
		}
		lastAddedTxn := new(big.Int).SetBytes(log.Topics[1].Bytes())
		identifier := log.Topics[2].Bytes()[31]
		parsedEvent.LastAddedTxn = lastAddedTxn
		parsedEvent.Identifier = identifier
		return &parsedEvent, "SybilTxEvent", nil

	case logSybilForgeBatch.Hex():
		var parsedEvent bindings.SybilForgeBatch
		err := sybilABI.UnpackIntoInterface(&parsedEvent, "ForgeBatch", log.Data)
		if err != nil {
			return nil, "", fmt.Errorf("ParseEvent: failed to unpack SybilForgeBatch: %w", err)
		}
		return &parsedEvent, "SybilForgeBatch", nil

	case logSybilExplodeAmountUpdated.Hex():
		var parsedEvent bindings.SybilExplodeAmountUpdated
		err := sybilABI.UnpackIntoInterface(&parsedEvent, "ExplodeAmountUpdated", log.Data)
		if err != nil {
			return nil, "", fmt.Errorf("ParseEvent: failed to unpack SybilExplodeAmountUpdated: %w", err)
		}
		return &parsedEvent, "SybilExplodeAmountUpdated", nil

	case logSybilScoringRequiredBalanceUpdated.Hex():
		var parsedEvent bindings.SybilScoringRequiredBalanceUpdated
		err := sybilABI.UnpackIntoInterface(&parsedEvent, "ScoringRequiredBalanceUpdated", log.Data)
		if err != nil {
			return nil, "", fmt.Errorf("ParseEvent: failed to unpack SybilScoringRequiredBalanceUpdated: %w", err)
		}
		return &parsedEvent, "SybilScoringRequiredBalanceUpdated", nil
	}

	// If we get here, it's an unknown event
	return nil, fmt.Sprintf("UnknownEvent(%s)", log.Topics[0].Hex()[:10]), nil
}

// ParseL1UserTxData parses the transaction data from L1UserTx
// The data format is:
// - txnType (1 byte)
// - fromEthAddress (20 bytes)
// - toEthAddress (20 bytes)
// - amount (32 bytes)
func (s *Synchronizer) ParseL1UserTxData(eventData *bindings.SybilTxEvent, sender ethCommon.Address) (string, ethCommon.Address, ethCommon.Address, *big.Int, common.AccountIdx, common.AccountIdx, error) {
	var fromIdx common.AccountIdx
	var toIdx common.AccountIdx
	var toEthAddr ethCommon.Address

	txType, err := SetType(eventData.Identifier)
	if err != nil {
		return "", ethCommon.Address{}, ethCommon.Address{}, nil, 0, 0, fmt.Errorf("ParseTxData: failed to set transaction type: %w", err)
	}

	fromEthAddr := sender
	amount := eventData.Amount
	fromIdx = common.AccountIdx(uint32(eventData.From.Uint64()))

	if txType == common.TxTypeCreateAccountDeposit || txType == common.TxTypeDeposit || txType == common.TxTypeWithdraw {
		toEthAddr = ethCommon.Address{}
		toIdx = 0
	} else {
		toAccount, err := s.historydb.GetAccountByIdx(uint32(eventData.To.Uint64()))
		if err != nil {
			return "", ethCommon.Address{}, ethCommon.Address{}, nil, 0, 0, fmt.Errorf("ParseTxData: failed to get to account index for %s: %w", toEthAddr.Hex(), err)
		}
		toEthAddr = toAccount.EthAddr
		toIdx = toAccount.Idx
	}
	return txType, fromEthAddr, toEthAddr, amount, fromIdx, toIdx, nil
}

// ParseForgeBatchTxData parses the transaction data from ForgeBatchTx
// The data format is:
// - lastForgedBatch (4 byte)
// - lastForgedTxn (32 bytes)
// - batchSize (32 bytes)
// - txnData
func (s *Synchronizer) ParseForgeBatchTxData(eventData *bindings.SybilForgeBatch) (uint32, *big.Int, *big.Int, []common.Tx, error) {

	out, err := decodeTxnData(eventData.TxnData)
	if err != nil {
		return 0, nil, nil, nil, err
	}

	s.logger.Println("Parsed Transaction Data: ", out)

	txnData := make([]common.Tx, eventData.BatchSize.Uint64())

	for i, tx := range out {
		var txType string
		switch tx.Identifier {
		case 0:
			txType = common.TxTypeCreateAccountDeposit
		case 1:
			txType = common.TxTypeDeposit
		case 2:
			txType = common.TxTypeWithdraw
		case 3:
			txType = common.TxTypeVouch
		case 4:
			txType = common.TxTypeUnvouch
		case 5:
			txType = common.TxTypeExplode
		default:
			txType = common.TxTypeUnknown
		}
		txnData[i].Type = txType
		txnData[i].BatchNum = eventData.LastForgedBatch + 1
		txnData[i].FromIdx = common.AccountIdx(tx.From)
		fromAcc, err := s.historydb.GetAccountByIdx(uint32(tx.From))
		if err != nil {
			return 0, nil, nil, nil, err
		}
		txnData[i].FromEthAddr = fromAcc.EthAddr[:]

		if txType == common.TxTypeVouch ||
			txType == common.TxTypeUnvouch ||
			txType == common.TxTypeExplode {
			txnData[i].ToIdx = common.AccountIdx(tx.To)
			toAcc, err := s.historydb.GetAccountByIdx(uint32(tx.To))
			if err != nil {
				return 0, nil, nil, nil, err
			}
			txnData[i].ToEthAddr = toAcc.EthAddr[:]
		}
		txnData[i].Amount = tx.Amount
	}

	return eventData.LastForgedBatch, eventData.LastForgedTxn, eventData.BatchSize, txnData, nil
}

// SetType determines the transaction type based on the first byte
func SetType(firstByte byte) (string, error) {
	switch firstByte {
	case 0:
		return common.TxTypeCreateAccountDeposit, nil
	case 1:
		return common.TxTypeDeposit, nil
	case 2:
		return common.TxTypeWithdraw, nil
	case 3:
		return common.TxTypeVouch, nil
	case 4:
		return common.TxTypeUnvouch, nil
	case 5:
		return common.TxTypeExplode, nil
	default:
		return common.TxTypeUnknown, common.Wrap(fmt.Errorf("Unknown transaction type with first byte: %d (hex: 0x%02x)",
			firstByte, firstByte))
	}
}

type Transaction struct {
	Identifier uint8
	From       uint32
	To         uint32
	Amount     *big.Int
}

func decodeTxnData(txnData []byte) ([]Transaction, error) {
	const abiJSON = `[
		{
			"name": "dummy",
			"type": "function",
			"inputs": [],
			"outputs": [
				{
					"name": "transactions",
					"type": "tuple[]",
					"components": [
						{"name": "identifier", "type": "uint8"},
						{"name": "from", "type": "uint24"},
						{"name": "to", "type": "uint24"},
						{"name": "amount", "type": "uint128"}
					]
				}
			],
			"stateMutability": "view"
		}
	]`

	parsedABI, err := abi.JSON(strings.NewReader(abiJSON))
	if err != nil {
		return nil, err
	}

	var out []struct {
		Identifier uint8
		From       *big.Int
		To         *big.Int
		Amount     *big.Int
	}

	err = parsedABI.UnpackIntoInterface(&out, "dummy", txnData)
	if err != nil {
		return nil, err
	}

	txns := make([]Transaction, len(out))
	for i, d := range out {
		id := d.Identifier
		from := uint32(d.From.Uint64()) // fits uint24 < 2^24
		to := uint32(d.To.Uint64())
		amount := d.Amount

		txns[i] = Transaction{
			Identifier: id,
			From:       from,
			To:         to,
			Amount:     amount,
		}
	}
	return txns, nil
}
