package models

import "time"

// Transaction represents a blockchain transaction
type Transaction struct {
	ID          int64
	TxHash      string
	TxName      string
	BlockNumber int64
	CreatedAt   time.Time
}
