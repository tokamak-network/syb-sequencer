package common

import (
	"math/big"

	ethCommon "github.com/ethereum/go-ethereum/common"
)

const (
	// TxTypeDeposit
	TxTypeDeposit string = "Deposit"
	// TxTypeCreateAccountDeposit represents creation of a new leaf in the state tree
	TxTypeCreateAccountDeposit string = "CreateAccountDeposit"
	// TxTypeForceExit TBD
	TxTypeForceExit string = "ForceExit"
	// TxTypeCreateVouch
	TxTypeCreateVouch string = "CreateVouch"
	// TxTypeDeleteVouch
	TxTypeUnVouch string = "UnVouch"
)

// Tx represents a transaction
type Tx struct {
	ItemID      int64    `json:"item_id"`
	BatchNum    int64    `json:"batch_num"`
	Position    int      `json:"position"`
	Type        string   `json:"type"`
	FromIdx     *int64   `json:"from_idx,omitempty"`
	FromEthAddr []byte   `json:"from_eth_addr,omitempty"`
	ToIdx       int64    `json:"to_idx"`
	ToEthAddr   []byte   `json:"to_eth_addr,omitempty"`
	Amount      *big.Int `json:"amount"`
	BlockNumber uint64   `json:"block_number"`
	Timestamp   uint64   `json:"timestamp"`
	GasFee      *big.Int `json:"gas_fee,omitempty"`
}

// EthAddrToBigInt returns a *big.Int from a given ethereum common.Address.
func EthAddrToBigInt(a ethCommon.Address) *big.Int {
	return new(big.Int).SetBytes(a.Bytes())
}
