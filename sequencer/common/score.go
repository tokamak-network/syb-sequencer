package common

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"math/big"
	"strconv"

	"github.com/iden3/go-iden3-crypto/poseidon"
	cryptoUtils "github.com/iden3/go-iden3-crypto/utils"
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
	var idxBytes [4]byte
	binary.BigEndian.PutUint32(idxBytes[:], uint32(idx))
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

// Score is a struct that gives information of the holdings of an address.
// Is the data structure that generates the Value stored in
// the leaf of the MerkleTree
type Score struct {
	Idx   ScoreIdx `json:"idx"`
	Score *big.Int `json:"score,bigint"`
}

func (s *Score) String() string {
	buf := bytes.NewBufferString("")
	fmt.Fprintf(buf, "Idx: %v, ", s.Idx)
	fmt.Fprintf(buf, "Score: %s, ", s.Score.String())
	return buf.String()
}

// Bytes returns the bytes representing the Score, in a way that each BigInt
// is represented by 32 bytes, in spite of the BigInt could be represented in
// less bytes (due a small big.Int), so in this way each BigInt is always 32
// bytes and can be automatically parsed from a byte array.
func (s *Score) Bytes() ([32]byte, error) {
	var b [32]byte

	scoreBytes := s.Score.Bytes()
	if len(scoreBytes) > 4 {
		return b, Wrap(fmt.Errorf("score overflow: needs <= 4 bytes"))
	}
	copy(b[32-len(scoreBytes):32], scoreBytes)

	return b, nil
}

// BigInts returns the [1]*big.Int, where each *big.Int is inside the Finite Field
func (s *Score) BigInts() ([1]*big.Int, error) {
	e := [1]*big.Int{}

	b, err := s.Bytes()
	if err != nil {
		return e, Wrap(err)
	}

	e[0] = new(big.Int).SetBytes(b[0:32])

	return e, nil
}

// HashValue returns the value of the Score, which is the Poseidon hash of its
// *big.Int representation
func (s *Score) HashValue() (*big.Int, error) {
	bi, err := s.BigInts()
	if err != nil {
		return nil, Wrap(err)
	}
	return poseidon.Hash(bi[:])
}

// ScoreFromBigInts returns a Account from a [1]*big.Int
func ScoreFromBigInts(e [1]*big.Int) (*Score, error) {
	if !cryptoUtils.CheckBigIntArrayInField(e[:]) {
		return nil, Wrap(ErrNotInFF)
	}
	e0B := e[0].Bytes()
	var b [32]byte
	copy(b[32-len(e0B):32], e0B)

	return ScoreFromBytes(b)
}

// ScoreFromBytes returns a Score from a byte array
func ScoreFromBytes(b [32]byte) (*Score, error) {
	score := new(big.Int).SetBytes(b[28:32])

	a := Score{
		Score: score,
	}
	return &a, nil
}
