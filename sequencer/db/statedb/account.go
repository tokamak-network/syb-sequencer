package statedb

import (
	"errors"
	"fmt"
	"math/big"

	ethCommon "github.com/ethereum/go-ethereum/common"
	"github.com/iden3/go-merkletree"
	"github.com/iden3/go-merkletree/db"
	"github.com/tokamak-network/syb-sequencer/sequencer/common"
)

var (
	// ErrAccountAlreadyExists is used when CreateAccount is called and the
	// Account already exists
	ErrAccountAlreadyExists = errors.New("cannot CreateAccount because Account already exists")
	// PrefixKeyAccIdx is the key prefix for accountIdx in the db
	PrefixKeyAccIdx = []byte("i:")
	// PrefixKeyAccHash is the key prefix for account hash in the db
	PrefixKeyAccHash = []byte("h:")
	// PrefixKeyAddr is the key prefix for address in the db
	PrefixKeyAddr = []byte("a:")
)

// CreateAccount creates a new Account in the StateDB for the given Idx.  If
// StateDB.MT==nil, MerkleTree is not affected, otherwise updates the
// MerkleTree, returning a CircomProcessorProof.
func (s *StateDB) CreateAccount(idx common.AccountIdx, account *common.Account) (
	*merkletree.CircomProcessorProof, error) {
	cpp, err := CreateAccountInTreeDB(s.db.DB(), s.AT, idx, account)
	if err != nil {
		return cpp, common.Wrap(err)
	}
	// store idx by EthAddr
	err = s.setAccountIdxByEthAddr(idx, account.EthAddr)
	return cpp, common.Wrap(err)
}

