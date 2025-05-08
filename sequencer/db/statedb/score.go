package statedb

import (
	"errors"
	"math/big"

	"github.com/iden3/go-merkletree"
	"github.com/iden3/go-merkletree/db"
	"github.com/tokamak-network/syb-sequencer/sequencer/common"
)

var (
	// ErrScoreAlreadyExists is used when CreateScore is called and the
	// Score already exists
	ErrScoreAlreadyExists = errors.New("cannot CreateScore because Score already exists")
	// PrefixKeyScoIdx is the key prefix for ScoreIdx in the db
	PrefixKeyScoIdx = []byte("i:")
	// PrefixKeyScoHash is the key prefix for Score hash in the db
	PrefixKeyScoHash = []byte("h:")
)

// CreateScore creates a new Score in the StateDB for the given Idx.  If
// StateDB.ST==nil, MerkleTree is not affected, otherwise updates the
// MerkleTree, returning a CircomProcessorProof.
func (s *StateDB) CreateScore(idx common.ScoreIdx, score *common.Score) (
	*merkletree.CircomProcessorProof, error) {
	cpp, err := CreateScoreInTreeDB(s.db.DB(), s.ST, idx, score)
	if err != nil {
		return cpp, common.Wrap(err)
	}
	return cpp, nil
}

// CreateScoreInTreeDB is abstracted from StateDB to be used from StateDB.
// Creates a new Score in the StateDB for the given Idx.  If
// StateDB.ST==nil, MerkleTree is not affected, otherwise updates the
// MerkleTree, returning a CircomProcessorProof.
func CreateScoreInTreeDB(sto db.Storage, mt *merkletree.MerkleTree, idx common.ScoreIdx,
	score *common.Score) (*merkletree.CircomProcessorProof, error) {
	// store at the DB the key: v, and value: leaf.Bytes()
	v, err := score.HashValue()
	if err != nil {
		return nil, common.Wrap(err)
	}
	scoreBytes, err := score.Bytes()
	if err != nil {
		return nil, common.Wrap(err)
	}

	// store the Leaf value
	tx, err := sto.NewTx()
	if err != nil {
		return nil, common.Wrap(err)
	}

	idxBytes, err := idx.Bytes()
	if err != nil {
		return nil, common.Wrap(err)
	}
	_, err = tx.Get(append(PrefixKeyScoIdx, idxBytes[:]...))
	if common.Unwrap(err) != db.ErrNotFound {
		return nil, common.Wrap(ErrScoreAlreadyExists)
	}

	err = tx.Put(append(PrefixKeyScoHash, v.Bytes()...), scoreBytes[:])
	if err != nil {
		return nil, common.Wrap(err)
	}
	err = tx.Put(append(PrefixKeyScoIdx, idxBytes[:]...), v.Bytes())
	if err != nil {
		return nil, common.Wrap(err)
	}

	if err := tx.Commit(); err != nil {
		return nil, common.Wrap(err)
	}

	if mt != nil {
		return mt.AddAndGetCircomProof(idx.BigInt(), v)
	}

	return nil, nil
}

// UpdateScore updates the Score in the StateDB for the given Idx.  If
// StateDB.mt==nil, MerkleTree is not affected, otherwise updates the
// MerkleTree, returning a CircomProcessorProof.
func (s *StateDB) UpdateScore(idx common.ScoreIdx, score *common.Score) (
	*merkletree.CircomProcessorProof, error) {
	return UpdateScoreInTreeDB(s.db.DB(), s.ST, idx, score)
}

// UpdateScoreInTreeDB is abstracted from StateDB to be used from StateDB.
// Updates the Score in the StateDB for the given Idx.  If
// StateDB.mt==nil, MerkleTree is not affected, otherwise updates the
// MerkleTree, returning a CircomProcessorProof.
func UpdateScoreInTreeDB(sto db.Storage, mt *merkletree.MerkleTree, idx common.ScoreIdx,
	score *common.Score) (*merkletree.CircomProcessorProof, error) {
	// store at the DB the key: v, and value: Score.Bytes()
	v, err := score.HashValue()
	if err != nil {
		return nil, common.Wrap(err)
	}
	scoreBytes, err := score.Bytes()
	if err != nil {
		return nil, common.Wrap(err)
	}

	tx, err := sto.NewTx()
	if err != nil {
		return nil, common.Wrap(err)
	}
	err = tx.Put(append(PrefixKeyScoHash, v.Bytes()...), scoreBytes[:])
	if err != nil {
		return nil, common.Wrap(err)
	}
	idxBytes, err := idx.Bytes()
	if err != nil {
		return nil, common.Wrap(err)
	}
	err = tx.Put(append(PrefixKeyScoIdx, idxBytes[:]...), v.Bytes())
	if err != nil {
		return nil, common.Wrap(err)
	}

	if err := tx.Commit(); err != nil {
		return nil, common.Wrap(err)
	}

	if mt != nil {
		proof, err := mt.Update(idx.BigInt(), v)
		return proof, common.Wrap(err)
	}
	return nil, nil
}

// STGetProof returns the CircomVerifierProof for a given scoreIdx
func (s *StateDB) STGetProof(idx common.ScoreIdx) (*merkletree.CircomVerifierProof, error) {
	if s.ST == nil {
		return nil, common.Wrap(ErrStateDBWithoutMT)
	}
	p, err := s.ST.GenerateSCVerifierProof(idx.BigInt(), s.ST.Root())
	if err != nil {
		return nil, common.Wrap(err)
	}
	return p, nil
}

// GetScore returns the score for the given Idx
func (s *StateDB) GetScore(idx common.ScoreIdx) (*common.Score, error) {
	return GetScoreInTreeDB(s.db.DB(), idx)
}

func ScoresIter(db db.Storage, fn func(a *common.Score) (bool, error)) error {
	idxDB := db.WithPrefix(PrefixKeyScoIdx)
	if err := idxDB.Iterate(func(k []byte, v []byte) (bool, error) {
		idx, err := common.ScoreIdxFromBytes(k)
		if err != nil {
			return false, common.Wrap(err)
		}
		acc, err := GetScoreInTreeDB(db, idx)
		if err != nil {
			return false, common.Wrap(err)
		}
		ok, err := fn(acc)
		if err != nil {
			return false, common.Wrap(err)
		}
		return ok, nil
	}); err != nil {
		return common.Wrap(err)
	}
	return nil
}

// GetScoreInTreeDB is abstracted from StateDB to be used from StateDB.
// GetScore returns the Score for the given Idx
func GetScoreInTreeDB(sto db.Storage, idx common.ScoreIdx) (*common.Score, error) {
	idxBytes, err := idx.Bytes()
	if err != nil {
		return nil, common.Wrap(err)
	}
	vBytes, err := sto.Get(append(PrefixKeyScoIdx, idxBytes[:]...))
	if err != nil {
		return nil, common.Wrap(err)
	}
	accBytes, err := sto.Get(append(PrefixKeyScoHash, vBytes...))
	if err != nil {
		return nil, common.Wrap(err)
	}
	var b [32]byte
	copy(b[:], accBytes)
	score, err := common.ScoreFromBytes(b)
	if err != nil {
		return nil, common.Wrap(err)
	}
	score.Idx = idx
	return score, nil
}

// GetSTRoot returns the root of the Score Merkle Tree
func (s *StateDB) GetSTRoot() *big.Int {
	return s.ST.Root().BigInt()
}
