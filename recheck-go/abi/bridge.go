// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package abi

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
)

// BridgeMetaData contains all meta data concerning the Bridge contract.
var BridgeMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"stateMutability\":\"nonpayable\",\"inputs\":[{\"type\":\"address\",\"name\":\"_owner\",\"internalType\":\"address\"}]},{\"type\":\"event\",\"name\":\"Deposit\",\"inputs\":[{\"type\":\"uint256\",\"name\":\"toChainId\",\"internalType\":\"uint256\",\"indexed\":false},{\"type\":\"address\",\"name\":\"fromToken\",\"internalType\":\"address\",\"indexed\":false},{\"type\":\"address\",\"name\":\"toToken\",\"internalType\":\"address\",\"indexed\":false},{\"type\":\"address\",\"name\":\"recipient\",\"internalType\":\"address\",\"indexed\":false},{\"type\":\"uint256\",\"name\":\"value\",\"internalType\":\"uint256\",\"indexed\":false}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"DepositNative\",\"inputs\":[{\"type\":\"uint256\",\"name\":\"toChainId\",\"internalType\":\"uint256\",\"indexed\":false},{\"type\":\"bool\",\"name\":\"isMain\",\"internalType\":\"bool\",\"indexed\":false},{\"type\":\"address\",\"name\":\"recipient\",\"internalType\":\"address\",\"indexed\":false},{\"type\":\"uint256\",\"name\":\"value\",\"internalType\":\"uint256\",\"indexed\":false}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"WithdrawDone\",\"inputs\":[{\"type\":\"uint256\",\"name\":\"toChainId\",\"internalType\":\"uint256\",\"indexed\":false},{\"type\":\"address\",\"name\":\"fromToken\",\"internalType\":\"address\",\"indexed\":false},{\"type\":\"address\",\"name\":\"toToken\",\"internalType\":\"address\",\"indexed\":false},{\"type\":\"address\",\"name\":\"recipient\",\"internalType\":\"address\",\"indexed\":false},{\"type\":\"uint256\",\"name\":\"value\",\"internalType\":\"uint256\",\"indexed\":false},{\"type\":\"bytes\",\"name\":\"depositHash\",\"internalType\":\"bytes\",\"indexed\":false}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"WithdrawNativeDone\",\"inputs\":[{\"type\":\"uint256\",\"name\":\"fromChainId\",\"internalType\":\"uint256\",\"indexed\":false},{\"type\":\"address\",\"name\":\"recipient\",\"internalType\":\"address\",\"indexed\":false},{\"type\":\"bool\",\"name\":\"isMain\",\"internalType\":\"bool\",\"indexed\":false},{\"type\":\"uint256\",\"name\":\"value\",\"internalType\":\"uint256\",\"indexed\":false},{\"type\":\"bytes\",\"name\":\"depositHash\",\"internalType\":\"bytes\",\"indexed\":false}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"adminChanged\",\"inputs\":[{\"type\":\"address\",\"name\":\"_address\",\"internalType\":\"address\",\"indexed\":false}],\"anonymous\":false},{\"type\":\"function\",\"stateMutability\":\"view\",\"outputs\":[{\"type\":\"address\",\"name\":\"\",\"internalType\":\"address\"}],\"name\":\"admin\",\"inputs\":[]},{\"type\":\"function\",\"stateMutability\":\"nonpayable\",\"outputs\":[],\"name\":\"deposit\",\"inputs\":[{\"type\":\"uint256\",\"name\":\"chainId\",\"internalType\":\"uint256\"},{\"type\":\"address\",\"name\":\"toToken\",\"internalType\":\"address\"},{\"type\":\"uint256\",\"name\":\"value\",\"internalType\":\"uint256\"}]},{\"type\":\"function\",\"stateMutability\":\"payable\",\"outputs\":[],\"name\":\"depositNative\",\"inputs\":[{\"type\":\"uint256\",\"name\":\"toChainId\",\"internalType\":\"uint256\"},{\"type\":\"bool\",\"name\":\"isMain\",\"internalType\":\"bool\"},{\"type\":\"uint256\",\"name\":\"value\",\"internalType\":\"uint256\"}]},{\"type\":\"function\",\"stateMutability\":\"view\",\"outputs\":[{\"type\":\"address\",\"name\":\"\",\"internalType\":\"address\"}],\"name\":\"manager\",\"inputs\":[]},{\"type\":\"function\",\"stateMutability\":\"nonpayable\",\"outputs\":[],\"name\":\"nativeInsert\",\"inputs\":[{\"type\":\"uint256\",\"name\":\"toChainId\",\"internalType\":\"uint256\"},{\"type\":\"bool\",\"name\":\"isRun\",\"internalType\":\"bool\"},{\"type\":\"bool\",\"name\":\"isMain\",\"internalType\":\"bool\"},{\"type\":\"address\",\"name\":\"fromAddress\",\"internalType\":\"address\"}]},{\"type\":\"function\",\"stateMutability\":\"nonpayable\",\"outputs\":[],\"name\":\"nativeTransfer\",\"inputs\":[{\"type\":\"address\",\"name\":\"recipient\",\"internalType\":\"addresspayable\"},{\"type\":\"uint256\",\"name\":\"value\",\"internalType\":\"uint256\"}]},{\"type\":\"function\",\"stateMutability\":\"view\",\"outputs\":[{\"type\":\"bool\",\"name\":\"isRun\",\"internalType\":\"bool\"},{\"type\":\"bool\",\"name\":\"isMain\",\"internalType\":\"bool\"},{\"type\":\"address\",\"name\":\"local\",\"internalType\":\"address\"}],\"name\":\"natives\",\"inputs\":[{\"type\":\"uint256\",\"name\":\"\",\"internalType\":\"uint256\"},{\"type\":\"bool\",\"name\":\"\",\"internalType\":\"bool\"}]},{\"type\":\"function\",\"stateMutability\":\"view\",\"outputs\":[{\"type\":\"address\",\"name\":\"\",\"internalType\":\"address\"}],\"name\":\"owner\",\"inputs\":[]},{\"type\":\"function\",\"stateMutability\":\"nonpayable\",\"outputs\":[],\"name\":\"setAdmin\",\"inputs\":[{\"type\":\"address\",\"name\":\"newAdmin\",\"internalType\":\"addresspayable\"}]},{\"type\":\"function\",\"stateMutability\":\"nonpayable\",\"outputs\":[],\"name\":\"setManager\",\"inputs\":[{\"type\":\"address\",\"name\":\"_manager\",\"internalType\":\"address\"}]},{\"type\":\"function\",\"stateMutability\":\"nonpayable\",\"outputs\":[],\"name\":\"setNativeIsRun\",\"inputs\":[{\"type\":\"uint256\",\"name\":\"toChainId\",\"internalType\":\"uint256\"},{\"type\":\"bool\",\"name\":\"isMain\",\"internalType\":\"bool\"},{\"type\":\"bool\",\"name\":\"state\",\"internalType\":\"bool\"}]},{\"type\":\"function\",\"stateMutability\":\"nonpayable\",\"outputs\":[],\"name\":\"setOwner\",\"inputs\":[{\"type\":\"address\",\"name\":\"newOwner\",\"internalType\":\"addresspayable\"}]},{\"type\":\"function\",\"stateMutability\":\"nonpayable\",\"outputs\":[],\"name\":\"setTokenIsRun\",\"inputs\":[{\"type\":\"uint256\",\"name\":\"toChainId\",\"internalType\":\"uint256\"},{\"type\":\"address\",\"name\":\"toToken\",\"internalType\":\"address\"},{\"type\":\"bool\",\"name\":\"state\",\"internalType\":\"bool\"}]},{\"type\":\"function\",\"stateMutability\":\"nonpayable\",\"outputs\":[],\"name\":\"tokenInsert\",\"inputs\":[{\"type\":\"uint256\",\"name\":\"toChainId\",\"internalType\":\"uint256\"},{\"type\":\"address\",\"name\":\"toToken\",\"internalType\":\"address\"},{\"type\":\"address\",\"name\":\"fromToken\",\"internalType\":\"address\"},{\"type\":\"bool\",\"name\":\"isRun\",\"internalType\":\"bool\"},{\"type\":\"bool\",\"name\":\"isMain\",\"internalType\":\"bool\"}]},{\"type\":\"function\",\"stateMutability\":\"nonpayable\",\"outputs\":[],\"name\":\"tokenTransfer\",\"inputs\":[{\"type\":\"address\",\"name\":\"fromToken\",\"internalType\":\"address\"},{\"type\":\"address\",\"name\":\"recipient\",\"internalType\":\"address\"},{\"type\":\"uint256\",\"name\":\"value\",\"internalType\":\"uint256\"}]},{\"type\":\"function\",\"stateMutability\":\"view\",\"outputs\":[{\"type\":\"bool\",\"name\":\"isRun\",\"internalType\":\"bool\"},{\"type\":\"bool\",\"name\":\"isMain\",\"internalType\":\"bool\"},{\"type\":\"address\",\"name\":\"local\",\"internalType\":\"address\"}],\"name\":\"tokens\",\"inputs\":[{\"type\":\"uint256\",\"name\":\"\",\"internalType\":\"uint256\"},{\"type\":\"address\",\"name\":\"\",\"internalType\":\"address\"}]},{\"type\":\"function\",\"stateMutability\":\"nonpayable\",\"outputs\":[],\"name\":\"withdraw\",\"inputs\":[{\"type\":\"uint256\",\"name\":\"toChainId\",\"internalType\":\"uint256\"},{\"type\":\"address\",\"name\":\"toToken\",\"internalType\":\"address\"},{\"type\":\"address\",\"name\":\"recipient\",\"internalType\":\"address\"},{\"type\":\"uint256\",\"name\":\"value\",\"internalType\":\"uint256\"},{\"type\":\"bytes\",\"name\":\"depositHash\",\"internalType\":\"bytes\"}]},{\"type\":\"function\",\"stateMutability\":\"nonpayable\",\"outputs\":[],\"name\":\"withdrawNative\",\"inputs\":[{\"type\":\"uint256\",\"name\":\"toChainId\",\"internalType\":\"uint256\"},{\"type\":\"address\",\"name\":\"recipient\",\"internalType\":\"addresspayable\"},{\"type\":\"bool\",\"name\":\"isMain\",\"internalType\":\"bool\"},{\"type\":\"uint256\",\"name\":\"value\",\"internalType\":\"uint256\"},{\"type\":\"bytes\",\"name\":\"depositHash\",\"internalType\":\"bytes\"}]},{\"type\":\"receive\",\"stateMutability\":\"payable\"}]",
}

