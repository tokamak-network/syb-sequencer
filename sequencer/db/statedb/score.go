package statedb

import (
	"errors"
	"fmt"
	"log"
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
	PrefixKeyScoIdx = []byte("si:")
	// PrefixKeyScoHash is the key prefix for Score hash in the db
	PrefixKeyScoHash = []byte("sh:")
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

// GetScores return array of scores from Score Merkle Tree
func (s *StateDB) GetScores() ([]*big.Int, error) {
	num := 1 << (MaxNLevels - 1)

	scores := make([]*big.Int, num)

	for i := 0; i < num; i++ {
		idx := common.ScoreIdx(i)
		score, err := s.GetScore(idx)
		if err != nil {
			if errors.Is(err, ErrKeyNotFound) {
				scores[i] = big.NewInt(0)
				continue
			}
			return nil, fmt.Errorf("failed to get score %d: %w", i, err)
		}
		scores[i] = score.Score
	}

	return scores, nil
}

func (s *StateDB) CalculateScore() ([]*big.Int, error) {
	num := 1 << (MaxNLevels - 1)

	balances, err := s.GetBalances()
	if err != nil {
		return nil, err
	}
	log.Println(balances)

	vouches, err := s.GetVouches()
	if err != nil {
		return nil, err
	}
	log.Println(vouches)

	scores, err := s.GetScores()
	if err != nil {
		return nil, err
	}
	log.Println(scores)

	newScores := make([]*big.Int, num)

	for i := 0; i < num; i++ {
		// Mock Score Calculation
		newScores[i].Add(newScores[i], big.NewInt(1))
	}

	return newScores, nil
}
