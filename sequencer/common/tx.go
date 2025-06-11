package common

import (
	"math/big"
)

// Transaction types
const (
	TxTypeCreateAccountDeposit string = "CreateAccountDeposit"
	TxTypeDeposit              string = "Deposit"
	TxTypeWithdraw             string = "Withdraw"
	TxTypeVouch                string = "Vouch"
	TxTypeUnvouch              string = "Unvouch"
	TxTypeExplode              string = "Explode"
	TxTypeUnknown              string = "Unknown"
)

// Tx represents a transaction
type Tx struct {
	ItemID      int64      `json:"item_id"`
	BatchNum    uint32     `json:"batch_num"`
	Position    *big.Int   `json:"position"`
	Type        string     `json:"type"`
	FromIdx     AccountIdx `json:"from_idx,omitempty"`
	FromEthAddr []byte     `json:"from_eth_addr,omitempty"`
	ToIdx       AccountIdx `json:"to_idx"`
	ToEthAddr   []byte     `json:"to_eth_addr,omitempty"`
	Amount      *big.Int   `json:"amount"`
	BlockNumber uint64     `json:"block_number"`
	Timestamp   uint64     `json:"timestamp"`
	GasFee      *big.Int   `json:"gas_fee,omitempty"`
	TxHash      []byte     `json:"tx_hash,omitempty"`
}
