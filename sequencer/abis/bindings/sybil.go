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

// BindingsMetaData contains all meta data concerning the Bindings contract.
var BindingsMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"AccessControlBadConfirmation\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"neededRole\",\"type\":\"bytes32\"}],\"name\":\"AccessControlUnauthorizedAccount\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"EthTransferFailed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InsufficientBalance\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InsufficientETH\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidInitialization\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"elementType\",\"type\":\"string\"}],\"name\":\"InvalidPoseidonAddress\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidProof\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidVerifierAddress\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"LimitAmountExceeded\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotInitializing\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"NotVouched\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ReceiverHasZeroBalance\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"SenderHasZeroBalance\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"SmtProofInvalid\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"explodeAmount\",\"type\":\"uint256\"}],\"name\":\"ExplodeAmountUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint32\",\"name\":\"batchNum\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"uint16\",\"name\":\"l1UserTxsLen\",\"type\":\"uint16\"}],\"name\":\"ForgeBatch\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"version\",\"type\":\"uint64\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint32\",\"name\":\"queueIndex\",\"type\":\"uint32\"},{\"indexed\":true,\"internalType\":\"uint8\",\"name\":\"position\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"l1UserTx\",\"type\":\"bytes\"}],\"name\":\"L1UserTxEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"minBalance\",\"type\":\"uint256\"}],\"name\":\"MinBalanceUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"previousAdminRole\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"newAdminRole\",\"type\":\"bytes32\"}],\"name\":\"RoleAdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleGranted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleRevoked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint48\",\"name\":\"idx\",\"type\":\"uint48\"},{\"indexed\":true,\"internalType\":\"uint32\",\"name\":\"numExitRoot\",\"type\":\"uint32\"}],\"name\":\"WithdrawEvent\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"DEFAULT_ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"key\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"_hashFinalNode\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"left\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"right\",\"type\":\"uint256\"}],\"name\":\"_hashNode\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"name\":\"accountRootMap\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"balances\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"currentFillingBatch\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"deposit\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"},{\"internalType\":\"uint48\",\"name\":\"\",\"type\":\"uint48\"}],\"name\":\"exitNullifierMap\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"name\":\"exitRootMap\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"explodeAmount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"toEthAddrs\",\"type\":\"address[]\"}],\"name\":\"explodeMultiple\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newAccountRoot\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"newVouchRoot\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"newScoreRoot\",\"type\":\"uint256\"},{\"internalType\":\"uint256[2]\",\"name\":\"proofA\",\"type\":\"uint256[2]\"},{\"internalType\":\"uint256[2][2]\",\"name\":\"proofB\",\"type\":\"uint256[2][2]\"},{\"internalType\":\"uint256[2]\",\"name\":\"proofC\",\"type\":\"uint256[2]\"}],\"name\":\"forgeBatch\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getQueueLength\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleAdmin\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"grantRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"hasRole\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_verifier\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"maxTx\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nLevel\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_poseidon2Elements\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_poseidon3Elements\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_adminRole\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lastForgedBatch\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"minBalance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"numScoreRoot\",\"type\":\"uint32\"},{\"internalType\":\"uint24\",\"name\":\"idx\",\"type\":\"uint24\"},{\"internalType\":\"uint32\",\"name\":\"score\",\"type\":\"uint32\"},{\"internalType\":\"uint256[]\",\"name\":\"siblings\",\"type\":\"uint256[]\"}],\"name\":\"proveScoreMerkleProof\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"callerConfirmation\",\"type\":\"address\"}],\"name\":\"renounceRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"revokeRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"name\":\"scoreRootMap\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"scoreSnapshots\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"score\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"batchNum\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"name\":\"txsDataHashMap\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"name\":\"unprocessedBatchesMap\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"toEthAddr\",\"type\":\"address\"}],\"name\":\"unvouch\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_explodeAmount\",\"type\":\"uint256\"}],\"name\":\"updateExplodeAmount\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_minBalance\",\"type\":\"uint256\"}],\"name\":\"updateMinBalance\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"verifier\",\"outputs\":[{\"internalType\":\"contractIVerifier\",\"name\":\"verifierInterface\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"maxTx\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nLevel\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"toEthAddr\",\"type\":\"address\"}],\"name\":\"vouch\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"name\":\"vouchRootMap\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"vouches\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// BindingsABI is the input ABI used to generate the binding from.
// Deprecated: Use BindingsMetaData.ABI instead.
var BindingsABI = BindingsMetaData.ABI

// Bindings is an auto generated Go binding around an Ethereum contract.
type Bindings struct {
	BindingsCaller     // Read-only binding to the contract
	BindingsTransactor // Write-only binding to the contract
	BindingsFilterer   // Log filterer for contract events
}

// BindingsCaller is an auto generated read-only Go binding around an Ethereum contract.
type BindingsCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BindingsTransactor is an auto generated write-only Go binding around an Ethereum contract.
type BindingsTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BindingsFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type BindingsFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BindingsSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type BindingsSession struct {
	Contract     *Bindings         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BindingsCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type BindingsCallerSession struct {
	Contract *BindingsCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// BindingsTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type BindingsTransactorSession struct {
	Contract     *BindingsTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// BindingsRaw is an auto generated low-level Go binding around an Ethereum contract.
type BindingsRaw struct {
	Contract *Bindings // Generic contract binding to access the raw methods on
}

// BindingsCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type BindingsCallerRaw struct {
	Contract *BindingsCaller // Generic read-only contract binding to access the raw methods on
}

// BindingsTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type BindingsTransactorRaw struct {
	Contract *BindingsTransactor // Generic write-only contract binding to access the raw methods on
}