// CreateAccountInTreeDB is abstracted from StateDB to be used from StateDB.
// Creates a new Account in the StateDB for the given Idx.  If
// StateDB.AT==nil, MerkleTree is not affected, otherwise updates the
// MerkleTree, returning a CircomProcessorProof.
func CreateAccountInTreeDB(sto db.Storage, mt *merkletree.MerkleTree, idx common.AccountIdx,
	account *common.Account) (*merkletree.CircomProcessorProof, error) {
	// store at the DB the key: v, and value: leaf.Bytes()
	v, err := account.HashValue()
	if err != nil {
		return nil, common.Wrap(err)
	}
	accountBytes, err := account.Bytes()
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
	_, err = tx.Get(append(PrefixKeyAccIdx, idxBytes[:]...))
	if common.Unwrap(err) != db.ErrNotFound {
		return nil, common.Wrap(ErrAccountAlreadyExists)
	}

	err = tx.Put(append(PrefixKeyAccHash, v.Bytes()...), accountBytes[:])
	if err != nil {
		return nil, common.Wrap(err)
	}
	err = tx.Put(append(PrefixKeyAccIdx, idxBytes[:]...), v.Bytes())
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

// UpdateAccount updates the Account in the StateDB for the given Idx.  If
// StateDB.mt==nil, MerkleTree is not affected, otherwise updates the
// MerkleTree, returning a CircomProcessorProof.
func (s *StateDB) UpdateAccount(idx common.AccountIdx, account *common.Account) (
	*merkletree.CircomProcessorProof, error) {
	return UpdateAccountInTreeDB(s.db.DB(), s.AT, idx, account)
}

// UpdateAccountInTreeDB is abstracted from StateDB to be used from StateDB.
// Updates the Account in the StateDB for the given Idx.  If
// StateDB.mt==nil, MerkleTree is not affected, otherwise updates the
// MerkleTree, returning a CircomProcessorProof.
func UpdateAccountInTreeDB(sto db.Storage, mt *merkletree.MerkleTree, idx common.AccountIdx,
	account *common.Account) (*merkletree.CircomProcessorProof, error) {
	// store at the DB the key: v, and value: account.Bytes()
	v, err := account.HashValue()
	if err != nil {
		return nil, common.Wrap(err)
	}
	accountBytes, err := account.Bytes()
	if err != nil {
		return nil, common.Wrap(err)
	}

	tx, err := sto.NewTx()
	if err != nil {
		return nil, common.Wrap(err)
	}
	err = tx.Put(append(PrefixKeyAccHash, v.Bytes()...), accountBytes[:])
	if err != nil {
		return nil, common.Wrap(err)
	}
	idxBytes, err := idx.Bytes()
	if err != nil {
		return nil, common.Wrap(err)
	}
	err = tx.Put(append(PrefixKeyAccIdx, idxBytes[:]...), v.Bytes())
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

// ATGetProof returns the CircomVerifierProof for a given accountIdx
func (s *StateDB) ATGetProof(idx common.AccountIdx) (*merkletree.CircomVerifierProof, error) {
	if s.AT == nil {
		return nil, common.Wrap(ErrStateDBWithoutMT)
	}
	p, err := s.AT.GenerateSCVerifierProof(idx.BigInt(), s.AT.Root())
	if err != nil {
		return nil, common.Wrap(err)
	}
	return p, nil
}

// LastGetAccount is a thread-safe method to query an account in the last
// checkpoint of the StateDB.
func (s *StateDB) LastGetAccount(idx common.AccountIdx) (*common.Account, error) {
	var account *common.Account
	if err := s.LastRead(func(sdb *Last) error {
		var err error
		account, err = sdb.GetAccount(idx)
		return err
	}); err != nil {
		return nil, common.Wrap(err)
	}
	return account, nil
}

// LastATGetRoot returns the root of the underlying Account Tree in the last
// checkpoint of the StateDB.
func (s *StateDB) LastATGetRoot() (*big.Int, error) {
	var root *big.Int
	if err := s.LastRead(func(sdb *Last) error {
		mt, err := merkletree.NewMerkleTree(sdb.DB().WithPrefix(PrefixKeyAT), s.cfg.NLevels)
		if err != nil {
			return common.Wrap(err)
		}
		root = mt.Root().BigInt()
		return nil
	}); err != nil {
		return nil, common.Wrap(err)
	}
	return root, nil
}

// GetAccount returns the account for the given Idx
func (s *StateDB) GetAccount(idx common.AccountIdx) (*common.Account, error) {
	return GetAccountInTreeDB(s.db.DB(), idx)
}

func accountsIter(db db.Storage, fn func(a *common.Account) (bool, error)) error {
	idxDB := db.WithPrefix(PrefixKeyAccIdx)
	if err := idxDB.Iterate(func(k []byte, v []byte) (bool, error) {
		idx, err := common.AccountIdxFromBytes(k)
		if err != nil {
			return false, common.Wrap(err)
		}
		acc, err := GetAccountInTreeDB(db, idx)
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

func getAccounts(db db.Storage) ([]common.Account, error) {
	accs := []common.Account{}
	if err := accountsIter(
		db,
		func(a *common.Account) (bool, error) {
			accs = append(accs, *a)
			return true, nil
		},
	); err != nil {
		return nil, common.Wrap(err)
	}
	return accs, nil
}

// TestGetAccounts returns all the accounts in the db.  Use only in tests.
// Outside tests getting all the accounts is discouraged because it's an
// expensive operation, but if you must do it, use `LastRead()` method to get a
// thread-safe and consistent view of the stateDB.
func (s *StateDB) TestGetAccounts() ([]common.Account, error) {
	return getAccounts(s.db.DB())
}

// GetAccountInTreeDB is abstracted from StateDB to be used from StateDB.
// GetAccount returns the account for the given Idx
func GetAccountInTreeDB(sto db.Storage, idx common.AccountIdx) (*common.Account, error) {
	idxBytes, err := idx.Bytes()
	if err != nil {
		return nil, common.Wrap(err)
	}
	vBytes, err := sto.Get(append(PrefixKeyAccIdx, idxBytes[:]...))
	if err != nil {
		return nil, common.Wrap(err)
	}
	accBytes, err := sto.Get(append(PrefixKeyAccHash, vBytes...))
	if err != nil {
		return nil, common.Wrap(err)
	}
	var b [32 * common.NLeafElems]byte
	copy(b[:], accBytes)
	account, err := common.AccountFromBytes(b)
	if err != nil {
		return nil, common.Wrap(err)
	}
	account.Idx = idx
	return account, nil
}

// CurrentAccountIdx returns the current in-memory CurrentIdx of the StateDB.db
func (s *StateDB) CurrentAccountIdx() common.AccountIdx {
	return s.db.CurrentAccountIdx
}

// SetCurrentAccountIdx stores Idx in the StateDB
func (s *StateDB) SetCurrentAccountIdx(idx common.AccountIdx) error {
	return s.db.SetCurrentAccountIdx(idx)
}

// GetCurrentAccountIdx returns the stored Idx from the localStateDB, which is the
// last Idx used for an Account in the localStateDB.
func (s *StateDB) GetCurrentAccountIdx() (common.AccountIdx, error) {
	return s.db.GetCurrentAccountIdx()
}

// GetATRoot returns the root of the Account Merkle Tree
func (s *StateDB) GetATRoot() *big.Int {
	return s.AT.Root().BigInt()
}

// setAccountIdxByEthAddr stores the given Idx in the StateDB as follows:
// - key: Eth Address, value: idx
// If Idx already exist for the given EthAddr, the remaining Idx will be
// always the smallest one.
func (s *StateDB) setAccountIdxByEthAddr(idx common.AccountIdx, addr ethCommon.Address) error {
	oldIdx, err := s.GetAccountIdxByEthAddr(addr)
	if err == nil {
		// EthAddr already have an Idx
		// check which Idx is smaller
		// if new idx is smaller, store the new one
		// if new idx is bigger, don't store and return, as the used one will be the old
		if idx >= oldIdx {
			fmt.Println("StateDB.setIdxByEthAddr: Idx not stored because there " +
				"already exist a smaller Idx for the given EthAddr")
			return nil
		}
	}

	// store idx for EthAddr assuming that EthAddr still don't have
	// an AccountIdx stored in the DB, and if so, the already stored Idx is
	// bigger than the given one, so should be updated to the new one
	// (smaller)
	tx, err := s.db.DB().NewTx()
	if err != nil {
		return common.Wrap(err)
	}
	idxBytes, err := idx.Bytes()
	if err != nil {
		return common.Wrap(err)
	}

	// store Addr-idx
	err = tx.Put(append(PrefixKeyAddr, addr.Bytes()...), idxBytes[:])
	if err != nil {
		return common.Wrap(err)
	}
	err = tx.Commit()
	if err != nil {
		return common.Wrap(err)
	}
	return nil
}

// GetAccountIdxByEthAddr returns the smallest Idx in the StateDB for the given
// Ethereum Address. Will return common.Idx(0) and error in case that Idx is
// not found in the StateDB.
func (s *StateDB) GetAccountIdxByEthAddr(addr ethCommon.Address) (common.AccountIdx, error) {
	b, err := s.db.DB().Get(append(PrefixKeyAddr, addr.Bytes()...))
	if err != nil {
		return common.AccountIdx(0), common.Wrap(fmt.Errorf("GetIdxByEthAddr: %s: ToEthAddr: %s",
			ErrIdxNotFound, addr.Hex()))
	}
	idx, err := common.AccountIdxFromBytes(b)
	if err != nil {
		return common.AccountIdx(0), common.Wrap(fmt.Errorf("GetIdxByEthAddr: %s: ToEthAddr: %s",
			err, addr.Hex()))
	}
	return idx, nil
}
