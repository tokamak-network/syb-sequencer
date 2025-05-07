package statedb

import (
	"errors"

	"github.com/iden3/go-merkletree"
	"github.com/tokamak-network/syb-sequencer/sequencer/common"
	"github.com/tokamak-network/syb-sequencer/sequencer/db/kvdb"
)

var (
	// ErrStateDBWithoutMT is used when a method that requires a MerkleTree
	// is called in a StateDB that does not have a MerkleTree defined
	ErrStateDBWithoutMT = errors.New(
		"cannot call method to use MerkleTree in a StateDB without MerkleTree")
	// ErrIdxNotFound is used when trying to get the Idx from EthAddr or
	// EthAddr&ToBJJ
	ErrIdxNotFound = errors.New("idx can not be found")
	// ErrGetIdxNoCase is used when trying to get the Idx from EthAddr &
	// BJJ with not compatible combination
	ErrGetIdxNoCase = errors.New(
		"cannot get Idx due unexpected combination of ethereum Address & BabyJubJub PublicKey")

	// PrefixKeyMTAcc is the key prefix for account merkle tree in the db
	PrefixKeyMTAcc = []byte("ma:")
	// PrefixKeyMTVoc is the key prefix for vouch merkle tree in the db
	PrefixKeyMTVoc = []byte("mv:")
	// PrefixKeyMTSco is the key prefix for score merkle tree in the db
	PrefixKeyMTSco = []byte("ms:")
)

const (
	// TypeSynchronizer defines a StateDB used by the Synchronizer, that
	// generates the ExitTree when processing the txs
	TypeSynchronizer = "synchronizer"
	// TypeBatchBuilder defines a StateDB used by the BatchBuilder, that
	// generates the ExitTree and the ZKInput when processing the txs
	TypeBatchBuilder = "batchbuilder"
	// MaxNLevels is the maximum value of NLevels for the merkle tree,
	// which comes from the fact that AccountIdx has 24 bits.
	MaxNLevels = 24
)

// TypeStateDB determines the type of StateDB
type TypeStateDB string

// Config of the StateDB
type Config struct {
	// Path where the checkpoints will be stored
	Path string
	// Keep is the number of old checkpoints to keep.  If 0, all
	// checkpoints are kept.
	Keep int
	// NoLast skips having an opened DB with a checkpoint to the last
	// batchNum for thread-safe reads.
	NoLast bool
	// Type of StateDB (
	Type TypeStateDB
	// NLevels is the number of merkle tree levels in case the Type uses a
	// merkle tree.  If the Type doesn't use a merkle tree, NLevels should
	// be 0.
	NLevels int
	// At every checkpoint, check that there are no gaps between the
	// checkpoints
	noGapsCheck bool
}

// StateDB represents the state database with an integrated Merkle tree.
type StateDB struct {
	cfg         Config
	db          *kvdb.KVDB
	AccountTree *merkletree.MerkleTree
	VouchTree   *merkletree.MerkleTree
	ScoreTree   *merkletree.MerkleTree
}

// NewStateDB initializes a new StateDB.
func NewStateDB(cfg Config) (*StateDB, error) {
	var kv *kvdb.KVDB
	var err error

	kv, err = kvdb.NewKVDB(kvdb.Config{Path: cfg.Path, Keep: cfg.Keep,
		NoGapsCheck: cfg.noGapsCheck, NoLast: cfg.NoLast})
	if err != nil {
		return nil, common.Wrap(err)
	}

	mtAccount, _ := merkletree.NewMerkleTree(kv.StorageWithPrefix(PrefixKeyMTAcc), 24)
	mtVouch, _ := merkletree.NewMerkleTree(kv.StorageWithPrefix(PrefixKeyMTVoc), 24)
	mtScore, _ := merkletree.NewMerkleTree(kv.StorageWithPrefix(PrefixKeyMTSco), 24)
	return &StateDB{
		cfg:         cfg,
		db:          kv,
		AccountTree: mtAccount,
		VouchTree:   mtVouch,
		ScoreTree:   mtScore,
	}, nil
}
