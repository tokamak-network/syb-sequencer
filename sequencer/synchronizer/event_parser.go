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

// Event types
type L1UserTxEvent struct {
	QueueIndex uint32
	Position   uint32
	L1UserTx   []byte
}

// Transaction types
const (
	TxTypeCreateAccountDeposit string = "CreateAccountDeposit"
	TxTypeDeposit              string = "Deposit"
	TxTypeWithdraw             string = "Withdraw"
	TxTypeCreateVouch          string = "CreateVouch"
	TxTypeUnvouch              string = "Unvouch"
	TxTypeExplode              string = "Explode"
	TxTypeUnknown              string = "Unknown"
)

var (
	logSYBL1UserTxEvent = crypto.Keccak256Hash([]byte(
		"L1UserTxEvent(uint32,uint8,bytes)"))
)

// ParseEvent parses data of events found on blockchain in required structure
func ParseEvent(log *types.Log) (interface{}, string, error) {
	// Parse ABI
	sybilABI, err := abi.JSON(strings.NewReader(bindings.BindingsMetaData.ABI))
	if err != nil {
		return nil, "", fmt.Errorf("parseEvent: failed to parse ABI: %w", err)
	}

	var l1UserTx L1UserTxEvent

	// Check which event it is and unpack accordingly
	switch log.Topics[0] {
	case logSYBL1UserTxEvent:
		err = sybilABI.UnpackIntoInterface(&l1UserTx, "L1UserTxEvent", log.Data)
		if err != nil {
			return nil, "", fmt.Errorf("parseEvent: failed to unpack L1UserTxEvent: %w", err)
		}
		fmt.Println(log.Topics[0])
		return l1UserTx, "L1UserTxEvent", nil
	}

	// If we get here, it's an unknown event
	return nil, fmt.Sprintf("Unknown(%s)", log.Topics[0].Hex()[:10]), nil
}

// ParseTxData parses the transaction data from L1UserTx
// The data format is:
// - txnType (1 byte)
// - fromEthAddress (20 bytes)
// - toEthAddress (20 bytes)
// - amount (32 bytes)
func ParseTxData(txData []byte) (string, ethCommon.Address, ethCommon.Address, *big.Int, error) {
	// Check if we have enough data
	if len(txData) < 73 {
		return "", ethCommon.Address{}, ethCommon.Address{}, nil, fmt.Errorf("transaction data too short: %d bytes, expected 73", len(txData))
	}

	// Extract transaction type (first byte)
	txType, _ := SetType(txData[0])

	// Extract from Ethereum address (next 20 bytes)
	fromEthAddr := ethCommon.BytesToAddress(txData[1:21])

	// Extract to Ethereum address (next 20 bytes)
	toEthAddr := ethCommon.BytesToAddress(txData[21:41])

	// Extract amount (next 32 bytes)
	amount := new(big.Int).SetBytes(txData[41:73])

	return txType, fromEthAddr, toEthAddr, amount, nil
}

// SetType determines the transaction type based on the first byte
func SetType(firstByte byte) (string, error) {
	switch firstByte {
	case 0:
		return TxTypeCreateAccountDeposit, nil
	case 1:
		return TxTypeDeposit, nil
	case 2:
		return TxTypeWithdraw, nil
	case 3:
		return TxTypeCreateVouch, nil
	case 4:
		return TxTypeUnvouch, nil
	case 5:
		return TxTypeExplode, nil
	default:
		return TxTypeUnknown, common.Wrap(fmt.Errorf("Unknown transaction type with first byte: %d (hex: 0x%02x)",
			firstByte, firstByte))
	}
}
