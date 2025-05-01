package synchronizer

import (
	"fmt"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/tokamak-network/syb-sequencer/sequencer/abis/bindings"
	"github.com/tokamak-network/syb-sequencer/sequencer/common"
	"github.com/tokamak-network/syb-sequencer/sequencer/db/historydb"
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

const (
	TxTypeDeposit              string = "Deposit"
	TxTypeCreateAccountDeposit string = "CreateAccountDeposit"
	TxTypeForceExit            string = "ForceExit"
	TxTypeCreateVouch          string = "CreateVouch"
	TxTypeDeleteVouch          string = "DeleteVouch"
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
		fmt.Println(log.Topics[0])
		return l1UserTx, "L1UserTxEvent", nil
	}

	// If we get here, it's an unknown event
	return nil, fmt.Sprintf("Unknown(%s)", log.Topics[0].Hex()[:10]), nil
}

// SetType sets the type of the transaction
// TODO: Update this once all the events are in place
func SetType(tx *historydb.Tx) error {
	// Check if FromIdx is nil or 0
	if tx.FromIdx == nil || *tx.FromIdx == 0 {
		if tx.ToIdx == 0 { // Assuming ToIdx is not a pointer
			tx.Type = TxTypeCreateAccountDeposit
			return nil
		} else {
			return common.Wrap(fmt.Errorf(
				"cannot determine type of L1Tx, invalid ToIdx value: %d", tx.ToIdx))
		}
	} else if tx.FromIdx != nil && *tx.FromIdx > 0 {
		if tx.ToIdx == 0 {
			tx.Type = TxTypeDeposit
			return nil
		} else if tx.ToIdx == 1 {
			tx.Type = TxTypeForceExit
			return nil
		} else if tx.Type == TxTypeCreateVouch || tx.Type == TxTypeDeleteVouch {
			return nil
		} else {
			return common.Wrap(fmt.Errorf(
				"cannot determine type of L1Tx, invalid ToIdx value: %d", tx.ToIdx))
		}
	} else {
		return common.Wrap(fmt.Errorf(
			"cannot determine type of L1Tx, invalid FromIdx value: %d", *tx.FromIdx))
	}
}
