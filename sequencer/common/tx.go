package common

import (
	"math/big"

	ethCommon "github.com/ethereum/go-ethereum/common"
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
	ItemID      int64             `json:"item_id"`
	BatchNum    int64             `json:"batch_num"`
	Position    int               `json:"position"`
	Type        string            `json:"type"`
	FromIdx     AccountIdx        `json:"from_idx,omitempty"`
	FromEthAddr ethCommon.Address `json:"from_eth_addr,omitempty"`
	ToIdx       AccountIdx        `json:"to_idx"`
	ToEthAddr   ethCommon.Address `json:"to_eth_addr,omitempty"`
	Amount      *big.Int          `json:"amount"`
	BlockNumber uint64            `json:"block_number"`
	Timestamp   uint64            `json:"timestamp"`
	GasFee      *big.Int          `json:"gas_fee,omitempty"`
}