// BridgeABI is the input ABI used to generate the binding from.
// Deprecated: Use BridgeMetaData.ABI instead.
var BridgeABI = BridgeMetaData.ABI

// Bridge is an auto generated Go binding around an Ethereum contract.
type Bridge struct {
	BridgeCaller     // Read-only binding to the contract
	BridgeTransactor // Write-only binding to the contract
	BridgeFilterer   // Log filterer for contract events
}

// BridgeCaller is an auto generated read-only Go binding around an Ethereum contract.
type BridgeCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BridgeTransactor is an auto generated write-only Go binding around an Ethereum contract.
type BridgeTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BridgeFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type BridgeFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BridgeSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type BridgeSession struct {
	Contract     *Bridge           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BridgeCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type BridgeCallerSession struct {
	Contract *BridgeCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// BridgeTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type BridgeTransactorSession struct {
	Contract     *BridgeTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BridgeRaw is an auto generated low-level Go binding around an Ethereum contract.
type BridgeRaw struct {
	Contract *Bridge // Generic contract binding to access the raw methods on
}

// BridgeCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type BridgeCallerRaw struct {
	Contract *BridgeCaller // Generic read-only contract binding to access the raw methods on
}

// BridgeTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type BridgeTransactorRaw struct {
	Contract *BridgeTransactor // Generic write-only contract binding to access the raw methods on
}

