package synchronizer

import (
	"fmt"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/tokamak-network/syb-sequencer/sequencer/abis/bindings"
)

// Event types
type L1UserTxEvent struct {
	QueueIndex uint32
	Position   uint32
	L1UserTx   []byte
}

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

	// Event structs
	var l1UserTx L1UserTxEvent

	// Check which event it is and unpack accordingly
	switch log.Topics[0] {
	case logSYBL1UserTxEvent:
		err = sybilABI.UnpackIntoInterface(&l1UserTx, "L1UserTxEvent", log.Data)
		if err != nil {
			return nil, "", fmt.Errorf("parseEvent: failed to unpack L1UserTxEvent: %w", err)
		}
		return l1UserTx, "L1UserTxEvent", nil
	}

	// If we get here, it's an unknown event
	return nil, fmt.Sprintf("Unknown(%s)", log.Topics[0].Hex()[:10]), nil
}
