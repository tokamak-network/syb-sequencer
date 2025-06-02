// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package bindings

import (
	"errors"
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
	_ = abi.ConvertType
)

// SybilMetaData contains all meta data concerning the Sybil contract.
var SybilMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"function\",\"name\":\"ADMIN_ROLE\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"DEFAULT_ADMIN_ROLE\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"_MIN_BALANCE\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"_hashFinalNode\",\"inputs\":[{\"name\":\"key\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"value\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"_hashNode\",\"inputs\":[{\"name\":\"left\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"right\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"accountInfo\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"balance\",\"type\":\"uint192\",\"internalType\":\"uint192\"},{\"name\":\"idx\",\"type\":\"uint24\",\"internalType\":\"uint24\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"accountRootMap\",\"inputs\":[{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"batchSize\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"deposit\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"exitRootMap\",\"inputs\":[{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"explodeAmount\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"explodeMultiple\",\"inputs\":[{\"name\":\"toEthAddrs\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"forgeBatch\",\"inputs\":[{\"name\":\"newAccountRoot\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"newVouchRoot\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"newScoreRoot\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"proofA\",\"type\":\"uint256[2]\",\"internalType\":\"uint256[2]\"},{\"name\":\"proofB\",\"type\":\"uint256[2][2]\",\"internalType\":\"uint256[2][2]\"},{\"name\":\"proofC\",\"type\":\"uint256[2]\",\"internalType\":\"uint256[2]\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"getQueueLength\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getRoleAdmin\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"grantRole\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"hasRole\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"initialize\",\"inputs\":[{\"name\":\"_verifier\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"maxTx\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"nLevel\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_poseidon2Elements\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_poseidon3Elements\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_adminRole\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"lastAddedTxn\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"lastForgedBatch\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"lastForgedTxn\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"lastIdx\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint24\",\"internalType\":\"uint24\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"proveScoreMerkleProof\",\"inputs\":[{\"name\":\"numScoreRoot\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"idx\",\"type\":\"uint24\",\"internalType\":\"uint24\"},{\"name\":\"score\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"siblings\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"renounceRole\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"callerConfirmation\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"revokeRole\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"scoreRootMap\",\"inputs\":[{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"scoreSnapshots\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"score\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"batchNum\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"scoringRequiredBalance\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"supportsInterface\",\"inputs\":[{\"name\":\"interfaceId\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"txsDataHashMap\",\"inputs\":[{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"unprocessedBatchesMap\",\"inputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"identifier\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"from\",\"type\":\"uint24\",\"internalType\":\"uint24\"},{\"name\":\"to\",\"type\":\"uint24\",\"internalType\":\"uint24\"},{\"name\":\"amount\",\"type\":\"uint128\",\"internalType\":\"uint128\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"unvouch\",\"inputs\":[{\"name\":\"toEthAddr\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"updateExplodeAmount\",\"inputs\":[{\"name\":\"_explodeAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"updateScoringRequiredBalance\",\"inputs\":[{\"name\":\"_scoringRequiredBalance\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"verifier\",\"inputs\":[],\"outputs\":[{\"name\":\"verifierInterface\",\"type\":\"address\",\"internalType\":\"contractIVerifier\"},{\"name\":\"maxTx\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"nLevel\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"vouch\",\"inputs\":[{\"name\":\"toEthAddr\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"vouchRootMap\",\"inputs\":[{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"vouches\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"withdraw\",\"inputs\":[{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"ExplodeAmountUpdated\",\"inputs\":[{\"name\":\"explodeAmount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ForgeBatch\",\"inputs\":[{\"name\":\"lastForgedBatch\",\"type\":\"uint32\",\"indexed\":true,\"internalType\":\"uint32\"},{\"name\":\"lastForgedTxn\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"batchSize\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Initialized\",\"inputs\":[{\"name\":\"version\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RoleAdminChanged\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"previousAdminRole\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"newAdminRole\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RoleGranted\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"sender\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RoleRevoked\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"sender\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ScoringRequiredBalanceUpdated\",\"inputs\":[{\"name\":\"newBalance\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"TxEvent\",\"inputs\":[{\"name\":\"lastAddedTxn\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"},{\"name\":\"identifier\",\"type\":\"uint8\",\"indexed\":true,\"internalType\":\"uint8\"},{\"name\":\"from\",\"type\":\"uint24\",\"indexed\":false,\"internalType\":\"uint24\"},{\"name\":\"to\",\"type\":\"uint24\",\"indexed\":false,\"internalType\":\"uint24\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"AccessControlBadConfirmation\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"AccessControlUnauthorizedAccount\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"neededRole\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"type\":\"error\",\"name\":\"BatchNotFull\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"EthTransferFailed\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InsufficientBalance\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InsufficientETH\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidInitialization\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidPoseidon2Address\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidPoseidon3Address\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidProof\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidVerifierAddress\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"LimitAmountExceeded\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NotInitializing\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NotVouched\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"ReceiverHasZeroBalance\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"SelfVouch\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"SenderHasZeroBalance\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"SmtProofInvalid\",\"inputs\":[]}]",
}

// SybilABI is the input ABI used to generate the binding from.
// Deprecated: Use SybilMetaData.ABI instead.
var SybilABI = SybilMetaData.ABI

// Sybil is an auto generated Go binding around an Ethereum contract.
type Sybil struct {
	SybilCaller     // Read-only binding to the contract
	SybilTransactor // Write-only binding to the contract
	SybilFilterer   // Log filterer for contract events
}

// SybilCaller is an auto generated read-only Go binding around an Ethereum contract.
type SybilCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SybilTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SybilTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SybilFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SybilFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SybilSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SybilSession struct {
	Contract     *Sybil            // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SybilCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SybilCallerSession struct {
	Contract *SybilCaller  // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// SybilTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SybilTransactorSession struct {
	Contract     *SybilTransactor  // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SybilRaw is an auto generated low-level Go binding around an Ethereum contract.
type SybilRaw struct {
	Contract *Sybil // Generic contract binding to access the raw methods on
}

// SybilCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SybilCallerRaw struct {
	Contract *SybilCaller // Generic read-only contract binding to access the raw methods on
}

// SybilTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SybilTransactorRaw struct {
	Contract *SybilTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSybil creates a new instance of Sybil, bound to a specific deployed contract.
