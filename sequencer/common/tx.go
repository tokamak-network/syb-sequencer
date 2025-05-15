package common

import "math/big"

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
}
