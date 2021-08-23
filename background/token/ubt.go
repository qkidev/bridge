// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package token

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

// UbtABI is the input ABI used to generate the binding from.
// Deprecated: Use UbtMetaData.ABI instead.
const UbtABI = "[{\"inputs\":[{\"internalType\":\"string\",\"name\":\"tokenName\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"tokenSymbol\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"epoch_time\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_burn_price\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_decimal_usdt\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Burn\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Freeze\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Minted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"CoinBalanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"TokenBalanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"auto_update_burn_price\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"burn\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"burn_price\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"burn_swap_address\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decimal_bt\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decimal_usdt\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"deprecated\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"epoch_base\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"getSwapPrice\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"invite\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"inviteCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"last_miner\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"mint\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"mint_reward\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"usePower\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"addresspayable\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"power\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"invite_address\",\"type\":\"address\"}],\"name\":\"registration\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"rewardCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_oracle_address\",\"type\":\"address\"}],\"name\":\"setOracle\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"new_owner\",\"type\":\"address\"}],\"name\":\"setOwner\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_swap_address\",\"type\":\"address\"}],\"name\":\"set_swap_address\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalPower\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalUsersAmount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"transfer_fee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"fee\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"update_burn_price\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"upgradedAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"users_epoch\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"users_start_time\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]"

// Ubt is an auto generated Go binding around an Ethereum contract.
type Ubt struct {
	UbtCaller     // Read-only binding to the contract
	UbtTransactor // Write-only binding to the contract
	UbtFilterer   // Log filterer for contract events
}

// UbtCaller is an auto generated read-only Go binding around an Ethereum contract.
type UbtCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// UbtTransactor is an auto generated write-only Go binding around an Ethereum contract.
type UbtTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// UbtFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type UbtFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// UbtSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type UbtSession struct {
	Contract     *Ubt              // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// UbtCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type UbtCallerSession struct {
	Contract *UbtCaller    // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// UbtTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type UbtTransactorSession struct {
	Contract     *UbtTransactor    // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// UbtRaw is an auto generated low-level Go binding around an Ethereum contract.
type UbtRaw struct {
	Contract *Ubt // Generic contract binding to access the raw methods on
}

// UbtCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type UbtCallerRaw struct {
	Contract *UbtCaller // Generic read-only contract binding to access the raw methods on
}

// UbtTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type UbtTransactorRaw struct {
	Contract *UbtTransactor // Generic write-only contract binding to access the raw methods on
}

