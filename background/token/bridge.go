// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package token

import (
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
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// TokenABI is the input ABI used to generate the binding from.
const TokenABI = "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"toChainId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"fromToken\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"toToken\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Deposit\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"toChainId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"isMain\",\"type\":\"bool\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"DepositNative\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"toChainId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"fromToken\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"toToken\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"depositHash\",\"type\":\"bytes\"}],\"name\":\"WithdrawDone\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"fromChainId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"isMain\",\"type\":\"bool\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"depositHash\",\"type\":\"bytes\"}],\"name\":\"WithdrawNativeDone\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_address\",\"type\":\"address\"}],\"name\":\"adminChanged\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"admin\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"toToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"deposit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"toChainId\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"isMain\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"depositNative\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"manager\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"toChainId\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"isRun\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"isMain\",\"type\":\"bool\"},{\"internalType\":\"address\",\"name\":\"fromAddress\",\"type\":\"address\"}],\"name\":\"nativeInsert\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"nativeTransfer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"name\":\"natives\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"isRun\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"isMain\",\"type\":\"bool\"},{\"internalType\":\"address\",\"name\":\"local\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"newAdmin\",\"type\":\"address\"}],\"name\":\"setAdmin\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_manager\",\"type\":\"address\"}],\"name\":\"setManager\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"toChainId\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"isMain\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"state\",\"type\":\"bool\"}],\"name\":\"setNativeIsRun\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"setOwner\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"toChainId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"toToken\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"state\",\"type\":\"bool\"}],\"name\":\"setTokenIsRun\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"toChainId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"toToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"fromToken\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"isRun\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"isMain\",\"type\":\"bool\"}],\"name\":\"tokenInsert\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"fromToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"tokenTransfer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"tokens\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"isRun\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"isMain\",\"type\":\"bool\"},{\"internalType\":\"address\",\"name\":\"local\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"toChainId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"toToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"depositHash\",\"type\":\"bytes\"}],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"toChainId\",\"type\":\"uint256\"},{\"internalType\":\"addresspayable\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"isMain\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"depositHash\",\"type\":\"bytes\"}],\"name\":\"withdrawNative\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]"

// Token is an auto generated Go binding around an Ethereum contract.
type Token struct {
	TokenCaller     // Read-only binding to the contract
	TokenTransactor // Write-only binding to the contract
	TokenFilterer   // Log filterer for contract events
}

// TokenCaller is an auto generated read-only Go binding around an Ethereum contract.
type TokenCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TokenTransactor is an auto generated write-only Go binding around an Ethereum contract.
type TokenTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TokenFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type TokenFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TokenSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type TokenSession struct {
	Contract     *Token            // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// TokenCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type TokenCallerSession struct {
	Contract *TokenCaller  // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// TokenTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type TokenTransactorSession struct {
	Contract     *TokenTransactor  // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// TokenRaw is an auto generated low-level Go binding around an Ethereum contract.
type TokenRaw struct {
	Contract *Token // Generic contract binding to access the raw methods on
}

// TokenCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type TokenCallerRaw struct {
	Contract *TokenCaller // Generic read-only contract binding to access the raw methods on
}

// TokenTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type TokenTransactorRaw struct {
	Contract *TokenTransactor // Generic write-only contract binding to access the raw methods on
}