// NewBindings creates a new instance of Bindings, bound to a specific deployed contract.
func NewBindings(address common.Address, backend bind.ContractBackend) (*Bindings, error) {
	contract, err := bindBindings(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Bindings{BindingsCaller: BindingsCaller{contract: contract}, BindingsTransactor: BindingsTransactor{contract: contract}, BindingsFilterer: BindingsFilterer{contract: contract}}, nil
}

// NewBindingsCaller creates a new read-only instance of Bindings, bound to a specific deployed contract.
func NewBindingsCaller(address common.Address, caller bind.ContractCaller) (*BindingsCaller, error) {
	contract, err := bindBindings(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &BindingsCaller{contract: contract}, nil
}

// NewBindingsTransactor creates a new write-only instance of Bindings, bound to a specific deployed contract.
func NewBindingsTransactor(address common.Address, transactor bind.ContractTransactor) (*BindingsTransactor, error) {
	contract, err := bindBindings(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &BindingsTransactor{contract: contract}, nil
}

// NewBindingsFilterer creates a new log filterer instance of Bindings, bound to a specific deployed contract.
func NewBindingsFilterer(address common.Address, filterer bind.ContractFilterer) (*BindingsFilterer, error) {
	contract, err := bindBindings(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &BindingsFilterer{contract: contract}, nil
}

// bindBindings binds a generic wrapper to an already deployed contract.
func bindBindings(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := BindingsMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Bindings *BindingsRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Bindings.Contract.BindingsCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Bindings *BindingsRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Bindings.Contract.BindingsTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Bindings *BindingsRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Bindings.Contract.BindingsTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Bindings *BindingsCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Bindings.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Bindings *BindingsTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Bindings.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Bindings *BindingsTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Bindings.Contract.contract.Transact(opts, method, params...)
}

// ADMINROLE is a free data retrieval call binding the contract method 0x75b238fc.
//
// Solidity: function ADMIN_ROLE() view returns(bytes32)
func (_Bindings *BindingsCaller) ADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Bindings.contract.Call(opts, &out, "ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ADMINROLE is a free data retrieval call binding the contract method 0x75b238fc.
//
// Solidity: function ADMIN_ROLE() view returns(bytes32)
func (_Bindings *BindingsSession) ADMINROLE() ([32]byte, error) {
	return _Bindings.Contract.ADMINROLE(&_Bindings.CallOpts)
}

// ADMINROLE is a free data retrieval call binding the contract method 0x75b238fc.
//
// Solidity: function ADMIN_ROLE() view returns(bytes32)
func (_Bindings *BindingsCallerSession) ADMINROLE() ([32]byte, error) {
	return _Bindings.Contract.ADMINROLE(&_Bindings.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_Bindings *BindingsCaller) DEFAULTADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Bindings.contract.Call(opts, &out, "DEFAULT_ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_Bindings *BindingsSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _Bindings.Contract.DEFAULTADMINROLE(&_Bindings.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_Bindings *BindingsCallerSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _Bindings.Contract.DEFAULTADMINROLE(&_Bindings.CallOpts)
}

// HashFinalNode is a free data retrieval call binding the contract method 0xbbe5a375.
//
// Solidity: function _hashFinalNode(uint256 key, uint256 value) view returns(uint256)
func (_Bindings *BindingsCaller) HashFinalNode(opts *bind.CallOpts, key *big.Int, value *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Bindings.contract.Call(opts, &out, "_hashFinalNode", key, value)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// HashFinalNode is a free data retrieval call binding the contract method 0xbbe5a375.
//
// Solidity: function _hashFinalNode(uint256 key, uint256 value) view returns(uint256)
func (_Bindings *BindingsSession) HashFinalNode(key *big.Int, value *big.Int) (*big.Int, error) {
	return _Bindings.Contract.HashFinalNode(&_Bindings.CallOpts, key, value)
}

// HashFinalNode is a free data retrieval call binding the contract method 0xbbe5a375.
//
// Solidity: function _hashFinalNode(uint256 key, uint256 value) view returns(uint256)
func (_Bindings *BindingsCallerSession) HashFinalNode(key *big.Int, value *big.Int) (*big.Int, error) {
	return _Bindings.Contract.HashFinalNode(&_Bindings.CallOpts, key, value)
}

// HashNode is a free data retrieval call binding the contract method 0xc0b55ae4.
//
// Solidity: function _hashNode(uint256 left, uint256 right) view returns(uint256)
func (_Bindings *BindingsCaller) HashNode(opts *bind.CallOpts, left *big.Int, right *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Bindings.contract.Call(opts, &out, "_hashNode", left, right)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// HashNode is a free data retrieval call binding the contract method 0xc0b55ae4.
//
// Solidity: function _hashNode(uint256 left, uint256 right) view returns(uint256)
func (_Bindings *BindingsSession) HashNode(left *big.Int, right *big.Int) (*big.Int, error) {
	return _Bindings.Contract.HashNode(&_Bindings.CallOpts, left, right)
}

// HashNode is a free data retrieval call binding the contract method 0xc0b55ae4.
//
// Solidity: function _hashNode(uint256 left, uint256 right) view returns(uint256)
func (_Bindings *BindingsCallerSession) HashNode(left *big.Int, right *big.Int) (*big.Int, error) {
	return _Bindings.Contract.HashNode(&_Bindings.CallOpts, left, right)
}

// AccountRootMap is a free data retrieval call binding the contract method 0x795c6167.
//
// Solidity: function accountRootMap(uint32 ) view returns(uint256)
func (_Bindings *BindingsCaller) AccountRootMap(opts *bind.CallOpts, arg0 uint32) (*big.Int, error) {
	var out []interface{}
	err := _Bindings.contract.Call(opts, &out, "accountRootMap", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// AccountRootMap is a free data retrieval call binding the contract method 0x795c6167.
//
// Solidity: function accountRootMap(uint32 ) view returns(uint256)
func (_Bindings *BindingsSession) AccountRootMap(arg0 uint32) (*big.Int, error) {
	return _Bindings.Contract.AccountRootMap(&_Bindings.CallOpts, arg0)
}

// AccountRootMap is a free data retrieval call binding the contract method 0x795c6167.
//
// Solidity: function accountRootMap(uint32 ) view returns(uint256)
func (_Bindings *BindingsCallerSession) AccountRootMap(arg0 uint32) (*big.Int, error) {
	return _Bindings.Contract.AccountRootMap(&_Bindings.CallOpts, arg0)
}

// Balances is a free data retrieval call binding the contract method 0x27e235e3.
//
// Solidity: function balances(address ) view returns(uint256)
func (_Bindings *BindingsCaller) Balances(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Bindings.contract.Call(opts, &out, "balances", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Balances is a free data retrieval call binding the contract method 0x27e235e3.
//
// Solidity: function balances(address ) view returns(uint256)
func (_Bindings *BindingsSession) Balances(arg0 common.Address) (*big.Int, error) {
	return _Bindings.Contract.Balances(&_Bindings.CallOpts, arg0)
}

// Balances is a free data retrieval call binding the contract method 0x27e235e3.
//
// Solidity: function balances(address ) view returns(uint256)
func (_Bindings *BindingsCallerSession) Balances(arg0 common.Address) (*big.Int, error) {
	return _Bindings.Contract.Balances(&_Bindings.CallOpts, arg0)
}

// CurrentFillingBatch is a free data retrieval call binding the contract method 0xc25d5789.
//
// Solidity: function currentFillingBatch() view returns(uint32)
func (_Bindings *BindingsCaller) CurrentFillingBatch(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _Bindings.contract.Call(opts, &out, "currentFillingBatch")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// CurrentFillingBatch is a free data retrieval call binding the contract method 0xc25d5789.
//
// Solidity: function currentFillingBatch() view returns(uint32)
func (_Bindings *BindingsSession) CurrentFillingBatch() (uint32, error) {
	return _Bindings.Contract.CurrentFillingBatch(&_Bindings.CallOpts)
}

// CurrentFillingBatch is a free data retrieval call binding the contract method 0xc25d5789.
//
// Solidity: function currentFillingBatch() view returns(uint32)
func (_Bindings *BindingsCallerSession) CurrentFillingBatch() (uint32, error) {
	return _Bindings.Contract.CurrentFillingBatch(&_Bindings.CallOpts)
}

// ExitNullifierMap is a free data retrieval call binding the contract method 0xf84f92ee.
//
// Solidity: function exitNullifierMap(uint32 , uint48 ) view returns(bool)
func (_Bindings *BindingsCaller) ExitNullifierMap(opts *bind.CallOpts, arg0 uint32, arg1 *big.Int) (bool, error) {
	var out []interface{}
	err := _Bindings.contract.Call(opts, &out, "exitNullifierMap", arg0, arg1)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// ExitNullifierMap is a free data retrieval call binding the contract method 0xf84f92ee.
//
// Solidity: function exitNullifierMap(uint32 , uint48 ) view returns(bool)
func (_Bindings *BindingsSession) ExitNullifierMap(arg0 uint32, arg1 *big.Int) (bool, error) {
	return _Bindings.Contract.ExitNullifierMap(&_Bindings.CallOpts, arg0, arg1)
}

// ExitNullifierMap is a free data retrieval call binding the contract method 0xf84f92ee.
//
// Solidity: function exitNullifierMap(uint32 , uint48 ) view returns(bool)
func (_Bindings *BindingsCallerSession) ExitNullifierMap(arg0 uint32, arg1 *big.Int) (bool, error) {
	return _Bindings.Contract.ExitNullifierMap(&_Bindings.CallOpts, arg0, arg1)
}

// ExitRootMap is a free data retrieval call binding the contract method 0xa5e2ec5b.
//
// Solidity: function exitRootMap(uint32 ) view returns(uint256)
func (_Bindings *BindingsCaller) ExitRootMap(opts *bind.CallOpts, arg0 uint32) (*big.Int, error) {
	var out []interface{}
	err := _Bindings.contract.Call(opts, &out, "exitRootMap", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ExitRootMap is a free data retrieval call binding the contract method 0xa5e2ec5b.
//
// Solidity: function exitRootMap(uint32 ) view returns(uint256)
func (_Bindings *BindingsSession) ExitRootMap(arg0 uint32) (*big.Int, error) {
	return _Bindings.Contract.ExitRootMap(&_Bindings.CallOpts, arg0)
}

// ExitRootMap is a free data retrieval call binding the contract method 0xa5e2ec5b.
//
// Solidity: function exitRootMap(uint32 ) view returns(uint256)
func (_Bindings *BindingsCallerSession) ExitRootMap(arg0 uint32) (*big.Int, error) {
	return _Bindings.Contract.ExitRootMap(&_Bindings.CallOpts, arg0)
}

// ExplodeAmount is a free data retrieval call binding the contract method 0x1dbceceb.
//
// Solidity: function explodeAmount() view returns(uint256)
func (_Bindings *BindingsCaller) ExplodeAmount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Bindings.contract.Call(opts, &out, "explodeAmount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ExplodeAmount is a free data retrieval call binding the contract method 0x1dbceceb.
//
// Solidity: function explodeAmount() view returns(uint256)
func (_Bindings *BindingsSession) ExplodeAmount() (*big.Int, error) {
	return _Bindings.Contract.ExplodeAmount(&_Bindings.CallOpts)
}

// ExplodeAmount is a free data retrieval call binding the contract method 0x1dbceceb.
//
// Solidity: function explodeAmount() view returns(uint256)
func (_Bindings *BindingsCallerSession) ExplodeAmount() (*big.Int, error) {
	return _Bindings.Contract.ExplodeAmount(&_Bindings.CallOpts)
}

// GetQueueLength is a free data retrieval call binding the contract method 0xb8f77005.
//
// Solidity: function getQueueLength() view returns(uint32)
func (_Bindings *BindingsCaller) GetQueueLength(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _Bindings.contract.Call(opts, &out, "getQueueLength")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// GetQueueLength is a free data retrieval call binding the contract method 0xb8f77005.
//
// Solidity: function getQueueLength() view returns(uint32)
func (_Bindings *BindingsSession) GetQueueLength() (uint32, error) {
	return _Bindings.Contract.GetQueueLength(&_Bindings.CallOpts)
}

// GetQueueLength is a free data retrieval call binding the contract method 0xb8f77005.
//
// Solidity: function getQueueLength() view returns(uint32)
func (_Bindings *BindingsCallerSession) GetQueueLength() (uint32, error) {
	return _Bindings.Contract.GetQueueLength(&_Bindings.CallOpts)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_Bindings *BindingsCaller) GetRoleAdmin(opts *bind.CallOpts, role [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _Bindings.contract.Call(opts, &out, "getRoleAdmin", role)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_Bindings *BindingsSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _Bindings.Contract.GetRoleAdmin(&_Bindings.CallOpts, role)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_Bindings *BindingsCallerSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _Bindings.Contract.GetRoleAdmin(&_Bindings.CallOpts, role)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_Bindings *BindingsCaller) HasRole(opts *bind.CallOpts, role [32]byte, account common.Address) (bool, error) {
	var out []interface{}
	err := _Bindings.contract.Call(opts, &out, "hasRole", role, account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_Bindings *BindingsSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _Bindings.Contract.HasRole(&_Bindings.CallOpts, role, account)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_Bindings *BindingsCallerSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _Bindings.Contract.HasRole(&_Bindings.CallOpts, role, account)
}

// LastForgedBatch is a free data retrieval call binding the contract method 0x44e0b2ce.
//
// Solidity: function lastForgedBatch() view returns(uint32)
func (_Bindings *BindingsCaller) LastForgedBatch(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _Bindings.contract.Call(opts, &out, "lastForgedBatch")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// LastForgedBatch is a free data retrieval call binding the contract method 0x44e0b2ce.
//
// Solidity: function lastForgedBatch() view returns(uint32)
func (_Bindings *BindingsSession) LastForgedBatch() (uint32, error) {
	return _Bindings.Contract.LastForgedBatch(&_Bindings.CallOpts)
}

// LastForgedBatch is a free data retrieval call binding the contract method 0x44e0b2ce.
//
// Solidity: function lastForgedBatch() view returns(uint32)
func (_Bindings *BindingsCallerSession) LastForgedBatch() (uint32, error) {
	return _Bindings.Contract.LastForgedBatch(&_Bindings.CallOpts)
}

// MinBalance is a free data retrieval call binding the contract method 0xc5bb8758.
//
// Solidity: function minBalance() view returns(uint256)
func (_Bindings *BindingsCaller) MinBalance(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Bindings.contract.Call(opts, &out, "minBalance")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MinBalance is a free data retrieval call binding the contract method 0xc5bb8758.
//
// Solidity: function minBalance() view returns(uint256)
func (_Bindings *BindingsSession) MinBalance() (*big.Int, error) {
	return _Bindings.Contract.MinBalance(&_Bindings.CallOpts)
}

// MinBalance is a free data retrieval call binding the contract method 0xc5bb8758.
//
// Solidity: function minBalance() view returns(uint256)
func (_Bindings *BindingsCallerSession) MinBalance() (*big.Int, error) {
	return _Bindings.Contract.MinBalance(&_Bindings.CallOpts)
}

// ScoreRootMap is a free data retrieval call binding the contract method 0xbd8a4a61.
//
// Solidity: function scoreRootMap(uint32 ) view returns(uint256)
func (_Bindings *BindingsCaller) ScoreRootMap(opts *bind.CallOpts, arg0 uint32) (*big.Int, error) {
	var out []interface{}
	err := _Bindings.contract.Call(opts, &out, "scoreRootMap", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ScoreRootMap is a free data retrieval call binding the contract method 0xbd8a4a61.
//
// Solidity: function scoreRootMap(uint32 ) view returns(uint256)
func (_Bindings *BindingsSession) ScoreRootMap(arg0 uint32) (*big.Int, error) {
	return _Bindings.Contract.ScoreRootMap(&_Bindings.CallOpts, arg0)
}

// ScoreRootMap is a free data retrieval call binding the contract method 0xbd8a4a61.
//
// Solidity: function scoreRootMap(uint32 ) view returns(uint256)
func (_Bindings *BindingsCallerSession) ScoreRootMap(arg0 uint32) (*big.Int, error) {
	return _Bindings.Contract.ScoreRootMap(&_Bindings.CallOpts, arg0)
}

// ScoreSnapshots is a free data retrieval call binding the contract method 0xdf4c5fe2.
//
// Solidity: function scoreSnapshots(address ) view returns(uint32 score, uint32 batchNum)
func (_Bindings *BindingsCaller) ScoreSnapshots(opts *bind.CallOpts, arg0 common.Address) (struct {
	Score    uint32
	BatchNum uint32
}, error) {
	var out []interface{}
	err := _Bindings.contract.Call(opts, &out, "scoreSnapshots", arg0)

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
func (_Bindings *BindingsSession) ScoreSnapshots(arg0 common.Address) (struct {
	Score    uint32
	BatchNum uint32
}, error) {
	return _Bindings.Contract.ScoreSnapshots(&_Bindings.CallOpts, arg0)
}

// ScoreSnapshots is a free data retrieval call binding the contract method 0xdf4c5fe2.
//
// Solidity: function scoreSnapshots(address ) view returns(uint32 score, uint32 batchNum)
func (_Bindings *BindingsCallerSession) ScoreSnapshots(arg0 common.Address) (struct {
	Score    uint32
	BatchNum uint32
}, error) {
	return _Bindings.Contract.ScoreSnapshots(&_Bindings.CallOpts, arg0)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Bindings *BindingsCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _Bindings.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Bindings *BindingsSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _Bindings.Contract.SupportsInterface(&_Bindings.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Bindings *BindingsCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _Bindings.Contract.SupportsInterface(&_Bindings.CallOpts, interfaceId)
}

// TxsDataHashMap is a free data retrieval call binding the contract method 0xf872ecdd.
//
// Solidity: function txsDataHashMap(uint32 ) view returns(bytes32)
func (_Bindings *BindingsCaller) TxsDataHashMap(opts *bind.CallOpts, arg0 uint32) ([32]byte, error) {
	var out []interface{}
	err := _Bindings.contract.Call(opts, &out, "txsDataHashMap", arg0)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// TxsDataHashMap is a free data retrieval call binding the contract method 0xf872ecdd.
//
// Solidity: function txsDataHashMap(uint32 ) view returns(bytes32)
func (_Bindings *BindingsSession) TxsDataHashMap(arg0 uint32) ([32]byte, error) {
	return _Bindings.Contract.TxsDataHashMap(&_Bindings.CallOpts, arg0)
}

// TxsDataHashMap is a free data retrieval call binding the contract method 0xf872ecdd.
//
// Solidity: function txsDataHashMap(uint32 ) view returns(bytes32)
func (_Bindings *BindingsCallerSession) TxsDataHashMap(arg0 uint32) ([32]byte, error) {
	return _Bindings.Contract.TxsDataHashMap(&_Bindings.CallOpts, arg0)
}

// UnprocessedBatchesMap is a free data retrieval call binding the contract method 0xef8140b5.
//
// Solidity: function unprocessedBatchesMap(uint32 ) view returns(bytes)
func (_Bindings *BindingsCaller) UnprocessedBatchesMap(opts *bind.CallOpts, arg0 uint32) ([]byte, error) {
	var out []interface{}
	err := _Bindings.contract.Call(opts, &out, "unprocessedBatchesMap", arg0)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// UnprocessedBatchesMap is a free data retrieval call binding the contract method 0xef8140b5.
//
// Solidity: function unprocessedBatchesMap(uint32 ) view returns(bytes)
func (_Bindings *BindingsSession) UnprocessedBatchesMap(arg0 uint32) ([]byte, error) {
	return _Bindings.Contract.UnprocessedBatchesMap(&_Bindings.CallOpts, arg0)
}

// UnprocessedBatchesMap is a free data retrieval call binding the contract method 0xef8140b5.
//
// Solidity: function unprocessedBatchesMap(uint32 ) view returns(bytes)
func (_Bindings *BindingsCallerSession) UnprocessedBatchesMap(arg0 uint32) ([]byte, error) {
	return _Bindings.Contract.UnprocessedBatchesMap(&_Bindings.CallOpts, arg0)
}

// Verifier is a free data retrieval call binding the contract method 0x2b7ac3f3.
//
// Solidity: function verifier() view returns(address verifierInterface, uint256 maxTx, uint256 nLevel)
func (_Bindings *BindingsCaller) Verifier(opts *bind.CallOpts) (struct {
	VerifierInterface common.Address
	MaxTx             *big.Int
	NLevel            *big.Int
}, error) {
	var out []interface{}
	err := _Bindings.contract.Call(opts, &out, "verifier")

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
func (_Bindings *BindingsSession) Verifier() (struct {
	VerifierInterface common.Address
	MaxTx             *big.Int
	NLevel            *big.Int
}, error) {
	return _Bindings.Contract.Verifier(&_Bindings.CallOpts)
}

// Verifier is a free data retrieval call binding the contract method 0x2b7ac3f3.
//
// Solidity: function verifier() view returns(address verifierInterface, uint256 maxTx, uint256 nLevel)
func (_Bindings *BindingsCallerSession) Verifier() (struct {
	VerifierInterface common.Address
	MaxTx             *big.Int
	NLevel            *big.Int
}, error) {
	return _Bindings.Contract.Verifier(&_Bindings.CallOpts)
}

// VouchRootMap is a free data retrieval call binding the contract method 0xadacd33b.
//
// Solidity: function vouchRootMap(uint32 ) view returns(uint256)
func (_Bindings *BindingsCaller) VouchRootMap(opts *bind.CallOpts, arg0 uint32) (*big.Int, error) {
	var out []interface{}
	err := _Bindings.contract.Call(opts, &out, "vouchRootMap", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// VouchRootMap is a free data retrieval call binding the contract method 0xadacd33b.
//
// Solidity: function vouchRootMap(uint32 ) view returns(uint256)
func (_Bindings *BindingsSession) VouchRootMap(arg0 uint32) (*big.Int, error) {
	return _Bindings.Contract.VouchRootMap(&_Bindings.CallOpts, arg0)
}

// VouchRootMap is a free data retrieval call binding the contract method 0xadacd33b.
//
// Solidity: function vouchRootMap(uint32 ) view returns(uint256)
func (_Bindings *BindingsCallerSession) VouchRootMap(arg0 uint32) (*big.Int, error) {
	return _Bindings.Contract.VouchRootMap(&_Bindings.CallOpts, arg0)
}

// Vouches is a free data retrieval call binding the contract method 0x0b337be6.
//
// Solidity: function vouches(address , address ) view returns(bool)
func (_Bindings *BindingsCaller) Vouches(opts *bind.CallOpts, arg0 common.Address, arg1 common.Address) (bool, error) {
	var out []interface{}
	err := _Bindings.contract.Call(opts, &out, "vouches", arg0, arg1)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Vouches is a free data retrieval call binding the contract method 0x0b337be6.
//
// Solidity: function vouches(address , address ) view returns(bool)
func (_Bindings *BindingsSession) Vouches(arg0 common.Address, arg1 common.Address) (bool, error) {
	return _Bindings.Contract.Vouches(&_Bindings.CallOpts, arg0, arg1)
}

// Vouches is a free data retrieval call binding the contract method 0x0b337be6.
//
// Solidity: function vouches(address , address ) view returns(bool)
func (_Bindings *BindingsCallerSession) Vouches(arg0 common.Address, arg1 common.Address) (bool, error) {
	return _Bindings.Contract.Vouches(&_Bindings.CallOpts, arg0, arg1)
}

// Deposit is a paid mutator transaction binding the contract method 0xd0e30db0.
//
// Solidity: function deposit() payable returns()
func (_Bindings *BindingsTransactor) Deposit(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Bindings.contract.Transact(opts, "deposit")
}

// Deposit is a paid mutator transaction binding the contract method 0xd0e30db0.
//
// Solidity: function deposit() payable returns()
func (_Bindings *BindingsSession) Deposit() (*types.Transaction, error) {
	return _Bindings.Contract.Deposit(&_Bindings.TransactOpts)
}

// Deposit is a paid mutator transaction binding the contract method 0xd0e30db0.
//
// Solidity: function deposit() payable returns()
func (_Bindings *BindingsTransactorSession) Deposit() (*types.Transaction, error) {
	return _Bindings.Contract.Deposit(&_Bindings.TransactOpts)
}

// ExplodeMultiple is a paid mutator transaction binding the contract method 0xd847d0ca.
//
// Solidity: function explodeMultiple(address[] toEthAddrs) returns()
func (_Bindings *BindingsTransactor) ExplodeMultiple(opts *bind.TransactOpts, toEthAddrs []common.Address) (*types.Transaction, error) {
	return _Bindings.contract.Transact(opts, "explodeMultiple", toEthAddrs)
}

// ExplodeMultiple is a paid mutator transaction binding the contract method 0xd847d0ca.
//
// Solidity: function explodeMultiple(address[] toEthAddrs) returns()
func (_Bindings *BindingsSession) ExplodeMultiple(toEthAddrs []common.Address) (*types.Transaction, error) {
	return _Bindings.Contract.ExplodeMultiple(&_Bindings.TransactOpts, toEthAddrs)
}

// ExplodeMultiple is a paid mutator transaction binding the contract method 0xd847d0ca.
//
// Solidity: function explodeMultiple(address[] toEthAddrs) returns()
func (_Bindings *BindingsTransactorSession) ExplodeMultiple(toEthAddrs []common.Address) (*types.Transaction, error) {
	return _Bindings.Contract.ExplodeMultiple(&_Bindings.TransactOpts, toEthAddrs)
}

// ForgeBatch is a paid mutator transaction binding the contract method 0xcc7e445e.
//
// Solidity: function forgeBatch(uint256 newAccountRoot, uint256 newVouchRoot, uint256 newScoreRoot, uint256[2] proofA, uint256[2][2] proofB, uint256[2] proofC) returns()
func (_Bindings *BindingsTransactor) ForgeBatch(opts *bind.TransactOpts, newAccountRoot *big.Int, newVouchRoot *big.Int, newScoreRoot *big.Int, proofA [2]*big.Int, proofB [2][2]*big.Int, proofC [2]*big.Int) (*types.Transaction, error) {
	return _Bindings.contract.Transact(opts, "forgeBatch", newAccountRoot, newVouchRoot, newScoreRoot, proofA, proofB, proofC)
}

// ForgeBatch is a paid mutator transaction binding the contract method 0xcc7e445e.
//
// Solidity: function forgeBatch(uint256 newAccountRoot, uint256 newVouchRoot, uint256 newScoreRoot, uint256[2] proofA, uint256[2][2] proofB, uint256[2] proofC) returns()
func (_Bindings *BindingsSession) ForgeBatch(newAccountRoot *big.Int, newVouchRoot *big.Int, newScoreRoot *big.Int, proofA [2]*big.Int, proofB [2][2]*big.Int, proofC [2]*big.Int) (*types.Transaction, error) {
	return _Bindings.Contract.ForgeBatch(&_Bindings.TransactOpts, newAccountRoot, newVouchRoot, newScoreRoot, proofA, proofB, proofC)
}

// ForgeBatch is a paid mutator transaction binding the contract method 0xcc7e445e.
//
// Solidity: function forgeBatch(uint256 newAccountRoot, uint256 newVouchRoot, uint256 newScoreRoot, uint256[2] proofA, uint256[2][2] proofB, uint256[2] proofC) returns()
func (_Bindings *BindingsTransactorSession) ForgeBatch(newAccountRoot *big.Int, newVouchRoot *big.Int, newScoreRoot *big.Int, proofA [2]*big.Int, proofB [2][2]*big.Int, proofC [2]*big.Int) (*types.Transaction, error) {
	return _Bindings.Contract.ForgeBatch(&_Bindings.TransactOpts, newAccountRoot, newVouchRoot, newScoreRoot, proofA, proofB, proofC)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_Bindings *BindingsTransactor) GrantRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Bindings.contract.Transact(opts, "grantRole", role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_Bindings *BindingsSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Bindings.Contract.GrantRole(&_Bindings.TransactOpts, role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_Bindings *BindingsTransactorSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Bindings.Contract.GrantRole(&_Bindings.TransactOpts, role, account)
}

// Initialize is a paid mutator transaction binding the contract method 0x11917b1d.
//
// Solidity: function initialize(address _verifier, uint256 maxTx, uint256 nLevel, address _poseidon2Elements, address _poseidon3Elements, address _adminRole) returns()
func (_Bindings *BindingsTransactor) Initialize(opts *bind.TransactOpts, _verifier common.Address, maxTx *big.Int, nLevel *big.Int, _poseidon2Elements common.Address, _poseidon3Elements common.Address, _adminRole common.Address) (*types.Transaction, error) {
	return _Bindings.contract.Transact(opts, "initialize", _verifier, maxTx, nLevel, _poseidon2Elements, _poseidon3Elements, _adminRole)
}

// Initialize is a paid mutator transaction binding the contract method 0x11917b1d.
//
// Solidity: function initialize(address _verifier, uint256 maxTx, uint256 nLevel, address _poseidon2Elements, address _poseidon3Elements, address _adminRole) returns()
func (_Bindings *BindingsSession) Initialize(_verifier common.Address, maxTx *big.Int, nLevel *big.Int, _poseidon2Elements common.Address, _poseidon3Elements common.Address, _adminRole common.Address) (*types.Transaction, error) {
	return _Bindings.Contract.Initialize(&_Bindings.TransactOpts, _verifier, maxTx, nLevel, _poseidon2Elements, _poseidon3Elements, _adminRole)
}

// Initialize is a paid mutator transaction binding the contract method 0x11917b1d.
//
// Solidity: function initialize(address _verifier, uint256 maxTx, uint256 nLevel, address _poseidon2Elements, address _poseidon3Elements, address _adminRole) returns()
func (_Bindings *BindingsTransactorSession) Initialize(_verifier common.Address, maxTx *big.Int, nLevel *big.Int, _poseidon2Elements common.Address, _poseidon3Elements common.Address, _adminRole common.Address) (*types.Transaction, error) {
	return _Bindings.Contract.Initialize(&_Bindings.TransactOpts, _verifier, maxTx, nLevel, _poseidon2Elements, _poseidon3Elements, _adminRole)
}

// ProveScoreMerkleProof is a paid mutator transaction binding the contract method 0x23add439.
//
// Solidity: function proveScoreMerkleProof(uint32 numScoreRoot, uint24 idx, uint32 score, uint256[] siblings) returns()
func (_Bindings *BindingsTransactor) ProveScoreMerkleProof(opts *bind.TransactOpts, numScoreRoot uint32, idx *big.Int, score uint32, siblings []*big.Int) (*types.Transaction, error) {
	return _Bindings.contract.Transact(opts, "proveScoreMerkleProof", numScoreRoot, idx, score, siblings)
}

// ProveScoreMerkleProof is a paid mutator transaction binding the contract method 0x23add439.
//
// Solidity: function proveScoreMerkleProof(uint32 numScoreRoot, uint24 idx, uint32 score, uint256[] siblings) returns()
func (_Bindings *BindingsSession) ProveScoreMerkleProof(numScoreRoot uint32, idx *big.Int, score uint32, siblings []*big.Int) (*types.Transaction, error) {
	return _Bindings.Contract.ProveScoreMerkleProof(&_Bindings.TransactOpts, numScoreRoot, idx, score, siblings)
}

// ProveScoreMerkleProof is a paid mutator transaction binding the contract method 0x23add439.
//
// Solidity: function proveScoreMerkleProof(uint32 numScoreRoot, uint24 idx, uint32 score, uint256[] siblings) returns()
func (_Bindings *BindingsTransactorSession) ProveScoreMerkleProof(numScoreRoot uint32, idx *big.Int, score uint32, siblings []*big.Int) (*types.Transaction, error) {
	return _Bindings.Contract.ProveScoreMerkleProof(&_Bindings.TransactOpts, numScoreRoot, idx, score, siblings)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_Bindings *BindingsTransactor) RenounceRole(opts *bind.TransactOpts, role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _Bindings.contract.Transact(opts, "renounceRole", role, callerConfirmation)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_Bindings *BindingsSession) RenounceRole(role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _Bindings.Contract.RenounceRole(&_Bindings.TransactOpts, role, callerConfirmation)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_Bindings *BindingsTransactorSession) RenounceRole(role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _Bindings.Contract.RenounceRole(&_Bindings.TransactOpts, role, callerConfirmation)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_Bindings *BindingsTransactor) RevokeRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Bindings.contract.Transact(opts, "revokeRole", role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_Bindings *BindingsSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Bindings.Contract.RevokeRole(&_Bindings.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_Bindings *BindingsTransactorSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Bindings.Contract.RevokeRole(&_Bindings.TransactOpts, role, account)
}

// Unvouch is a paid mutator transaction binding the contract method 0xab43cc36.
//
// Solidity: function unvouch(address toEthAddr) returns()
func (_Bindings *BindingsTransactor) Unvouch(opts *bind.TransactOpts, toEthAddr common.Address) (*types.Transaction, error) {
	return _Bindings.contract.Transact(opts, "unvouch", toEthAddr)
}

// Unvouch is a paid mutator transaction binding the contract method 0xab43cc36.
//
// Solidity: function unvouch(address toEthAddr) returns()
func (_Bindings *BindingsSession) Unvouch(toEthAddr common.Address) (*types.Transaction, error) {
	return _Bindings.Contract.Unvouch(&_Bindings.TransactOpts, toEthAddr)
}

// Unvouch is a paid mutator transaction binding the contract method 0xab43cc36.
//
// Solidity: function unvouch(address toEthAddr) returns()
func (_Bindings *BindingsTransactorSession) Unvouch(toEthAddr common.Address) (*types.Transaction, error) {
	return _Bindings.Contract.Unvouch(&_Bindings.TransactOpts, toEthAddr)
}

// UpdateExplodeAmount is a paid mutator transaction binding the contract method 0xaa4f9116.
//
// Solidity: function updateExplodeAmount(uint256 _explodeAmount) returns()
func (_Bindings *BindingsTransactor) UpdateExplodeAmount(opts *bind.TransactOpts, _explodeAmount *big.Int) (*types.Transaction, error) {
	return _Bindings.contract.Transact(opts, "updateExplodeAmount", _explodeAmount)
}

// UpdateExplodeAmount is a paid mutator transaction binding the contract method 0xaa4f9116.
//
// Solidity: function updateExplodeAmount(uint256 _explodeAmount) returns()
func (_Bindings *BindingsSession) UpdateExplodeAmount(_explodeAmount *big.Int) (*types.Transaction, error) {
	return _Bindings.Contract.UpdateExplodeAmount(&_Bindings.TransactOpts, _explodeAmount)
}

// UpdateExplodeAmount is a paid mutator transaction binding the contract method 0xaa4f9116.
//
// Solidity: function updateExplodeAmount(uint256 _explodeAmount) returns()
func (_Bindings *BindingsTransactorSession) UpdateExplodeAmount(_explodeAmount *big.Int) (*types.Transaction, error) {
	return _Bindings.Contract.UpdateExplodeAmount(&_Bindings.TransactOpts, _explodeAmount)
}

// UpdateMinBalance is a paid mutator transaction binding the contract method 0xd83567ab.
//
// Solidity: function updateMinBalance(uint256 _minBalance) returns()
func (_Bindings *BindingsTransactor) UpdateMinBalance(opts *bind.TransactOpts, _minBalance *big.Int) (*types.Transaction, error) {
	return _Bindings.contract.Transact(opts, "updateMinBalance", _minBalance)
}

// UpdateMinBalance is a paid mutator transaction binding the contract method 0xd83567ab.
//
// Solidity: function updateMinBalance(uint256 _minBalance) returns()
func (_Bindings *BindingsSession) UpdateMinBalance(_minBalance *big.Int) (*types.Transaction, error) {
	return _Bindings.Contract.UpdateMinBalance(&_Bindings.TransactOpts, _minBalance)
}

// UpdateMinBalance is a paid mutator transaction binding the contract method 0xd83567ab.
//
// Solidity: function updateMinBalance(uint256 _minBalance) returns()
func (_Bindings *BindingsTransactorSession) UpdateMinBalance(_minBalance *big.Int) (*types.Transaction, error) {
	return _Bindings.Contract.UpdateMinBalance(&_Bindings.TransactOpts, _minBalance)
}

// Vouch is a paid mutator transaction binding the contract method 0xdd66e16b.
//
// Solidity: function vouch(address toEthAddr) returns()
func (_Bindings *BindingsTransactor) Vouch(opts *bind.TransactOpts, toEthAddr common.Address) (*types.Transaction, error) {
	return _Bindings.contract.Transact(opts, "vouch", toEthAddr)
}

// Vouch is a paid mutator transaction binding the contract method 0xdd66e16b.
//
// Solidity: function vouch(address toEthAddr) returns()
func (_Bindings *BindingsSession) Vouch(toEthAddr common.Address) (*types.Transaction, error) {
	return _Bindings.Contract.Vouch(&_Bindings.TransactOpts, toEthAddr)
}

// Vouch is a paid mutator transaction binding the contract method 0xdd66e16b.
//
// Solidity: function vouch(address toEthAddr) returns()
func (_Bindings *BindingsTransactorSession) Vouch(toEthAddr common.Address) (*types.Transaction, error) {
	return _Bindings.Contract.Vouch(&_Bindings.TransactOpts, toEthAddr)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 amount) returns()
func (_Bindings *BindingsTransactor) Withdraw(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _Bindings.contract.Transact(opts, "withdraw", amount)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 amount) returns()
func (_Bindings *BindingsSession) Withdraw(amount *big.Int) (*types.Transaction, error) {
	return _Bindings.Contract.Withdraw(&_Bindings.TransactOpts, amount)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 amount) returns()
func (_Bindings *BindingsTransactorSession) Withdraw(amount *big.Int) (*types.Transaction, error) {
	return _Bindings.Contract.Withdraw(&_Bindings.TransactOpts, amount)
}

// BindingsExplodeAmountUpdatedIterator is returned from FilterExplodeAmountUpdated and is used to iterate over the raw logs and unpacked data for ExplodeAmountUpdated events raised by the Bindings contract.
type BindingsExplodeAmountUpdatedIterator struct {
	Event *BindingsExplodeAmountUpdated // Event containing the contract specifics and raw log

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
func (it *BindingsExplodeAmountUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BindingsExplodeAmountUpdated)
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
		it.Event = new(BindingsExplodeAmountUpdated)
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
func (it *BindingsExplodeAmountUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BindingsExplodeAmountUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BindingsExplodeAmountUpdated represents a ExplodeAmountUpdated event raised by the Bindings contract.
type BindingsExplodeAmountUpdated struct {
	ExplodeAmount *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterExplodeAmountUpdated is a free log retrieval operation binding the contract event 0xe4d07ddba7bee2524330d02dcb17edc18e025341f913484d72c85301628b4a79.
//
// Solidity: event ExplodeAmountUpdated(uint256 explodeAmount)
func (_Bindings *BindingsFilterer) FilterExplodeAmountUpdated(opts *bind.FilterOpts) (*BindingsExplodeAmountUpdatedIterator, error) {

	logs, sub, err := _Bindings.contract.FilterLogs(opts, "ExplodeAmountUpdated")
	if err != nil {
		return nil, err
	}
	return &BindingsExplodeAmountUpdatedIterator{contract: _Bindings.contract, event: "ExplodeAmountUpdated", logs: logs, sub: sub}, nil
}

// WatchExplodeAmountUpdated is a free log subscription operation binding the contract event 0xe4d07ddba7bee2524330d02dcb17edc18e025341f913484d72c85301628b4a79.
//
// Solidity: event ExplodeAmountUpdated(uint256 explodeAmount)
func (_Bindings *BindingsFilterer) WatchExplodeAmountUpdated(opts *bind.WatchOpts, sink chan<- *BindingsExplodeAmountUpdated) (event.Subscription, error) {

	logs, sub, err := _Bindings.contract.WatchLogs(opts, "ExplodeAmountUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BindingsExplodeAmountUpdated)
				if err := _Bindings.contract.UnpackLog(event, "ExplodeAmountUpdated", log); err != nil {
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
func (_Bindings *BindingsFilterer) ParseExplodeAmountUpdated(log types.Log) (*BindingsExplodeAmountUpdated, error) {
	event := new(BindingsExplodeAmountUpdated)
	if err := _Bindings.contract.UnpackLog(event, "ExplodeAmountUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BindingsForgeBatchIterator is returned from FilterForgeBatch and is used to iterate over the raw logs and unpacked data for ForgeBatch events raised by the Bindings contract.
type BindingsForgeBatchIterator struct {
	Event *BindingsForgeBatch // Event containing the contract specifics and raw log

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
func (it *BindingsForgeBatchIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BindingsForgeBatch)
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
		it.Event = new(BindingsForgeBatch)
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
func (it *BindingsForgeBatchIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BindingsForgeBatchIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BindingsForgeBatch represents a ForgeBatch event raised by the Bindings contract.
type BindingsForgeBatch struct {
	BatchNum     uint32
	L1UserTxsLen uint16
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterForgeBatch is a free log retrieval operation binding the contract event 0xe00040c8a3b0bf905636c26924e90520eafc5003324138236fddee2d34588618.
//
// Solidity: event ForgeBatch(uint32 indexed batchNum, uint16 l1UserTxsLen)
func (_Bindings *BindingsFilterer) FilterForgeBatch(opts *bind.FilterOpts, batchNum []uint32) (*BindingsForgeBatchIterator, error) {

	var batchNumRule []interface{}
	for _, batchNumItem := range batchNum {
		batchNumRule = append(batchNumRule, batchNumItem)
	}

	logs, sub, err := _Bindings.contract.FilterLogs(opts, "ForgeBatch", batchNumRule)
	if err != nil {
		return nil, err
	}
	return &BindingsForgeBatchIterator{contract: _Bindings.contract, event: "ForgeBatch", logs: logs, sub: sub}, nil
}

// WatchForgeBatch is a free log subscription operation binding the contract event 0xe00040c8a3b0bf905636c26924e90520eafc5003324138236fddee2d34588618.
//
// Solidity: event ForgeBatch(uint32 indexed batchNum, uint16 l1UserTxsLen)
func (_Bindings *BindingsFilterer) WatchForgeBatch(opts *bind.WatchOpts, sink chan<- *BindingsForgeBatch, batchNum []uint32) (event.Subscription, error) {

	var batchNumRule []interface{}
	for _, batchNumItem := range batchNum {
		batchNumRule = append(batchNumRule, batchNumItem)
	}

	logs, sub, err := _Bindings.contract.WatchLogs(opts, "ForgeBatch", batchNumRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BindingsForgeBatch)
				if err := _Bindings.contract.UnpackLog(event, "ForgeBatch", log); err != nil {
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

// ParseForgeBatch is a log parse operation binding the contract event 0xe00040c8a3b0bf905636c26924e90520eafc5003324138236fddee2d34588618.
//
// Solidity: event ForgeBatch(uint32 indexed batchNum, uint16 l1UserTxsLen)
func (_Bindings *BindingsFilterer) ParseForgeBatch(log types.Log) (*BindingsForgeBatch, error) {
	event := new(BindingsForgeBatch)
	if err := _Bindings.contract.UnpackLog(event, "ForgeBatch", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BindingsInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the Bindings contract.
type BindingsInitializedIterator struct {
	Event *BindingsInitialized // Event containing the contract specifics and raw log

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
func (it *BindingsInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BindingsInitialized)
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
		it.Event = new(BindingsInitialized)
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
func (it *BindingsInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BindingsInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BindingsInitialized represents a Initialized event raised by the Bindings contract.
type BindingsInitialized struct {
	Version uint64
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_Bindings *BindingsFilterer) FilterInitialized(opts *bind.FilterOpts) (*BindingsInitializedIterator, error) {

	logs, sub, err := _Bindings.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &BindingsInitializedIterator{contract: _Bindings.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_Bindings *BindingsFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *BindingsInitialized) (event.Subscription, error) {

	logs, sub, err := _Bindings.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BindingsInitialized)
				if err := _Bindings.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_Bindings *BindingsFilterer) ParseInitialized(log types.Log) (*BindingsInitialized, error) {
	event := new(BindingsInitialized)
	if err := _Bindings.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BindingsL1UserTxEventIterator is returned from FilterL1UserTxEvent and is used to iterate over the raw logs and unpacked data for L1UserTxEvent events raised by the Bindings contract.
type BindingsL1UserTxEventIterator struct {
	Event *BindingsL1UserTxEvent // Event containing the contract specifics and raw log

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
func (it *BindingsL1UserTxEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BindingsL1UserTxEvent)
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
		it.Event = new(BindingsL1UserTxEvent)
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
func (it *BindingsL1UserTxEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BindingsL1UserTxEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BindingsL1UserTxEvent represents a L1UserTxEvent event raised by the Bindings contract.
type BindingsL1UserTxEvent struct {
	QueueIndex uint32
	Position   uint8
	L1UserTx   []byte
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterL1UserTxEvent is a free log retrieval operation binding the contract event 0xdd5c7c5ea02d3c5d1621513faa6de53d474ee6f111eda6352a63e3dfe8c40119.
//
// Solidity: event L1UserTxEvent(uint32 indexed queueIndex, uint8 indexed position, bytes l1UserTx)
func (_Bindings *BindingsFilterer) FilterL1UserTxEvent(opts *bind.FilterOpts, queueIndex []uint32, position []uint8) (*BindingsL1UserTxEventIterator, error) {

	var queueIndexRule []interface{}
	for _, queueIndexItem := range queueIndex {
		queueIndexRule = append(queueIndexRule, queueIndexItem)
	}
	var positionRule []interface{}
	for _, positionItem := range position {
		positionRule = append(positionRule, positionItem)
	}

	logs, sub, err := _Bindings.contract.FilterLogs(opts, "L1UserTxEvent", queueIndexRule, positionRule)
	if err != nil {
		return nil, err
	}
	return &BindingsL1UserTxEventIterator{contract: _Bindings.contract, event: "L1UserTxEvent", logs: logs, sub: sub}, nil
}

// WatchL1UserTxEvent is a free log subscription operation binding the contract event 0xdd5c7c5ea02d3c5d1621513faa6de53d474ee6f111eda6352a63e3dfe8c40119.
//
// Solidity: event L1UserTxEvent(uint32 indexed queueIndex, uint8 indexed position, bytes l1UserTx)
func (_Bindings *BindingsFilterer) WatchL1UserTxEvent(opts *bind.WatchOpts, sink chan<- *BindingsL1UserTxEvent, queueIndex []uint32, position []uint8) (event.Subscription, error) {

	var queueIndexRule []interface{}
	for _, queueIndexItem := range queueIndex {
		queueIndexRule = append(queueIndexRule, queueIndexItem)
	}
	var positionRule []interface{}
	for _, positionItem := range position {
		positionRule = append(positionRule, positionItem)
	}

	logs, sub, err := _Bindings.contract.WatchLogs(opts, "L1UserTxEvent", queueIndexRule, positionRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BindingsL1UserTxEvent)
				if err := _Bindings.contract.UnpackLog(event, "L1UserTxEvent", log); err != nil {
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

// ParseL1UserTxEvent is a log parse operation binding the contract event 0xdd5c7c5ea02d3c5d1621513faa6de53d474ee6f111eda6352a63e3dfe8c40119.
//
// Solidity: event L1UserTxEvent(uint32 indexed queueIndex, uint8 indexed position, bytes l1UserTx)
func (_Bindings *BindingsFilterer) ParseL1UserTxEvent(log types.Log) (*BindingsL1UserTxEvent, error) {
	event := new(BindingsL1UserTxEvent)
	if err := _Bindings.contract.UnpackLog(event, "L1UserTxEvent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BindingsMinBalanceUpdatedIterator is returned from FilterMinBalanceUpdated and is used to iterate over the raw logs and unpacked data for MinBalanceUpdated events raised by the Bindings contract.
type BindingsMinBalanceUpdatedIterator struct {
	Event *BindingsMinBalanceUpdated // Event containing the contract specifics and raw log

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
func (it *BindingsMinBalanceUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BindingsMinBalanceUpdated)
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
		it.Event = new(BindingsMinBalanceUpdated)
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
func (it *BindingsMinBalanceUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BindingsMinBalanceUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BindingsMinBalanceUpdated represents a MinBalanceUpdated event raised by the Bindings contract.
type BindingsMinBalanceUpdated struct {
	MinBalance *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterMinBalanceUpdated is a free log retrieval operation binding the contract event 0x4e1cd0a17dbc393262d4d9b66380671f5273c5f0a34fed0ed36c50ba6b1f0e16.
//
// Solidity: event MinBalanceUpdated(uint256 minBalance)
func (_Bindings *BindingsFilterer) FilterMinBalanceUpdated(opts *bind.FilterOpts) (*BindingsMinBalanceUpdatedIterator, error) {

	logs, sub, err := _Bindings.contract.FilterLogs(opts, "MinBalanceUpdated")
	if err != nil {
		return nil, err
	}
	return &BindingsMinBalanceUpdatedIterator{contract: _Bindings.contract, event: "MinBalanceUpdated", logs: logs, sub: sub}, nil
}

// WatchMinBalanceUpdated is a free log subscription operation binding the contract event 0x4e1cd0a17dbc393262d4d9b66380671f5273c5f0a34fed0ed36c50ba6b1f0e16.
//
// Solidity: event MinBalanceUpdated(uint256 minBalance)
func (_Bindings *BindingsFilterer) WatchMinBalanceUpdated(opts *bind.WatchOpts, sink chan<- *BindingsMinBalanceUpdated) (event.Subscription, error) {

	logs, sub, err := _Bindings.contract.WatchLogs(opts, "MinBalanceUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BindingsMinBalanceUpdated)
				if err := _Bindings.contract.UnpackLog(event, "MinBalanceUpdated", log); err != nil {
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

// ParseMinBalanceUpdated is a log parse operation binding the contract event 0x4e1cd0a17dbc393262d4d9b66380671f5273c5f0a34fed0ed36c50ba6b1f0e16.
//
// Solidity: event MinBalanceUpdated(uint256 minBalance)
func (_Bindings *BindingsFilterer) ParseMinBalanceUpdated(log types.Log) (*BindingsMinBalanceUpdated, error) {
	event := new(BindingsMinBalanceUpdated)
	if err := _Bindings.contract.UnpackLog(event, "MinBalanceUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BindingsRoleAdminChangedIterator is returned from FilterRoleAdminChanged and is used to iterate over the raw logs and unpacked data for RoleAdminChanged events raised by the Bindings contract.
type BindingsRoleAdminChangedIterator struct {
	Event *BindingsRoleAdminChanged // Event containing the contract specifics and raw log

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
func (it *BindingsRoleAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BindingsRoleAdminChanged)
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
		it.Event = new(BindingsRoleAdminChanged)
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
func (it *BindingsRoleAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BindingsRoleAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BindingsRoleAdminChanged represents a RoleAdminChanged event raised by the Bindings contract.
type BindingsRoleAdminChanged struct {
	Role              [32]byte
	PreviousAdminRole [32]byte
	NewAdminRole      [32]byte
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterRoleAdminChanged is a free log retrieval operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_Bindings *BindingsFilterer) FilterRoleAdminChanged(opts *bind.FilterOpts, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (*BindingsRoleAdminChangedIterator, error) {

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

	logs, sub, err := _Bindings.contract.FilterLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return &BindingsRoleAdminChangedIterator{contract: _Bindings.contract, event: "RoleAdminChanged", logs: logs, sub: sub}, nil
}

// WatchRoleAdminChanged is a free log subscription operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_Bindings *BindingsFilterer) WatchRoleAdminChanged(opts *bind.WatchOpts, sink chan<- *BindingsRoleAdminChanged, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (event.Subscription, error) {

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

	logs, sub, err := _Bindings.contract.WatchLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BindingsRoleAdminChanged)
				if err := _Bindings.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
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
func (_Bindings *BindingsFilterer) ParseRoleAdminChanged(log types.Log) (*BindingsRoleAdminChanged, error) {
	event := new(BindingsRoleAdminChanged)
	if err := _Bindings.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BindingsRoleGrantedIterator is returned from FilterRoleGranted and is used to iterate over the raw logs and unpacked data for RoleGranted events raised by the Bindings contract.
type BindingsRoleGrantedIterator struct {
	Event *BindingsRoleGranted // Event containing the contract specifics and raw log

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
func (it *BindingsRoleGrantedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BindingsRoleGranted)
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
		it.Event = new(BindingsRoleGranted)
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
func (it *BindingsRoleGrantedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BindingsRoleGrantedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BindingsRoleGranted represents a RoleGranted event raised by the Bindings contract.
type BindingsRoleGranted struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleGranted is a free log retrieval operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_Bindings *BindingsFilterer) FilterRoleGranted(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*BindingsRoleGrantedIterator, error) {

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

	logs, sub, err := _Bindings.contract.FilterLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &BindingsRoleGrantedIterator{contract: _Bindings.contract, event: "RoleGranted", logs: logs, sub: sub}, nil
}

// WatchRoleGranted is a free log subscription operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_Bindings *BindingsFilterer) WatchRoleGranted(opts *bind.WatchOpts, sink chan<- *BindingsRoleGranted, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _Bindings.contract.WatchLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BindingsRoleGranted)
				if err := _Bindings.contract.UnpackLog(event, "RoleGranted", log); err != nil {
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
func (_Bindings *BindingsFilterer) ParseRoleGranted(log types.Log) (*BindingsRoleGranted, error) {
	event := new(BindingsRoleGranted)
	if err := _Bindings.contract.UnpackLog(event, "RoleGranted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BindingsRoleRevokedIterator is returned from FilterRoleRevoked and is used to iterate over the raw logs and unpacked data for RoleRevoked events raised by the Bindings contract.
type BindingsRoleRevokedIterator struct {
	Event *BindingsRoleRevoked // Event containing the contract specifics and raw log

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
func (it *BindingsRoleRevokedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BindingsRoleRevoked)
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
		it.Event = new(BindingsRoleRevoked)
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
func (it *BindingsRoleRevokedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BindingsRoleRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BindingsRoleRevoked represents a RoleRevoked event raised by the Bindings contract.
type BindingsRoleRevoked struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleRevoked is a free log retrieval operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_Bindings *BindingsFilterer) FilterRoleRevoked(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*BindingsRoleRevokedIterator, error) {

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

	logs, sub, err := _Bindings.contract.FilterLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &BindingsRoleRevokedIterator{contract: _Bindings.contract, event: "RoleRevoked", logs: logs, sub: sub}, nil
}

// WatchRoleRevoked is a free log subscription operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_Bindings *BindingsFilterer) WatchRoleRevoked(opts *bind.WatchOpts, sink chan<- *BindingsRoleRevoked, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _Bindings.contract.WatchLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BindingsRoleRevoked)
				if err := _Bindings.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
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
func (_Bindings *BindingsFilterer) ParseRoleRevoked(log types.Log) (*BindingsRoleRevoked, error) {
	event := new(BindingsRoleRevoked)
	if err := _Bindings.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BindingsWithdrawEventIterator is returned from FilterWithdrawEvent and is used to iterate over the raw logs and unpacked data for WithdrawEvent events raised by the Bindings contract.
type BindingsWithdrawEventIterator struct {
	Event *BindingsWithdrawEvent // Event containing the contract specifics and raw log

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
func (it *BindingsWithdrawEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BindingsWithdrawEvent)
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
		it.Event = new(BindingsWithdrawEvent)
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
func (it *BindingsWithdrawEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BindingsWithdrawEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BindingsWithdrawEvent represents a WithdrawEvent event raised by the Bindings contract.
type BindingsWithdrawEvent struct {
	Idx         *big.Int
	NumExitRoot uint32
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterWithdrawEvent is a free log retrieval operation binding the contract event 0x102db758451b2f65238246a452d00c0c4c8f59d8c623aff254111079418e57ec.
//
// Solidity: event WithdrawEvent(uint48 indexed idx, uint32 indexed numExitRoot)
func (_Bindings *BindingsFilterer) FilterWithdrawEvent(opts *bind.FilterOpts, idx []*big.Int, numExitRoot []uint32) (*BindingsWithdrawEventIterator, error) {

	var idxRule []interface{}
	for _, idxItem := range idx {
		idxRule = append(idxRule, idxItem)
	}
	var numExitRootRule []interface{}
	for _, numExitRootItem := range numExitRoot {
		numExitRootRule = append(numExitRootRule, numExitRootItem)
	}

	logs, sub, err := _Bindings.contract.FilterLogs(opts, "WithdrawEvent", idxRule, numExitRootRule)
	if err != nil {
		return nil, err
	}
	return &BindingsWithdrawEventIterator{contract: _Bindings.contract, event: "WithdrawEvent", logs: logs, sub: sub}, nil
}

// WatchWithdrawEvent is a free log subscription operation binding the contract event 0x102db758451b2f65238246a452d00c0c4c8f59d8c623aff254111079418e57ec.
//
// Solidity: event WithdrawEvent(uint48 indexed idx, uint32 indexed numExitRoot)
func (_Bindings *BindingsFilterer) WatchWithdrawEvent(opts *bind.WatchOpts, sink chan<- *BindingsWithdrawEvent, idx []*big.Int, numExitRoot []uint32) (event.Subscription, error) {

	var idxRule []interface{}
	for _, idxItem := range idx {
		idxRule = append(idxRule, idxItem)
	}
	var numExitRootRule []interface{}
	for _, numExitRootItem := range numExitRoot {
		numExitRootRule = append(numExitRootRule, numExitRootItem)
	}

	logs, sub, err := _Bindings.contract.WatchLogs(opts, "WithdrawEvent", idxRule, numExitRootRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BindingsWithdrawEvent)
				if err := _Bindings.contract.UnpackLog(event, "WithdrawEvent", log); err != nil {
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

// ParseWithdrawEvent is a log parse operation binding the contract event 0x102db758451b2f65238246a452d00c0c4c8f59d8c623aff254111079418e57ec.
//
// Solidity: event WithdrawEvent(uint48 indexed idx, uint32 indexed numExitRoot)
func (_Bindings *BindingsFilterer) ParseWithdrawEvent(log types.Log) (*BindingsWithdrawEvent, error) {
	event := new(BindingsWithdrawEvent)
	if err := _Bindings.contract.UnpackLog(event, "WithdrawEvent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
