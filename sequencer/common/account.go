package common

import (
	"encoding/binary"
	"fmt"
	"math/big"
	"strconv"

	ethCommon "github.com/ethereum/go-ethereum/common"
)

const (
	// RollupConstReservedIDx First 256 indexes reserved, first user index will be the 256
	RollupConstReservedIDx = 255

	// AccountIdxBytesLen idx bytes
	AccountIdxBytesLen = 3
	// maxIdxValue is the maximum value that Idx can have (24 bits:
	// maxIdxValue=2**24-1)
	maxIdxValue = 0xffffff
)

// AccountIdx represents the account Index in the MerkleTree
type AccountIdx uint32

// String returns a string representation of the account Idx
func (idx AccountIdx) String() string {
	return strconv.Itoa(int(idx))
}

// Bytes returns a byte array representing the account Idx
func (idx AccountIdx) Bytes() ([3]byte, error) {
	if idx > maxIdxValue {
		return [3]byte{}, Wrap(ErrIdxOverflow)
	}
	var idxBytes [3]byte
	binary.BigEndian.PutUint64(idxBytes[:], uint64(idx))
	var b [3]byte
	copy(b[:], idxBytes[1:])
	return b, nil
}

// BigInt returns a *big.Int representing the account Idx
func (idx AccountIdx) BigInt() *big.Int {
	return big.NewInt(int64(idx))
}

// AccountIdxFromBigInt converts a *big.Int to account Idx type
func AccountIdxFromBigInt(b *big.Int) (AccountIdx, error) {
	if b.Int64() > maxIdxValue {
		return 0, Wrap(ErrNumOverflow)
	}
	return AccountIdx(uint32(b.Int64())), nil
}

// AccountIdxFromBytes returns account Idx from a byte array
func AccountIdxFromBytes(b []byte) (AccountIdx, error) {
	if len(b) != AccountIdxBytesLen {
		return 0, Wrap(fmt.Errorf("can not parse Idx, bytes len %d, expected %d",
			len(b), AccountIdxBytesLen))
	}
	var accountIdxBytes [4]byte
	copy(accountIdxBytes[4-AccountIdxBytesLen:], b[:])
	idx := binary.BigEndian.Uint32(accountIdxBytes[:])
	return AccountIdx(idx), nil
}

// Account is a struct that gives information of the holdings of an address.
// Is the data structure that generates the Value stored in
// the leaf of the MerkleTree
type Account struct {
	Idx     AccountIdx        `json:"idx"`
	EthAddr ethCommon.Address `json:"eth_addr"`
	Balance *big.Int          `json:"balance,bigint"`
	Score   *big.Int          `json:"score,bigint"`
}
