package common

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"math/big"
	"strconv"

	ethCommon "github.com/ethereum/go-ethereum/common"
	"github.com/iden3/go-iden3-crypto/poseidon"
	cryptoUtils "github.com/iden3/go-iden3-crypto/utils"
)

const (
	// NLeafElems is the number of elements for a leaf
	NLeafElems = 2

	// maxBalanceBytes is the maximum bytes that can use the
	// Account.Balance *big.Int
	maxBalanceBytes = 24

	// RollupConstReservedIDx First 256 indexes reserved, first user index will be the 256
	RollupConstReservedIDx = 255

	// AccountIdxBytesLen idx bytes
	AccountIdxBytesLen = 3
	// maxIdxValue is the maximum value that Idx can have (24 bits:
	// maxIdxValue=2**24-1)
	maxIdxValue = 0xffffff
)

// AccountIdx represents the account Index in the MerkleTree
type AccountIdx int64

// String returns a string representation of the account Idx
func (idx AccountIdx) String() string {
	return strconv.Itoa(int(idx))
}

// Bytes returns a byte array representing the account Idx
func (idx AccountIdx) Bytes() ([3]byte, error) {
	if idx > maxIdxValue {
		return [3]byte{}, Wrap(ErrIdxOverflow)
	}
	var idxBytes [4]byte
	binary.BigEndian.PutUint32(idxBytes[:], uint32(idx))
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
	Idx           AccountIdx        `json:"idx"`
	EthAddr       ethCommon.Address `json:"ethAddr"`
	Balance       *big.Int          `json:"balance"`
	Score         *big.Int          `json:"score"`
	ScoreSiblings []*big.Int        `json:"scoreSiblings"`
}

func (a *Account) String() string {
	buf := bytes.NewBufferString("")
	fmt.Fprintf(buf, "Idx: %v, ", a.Idx)
	fmt.Fprintf(buf, "EthAddr: %s..., ", a.EthAddr.String()[:10])
	fmt.Fprintf(buf, "Balance: %s, ", a.Balance.String())
	return buf.String()
}

// Bytes returns the bytes representing the Account, in a way that each BigInt
// is represented by 32 bytes, in spite of the BigInt could be represented in
// less bytes (due a small big.Int), so in this way each BigInt is always 32
// bytes and can be automatically parsed from a byte array.
func (a *Account) Bytes() ([32 * NLeafElems]byte, error) {
	var b [32 * NLeafElems]byte

	if len(a.Balance.Bytes()) > maxBalanceBytes {
		return b, Wrap(fmt.Errorf("%s Balance", ErrNumOverflow))
	}

	copy(b[12:32], a.EthAddr.Bytes())
	balanceBytes := a.Balance.Bytes()
	copy(b[64-len(balanceBytes):64], balanceBytes)

	return b, nil
}

// BigInts returns the [2]*big.Int, where each *big.Int is inside the Finite Field
func (a *Account) BigInts() ([NLeafElems]*big.Int, error) {
	e := [NLeafElems]*big.Int{}

	b, err := a.Bytes()
	if err != nil {
		return e, Wrap(err)
	}

	e[0] = new(big.Int).SetBytes(b[0:32])
	e[1] = new(big.Int).SetBytes(b[32:64])

	return e, nil
}

// HashValue returns the value of the Account, which is the Poseidon hash of its
// *big.Int representation
func (a *Account) HashValue() (*big.Int, error) {
	bi, err := a.BigInts()
	if err != nil {
		return nil, Wrap(err)
	}
	return poseidon.Hash(bi[:])
}

// AccountFromBigInts returns a Account from a [2]*big.Int
func AccountFromBigInts(e [NLeafElems]*big.Int) (*Account, error) {
	if !cryptoUtils.CheckBigIntArrayInField(e[:]) {
		return nil, Wrap(ErrNotInFF)
	}
	e0B := e[0].Bytes()
	e1B := e[1].Bytes()
	var b [32 * NLeafElems]byte
	copy(b[32-len(e0B):32], e0B)
	copy(b[64-len(e1B):64], e1B)

	return AccountFromBytes(b)
}

// AccountFromBytes returns a Account from a byte array
func AccountFromBytes(b [32 * NLeafElems]byte) (*Account, error) {
	ethAddr := ethCommon.BytesToAddress(b[12:32])

	balance := new(big.Int).SetBytes(b[40:64])
	// Balance is max of 192 bits (24 bytes)
	if !bytes.Equal(b[32:40], []byte{0, 0, 0, 0, 0, 0, 0, 0}) {
		return nil, Wrap(fmt.Errorf("%s Balance", ErrNumOverflow))
	}

	if !cryptoUtils.CheckBigIntInField(balance) {
		return nil, Wrap(ErrNotInFF)
	}

	a := Account{
		Balance: balance,
		EthAddr: ethAddr,
	}
	return &a, nil
}

func EthAddrToBigInt(a ethCommon.Address) *big.Int {
	return new(big.Int).SetBytes(a.Bytes())
}
