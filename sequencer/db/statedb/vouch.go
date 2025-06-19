package statedb

import (
	"errors"
	"fmt"
	"math/big"

	"github.com/iden3/go-merkletree"
	"github.com/iden3/go-merkletree/db"
	"github.com/tokamak-network/syb-sequencer/sequencer/common"
)

var (
	ErrAlreadyVouched = errors.New("can not Vouch because already vouched")
	// PrefixKeyVocIdx is the key prefix for vouchIdx in the db
	PrefixKeyVocIdx = []byte("v:")
)

// Vouch creates a new Vouch in the StateDB for the given Idx. If
// StateDB.VT==nil, MerkleTree is not affected, otherwise updates the
// MerkleTree, returning a CircomProcessorProof.
func (s *StateDB) Vouch(idx common.VouchIdx, vouch *common.Vouch) (
	*merkletree.CircomProcessorProof, error) {
	cpp, err := CreateVouchInTreeDB(s.db.DB(), s.VT, idx)
	if err != nil {
		return cpp, common.Wrap(err)
	}
	return cpp, nil
}

// CreateVouchInTreeDB is abstracted from StateDB to be used from StateDB.
// Creates a new Vouch in the StateDB for the given Idx. If
// StateDB.VT==nil, MerkleTree is no affected, otherwise updates the
// MerkleTree, returning a CircomProcessorProof
func CreateVouchInTreeDB(sto db.Storage, mt *merkletree.MerkleTree, idx common.VouchIdx) (*merkletree.CircomProcessorProof, error) {
	// store the Leaf value
	tx, err := sto.NewTx()
	if err != nil {
		return nil, common.Wrap(err)
	}

	idxBytes, err := idx.Bytes()
	if err != nil {
		return nil, common.Wrap(err)
	}
	err = tx.Put(append(PrefixKeyVocIdx, idxBytes[:]...), []byte{1})
	if err != nil {
		return nil, common.Wrap(err)
	}
	if err := tx.Commit(); err != nil {
		return nil, common.Wrap(err)
	}

	if mt != nil {
		return mt.AddAndGetCircomProof(idx.BigInt(), big.NewInt(1))
	}

	return nil, nil
}

// UnVouch updates the Vouch in the StateDB for the given Idx.  If
// StateDB.VT==nil, MerkleTree is not affected, otherwise updates the
// MerkleTree, returning a CircomProcessorProof.
func (s *StateDB) UnVouch(idx common.VouchIdx) (
	*merkletree.CircomProcessorProof, error) {
	return DeleteVouchInTreeDB(s.db.DB(), s.VT, idx)
}

// DeleteVouchInTreeDB is abstracted from StateDB to be used from StateDB.
// Deletes the Vouch in the StateDB for the given Idx.  If
// StateDB.VT==nil, MerkleTree is not affected, otherwise updates the
// MerkleTree, returning a CircomProcessorProof.
func DeleteVouchInTreeDB(sto db.Storage, mt *merkletree.MerkleTree, idx common.VouchIdx) (*merkletree.CircomProcessorProof, error) {
	// store the Leaf value
	tx, err := sto.NewTx()
	if err != nil {
		return nil, common.Wrap(err)
	}

	idxBytes, err := idx.Bytes()
	if err != nil {
		return nil, common.Wrap(err)
	}
	err = tx.Put(append(PrefixKeyVocIdx, idxBytes[:]...), []byte{0})
	if err != nil {
		return nil, common.Wrap(err)
	}
	if err := tx.Commit(); err != nil {
		return nil, common.Wrap(err)
	}

	if mt != nil {
		return mt.AddAndGetCircomProof(idx.BigInt(), big.NewInt(0))
	}

	return nil, nil
}

// VTGetProof returns the CircomVerifierProof for a given vouchIdx
func (s *StateDB) VTGetProof(idx common.VouchIdx) (*merkletree.CircomVerifierProof, error) {
	if s.VT == nil {
		return nil, common.Wrap(ErrStateDBWithoutMT)
	}
	p, err := s.VT.GenerateSCVerifierProof(idx.BigInt(), s.VT.Root())
	if err != nil {
		return nil, common.Wrap(err)
	}
	return p, nil
}

// GetVouch returns the vouch for the given Idx
func (s *StateDB) GetVouch(idx common.VouchIdx) (*common.Vouch, error) {
	return GetVouchInTreeDB(s.db.DB(), idx)
}

// GetVouchInTreeDB is abstracted from StateDB to be used from StateDB.
// GetVouch returns the vouch for the given Idx
func GetVouchInTreeDB(sto db.Storage, idx common.VouchIdx) (*common.Vouch, error) {
	idxBytes, err := idx.Bytes()
	if err != nil {
		return nil, common.Wrap(err)
	}
	vBytes, err := sto.Get(append(PrefixKeyAccIdx, idxBytes[:]...))
	if err != nil {
		return nil, common.Wrap(err)
	}

	var b [1]byte
	copy(b[:], vBytes)
	if b[0] == 0 {
		return nil, nil
	}

	var vouch *common.Vouch

	// Split the byte array for FromIdx and ToIdx
	halfLen := common.VouchIdxBytesLen / 2
	fromIdx, err := common.AccountIdxFromBytes(b[:halfLen])
	if err != nil {
		return nil, common.Wrap(err)
	}
	vouch.FromIdx = fromIdx

	// Retrieve from account details
	fromAccount, err := GetAccountInTreeDB(sto, fromIdx)
	if err != nil {
		return nil, common.Wrap(err)
	}
	vouch.FromEthAddr = fromAccount.EthAddr

	toIdx, err := common.AccountIdxFromBytes(b[halfLen:])
	if err != nil {
		return nil, common.Wrap(err)
	}
	vouch.ToIdx = toIdx

	// Retrieve to account details
	toAccount, err := GetAccountInTreeDB(sto, toIdx)
	if err != nil {
		return nil, common.Wrap(err)
	}
	vouch.ToEthAddr = toAccount.EthAddr

	return vouch, nil
}

// GetATRoot returns the root of the Account Merkle Tree
func (s *StateDB) GetVTRoot() *big.Int {
	return s.VT.Root().BigInt()
}

// GetVouches returns vouch array from Vouch Merkle Tree
func (s *StateDB) GetVouches() ([]bool, error) {
	num := 1 << (MaxNLevels - 1)

	vouches := make([]bool, num)

	for i := 0; i < num; i++ {
		idx := common.VouchIdx(i)
		_, err := s.GetVouch(idx)
		if err != nil {
			if errors.Is(err, ErrKeyNotFound) {
				vouches[i] = false
				continue
			}
			return nil, fmt.Errorf("failed to get vouch %d: %w", i, err)
		}
		vouches[i] = true
	}

	return vouches, nil
}
