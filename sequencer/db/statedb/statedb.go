package statedb

import (
	"errors"

	"github.com/iden3/go-merkletree"
	"github.com/iden3/go-merkletree/db"
	"github.com/iden3/go-merkletree/db/pebble"
	"github.com/tokamak-network/syb-sequencer/sequencer/common"
	"github.com/tokamak-network/syb-sequencer/sequencer/db/kvdb"
	"github.com/tokamak-network/syb-sequencer/sequencer/log"
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
		"cannot get Idx due unexpected combination of ethereum Address")

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

// Last offers a subset of view methods of the StateDB that can be
// called via the LastRead method of StateDB in a thread-safe manner to obtain
// a consistent view to the last batch of the StateDB.
type Last struct {
	db db.Storage
}

// GetAccount returns the account for the given Idx
func (s *Last) GetAccount(idx common.AccountIdx) (*common.Account, error) {
	return GetAccountInTreeDB(s.db, idx)
}

// GetCurrentBatch returns the current BatchNum stored in Last.db
func (s *Last) GetCurrentBatch() (common.BatchNum, error) {
	cbBytes, err := s.db.Get(kvdb.KeyCurrentBatch)
	if common.Unwrap(err) == db.ErrNotFound {
		return 0, nil
	} else if err != nil {
		return 0, common.Wrap(err)
	}
	return common.BatchNumFromBytes(cbBytes)
}

// DB returns the underlying storage of Last
func (s *Last) DB() db.Storage {
	return s.db
}

// GetAccounts returns all the accounts in the db.  Use for debugging pruposes
// only.
func (s *Last) GetAccounts() ([]common.Account, error) {
	return getAccounts(s.db)
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

// LastRead is a thread-safe method to query the last checkpoint of the StateDB
// via the Last type methods
func (s *StateDB) LastRead(fn func(sdbLast *Last) error) error {
	return s.db.LastRead(
		func(db *pebble.Storage) error {
			return fn(&Last{
				db: db,
			})
		},
	)
}

// MakeCheckpoint does a checkpoint at the given batchNum in the defined path.
// Internally this advances & stores the current BatchNum, and then stores a
// Checkpoint of the current state of the StateDB.
func (s *StateDB) MakeCheckpoint() error {
	log.Debugw("Making StateDB checkpoint", "batch", s.CurrentBatch()+1, "type", s.cfg.Type)
	return s.db.MakeCheckpoint()
}

// DeleteOldCheckpoints deletes old checkpoints when there are more than
// `cfg.keep` checkpoints
func (s *StateDB) DeleteOldCheckpoints() error {
	return s.db.DeleteOldCheckpoints()
}

// CurrentBatch returns the current in-memory CurrentBatch of the StateDB.db
func (s *StateDB) CurrentBatch() common.BatchNum {
	return s.db.CurrentBatch
}

// getCurrentBatch returns the current BatchNum stored in the StateDB.db
func (s *StateDB) getCurrentBatch() (common.BatchNum, error) {
	return s.db.GetCurrentBatch()
}

// Reset resets the StateDB to the checkpoint at the given batchNum. Reset
// does not delete the checkpoints between old current and the new current,
// those checkpoints will remain in the storage, and eventually will be
// deleted when MakeCheckpoint overwrites them.
func (s *StateDB) Reset(batchNum common.BatchNum) error {
	log.Debugw("Making StateDB Reset", "batch", batchNum, "type", s.cfg.Type)
	if err := s.db.Reset(batchNum); err != nil {
		return common.Wrap(err)
	}
	if s.AccountTree != nil {
		// open the AccountTree for the current s.db
		mt, err := merkletree.NewMerkleTree(s.db.StorageWithPrefix(PrefixKeyMTAcc), s.AccountTree.MaxLevels())
		if err != nil {
			return common.Wrap(err)
		}
		s.AccountTree = mt
	}
	if s.VouchTree != nil {
		// open the VouchTree for the current s.db
		mt, err := merkletree.NewMerkleTree(s.db.StorageWithPrefix(PrefixKeyMTVoc), s.VouchTree.MaxLevels())
		if err != nil {
			return common.Wrap(err)
		}
		s.VouchTree = mt
	}
	if s.ScoreTree != nil {
		// open the ScoreTree for the current s.db
		mt, err := merkletree.NewMerkleTree(s.db.StorageWithPrefix(PrefixKeyMTSco), s.ScoreTree.MaxLevels())
		if err != nil {
			return common.Wrap(err)
		}
		s.ScoreTree = mt
	}
	return nil
}

// LocalStateDB represents the local StateDB which allows to make copies from
// the synchronizer StateDB, and is used by the tx-selector and the
// batch-builder. LocalStateDB is an in-memory storage.
type LocalStateDB struct {
	*StateDB
	synchronizerStateDB *StateDB
}

// NewLocalStateDB returns a new LocalStateDB connected to the given
// synchronizerDB.  Checkpoints older than the value defined by `keep` will be
// deleted.
func NewLocalStateDB(cfg Config, synchronizerDB *StateDB) (*LocalStateDB, error) {
	cfg.noGapsCheck = true
	cfg.NoLast = true
	s, err := NewStateDB(cfg)
	if err != nil {
		return nil, common.Wrap(err)
	}
	return &LocalStateDB{
		s,
		synchronizerDB,
	}, nil
}

// CheckpointExists returns true if the checkpoint exists
func (l *LocalStateDB) CheckpointExists(batchNum common.BatchNum) (bool, error) {
	return l.db.CheckpointExists(batchNum)
}

// Reset performs a reset in the LocalStateDB. If fromSynchronizer is true, it
// gets the state from LocalStateDB.synchronizerStateDB for the given batchNum.
// If fromSynchronizer is false, get the state from LocalStateDB checkpoints.
func (l *LocalStateDB) Reset(batchNum common.BatchNum, fromSynchronizer bool) error {
	if fromSynchronizer {
		log.Debugw("Making StateDB ResetFromSynchronizer", "batch", batchNum, "type", l.cfg.Type)
		if err := l.db.ResetFromSynchronizer(batchNum, l.synchronizerStateDB.db); err != nil {
			return common.Wrap(err)
		}
		// open the AccountTree for the current s.db
		if l.AccountTree != nil {
			mt, err := merkletree.NewMerkleTree(l.db.StorageWithPrefix(PrefixKeyMTAcc),
				l.AccountTree.MaxLevels())
			if err != nil {
				return common.Wrap(err)
			}
			l.AccountTree = mt
		}
		// open the MT for the current s.db
		if l.VouchTree != nil {
			mt, err := merkletree.NewMerkleTree(l.db.StorageWithPrefix(PrefixKeyMTVoc),
				l.VouchTree.MaxLevels())
			if err != nil {
				return common.Wrap(err)
			}
			l.VouchTree = mt
		}
		// open the MT for the current s.db
		if l.ScoreTree != nil {
			mt, err := merkletree.NewMerkleTree(l.db.StorageWithPrefix(PrefixKeyMTSco),
				l.ScoreTree.MaxLevels())
			if err != nil {
				return common.Wrap(err)
			}
			l.ScoreTree = mt
		}
		return nil
	}
	// use checkpoint from LocalStateDB
	return l.StateDB.Reset(batchNum)
}