func NewSybil(address common.Address, backend bind.ContractBackend) (*Sybil, error) {
	contract, err := bindSybil(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Sybil{SybilCaller: SybilCaller{contract: contract}, SybilTransactor: SybilTransactor{contract: contract}, SybilFilterer: SybilFilterer{contract: contract}}, nil
}

// NewSybilCaller creates a new read-only instance of Sybil, bound to a specific deployed contract.
func NewSybilCaller(address common.Address, caller bind.ContractCaller) (*SybilCaller, error) {
	contract, err := bindSybil(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SybilCaller{contract: contract}, nil
}

// NewSybilTransactor creates a new write-only instance of Sybil, bound to a specific deployed contract.
func NewSybilTransactor(address common.Address, transactor bind.ContractTransactor) (*SybilTransactor, error) {
	contract, err := bindSybil(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SybilTransactor{contract: contract}, nil
}

// NewSybilFilterer creates a new log filterer instance of Sybil, bound to a specific deployed contract.
func NewSybilFilterer(address common.Address, filterer bind.ContractFilterer) (*SybilFilterer, error) {
	contract, err := bindSybil(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SybilFilterer{contract: contract}, nil
}

// bindSybil binds a generic wrapper to an already deployed contract.
func bindSybil(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := SybilMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Sybil *SybilRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Sybil.Contract.SybilCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Sybil *SybilRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Sybil.Contract.SybilTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Sybil *SybilRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Sybil.Contract.SybilTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Sybil *SybilCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Sybil.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Sybil *SybilTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Sybil.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Sybil *SybilTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Sybil.Contract.contract.Transact(opts, method, params...)
}

// ADMINROLE is a free data retrieval call binding the contract method 0x75b238fc.
//
// Solidity: function ADMIN_ROLE() view returns(bytes32)
func (_Sybil *SybilCaller) ADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Sybil.contract.Call(opts, &out, "ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ADMINROLE is a free data retrieval call binding the contract method 0x75b238fc.
//
// Solidity: function ADMIN_ROLE() view returns(bytes32)
func (_Sybil *SybilSession) ADMINROLE() ([32]byte, error) {
	return _Sybil.Contract.ADMINROLE(&_Sybil.CallOpts)
}

// ADMINROLE is a free data retrieval call binding the contract method 0x75b238fc.
//
// Solidity: function ADMIN_ROLE() view returns(bytes32)
func (_Sybil *SybilCallerSession) ADMINROLE() ([32]byte, error) {
	return _Sybil.Contract.ADMINROLE(&_Sybil.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_Sybil *SybilCaller) DEFAULTADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Sybil.contract.Call(opts, &out, "DEFAULT_ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_Sybil *SybilSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _Sybil.Contract.DEFAULTADMINROLE(&_Sybil.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_Sybil *SybilCallerSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _Sybil.Contract.DEFAULTADMINROLE(&_Sybil.CallOpts)
}

// MINBALANCE is a free data retrieval call binding the contract method 0x45ddebae.
//
// Solidity: function _MIN_BALANCE() view returns(uint256)
func (_Sybil *SybilCaller) MINBALANCE(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Sybil.contract.Call(opts, &out, "_MIN_BALANCE")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MINBALANCE is a free data retrieval call binding the contract method 0x45ddebae.
//
// Solidity: function _MIN_BALANCE() view returns(uint256)
func (_Sybil *SybilSession) MINBALANCE() (*big.Int, error) {
	return _Sybil.Contract.MINBALANCE(&_Sybil.CallOpts)
}

// MINBALANCE is a free data retrieval call binding the contract method 0x45ddebae.
//
// Solidity: function _MIN_BALANCE() view returns(uint256)
func (_Sybil *SybilCallerSession) MINBALANCE() (*big.Int, error) {
	return _Sybil.Contract.MINBALANCE(&_Sybil.CallOpts)
}

// HashFinalNode is a free data retrieval call binding the contract method 0xbbe5a375.
//
// Solidity: function _hashFinalNode(uint256 key, uint256 value) view returns(uint256)
func (_Sybil *SybilCaller) HashFinalNode(opts *bind.CallOpts, key *big.Int, value *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Sybil.contract.Call(opts, &out, "_hashFinalNode", key, value)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// HashFinalNode is a free data retrieval call binding the contract method 0xbbe5a375.
//
// Solidity: function _hashFinalNode(uint256 key, uint256 value) view returns(uint256)
func (_Sybil *SybilSession) HashFinalNode(key *big.Int, value *big.Int) (*big.Int, error) {
	return _Sybil.Contract.HashFinalNode(&_Sybil.CallOpts, key, value)
}

// HashFinalNode is a free data retrieval call binding the contract method 0xbbe5a375.
//
// Solidity: function _hashFinalNode(uint256 key, uint256 value) view returns(uint256)
func (_Sybil *SybilCallerSession) HashFinalNode(key *big.Int, value *big.Int) (*big.Int, error) {
	return _Sybil.Contract.HashFinalNode(&_Sybil.CallOpts, key, value)
}

// HashNode is a free data retrieval call binding the contract method 0xc0b55ae4.
//
// Solidity: function _hashNode(uint256 left, uint256 right) view returns(uint256)
func (_Sybil *SybilCaller) HashNode(opts *bind.CallOpts, left *big.Int, right *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Sybil.contract.Call(opts, &out, "_hashNode", left, right)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// HashNode is a free data retrieval call binding the contract method 0xc0b55ae4.
//
// Solidity: function _hashNode(uint256 left, uint256 right) view returns(uint256)
func (_Sybil *SybilSession) HashNode(left *big.Int, right *big.Int) (*big.Int, error) {
	return _Sybil.Contract.HashNode(&_Sybil.CallOpts, left, right)
}

// HashNode is a free data retrieval call binding the contract method 0xc0b55ae4.
//
// Solidity: function _hashNode(uint256 left, uint256 right) view returns(uint256)
func (_Sybil *SybilCallerSession) HashNode(left *big.Int, right *big.Int) (*big.Int, error) {
	return _Sybil.Contract.HashNode(&_Sybil.CallOpts, left, right)
}

// AccountInfo is a free data retrieval call binding the contract method 0xa7310b58.
//
// Solidity: function accountInfo(address ) view returns(uint192 balance, uint24 idx)
func (_Sybil *SybilCaller) AccountInfo(opts *bind.CallOpts, arg0 common.Address) (struct {
	Balance *big.Int
	Idx     *big.Int
}, error) {
	var out []interface{}
	err := _Sybil.contract.Call(opts, &out, "accountInfo", arg0)

	outstruct := new(struct {
		Balance *big.Int
		Idx     *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Balance = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Idx = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// AccountInfo is a free data retrieval call binding the contract method 0xa7310b58.
//
// Solidity: function accountInfo(address ) view returns(uint192 balance, uint24 idx)
func (_Sybil *SybilSession) AccountInfo(arg0 common.Address) (struct {
	Balance *big.Int
	Idx     *big.Int
}, error) {
	return _Sybil.Contract.AccountInfo(&_Sybil.CallOpts, arg0)
}

// AccountInfo is a free data retrieval call binding the contract method 0xa7310b58.
//
// Solidity: function accountInfo(address ) view returns(uint192 balance, uint24 idx)
func (_Sybil *SybilCallerSession) AccountInfo(arg0 common.Address) (struct {
	Balance *big.Int
	Idx     *big.Int
}, error) {
	return _Sybil.Contract.AccountInfo(&_Sybil.CallOpts, arg0)
}

// AccountRootMap is a free data retrieval call binding the contract method 0x795c6167.
//
// Solidity: function accountRootMap(uint32 ) view returns(uint256)
func (_Sybil *SybilCaller) AccountRootMap(opts *bind.CallOpts, arg0 uint32) (*big.Int, error) {
	var out []interface{}
	err := _Sybil.contract.Call(opts, &out, "accountRootMap", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// AccountRootMap is a free data retrieval call binding the contract method 0x795c6167.
//
// Solidity: function accountRootMap(uint32 ) view returns(uint256)
func (_Sybil *SybilSession) AccountRootMap(arg0 uint32) (*big.Int, error) {
	return _Sybil.Contract.AccountRootMap(&_Sybil.CallOpts, arg0)
}

// AccountRootMap is a free data retrieval call binding the contract method 0x795c6167.
//
// Solidity: function accountRootMap(uint32 ) view returns(uint256)
func (_Sybil *SybilCallerSession) AccountRootMap(arg0 uint32) (*big.Int, error) {
	return _Sybil.Contract.AccountRootMap(&_Sybil.CallOpts, arg0)
}

// BatchSize is a free data retrieval call binding the contract method 0xf4daaba1.
//
// Solidity: function batchSize() view returns(uint256)
func (_Sybil *SybilCaller) BatchSize(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Sybil.contract.Call(opts, &out, "batchSize")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BatchSize is a free data retrieval call binding the contract method 0xf4daaba1.
//
// Solidity: function batchSize() view returns(uint256)
func (_Sybil *SybilSession) BatchSize() (*big.Int, error) {
	return _Sybil.Contract.BatchSize(&_Sybil.CallOpts)
}

// BatchSize is a free data retrieval call binding the contract method 0xf4daaba1.
//
// Solidity: function batchSize() view returns(uint256)
func (_Sybil *SybilCallerSession) BatchSize() (*big.Int, error) {
	return _Sybil.Contract.BatchSize(&_Sybil.CallOpts)
}

// ExitRootMap is a free data retrieval call binding the contract method 0xa5e2ec5b.
//
// Solidity: function exitRootMap(uint32 ) view returns(uint256)
func (_Sybil *SybilCaller) ExitRootMap(opts *bind.CallOpts, arg0 uint32) (*big.Int, error) {
	var out []interface{}
	err := _Sybil.contract.Call(opts, &out, "exitRootMap", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ExitRootMap is a free data retrieval call binding the contract method 0xa5e2ec5b.
//
// Solidity: function exitRootMap(uint32 ) view returns(uint256)
func (_Sybil *SybilSession) ExitRootMap(arg0 uint32) (*big.Int, error) {
	return _Sybil.Contract.ExitRootMap(&_Sybil.CallOpts, arg0)
}

// ExitRootMap is a free data retrieval call binding the contract method 0xa5e2ec5b.
//
// Solidity: function exitRootMap(uint32 ) view returns(uint256)
func (_Sybil *SybilCallerSession) ExitRootMap(arg0 uint32) (*big.Int, error) {
	return _Sybil.Contract.ExitRootMap(&_Sybil.CallOpts, arg0)
}

// ExplodeAmount is a free data retrieval call binding the contract method 0x1dbceceb.
//
// Solidity: function explodeAmount() view returns(uint256)
func (_Sybil *SybilCaller) ExplodeAmount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Sybil.contract.Call(opts, &out, "explodeAmount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ExplodeAmount is a free data retrieval call binding the contract method 0x1dbceceb.
//
// Solidity: function explodeAmount() view returns(uint256)
func (_Sybil *SybilSession) ExplodeAmount() (*big.Int, error) {
	return _Sybil.Contract.ExplodeAmount(&_Sybil.CallOpts)
}

// ExplodeAmount is a free data retrieval call binding the contract method 0x1dbceceb.
//
// Solidity: function explodeAmount() view returns(uint256)
func (_Sybil *SybilCallerSession) ExplodeAmount() (*big.Int, error) {
	return _Sybil.Contract.ExplodeAmount(&_Sybil.CallOpts)
}

// GetQueueLength is a free data retrieval call binding the contract method 0xb8f77005.
//
// Solidity: function getQueueLength() view returns(uint256)
func (_Sybil *SybilCaller) GetQueueLength(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Sybil.contract.Call(opts, &out, "getQueueLength")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetQueueLength is a free data retrieval call binding the contract method 0xb8f77005.
//
// Solidity: function getQueueLength() view returns(uint256)
func (_Sybil *SybilSession) GetQueueLength() (*big.Int, error) {
	return _Sybil.Contract.GetQueueLength(&_Sybil.CallOpts)
}

// GetQueueLength is a free data retrieval call binding the contract method 0xb8f77005.
//
// Solidity: function getQueueLength() view returns(uint256)
func (_Sybil *SybilCallerSession) GetQueueLength() (*big.Int, error) {
	return _Sybil.Contract.GetQueueLength(&_Sybil.CallOpts)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_Sybil *SybilCaller) GetRoleAdmin(opts *bind.CallOpts, role [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _Sybil.contract.Call(opts, &out, "getRoleAdmin", role)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_Sybil *SybilSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _Sybil.Contract.GetRoleAdmin(&_Sybil.CallOpts, role)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_Sybil *SybilCallerSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _Sybil.Contract.GetRoleAdmin(&_Sybil.CallOpts, role)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_Sybil *SybilCaller) HasRole(opts *bind.CallOpts, role [32]byte, account common.Address) (bool, error) {
	var out []interface{}
	err := _Sybil.contract.Call(opts, &out, "hasRole", role, account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_Sybil *SybilSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _Sybil.Contract.HasRole(&_Sybil.CallOpts, role, account)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_Sybil *SybilCallerSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _Sybil.Contract.HasRole(&_Sybil.CallOpts, role, account)
}

// LastAddedTxn is a free data retrieval call binding the contract method 0xb1de0ae1.
//
// Solidity: function lastAddedTxn() view returns(uint256)
func (_Sybil *SybilCaller) LastAddedTxn(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Sybil.contract.Call(opts, &out, "lastAddedTxn")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LastAddedTxn is a free data retrieval call binding the contract method 0xb1de0ae1.
//
// Solidity: function lastAddedTxn() view returns(uint256)
func (_Sybil *SybilSession) LastAddedTxn() (*big.Int, error) {
	return _Sybil.Contract.LastAddedTxn(&_Sybil.CallOpts)
}

// LastAddedTxn is a free data retrieval call binding the contract method 0xb1de0ae1.
//
// Solidity: function lastAddedTxn() view returns(uint256)
func (_Sybil *SybilCallerSession) LastAddedTxn() (*big.Int, error) {
	return _Sybil.Contract.LastAddedTxn(&_Sybil.CallOpts)
}

// LastForgedBatch is a free data retrieval call binding the contract method 0x44e0b2ce.
//
// Solidity: function lastForgedBatch() view returns(uint32)
func (_Sybil *SybilCaller) LastForgedBatch(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _Sybil.contract.Call(opts, &out, "lastForgedBatch")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// LastForgedBatch is a free data retrieval call binding the contract method 0x44e0b2ce.
//
// Solidity: function lastForgedBatch() view returns(uint32)
func (_Sybil *SybilSession) LastForgedBatch() (uint32, error) {
	return _Sybil.Contract.LastForgedBatch(&_Sybil.CallOpts)
}

// LastForgedBatch is a free data retrieval call binding the contract method 0x44e0b2ce.
//
// Solidity: function lastForgedBatch() view returns(uint32)
func (_Sybil *SybilCallerSession) LastForgedBatch() (uint32, error) {
	return _Sybil.Contract.LastForgedBatch(&_Sybil.CallOpts)
}

// LastForgedTxn is a free data retrieval call binding the contract method 0x4254b140.
//
// Solidity: function lastForgedTxn() view returns(uint256)
func (_Sybil *SybilCaller) LastForgedTxn(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Sybil.contract.Call(opts, &out, "lastForgedTxn")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LastForgedTxn is a free data retrieval call binding the contract method 0x4254b140.
//
// Solidity: function lastForgedTxn() view returns(uint256)
func (_Sybil *SybilSession) LastForgedTxn() (*big.Int, error) {
	return _Sybil.Contract.LastForgedTxn(&_Sybil.CallOpts)
}

// LastForgedTxn is a free data retrieval call binding the contract method 0x4254b140.
//
// Solidity: function lastForgedTxn() view returns(uint256)
func (_Sybil *SybilCallerSession) LastForgedTxn() (*big.Int, error) {
	return _Sybil.Contract.LastForgedTxn(&_Sybil.CallOpts)
}

// LastIdx is a free data retrieval call binding the contract method 0xd486645c.
//
// Solidity: function lastIdx() view returns(uint24)
func (_Sybil *SybilCaller) LastIdx(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Sybil.contract.Call(opts, &out, "lastIdx")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LastIdx is a free data retrieval call binding the contract method 0xd486645c.
//
// Solidity: function lastIdx() view returns(uint24)
func (_Sybil *SybilSession) LastIdx() (*big.Int, error) {
	return _Sybil.Contract.LastIdx(&_Sybil.CallOpts)
}

// LastIdx is a free data retrieval call binding the contract method 0xd486645c.
//
// Solidity: function lastIdx() view returns(uint24)
func (_Sybil *SybilCallerSession) LastIdx() (*big.Int, error) {
	return _Sybil.Contract.LastIdx(&_Sybil.CallOpts)
}

// ScoreRootMap is a free data retrieval call binding the contract method 0xbd8a4a61.
//
// Solidity: function scoreRootMap(uint32 ) view returns(uint256)
func (_Sybil *SybilCaller) ScoreRootMap(opts *bind.CallOpts, arg0 uint32) (*big.Int, error) {
	var out []interface{}
	err := _Sybil.contract.Call(opts, &out, "scoreRootMap", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ScoreRootMap is a free data retrieval call binding the contract method 0xbd8a4a61.
//
// Solidity: function scoreRootMap(uint32 ) view returns(uint256)
func (_Sybil *SybilSession) ScoreRootMap(arg0 uint32) (*big.Int, error) {
	return _Sybil.Contract.ScoreRootMap(&_Sybil.CallOpts, arg0)
}

// ScoreRootMap is a free data retrieval call binding the contract method 0xbd8a4a61.
//
// Solidity: function scoreRootMap(uint32 ) view returns(uint256)
func (_Sybil *SybilCallerSession) ScoreRootMap(arg0 uint32) (*big.Int, error) {
	return _Sybil.Contract.ScoreRootMap(&_Sybil.CallOpts, arg0)
}

// ScoreSnapshots is a free data retrieval call binding the contract method 0xdf4c5fe2.
//
// Solidity: function scoreSnapshots(address ) view returns(uint32 score, uint32 batchNum)
func (_Sybil *SybilCaller) ScoreSnapshots(opts *bind.CallOpts, arg0 common.Address) (struct {
	Score    uint32
	BatchNum uint32
}, error) {
	var out []interface{}
	err := _Sybil.contract.Call(opts, &out, "scoreSnapshots", arg0)

	outstruct := new(struct {
		Score    uint32
		BatchNum uint32
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Score = *abi.ConvertType(out[0], new(uint32)).(*uint32)
	outstruct.BatchNum = *abi.ConvertType(out[1], new(uint32)).(*uint32)

	return *outstruct, err

}

// ScoreSnapshots is a free data retrieval call binding the contract method 0xdf4c5fe2.
//
// Solidity: function scoreSnapshots(address ) view returns(uint32 score, uint32 batchNum)
func (_Sybil *SybilSession) ScoreSnapshots(arg0 common.Address) (struct {
	Score    uint32
	BatchNum uint32
}, error) {
	return _Sybil.Contract.ScoreSnapshots(&_Sybil.CallOpts, arg0)
}

// ScoreSnapshots is a free data retrieval call binding the contract method 0xdf4c5fe2.
//
// Solidity: function scoreSnapshots(address ) view returns(uint32 score, uint32 batchNum)
func (_Sybil *SybilCallerSession) ScoreSnapshots(arg0 common.Address) (struct {
	Score    uint32
	BatchNum uint32
}, error) {
	return _Sybil.Contract.ScoreSnapshots(&_Sybil.CallOpts, arg0)
}

// ScoringRequiredBalance is a free data retrieval call binding the contract method 0xddf9e619.
//
// Solidity: function scoringRequiredBalance() view returns(uint256)
func (_Sybil *SybilCaller) ScoringRequiredBalance(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Sybil.contract.Call(opts, &out, "scoringRequiredBalance")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ScoringRequiredBalance is a free data retrieval call binding the contract method 0xddf9e619.
//
// Solidity: function scoringRequiredBalance() view returns(uint256)
func (_Sybil *SybilSession) ScoringRequiredBalance() (*big.Int, error) {
	return _Sybil.Contract.ScoringRequiredBalance(&_Sybil.CallOpts)
}

// ScoringRequiredBalance is a free data retrieval call binding the contract method 0xddf9e619.
//
// Solidity: function scoringRequiredBalance() view returns(uint256)
func (_Sybil *SybilCallerSession) ScoringRequiredBalance() (*big.Int, error) {
	return _Sybil.Contract.ScoringRequiredBalance(&_Sybil.CallOpts)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Sybil *SybilCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _Sybil.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Sybil *SybilSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _Sybil.Contract.SupportsInterface(&_Sybil.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Sybil *SybilCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _Sybil.Contract.SupportsInterface(&_Sybil.CallOpts, interfaceId)
}

// TxsDataHashMap is a free data retrieval call binding the contract method 0xf872ecdd.
//
// Solidity: function txsDataHashMap(uint32 ) view returns(bytes32)
func (_Sybil *SybilCaller) TxsDataHashMap(opts *bind.CallOpts, arg0 uint32) ([32]byte, error) {
	var out []interface{}
	err := _Sybil.contract.Call(opts, &out, "txsDataHashMap", arg0)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// TxsDataHashMap is a free data retrieval call binding the contract method 0xf872ecdd.
//
// Solidity: function txsDataHashMap(uint32 ) view returns(bytes32)
func (_Sybil *SybilSession) TxsDataHashMap(arg0 uint32) ([32]byte, error) {
	return _Sybil.Contract.TxsDataHashMap(&_Sybil.CallOpts, arg0)
}

// TxsDataHashMap is a free data retrieval call binding the contract method 0xf872ecdd.
//
// Solidity: function txsDataHashMap(uint32 ) view returns(bytes32)
func (_Sybil *SybilCallerSession) TxsDataHashMap(arg0 uint32) ([32]byte, error) {
	return _Sybil.Contract.TxsDataHashMap(&_Sybil.CallOpts, arg0)
}

// UnprocessedBatchesMap is a free data retrieval call binding the contract method 0xd09dfc99.
//
// Solidity: function unprocessedBatchesMap(uint256 ) view returns(uint8 identifier, uint24 from, uint24 to, uint128 amount)
func (_Sybil *SybilCaller) UnprocessedBatchesMap(opts *bind.CallOpts, arg0 *big.Int) (struct {
	Identifier uint8
	From       *big.Int
	To         *big.Int
	Amount     *big.Int
}, error) {
	var out []interface{}
	err := _Sybil.contract.Call(opts, &out, "unprocessedBatchesMap", arg0)

	outstruct := new(struct {
		Identifier uint8
		From       *big.Int
		To         *big.Int
		Amount     *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Identifier = *abi.ConvertType(out[0], new(uint8)).(*uint8)
	outstruct.From = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.To = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.Amount = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// UnprocessedBatchesMap is a free data retrieval call binding the contract method 0xd09dfc99.
//
// Solidity: function unprocessedBatchesMap(uint256 ) view returns(uint8 identifier, uint24 from, uint24 to, uint128 amount)
func (_Sybil *SybilSession) UnprocessedBatchesMap(arg0 *big.Int) (struct {
	Identifier uint8
	From       *big.Int
	To         *big.Int
	Amount     *big.Int
}, error) {
	return _Sybil.Contract.UnprocessedBatchesMap(&_Sybil.CallOpts, arg0)
}

// UnprocessedBatchesMap is a free data retrieval call binding the contract method 0xd09dfc99.
//
// Solidity: function unprocessedBatchesMap(uint256 ) view returns(uint8 identifier, uint24 from, uint24 to, uint128 amount)
func (_Sybil *SybilCallerSession) UnprocessedBatchesMap(arg0 *big.Int) (struct {
	Identifier uint8
	From       *big.Int
	To         *big.Int
	Amount     *big.Int
}, error) {
	return _Sybil.Contract.UnprocessedBatchesMap(&_Sybil.CallOpts, arg0)
}

// Verifier is a free data retrieval call binding the contract method 0x2b7ac3f3.
//
// Solidity: function verifier() view returns(address verifierInterface, uint256 maxTx, uint256 nLevel)
func (_Sybil *SybilCaller) Verifier(opts *bind.CallOpts) (struct {
	VerifierInterface common.Address
	MaxTx             *big.Int
	NLevel            *big.Int
}, error) {
	var out []interface{}
	err := _Sybil.contract.Call(opts, &out, "verifier")

	outstruct := new(struct {
		VerifierInterface common.Address
		MaxTx             *big.Int
		NLevel            *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.VerifierInterface = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.MaxTx = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.NLevel = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// Verifier is a free data retrieval call binding the contract method 0x2b7ac3f3.
//
// Solidity: function verifier() view returns(address verifierInterface, uint256 maxTx, uint256 nLevel)
func (_Sybil *SybilSession) Verifier() (struct {
	VerifierInterface common.Address
	MaxTx             *big.Int
	NLevel            *big.Int
}, error) {
	return _Sybil.Contract.Verifier(&_Sybil.CallOpts)
}

// Verifier is a free data retrieval call binding the contract method 0x2b7ac3f3.
//
// Solidity: function verifier() view returns(address verifierInterface, uint256 maxTx, uint256 nLevel)
func (_Sybil *SybilCallerSession) Verifier() (struct {
	VerifierInterface common.Address
	MaxTx             *big.Int
	NLevel            *big.Int
}, error) {
	return _Sybil.Contract.Verifier(&_Sybil.CallOpts)
}

// VouchRootMap is a free data retrieval call binding the contract method 0xadacd33b.
//
// Solidity: function vouchRootMap(uint32 ) view returns(uint256)
func (_Sybil *SybilCaller) VouchRootMap(opts *bind.CallOpts, arg0 uint32) (*big.Int, error) {
	var out []interface{}
	err := _Sybil.contract.Call(opts, &out, "vouchRootMap", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// VouchRootMap is a free data retrieval call binding the contract method 0xadacd33b.
//
// Solidity: function vouchRootMap(uint32 ) view returns(uint256)
func (_Sybil *SybilSession) VouchRootMap(arg0 uint32) (*big.Int, error) {
	return _Sybil.Contract.VouchRootMap(&_Sybil.CallOpts, arg0)
}

// VouchRootMap is a free data retrieval call binding the contract method 0xadacd33b.
//
// Solidity: function vouchRootMap(uint32 ) view returns(uint256)
func (_Sybil *SybilCallerSession) VouchRootMap(arg0 uint32) (*big.Int, error) {
	return _Sybil.Contract.VouchRootMap(&_Sybil.CallOpts, arg0)
}

// Vouches is a free data retrieval call binding the contract method 0x0b337be6.
//
// Solidity: function vouches(address , address ) view returns(bool)
func (_Sybil *SybilCaller) Vouches(opts *bind.CallOpts, arg0 common.Address, arg1 common.Address) (bool, error) {
	var out []interface{}
	err := _Sybil.contract.Call(opts, &out, "vouches", arg0, arg1)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Vouches is a free data retrieval call binding the contract method 0x0b337be6.
//
// Solidity: function vouches(address , address ) view returns(bool)
func (_Sybil *SybilSession) Vouches(arg0 common.Address, arg1 common.Address) (bool, error) {
	return _Sybil.Contract.Vouches(&_Sybil.CallOpts, arg0, arg1)
}

// Vouches is a free data retrieval call binding the contract method 0x0b337be6.
//
// Solidity: function vouches(address , address ) view returns(bool)
func (_Sybil *SybilCallerSession) Vouches(arg0 common.Address, arg1 common.Address) (bool, error) {
	return _Sybil.Contract.Vouches(&_Sybil.CallOpts, arg0, arg1)
}

// Deposit is a paid mutator transaction binding the contract method 0xd0e30db0.
//
// Solidity: function deposit() payable returns()
func (_Sybil *SybilTransactor) Deposit(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Sybil.contract.Transact(opts, "deposit")
}

// Deposit is a paid mutator transaction binding the contract method 0xd0e30db0.
//
// Solidity: function deposit() payable returns()
func (_Sybil *SybilSession) Deposit() (*types.Transaction, error) {
	return _Sybil.Contract.Deposit(&_Sybil.TransactOpts)
}

// Deposit is a paid mutator transaction binding the contract method 0xd0e30db0.
//
// Solidity: function deposit() payable returns()
func (_Sybil *SybilTransactorSession) Deposit() (*types.Transaction, error) {
	return _Sybil.Contract.Deposit(&_Sybil.TransactOpts)
}

// ExplodeMultiple is a paid mutator transaction binding the contract method 0xd847d0ca.
//
// Solidity: function explodeMultiple(address[] toEthAddrs) returns()
func (_Sybil *SybilTransactor) ExplodeMultiple(opts *bind.TransactOpts, toEthAddrs []common.Address) (*types.Transaction, error) {
	return _Sybil.contract.Transact(opts, "explodeMultiple", toEthAddrs)
}

// ExplodeMultiple is a paid mutator transaction binding the contract method 0xd847d0ca.
//
// Solidity: function explodeMultiple(address[] toEthAddrs) returns()
func (_Sybil *SybilSession) ExplodeMultiple(toEthAddrs []common.Address) (*types.Transaction, error) {
	return _Sybil.Contract.ExplodeMultiple(&_Sybil.TransactOpts, toEthAddrs)
}

// ExplodeMultiple is a paid mutator transaction binding the contract method 0xd847d0ca.
//
// Solidity: function explodeMultiple(address[] toEthAddrs) returns()
func (_Sybil *SybilTransactorSession) ExplodeMultiple(toEthAddrs []common.Address) (*types.Transaction, error) {
	return _Sybil.Contract.ExplodeMultiple(&_Sybil.TransactOpts, toEthAddrs)
}

// ForgeBatch is a paid mutator transaction binding the contract method 0xcc7e445e.
//
// Solidity: function forgeBatch(uint256 newAccountRoot, uint256 newVouchRoot, uint256 newScoreRoot, uint256[2] proofA, uint256[2][2] proofB, uint256[2] proofC) returns()
func (_Sybil *SybilTransactor) ForgeBatch(opts *bind.TransactOpts, newAccountRoot *big.Int, newVouchRoot *big.Int, newScoreRoot *big.Int, proofA [2]*big.Int, proofB [2][2]*big.Int, proofC [2]*big.Int) (*types.Transaction, error) {
	return _Sybil.contract.Transact(opts, "forgeBatch", newAccountRoot, newVouchRoot, newScoreRoot, proofA, proofB, proofC)
}

// ForgeBatch is a paid mutator transaction binding the contract method 0xcc7e445e.
//
// Solidity: function forgeBatch(uint256 newAccountRoot, uint256 newVouchRoot, uint256 newScoreRoot, uint256[2] proofA, uint256[2][2] proofB, uint256[2] proofC) returns()
func (_Sybil *SybilSession) ForgeBatch(newAccountRoot *big.Int, newVouchRoot *big.Int, newScoreRoot *big.Int, proofA [2]*big.Int, proofB [2][2]*big.Int, proofC [2]*big.Int) (*types.Transaction, error) {
	return _Sybil.Contract.ForgeBatch(&_Sybil.TransactOpts, newAccountRoot, newVouchRoot, newScoreRoot, proofA, proofB, proofC)
}

// ForgeBatch is a paid mutator transaction binding the contract method 0xcc7e445e.
//
// Solidity: function forgeBatch(uint256 newAccountRoot, uint256 newVouchRoot, uint256 newScoreRoot, uint256[2] proofA, uint256[2][2] proofB, uint256[2] proofC) returns()
func (_Sybil *SybilTransactorSession) ForgeBatch(newAccountRoot *big.Int, newVouchRoot *big.Int, newScoreRoot *big.Int, proofA [2]*big.Int, proofB [2][2]*big.Int, proofC [2]*big.Int) (*types.Transaction, error) {
	return _Sybil.Contract.ForgeBatch(&_Sybil.TransactOpts, newAccountRoot, newVouchRoot, newScoreRoot, proofA, proofB, proofC)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_Sybil *SybilTransactor) GrantRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Sybil.contract.Transact(opts, "grantRole", role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_Sybil *SybilSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Sybil.Contract.GrantRole(&_Sybil.TransactOpts, role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_Sybil *SybilTransactorSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Sybil.Contract.GrantRole(&_Sybil.TransactOpts, role, account)
}

// Initialize is a paid mutator transaction binding the contract method 0x11917b1d.
//
// Solidity: function initialize(address _verifier, uint256 maxTx, uint256 nLevel, address _poseidon2Elements, address _poseidon3Elements, address _adminRole) returns()
func (_Sybil *SybilTransactor) Initialize(opts *bind.TransactOpts, _verifier common.Address, maxTx *big.Int, nLevel *big.Int, _poseidon2Elements common.Address, _poseidon3Elements common.Address, _adminRole common.Address) (*types.Transaction, error) {
	return _Sybil.contract.Transact(opts, "initialize", _verifier, maxTx, nLevel, _poseidon2Elements, _poseidon3Elements, _adminRole)
}

// Initialize is a paid mutator transaction binding the contract method 0x11917b1d.
//
// Solidity: function initialize(address _verifier, uint256 maxTx, uint256 nLevel, address _poseidon2Elements, address _poseidon3Elements, address _adminRole) returns()
func (_Sybil *SybilSession) Initialize(_verifier common.Address, maxTx *big.Int, nLevel *big.Int, _poseidon2Elements common.Address, _poseidon3Elements common.Address, _adminRole common.Address) (*types.Transaction, error) {
	return _Sybil.Contract.Initialize(&_Sybil.TransactOpts, _verifier, maxTx, nLevel, _poseidon2Elements, _poseidon3Elements, _adminRole)
}

// Initialize is a paid mutator transaction binding the contract method 0x11917b1d.
//
// Solidity: function initialize(address _verifier, uint256 maxTx, uint256 nLevel, address _poseidon2Elements, address _poseidon3Elements, address _adminRole) returns()
func (_Sybil *SybilTransactorSession) Initialize(_verifier common.Address, maxTx *big.Int, nLevel *big.Int, _poseidon2Elements common.Address, _poseidon3Elements common.Address, _adminRole common.Address) (*types.Transaction, error) {
	return _Sybil.Contract.Initialize(&_Sybil.TransactOpts, _verifier, maxTx, nLevel, _poseidon2Elements, _poseidon3Elements, _adminRole)
}

// ProveScoreMerkleProof is a paid mutator transaction binding the contract method 0x23add439.
//
// Solidity: function proveScoreMerkleProof(uint32 numScoreRoot, uint24 idx, uint32 score, uint256[] siblings) returns()
func (_Sybil *SybilTransactor) ProveScoreMerkleProof(opts *bind.TransactOpts, numScoreRoot uint32, idx *big.Int, score uint32, siblings []*big.Int) (*types.Transaction, error) {
	return _Sybil.contract.Transact(opts, "proveScoreMerkleProof", numScoreRoot, idx, score, siblings)
}

// ProveScoreMerkleProof is a paid mutator transaction binding the contract method 0x23add439.
//
// Solidity: function proveScoreMerkleProof(uint32 numScoreRoot, uint24 idx, uint32 score, uint256[] siblings) returns()
func (_Sybil *SybilSession) ProveScoreMerkleProof(numScoreRoot uint32, idx *big.Int, score uint32, siblings []*big.Int) (*types.Transaction, error) {
	return _Sybil.Contract.ProveScoreMerkleProof(&_Sybil.TransactOpts, numScoreRoot, idx, score, siblings)
}

// ProveScoreMerkleProof is a paid mutator transaction binding the contract method 0x23add439.
//
// Solidity: function proveScoreMerkleProof(uint32 numScoreRoot, uint24 idx, uint32 score, uint256[] siblings) returns()
func (_Sybil *SybilTransactorSession) ProveScoreMerkleProof(numScoreRoot uint32, idx *big.Int, score uint32, siblings []*big.Int) (*types.Transaction, error) {
	return _Sybil.Contract.ProveScoreMerkleProof(&_Sybil.TransactOpts, numScoreRoot, idx, score, siblings)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_Sybil *SybilTransactor) RenounceRole(opts *bind.TransactOpts, role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _Sybil.contract.Transact(opts, "renounceRole", role, callerConfirmation)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_Sybil *SybilSession) RenounceRole(role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _Sybil.Contract.RenounceRole(&_Sybil.TransactOpts, role, callerConfirmation)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_Sybil *SybilTransactorSession) RenounceRole(role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _Sybil.Contract.RenounceRole(&_Sybil.TransactOpts, role, callerConfirmation)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_Sybil *SybilTransactor) RevokeRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Sybil.contract.Transact(opts, "revokeRole", role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_Sybil *SybilSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Sybil.Contract.RevokeRole(&_Sybil.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_Sybil *SybilTransactorSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Sybil.Contract.RevokeRole(&_Sybil.TransactOpts, role, account)
}

// Unvouch is a paid mutator transaction binding the contract method 0xab43cc36.
//
// Solidity: function unvouch(address toEthAddr) returns()
func (_Sybil *SybilTransactor) Unvouch(opts *bind.TransactOpts, toEthAddr common.Address) (*types.Transaction, error) {
	return _Sybil.contract.Transact(opts, "unvouch", toEthAddr)
}

// Unvouch is a paid mutator transaction binding the contract method 0xab43cc36.
//
// Solidity: function unvouch(address toEthAddr) returns()
func (_Sybil *SybilSession) Unvouch(toEthAddr common.Address) (*types.Transaction, error) {
	return _Sybil.Contract.Unvouch(&_Sybil.TransactOpts, toEthAddr)
}

// Unvouch is a paid mutator transaction binding the contract method 0xab43cc36.
//
// Solidity: function unvouch(address toEthAddr) returns()
func (_Sybil *SybilTransactorSession) Unvouch(toEthAddr common.Address) (*types.Transaction, error) {
	return _Sybil.Contract.Unvouch(&_Sybil.TransactOpts, toEthAddr)
}

// UpdateExplodeAmount is a paid mutator transaction binding the contract method 0xaa4f9116.
//
// Solidity: function updateExplodeAmount(uint256 _explodeAmount) returns()
func (_Sybil *SybilTransactor) UpdateExplodeAmount(opts *bind.TransactOpts, _explodeAmount *big.Int) (*types.Transaction, error) {
	return _Sybil.contract.Transact(opts, "updateExplodeAmount", _explodeAmount)
}

// UpdateExplodeAmount is a paid mutator transaction binding the contract method 0xaa4f9116.
//
// Solidity: function updateExplodeAmount(uint256 _explodeAmount) returns()
func (_Sybil *SybilSession) UpdateExplodeAmount(_explodeAmount *big.Int) (*types.Transaction, error) {
	return _Sybil.Contract.UpdateExplodeAmount(&_Sybil.TransactOpts, _explodeAmount)
}

// UpdateExplodeAmount is a paid mutator transaction binding the contract method 0xaa4f9116.
//
// Solidity: function updateExplodeAmount(uint256 _explodeAmount) returns()
func (_Sybil *SybilTransactorSession) UpdateExplodeAmount(_explodeAmount *big.Int) (*types.Transaction, error) {
	return _Sybil.Contract.UpdateExplodeAmount(&_Sybil.TransactOpts, _explodeAmount)
}

// UpdateScoringRequiredBalance is a paid mutator transaction binding the contract method 0x5df89957.
//
// Solidity: function updateScoringRequiredBalance(uint256 _scoringRequiredBalance) returns()
func (_Sybil *SybilTransactor) UpdateScoringRequiredBalance(opts *bind.TransactOpts, _scoringRequiredBalance *big.Int) (*types.Transaction, error) {
	return _Sybil.contract.Transact(opts, "updateScoringRequiredBalance", _scoringRequiredBalance)
}

// UpdateScoringRequiredBalance is a paid mutator transaction binding the contract method 0x5df89957.
//
// Solidity: function updateScoringRequiredBalance(uint256 _scoringRequiredBalance) returns()
func (_Sybil *SybilSession) UpdateScoringRequiredBalance(_scoringRequiredBalance *big.Int) (*types.Transaction, error) {
	return _Sybil.Contract.UpdateScoringRequiredBalance(&_Sybil.TransactOpts, _scoringRequiredBalance)
}

// UpdateScoringRequiredBalance is a paid mutator transaction binding the contract method 0x5df89957.
//
// Solidity: function updateScoringRequiredBalance(uint256 _scoringRequiredBalance) returns()
func (_Sybil *SybilTransactorSession) UpdateScoringRequiredBalance(_scoringRequiredBalance *big.Int) (*types.Transaction, error) {
	return _Sybil.Contract.UpdateScoringRequiredBalance(&_Sybil.TransactOpts, _scoringRequiredBalance)
}

// Vouch is a paid mutator transaction binding the contract method 0xdd66e16b.
//
// Solidity: function vouch(address toEthAddr) returns()
func (_Sybil *SybilTransactor) Vouch(opts *bind.TransactOpts, toEthAddr common.Address) (*types.Transaction, error) {
	return _Sybil.contract.Transact(opts, "vouch", toEthAddr)
}

// Vouch is a paid mutator transaction binding the contract method 0xdd66e16b.
//
// Solidity: function vouch(address toEthAddr) returns()
func (_Sybil *SybilSession) Vouch(toEthAddr common.Address) (*types.Transaction, error) {
	return _Sybil.Contract.Vouch(&_Sybil.TransactOpts, toEthAddr)
}

// Vouch is a paid mutator transaction binding the contract method 0xdd66e16b.
//
// Solidity: function vouch(address toEthAddr) returns()
func (_Sybil *SybilTransactorSession) Vouch(toEthAddr common.Address) (*types.Transaction, error) {
	return _Sybil.Contract.Vouch(&_Sybil.TransactOpts, toEthAddr)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 amount) returns()
func (_Sybil *SybilTransactor) Withdraw(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _Sybil.contract.Transact(opts, "withdraw", amount)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 amount) returns()
func (_Sybil *SybilSession) Withdraw(amount *big.Int) (*types.Transaction, error) {
	return _Sybil.Contract.Withdraw(&_Sybil.TransactOpts, amount)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 amount) returns()
func (_Sybil *SybilTransactorSession) Withdraw(amount *big.Int) (*types.Transaction, error) {
	return _Sybil.Contract.Withdraw(&_Sybil.TransactOpts, amount)
}

// SybilExplodeAmountUpdatedIterator is returned from FilterExplodeAmountUpdated and is used to iterate over the raw logs and unpacked data for ExplodeAmountUpdated events raised by the Sybil contract.
type SybilExplodeAmountUpdatedIterator struct {
	Event *SybilExplodeAmountUpdated // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *SybilExplodeAmountUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SybilExplodeAmountUpdated)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(SybilExplodeAmountUpdated)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *SybilExplodeAmountUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SybilExplodeAmountUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SybilExplodeAmountUpdated represents a ExplodeAmountUpdated event raised by the Sybil contract.
type SybilExplodeAmountUpdated struct {
	ExplodeAmount *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterExplodeAmountUpdated is a free log retrieval operation binding the contract event 0xe4d07ddba7bee2524330d02dcb17edc18e025341f913484d72c85301628b4a79.
//
// Solidity: event ExplodeAmountUpdated(uint256 explodeAmount)
func (_Sybil *SybilFilterer) FilterExplodeAmountUpdated(opts *bind.FilterOpts) (*SybilExplodeAmountUpdatedIterator, error) {

	logs, sub, err := _Sybil.contract.FilterLogs(opts, "ExplodeAmountUpdated")
	if err != nil {
		return nil, err
	}
	return &SybilExplodeAmountUpdatedIterator{contract: _Sybil.contract, event: "ExplodeAmountUpdated", logs: logs, sub: sub}, nil
}

// WatchExplodeAmountUpdated is a free log subscription operation binding the contract event 0xe4d07ddba7bee2524330d02dcb17edc18e025341f913484d72c85301628b4a79.
//
// Solidity: event ExplodeAmountUpdated(uint256 explodeAmount)
func (_Sybil *SybilFilterer) WatchExplodeAmountUpdated(opts *bind.WatchOpts, sink chan<- *SybilExplodeAmountUpdated) (event.Subscription, error) {

	logs, sub, err := _Sybil.contract.WatchLogs(opts, "ExplodeAmountUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SybilExplodeAmountUpdated)
				if err := _Sybil.contract.UnpackLog(event, "ExplodeAmountUpdated", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseExplodeAmountUpdated is a log parse operation binding the contract event 0xe4d07ddba7bee2524330d02dcb17edc18e025341f913484d72c85301628b4a79.
//
// Solidity: event ExplodeAmountUpdated(uint256 explodeAmount)
func (_Sybil *SybilFilterer) ParseExplodeAmountUpdated(log types.Log) (*SybilExplodeAmountUpdated, error) {
	event := new(SybilExplodeAmountUpdated)
	if err := _Sybil.contract.UnpackLog(event, "ExplodeAmountUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SybilForgeBatchIterator is returned from FilterForgeBatch and is used to iterate over the raw logs and unpacked data for ForgeBatch events raised by the Sybil contract.
type SybilForgeBatchIterator struct {
	Event *SybilForgeBatch // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *SybilForgeBatchIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SybilForgeBatch)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(SybilForgeBatch)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *SybilForgeBatchIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SybilForgeBatchIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SybilForgeBatch represents a ForgeBatch event raised by the Sybil contract.
type SybilForgeBatch struct {
	LastForgedBatch uint32
	LastForgedTxn   *big.Int
	BatchSize       *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterForgeBatch is a free log retrieval operation binding the contract event 0x1059ec8d4de0f330dc555e91c43e9a0c8a207c0f0631bfe781d5caccb657e651.
//
// Solidity: event ForgeBatch(uint32 indexed lastForgedBatch, uint256 lastForgedTxn, uint256 batchSize)
func (_Sybil *SybilFilterer) FilterForgeBatch(opts *bind.FilterOpts, lastForgedBatch []uint32) (*SybilForgeBatchIterator, error) {

	var lastForgedBatchRule []interface{}
	for _, lastForgedBatchItem := range lastForgedBatch {
		lastForgedBatchRule = append(lastForgedBatchRule, lastForgedBatchItem)
	}

	logs, sub, err := _Sybil.contract.FilterLogs(opts, "ForgeBatch", lastForgedBatchRule)
	if err != nil {
		return nil, err
	}
	return &SybilForgeBatchIterator{contract: _Sybil.contract, event: "ForgeBatch", logs: logs, sub: sub}, nil
}

// WatchForgeBatch is a free log subscription operation binding the contract event 0x1059ec8d4de0f330dc555e91c43e9a0c8a207c0f0631bfe781d5caccb657e651.
//
// Solidity: event ForgeBatch(uint32 indexed lastForgedBatch, uint256 lastForgedTxn, uint256 batchSize)
func (_Sybil *SybilFilterer) WatchForgeBatch(opts *bind.WatchOpts, sink chan<- *SybilForgeBatch, lastForgedBatch []uint32) (event.Subscription, error) {

	var lastForgedBatchRule []interface{}
	for _, lastForgedBatchItem := range lastForgedBatch {
		lastForgedBatchRule = append(lastForgedBatchRule, lastForgedBatchItem)
	}

	logs, sub, err := _Sybil.contract.WatchLogs(opts, "ForgeBatch", lastForgedBatchRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SybilForgeBatch)
				if err := _Sybil.contract.UnpackLog(event, "ForgeBatch", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseForgeBatch is a log parse operation binding the contract event 0x1059ec8d4de0f330dc555e91c43e9a0c8a207c0f0631bfe781d5caccb657e651.
//
// Solidity: event ForgeBatch(uint32 indexed lastForgedBatch, uint256 lastForgedTxn, uint256 batchSize)
func (_Sybil *SybilFilterer) ParseForgeBatch(log types.Log) (*SybilForgeBatch, error) {
	event := new(SybilForgeBatch)
	if err := _Sybil.contract.UnpackLog(event, "ForgeBatch", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SybilInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the Sybil contract.
type SybilInitializedIterator struct {
	Event *SybilInitialized // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *SybilInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SybilInitialized)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(SybilInitialized)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *SybilInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SybilInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SybilInitialized represents a Initialized event raised by the Sybil contract.
type SybilInitialized struct {
	Version uint64
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_Sybil *SybilFilterer) FilterInitialized(opts *bind.FilterOpts) (*SybilInitializedIterator, error) {

	logs, sub, err := _Sybil.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &SybilInitializedIterator{contract: _Sybil.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_Sybil *SybilFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *SybilInitialized) (event.Subscription, error) {

	logs, sub, err := _Sybil.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SybilInitialized)
				if err := _Sybil.contract.UnpackLog(event, "Initialized", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseInitialized is a log parse operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_Sybil *SybilFilterer) ParseInitialized(log types.Log) (*SybilInitialized, error) {
	event := new(SybilInitialized)
	if err := _Sybil.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SybilRoleAdminChangedIterator is returned from FilterRoleAdminChanged and is used to iterate over the raw logs and unpacked data for RoleAdminChanged events raised by the Sybil contract.
type SybilRoleAdminChangedIterator struct {
	Event *SybilRoleAdminChanged // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *SybilRoleAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SybilRoleAdminChanged)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(SybilRoleAdminChanged)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *SybilRoleAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SybilRoleAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SybilRoleAdminChanged represents a RoleAdminChanged event raised by the Sybil contract.
type SybilRoleAdminChanged struct {
	Role              [32]byte
	PreviousAdminRole [32]byte
	NewAdminRole      [32]byte
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterRoleAdminChanged is a free log retrieval operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_Sybil *SybilFilterer) FilterRoleAdminChanged(opts *bind.FilterOpts, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (*SybilRoleAdminChangedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var previousAdminRoleRule []interface{}
	for _, previousAdminRoleItem := range previousAdminRole {
		previousAdminRoleRule = append(previousAdminRoleRule, previousAdminRoleItem)
	}
	var newAdminRoleRule []interface{}
	for _, newAdminRoleItem := range newAdminRole {
		newAdminRoleRule = append(newAdminRoleRule, newAdminRoleItem)
	}

	logs, sub, err := _Sybil.contract.FilterLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return &SybilRoleAdminChangedIterator{contract: _Sybil.contract, event: "RoleAdminChanged", logs: logs, sub: sub}, nil
}

// WatchRoleAdminChanged is a free log subscription operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_Sybil *SybilFilterer) WatchRoleAdminChanged(opts *bind.WatchOpts, sink chan<- *SybilRoleAdminChanged, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var previousAdminRoleRule []interface{}
	for _, previousAdminRoleItem := range previousAdminRole {
		previousAdminRoleRule = append(previousAdminRoleRule, previousAdminRoleItem)
	}
	var newAdminRoleRule []interface{}
	for _, newAdminRoleItem := range newAdminRole {
		newAdminRoleRule = append(newAdminRoleRule, newAdminRoleItem)
	}

	logs, sub, err := _Sybil.contract.WatchLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SybilRoleAdminChanged)
				if err := _Sybil.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseRoleAdminChanged is a log parse operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_Sybil *SybilFilterer) ParseRoleAdminChanged(log types.Log) (*SybilRoleAdminChanged, error) {
	event := new(SybilRoleAdminChanged)
	if err := _Sybil.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SybilRoleGrantedIterator is returned from FilterRoleGranted and is used to iterate over the raw logs and unpacked data for RoleGranted events raised by the Sybil contract.
type SybilRoleGrantedIterator struct {
	Event *SybilRoleGranted // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *SybilRoleGrantedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SybilRoleGranted)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(SybilRoleGranted)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *SybilRoleGrantedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SybilRoleGrantedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SybilRoleGranted represents a RoleGranted event raised by the Sybil contract.
type SybilRoleGranted struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleGranted is a free log retrieval operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_Sybil *SybilFilterer) FilterRoleGranted(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*SybilRoleGrantedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _Sybil.contract.FilterLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &SybilRoleGrantedIterator{contract: _Sybil.contract, event: "RoleGranted", logs: logs, sub: sub}, nil
}

// WatchRoleGranted is a free log subscription operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_Sybil *SybilFilterer) WatchRoleGranted(opts *bind.WatchOpts, sink chan<- *SybilRoleGranted, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _Sybil.contract.WatchLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SybilRoleGranted)
				if err := _Sybil.contract.UnpackLog(event, "RoleGranted", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseRoleGranted is a log parse operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_Sybil *SybilFilterer) ParseRoleGranted(log types.Log) (*SybilRoleGranted, error) {
	event := new(SybilRoleGranted)
	if err := _Sybil.contract.UnpackLog(event, "RoleGranted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SybilRoleRevokedIterator is returned from FilterRoleRevoked and is used to iterate over the raw logs and unpacked data for RoleRevoked events raised by the Sybil contract.
type SybilRoleRevokedIterator struct {
	Event *SybilRoleRevoked // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *SybilRoleRevokedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SybilRoleRevoked)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(SybilRoleRevoked)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *SybilRoleRevokedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SybilRoleRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SybilRoleRevoked represents a RoleRevoked event raised by the Sybil contract.
type SybilRoleRevoked struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleRevoked is a free log retrieval operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_Sybil *SybilFilterer) FilterRoleRevoked(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*SybilRoleRevokedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _Sybil.contract.FilterLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &SybilRoleRevokedIterator{contract: _Sybil.contract, event: "RoleRevoked", logs: logs, sub: sub}, nil
}

// WatchRoleRevoked is a free log subscription operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_Sybil *SybilFilterer) WatchRoleRevoked(opts *bind.WatchOpts, sink chan<- *SybilRoleRevoked, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _Sybil.contract.WatchLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SybilRoleRevoked)
				if err := _Sybil.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseRoleRevoked is a log parse operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_Sybil *SybilFilterer) ParseRoleRevoked(log types.Log) (*SybilRoleRevoked, error) {
	event := new(SybilRoleRevoked)
	if err := _Sybil.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SybilScoringRequiredBalanceUpdatedIterator is returned from FilterScoringRequiredBalanceUpdated and is used to iterate over the raw logs and unpacked data for ScoringRequiredBalanceUpdated events raised by the Sybil contract.
type SybilScoringRequiredBalanceUpdatedIterator struct {
	Event *SybilScoringRequiredBalanceUpdated // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *SybilScoringRequiredBalanceUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SybilScoringRequiredBalanceUpdated)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(SybilScoringRequiredBalanceUpdated)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *SybilScoringRequiredBalanceUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SybilScoringRequiredBalanceUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SybilScoringRequiredBalanceUpdated represents a ScoringRequiredBalanceUpdated event raised by the Sybil contract.
type SybilScoringRequiredBalanceUpdated struct {
	NewBalance *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterScoringRequiredBalanceUpdated is a free log retrieval operation binding the contract event 0x18b39a199aab6d2ebd7c7d0d55c6462e08c66ea0a6e1a486df135a27cb591e4e.
//
// Solidity: event ScoringRequiredBalanceUpdated(uint256 newBalance)
func (_Sybil *SybilFilterer) FilterScoringRequiredBalanceUpdated(opts *bind.FilterOpts) (*SybilScoringRequiredBalanceUpdatedIterator, error) {

	logs, sub, err := _Sybil.contract.FilterLogs(opts, "ScoringRequiredBalanceUpdated")
	if err != nil {
		return nil, err
	}
	return &SybilScoringRequiredBalanceUpdatedIterator{contract: _Sybil.contract, event: "ScoringRequiredBalanceUpdated", logs: logs, sub: sub}, nil
}

// WatchScoringRequiredBalanceUpdated is a free log subscription operation binding the contract event 0x18b39a199aab6d2ebd7c7d0d55c6462e08c66ea0a6e1a486df135a27cb591e4e.
//
// Solidity: event ScoringRequiredBalanceUpdated(uint256 newBalance)
func (_Sybil *SybilFilterer) WatchScoringRequiredBalanceUpdated(opts *bind.WatchOpts, sink chan<- *SybilScoringRequiredBalanceUpdated) (event.Subscription, error) {

	logs, sub, err := _Sybil.contract.WatchLogs(opts, "ScoringRequiredBalanceUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SybilScoringRequiredBalanceUpdated)
				if err := _Sybil.contract.UnpackLog(event, "ScoringRequiredBalanceUpdated", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseScoringRequiredBalanceUpdated is a log parse operation binding the contract event 0x18b39a199aab6d2ebd7c7d0d55c6462e08c66ea0a6e1a486df135a27cb591e4e.
//
// Solidity: event ScoringRequiredBalanceUpdated(uint256 newBalance)
func (_Sybil *SybilFilterer) ParseScoringRequiredBalanceUpdated(log types.Log) (*SybilScoringRequiredBalanceUpdated, error) {
	event := new(SybilScoringRequiredBalanceUpdated)
	if err := _Sybil.contract.UnpackLog(event, "ScoringRequiredBalanceUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SybilTxEventIterator is returned from FilterTxEvent and is used to iterate over the raw logs and unpacked data for TxEvent events raised by the Sybil contract.
type SybilTxEventIterator struct {
	Event *SybilTxEvent // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *SybilTxEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SybilTxEvent)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(SybilTxEvent)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *SybilTxEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SybilTxEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SybilTxEvent represents a TxEvent event raised by the Sybil contract.
type SybilTxEvent struct {
	LastAddedTxn *big.Int
	Identifier   uint8
	From         *big.Int
	To           *big.Int
	Amount       *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterTxEvent is a free log retrieval operation binding the contract event 0x2bc985344d14a1f1151b7d0e8f07e9277768fcf9a34ef9f07ec88f25a5cd2668.
//
// Solidity: event TxEvent(uint256 indexed lastAddedTxn, uint8 indexed identifier, uint24 from, uint24 to, uint256 amount)
func (_Sybil *SybilFilterer) FilterTxEvent(opts *bind.FilterOpts, lastAddedTxn []*big.Int, identifier []uint8) (*SybilTxEventIterator, error) {

	var lastAddedTxnRule []interface{}
	for _, lastAddedTxnItem := range lastAddedTxn {
		lastAddedTxnRule = append(lastAddedTxnRule, lastAddedTxnItem)
	}
	var identifierRule []interface{}
	for _, identifierItem := range identifier {
		identifierRule = append(identifierRule, identifierItem)
	}

	logs, sub, err := _Sybil.contract.FilterLogs(opts, "TxEvent", lastAddedTxnRule, identifierRule)
	if err != nil {
		return nil, err
	}
	return &SybilTxEventIterator{contract: _Sybil.contract, event: "TxEvent", logs: logs, sub: sub}, nil
}

// WatchTxEvent is a free log subscription operation binding the contract event 0x2bc985344d14a1f1151b7d0e8f07e9277768fcf9a34ef9f07ec88f25a5cd2668.
//
// Solidity: event TxEvent(uint256 indexed lastAddedTxn, uint8 indexed identifier, uint24 from, uint24 to, uint256 amount)
func (_Sybil *SybilFilterer) WatchTxEvent(opts *bind.WatchOpts, sink chan<- *SybilTxEvent, lastAddedTxn []*big.Int, identifier []uint8) (event.Subscription, error) {

	var lastAddedTxnRule []interface{}
	for _, lastAddedTxnItem := range lastAddedTxn {
		lastAddedTxnRule = append(lastAddedTxnRule, lastAddedTxnItem)
	}
	var identifierRule []interface{}
	for _, identifierItem := range identifier {
		identifierRule = append(identifierRule, identifierItem)
	}

	logs, sub, err := _Sybil.contract.WatchLogs(opts, "TxEvent", lastAddedTxnRule, identifierRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SybilTxEvent)
				if err := _Sybil.contract.UnpackLog(event, "TxEvent", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseTxEvent is a log parse operation binding the contract event 0x2bc985344d14a1f1151b7d0e8f07e9277768fcf9a34ef9f07ec88f25a5cd2668.
//
// Solidity: event TxEvent(uint256 indexed lastAddedTxn, uint8 indexed identifier, uint24 from, uint24 to, uint256 amount)
func (_Sybil *SybilFilterer) ParseTxEvent(log types.Log) (*SybilTxEvent, error) {
	event := new(SybilTxEvent)
	if err := _Sybil.contract.UnpackLog(event, "TxEvent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