// NewUbt creates a new instance of Ubt, bound to a specific deployed contract.
func NewUbt(address common.Address, backend bind.ContractBackend) (*Ubt, error) {
	contract, err := bindUbt(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Ubt{UbtCaller: UbtCaller{contract: contract}, UbtTransactor: UbtTransactor{contract: contract}, UbtFilterer: UbtFilterer{contract: contract}}, nil
}

// NewUbtCaller creates a new read-only instance of Ubt, bound to a specific deployed contract.
func NewUbtCaller(address common.Address, caller bind.ContractCaller) (*UbtCaller, error) {
	contract, err := bindUbt(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &UbtCaller{contract: contract}, nil
}

// NewUbtTransactor creates a new write-only instance of Ubt, bound to a specific deployed contract.
func NewUbtTransactor(address common.Address, transactor bind.ContractTransactor) (*UbtTransactor, error) {
	contract, err := bindUbt(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &UbtTransactor{contract: contract}, nil
}

// NewUbtFilterer creates a new log filterer instance of Ubt, bound to a specific deployed contract.
func NewUbtFilterer(address common.Address, filterer bind.ContractFilterer) (*UbtFilterer, error) {
	contract, err := bindUbt(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &UbtFilterer{contract: contract}, nil
}

// bindUbt binds a generic wrapper to an already deployed contract.
func bindUbt(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(UbtABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Ubt *UbtRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Ubt.Contract.UbtCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Ubt *UbtRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Ubt.Contract.UbtTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Ubt *UbtRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Ubt.Contract.UbtTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Ubt *UbtCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Ubt.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Ubt *UbtTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Ubt.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Ubt *UbtTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Ubt.Contract.contract.Transact(opts, method, params...)
}

// CoinBalanceOf is a free data retrieval call binding the contract method 0xb370d4e0.
//
// Solidity: function CoinBalanceOf(address ) view returns(uint256)
func (_Ubt *UbtCaller) CoinBalanceOf(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Ubt.contract.Call(opts, &out, "CoinBalanceOf", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CoinBalanceOf is a free data retrieval call binding the contract method 0xb370d4e0.
//
// Solidity: function CoinBalanceOf(address ) view returns(uint256)
func (_Ubt *UbtSession) CoinBalanceOf(arg0 common.Address) (*big.Int, error) {
	return _Ubt.Contract.CoinBalanceOf(&_Ubt.CallOpts, arg0)
}

// CoinBalanceOf is a free data retrieval call binding the contract method 0xb370d4e0.
//
// Solidity: function CoinBalanceOf(address ) view returns(uint256)
func (_Ubt *UbtCallerSession) CoinBalanceOf(arg0 common.Address) (*big.Int, error) {
	return _Ubt.Contract.CoinBalanceOf(&_Ubt.CallOpts, arg0)
}

// TokenBalanceOf is a free data retrieval call binding the contract method 0x2f2617cc.
//
// Solidity: function TokenBalanceOf(address , address ) view returns(uint256)
func (_Ubt *UbtCaller) TokenBalanceOf(opts *bind.CallOpts, arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Ubt.contract.Call(opts, &out, "TokenBalanceOf", arg0, arg1)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TokenBalanceOf is a free data retrieval call binding the contract method 0x2f2617cc.
//
// Solidity: function TokenBalanceOf(address , address ) view returns(uint256)
func (_Ubt *UbtSession) TokenBalanceOf(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	return _Ubt.Contract.TokenBalanceOf(&_Ubt.CallOpts, arg0, arg1)
}

// TokenBalanceOf is a free data retrieval call binding the contract method 0x2f2617cc.
//
// Solidity: function TokenBalanceOf(address , address ) view returns(uint256)
func (_Ubt *UbtCallerSession) TokenBalanceOf(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	return _Ubt.Contract.TokenBalanceOf(&_Ubt.CallOpts, arg0, arg1)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address , address ) view returns(uint256)
func (_Ubt *UbtCaller) Allowance(opts *bind.CallOpts, arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Ubt.contract.Call(opts, &out, "allowance", arg0, arg1)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address , address ) view returns(uint256)
func (_Ubt *UbtSession) Allowance(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	return _Ubt.Contract.Allowance(&_Ubt.CallOpts, arg0, arg1)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address , address ) view returns(uint256)
func (_Ubt *UbtCallerSession) Allowance(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	return _Ubt.Contract.Allowance(&_Ubt.CallOpts, arg0, arg1)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address ) view returns(uint256)
func (_Ubt *UbtCaller) BalanceOf(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Ubt.contract.Call(opts, &out, "balanceOf", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address ) view returns(uint256)
func (_Ubt *UbtSession) BalanceOf(arg0 common.Address) (*big.Int, error) {
	return _Ubt.Contract.BalanceOf(&_Ubt.CallOpts, arg0)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address ) view returns(uint256)
func (_Ubt *UbtCallerSession) BalanceOf(arg0 common.Address) (*big.Int, error) {
	return _Ubt.Contract.BalanceOf(&_Ubt.CallOpts, arg0)
}

// BurnPrice is a free data retrieval call binding the contract method 0xc701b774.
//
// Solidity: function burn_price() view returns(uint256)
func (_Ubt *UbtCaller) BurnPrice(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Ubt.contract.Call(opts, &out, "burn_price")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BurnPrice is a free data retrieval call binding the contract method 0xc701b774.
//
// Solidity: function burn_price() view returns(uint256)
func (_Ubt *UbtSession) BurnPrice() (*big.Int, error) {
	return _Ubt.Contract.BurnPrice(&_Ubt.CallOpts)
}

// BurnPrice is a free data retrieval call binding the contract method 0xc701b774.
//
// Solidity: function burn_price() view returns(uint256)
func (_Ubt *UbtCallerSession) BurnPrice() (*big.Int, error) {
	return _Ubt.Contract.BurnPrice(&_Ubt.CallOpts)
}

// BurnSwapAddress is a free data retrieval call binding the contract method 0x3600e5f4.
//
// Solidity: function burn_swap_address(address ) view returns(bool)
func (_Ubt *UbtCaller) BurnSwapAddress(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _Ubt.contract.Call(opts, &out, "burn_swap_address", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// BurnSwapAddress is a free data retrieval call binding the contract method 0x3600e5f4.
//
// Solidity: function burn_swap_address(address ) view returns(bool)
func (_Ubt *UbtSession) BurnSwapAddress(arg0 common.Address) (bool, error) {
	return _Ubt.Contract.BurnSwapAddress(&_Ubt.CallOpts, arg0)
}

// BurnSwapAddress is a free data retrieval call binding the contract method 0x3600e5f4.
//
// Solidity: function burn_swap_address(address ) view returns(bool)
func (_Ubt *UbtCallerSession) BurnSwapAddress(arg0 common.Address) (bool, error) {
	return _Ubt.Contract.BurnSwapAddress(&_Ubt.CallOpts, arg0)
}

// DecimalBt is a free data retrieval call binding the contract method 0x45a27ed0.
//
// Solidity: function decimal_bt() view returns(uint256)
func (_Ubt *UbtCaller) DecimalBt(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Ubt.contract.Call(opts, &out, "decimal_bt")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// DecimalBt is a free data retrieval call binding the contract method 0x45a27ed0.
//
// Solidity: function decimal_bt() view returns(uint256)
func (_Ubt *UbtSession) DecimalBt() (*big.Int, error) {
	return _Ubt.Contract.DecimalBt(&_Ubt.CallOpts)
}

// DecimalBt is a free data retrieval call binding the contract method 0x45a27ed0.
//
// Solidity: function decimal_bt() view returns(uint256)
func (_Ubt *UbtCallerSession) DecimalBt() (*big.Int, error) {
	return _Ubt.Contract.DecimalBt(&_Ubt.CallOpts)
}

// DecimalUsdt is a free data retrieval call binding the contract method 0xf764dd22.
//
// Solidity: function decimal_usdt() view returns(uint256)
func (_Ubt *UbtCaller) DecimalUsdt(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Ubt.contract.Call(opts, &out, "decimal_usdt")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// DecimalUsdt is a free data retrieval call binding the contract method 0xf764dd22.
//
// Solidity: function decimal_usdt() view returns(uint256)
func (_Ubt *UbtSession) DecimalUsdt() (*big.Int, error) {
	return _Ubt.Contract.DecimalUsdt(&_Ubt.CallOpts)
}

// DecimalUsdt is a free data retrieval call binding the contract method 0xf764dd22.
//
// Solidity: function decimal_usdt() view returns(uint256)
func (_Ubt *UbtCallerSession) DecimalUsdt() (*big.Int, error) {
	return _Ubt.Contract.DecimalUsdt(&_Ubt.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_Ubt *UbtCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _Ubt.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_Ubt *UbtSession) Decimals() (uint8, error) {
	return _Ubt.Contract.Decimals(&_Ubt.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_Ubt *UbtCallerSession) Decimals() (uint8, error) {
	return _Ubt.Contract.Decimals(&_Ubt.CallOpts)
}

// Deprecated is a free data retrieval call binding the contract method 0x0e136b19.
//
// Solidity: function deprecated() view returns(bool)
func (_Ubt *UbtCaller) Deprecated(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Ubt.contract.Call(opts, &out, "deprecated")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Deprecated is a free data retrieval call binding the contract method 0x0e136b19.
//
// Solidity: function deprecated() view returns(bool)
func (_Ubt *UbtSession) Deprecated() (bool, error) {
	return _Ubt.Contract.Deprecated(&_Ubt.CallOpts)
}

// Deprecated is a free data retrieval call binding the contract method 0x0e136b19.
//
// Solidity: function deprecated() view returns(bool)
func (_Ubt *UbtCallerSession) Deprecated() (bool, error) {
	return _Ubt.Contract.Deprecated(&_Ubt.CallOpts)
}

// EpochBase is a free data retrieval call binding the contract method 0xaf2872c7.
//
// Solidity: function epoch_base() view returns(uint256)
func (_Ubt *UbtCaller) EpochBase(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Ubt.contract.Call(opts, &out, "epoch_base")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// EpochBase is a free data retrieval call binding the contract method 0xaf2872c7.
//
// Solidity: function epoch_base() view returns(uint256)
func (_Ubt *UbtSession) EpochBase() (*big.Int, error) {
	return _Ubt.Contract.EpochBase(&_Ubt.CallOpts)
}

// EpochBase is a free data retrieval call binding the contract method 0xaf2872c7.
//
// Solidity: function epoch_base() view returns(uint256)
func (_Ubt *UbtCallerSession) EpochBase() (*big.Int, error) {
	return _Ubt.Contract.EpochBase(&_Ubt.CallOpts)
}

// GetSwapPrice is a free data retrieval call binding the contract method 0xda04100b.
//
// Solidity: function getSwapPrice(uint256 value) view returns(uint256)
func (_Ubt *UbtCaller) GetSwapPrice(opts *bind.CallOpts, value *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Ubt.contract.Call(opts, &out, "getSwapPrice", value)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetSwapPrice is a free data retrieval call binding the contract method 0xda04100b.
//
// Solidity: function getSwapPrice(uint256 value) view returns(uint256)
func (_Ubt *UbtSession) GetSwapPrice(value *big.Int) (*big.Int, error) {
	return _Ubt.Contract.GetSwapPrice(&_Ubt.CallOpts, value)
}

// GetSwapPrice is a free data retrieval call binding the contract method 0xda04100b.
//
// Solidity: function getSwapPrice(uint256 value) view returns(uint256)
func (_Ubt *UbtCallerSession) GetSwapPrice(value *big.Int) (*big.Int, error) {
	return _Ubt.Contract.GetSwapPrice(&_Ubt.CallOpts, value)
}

// Invite is a free data retrieval call binding the contract method 0x4b77c468.
//
// Solidity: function invite(address ) view returns(address)
func (_Ubt *UbtCaller) Invite(opts *bind.CallOpts, arg0 common.Address) (common.Address, error) {
	var out []interface{}
	err := _Ubt.contract.Call(opts, &out, "invite", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Invite is a free data retrieval call binding the contract method 0x4b77c468.
//
// Solidity: function invite(address ) view returns(address)
func (_Ubt *UbtSession) Invite(arg0 common.Address) (common.Address, error) {
	return _Ubt.Contract.Invite(&_Ubt.CallOpts, arg0)
}

// Invite is a free data retrieval call binding the contract method 0x4b77c468.
//
// Solidity: function invite(address ) view returns(address)
func (_Ubt *UbtCallerSession) Invite(arg0 common.Address) (common.Address, error) {
	return _Ubt.Contract.Invite(&_Ubt.CallOpts, arg0)
}

// InviteCount is a free data retrieval call binding the contract method 0x57b24e6b.
//
// Solidity: function inviteCount(address ) view returns(uint256)
func (_Ubt *UbtCaller) InviteCount(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Ubt.contract.Call(opts, &out, "inviteCount", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// InviteCount is a free data retrieval call binding the contract method 0x57b24e6b.
//
// Solidity: function inviteCount(address ) view returns(uint256)
func (_Ubt *UbtSession) InviteCount(arg0 common.Address) (*big.Int, error) {
	return _Ubt.Contract.InviteCount(&_Ubt.CallOpts, arg0)
}

// InviteCount is a free data retrieval call binding the contract method 0x57b24e6b.
//
// Solidity: function inviteCount(address ) view returns(uint256)
func (_Ubt *UbtCallerSession) InviteCount(arg0 common.Address) (*big.Int, error) {
	return _Ubt.Contract.InviteCount(&_Ubt.CallOpts, arg0)
}

// LastMiner is a free data retrieval call binding the contract method 0x1f251a62.
//
// Solidity: function last_miner(address ) view returns(uint256)
func (_Ubt *UbtCaller) LastMiner(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Ubt.contract.Call(opts, &out, "last_miner", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LastMiner is a free data retrieval call binding the contract method 0x1f251a62.
//
// Solidity: function last_miner(address ) view returns(uint256)
func (_Ubt *UbtSession) LastMiner(arg0 common.Address) (*big.Int, error) {
	return _Ubt.Contract.LastMiner(&_Ubt.CallOpts, arg0)
}

// LastMiner is a free data retrieval call binding the contract method 0x1f251a62.
//
// Solidity: function last_miner(address ) view returns(uint256)
func (_Ubt *UbtCallerSession) LastMiner(arg0 common.Address) (*big.Int, error) {
	return _Ubt.Contract.LastMiner(&_Ubt.CallOpts, arg0)
}

// MintReward is a free data retrieval call binding the contract method 0x666ae7b6.
//
// Solidity: function mint_reward() view returns(uint256 usePower, uint256 amount)
func (_Ubt *UbtCaller) MintReward(opts *bind.CallOpts) (struct {
	UsePower *big.Int
	Amount   *big.Int
}, error) {
	var out []interface{}
	err := _Ubt.contract.Call(opts, &out, "mint_reward")

	outstruct := new(struct {
		UsePower *big.Int
		Amount   *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.UsePower = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Amount = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// MintReward is a free data retrieval call binding the contract method 0x666ae7b6.
//
// Solidity: function mint_reward() view returns(uint256 usePower, uint256 amount)
func (_Ubt *UbtSession) MintReward() (struct {
	UsePower *big.Int
	Amount   *big.Int
}, error) {
	return _Ubt.Contract.MintReward(&_Ubt.CallOpts)
}

// MintReward is a free data retrieval call binding the contract method 0x666ae7b6.
//
// Solidity: function mint_reward() view returns(uint256 usePower, uint256 amount)
func (_Ubt *UbtCallerSession) MintReward() (struct {
	UsePower *big.Int
	Amount   *big.Int
}, error) {
	return _Ubt.Contract.MintReward(&_Ubt.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Ubt *UbtCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Ubt.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Ubt *UbtSession) Name() (string, error) {
	return _Ubt.Contract.Name(&_Ubt.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Ubt *UbtCallerSession) Name() (string, error) {
	return _Ubt.Contract.Name(&_Ubt.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Ubt *UbtCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Ubt.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Ubt *UbtSession) Owner() (common.Address, error) {
	return _Ubt.Contract.Owner(&_Ubt.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Ubt *UbtCallerSession) Owner() (common.Address, error) {
	return _Ubt.Contract.Owner(&_Ubt.CallOpts)
}

// Power is a free data retrieval call binding the contract method 0x503371a5.
//
// Solidity: function power(address ) view returns(uint256)
func (_Ubt *UbtCaller) Power(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Ubt.contract.Call(opts, &out, "power", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Power is a free data retrieval call binding the contract method 0x503371a5.
//
// Solidity: function power(address ) view returns(uint256)
func (_Ubt *UbtSession) Power(arg0 common.Address) (*big.Int, error) {
	return _Ubt.Contract.Power(&_Ubt.CallOpts, arg0)
}

// Power is a free data retrieval call binding the contract method 0x503371a5.
//
// Solidity: function power(address ) view returns(uint256)
func (_Ubt *UbtCallerSession) Power(arg0 common.Address) (*big.Int, error) {
	return _Ubt.Contract.Power(&_Ubt.CallOpts, arg0)
}

// RewardCount is a free data retrieval call binding the contract method 0xb5af960d.
//
// Solidity: function rewardCount(address ) view returns(uint256)
func (_Ubt *UbtCaller) RewardCount(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Ubt.contract.Call(opts, &out, "rewardCount", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// RewardCount is a free data retrieval call binding the contract method 0xb5af960d.
//
// Solidity: function rewardCount(address ) view returns(uint256)
func (_Ubt *UbtSession) RewardCount(arg0 common.Address) (*big.Int, error) {
	return _Ubt.Contract.RewardCount(&_Ubt.CallOpts, arg0)
}

// RewardCount is a free data retrieval call binding the contract method 0xb5af960d.
//
// Solidity: function rewardCount(address ) view returns(uint256)
func (_Ubt *UbtCallerSession) RewardCount(arg0 common.Address) (*big.Int, error) {
	return _Ubt.Contract.RewardCount(&_Ubt.CallOpts, arg0)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Ubt *UbtCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Ubt.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Ubt *UbtSession) Symbol() (string, error) {
	return _Ubt.Contract.Symbol(&_Ubt.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Ubt *UbtCallerSession) Symbol() (string, error) {
	return _Ubt.Contract.Symbol(&_Ubt.CallOpts)
}

// TotalPower is a free data retrieval call binding the contract method 0xdb3ad22c.
//
// Solidity: function totalPower() view returns(uint256)
func (_Ubt *UbtCaller) TotalPower(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Ubt.contract.Call(opts, &out, "totalPower")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalPower is a free data retrieval call binding the contract method 0xdb3ad22c.
//
// Solidity: function totalPower() view returns(uint256)
func (_Ubt *UbtSession) TotalPower() (*big.Int, error) {
	return _Ubt.Contract.TotalPower(&_Ubt.CallOpts)
}

// TotalPower is a free data retrieval call binding the contract method 0xdb3ad22c.
//
// Solidity: function totalPower() view returns(uint256)
func (_Ubt *UbtCallerSession) TotalPower() (*big.Int, error) {
	return _Ubt.Contract.TotalPower(&_Ubt.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Ubt *UbtCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Ubt.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Ubt *UbtSession) TotalSupply() (*big.Int, error) {
	return _Ubt.Contract.TotalSupply(&_Ubt.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Ubt *UbtCallerSession) TotalSupply() (*big.Int, error) {
	return _Ubt.Contract.TotalSupply(&_Ubt.CallOpts)
}

// TotalUsersAmount is a free data retrieval call binding the contract method 0xdd94a29d.
//
// Solidity: function totalUsersAmount() view returns(uint256)
func (_Ubt *UbtCaller) TotalUsersAmount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Ubt.contract.Call(opts, &out, "totalUsersAmount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalUsersAmount is a free data retrieval call binding the contract method 0xdd94a29d.
//
// Solidity: function totalUsersAmount() view returns(uint256)
func (_Ubt *UbtSession) TotalUsersAmount() (*big.Int, error) {
	return _Ubt.Contract.TotalUsersAmount(&_Ubt.CallOpts)
}

// TotalUsersAmount is a free data retrieval call binding the contract method 0xdd94a29d.
//
// Solidity: function totalUsersAmount() view returns(uint256)
func (_Ubt *UbtCallerSession) TotalUsersAmount() (*big.Int, error) {
	return _Ubt.Contract.TotalUsersAmount(&_Ubt.CallOpts)
}

// TransferFee is a free data retrieval call binding the contract method 0x5914dc6a.
//
// Solidity: function transfer_fee(address _from, address _to, uint256 _value) view returns(uint256 fee)
func (_Ubt *UbtCaller) TransferFee(opts *bind.CallOpts, _from common.Address, _to common.Address, _value *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Ubt.contract.Call(opts, &out, "transfer_fee", _from, _to, _value)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TransferFee is a free data retrieval call binding the contract method 0x5914dc6a.
//
// Solidity: function transfer_fee(address _from, address _to, uint256 _value) view returns(uint256 fee)
func (_Ubt *UbtSession) TransferFee(_from common.Address, _to common.Address, _value *big.Int) (*big.Int, error) {
	return _Ubt.Contract.TransferFee(&_Ubt.CallOpts, _from, _to, _value)
}

// TransferFee is a free data retrieval call binding the contract method 0x5914dc6a.
//
// Solidity: function transfer_fee(address _from, address _to, uint256 _value) view returns(uint256 fee)
func (_Ubt *UbtCallerSession) TransferFee(_from common.Address, _to common.Address, _value *big.Int) (*big.Int, error) {
	return _Ubt.Contract.TransferFee(&_Ubt.CallOpts, _from, _to, _value)
}

// UpgradedAddress is a free data retrieval call binding the contract method 0x26976e3f.
//
// Solidity: function upgradedAddress() view returns(address)
func (_Ubt *UbtCaller) UpgradedAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Ubt.contract.Call(opts, &out, "upgradedAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// UpgradedAddress is a free data retrieval call binding the contract method 0x26976e3f.
//
// Solidity: function upgradedAddress() view returns(address)
func (_Ubt *UbtSession) UpgradedAddress() (common.Address, error) {
	return _Ubt.Contract.UpgradedAddress(&_Ubt.CallOpts)
}

// UpgradedAddress is a free data retrieval call binding the contract method 0x26976e3f.
//
// Solidity: function upgradedAddress() view returns(address)
func (_Ubt *UbtCallerSession) UpgradedAddress() (common.Address, error) {
	return _Ubt.Contract.UpgradedAddress(&_Ubt.CallOpts)
}

// UsersEpoch is a free data retrieval call binding the contract method 0xa6f60cbd.
//
// Solidity: function users_epoch(address ) view returns(uint256)
func (_Ubt *UbtCaller) UsersEpoch(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Ubt.contract.Call(opts, &out, "users_epoch", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// UsersEpoch is a free data retrieval call binding the contract method 0xa6f60cbd.
//
// Solidity: function users_epoch(address ) view returns(uint256)
func (_Ubt *UbtSession) UsersEpoch(arg0 common.Address) (*big.Int, error) {
	return _Ubt.Contract.UsersEpoch(&_Ubt.CallOpts, arg0)
}

// UsersEpoch is a free data retrieval call binding the contract method 0xa6f60cbd.
//
// Solidity: function users_epoch(address ) view returns(uint256)
func (_Ubt *UbtCallerSession) UsersEpoch(arg0 common.Address) (*big.Int, error) {
	return _Ubt.Contract.UsersEpoch(&_Ubt.CallOpts, arg0)
}

// UsersStartTime is a free data retrieval call binding the contract method 0x05908b3e.
//
// Solidity: function users_start_time(address ) view returns(uint256)
func (_Ubt *UbtCaller) UsersStartTime(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Ubt.contract.Call(opts, &out, "users_start_time", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// UsersStartTime is a free data retrieval call binding the contract method 0x05908b3e.
//
// Solidity: function users_start_time(address ) view returns(uint256)
func (_Ubt *UbtSession) UsersStartTime(arg0 common.Address) (*big.Int, error) {
	return _Ubt.Contract.UsersStartTime(&_Ubt.CallOpts, arg0)
}

// UsersStartTime is a free data retrieval call binding the contract method 0x05908b3e.
//
// Solidity: function users_start_time(address ) view returns(uint256)
func (_Ubt *UbtCallerSession) UsersStartTime(arg0 common.Address) (*big.Int, error) {
	return _Ubt.Contract.UsersStartTime(&_Ubt.CallOpts, arg0)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address _spender, uint256 _value) returns(bool success)
func (_Ubt *UbtTransactor) Approve(opts *bind.TransactOpts, _spender common.Address, _value *big.Int) (*types.Transaction, error) {
	return _Ubt.contract.Transact(opts, "approve", _spender, _value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address _spender, uint256 _value) returns(bool success)
func (_Ubt *UbtSession) Approve(_spender common.Address, _value *big.Int) (*types.Transaction, error) {
	return _Ubt.Contract.Approve(&_Ubt.TransactOpts, _spender, _value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address _spender, uint256 _value) returns(bool success)
func (_Ubt *UbtTransactorSession) Approve(_spender common.Address, _value *big.Int) (*types.Transaction, error) {
	return _Ubt.Contract.Approve(&_Ubt.TransactOpts, _spender, _value)
}

// AutoUpdateBurnPrice is a paid mutator transaction binding the contract method 0x68735da1.
//
// Solidity: function auto_update_burn_price() returns(bool success)
func (_Ubt *UbtTransactor) AutoUpdateBurnPrice(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Ubt.contract.Transact(opts, "auto_update_burn_price")
}

// AutoUpdateBurnPrice is a paid mutator transaction binding the contract method 0x68735da1.
//
// Solidity: function auto_update_burn_price() returns(bool success)
func (_Ubt *UbtSession) AutoUpdateBurnPrice() (*types.Transaction, error) {
	return _Ubt.Contract.AutoUpdateBurnPrice(&_Ubt.TransactOpts)
}

// AutoUpdateBurnPrice is a paid mutator transaction binding the contract method 0x68735da1.
//
// Solidity: function auto_update_burn_price() returns(bool success)
func (_Ubt *UbtTransactorSession) AutoUpdateBurnPrice() (*types.Transaction, error) {
	return _Ubt.Contract.AutoUpdateBurnPrice(&_Ubt.TransactOpts)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(uint256 _value) returns(bool success)
func (_Ubt *UbtTransactor) Burn(opts *bind.TransactOpts, _value *big.Int) (*types.Transaction, error) {
	return _Ubt.contract.Transact(opts, "burn", _value)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(uint256 _value) returns(bool success)
func (_Ubt *UbtSession) Burn(_value *big.Int) (*types.Transaction, error) {
	return _Ubt.Contract.Burn(&_Ubt.TransactOpts, _value)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(uint256 _value) returns(bool success)
func (_Ubt *UbtTransactorSession) Burn(_value *big.Int) (*types.Transaction, error) {
	return _Ubt.Contract.Burn(&_Ubt.TransactOpts, _value)
}

// Mint is a paid mutator transaction binding the contract method 0x1249c58b.
//
// Solidity: function mint() returns(bool success)
func (_Ubt *UbtTransactor) Mint(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Ubt.contract.Transact(opts, "mint")
}

// Mint is a paid mutator transaction binding the contract method 0x1249c58b.
//
// Solidity: function mint() returns(bool success)
func (_Ubt *UbtSession) Mint() (*types.Transaction, error) {
	return _Ubt.Contract.Mint(&_Ubt.TransactOpts)
}

// Mint is a paid mutator transaction binding the contract method 0x1249c58b.
//
// Solidity: function mint() returns(bool success)
func (_Ubt *UbtTransactorSession) Mint() (*types.Transaction, error) {
	return _Ubt.Contract.Mint(&_Ubt.TransactOpts)
}

// Registration is a paid mutator transaction binding the contract method 0x0840605a.
//
// Solidity: function registration(address invite_address) returns(bool success)
func (_Ubt *UbtTransactor) Registration(opts *bind.TransactOpts, invite_address common.Address) (*types.Transaction, error) {
	return _Ubt.contract.Transact(opts, "registration", invite_address)
}

// Registration is a paid mutator transaction binding the contract method 0x0840605a.
//
// Solidity: function registration(address invite_address) returns(bool success)
func (_Ubt *UbtSession) Registration(invite_address common.Address) (*types.Transaction, error) {
	return _Ubt.Contract.Registration(&_Ubt.TransactOpts, invite_address)
}

// Registration is a paid mutator transaction binding the contract method 0x0840605a.
//
// Solidity: function registration(address invite_address) returns(bool success)
func (_Ubt *UbtTransactorSession) Registration(invite_address common.Address) (*types.Transaction, error) {
	return _Ubt.Contract.Registration(&_Ubt.TransactOpts, invite_address)
}

// SetOracle is a paid mutator transaction binding the contract method 0x7adbf973.
//
// Solidity: function setOracle(address _oracle_address) returns()
func (_Ubt *UbtTransactor) SetOracle(opts *bind.TransactOpts, _oracle_address common.Address) (*types.Transaction, error) {
	return _Ubt.contract.Transact(opts, "setOracle", _oracle_address)
}

// SetOracle is a paid mutator transaction binding the contract method 0x7adbf973.
//
// Solidity: function setOracle(address _oracle_address) returns()
func (_Ubt *UbtSession) SetOracle(_oracle_address common.Address) (*types.Transaction, error) {
	return _Ubt.Contract.SetOracle(&_Ubt.TransactOpts, _oracle_address)
}

// SetOracle is a paid mutator transaction binding the contract method 0x7adbf973.
//
// Solidity: function setOracle(address _oracle_address) returns()
func (_Ubt *UbtTransactorSession) SetOracle(_oracle_address common.Address) (*types.Transaction, error) {
	return _Ubt.Contract.SetOracle(&_Ubt.TransactOpts, _oracle_address)
}

// SetOwner is a paid mutator transaction binding the contract method 0x13af4035.
//
// Solidity: function setOwner(address new_owner) returns()
func (_Ubt *UbtTransactor) SetOwner(opts *bind.TransactOpts, new_owner common.Address) (*types.Transaction, error) {
	return _Ubt.contract.Transact(opts, "setOwner", new_owner)
}

// SetOwner is a paid mutator transaction binding the contract method 0x13af4035.
//
// Solidity: function setOwner(address new_owner) returns()
func (_Ubt *UbtSession) SetOwner(new_owner common.Address) (*types.Transaction, error) {
	return _Ubt.Contract.SetOwner(&_Ubt.TransactOpts, new_owner)
}

// SetOwner is a paid mutator transaction binding the contract method 0x13af4035.
//
// Solidity: function setOwner(address new_owner) returns()
func (_Ubt *UbtTransactorSession) SetOwner(new_owner common.Address) (*types.Transaction, error) {
	return _Ubt.Contract.SetOwner(&_Ubt.TransactOpts, new_owner)
}

// SetSwapAddress is a paid mutator transaction binding the contract method 0xf457b734.
//
// Solidity: function set_swap_address(address _swap_address) returns()
func (_Ubt *UbtTransactor) SetSwapAddress(opts *bind.TransactOpts, _swap_address common.Address) (*types.Transaction, error) {
	return _Ubt.contract.Transact(opts, "set_swap_address", _swap_address)
}

// SetSwapAddress is a paid mutator transaction binding the contract method 0xf457b734.
//
// Solidity: function set_swap_address(address _swap_address) returns()
func (_Ubt *UbtSession) SetSwapAddress(_swap_address common.Address) (*types.Transaction, error) {
	return _Ubt.Contract.SetSwapAddress(&_Ubt.TransactOpts, _swap_address)
}

// SetSwapAddress is a paid mutator transaction binding the contract method 0xf457b734.
//
// Solidity: function set_swap_address(address _swap_address) returns()
func (_Ubt *UbtTransactorSession) SetSwapAddress(_swap_address common.Address) (*types.Transaction, error) {
	return _Ubt.Contract.SetSwapAddress(&_Ubt.TransactOpts, _swap_address)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address _to, uint256 _value) returns(bool success)
func (_Ubt *UbtTransactor) Transfer(opts *bind.TransactOpts, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _Ubt.contract.Transact(opts, "transfer", _to, _value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address _to, uint256 _value) returns(bool success)
func (_Ubt *UbtSession) Transfer(_to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _Ubt.Contract.Transfer(&_Ubt.TransactOpts, _to, _value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address _to, uint256 _value) returns(bool success)
func (_Ubt *UbtTransactorSession) Transfer(_to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _Ubt.Contract.Transfer(&_Ubt.TransactOpts, _to, _value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address _from, address _to, uint256 _value) returns(bool success)
func (_Ubt *UbtTransactor) TransferFrom(opts *bind.TransactOpts, _from common.Address, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _Ubt.contract.Transact(opts, "transferFrom", _from, _to, _value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address _from, address _to, uint256 _value) returns(bool success)
func (_Ubt *UbtSession) TransferFrom(_from common.Address, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _Ubt.Contract.TransferFrom(&_Ubt.TransactOpts, _from, _to, _value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address _from, address _to, uint256 _value) returns(bool success)
func (_Ubt *UbtTransactorSession) TransferFrom(_from common.Address, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _Ubt.Contract.TransferFrom(&_Ubt.TransactOpts, _from, _to, _value)
}

// UpdateBurnPrice is a paid mutator transaction binding the contract method 0xb0c8a16c.
//
// Solidity: function update_burn_price() returns(bool success)
func (_Ubt *UbtTransactor) UpdateBurnPrice(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Ubt.contract.Transact(opts, "update_burn_price")
}

// UpdateBurnPrice is a paid mutator transaction binding the contract method 0xb0c8a16c.
//
// Solidity: function update_burn_price() returns(bool success)
func (_Ubt *UbtSession) UpdateBurnPrice() (*types.Transaction, error) {
	return _Ubt.Contract.UpdateBurnPrice(&_Ubt.TransactOpts)
}

// UpdateBurnPrice is a paid mutator transaction binding the contract method 0xb0c8a16c.
//
// Solidity: function update_burn_price() returns(bool success)
func (_Ubt *UbtTransactorSession) UpdateBurnPrice() (*types.Transaction, error) {
	return _Ubt.Contract.UpdateBurnPrice(&_Ubt.TransactOpts)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 amount) returns()
func (_Ubt *UbtTransactor) Withdraw(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _Ubt.contract.Transact(opts, "withdraw", amount)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 amount) returns()
func (_Ubt *UbtSession) Withdraw(amount *big.Int) (*types.Transaction, error) {
	return _Ubt.Contract.Withdraw(&_Ubt.TransactOpts, amount)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 amount) returns()
func (_Ubt *UbtTransactorSession) Withdraw(amount *big.Int) (*types.Transaction, error) {
	return _Ubt.Contract.Withdraw(&_Ubt.TransactOpts, amount)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Ubt *UbtTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Ubt.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Ubt *UbtSession) Receive() (*types.Transaction, error) {
	return _Ubt.Contract.Receive(&_Ubt.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Ubt *UbtTransactorSession) Receive() (*types.Transaction, error) {
	return _Ubt.Contract.Receive(&_Ubt.TransactOpts)
}

// UbtBurnIterator is returned from FilterBurn and is used to iterate over the raw logs and unpacked data for Burn events raised by the Ubt contract.
type UbtBurnIterator struct {
	Event *UbtBurn // Event containing the contract specifics and raw log

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
func (it *UbtBurnIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(UbtBurn)
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
		it.Event = new(UbtBurn)
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
func (it *UbtBurnIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *UbtBurnIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// UbtBurn represents a Burn event raised by the Ubt contract.
type UbtBurn struct {
	From  common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterBurn is a free log retrieval operation binding the contract event 0xcc16f5dbb4873280815c1ee09dbd06736cffcc184412cf7a71a0fdb75d397ca5.
//
// Solidity: event Burn(address indexed from, uint256 value)
func (_Ubt *UbtFilterer) FilterBurn(opts *bind.FilterOpts, from []common.Address) (*UbtBurnIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}

	logs, sub, err := _Ubt.contract.FilterLogs(opts, "Burn", fromRule)
	if err != nil {
		return nil, err
	}
	return &UbtBurnIterator{contract: _Ubt.contract, event: "Burn", logs: logs, sub: sub}, nil
}

// WatchBurn is a free log subscription operation binding the contract event 0xcc16f5dbb4873280815c1ee09dbd06736cffcc184412cf7a71a0fdb75d397ca5.
//
// Solidity: event Burn(address indexed from, uint256 value)
func (_Ubt *UbtFilterer) WatchBurn(opts *bind.WatchOpts, sink chan<- *UbtBurn, from []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}

	logs, sub, err := _Ubt.contract.WatchLogs(opts, "Burn", fromRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(UbtBurn)
				if err := _Ubt.contract.UnpackLog(event, "Burn", log); err != nil {
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

// ParseBurn is a log parse operation binding the contract event 0xcc16f5dbb4873280815c1ee09dbd06736cffcc184412cf7a71a0fdb75d397ca5.
//
// Solidity: event Burn(address indexed from, uint256 value)
func (_Ubt *UbtFilterer) ParseBurn(log types.Log) (*UbtBurn, error) {
	event := new(UbtBurn)
	if err := _Ubt.contract.UnpackLog(event, "Burn", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// UbtFreezeIterator is returned from FilterFreeze and is used to iterate over the raw logs and unpacked data for Freeze events raised by the Ubt contract.
type UbtFreezeIterator struct {
	Event *UbtFreeze // Event containing the contract specifics and raw log

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
func (it *UbtFreezeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(UbtFreeze)
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
		it.Event = new(UbtFreeze)
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
func (it *UbtFreezeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *UbtFreezeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// UbtFreeze represents a Freeze event raised by the Ubt contract.
type UbtFreeze struct {
	From  common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterFreeze is a free log retrieval operation binding the contract event 0xf97a274face0b5517365ad396b1fdba6f68bd3135ef603e44272adba3af5a1e0.
//
// Solidity: event Freeze(address indexed from, uint256 value)
func (_Ubt *UbtFilterer) FilterFreeze(opts *bind.FilterOpts, from []common.Address) (*UbtFreezeIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}

	logs, sub, err := _Ubt.contract.FilterLogs(opts, "Freeze", fromRule)
	if err != nil {
		return nil, err
	}
	return &UbtFreezeIterator{contract: _Ubt.contract, event: "Freeze", logs: logs, sub: sub}, nil
}

// WatchFreeze is a free log subscription operation binding the contract event 0xf97a274face0b5517365ad396b1fdba6f68bd3135ef603e44272adba3af5a1e0.
//
// Solidity: event Freeze(address indexed from, uint256 value)
func (_Ubt *UbtFilterer) WatchFreeze(opts *bind.WatchOpts, sink chan<- *UbtFreeze, from []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}

	logs, sub, err := _Ubt.contract.WatchLogs(opts, "Freeze", fromRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(UbtFreeze)
				if err := _Ubt.contract.UnpackLog(event, "Freeze", log); err != nil {
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

// ParseFreeze is a log parse operation binding the contract event 0xf97a274face0b5517365ad396b1fdba6f68bd3135ef603e44272adba3af5a1e0.
//
// Solidity: event Freeze(address indexed from, uint256 value)
func (_Ubt *UbtFilterer) ParseFreeze(log types.Log) (*UbtFreeze, error) {
	event := new(UbtFreeze)
	if err := _Ubt.contract.UnpackLog(event, "Freeze", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// UbtMintedIterator is returned from FilterMinted and is used to iterate over the raw logs and unpacked data for Minted events raised by the Ubt contract.
type UbtMintedIterator struct {
	Event *UbtMinted // Event containing the contract specifics and raw log

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
func (it *UbtMintedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(UbtMinted)
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
		it.Event = new(UbtMinted)
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
func (it *UbtMintedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *UbtMintedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// UbtMinted represents a Minted event raised by the Ubt contract.
type UbtMinted struct {
	Operator common.Address
	To       common.Address
	Amount   *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterMinted is a free log retrieval operation binding the contract event 0x9d228d69b5fdb8d273a2336f8fb8612d039631024ea9bf09c424a9503aa078f0.
//
// Solidity: event Minted(address indexed operator, address indexed to, uint256 amount)
func (_Ubt *UbtFilterer) FilterMinted(opts *bind.FilterOpts, operator []common.Address, to []common.Address) (*UbtMintedIterator, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Ubt.contract.FilterLogs(opts, "Minted", operatorRule, toRule)
	if err != nil {
		return nil, err
	}
	return &UbtMintedIterator{contract: _Ubt.contract, event: "Minted", logs: logs, sub: sub}, nil
}

// WatchMinted is a free log subscription operation binding the contract event 0x9d228d69b5fdb8d273a2336f8fb8612d039631024ea9bf09c424a9503aa078f0.
//
// Solidity: event Minted(address indexed operator, address indexed to, uint256 amount)
func (_Ubt *UbtFilterer) WatchMinted(opts *bind.WatchOpts, sink chan<- *UbtMinted, operator []common.Address, to []common.Address) (event.Subscription, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Ubt.contract.WatchLogs(opts, "Minted", operatorRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(UbtMinted)
				if err := _Ubt.contract.UnpackLog(event, "Minted", log); err != nil {
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

// ParseMinted is a log parse operation binding the contract event 0x9d228d69b5fdb8d273a2336f8fb8612d039631024ea9bf09c424a9503aa078f0.
//
// Solidity: event Minted(address indexed operator, address indexed to, uint256 amount)
func (_Ubt *UbtFilterer) ParseMinted(log types.Log) (*UbtMinted, error) {
	event := new(UbtMinted)
	if err := _Ubt.contract.UnpackLog(event, "Minted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// UbtTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the Ubt contract.
type UbtTransferIterator struct {
	Event *UbtTransfer // Event containing the contract specifics and raw log

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
func (it *UbtTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(UbtTransfer)
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
		it.Event = new(UbtTransfer)
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
func (it *UbtTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *UbtTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// UbtTransfer represents a Transfer event raised by the Ubt contract.
type UbtTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_Ubt *UbtFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*UbtTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Ubt.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &UbtTransferIterator{contract: _Ubt.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_Ubt *UbtFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *UbtTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Ubt.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(UbtTransfer)
				if err := _Ubt.contract.UnpackLog(event, "Transfer", log); err != nil {
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

// ParseTransfer is a log parse operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_Ubt *UbtFilterer) ParseTransfer(log types.Log) (*UbtTransfer, error) {
	event := new(UbtTransfer)
	if err := _Ubt.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
