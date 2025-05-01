package common

import (
	"math/big"

	ethCommon "github.com/ethereum/go-ethereum/common"
	"github.com/iden3/go-iden3-crypto/babyjub"
)

// Account is a struct that gives information of the holdings of an address.
// Is the data structure that generates the Value stored in
// the leaf of the MerkleTree
type Account struct {
	Idx      AccountIdx            `meddler:"idx"`
	BatchNum BatchNum              `meddler:"batch_num"`
	BJJ      babyjub.PublicKeyComp `meddler:"bjj"`
	EthAddr  ethCommon.Address     `meddler:"eth_addr"`
	Nonce    Nonce                 `meddler:"-"` // max of 40 bits used
	Balance  *big.Int              `meddler:"-"` // max of 192 bits used
}

// AccountIdx represents the account Index in the MerkleTree
type AccountIdx uint32

const (
	// NAccountLeafElems is the number of elements for a leaf in account tree
	NAccountLeafElems = 4

	// maxBalanceBytes is the maximum bytes that can use the
	// Account.Balance *big.Int
	maxBalanceBytes = 24

	// AccountIdxBytesLen idx bytes
	AccountIdxBytesLen = 6

	// maxAccountIdxValue is the maximum value that AccountIdx can have (24 bits:
	// maxAccountIdxValue=2**24-1)
	maxAccountIdxValue = 0xffffff

	// UserThreshold determines the threshold from the User Idxs can be
	UserThreshold = 256
	// IdxUserThreshold is a Idx type value that determines the threshold
	// from the User Idxs can be
	IdxUserThreshold = AccountIdx(UserThreshold)
)