// NewBridge creates a new instance of Bridge, bound to a specific deployed contract.
func NewBridge(address common.Address, backend bind.ContractBackend) (*Bridge, error) {
	contract, err := bindBridge(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Bridge{BridgeCaller: BridgeCaller{contract: contract}, BridgeTransactor: BridgeTransactor{contract: contract}, BridgeFilterer: BridgeFilterer{contract: contract}}, nil
}

// NewBridgeCaller creates a new read-only instance of Bridge, bound to a specific deployed contract.
func NewBridgeCaller(address common.Address, caller bind.ContractCaller) (*BridgeCaller, error) {
	contract, err := bindBridge(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &BridgeCaller{contract: contract}, nil
}

// NewBridgeTransactor creates a new write-only instance of Bridge, bound to a specific deployed contract.
func NewBridgeTransactor(address common.Address, transactor bind.ContractTransactor) (*BridgeTransactor, error) {
	contract, err := bindBridge(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &BridgeTransactor{contract: contract}, nil
}

// NewBridgeFilterer creates a new log filterer instance of Bridge, bound to a specific deployed contract.
func NewBridgeFilterer(address common.Address, filterer bind.ContractFilterer) (*BridgeFilterer, error) {
	contract, err := bindBridge(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &BridgeFilterer{contract: contract}, nil
}

// bindBridge binds a generic wrapper to an already deployed contract.
func bindBridge(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(BridgeABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Bridge *BridgeRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Bridge.Contract.BridgeCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Bridge *BridgeRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Bridge.Contract.BridgeTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Bridge *BridgeRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Bridge.Contract.BridgeTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Bridge *BridgeCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Bridge.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Bridge *BridgeTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Bridge.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Bridge *BridgeTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Bridge.Contract.contract.Transact(opts, method, params...)
}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() view returns(address)
func (_Bridge *BridgeCaller) Admin(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Bridge.contract.Call(opts, &out, "admin")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() view returns(address)
func (_Bridge *BridgeSession) Admin() (common.Address, error) {
	return _Bridge.Contract.Admin(&_Bridge.CallOpts)
}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() view returns(address)
func (_Bridge *BridgeCallerSession) Admin() (common.Address, error) {
	return _Bridge.Contract.Admin(&_Bridge.CallOpts)
}

// Manager is a free data retrieval call binding the contract method 0x481c6a75.
//
// Solidity: function manager() view returns(address)
func (_Bridge *BridgeCaller) Manager(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Bridge.contract.Call(opts, &out, "manager")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Manager is a free data retrieval call binding the contract method 0x481c6a75.
//
// Solidity: function manager() view returns(address)
func (_Bridge *BridgeSession) Manager() (common.Address, error) {
	return _Bridge.Contract.Manager(&_Bridge.CallOpts)
}

// Manager is a free data retrieval call binding the contract method 0x481c6a75.
//
// Solidity: function manager() view returns(address)
func (_Bridge *BridgeCallerSession) Manager() (common.Address, error) {
	return _Bridge.Contract.Manager(&_Bridge.CallOpts)
}

// Natives is a free data retrieval call binding the contract method 0x7aabe0a6.
//
// Solidity: function natives(uint256 , bool ) view returns(bool isRun, bool isMain, address local)
func (_Bridge *BridgeCaller) Natives(opts *bind.CallOpts, arg0 *big.Int, arg1 bool) (struct {
	IsRun  bool
	IsMain bool
	Local  common.Address
}, error) {
	var out []interface{}
	err := _Bridge.contract.Call(opts, &out, "natives", arg0, arg1)

	outstruct := new(struct {
		IsRun  bool
		IsMain bool
		Local  common.Address
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.IsRun = *abi.ConvertType(out[0], new(bool)).(*bool)
	outstruct.IsMain = *abi.ConvertType(out[1], new(bool)).(*bool)
	outstruct.Local = *abi.ConvertType(out[2], new(common.Address)).(*common.Address)

	return *outstruct, err

}

// Natives is a free data retrieval call binding the contract method 0x7aabe0a6.
//
// Solidity: function natives(uint256 , bool ) view returns(bool isRun, bool isMain, address local)
func (_Bridge *BridgeSession) Natives(arg0 *big.Int, arg1 bool) (struct {
	IsRun  bool
	IsMain bool
	Local  common.Address
}, error) {
	return _Bridge.Contract.Natives(&_Bridge.CallOpts, arg0, arg1)
}

// Natives is a free data retrieval call binding the contract method 0x7aabe0a6.
//
// Solidity: function natives(uint256 , bool ) view returns(bool isRun, bool isMain, address local)
func (_Bridge *BridgeCallerSession) Natives(arg0 *big.Int, arg1 bool) (struct {
	IsRun  bool
	IsMain bool
	Local  common.Address
}, error) {
	return _Bridge.Contract.Natives(&_Bridge.CallOpts, arg0, arg1)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Bridge *BridgeCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Bridge.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Bridge *BridgeSession) Owner() (common.Address, error) {
	return _Bridge.Contract.Owner(&_Bridge.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Bridge *BridgeCallerSession) Owner() (common.Address, error) {
	return _Bridge.Contract.Owner(&_Bridge.CallOpts)
}

// Tokens is a free data retrieval call binding the contract method 0xae68a2ba.
//
// Solidity: function tokens(uint256 , address ) view returns(bool isRun, bool isMain, address local)
func (_Bridge *BridgeCaller) Tokens(opts *bind.CallOpts, arg0 *big.Int, arg1 common.Address) (struct {
	IsRun  bool
	IsMain bool
	Local  common.Address
}, error) {
	var out []interface{}
	err := _Bridge.contract.Call(opts, &out, "tokens", arg0, arg1)

	outstruct := new(struct {
		IsRun  bool
		IsMain bool
		Local  common.Address
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.IsRun = *abi.ConvertType(out[0], new(bool)).(*bool)
	outstruct.IsMain = *abi.ConvertType(out[1], new(bool)).(*bool)
	outstruct.Local = *abi.ConvertType(out[2], new(common.Address)).(*common.Address)

	return *outstruct, err

}

// Tokens is a free data retrieval call binding the contract method 0xae68a2ba.
//
// Solidity: function tokens(uint256 , address ) view returns(bool isRun, bool isMain, address local)
func (_Bridge *BridgeSession) Tokens(arg0 *big.Int, arg1 common.Address) (struct {
	IsRun  bool
	IsMain bool
	Local  common.Address
}, error) {
	return _Bridge.Contract.Tokens(&_Bridge.CallOpts, arg0, arg1)
}

// Tokens is a free data retrieval call binding the contract method 0xae68a2ba.
//
// Solidity: function tokens(uint256 , address ) view returns(bool isRun, bool isMain, address local)
func (_Bridge *BridgeCallerSession) Tokens(arg0 *big.Int, arg1 common.Address) (struct {
	IsRun  bool
	IsMain bool
	Local  common.Address
}, error) {
	return _Bridge.Contract.Tokens(&_Bridge.CallOpts, arg0, arg1)
}

// Deposit is a paid mutator transaction binding the contract method 0xbc157ac1.
//
// Solidity: function deposit(uint256 chainId, address toToken, uint256 value) returns()
func (_Bridge *BridgeTransactor) Deposit(opts *bind.TransactOpts, chainId *big.Int, toToken common.Address, value *big.Int) (*types.Transaction, error) {
	return _Bridge.contract.Transact(opts, "deposit", chainId, toToken, value)
}

// Deposit is a paid mutator transaction binding the contract method 0xbc157ac1.
//
// Solidity: function deposit(uint256 chainId, address toToken, uint256 value) returns()
func (_Bridge *BridgeSession) Deposit(chainId *big.Int, toToken common.Address, value *big.Int) (*types.Transaction, error) {
	return _Bridge.Contract.Deposit(&_Bridge.TransactOpts, chainId, toToken, value)
}

// Deposit is a paid mutator transaction binding the contract method 0xbc157ac1.
//
// Solidity: function deposit(uint256 chainId, address toToken, uint256 value) returns()
func (_Bridge *BridgeTransactorSession) Deposit(chainId *big.Int, toToken common.Address, value *big.Int) (*types.Transaction, error) {
	return _Bridge.Contract.Deposit(&_Bridge.TransactOpts, chainId, toToken, value)
}

// DepositNative is a paid mutator transaction binding the contract method 0x02367ad2.
//
// Solidity: function depositNative(uint256 toChainId, bool isMain, uint256 value) payable returns()
func (_Bridge *BridgeTransactor) DepositNative(opts *bind.TransactOpts, toChainId *big.Int, isMain bool, value *big.Int) (*types.Transaction, error) {
	return _Bridge.contract.Transact(opts, "depositNative", toChainId, isMain, value)
}

// DepositNative is a paid mutator transaction binding the contract method 0x02367ad2.
//
// Solidity: function depositNative(uint256 toChainId, bool isMain, uint256 value) payable returns()
func (_Bridge *BridgeSession) DepositNative(toChainId *big.Int, isMain bool, value *big.Int) (*types.Transaction, error) {
	return _Bridge.Contract.DepositNative(&_Bridge.TransactOpts, toChainId, isMain, value)
}

// DepositNative is a paid mutator transaction binding the contract method 0x02367ad2.
//
// Solidity: function depositNative(uint256 toChainId, bool isMain, uint256 value) payable returns()
func (_Bridge *BridgeTransactorSession) DepositNative(toChainId *big.Int, isMain bool, value *big.Int) (*types.Transaction, error) {
	return _Bridge.Contract.DepositNative(&_Bridge.TransactOpts, toChainId, isMain, value)
}

// NativeInsert is a paid mutator transaction binding the contract method 0xffb4aedc.
//
// Solidity: function nativeInsert(uint256 toChainId, bool isRun, bool isMain, address fromAddress) returns()
func (_Bridge *BridgeTransactor) NativeInsert(opts *bind.TransactOpts, toChainId *big.Int, isRun bool, isMain bool, fromAddress common.Address) (*types.Transaction, error) {
	return _Bridge.contract.Transact(opts, "nativeInsert", toChainId, isRun, isMain, fromAddress)
}

// NativeInsert is a paid mutator transaction binding the contract method 0xffb4aedc.
//
// Solidity: function nativeInsert(uint256 toChainId, bool isRun, bool isMain, address fromAddress) returns()
func (_Bridge *BridgeSession) NativeInsert(toChainId *big.Int, isRun bool, isMain bool, fromAddress common.Address) (*types.Transaction, error) {
	return _Bridge.Contract.NativeInsert(&_Bridge.TransactOpts, toChainId, isRun, isMain, fromAddress)
}

// NativeInsert is a paid mutator transaction binding the contract method 0xffb4aedc.
//
// Solidity: function nativeInsert(uint256 toChainId, bool isRun, bool isMain, address fromAddress) returns()
func (_Bridge *BridgeTransactorSession) NativeInsert(toChainId *big.Int, isRun bool, isMain bool, fromAddress common.Address) (*types.Transaction, error) {
	return _Bridge.Contract.NativeInsert(&_Bridge.TransactOpts, toChainId, isRun, isMain, fromAddress)
}

// NativeTransfer is a paid mutator transaction binding the contract method 0xf2522bcd.
//
// Solidity: function nativeTransfer(address recipient, uint256 value) returns()
func (_Bridge *BridgeTransactor) NativeTransfer(opts *bind.TransactOpts, recipient common.Address, value *big.Int) (*types.Transaction, error) {
	return _Bridge.contract.Transact(opts, "nativeTransfer", recipient, value)
}

// NativeTransfer is a paid mutator transaction binding the contract method 0xf2522bcd.
//
// Solidity: function nativeTransfer(address recipient, uint256 value) returns()
func (_Bridge *BridgeSession) NativeTransfer(recipient common.Address, value *big.Int) (*types.Transaction, error) {
	return _Bridge.Contract.NativeTransfer(&_Bridge.TransactOpts, recipient, value)
}

// NativeTransfer is a paid mutator transaction binding the contract method 0xf2522bcd.
//
// Solidity: function nativeTransfer(address recipient, uint256 value) returns()
func (_Bridge *BridgeTransactorSession) NativeTransfer(recipient common.Address, value *big.Int) (*types.Transaction, error) {
	return _Bridge.Contract.NativeTransfer(&_Bridge.TransactOpts, recipient, value)
}

// SetAdmin is a paid mutator transaction binding the contract method 0x704b6c02.
//
// Solidity: function setAdmin(address newAdmin) returns()
func (_Bridge *BridgeTransactor) SetAdmin(opts *bind.TransactOpts, newAdmin common.Address) (*types.Transaction, error) {
	return _Bridge.contract.Transact(opts, "setAdmin", newAdmin)
}

// SetAdmin is a paid mutator transaction binding the contract method 0x704b6c02.
//
// Solidity: function setAdmin(address newAdmin) returns()
func (_Bridge *BridgeSession) SetAdmin(newAdmin common.Address) (*types.Transaction, error) {
	return _Bridge.Contract.SetAdmin(&_Bridge.TransactOpts, newAdmin)
}

// SetAdmin is a paid mutator transaction binding the contract method 0x704b6c02.
//
// Solidity: function setAdmin(address newAdmin) returns()
func (_Bridge *BridgeTransactorSession) SetAdmin(newAdmin common.Address) (*types.Transaction, error) {
	return _Bridge.Contract.SetAdmin(&_Bridge.TransactOpts, newAdmin)
}

// SetManager is a paid mutator transaction binding the contract method 0xd0ebdbe7.
//
// Solidity: function setManager(address _manager) returns()
func (_Bridge *BridgeTransactor) SetManager(opts *bind.TransactOpts, _manager common.Address) (*types.Transaction, error) {
	return _Bridge.contract.Transact(opts, "setManager", _manager)
}

// SetManager is a paid mutator transaction binding the contract method 0xd0ebdbe7.
//
// Solidity: function setManager(address _manager) returns()
func (_Bridge *BridgeSession) SetManager(_manager common.Address) (*types.Transaction, error) {
	return _Bridge.Contract.SetManager(&_Bridge.TransactOpts, _manager)
}

// SetManager is a paid mutator transaction binding the contract method 0xd0ebdbe7.
//
// Solidity: function setManager(address _manager) returns()
func (_Bridge *BridgeTransactorSession) SetManager(_manager common.Address) (*types.Transaction, error) {
	return _Bridge.Contract.SetManager(&_Bridge.TransactOpts, _manager)
}

// SetNativeIsRun is a paid mutator transaction binding the contract method 0x9a234732.
//
// Solidity: function setNativeIsRun(uint256 toChainId, bool isMain, bool state) returns()
func (_Bridge *BridgeTransactor) SetNativeIsRun(opts *bind.TransactOpts, toChainId *big.Int, isMain bool, state bool) (*types.Transaction, error) {
	return _Bridge.contract.Transact(opts, "setNativeIsRun", toChainId, isMain, state)
}

// SetNativeIsRun is a paid mutator transaction binding the contract method 0x9a234732.
//
// Solidity: function setNativeIsRun(uint256 toChainId, bool isMain, bool state) returns()
func (_Bridge *BridgeSession) SetNativeIsRun(toChainId *big.Int, isMain bool, state bool) (*types.Transaction, error) {
	return _Bridge.Contract.SetNativeIsRun(&_Bridge.TransactOpts, toChainId, isMain, state)
}

// SetNativeIsRun is a paid mutator transaction binding the contract method 0x9a234732.
//
// Solidity: function setNativeIsRun(uint256 toChainId, bool isMain, bool state) returns()
func (_Bridge *BridgeTransactorSession) SetNativeIsRun(toChainId *big.Int, isMain bool, state bool) (*types.Transaction, error) {
	return _Bridge.Contract.SetNativeIsRun(&_Bridge.TransactOpts, toChainId, isMain, state)
}

// SetOwner is a paid mutator transaction binding the contract method 0x13af4035.
//
// Solidity: function setOwner(address newOwner) returns()
func (_Bridge *BridgeTransactor) SetOwner(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Bridge.contract.Transact(opts, "setOwner", newOwner)
}

// SetOwner is a paid mutator transaction binding the contract method 0x13af4035.
//
// Solidity: function setOwner(address newOwner) returns()
func (_Bridge *BridgeSession) SetOwner(newOwner common.Address) (*types.Transaction, error) {
	return _Bridge.Contract.SetOwner(&_Bridge.TransactOpts, newOwner)
}

// SetOwner is a paid mutator transaction binding the contract method 0x13af4035.
//
// Solidity: function setOwner(address newOwner) returns()
func (_Bridge *BridgeTransactorSession) SetOwner(newOwner common.Address) (*types.Transaction, error) {
	return _Bridge.Contract.SetOwner(&_Bridge.TransactOpts, newOwner)
}

// SetTokenIsRun is a paid mutator transaction binding the contract method 0x2eef73b3.
//
// Solidity: function setTokenIsRun(uint256 toChainId, address toToken, bool state) returns()
func (_Bridge *BridgeTransactor) SetTokenIsRun(opts *bind.TransactOpts, toChainId *big.Int, toToken common.Address, state bool) (*types.Transaction, error) {
	return _Bridge.contract.Transact(opts, "setTokenIsRun", toChainId, toToken, state)
}

// SetTokenIsRun is a paid mutator transaction binding the contract method 0x2eef73b3.
//
// Solidity: function setTokenIsRun(uint256 toChainId, address toToken, bool state) returns()
func (_Bridge *BridgeSession) SetTokenIsRun(toChainId *big.Int, toToken common.Address, state bool) (*types.Transaction, error) {
	return _Bridge.Contract.SetTokenIsRun(&_Bridge.TransactOpts, toChainId, toToken, state)
}

// SetTokenIsRun is a paid mutator transaction binding the contract method 0x2eef73b3.
//
// Solidity: function setTokenIsRun(uint256 toChainId, address toToken, bool state) returns()
func (_Bridge *BridgeTransactorSession) SetTokenIsRun(toChainId *big.Int, toToken common.Address, state bool) (*types.Transaction, error) {
	return _Bridge.Contract.SetTokenIsRun(&_Bridge.TransactOpts, toChainId, toToken, state)
}

// TokenInsert is a paid mutator transaction binding the contract method 0xd523a97e.
//
// Solidity: function tokenInsert(uint256 toChainId, address toToken, address fromToken, bool isRun, bool isMain) returns()
func (_Bridge *BridgeTransactor) TokenInsert(opts *bind.TransactOpts, toChainId *big.Int, toToken common.Address, fromToken common.Address, isRun bool, isMain bool) (*types.Transaction, error) {
	return _Bridge.contract.Transact(opts, "tokenInsert", toChainId, toToken, fromToken, isRun, isMain)
}

// TokenInsert is a paid mutator transaction binding the contract method 0xd523a97e.
//
// Solidity: function tokenInsert(uint256 toChainId, address toToken, address fromToken, bool isRun, bool isMain) returns()
func (_Bridge *BridgeSession) TokenInsert(toChainId *big.Int, toToken common.Address, fromToken common.Address, isRun bool, isMain bool) (*types.Transaction, error) {
	return _Bridge.Contract.TokenInsert(&_Bridge.TransactOpts, toChainId, toToken, fromToken, isRun, isMain)
}

// TokenInsert is a paid mutator transaction binding the contract method 0xd523a97e.
//
// Solidity: function tokenInsert(uint256 toChainId, address toToken, address fromToken, bool isRun, bool isMain) returns()
func (_Bridge *BridgeTransactorSession) TokenInsert(toChainId *big.Int, toToken common.Address, fromToken common.Address, isRun bool, isMain bool) (*types.Transaction, error) {
	return _Bridge.Contract.TokenInsert(&_Bridge.TransactOpts, toChainId, toToken, fromToken, isRun, isMain)
}

// TokenTransfer is a paid mutator transaction binding the contract method 0x15eaef6b.
//
// Solidity: function tokenTransfer(address fromToken, address recipient, uint256 value) returns()
func (_Bridge *BridgeTransactor) TokenTransfer(opts *bind.TransactOpts, fromToken common.Address, recipient common.Address, value *big.Int) (*types.Transaction, error) {
	return _Bridge.contract.Transact(opts, "tokenTransfer", fromToken, recipient, value)
}

// TokenTransfer is a paid mutator transaction binding the contract method 0x15eaef6b.
//
// Solidity: function tokenTransfer(address fromToken, address recipient, uint256 value) returns()
func (_Bridge *BridgeSession) TokenTransfer(fromToken common.Address, recipient common.Address, value *big.Int) (*types.Transaction, error) {
	return _Bridge.Contract.TokenTransfer(&_Bridge.TransactOpts, fromToken, recipient, value)
}

// TokenTransfer is a paid mutator transaction binding the contract method 0x15eaef6b.
//
// Solidity: function tokenTransfer(address fromToken, address recipient, uint256 value) returns()
func (_Bridge *BridgeTransactorSession) TokenTransfer(fromToken common.Address, recipient common.Address, value *big.Int) (*types.Transaction, error) {
	return _Bridge.Contract.TokenTransfer(&_Bridge.TransactOpts, fromToken, recipient, value)
}

// Withdraw is a paid mutator transaction binding the contract method 0xae57de0c.
//
// Solidity: function withdraw(uint256 toChainId, address toToken, address recipient, uint256 value, bytes depositHash) returns()
func (_Bridge *BridgeTransactor) Withdraw(opts *bind.TransactOpts, toChainId *big.Int, toToken common.Address, recipient common.Address, value *big.Int, depositHash []byte) (*types.Transaction, error) {
	return _Bridge.contract.Transact(opts, "withdraw", toChainId, toToken, recipient, value, depositHash)
}

// Withdraw is a paid mutator transaction binding the contract method 0xae57de0c.
//
// Solidity: function withdraw(uint256 toChainId, address toToken, address recipient, uint256 value, bytes depositHash) returns()
func (_Bridge *BridgeSession) Withdraw(toChainId *big.Int, toToken common.Address, recipient common.Address, value *big.Int, depositHash []byte) (*types.Transaction, error) {
	return _Bridge.Contract.Withdraw(&_Bridge.TransactOpts, toChainId, toToken, recipient, value, depositHash)
}

// Withdraw is a paid mutator transaction binding the contract method 0xae57de0c.
//
// Solidity: function withdraw(uint256 toChainId, address toToken, address recipient, uint256 value, bytes depositHash) returns()
func (_Bridge *BridgeTransactorSession) Withdraw(toChainId *big.Int, toToken common.Address, recipient common.Address, value *big.Int, depositHash []byte) (*types.Transaction, error) {
	return _Bridge.Contract.Withdraw(&_Bridge.TransactOpts, toChainId, toToken, recipient, value, depositHash)
}

// WithdrawNative is a paid mutator transaction binding the contract method 0xf405250f.
//
// Solidity: function withdrawNative(uint256 toChainId, address recipient, bool isMain, uint256 value, bytes depositHash) returns()
func (_Bridge *BridgeTransactor) WithdrawNative(opts *bind.TransactOpts, toChainId *big.Int, recipient common.Address, isMain bool, value *big.Int, depositHash []byte) (*types.Transaction, error) {
	return _Bridge.contract.Transact(opts, "withdrawNative", toChainId, recipient, isMain, value, depositHash)
}

// WithdrawNative is a paid mutator transaction binding the contract method 0xf405250f.
//
// Solidity: function withdrawNative(uint256 toChainId, address recipient, bool isMain, uint256 value, bytes depositHash) returns()
func (_Bridge *BridgeSession) WithdrawNative(toChainId *big.Int, recipient common.Address, isMain bool, value *big.Int, depositHash []byte) (*types.Transaction, error) {
	return _Bridge.Contract.WithdrawNative(&_Bridge.TransactOpts, toChainId, recipient, isMain, value, depositHash)
}

// WithdrawNative is a paid mutator transaction binding the contract method 0xf405250f.
//
// Solidity: function withdrawNative(uint256 toChainId, address recipient, bool isMain, uint256 value, bytes depositHash) returns()
func (_Bridge *BridgeTransactorSession) WithdrawNative(toChainId *big.Int, recipient common.Address, isMain bool, value *big.Int, depositHash []byte) (*types.Transaction, error) {
	return _Bridge.Contract.WithdrawNative(&_Bridge.TransactOpts, toChainId, recipient, isMain, value, depositHash)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Bridge *BridgeTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Bridge.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Bridge *BridgeSession) Receive() (*types.Transaction, error) {
	return _Bridge.Contract.Receive(&_Bridge.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Bridge *BridgeTransactorSession) Receive() (*types.Transaction, error) {
	return _Bridge.Contract.Receive(&_Bridge.TransactOpts)
}

// BridgeDepositIterator is returned from FilterDeposit and is used to iterate over the raw logs and unpacked data for Deposit events raised by the Bridge contract.
type BridgeDepositIterator struct {
	Event *BridgeDeposit // Event containing the contract specifics and raw log

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
func (it *BridgeDepositIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BridgeDeposit)
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
		it.Event = new(BridgeDeposit)
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
func (it *BridgeDepositIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BridgeDepositIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BridgeDeposit represents a Deposit event raised by the Bridge contract.
type BridgeDeposit struct {
	ToChainId *big.Int
	FromToken common.Address
	ToToken   common.Address
	Recipient common.Address
	Value     *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterDeposit is a free log retrieval operation binding the contract event 0x7189490b11cdaec9f21762aedf4a502b07ca019f0f6da42c6da64a60e4377f91.
//
// Solidity: event Deposit(uint256 toChainId, address fromToken, address toToken, address recipient, uint256 value)
func (_Bridge *BridgeFilterer) FilterDeposit(opts *bind.FilterOpts) (*BridgeDepositIterator, error) {

	logs, sub, err := _Bridge.contract.FilterLogs(opts, "Deposit")
	if err != nil {
		return nil, err
	}
	return &BridgeDepositIterator{contract: _Bridge.contract, event: "Deposit", logs: logs, sub: sub}, nil
}

// WatchDeposit is a free log subscription operation binding the contract event 0x7189490b11cdaec9f21762aedf4a502b07ca019f0f6da42c6da64a60e4377f91.
//
// Solidity: event Deposit(uint256 toChainId, address fromToken, address toToken, address recipient, uint256 value)
func (_Bridge *BridgeFilterer) WatchDeposit(opts *bind.WatchOpts, sink chan<- *BridgeDeposit) (event.Subscription, error) {

	logs, sub, err := _Bridge.contract.WatchLogs(opts, "Deposit")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BridgeDeposit)
				if err := _Bridge.contract.UnpackLog(event, "Deposit", log); err != nil {
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

// ParseDeposit is a log parse operation binding the contract event 0x7189490b11cdaec9f21762aedf4a502b07ca019f0f6da42c6da64a60e4377f91.
//
// Solidity: event Deposit(uint256 toChainId, address fromToken, address toToken, address recipient, uint256 value)
func (_Bridge *BridgeFilterer) ParseDeposit(log types.Log) (*BridgeDeposit, error) {
	event := new(BridgeDeposit)
	if err := _Bridge.contract.UnpackLog(event, "Deposit", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BridgeDepositNativeIterator is returned from FilterDepositNative and is used to iterate over the raw logs and unpacked data for DepositNative events raised by the Bridge contract.
type BridgeDepositNativeIterator struct {
	Event *BridgeDepositNative // Event containing the contract specifics and raw log

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
func (it *BridgeDepositNativeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BridgeDepositNative)
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
		it.Event = new(BridgeDepositNative)
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
func (it *BridgeDepositNativeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BridgeDepositNativeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BridgeDepositNative represents a DepositNative event raised by the Bridge contract.
type BridgeDepositNative struct {
	ToChainId *big.Int
	IsMain    bool
	Recipient common.Address
	Value     *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterDepositNative is a free log retrieval operation binding the contract event 0x41167412b560cf27acf5e10143573cc98e0677a05a9802e8add041a252241fc7.
//
// Solidity: event DepositNative(uint256 toChainId, bool isMain, address recipient, uint256 value)
func (_Bridge *BridgeFilterer) FilterDepositNative(opts *bind.FilterOpts) (*BridgeDepositNativeIterator, error) {

	logs, sub, err := _Bridge.contract.FilterLogs(opts, "DepositNative")
	if err != nil {
		return nil, err
	}
	return &BridgeDepositNativeIterator{contract: _Bridge.contract, event: "DepositNative", logs: logs, sub: sub}, nil
}

// WatchDepositNative is a free log subscription operation binding the contract event 0x41167412b560cf27acf5e10143573cc98e0677a05a9802e8add041a252241fc7.
//
// Solidity: event DepositNative(uint256 toChainId, bool isMain, address recipient, uint256 value)
func (_Bridge *BridgeFilterer) WatchDepositNative(opts *bind.WatchOpts, sink chan<- *BridgeDepositNative) (event.Subscription, error) {

	logs, sub, err := _Bridge.contract.WatchLogs(opts, "DepositNative")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BridgeDepositNative)
				if err := _Bridge.contract.UnpackLog(event, "DepositNative", log); err != nil {
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

// ParseDepositNative is a log parse operation binding the contract event 0x41167412b560cf27acf5e10143573cc98e0677a05a9802e8add041a252241fc7.
//
// Solidity: event DepositNative(uint256 toChainId, bool isMain, address recipient, uint256 value)
func (_Bridge *BridgeFilterer) ParseDepositNative(log types.Log) (*BridgeDepositNative, error) {
	event := new(BridgeDepositNative)
	if err := _Bridge.contract.UnpackLog(event, "DepositNative", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BridgeWithdrawDoneIterator is returned from FilterWithdrawDone and is used to iterate over the raw logs and unpacked data for WithdrawDone events raised by the Bridge contract.
type BridgeWithdrawDoneIterator struct {
	Event *BridgeWithdrawDone // Event containing the contract specifics and raw log

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
func (it *BridgeWithdrawDoneIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BridgeWithdrawDone)
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
		it.Event = new(BridgeWithdrawDone)
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
func (it *BridgeWithdrawDoneIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BridgeWithdrawDoneIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BridgeWithdrawDone represents a WithdrawDone event raised by the Bridge contract.
type BridgeWithdrawDone struct {
	ToChainId   *big.Int
	FromToken   common.Address
	ToToken     common.Address
	Recipient   common.Address
	Value       *big.Int
	DepositHash []byte
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterWithdrawDone is a free log retrieval operation binding the contract event 0x1343453926c400b225f44cca53c3e771b9eb74201d8b2b5c8e2c63424ef2373e.
//
// Solidity: event WithdrawDone(uint256 toChainId, address fromToken, address toToken, address recipient, uint256 value, bytes depositHash)
func (_Bridge *BridgeFilterer) FilterWithdrawDone(opts *bind.FilterOpts) (*BridgeWithdrawDoneIterator, error) {

	logs, sub, err := _Bridge.contract.FilterLogs(opts, "WithdrawDone")
	if err != nil {
		return nil, err
	}
	return &BridgeWithdrawDoneIterator{contract: _Bridge.contract, event: "WithdrawDone", logs: logs, sub: sub}, nil
}

// WatchWithdrawDone is a free log subscription operation binding the contract event 0x1343453926c400b225f44cca53c3e771b9eb74201d8b2b5c8e2c63424ef2373e.
//
// Solidity: event WithdrawDone(uint256 toChainId, address fromToken, address toToken, address recipient, uint256 value, bytes depositHash)
func (_Bridge *BridgeFilterer) WatchWithdrawDone(opts *bind.WatchOpts, sink chan<- *BridgeWithdrawDone) (event.Subscription, error) {

	logs, sub, err := _Bridge.contract.WatchLogs(opts, "WithdrawDone")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BridgeWithdrawDone)
				if err := _Bridge.contract.UnpackLog(event, "WithdrawDone", log); err != nil {
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

// ParseWithdrawDone is a log parse operation binding the contract event 0x1343453926c400b225f44cca53c3e771b9eb74201d8b2b5c8e2c63424ef2373e.
//
// Solidity: event WithdrawDone(uint256 toChainId, address fromToken, address toToken, address recipient, uint256 value, bytes depositHash)
func (_Bridge *BridgeFilterer) ParseWithdrawDone(log types.Log) (*BridgeWithdrawDone, error) {
	event := new(BridgeWithdrawDone)
	if err := _Bridge.contract.UnpackLog(event, "WithdrawDone", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BridgeWithdrawNativeDoneIterator is returned from FilterWithdrawNativeDone and is used to iterate over the raw logs and unpacked data for WithdrawNativeDone events raised by the Bridge contract.
type BridgeWithdrawNativeDoneIterator struct {
	Event *BridgeWithdrawNativeDone // Event containing the contract specifics and raw log

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
func (it *BridgeWithdrawNativeDoneIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BridgeWithdrawNativeDone)
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
		it.Event = new(BridgeWithdrawNativeDone)
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
func (it *BridgeWithdrawNativeDoneIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BridgeWithdrawNativeDoneIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BridgeWithdrawNativeDone represents a WithdrawNativeDone event raised by the Bridge contract.
type BridgeWithdrawNativeDone struct {
	FromChainId *big.Int
	Recipient   common.Address
	IsMain      bool
	Value       *big.Int
	DepositHash []byte
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterWithdrawNativeDone is a free log retrieval operation binding the contract event 0xbe024f80850ec9290b6c486221280c8c38d2fb057dc3c0e9566609a1a5714fd4.
//
// Solidity: event WithdrawNativeDone(uint256 fromChainId, address recipient, bool isMain, uint256 value, bytes depositHash)
func (_Bridge *BridgeFilterer) FilterWithdrawNativeDone(opts *bind.FilterOpts) (*BridgeWithdrawNativeDoneIterator, error) {

	logs, sub, err := _Bridge.contract.FilterLogs(opts, "WithdrawNativeDone")
	if err != nil {
		return nil, err
	}
	return &BridgeWithdrawNativeDoneIterator{contract: _Bridge.contract, event: "WithdrawNativeDone", logs: logs, sub: sub}, nil
}

// WatchWithdrawNativeDone is a free log subscription operation binding the contract event 0xbe024f80850ec9290b6c486221280c8c38d2fb057dc3c0e9566609a1a5714fd4.
//
// Solidity: event WithdrawNativeDone(uint256 fromChainId, address recipient, bool isMain, uint256 value, bytes depositHash)
func (_Bridge *BridgeFilterer) WatchWithdrawNativeDone(opts *bind.WatchOpts, sink chan<- *BridgeWithdrawNativeDone) (event.Subscription, error) {

	logs, sub, err := _Bridge.contract.WatchLogs(opts, "WithdrawNativeDone")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BridgeWithdrawNativeDone)
				if err := _Bridge.contract.UnpackLog(event, "WithdrawNativeDone", log); err != nil {
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

// ParseWithdrawNativeDone is a log parse operation binding the contract event 0xbe024f80850ec9290b6c486221280c8c38d2fb057dc3c0e9566609a1a5714fd4.
//
// Solidity: event WithdrawNativeDone(uint256 fromChainId, address recipient, bool isMain, uint256 value, bytes depositHash)
func (_Bridge *BridgeFilterer) ParseWithdrawNativeDone(log types.Log) (*BridgeWithdrawNativeDone, error) {
	event := new(BridgeWithdrawNativeDone)
	if err := _Bridge.contract.UnpackLog(event, "WithdrawNativeDone", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BridgeAdminChangedIterator is returned from FilterAdminChanged and is used to iterate over the raw logs and unpacked data for AdminChanged events raised by the Bridge contract.
type BridgeAdminChangedIterator struct {
	Event *BridgeAdminChanged // Event containing the contract specifics and raw log

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
func (it *BridgeAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BridgeAdminChanged)
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
		it.Event = new(BridgeAdminChanged)
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
func (it *BridgeAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BridgeAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BridgeAdminChanged represents a AdminChanged event raised by the Bridge contract.
type BridgeAdminChanged struct {
	Address common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterAdminChanged is a free log retrieval operation binding the contract event 0xad19681a09902b2e83e20447a2012dbb1bea01d48dd453e74021ebab67e66642.
//
// Solidity: event adminChanged(address _address)
func (_Bridge *BridgeFilterer) FilterAdminChanged(opts *bind.FilterOpts) (*BridgeAdminChangedIterator, error) {

	logs, sub, err := _Bridge.contract.FilterLogs(opts, "adminChanged")
	if err != nil {
		return nil, err
	}
	return &BridgeAdminChangedIterator{contract: _Bridge.contract, event: "adminChanged", logs: logs, sub: sub}, nil
}

// WatchAdminChanged is a free log subscription operation binding the contract event 0xad19681a09902b2e83e20447a2012dbb1bea01d48dd453e74021ebab67e66642.
//
// Solidity: event adminChanged(address _address)
func (_Bridge *BridgeFilterer) WatchAdminChanged(opts *bind.WatchOpts, sink chan<- *BridgeAdminChanged) (event.Subscription, error) {

	logs, sub, err := _Bridge.contract.WatchLogs(opts, "adminChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BridgeAdminChanged)
				if err := _Bridge.contract.UnpackLog(event, "adminChanged", log); err != nil {
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

// ParseAdminChanged is a log parse operation binding the contract event 0xad19681a09902b2e83e20447a2012dbb1bea01d48dd453e74021ebab67e66642.
//
// Solidity: event adminChanged(address _address)
func (_Bridge *BridgeFilterer) ParseAdminChanged(log types.Log) (*BridgeAdminChanged, error) {
	event := new(BridgeAdminChanged)
	if err := _Bridge.contract.UnpackLog(event, "adminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