// NewToken creates a new instance of Token, bound to a specific deployed contract.
func NewToken(address common.Address, backend bind.ContractBackend) (*Token, error) {
	contract, err := bindToken(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Token{TokenCaller: TokenCaller{contract: contract}, TokenTransactor: TokenTransactor{contract: contract}, TokenFilterer: TokenFilterer{contract: contract}}, nil
}

// NewTokenCaller creates a new read-only instance of Token, bound to a specific deployed contract.
func NewTokenCaller(address common.Address, caller bind.ContractCaller) (*TokenCaller, error) {
	contract, err := bindToken(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &TokenCaller{contract: contract}, nil
}

// NewTokenTransactor creates a new write-only instance of Token, bound to a specific deployed contract.
func NewTokenTransactor(address common.Address, transactor bind.ContractTransactor) (*TokenTransactor, error) {
	contract, err := bindToken(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &TokenTransactor{contract: contract}, nil
}

// NewTokenFilterer creates a new log filterer instance of Token, bound to a specific deployed contract.
func NewTokenFilterer(address common.Address, filterer bind.ContractFilterer) (*TokenFilterer, error) {
	contract, err := bindToken(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &TokenFilterer{contract: contract}, nil
}

// bindToken binds a generic wrapper to an already deployed contract.
func bindToken(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(TokenABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Token *TokenRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Token.Contract.TokenCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Token *TokenRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Token.Contract.TokenTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Token *TokenRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Token.Contract.TokenTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Token *TokenCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Token.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Token *TokenTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Token.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Token *TokenTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Token.Contract.contract.Transact(opts, method, params...)
}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() view returns(address)
func (_Token *TokenCaller) Admin(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Token.contract.Call(opts, &out, "admin")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() view returns(address)
func (_Token *TokenSession) Admin() (common.Address, error) {
	return _Token.Contract.Admin(&_Token.CallOpts)
}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() view returns(address)
func (_Token *TokenCallerSession) Admin() (common.Address, error) {
	return _Token.Contract.Admin(&_Token.CallOpts)
}

// Manager is a free data retrieval call binding the contract method 0x481c6a75.
//
// Solidity: function manager() view returns(address)
func (_Token *TokenCaller) Manager(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Token.contract.Call(opts, &out, "manager")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Manager is a free data retrieval call binding the contract method 0x481c6a75.
//
// Solidity: function manager() view returns(address)
func (_Token *TokenSession) Manager() (common.Address, error) {
	return _Token.Contract.Manager(&_Token.CallOpts)
}

// Manager is a free data retrieval call binding the contract method 0x481c6a75.
//
// Solidity: function manager() view returns(address)
func (_Token *TokenCallerSession) Manager() (common.Address, error) {
	return _Token.Contract.Manager(&_Token.CallOpts)
}

// Natives is a free data retrieval call binding the contract method 0x7aabe0a6.
//
// Solidity: function natives(uint256 , bool ) view returns(bool isRun, bool isMain, address local)
func (_Token *TokenCaller) Natives(opts *bind.CallOpts, arg0 *big.Int, arg1 bool) (struct {
	IsRun  bool
	IsMain bool
	Local  common.Address
}, error) {
	var out []interface{}
	err := _Token.contract.Call(opts, &out, "natives", arg0, arg1)

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
func (_Token *TokenSession) Natives(arg0 *big.Int, arg1 bool) (struct {
	IsRun  bool
	IsMain bool
	Local  common.Address
}, error) {
	return _Token.Contract.Natives(&_Token.CallOpts, arg0, arg1)
}

// Natives is a free data retrieval call binding the contract method 0x7aabe0a6.
//
// Solidity: function natives(uint256 , bool ) view returns(bool isRun, bool isMain, address local)
func (_Token *TokenCallerSession) Natives(arg0 *big.Int, arg1 bool) (struct {
	IsRun  bool
	IsMain bool
	Local  common.Address
}, error) {
	return _Token.Contract.Natives(&_Token.CallOpts, arg0, arg1)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Token *TokenCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Token.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Token *TokenSession) Owner() (common.Address, error) {
	return _Token.Contract.Owner(&_Token.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Token *TokenCallerSession) Owner() (common.Address, error) {
	return _Token.Contract.Owner(&_Token.CallOpts)
}

// Tokens is a free data retrieval call binding the contract method 0xae68a2ba.
//
// Solidity: function tokens(uint256 , address ) view returns(bool isRun, bool isMain, address local)
func (_Token *TokenCaller) Tokens(opts *bind.CallOpts, arg0 *big.Int, arg1 common.Address) (struct {
	IsRun  bool
	IsMain bool
	Local  common.Address
}, error) {
	var out []interface{}
	err := _Token.contract.Call(opts, &out, "tokens", arg0, arg1)

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
func (_Token *TokenSession) Tokens(arg0 *big.Int, arg1 common.Address) (struct {
	IsRun  bool
	IsMain bool
	Local  common.Address
}, error) {
	return _Token.Contract.Tokens(&_Token.CallOpts, arg0, arg1)
}

// Tokens is a free data retrieval call binding the contract method 0xae68a2ba.
//
// Solidity: function tokens(uint256 , address ) view returns(bool isRun, bool isMain, address local)
func (_Token *TokenCallerSession) Tokens(arg0 *big.Int, arg1 common.Address) (struct {
	IsRun  bool
	IsMain bool
	Local  common.Address
}, error) {
	return _Token.Contract.Tokens(&_Token.CallOpts, arg0, arg1)
}

// Deposit is a paid mutator transaction binding the contract method 0xbc157ac1.
//
// Solidity: function deposit(uint256 chainId, address toToken, uint256 value) returns()
func (_Token *TokenTransactor) Deposit(opts *bind.TransactOpts, chainId *big.Int, toToken common.Address, value *big.Int) (*types.Transaction, error) {
	return _Token.contract.Transact(opts, "deposit", chainId, toToken, value)
}

// Deposit is a paid mutator transaction binding the contract method 0xbc157ac1.
//
// Solidity: function deposit(uint256 chainId, address toToken, uint256 value) returns()
func (_Token *TokenSession) Deposit(chainId *big.Int, toToken common.Address, value *big.Int) (*types.Transaction, error) {
	return _Token.Contract.Deposit(&_Token.TransactOpts, chainId, toToken, value)
}

// Deposit is a paid mutator transaction binding the contract method 0xbc157ac1.
//
// Solidity: function deposit(uint256 chainId, address toToken, uint256 value) returns()
func (_Token *TokenTransactorSession) Deposit(chainId *big.Int, toToken common.Address, value *big.Int) (*types.Transaction, error) {
	return _Token.Contract.Deposit(&_Token.TransactOpts, chainId, toToken, value)
}

// DepositNative is a paid mutator transaction binding the contract method 0x02367ad2.
//
// Solidity: function depositNative(uint256 toChainId, bool isMain, uint256 value) payable returns()
func (_Token *TokenTransactor) DepositNative(opts *bind.TransactOpts, toChainId *big.Int, isMain bool, value *big.Int) (*types.Transaction, error) {
	return _Token.contract.Transact(opts, "depositNative", toChainId, isMain, value)
}

// DepositNative is a paid mutator transaction binding the contract method 0x02367ad2.
//
// Solidity: function depositNative(uint256 toChainId, bool isMain, uint256 value) payable returns()
func (_Token *TokenSession) DepositNative(toChainId *big.Int, isMain bool, value *big.Int) (*types.Transaction, error) {
	return _Token.Contract.DepositNative(&_Token.TransactOpts, toChainId, isMain, value)
}

// DepositNative is a paid mutator transaction binding the contract method 0x02367ad2.
//
// Solidity: function depositNative(uint256 toChainId, bool isMain, uint256 value) payable returns()
func (_Token *TokenTransactorSession) DepositNative(toChainId *big.Int, isMain bool, value *big.Int) (*types.Transaction, error) {
	return _Token.Contract.DepositNative(&_Token.TransactOpts, toChainId, isMain, value)
}

// NativeInsert is a paid mutator transaction binding the contract method 0xffb4aedc.
//
// Solidity: function nativeInsert(uint256 toChainId, bool isRun, bool isMain, address fromAddress) returns()
func (_Token *TokenTransactor) NativeInsert(opts *bind.TransactOpts, toChainId *big.Int, isRun bool, isMain bool, fromAddress common.Address) (*types.Transaction, error) {
	return _Token.contract.Transact(opts, "nativeInsert", toChainId, isRun, isMain, fromAddress)
}

// NativeInsert is a paid mutator transaction binding the contract method 0xffb4aedc.
//
// Solidity: function nativeInsert(uint256 toChainId, bool isRun, bool isMain, address fromAddress) returns()
func (_Token *TokenSession) NativeInsert(toChainId *big.Int, isRun bool, isMain bool, fromAddress common.Address) (*types.Transaction, error) {
	return _Token.Contract.NativeInsert(&_Token.TransactOpts, toChainId, isRun, isMain, fromAddress)
}

// NativeInsert is a paid mutator transaction binding the contract method 0xffb4aedc.
//
// Solidity: function nativeInsert(uint256 toChainId, bool isRun, bool isMain, address fromAddress) returns()
func (_Token *TokenTransactorSession) NativeInsert(toChainId *big.Int, isRun bool, isMain bool, fromAddress common.Address) (*types.Transaction, error) {
	return _Token.Contract.NativeInsert(&_Token.TransactOpts, toChainId, isRun, isMain, fromAddress)
}

// NativeTransfer is a paid mutator transaction binding the contract method 0xf2522bcd.
//
// Solidity: function nativeTransfer(address recipient, uint256 value) returns()
func (_Token *TokenTransactor) NativeTransfer(opts *bind.TransactOpts, recipient common.Address, value *big.Int) (*types.Transaction, error) {
	return _Token.contract.Transact(opts, "nativeTransfer", recipient, value)
}

// NativeTransfer is a paid mutator transaction binding the contract method 0xf2522bcd.
//
// Solidity: function nativeTransfer(address recipient, uint256 value) returns()
func (_Token *TokenSession) NativeTransfer(recipient common.Address, value *big.Int) (*types.Transaction, error) {
	return _Token.Contract.NativeTransfer(&_Token.TransactOpts, recipient, value)
}

// NativeTransfer is a paid mutator transaction binding the contract method 0xf2522bcd.
//
// Solidity: function nativeTransfer(address recipient, uint256 value) returns()
func (_Token *TokenTransactorSession) NativeTransfer(recipient common.Address, value *big.Int) (*types.Transaction, error) {
	return _Token.Contract.NativeTransfer(&_Token.TransactOpts, recipient, value)
}

// SetAdmin is a paid mutator transaction binding the contract method 0x704b6c02.
//
// Solidity: function setAdmin(address newAdmin) returns()
func (_Token *TokenTransactor) SetAdmin(opts *bind.TransactOpts, newAdmin common.Address) (*types.Transaction, error) {
	return _Token.contract.Transact(opts, "setAdmin", newAdmin)
}

// SetAdmin is a paid mutator transaction binding the contract method 0x704b6c02.
//
// Solidity: function setAdmin(address newAdmin) returns()
func (_Token *TokenSession) SetAdmin(newAdmin common.Address) (*types.Transaction, error) {
	return _Token.Contract.SetAdmin(&_Token.TransactOpts, newAdmin)
}

// SetAdmin is a paid mutator transaction binding the contract method 0x704b6c02.
//
// Solidity: function setAdmin(address newAdmin) returns()
func (_Token *TokenTransactorSession) SetAdmin(newAdmin common.Address) (*types.Transaction, error) {
	return _Token.Contract.SetAdmin(&_Token.TransactOpts, newAdmin)
}

// SetManager is a paid mutator transaction binding the contract method 0xd0ebdbe7.
//
// Solidity: function setManager(address _manager) returns()
func (_Token *TokenTransactor) SetManager(opts *bind.TransactOpts, _manager common.Address) (*types.Transaction, error) {
	return _Token.contract.Transact(opts, "setManager", _manager)
}

// SetManager is a paid mutator transaction binding the contract method 0xd0ebdbe7.
//
// Solidity: function setManager(address _manager) returns()
func (_Token *TokenSession) SetManager(_manager common.Address) (*types.Transaction, error) {
	return _Token.Contract.SetManager(&_Token.TransactOpts, _manager)
}

// SetManager is a paid mutator transaction binding the contract method 0xd0ebdbe7.
//
// Solidity: function setManager(address _manager) returns()
func (_Token *TokenTransactorSession) SetManager(_manager common.Address) (*types.Transaction, error) {
	return _Token.Contract.SetManager(&_Token.TransactOpts, _manager)
}

// SetNativeIsRun is a paid mutator transaction binding the contract method 0x9a234732.
//
// Solidity: function setNativeIsRun(uint256 toChainId, bool isMain, bool state) returns()
func (_Token *TokenTransactor) SetNativeIsRun(opts *bind.TransactOpts, toChainId *big.Int, isMain bool, state bool) (*types.Transaction, error) {
	return _Token.contract.Transact(opts, "setNativeIsRun", toChainId, isMain, state)
}

// SetNativeIsRun is a paid mutator transaction binding the contract method 0x9a234732.
//
// Solidity: function setNativeIsRun(uint256 toChainId, bool isMain, bool state) returns()
func (_Token *TokenSession) SetNativeIsRun(toChainId *big.Int, isMain bool, state bool) (*types.Transaction, error) {
	return _Token.Contract.SetNativeIsRun(&_Token.TransactOpts, toChainId, isMain, state)
}

// SetNativeIsRun is a paid mutator transaction binding the contract method 0x9a234732.
//
// Solidity: function setNativeIsRun(uint256 toChainId, bool isMain, bool state) returns()
func (_Token *TokenTransactorSession) SetNativeIsRun(toChainId *big.Int, isMain bool, state bool) (*types.Transaction, error) {
	return _Token.Contract.SetNativeIsRun(&_Token.TransactOpts, toChainId, isMain, state)
}

// SetOwner is a paid mutator transaction binding the contract method 0x13af4035.
//
// Solidity: function setOwner(address newOwner) returns()
func (_Token *TokenTransactor) SetOwner(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Token.contract.Transact(opts, "setOwner", newOwner)
}

// SetOwner is a paid mutator transaction binding the contract method 0x13af4035.
//
// Solidity: function setOwner(address newOwner) returns()
func (_Token *TokenSession) SetOwner(newOwner common.Address) (*types.Transaction, error) {
	return _Token.Contract.SetOwner(&_Token.TransactOpts, newOwner)
}

// SetOwner is a paid mutator transaction binding the contract method 0x13af4035.
//
// Solidity: function setOwner(address newOwner) returns()
func (_Token *TokenTransactorSession) SetOwner(newOwner common.Address) (*types.Transaction, error) {
	return _Token.Contract.SetOwner(&_Token.TransactOpts, newOwner)
}

// SetTokenIsRun is a paid mutator transaction binding the contract method 0x2eef73b3.
//
// Solidity: function setTokenIsRun(uint256 toChainId, address toToken, bool state) returns()
func (_Token *TokenTransactor) SetTokenIsRun(opts *bind.TransactOpts, toChainId *big.Int, toToken common.Address, state bool) (*types.Transaction, error) {
	return _Token.contract.Transact(opts, "setTokenIsRun", toChainId, toToken, state)
}

// SetTokenIsRun is a paid mutator transaction binding the contract method 0x2eef73b3.
//
// Solidity: function setTokenIsRun(uint256 toChainId, address toToken, bool state) returns()
func (_Token *TokenSession) SetTokenIsRun(toChainId *big.Int, toToken common.Address, state bool) (*types.Transaction, error) {
	return _Token.Contract.SetTokenIsRun(&_Token.TransactOpts, toChainId, toToken, state)
}

// SetTokenIsRun is a paid mutator transaction binding the contract method 0x2eef73b3.
//
// Solidity: function setTokenIsRun(uint256 toChainId, address toToken, bool state) returns()
func (_Token *TokenTransactorSession) SetTokenIsRun(toChainId *big.Int, toToken common.Address, state bool) (*types.Transaction, error) {
	return _Token.Contract.SetTokenIsRun(&_Token.TransactOpts, toChainId, toToken, state)
}

// TokenInsert is a paid mutator transaction binding the contract method 0xd523a97e.
//
// Solidity: function tokenInsert(uint256 toChainId, address toToken, address fromToken, bool isRun, bool isMain) returns()
func (_Token *TokenTransactor) TokenInsert(opts *bind.TransactOpts, toChainId *big.Int, toToken common.Address, fromToken common.Address, isRun bool, isMain bool) (*types.Transaction, error) {
	return _Token.contract.Transact(opts, "tokenInsert", toChainId, toToken, fromToken, isRun, isMain)
}

// TokenInsert is a paid mutator transaction binding the contract method 0xd523a97e.
//
// Solidity: function tokenInsert(uint256 toChainId, address toToken, address fromToken, bool isRun, bool isMain) returns()
func (_Token *TokenSession) TokenInsert(toChainId *big.Int, toToken common.Address, fromToken common.Address, isRun bool, isMain bool) (*types.Transaction, error) {
	return _Token.Contract.TokenInsert(&_Token.TransactOpts, toChainId, toToken, fromToken, isRun, isMain)
}

// TokenInsert is a paid mutator transaction binding the contract method 0xd523a97e.
//
// Solidity: function tokenInsert(uint256 toChainId, address toToken, address fromToken, bool isRun, bool isMain) returns()
func (_Token *TokenTransactorSession) TokenInsert(toChainId *big.Int, toToken common.Address, fromToken common.Address, isRun bool, isMain bool) (*types.Transaction, error) {
	return _Token.Contract.TokenInsert(&_Token.TransactOpts, toChainId, toToken, fromToken, isRun, isMain)
}

// TokenTransfer is a paid mutator transaction binding the contract method 0x15eaef6b.
//
// Solidity: function tokenTransfer(address fromToken, address recipient, uint256 value) returns()
func (_Token *TokenTransactor) TokenTransfer(opts *bind.TransactOpts, fromToken common.Address, recipient common.Address, value *big.Int) (*types.Transaction, error) {
	return _Token.contract.Transact(opts, "tokenTransfer", fromToken, recipient, value)
}

// TokenTransfer is a paid mutator transaction binding the contract method 0x15eaef6b.
//
// Solidity: function tokenTransfer(address fromToken, address recipient, uint256 value) returns()
func (_Token *TokenSession) TokenTransfer(fromToken common.Address, recipient common.Address, value *big.Int) (*types.Transaction, error) {
	return _Token.Contract.TokenTransfer(&_Token.TransactOpts, fromToken, recipient, value)
}

// TokenTransfer is a paid mutator transaction binding the contract method 0x15eaef6b.
//
// Solidity: function tokenTransfer(address fromToken, address recipient, uint256 value) returns()
func (_Token *TokenTransactorSession) TokenTransfer(fromToken common.Address, recipient common.Address, value *big.Int) (*types.Transaction, error) {
	return _Token.Contract.TokenTransfer(&_Token.TransactOpts, fromToken, recipient, value)
}

// Withdraw is a paid mutator transaction binding the contract method 0xae57de0c.
//
// Solidity: function withdraw(uint256 toChainId, address toToken, address recipient, uint256 value, bytes depositHash) returns()
func (_Token *TokenTransactor) Withdraw(opts *bind.TransactOpts, toChainId *big.Int, toToken common.Address, recipient common.Address, value *big.Int, depositHash []byte) (*types.Transaction, error) {
	return _Token.contract.Transact(opts, "withdraw", toChainId, toToken, recipient, value, depositHash)
}

// Withdraw is a paid mutator transaction binding the contract method 0xae57de0c.
//
// Solidity: function withdraw(uint256 toChainId, address toToken, address recipient, uint256 value, bytes depositHash) returns()
func (_Token *TokenSession) Withdraw(toChainId *big.Int, toToken common.Address, recipient common.Address, value *big.Int, depositHash []byte) (*types.Transaction, error) {
	return _Token.Contract.Withdraw(&_Token.TransactOpts, toChainId, toToken, recipient, value, depositHash)
}

// Withdraw is a paid mutator transaction binding the contract method 0xae57de0c.
//
// Solidity: function withdraw(uint256 toChainId, address toToken, address recipient, uint256 value, bytes depositHash) returns()
func (_Token *TokenTransactorSession) Withdraw(toChainId *big.Int, toToken common.Address, recipient common.Address, value *big.Int, depositHash []byte) (*types.Transaction, error) {
	return _Token.Contract.Withdraw(&_Token.TransactOpts, toChainId, toToken, recipient, value, depositHash)
}

// WithdrawNative is a paid mutator transaction binding the contract method 0xf405250f.
//
// Solidity: function withdrawNative(uint256 toChainId, address recipient, bool isMain, uint256 value, bytes depositHash) returns()
func (_Token *TokenTransactor) WithdrawNative(opts *bind.TransactOpts, toChainId *big.Int, recipient common.Address, isMain bool, value *big.Int, depositHash []byte) (*types.Transaction, error) {
	return _Token.contract.Transact(opts, "withdrawNative", toChainId, recipient, isMain, value, depositHash)
}

// WithdrawNative is a paid mutator transaction binding the contract method 0xf405250f.
//
// Solidity: function withdrawNative(uint256 toChainId, address recipient, bool isMain, uint256 value, bytes depositHash) returns()
func (_Token *TokenSession) WithdrawNative(toChainId *big.Int, recipient common.Address, isMain bool, value *big.Int, depositHash []byte) (*types.Transaction, error) {
	return _Token.Contract.WithdrawNative(&_Token.TransactOpts, toChainId, recipient, isMain, value, depositHash)
}

// WithdrawNative is a paid mutator transaction binding the contract method 0xf405250f.
//
// Solidity: function withdrawNative(uint256 toChainId, address recipient, bool isMain, uint256 value, bytes depositHash) returns()
func (_Token *TokenTransactorSession) WithdrawNative(toChainId *big.Int, recipient common.Address, isMain bool, value *big.Int, depositHash []byte) (*types.Transaction, error) {
	return _Token.Contract.WithdrawNative(&_Token.TransactOpts, toChainId, recipient, isMain, value, depositHash)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Token *TokenTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Token.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Token *TokenSession) Receive() (*types.Transaction, error) {
	return _Token.Contract.Receive(&_Token.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Token *TokenTransactorSession) Receive() (*types.Transaction, error) {
	return _Token.Contract.Receive(&_Token.TransactOpts)
}

// TokenDepositIterator is returned from FilterDeposit and is used to iterate over the raw logs and unpacked data for Deposit events raised by the Token contract.
type TokenDepositIterator struct {
	Event *TokenDeposit // Event containing the contract specifics and raw log

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
func (it *TokenDepositIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TokenDeposit)
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
		it.Event = new(TokenDeposit)
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
func (it *TokenDepositIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TokenDepositIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TokenDeposit represents a Deposit event raised by the Token contract.
type TokenDeposit struct {
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
func (_Token *TokenFilterer) FilterDeposit(opts *bind.FilterOpts) (*TokenDepositIterator, error) {

	logs, sub, err := _Token.contract.FilterLogs(opts, "Deposit")
	if err != nil {
		return nil, err
	}
	return &TokenDepositIterator{contract: _Token.contract, event: "Deposit", logs: logs, sub: sub}, nil
}

// WatchDeposit is a free log subscription operation binding the contract event 0x7189490b11cdaec9f21762aedf4a502b07ca019f0f6da42c6da64a60e4377f91.
//
// Solidity: event Deposit(uint256 toChainId, address fromToken, address toToken, address recipient, uint256 value)
func (_Token *TokenFilterer) WatchDeposit(opts *bind.WatchOpts, sink chan<- *TokenDeposit) (event.Subscription, error) {

	logs, sub, err := _Token.contract.WatchLogs(opts, "Deposit")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TokenDeposit)
				if err := _Token.contract.UnpackLog(event, "Deposit", log); err != nil {
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
func (_Token *TokenFilterer) ParseDeposit(log types.Log) (*TokenDeposit, error) {
	event := new(TokenDeposit)
	if err := _Token.contract.UnpackLog(event, "Deposit", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TokenDepositNativeIterator is returned from FilterDepositNative and is used to iterate over the raw logs and unpacked data for DepositNative events raised by the Token contract.
type TokenDepositNativeIterator struct {
	Event *TokenDepositNative // Event containing the contract specifics and raw log

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
func (it *TokenDepositNativeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TokenDepositNative)
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
		it.Event = new(TokenDepositNative)
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
func (it *TokenDepositNativeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TokenDepositNativeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TokenDepositNative represents a DepositNative event raised by the Token contract.
type TokenDepositNative struct {
	ToChainId *big.Int
	IsMain    bool
	Recipient common.Address
	Value     *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterDepositNative is a free log retrieval operation binding the contract event 0x41167412b560cf27acf5e10143573cc98e0677a05a9802e8add041a252241fc7.
//
// Solidity: event DepositNative(uint256 toChainId, bool isMain, address recipient, uint256 value)
func (_Token *TokenFilterer) FilterDepositNative(opts *bind.FilterOpts) (*TokenDepositNativeIterator, error) {

	logs, sub, err := _Token.contract.FilterLogs(opts, "DepositNative")
	if err != nil {
		return nil, err
	}
	return &TokenDepositNativeIterator{contract: _Token.contract, event: "DepositNative", logs: logs, sub: sub}, nil
}

// WatchDepositNative is a free log subscription operation binding the contract event 0x41167412b560cf27acf5e10143573cc98e0677a05a9802e8add041a252241fc7.
//
// Solidity: event DepositNative(uint256 toChainId, bool isMain, address recipient, uint256 value)
func (_Token *TokenFilterer) WatchDepositNative(opts *bind.WatchOpts, sink chan<- *TokenDepositNative) (event.Subscription, error) {

	logs, sub, err := _Token.contract.WatchLogs(opts, "DepositNative")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TokenDepositNative)
				if err := _Token.contract.UnpackLog(event, "DepositNative", log); err != nil {
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
func (_Token *TokenFilterer) ParseDepositNative(log types.Log) (*TokenDepositNative, error) {
	event := new(TokenDepositNative)
	if err := _Token.contract.UnpackLog(event, "DepositNative", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TokenWithdrawDoneIterator is returned from FilterWithdrawDone and is used to iterate over the raw logs and unpacked data for WithdrawDone events raised by the Token contract.
type TokenWithdrawDoneIterator struct {
	Event *TokenWithdrawDone // Event containing the contract specifics and raw log

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
func (it *TokenWithdrawDoneIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TokenWithdrawDone)
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
		it.Event = new(TokenWithdrawDone)
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
func (it *TokenWithdrawDoneIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TokenWithdrawDoneIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TokenWithdrawDone represents a WithdrawDone event raised by the Token contract.
type TokenWithdrawDone struct {
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
func (_Token *TokenFilterer) FilterWithdrawDone(opts *bind.FilterOpts) (*TokenWithdrawDoneIterator, error) {

	logs, sub, err := _Token.contract.FilterLogs(opts, "WithdrawDone")
	if err != nil {
		return nil, err
	}
	return &TokenWithdrawDoneIterator{contract: _Token.contract, event: "WithdrawDone", logs: logs, sub: sub}, nil
}

// WatchWithdrawDone is a free log subscription operation binding the contract event 0x1343453926c400b225f44cca53c3e771b9eb74201d8b2b5c8e2c63424ef2373e.
//
// Solidity: event WithdrawDone(uint256 toChainId, address fromToken, address toToken, address recipient, uint256 value, bytes depositHash)
func (_Token *TokenFilterer) WatchWithdrawDone(opts *bind.WatchOpts, sink chan<- *TokenWithdrawDone) (event.Subscription, error) {

	logs, sub, err := _Token.contract.WatchLogs(opts, "WithdrawDone")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TokenWithdrawDone)
				if err := _Token.contract.UnpackLog(event, "WithdrawDone", log); err != nil {
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
func (_Token *TokenFilterer) ParseWithdrawDone(log types.Log) (*TokenWithdrawDone, error) {
	event := new(TokenWithdrawDone)
	if err := _Token.contract.UnpackLog(event, "WithdrawDone", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TokenWithdrawNativeDoneIterator is returned from FilterWithdrawNativeDone and is used to iterate over the raw logs and unpacked data for WithdrawNativeDone events raised by the Token contract.
type TokenWithdrawNativeDoneIterator struct {
	Event *TokenWithdrawNativeDone // Event containing the contract specifics and raw log

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
func (it *TokenWithdrawNativeDoneIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TokenWithdrawNativeDone)
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
		it.Event = new(TokenWithdrawNativeDone)
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
func (it *TokenWithdrawNativeDoneIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TokenWithdrawNativeDoneIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TokenWithdrawNativeDone represents a WithdrawNativeDone event raised by the Token contract.
type TokenWithdrawNativeDone struct {
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
func (_Token *TokenFilterer) FilterWithdrawNativeDone(opts *bind.FilterOpts) (*TokenWithdrawNativeDoneIterator, error) {

	logs, sub, err := _Token.contract.FilterLogs(opts, "WithdrawNativeDone")
	if err != nil {
		return nil, err
	}
	return &TokenWithdrawNativeDoneIterator{contract: _Token.contract, event: "WithdrawNativeDone", logs: logs, sub: sub}, nil
}

// WatchWithdrawNativeDone is a free log subscription operation binding the contract event 0xbe024f80850ec9290b6c486221280c8c38d2fb057dc3c0e9566609a1a5714fd4.
//
// Solidity: event WithdrawNativeDone(uint256 fromChainId, address recipient, bool isMain, uint256 value, bytes depositHash)
func (_Token *TokenFilterer) WatchWithdrawNativeDone(opts *bind.WatchOpts, sink chan<- *TokenWithdrawNativeDone) (event.Subscription, error) {

	logs, sub, err := _Token.contract.WatchLogs(opts, "WithdrawNativeDone")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TokenWithdrawNativeDone)
				if err := _Token.contract.UnpackLog(event, "WithdrawNativeDone", log); err != nil {
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
func (_Token *TokenFilterer) ParseWithdrawNativeDone(log types.Log) (*TokenWithdrawNativeDone, error) {
	event := new(TokenWithdrawNativeDone)
	if err := _Token.contract.UnpackLog(event, "WithdrawNativeDone", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TokenAdminChangedIterator is returned from FilterAdminChanged and is used to iterate over the raw logs and unpacked data for AdminChanged events raised by the Token contract.
type TokenAdminChangedIterator struct {
	Event *TokenAdminChanged // Event containing the contract specifics and raw log

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
func (it *TokenAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TokenAdminChanged)
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
		it.Event = new(TokenAdminChanged)
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
func (it *TokenAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TokenAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TokenAdminChanged represents a AdminChanged event raised by the Token contract.
type TokenAdminChanged struct {
	Address common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterAdminChanged is a free log retrieval operation binding the contract event 0xad19681a09902b2e83e20447a2012dbb1bea01d48dd453e74021ebab67e66642.
//
// Solidity: event adminChanged(address _address)
func (_Token *TokenFilterer) FilterAdminChanged(opts *bind.FilterOpts) (*TokenAdminChangedIterator, error) {

	logs, sub, err := _Token.contract.FilterLogs(opts, "adminChanged")
	if err != nil {
		return nil, err
	}
	return &TokenAdminChangedIterator{contract: _Token.contract, event: "adminChanged", logs: logs, sub: sub}, nil
}

// WatchAdminChanged is a free log subscription operation binding the contract event 0xad19681a09902b2e83e20447a2012dbb1bea01d48dd453e74021ebab67e66642.
//
// Solidity: event adminChanged(address _address)
func (_Token *TokenFilterer) WatchAdminChanged(opts *bind.WatchOpts, sink chan<- *TokenAdminChanged) (event.Subscription, error) {

	logs, sub, err := _Token.contract.WatchLogs(opts, "adminChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TokenAdminChanged)
				if err := _Token.contract.UnpackLog(event, "adminChanged", log); err != nil {
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
func (_Token *TokenFilterer) ParseAdminChanged(log types.Log) (*TokenAdminChanged, error) {
	event := new(TokenAdminChanged)
	if err := _Token.contract.UnpackLog(event, "adminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
