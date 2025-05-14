package common

import (
	"encoding/binary"
	"fmt"
	"math/big"
)

const batchNumBytesLen = 8

// Batch is a struct that represents SYB sequencer batch
type Batch struct {
	BatchNum    BatchNum `json:"item_id"`
	AccountRoot *big.Int `json:"account_root,bigint"`
	VouchRoot   *big.Int `json:"vouch_root,bigint"`
	ScoreRoot   *big.Int `json:"score_root,bigint"`
}

func NewEmptyBatch() *Batch {
	return &Batch{
		BatchNum:    0,
		AccountRoot: big.NewInt(0),
		VouchRoot:   big.NewInt(0),
		ScoreRoot:   big.NewInt(0),
	}
}

type BatchNum int64

// Bytes returns a byte array of length 8 representing the BatchNum
func (bn BatchNum) Bytes() []byte {
	var batchNumBytes [batchNumBytesLen]byte
	binary.BigEndian.PutUint64(batchNumBytes[:], uint64(bn))
	return batchNumBytes[:]
}

// BigInt returns a *big.Int representing the BatchNum
func (bn BatchNum) BigInt() *big.Int {
	return big.NewInt(int64(bn))
}

// BatchNumFromBytes returns BatchNum from a []byte
func BatchNumFromBytes(b []byte) (BatchNum, error) {
	if len(b) != batchNumBytesLen {
		return 0,
			Wrap(fmt.Errorf("can not parse BatchNumFromBytes, bytes len %d, expected %d",
				len(b), batchNumBytesLen))
	}
	batchNum := binary.BigEndian.Uint32(b[:batchNumBytesLen])
	return BatchNum(batchNum), nil
}
