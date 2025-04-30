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
	ABI: "[{\"inputs\":[],\"name\":\"AccessControlBadConfirmation\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"neededRole\",\"type\":\"bytes32\"}],\"name\":\"AccessControlUnauthorizedAccount\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"AmountExceedsLimit\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"EthTransferFailed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidFromIdx\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidInitialization\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"elementType\",\"type\":\"string\"}],\"name\":\"InvalidPoseidonAddress\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidProof\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidToIdx\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidVerifierAddress\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"LoadAmountDoesNotMatch\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"LoadAmountExceedsLimit\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotInitializing\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"SmtProofInvalid\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"WithdrawAlreadyDone\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"explodeAmount\",\"type\":\"uint256\"}],\"name\":\"ExplodeAmountUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint32\",\"name\":\"batchNum\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"uint16\",\"name\":\"l1UserTxsLen\",\"type\":\"uint16\"}],\"name\":\"ForgeBatch\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"version\",\"type\":\"uint64\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint32\",\"name\":\"queueIndex\",\"type\":\"uint32\"},{\"indexed\":true,\"internalType\":\"uint8\",\"name\":\"position\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"l1UserTx\",\"type\":\"bytes\"}],\"name\":\"L1UserTxEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"minBalance\",\"type\":\"uint256\"}],\"name\":\"MinBalanceUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"previousAdminRole\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"newAdminRole\",\"type\":\"bytes32\"}],\"name\":\"RoleAdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleGranted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleRevoked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint48\",\"name\":\"idx\",\"type\":\"uint48\"},{\"indexed\":true,\"internalType\":\"uint32\",\"name\":\"numExitRoot\",\"type\":\"uint32\"}],\"name\":\"WithdrawEvent\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"DEFAULT_ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"ethAddress\",\"type\":\"address\"},{\"internalType\":\"uint48\",\"name\":\"fromIdx\",\"type\":\"uint48\"},{\"internalType\":\"uint40\",\"name\":\"loadAmountF\",\"type\":\"uint40\"},{\"internalType\":\"uint40\",\"name\":\"amountF\",\"type\":\"uint40\"},{\"internalType\":\"uint48\",\"name\":\"toIdx\",\"type\":\"uint48\"}],\"name\":\"_addTx\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"key\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"_hashFinalNode\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"left\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"right\",\"type\":\"uint256\"}],\"name\":\"_hashNode\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"name\":\"accountRootMap\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint40\",\"name\":\"loadAmountF\",\"type\":\"uint40\"}],\"name\":\"createAccountDeposit\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"currentFillingBatch\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint48\",\"name\":\"fromIdx\",\"type\":\"uint48\"},{\"internalType\":\"uint40\",\"name\":\"loadAmountF\",\"type\":\"uint40\"}],\"name\":\"deposit\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint48\",\"name\":\"fromIdx\",\"type\":\"uint48\"},{\"internalType\":\"uint40\",\"name\":\"amountF\",\"type\":\"uint40\"}],\"name\":\"exit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"},{\"internalType\":\"uint48\",\"name\":\"\",\"type\":\"uint48\"}],\"name\":\"exitNullifierMap\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"name\":\"exitRootMap\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"explodeAmount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint48\",\"name\":\"fromIdx\",\"type\":\"uint48\"},{\"internalType\":\"uint48[]\",\"name\":\"toIdxs\",\"type\":\"uint48[]\"}],\"name\":\"explodeMultiple\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint48\",\"name\":\"newLastIdx\",\"type\":\"uint48\"},{\"internalType\":\"uint256\",\"name\":\"newAccountRoot\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"newVouchRoot\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"newScoreRoot\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"newExitRoot\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"txsData\",\"type\":\"bytes\"},{\"internalType\":\"uint256[2]\",\"name\":\"proofA\",\"type\":\"uint256[2]\"},{\"internalType\":\"uint256[2][2]\",\"name\":\"proofB\",\"type\":\"uint256[2][2]\"},{\"internalType\":\"uint256[2]\",\"name\":\"proofC\",\"type\":\"uint256[2]\"}],\"name\":\"forgeBatch\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"queueIndex\",\"type\":\"uint32\"}],\"name\":\"getL1TransactionQueue\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getLastForgedBatch\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getQueueLength\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleAdmin\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"batchNum\",\"type\":\"uint32\"}],\"name\":\"getStateRoot\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"grantRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"hasRole\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_verifier\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"maxTx\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nLevel\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_poseidon2Elements\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_poseidon3Elements\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_adminRole\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lastForgedBatch\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lastIdx\",\"outputs\":[{\"internalType\":\"uint48\",\"name\":\"\",\"type\":\"uint48\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"minBalance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"callerConfirmation\",\"type\":\"address\"}],\"name\":\"renounceRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"revokeRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"name\":\"scoreRootMap\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"name\":\"txsDataHashMap\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"name\":\"unprocessedBatchesMap\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint48\",\"name\":\"fromIdx\",\"type\":\"uint48\"},{\"internalType\":\"uint48\",\"name\":\"toIdx\",\"type\":\"uint48\"}],\"name\":\"unvouch\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_explodeAmount\",\"type\":\"uint256\"}],\"name\":\"updateExplodeAmount\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_minBalance\",\"type\":\"uint256\"}],\"name\":\"updateMinBalance\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"verifier\",\"outputs\":[{\"internalType\":\"contractIVerifier\",\"name\":\"verifierInterface\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"maxTx\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nLevel\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint48\",\"name\":\"fromIdx\",\"type\":\"uint48\"},{\"internalType\":\"uint48\",\"name\":\"toIdx\",\"type\":\"uint48\"}],\"name\":\"vouch\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"name\":\"vouchRootMap\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint192\",\"name\":\"amount\",\"type\":\"uint192\"},{\"internalType\":\"uint32\",\"name\":\"numExitRoot\",\"type\":\"uint32\"},{\"internalType\":\"uint256[]\",\"name\":\"siblings\",\"type\":\"uint256[]\"},{\"internalType\":\"uint48\",\"name\":\"idx\",\"type\":\"uint48\"}],\"name\":\"withdrawMerkleProof\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
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

// GetL1TransactionQueue is a free data retrieval call binding the contract method 0xba2506df.
//
// Solidity: function getL1TransactionQueue(uint32 queueIndex) view returns(bytes)
func (_Bindings *BindingsCaller) GetL1TransactionQueue(opts *bind.CallOpts, queueIndex uint32) ([]byte, error) {
	var out []interface{}
	err := _Bindings.contract.Call(opts, &out, "getL1TransactionQueue", queueIndex)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// GetL1TransactionQueue is a free data retrieval call binding the contract method 0xba2506df.
//
// Solidity: function getL1TransactionQueue(uint32 queueIndex) view returns(bytes)
func (_Bindings *BindingsSession) GetL1TransactionQueue(queueIndex uint32) ([]byte, error) {
	return _Bindings.Contract.GetL1TransactionQueue(&_Bindings.CallOpts, queueIndex)
}

// GetL1TransactionQueue is a free data retrieval call binding the contract method 0xba2506df.
//
// Solidity: function getL1TransactionQueue(uint32 queueIndex) view returns(bytes)
func (_Bindings *BindingsCallerSession) GetL1TransactionQueue(queueIndex uint32) ([]byte, error) {
	return _Bindings.Contract.GetL1TransactionQueue(&_Bindings.CallOpts, queueIndex)
}

// GetLastForgedBatch is a free data retrieval call binding the contract method 0x1b78164b.
//
// Solidity: function getLastForgedBatch() view returns(uint32)
func (_Bindings *BindingsCaller) GetLastForgedBatch(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _Bindings.contract.Call(opts, &out, "getLastForgedBatch")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// GetLastForgedBatch is a free data retrieval call binding the contract method 0x1b78164b.
//
// Solidity: function getLastForgedBatch() view returns(uint32)
func (_Bindings *BindingsSession) GetLastForgedBatch() (uint32, error) {
	return _Bindings.Contract.GetLastForgedBatch(&_Bindings.CallOpts)
}

// GetLastForgedBatch is a free data retrieval call binding the contract method 0x1b78164b.
//
// Solidity: function getLastForgedBatch() view returns(uint32)
func (_Bindings *BindingsCallerSession) GetLastForgedBatch() (uint32, error) {
	return _Bindings.Contract.GetLastForgedBatch(&_Bindings.CallOpts)
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

// GetStateRoot is a free data retrieval call binding the contract method 0x3009c59f.
//
// Solidity: function getStateRoot(uint32 batchNum) view returns(uint256)
func (_Bindings *BindingsCaller) GetStateRoot(opts *bind.CallOpts, batchNum uint32) (*big.Int, error) {
	var out []interface{}
	err := _Bindings.contract.Call(opts, &out, "getStateRoot", batchNum)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetStateRoot is a free data retrieval call binding the contract method 0x3009c59f.
//
// Solidity: function getStateRoot(uint32 batchNum) view returns(uint256)
func (_Bindings *BindingsSession) GetStateRoot(batchNum uint32) (*big.Int, error) {
	return _Bindings.Contract.GetStateRoot(&_Bindings.CallOpts, batchNum)
}

// GetStateRoot is a free data retrieval call binding the contract method 0x3009c59f.
//
// Solidity: function getStateRoot(uint32 batchNum) view returns(uint256)
func (_Bindings *BindingsCallerSession) GetStateRoot(batchNum uint32) (*big.Int, error) {
	return _Bindings.Contract.GetStateRoot(&_Bindings.CallOpts, batchNum)
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

// LastIdx is a free data retrieval call binding the contract method 0xd486645c.
//
// Solidity: function lastIdx() view returns(uint48)
func (_Bindings *BindingsCaller) LastIdx(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Bindings.contract.Call(opts, &out, "lastIdx")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LastIdx is a free data retrieval call binding the contract method 0xd486645c.
//
// Solidity: function lastIdx() view returns(uint48)
func (_Bindings *BindingsSession) LastIdx() (*big.Int, error) {
	return _Bindings.Contract.LastIdx(&_Bindings.CallOpts)
}

// LastIdx is a free data retrieval call binding the contract method 0xd486645c.
//
// Solidity: function lastIdx() view returns(uint48)
func (_Bindings *BindingsCallerSession) LastIdx() (*big.Int, error) {
	return _Bindings.Contract.LastIdx(&_Bindings.CallOpts)
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

// AddTx is a paid mutator transaction binding the contract method 0x2f463f59.
//
// Solidity: function _addTx(address ethAddress, uint48 fromIdx, uint40 loadAmountF, uint40 amountF, uint48 toIdx) returns()
func (_Bindings *BindingsTransactor) AddTx(opts *bind.TransactOpts, ethAddress common.Address, fromIdx *big.Int, loadAmountF *big.Int, amountF *big.Int, toIdx *big.Int) (*types.Transaction, error) {
	return _Bindings.contract.Transact(opts, "_addTx", ethAddress, fromIdx, loadAmountF, amountF, toIdx)
}

// AddTx is a paid mutator transaction binding the contract method 0x2f463f59.
//
// Solidity: function _addTx(address ethAddress, uint48 fromIdx, uint40 loadAmountF, uint40 amountF, uint48 toIdx) returns()
func (_Bindings *BindingsSession) AddTx(ethAddress common.Address, fromIdx *big.Int, loadAmountF *big.Int, amountF *big.Int, toIdx *big.Int) (*types.Transaction, error) {
	return _Bindings.Contract.AddTx(&_Bindings.TransactOpts, ethAddress, fromIdx, loadAmountF, amountF, toIdx)
}

// AddTx is a paid mutator transaction binding the contract method 0x2f463f59.
//
// Solidity: function _addTx(address ethAddress, uint48 fromIdx, uint40 loadAmountF, uint40 amountF, uint48 toIdx) returns()
func (_Bindings *BindingsTransactorSession) AddTx(ethAddress common.Address, fromIdx *big.Int, loadAmountF *big.Int, amountF *big.Int, toIdx *big.Int) (*types.Transaction, error) {
	return _Bindings.Contract.AddTx(&_Bindings.TransactOpts, ethAddress, fromIdx, loadAmountF, amountF, toIdx)
}

// CreateAccountDeposit is a paid mutator transaction binding the contract method 0xfbb4a00f.
//
// Solidity: function createAccountDeposit(uint40 loadAmountF) payable returns()
func (_Bindings *BindingsTransactor) CreateAccountDeposit(opts *bind.TransactOpts, loadAmountF *big.Int) (*types.Transaction, error) {
	return _Bindings.contract.Transact(opts, "createAccountDeposit", loadAmountF)
}

// CreateAccountDeposit is a paid mutator transaction binding the contract method 0xfbb4a00f.
//
// Solidity: function createAccountDeposit(uint40 loadAmountF) payable returns()
func (_Bindings *BindingsSession) CreateAccountDeposit(loadAmountF *big.Int) (*types.Transaction, error) {
	return _Bindings.Contract.CreateAccountDeposit(&_Bindings.TransactOpts, loadAmountF)
}

// CreateAccountDeposit is a paid mutator transaction binding the contract method 0xfbb4a00f.
//
// Solidity: function createAccountDeposit(uint40 loadAmountF) payable returns()
func (_Bindings *BindingsTransactorSession) CreateAccountDeposit(loadAmountF *big.Int) (*types.Transaction, error) {
	return _Bindings.Contract.CreateAccountDeposit(&_Bindings.TransactOpts, loadAmountF)
}

// Deposit is a paid mutator transaction binding the contract method 0x212bafd7.
//
// Solidity: function deposit(uint48 fromIdx, uint40 loadAmountF) payable returns()
func (_Bindings *BindingsTransactor) Deposit(opts *bind.TransactOpts, fromIdx *big.Int, loadAmountF *big.Int) (*types.Transaction, error) {
	return _Bindings.contract.Transact(opts, "deposit", fromIdx, loadAmountF)
}

// Deposit is a paid mutator transaction binding the contract method 0x212bafd7.
//
// Solidity: function deposit(uint48 fromIdx, uint40 loadAmountF) payable returns()
func (_Bindings *BindingsSession) Deposit(fromIdx *big.Int, loadAmountF *big.Int) (*types.Transaction, error) {
	return _Bindings.Contract.Deposit(&_Bindings.TransactOpts, fromIdx, loadAmountF)
}

// Deposit is a paid mutator transaction binding the contract method 0x212bafd7.
//
// Solidity: function deposit(uint48 fromIdx, uint40 loadAmountF) payable returns()
func (_Bindings *BindingsTransactorSession) Deposit(fromIdx *big.Int, loadAmountF *big.Int) (*types.Transaction, error) {
	return _Bindings.Contract.Deposit(&_Bindings.TransactOpts, fromIdx, loadAmountF)
}

// Exit is a paid mutator transaction binding the contract method 0x8195b790.
//
// Solidity: function exit(uint48 fromIdx, uint40 amountF) returns()
func (_Bindings *BindingsTransactor) Exit(opts *bind.TransactOpts, fromIdx *big.Int, amountF *big.Int) (*types.Transaction, error) {
	return _Bindings.contract.Transact(opts, "exit", fromIdx, amountF)
}

// Exit is a paid mutator transaction binding the contract method 0x8195b790.
//
// Solidity: function exit(uint48 fromIdx, uint40 amountF) returns()
func (_Bindings *BindingsSession) Exit(fromIdx *big.Int, amountF *big.Int) (*types.Transaction, error) {
	return _Bindings.Contract.Exit(&_Bindings.TransactOpts, fromIdx, amountF)
}

// Exit is a paid mutator transaction binding the contract method 0x8195b790.
//
// Solidity: function exit(uint48 fromIdx, uint40 amountF) returns()
func (_Bindings *BindingsTransactorSession) Exit(fromIdx *big.Int, amountF *big.Int) (*types.Transaction, error) {
	return _Bindings.Contract.Exit(&_Bindings.TransactOpts, fromIdx, amountF)
}

// ExplodeMultiple is a paid mutator transaction binding the contract method 0x894bc2b8.
//
// Solidity: function explodeMultiple(uint48 fromIdx, uint48[] toIdxs) returns()
func (_Bindings *BindingsTransactor) ExplodeMultiple(opts *bind.TransactOpts, fromIdx *big.Int, toIdxs []*big.Int) (*types.Transaction, error) {
	return _Bindings.contract.Transact(opts, "explodeMultiple", fromIdx, toIdxs)
}

// ExplodeMultiple is a paid mutator transaction binding the contract method 0x894bc2b8.
//
// Solidity: function explodeMultiple(uint48 fromIdx, uint48[] toIdxs) returns()
func (_Bindings *BindingsSession) ExplodeMultiple(fromIdx *big.Int, toIdxs []*big.Int) (*types.Transaction, error) {
	return _Bindings.Contract.ExplodeMultiple(&_Bindings.TransactOpts, fromIdx, toIdxs)
}

// ExplodeMultiple is a paid mutator transaction binding the contract method 0x894bc2b8.
//
// Solidity: function explodeMultiple(uint48 fromIdx, uint48[] toIdxs) returns()
func (_Bindings *BindingsTransactorSession) ExplodeMultiple(fromIdx *big.Int, toIdxs []*big.Int) (*types.Transaction, error) {
	return _Bindings.Contract.ExplodeMultiple(&_Bindings.TransactOpts, fromIdx, toIdxs)
}

// ForgeBatch is a paid mutator transaction binding the contract method 0x83cd43e0.
//
// Solidity: function forgeBatch(uint48 newLastIdx, uint256 newAccountRoot, uint256 newVouchRoot, uint256 newScoreRoot, uint256 newExitRoot, bytes txsData, uint256[2] proofA, uint256[2][2] proofB, uint256[2] proofC) returns()
func (_Bindings *BindingsTransactor) ForgeBatch(opts *bind.TransactOpts, newLastIdx *big.Int, newAccountRoot *big.Int, newVouchRoot *big.Int, newScoreRoot *big.Int, newExitRoot *big.Int, txsData []byte, proofA [2]*big.Int, proofB [2][2]*big.Int, proofC [2]*big.Int) (*types.Transaction, error) {
	return _Bindings.contract.Transact(opts, "forgeBatch", newLastIdx, newAccountRoot, newVouchRoot, newScoreRoot, newExitRoot, txsData, proofA, proofB, proofC)
}

// ForgeBatch is a paid mutator transaction binding the contract method 0x83cd43e0.
//
// Solidity: function forgeBatch(uint48 newLastIdx, uint256 newAccountRoot, uint256 newVouchRoot, uint256 newScoreRoot, uint256 newExitRoot, bytes txsData, uint256[2] proofA, uint256[2][2] proofB, uint256[2] proofC) returns()
func (_Bindings *BindingsSession) ForgeBatch(newLastIdx *big.Int, newAccountRoot *big.Int, newVouchRoot *big.Int, newScoreRoot *big.Int, newExitRoot *big.Int, txsData []byte, proofA [2]*big.Int, proofB [2][2]*big.Int, proofC [2]*big.Int) (*types.Transaction, error) {
	return _Bindings.Contract.ForgeBatch(&_Bindings.TransactOpts, newLastIdx, newAccountRoot, newVouchRoot, newScoreRoot, newExitRoot, txsData, proofA, proofB, proofC)
}

// ForgeBatch is a paid mutator transaction binding the contract method 0x83cd43e0.
//
// Solidity: function forgeBatch(uint48 newLastIdx, uint256 newAccountRoot, uint256 newVouchRoot, uint256 newScoreRoot, uint256 newExitRoot, bytes txsData, uint256[2] proofA, uint256[2][2] proofB, uint256[2] proofC) returns()
func (_Bindings *BindingsTransactorSession) ForgeBatch(newLastIdx *big.Int, newAccountRoot *big.Int, newVouchRoot *big.Int, newScoreRoot *big.Int, newExitRoot *big.Int, txsData []byte, proofA [2]*big.Int, proofB [2][2]*big.Int, proofC [2]*big.Int) (*types.Transaction, error) {
	return _Bindings.Contract.ForgeBatch(&_Bindings.TransactOpts, newLastIdx, newAccountRoot, newVouchRoot, newScoreRoot, newExitRoot, txsData, proofA, proofB, proofC)
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

// Unvouch is a paid mutator transaction binding the contract method 0xc1b190c0.
//
// Solidity: function unvouch(uint48 fromIdx, uint48 toIdx) returns()
func (_Bindings *BindingsTransactor) Unvouch(opts *bind.TransactOpts, fromIdx *big.Int, toIdx *big.Int) (*types.Transaction, error) {
	return _Bindings.contract.Transact(opts, "unvouch", fromIdx, toIdx)
}

// Unvouch is a paid mutator transaction binding the contract method 0xc1b190c0.
//
// Solidity: function unvouch(uint48 fromIdx, uint48 toIdx) returns()
func (_Bindings *BindingsSession) Unvouch(fromIdx *big.Int, toIdx *big.Int) (*types.Transaction, error) {
	return _Bindings.Contract.Unvouch(&_Bindings.TransactOpts, fromIdx, toIdx)
}

// Unvouch is a paid mutator transaction binding the contract method 0xc1b190c0.
//
// Solidity: function unvouch(uint48 fromIdx, uint48 toIdx) returns()
func (_Bindings *BindingsTransactorSession) Unvouch(fromIdx *big.Int, toIdx *big.Int) (*types.Transaction, error) {
	return _Bindings.Contract.Unvouch(&_Bindings.TransactOpts, fromIdx, toIdx)
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

// Vouch is a paid mutator transaction binding the contract method 0x11954d3c.
//
// Solidity: function vouch(uint48 fromIdx, uint48 toIdx) returns()
func (_Bindings *BindingsTransactor) Vouch(opts *bind.TransactOpts, fromIdx *big.Int, toIdx *big.Int) (*types.Transaction, error) {
	return _Bindings.contract.Transact(opts, "vouch", fromIdx, toIdx)
}

// Vouch is a paid mutator transaction binding the contract method 0x11954d3c.
//
// Solidity: function vouch(uint48 fromIdx, uint48 toIdx) returns()
func (_Bindings *BindingsSession) Vouch(fromIdx *big.Int, toIdx *big.Int) (*types.Transaction, error) {
	return _Bindings.Contract.Vouch(&_Bindings.TransactOpts, fromIdx, toIdx)
}

// Vouch is a paid mutator transaction binding the contract method 0x11954d3c.
//
// Solidity: function vouch(uint48 fromIdx, uint48 toIdx) returns()
func (_Bindings *BindingsTransactorSession) Vouch(fromIdx *big.Int, toIdx *big.Int) (*types.Transaction, error) {
	return _Bindings.Contract.Vouch(&_Bindings.TransactOpts, fromIdx, toIdx)
}

// WithdrawMerkleProof is a paid mutator transaction binding the contract method 0x62332ee2.
//
// Solidity: function withdrawMerkleProof(uint192 amount, uint32 numExitRoot, uint256[] siblings, uint48 idx) returns()
func (_Bindings *BindingsTransactor) WithdrawMerkleProof(opts *bind.TransactOpts, amount *big.Int, numExitRoot uint32, siblings []*big.Int, idx *big.Int) (*types.Transaction, error) {
	return _Bindings.contract.Transact(opts, "withdrawMerkleProof", amount, numExitRoot, siblings, idx)
}

// WithdrawMerkleProof is a paid mutator transaction binding the contract method 0x62332ee2.
//
// Solidity: function withdrawMerkleProof(uint192 amount, uint32 numExitRoot, uint256[] siblings, uint48 idx) returns()
func (_Bindings *BindingsSession) WithdrawMerkleProof(amount *big.Int, numExitRoot uint32, siblings []*big.Int, idx *big.Int) (*types.Transaction, error) {
	return _Bindings.Contract.WithdrawMerkleProof(&_Bindings.TransactOpts, amount, numExitRoot, siblings, idx)
}

// WithdrawMerkleProof is a paid mutator transaction binding the contract method 0x62332ee2.
//
// Solidity: function withdrawMerkleProof(uint192 amount, uint32 numExitRoot, uint256[] siblings, uint48 idx) returns()
func (_Bindings *BindingsTransactorSession) WithdrawMerkleProof(amount *big.Int, numExitRoot uint32, siblings []*big.Int, idx *big.Int) (*types.Transaction, error) {
	return _Bindings.Contract.WithdrawMerkleProof(&_Bindings.TransactOpts, amount, numExitRoot, siblings, idx)
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
