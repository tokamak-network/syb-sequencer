package common

import (
	"encoding/binary"
	"fmt"
	"math/big"
	"strconv"
)

const (
	// ScoreIdxBytesLen idx bytes
	ScoreIdxBytesLen = 3
	// maxIdxValue is the maximum value that Idx can have (24 bits:
	// maxIdxValue=2**24-1)
	maxScoreIdxValue = 0xffffff
)

// ScoreIdx represents the score Index in the MerkleTree
type ScoreIdx uint32

// String returns a string representation of the score Idx
func (idx ScoreIdx) String() string {
	return strconv.Itoa(int(idx))
}

// Bytes returns a byte array representing the score Idx
func (idx ScoreIdx) Bytes() ([3]byte, error) {
	if idx > maxIdxValue {
		return [3]byte{}, Wrap(ErrIdxOverflow)
	}
	var idxBytes [3]byte
	binary.BigEndian.PutUint64(idxBytes[:], uint64(idx))
	var b [3]byte
	copy(b[:], idxBytes[1:])
	return b, nil
}

// BigInt returns a *big.Int representing the score Idx
func (idx ScoreIdx) BigInt() *big.Int {
	return big.NewInt(int64(idx))
}

// ScoreIdxFromBigInt converts a *big.Int to score Idx type
func ScoreIdxFromBigInt(b *big.Int) (ScoreIdx, error) {
	if b.Int64() > maxIdxValue {
		return 0, Wrap(ErrNumOverflow)
	}
	return ScoreIdx(uint32(b.Int64())), nil
}

// ScoreIdxFromBytes returns Idx from a byte array
func ScoreIdxFromBytes(b []byte) (ScoreIdx, error) {
	if len(b) != ScoreIdxBytesLen {
		return 0, Wrap(fmt.Errorf("can not parse Idx, bytes len %d, expected %d",
			len(b), ScoreIdxBytesLen))
	}
	var ScoreIdxBytes [4]byte
	copy(ScoreIdxBytes[4-ScoreIdxBytesLen:], b[:])
	idx := binary.BigEndian.Uint32(ScoreIdxBytes[:])
	return ScoreIdx(idx), nil
}
