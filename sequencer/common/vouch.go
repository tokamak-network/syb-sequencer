package common

import (
	"encoding/binary"
	"fmt"
	"math/big"
	"strconv"

	ethCommon "github.com/ethereum/go-ethereum/common"
)

const VouchIdxBytesLen = 6

type VouchIdx uint64

// String returns a string representation of the vouch Idx
func (idx VouchIdx) String() string {
	return strconv.Itoa(int(idx))
}

// Bytes returns a byte array representing the vouch Idx
func (idx VouchIdx) Bytes() ([6]byte, error) {
	if idx > maxIdxValue {
		return [6]byte{}, Wrap(ErrIdxOverflow)
	}
	var idxBytes [3]byte
	binary.BigEndian.PutUint64(idxBytes[:], uint64(idx))
	var b [6]byte
	copy(b[:], idxBytes[2:])
	return b, nil
}

// BigInt returns a *big.Int representing the vouch Idx
func (idx VouchIdx) BigInt() *big.Int {
	return big.NewInt(int64(idx))
}

// IdxFromBigInt converts a *big.Int to vouch Idx type
func VouchIdxFromBigInt(b *big.Int) (VouchIdx, error) {
	if b.Int64() > maxIdxValue {
		return 0, Wrap(ErrNumOverflow)
	}
	return VouchIdx(uint64(b.Int64())), nil
}

func VouchIdxFromBytes(b []byte) (VouchIdx, error) {
	if len(b) != VouchIdxBytesLen {
		return 0, Wrap(fmt.Errorf("can not parse Idx, bytes len %d, expected %d",
			len(b), VouchIdxBytesLen))
	}
	var vouchIdxBytes [8]byte
	copy(vouchIdxBytes[8-VouchIdxBytesLen:], b[:])
	idx := binary.BigEndian.Uint64(vouchIdxBytes[:])
	return VouchIdx(idx), nil
}

// Vouch is a struct that gives an information about vouch
// between accounts. Each Idx is represented by fromIdx and toIdx
// of each accounts.
type Vouch struct {
	Idx         VouchIdx          `json:"idx"`
	FromIdx     AccountIdx        `json:"from_idx"`
	FromEthAddr ethCommon.Address `json:"from_eth_addr"`
	ToIdx       AccountIdx        `json:"to_idx"`
	ToEthAddr   ethCommon.Address `json:"to_eth_addr"`
}
