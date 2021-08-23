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

// BtABI is the input ABI used to generate the binding from.
// Deprecated: Use BtMetaData.ABI instead.
const BtABI = "[{\"type\":\"constructor\",\"stateMutability\":\"nonpayable\",\"inputs\":[{\"type\":\"string\",\"name\":\"tokenName\",\"internalType\":\"string\"},{\"type\":\"string\",\"name\":\"tokenSymbol\",\"internalType\":\"string\"},{\"type\":\"uint256\",\"name\":\"epoch_time\",\"internalType\":\"uint256\"}]},{\"type\":\"event\",\"name\":\"Burn\",\"inputs\":[{\"type\":\"address\",\"name\":\"from\",\"internalType\":\"address\",\"indexed\":true},{\"type\":\"uint256\",\"name\":\"value\",\"internalType\":\"uint256\",\"indexed\":false}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Freeze\",\"inputs\":[{\"type\":\"address\",\"name\":\"from\",\"internalType\":\"address\",\"indexed\":true},{\"type\":\"uint256\",\"name\":\"value\",\"internalType\":\"uint256\",\"indexed\":false}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Minted\",\"inputs\":[{\"type\":\"address\",\"name\":\"operator\",\"internalType\":\"address\",\"indexed\":true},{\"type\":\"address\",\"name\":\"to\",\"internalType\":\"address\",\"indexed\":true},{\"type\":\"uint256\",\"name\":\"amount\",\"internalType\":\"uint256\",\"indexed\":false}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Transfer\",\"inputs\":[{\"type\":\"address\",\"name\":\"from\",\"internalType\":\"address\",\"indexed\":true},{\"type\":\"address\",\"name\":\"to\",\"internalType\":\"address\",\"indexed\":true},{\"type\":\"uint256\",\"name\":\"value\",\"internalType\":\"uint256\",\"indexed\":false}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Unfreeze\",\"inputs\":[{\"type\":\"address\",\"name\":\"from\",\"internalType\":\"address\",\"indexed\":true},{\"type\":\"uint256\",\"name\":\"value\",\"internalType\":\"uint256\",\"indexed\":false}],\"anonymous\":false},{\"type\":\"function\",\"stateMutability\":\"nonpayable\",\"outputs\":[],\"name\":\"airdrop\",\"inputs\":[{\"type\":\"address[]\",\"name\":\"address_array\",\"internalType\":\"address[]\"}]},{\"type\":\"function\",\"stateMutability\":\"view\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"\",\"internalType\":\"uint256\"}],\"name\":\"allowance\",\"inputs\":[{\"type\":\"address\",\"name\":\"\",\"internalType\":\"address\"},{\"type\":\"address\",\"name\":\"\",\"internalType\":\"address\"}]},{\"type\":\"function\",\"stateMutability\":\"view\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"\",\"internalType\":\"uint256\"}],\"name\":\"anti_bot\",\"inputs\":[]},{\"type\":\"function\",\"stateMutability\":\"nonpayable\",\"outputs\":[{\"type\":\"bool\",\"name\":\"success\",\"internalType\":\"bool\"}],\"name\":\"approve\",\"inputs\":[{\"type\":\"address\",\"name\":\"_spender\",\"internalType\":\"address\"},{\"type\":\"uint256\",\"name\":\"_value\",\"internalType\":\"uint256\"}]},{\"type\":\"function\",\"stateMutability\":\"view\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"\",\"internalType\":\"uint256\"}],\"name\":\"balanceOf\",\"inputs\":[{\"type\":\"address\",\"name\":\"\",\"internalType\":\"address\"}]},{\"type\":\"function\",\"stateMutability\":\"nonpayable\",\"outputs\":[{\"type\":\"bool\",\"name\":\"success\",\"internalType\":\"bool\"}],\"name\":\"burn\",\"inputs\":[{\"type\":\"uint256\",\"name\":\"_value\",\"internalType\":\"uint256\"}]},{\"type\":\"function\",\"stateMutability\":\"view\",\"outputs\":[{\"type\":\"uint8\",\"name\":\"\",\"internalType\":\"uint8\"}],\"name\":\"decimals\",\"inputs\":[]},{\"type\":\"function\",\"stateMutability\":\"view\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"\",\"internalType\":\"uint256\"}],\"name\":\"epoch\",\"inputs\":[]},{\"type\":\"function\",\"stateMutability\":\"view\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"\",\"internalType\":\"uint256\"}],\"name\":\"epoch_add\",\"inputs\":[]},{\"type\":\"function\",\"stateMutability\":\"view\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"\",\"internalType\":\"uint256\"}],\"name\":\"epoch_base\",\"inputs\":[]},{\"type\":\"function\",\"stateMutability\":\"nonpayable\",\"outputs\":[{\"type\":\"bool\",\"name\":\"success\",\"internalType\":\"bool\"}],\"name\":\"freeze\",\"inputs\":[{\"type\":\"uint256\",\"name\":\"_value\",\"internalType\":\"uint256\"}]},{\"type\":\"function\",\"stateMutability\":\"view\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"\",\"internalType\":\"uint256\"}],\"name\":\"freezeOf\",\"inputs\":[{\"type\":\"address\",\"name\":\"\",\"internalType\":\"address\"}]},{\"type\":\"function\",\"stateMutability\":\"view\",\"outputs\":[{\"type\":\"address\",\"name\":\"\",\"internalType\":\"address\"}],\"name\":\"invite\",\"inputs\":[{\"type\":\"address\",\"name\":\"\",\"internalType\":\"address\"}]},{\"type\":\"function\",\"stateMutability\":\"view\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"\",\"internalType\":\"uint256\"}],\"name\":\"inviteCount\",\"inputs\":[{\"type\":\"address\",\"name\":\"\",\"internalType\":\"address\"}]},{\"type\":\"function\",\"stateMutability\":\"view\",\"outputs\":[{\"type\":\"bool\",\"name\":\"\",\"internalType\":\"bool\"}],\"name\":\"is_airdrop\",\"inputs\":[]},{\"type\":\"function\",\"stateMutability\":\"view\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"\",\"internalType\":\"uint256\"}],\"name\":\"last_miner\",\"inputs\":[{\"type\":\"address\",\"name\":\"\",\"internalType\":\"address\"}]},{\"type\":\"function\",\"stateMutability\":\"nonpayable\",\"outputs\":[{\"type\":\"bool\",\"name\":\"success\",\"internalType\":\"bool\"}],\"name\":\"mint\",\"inputs\":[]},{\"type\":\"function\",\"stateMutability\":\"view\",\"outputs\":[{\"type\":\"string\",\"name\":\"\",\"internalType\":\"string\"}],\"name\":\"name\",\"inputs\":[]},{\"type\":\"function\",\"stateMutability\":\"view\",\"outputs\":[{\"type\":\"address\",\"name\":\"\",\"internalType\":\"addresspayable\"}],\"name\":\"owner\",\"inputs\":[]},{\"type\":\"function\",\"stateMutability\":\"view\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"\",\"internalType\":\"uint256\"}],\"name\":\"power\",\"inputs\":[{\"type\":\"address\",\"name\":\"\",\"internalType\":\"address\"}]},{\"type\":\"function\",\"stateMutability\":\"nonpayable\",\"outputs\":[{\"type\":\"bool\",\"name\":\"success\",\"internalType\":\"bool\"}],\"name\":\"registration\",\"inputs\":[{\"type\":\"address\",\"name\":\"invite_address\",\"internalType\":\"address\"}]},{\"type\":\"function\",\"stateMutability\":\"view\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"\",\"internalType\":\"uint256\"}],\"name\":\"rewardCount\",\"inputs\":[{\"type\":\"address\",\"name\":\"\",\"internalType\":\"address\"}]},{\"type\":\"function\",\"stateMutability\":\"nonpayable\",\"outputs\":[],\"name\":\"setOwner\",\"inputs\":[{\"type\":\"address\",\"name\":\"new_owner\",\"internalType\":\"addresspayable\"}]},{\"type\":\"function\",\"stateMutability\":\"nonpayable\",\"outputs\":[],\"name\":\"set_anti_bot\",\"inputs\":[{\"type\":\"uint256\",\"name\":\"_value\",\"internalType\":\"uint256\"}]},{\"type\":\"function\",\"stateMutability\":\"view\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"\",\"internalType\":\"uint256\"}],\"name\":\"start_time\",\"inputs\":[]},{\"type\":\"function\",\"stateMutability\":\"nonpayable\",\"outputs\":[],\"name\":\"stop_airdrop\",\"inputs\":[]},{\"type\":\"function\",\"stateMutability\":\"view\",\"outputs\":[{\"type\":\"string\",\"name\":\"\",\"internalType\":\"string\"}],\"name\":\"symbol\",\"inputs\":[]},{\"type\":\"function\",\"stateMutability\":\"view\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"\",\"internalType\":\"uint256\"}],\"name\":\"totalPower\",\"inputs\":[]},{\"type\":\"function\",\"stateMutability\":\"view\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"\",\"internalType\":\"uint256\"}],\"name\":\"totalSupply\",\"inputs\":[]},{\"type\":\"function\",\"stateMutability\":\"view\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"\",\"internalType\":\"uint256\"}],\"name\":\"totalUsersAmount\",\"inputs\":[]},{\"type\":\"function\",\"stateMutability\":\"nonpayable\",\"outputs\":[{\"type\":\"bool\",\"name\":\"success\",\"internalType\":\"bool\"}],\"name\":\"transfer\",\"inputs\":[{\"type\":\"address\",\"name\":\"_to\",\"internalType\":\"address\"},{\"type\":\"uint256\",\"name\":\"_value\",\"internalType\":\"uint256\"}]},{\"type\":\"function\",\"stateMutability\":\"nonpayable\",\"outputs\":[{\"type\":\"bool\",\"name\":\"success\",\"internalType\":\"bool\"}],\"name\":\"transferFrom\",\"inputs\":[{\"type\":\"address\",\"name\":\"_from\",\"internalType\":\"address\"},{\"type\":\"address\",\"name\":\"_to\",\"internalType\":\"address\"},{\"type\":\"uint256\",\"name\":\"_value\",\"internalType\":\"uint256\"}]},{\"type\":\"function\",\"stateMutability\":\"view\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"fee\",\"internalType\":\"uint256\"}],\"name\":\"transfer_fee\",\"inputs\":[{\"type\":\"address\",\"name\":\"_from\",\"internalType\":\"address\"},{\"type\":\"uint256\",\"name\":\"_value\",\"internalType\":\"uint256\"}]},{\"type\":\"function\",\"stateMutability\":\"nonpayable\",\"outputs\":[{\"type\":\"bool\",\"name\":\"success\",\"internalType\":\"bool\"}],\"name\":\"unfreeze\",\"inputs\":[{\"type\":\"uint256\",\"name\":\"_value\",\"internalType\":\"uint256\"}]},{\"type\":\"function\",\"stateMutability\":\"nonpayable\",\"outputs\":[],\"name\":\"withdraw\",\"inputs\":[{\"type\":\"uint256\",\"name\":\"amount\",\"internalType\":\"uint256\"}]},{\"type\":\"receive\",\"stateMutability\":\"payable\"}]"

// Bt is an auto generated Go binding around an Ethereum contract.
type Bt struct {
	BtCaller     // Read-only binding to the contract
	BtTransactor // Write-only binding to the contract
	BtFilterer   // Log filterer for contract events
}

// BtCaller is an auto generated read-only Go binding around an Ethereum contract.
type BtCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BtTransactor is an auto generated write-only Go binding around an Ethereum contract.
type BtTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BtFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type BtFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BtSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type BtSession struct {
	Contract     *Bt               // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BtCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type BtCallerSession struct {
	Contract *BtCaller     // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// BtTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type BtTransactorSession struct {
	Contract     *BtTransactor     // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BtRaw is an auto generated low-level Go binding around an Ethereum contract.
type BtRaw struct {
	Contract *Bt // Generic contract binding to access the raw methods on
}

// BtCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type BtCallerRaw struct {
	Contract *BtCaller // Generic read-only contract binding to access the raw methods on
}

// BtTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type BtTransactorRaw struct {
	Contract *BtTransactor // Generic write-only contract binding to access the raw methods on
}

// NewBt creates a new instance of Bt, bound to a specific deployed contract.
func NewBt(address common.Address, backend bind.ContractBackend) (*Bt, error) {
	contract, err := bindBt(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Bt{BtCaller: BtCaller{contract: contract}, BtTransactor: BtTransactor{contract: contract}, BtFilterer: BtFilterer{contract: contract}}, nil
}

// NewBtCaller creates a new read-only instance of Bt, bound to a specific deployed contract.
func NewBtCaller(address common.Address, caller bind.ContractCaller) (*BtCaller, error) {
	contract, err := bindBt(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &BtCaller{contract: contract}, nil
}

// NewBtTransactor creates a new write-only instance of Bt, bound to a specific deployed contract.
func NewBtTransactor(address common.Address, transactor bind.ContractTransactor) (*BtTransactor, error) {
	contract, err := bindBt(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &BtTransactor{contract: contract}, nil
}

// NewBtFilterer creates a new log filterer instance of Bt, bound to a specific deployed contract.
func NewBtFilterer(address common.Address, filterer bind.ContractFilterer) (*BtFilterer, error) {
	contract, err := bindBt(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &BtFilterer{contract: contract}, nil
}

// bindBt binds a generic wrapper to an already deployed contract.
func bindBt(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(BtABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Bt *BtRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Bt.Contract.BtCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Bt *BtRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Bt.Contract.BtTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Bt *BtRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Bt.Contract.BtTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Bt *BtCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Bt.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Bt *BtTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Bt.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Bt *BtTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Bt.Contract.contract.Transact(opts, method, params...)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address , address ) view returns(uint256)
func (_Bt *BtCaller) Allowance(opts *bind.CallOpts, arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Bt.contract.Call(opts, &out, "allowance", arg0, arg1)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address , address ) view returns(uint256)
func (_Bt *BtSession) Allowance(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	return _Bt.Contract.Allowance(&_Bt.CallOpts, arg0, arg1)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address , address ) view returns(uint256)
func (_Bt *BtCallerSession) Allowance(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	return _Bt.Contract.Allowance(&_Bt.CallOpts, arg0, arg1)
}

// AntiBot is a free data retrieval call binding the contract method 0x50b98f50.
//
// Solidity: function anti_bot() view returns(uint256)
func (_Bt *BtCaller) AntiBot(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Bt.contract.Call(opts, &out, "anti_bot")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// AntiBot is a free data retrieval call binding the contract method 0x50b98f50.
//
// Solidity: function anti_bot() view returns(uint256)
func (_Bt *BtSession) AntiBot() (*big.Int, error) {
	return _Bt.Contract.AntiBot(&_Bt.CallOpts)
}

// AntiBot is a free data retrieval call binding the contract method 0x50b98f50.
//
// Solidity: function anti_bot() view returns(uint256)
func (_Bt *BtCallerSession) AntiBot() (*big.Int, error) {
	return _Bt.Contract.AntiBot(&_Bt.CallOpts)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address ) view returns(uint256)
func (_Bt *BtCaller) BalanceOf(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Bt.contract.Call(opts, &out, "balanceOf", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address ) view returns(uint256)
func (_Bt *BtSession) BalanceOf(arg0 common.Address) (*big.Int, error) {
	return _Bt.Contract.BalanceOf(&_Bt.CallOpts, arg0)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address ) view returns(uint256)
func (_Bt *BtCallerSession) BalanceOf(arg0 common.Address) (*big.Int, error) {
	return _Bt.Contract.BalanceOf(&_Bt.CallOpts, arg0)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_Bt *BtCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _Bt.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_Bt *BtSession) Decimals() (uint8, error) {
	return _Bt.Contract.Decimals(&_Bt.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_Bt *BtCallerSession) Decimals() (uint8, error) {
	return _Bt.Contract.Decimals(&_Bt.CallOpts)
}

// Epoch is a free data retrieval call binding the contract method 0x900cf0cf.
//
// Solidity: function epoch() view returns(uint256)
func (_Bt *BtCaller) Epoch(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Bt.contract.Call(opts, &out, "epoch")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Epoch is a free data retrieval call binding the contract method 0x900cf0cf.
//
// Solidity: function epoch() view returns(uint256)
func (_Bt *BtSession) Epoch() (*big.Int, error) {
	return _Bt.Contract.Epoch(&_Bt.CallOpts)
}

// Epoch is a free data retrieval call binding the contract method 0x900cf0cf.
//
// Solidity: function epoch() view returns(uint256)
func (_Bt *BtCallerSession) Epoch() (*big.Int, error) {
	return _Bt.Contract.Epoch(&_Bt.CallOpts)
}

// EpochAdd is a free data retrieval call binding the contract method 0xe1f69814.
//
// Solidity: function epoch_add() view returns(uint256)
func (_Bt *BtCaller) EpochAdd(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Bt.contract.Call(opts, &out, "epoch_add")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// EpochAdd is a free data retrieval call binding the contract method 0xe1f69814.
//
// Solidity: function epoch_add() view returns(uint256)
func (_Bt *BtSession) EpochAdd() (*big.Int, error) {
	return _Bt.Contract.EpochAdd(&_Bt.CallOpts)
}

// EpochAdd is a free data retrieval call binding the contract method 0xe1f69814.
//
// Solidity: function epoch_add() view returns(uint256)
func (_Bt *BtCallerSession) EpochAdd() (*big.Int, error) {
	return _Bt.Contract.EpochAdd(&_Bt.CallOpts)
}

// EpochBase is a free data retrieval call binding the contract method 0xaf2872c7.
//
// Solidity: function epoch_base() view returns(uint256)
func (_Bt *BtCaller) EpochBase(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Bt.contract.Call(opts, &out, "epoch_base")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// EpochBase is a free data retrieval call binding the contract method 0xaf2872c7.
//
// Solidity: function epoch_base() view returns(uint256)
func (_Bt *BtSession) EpochBase() (*big.Int, error) {
	return _Bt.Contract.EpochBase(&_Bt.CallOpts)
}

// EpochBase is a free data retrieval call binding the contract method 0xaf2872c7.
//
// Solidity: function epoch_base() view returns(uint256)
func (_Bt *BtCallerSession) EpochBase() (*big.Int, error) {
	return _Bt.Contract.EpochBase(&_Bt.CallOpts)
}

// FreezeOf is a free data retrieval call binding the contract method 0xcd4217c1.
//
// Solidity: function freezeOf(address ) view returns(uint256)
func (_Bt *BtCaller) FreezeOf(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Bt.contract.Call(opts, &out, "freezeOf", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// FreezeOf is a free data retrieval call binding the contract method 0xcd4217c1.
//
// Solidity: function freezeOf(address ) view returns(uint256)
func (_Bt *BtSession) FreezeOf(arg0 common.Address) (*big.Int, error) {
	return _Bt.Contract.FreezeOf(&_Bt.CallOpts, arg0)
}

// FreezeOf is a free data retrieval call binding the contract method 0xcd4217c1.
//
// Solidity: function freezeOf(address ) view returns(uint256)
func (_Bt *BtCallerSession) FreezeOf(arg0 common.Address) (*big.Int, error) {
	return _Bt.Contract.FreezeOf(&_Bt.CallOpts, arg0)
}

// Invite is a free data retrieval call binding the contract method 0x4b77c468.
//
// Solidity: function invite(address ) view returns(address)
func (_Bt *BtCaller) Invite(opts *bind.CallOpts, arg0 common.Address) (common.Address, error) {
	var out []interface{}
	err := _Bt.contract.Call(opts, &out, "invite", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Invite is a free data retrieval call binding the contract method 0x4b77c468.
//
// Solidity: function invite(address ) view returns(address)
func (_Bt *BtSession) Invite(arg0 common.Address) (common.Address, error) {
	return _Bt.Contract.Invite(&_Bt.CallOpts, arg0)
}

// Invite is a free data retrieval call binding the contract method 0x4b77c468.
//
// Solidity: function invite(address ) view returns(address)
func (_Bt *BtCallerSession) Invite(arg0 common.Address) (common.Address, error) {
	return _Bt.Contract.Invite(&_Bt.CallOpts, arg0)
}

// InviteCount is a free data retrieval call binding the contract method 0x57b24e6b.
//
// Solidity: function inviteCount(address ) view returns(uint256)
func (_Bt *BtCaller) InviteCount(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Bt.contract.Call(opts, &out, "inviteCount", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// InviteCount is a free data retrieval call binding the contract method 0x57b24e6b.
//
// Solidity: function inviteCount(address ) view returns(uint256)
func (_Bt *BtSession) InviteCount(arg0 common.Address) (*big.Int, error) {
	return _Bt.Contract.InviteCount(&_Bt.CallOpts, arg0)
}

// InviteCount is a free data retrieval call binding the contract method 0x57b24e6b.
//
// Solidity: function inviteCount(address ) view returns(uint256)
func (_Bt *BtCallerSession) InviteCount(arg0 common.Address) (*big.Int, error) {
	return _Bt.Contract.InviteCount(&_Bt.CallOpts, arg0)
}

// IsAirdrop is a free data retrieval call binding the contract method 0xaae65119.
//
// Solidity: function is_airdrop() view returns(bool)
func (_Bt *BtCaller) IsAirdrop(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Bt.contract.Call(opts, &out, "is_airdrop")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsAirdrop is a free data retrieval call binding the contract method 0xaae65119.
//
// Solidity: function is_airdrop() view returns(bool)
func (_Bt *BtSession) IsAirdrop() (bool, error) {
	return _Bt.Contract.IsAirdrop(&_Bt.CallOpts)
}

// IsAirdrop is a free data retrieval call binding the contract method 0xaae65119.
//
// Solidity: function is_airdrop() view returns(bool)
func (_Bt *BtCallerSession) IsAirdrop() (bool, error) {
	return _Bt.Contract.IsAirdrop(&_Bt.CallOpts)
}

// LastMiner is a free data retrieval call binding the contract method 0x1f251a62.
//
// Solidity: function last_miner(address ) view returns(uint256)
func (_Bt *BtCaller) LastMiner(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Bt.contract.Call(opts, &out, "last_miner", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LastMiner is a free data retrieval call binding the contract method 0x1f251a62.
//
// Solidity: function last_miner(address ) view returns(uint256)
func (_Bt *BtSession) LastMiner(arg0 common.Address) (*big.Int, error) {
	return _Bt.Contract.LastMiner(&_Bt.CallOpts, arg0)
}

// LastMiner is a free data retrieval call binding the contract method 0x1f251a62.
//
// Solidity: function last_miner(address ) view returns(uint256)
func (_Bt *BtCallerSession) LastMiner(arg0 common.Address) (*big.Int, error) {
	return _Bt.Contract.LastMiner(&_Bt.CallOpts, arg0)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Bt *BtCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Bt.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Bt *BtSession) Name() (string, error) {
	return _Bt.Contract.Name(&_Bt.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Bt *BtCallerSession) Name() (string, error) {
	return _Bt.Contract.Name(&_Bt.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Bt *BtCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Bt.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Bt *BtSession) Owner() (common.Address, error) {
	return _Bt.Contract.Owner(&_Bt.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Bt *BtCallerSession) Owner() (common.Address, error) {
	return _Bt.Contract.Owner(&_Bt.CallOpts)
}

// Power is a free data retrieval call binding the contract method 0x503371a5.
//
// Solidity: function power(address ) view returns(uint256)
func (_Bt *BtCaller) Power(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Bt.contract.Call(opts, &out, "power", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Power is a free data retrieval call binding the contract method 0x503371a5.
//
// Solidity: function power(address ) view returns(uint256)
func (_Bt *BtSession) Power(arg0 common.Address) (*big.Int, error) {
	return _Bt.Contract.Power(&_Bt.CallOpts, arg0)
}

// Power is a free data retrieval call binding the contract method 0x503371a5.
//
// Solidity: function power(address ) view returns(uint256)
func (_Bt *BtCallerSession) Power(arg0 common.Address) (*big.Int, error) {
	return _Bt.Contract.Power(&_Bt.CallOpts, arg0)
}

// RewardCount is a free data retrieval call binding the contract method 0xb5af960d.
//
// Solidity: function rewardCount(address ) view returns(uint256)
func (_Bt *BtCaller) RewardCount(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Bt.contract.Call(opts, &out, "rewardCount", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// RewardCount is a free data retrieval call binding the contract method 0xb5af960d.
//
// Solidity: function rewardCount(address ) view returns(uint256)
func (_Bt *BtSession) RewardCount(arg0 common.Address) (*big.Int, error) {
	return _Bt.Contract.RewardCount(&_Bt.CallOpts, arg0)
}

// RewardCount is a free data retrieval call binding the contract method 0xb5af960d.
//
// Solidity: function rewardCount(address ) view returns(uint256)
func (_Bt *BtCallerSession) RewardCount(arg0 common.Address) (*big.Int, error) {
	return _Bt.Contract.RewardCount(&_Bt.CallOpts, arg0)
}

// StartTime is a free data retrieval call binding the contract method 0x834ee417.
//
// Solidity: function start_time() view returns(uint256)
func (_Bt *BtCaller) StartTime(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Bt.contract.Call(opts, &out, "start_time")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// StartTime is a free data retrieval call binding the contract method 0x834ee417.
//
// Solidity: function start_time() view returns(uint256)
func (_Bt *BtSession) StartTime() (*big.Int, error) {
	return _Bt.Contract.StartTime(&_Bt.CallOpts)
}

// StartTime is a free data retrieval call binding the contract method 0x834ee417.
//
// Solidity: function start_time() view returns(uint256)
func (_Bt *BtCallerSession) StartTime() (*big.Int, error) {
	return _Bt.Contract.StartTime(&_Bt.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Bt *BtCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Bt.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Bt *BtSession) Symbol() (string, error) {
	return _Bt.Contract.Symbol(&_Bt.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Bt *BtCallerSession) Symbol() (string, error) {
	return _Bt.Contract.Symbol(&_Bt.CallOpts)
}

// TotalPower is a free data retrieval call binding the contract method 0xdb3ad22c.
//
// Solidity: function totalPower() view returns(uint256)
func (_Bt *BtCaller) TotalPower(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Bt.contract.Call(opts, &out, "totalPower")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalPower is a free data retrieval call binding the contract method 0xdb3ad22c.
//
// Solidity: function totalPower() view returns(uint256)
func (_Bt *BtSession) TotalPower() (*big.Int, error) {
	return _Bt.Contract.TotalPower(&_Bt.CallOpts)
}

// TotalPower is a free data retrieval call binding the contract method 0xdb3ad22c.
//
// Solidity: function totalPower() view returns(uint256)
func (_Bt *BtCallerSession) TotalPower() (*big.Int, error) {
	return _Bt.Contract.TotalPower(&_Bt.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Bt *BtCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Bt.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Bt *BtSession) TotalSupply() (*big.Int, error) {
	return _Bt.Contract.TotalSupply(&_Bt.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Bt *BtCallerSession) TotalSupply() (*big.Int, error) {
	return _Bt.Contract.TotalSupply(&_Bt.CallOpts)
}

// TotalUsersAmount is a free data retrieval call binding the contract method 0xdd94a29d.
//
// Solidity: function totalUsersAmount() view returns(uint256)
func (_Bt *BtCaller) TotalUsersAmount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Bt.contract.Call(opts, &out, "totalUsersAmount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalUsersAmount is a free data retrieval call binding the contract method 0xdd94a29d.
//
// Solidity: function totalUsersAmount() view returns(uint256)
func (_Bt *BtSession) TotalUsersAmount() (*big.Int, error) {
	return _Bt.Contract.TotalUsersAmount(&_Bt.CallOpts)
}

// TotalUsersAmount is a free data retrieval call binding the contract method 0xdd94a29d.
//
// Solidity: function totalUsersAmount() view returns(uint256)
func (_Bt *BtCallerSession) TotalUsersAmount() (*big.Int, error) {
	return _Bt.Contract.TotalUsersAmount(&_Bt.CallOpts)
}

// TransferFee is a free data retrieval call binding the contract method 0x9e2f3836.
//
// Solidity: function transfer_fee(address _from, uint256 _value) view returns(uint256 fee)
func (_Bt *BtCaller) TransferFee(opts *bind.CallOpts, _from common.Address, _value *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Bt.contract.Call(opts, &out, "transfer_fee", _from, _value)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TransferFee is a free data retrieval call binding the contract method 0x9e2f3836.
//
// Solidity: function transfer_fee(address _from, uint256 _value) view returns(uint256 fee)
func (_Bt *BtSession) TransferFee(_from common.Address, _value *big.Int) (*big.Int, error) {
	return _Bt.Contract.TransferFee(&_Bt.CallOpts, _from, _value)
}

// TransferFee is a free data retrieval call binding the contract method 0x9e2f3836.
//
// Solidity: function transfer_fee(address _from, uint256 _value) view returns(uint256 fee)
func (_Bt *BtCallerSession) TransferFee(_from common.Address, _value *big.Int) (*big.Int, error) {
	return _Bt.Contract.TransferFee(&_Bt.CallOpts, _from, _value)
}

// Airdrop is a paid mutator transaction binding the contract method 0x729ad39e.
//
// Solidity: function airdrop(address[] address_array) returns()
func (_Bt *BtTransactor) Airdrop(opts *bind.TransactOpts, address_array []common.Address) (*types.Transaction, error) {
	return _Bt.contract.Transact(opts, "airdrop", address_array)
}

// Airdrop is a paid mutator transaction binding the contract method 0x729ad39e.
//
// Solidity: function airdrop(address[] address_array) returns()
func (_Bt *BtSession) Airdrop(address_array []common.Address) (*types.Transaction, error) {
	return _Bt.Contract.Airdrop(&_Bt.TransactOpts, address_array)
}

// Airdrop is a paid mutator transaction binding the contract method 0x729ad39e.
//
// Solidity: function airdrop(address[] address_array) returns()
func (_Bt *BtTransactorSession) Airdrop(address_array []common.Address) (*types.Transaction, error) {
	return _Bt.Contract.Airdrop(&_Bt.TransactOpts, address_array)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address _spender, uint256 _value) returns(bool success)
func (_Bt *BtTransactor) Approve(opts *bind.TransactOpts, _spender common.Address, _value *big.Int) (*types.Transaction, error) {
	return _Bt.contract.Transact(opts, "approve", _spender, _value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address _spender, uint256 _value) returns(bool success)
func (_Bt *BtSession) Approve(_spender common.Address, _value *big.Int) (*types.Transaction, error) {
	return _Bt.Contract.Approve(&_Bt.TransactOpts, _spender, _value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address _spender, uint256 _value) returns(bool success)
func (_Bt *BtTransactorSession) Approve(_spender common.Address, _value *big.Int) (*types.Transaction, error) {
	return _Bt.Contract.Approve(&_Bt.TransactOpts, _spender, _value)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(uint256 _value) returns(bool success)
func (_Bt *BtTransactor) Burn(opts *bind.TransactOpts, _value *big.Int) (*types.Transaction, error) {
	return _Bt.contract.Transact(opts, "burn", _value)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(uint256 _value) returns(bool success)
func (_Bt *BtSession) Burn(_value *big.Int) (*types.Transaction, error) {
	return _Bt.Contract.Burn(&_Bt.TransactOpts, _value)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(uint256 _value) returns(bool success)
func (_Bt *BtTransactorSession) Burn(_value *big.Int) (*types.Transaction, error) {
	return _Bt.Contract.Burn(&_Bt.TransactOpts, _value)
}

// Freeze is a paid mutator transaction binding the contract method 0xd7a78db8.
//
// Solidity: function freeze(uint256 _value) returns(bool success)
func (_Bt *BtTransactor) Freeze(opts *bind.TransactOpts, _value *big.Int) (*types.Transaction, error) {
	return _Bt.contract.Transact(opts, "freeze", _value)
}

// Freeze is a paid mutator transaction binding the contract method 0xd7a78db8.
//
// Solidity: function freeze(uint256 _value) returns(bool success)
func (_Bt *BtSession) Freeze(_value *big.Int) (*types.Transaction, error) {
	return _Bt.Contract.Freeze(&_Bt.TransactOpts, _value)
}

// Freeze is a paid mutator transaction binding the contract method 0xd7a78db8.
//
// Solidity: function freeze(uint256 _value) returns(bool success)
func (_Bt *BtTransactorSession) Freeze(_value *big.Int) (*types.Transaction, error) {
	return _Bt.Contract.Freeze(&_Bt.TransactOpts, _value)
}

// Mint is a paid mutator transaction binding the contract method 0x1249c58b.
//
// Solidity: function mint() returns(bool success)
func (_Bt *BtTransactor) Mint(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Bt.contract.Transact(opts, "mint")
}

// Mint is a paid mutator transaction binding the contract method 0x1249c58b.
//
// Solidity: function mint() returns(bool success)
func (_Bt *BtSession) Mint() (*types.Transaction, error) {
	return _Bt.Contract.Mint(&_Bt.TransactOpts)
}

// Mint is a paid mutator transaction binding the contract method 0x1249c58b.
//
// Solidity: function mint() returns(bool success)
func (_Bt *BtTransactorSession) Mint() (*types.Transaction, error) {
	return _Bt.Contract.Mint(&_Bt.TransactOpts)
}

// Registration is a paid mutator transaction binding the contract method 0x0840605a.
//
// Solidity: function registration(address invite_address) returns(bool success)
func (_Bt *BtTransactor) Registration(opts *bind.TransactOpts, invite_address common.Address) (*types.Transaction, error) {
	return _Bt.contract.Transact(opts, "registration", invite_address)
}

// Registration is a paid mutator transaction binding the contract method 0x0840605a.
//
// Solidity: function registration(address invite_address) returns(bool success)
func (_Bt *BtSession) Registration(invite_address common.Address) (*types.Transaction, error) {
	return _Bt.Contract.Registration(&_Bt.TransactOpts, invite_address)
}

// Registration is a paid mutator transaction binding the contract method 0x0840605a.
//
// Solidity: function registration(address invite_address) returns(bool success)
func (_Bt *BtTransactorSession) Registration(invite_address common.Address) (*types.Transaction, error) {
	return _Bt.Contract.Registration(&_Bt.TransactOpts, invite_address)
}

// SetOwner is a paid mutator transaction binding the contract method 0x13af4035.
//
// Solidity: function setOwner(address new_owner) returns()
func (_Bt *BtTransactor) SetOwner(opts *bind.TransactOpts, new_owner common.Address) (*types.Transaction, error) {
	return _Bt.contract.Transact(opts, "setOwner", new_owner)
}

// SetOwner is a paid mutator transaction binding the contract method 0x13af4035.
//
// Solidity: function setOwner(address new_owner) returns()
func (_Bt *BtSession) SetOwner(new_owner common.Address) (*types.Transaction, error) {
	return _Bt.Contract.SetOwner(&_Bt.TransactOpts, new_owner)
}

// SetOwner is a paid mutator transaction binding the contract method 0x13af4035.
//
// Solidity: function setOwner(address new_owner) returns()
func (_Bt *BtTransactorSession) SetOwner(new_owner common.Address) (*types.Transaction, error) {
	return _Bt.Contract.SetOwner(&_Bt.TransactOpts, new_owner)
}

// SetAntiBot is a paid mutator transaction binding the contract method 0xe216a2dd.
//
// Solidity: function set_anti_bot(uint256 _value) returns()
func (_Bt *BtTransactor) SetAntiBot(opts *bind.TransactOpts, _value *big.Int) (*types.Transaction, error) {
	return _Bt.contract.Transact(opts, "set_anti_bot", _value)
}

// SetAntiBot is a paid mutator transaction binding the contract method 0xe216a2dd.
//
// Solidity: function set_anti_bot(uint256 _value) returns()
func (_Bt *BtSession) SetAntiBot(_value *big.Int) (*types.Transaction, error) {
	return _Bt.Contract.SetAntiBot(&_Bt.TransactOpts, _value)
}

// SetAntiBot is a paid mutator transaction binding the contract method 0xe216a2dd.
//
// Solidity: function set_anti_bot(uint256 _value) returns()
func (_Bt *BtTransactorSession) SetAntiBot(_value *big.Int) (*types.Transaction, error) {
	return _Bt.Contract.SetAntiBot(&_Bt.TransactOpts, _value)
}

// StopAirdrop is a paid mutator transaction binding the contract method 0x6ad82705.
//
// Solidity: function stop_airdrop() returns()
func (_Bt *BtTransactor) StopAirdrop(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Bt.contract.Transact(opts, "stop_airdrop")
}

// StopAirdrop is a paid mutator transaction binding the contract method 0x6ad82705.
//
// Solidity: function stop_airdrop() returns()
func (_Bt *BtSession) StopAirdrop() (*types.Transaction, error) {
	return _Bt.Contract.StopAirdrop(&_Bt.TransactOpts)
}

// StopAirdrop is a paid mutator transaction binding the contract method 0x6ad82705.
//
// Solidity: function stop_airdrop() returns()
func (_Bt *BtTransactorSession) StopAirdrop() (*types.Transaction, error) {
	return _Bt.Contract.StopAirdrop(&_Bt.TransactOpts)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address _to, uint256 _value) returns(bool success)
func (_Bt *BtTransactor) Transfer(opts *bind.TransactOpts, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _Bt.contract.Transact(opts, "transfer", _to, _value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address _to, uint256 _value) returns(bool success)
func (_Bt *BtSession) Transfer(_to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _Bt.Contract.Transfer(&_Bt.TransactOpts, _to, _value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address _to, uint256 _value) returns(bool success)
func (_Bt *BtTransactorSession) Transfer(_to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _Bt.Contract.Transfer(&_Bt.TransactOpts, _to, _value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address _from, address _to, uint256 _value) returns(bool success)
func (_Bt *BtTransactor) TransferFrom(opts *bind.TransactOpts, _from common.Address, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _Bt.contract.Transact(opts, "transferFrom", _from, _to, _value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address _from, address _to, uint256 _value) returns(bool success)
func (_Bt *BtSession) TransferFrom(_from common.Address, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _Bt.Contract.TransferFrom(&_Bt.TransactOpts, _from, _to, _value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address _from, address _to, uint256 _value) returns(bool success)
func (_Bt *BtTransactorSession) TransferFrom(_from common.Address, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _Bt.Contract.TransferFrom(&_Bt.TransactOpts, _from, _to, _value)
}

// Unfreeze is a paid mutator transaction binding the contract method 0x6623fc46.
//
// Solidity: function unfreeze(uint256 _value) returns(bool success)
func (_Bt *BtTransactor) Unfreeze(opts *bind.TransactOpts, _value *big.Int) (*types.Transaction, error) {
	return _Bt.contract.Transact(opts, "unfreeze", _value)
}

// Unfreeze is a paid mutator transaction binding the contract method 0x6623fc46.
//
// Solidity: function unfreeze(uint256 _value) returns(bool success)
func (_Bt *BtSession) Unfreeze(_value *big.Int) (*types.Transaction, error) {
	return _Bt.Contract.Unfreeze(&_Bt.TransactOpts, _value)
}

// Unfreeze is a paid mutator transaction binding the contract method 0x6623fc46.
//
// Solidity: function unfreeze(uint256 _value) returns(bool success)
func (_Bt *BtTransactorSession) Unfreeze(_value *big.Int) (*types.Transaction, error) {
	return _Bt.Contract.Unfreeze(&_Bt.TransactOpts, _value)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 amount) returns()
func (_Bt *BtTransactor) Withdraw(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _Bt.contract.Transact(opts, "withdraw", amount)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 amount) returns()
func (_Bt *BtSession) Withdraw(amount *big.Int) (*types.Transaction, error) {
	return _Bt.Contract.Withdraw(&_Bt.TransactOpts, amount)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 amount) returns()
func (_Bt *BtTransactorSession) Withdraw(amount *big.Int) (*types.Transaction, error) {
	return _Bt.Contract.Withdraw(&_Bt.TransactOpts, amount)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Bt *BtTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Bt.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Bt *BtSession) Receive() (*types.Transaction, error) {
	return _Bt.Contract.Receive(&_Bt.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Bt *BtTransactorSession) Receive() (*types.Transaction, error) {
	return _Bt.Contract.Receive(&_Bt.TransactOpts)
}

// BtBurnIterator is returned from FilterBurn and is used to iterate over the raw logs and unpacked data for Burn events raised by the Bt contract.
type BtBurnIterator struct {
	Event *BtBurn // Event containing the contract specifics and raw log

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
func (it *BtBurnIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BtBurn)
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
		it.Event = new(BtBurn)
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
func (it *BtBurnIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BtBurnIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BtBurn represents a Burn event raised by the Bt contract.
type BtBurn struct {
	From  common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterBurn is a free log retrieval operation binding the contract event 0xcc16f5dbb4873280815c1ee09dbd06736cffcc184412cf7a71a0fdb75d397ca5.
//
// Solidity: event Burn(address indexed from, uint256 value)
func (_Bt *BtFilterer) FilterBurn(opts *bind.FilterOpts, from []common.Address) (*BtBurnIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}

	logs, sub, err := _Bt.contract.FilterLogs(opts, "Burn", fromRule)
	if err != nil {
		return nil, err
	}
	return &BtBurnIterator{contract: _Bt.contract, event: "Burn", logs: logs, sub: sub}, nil
}

// WatchBurn is a free log subscription operation binding the contract event 0xcc16f5dbb4873280815c1ee09dbd06736cffcc184412cf7a71a0fdb75d397ca5.
//
// Solidity: event Burn(address indexed from, uint256 value)
func (_Bt *BtFilterer) WatchBurn(opts *bind.WatchOpts, sink chan<- *BtBurn, from []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}

	logs, sub, err := _Bt.contract.WatchLogs(opts, "Burn", fromRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BtBurn)
				if err := _Bt.contract.UnpackLog(event, "Burn", log); err != nil {
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
func (_Bt *BtFilterer) ParseBurn(log types.Log) (*BtBurn, error) {
	event := new(BtBurn)
	if err := _Bt.contract.UnpackLog(event, "Burn", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BtFreezeIterator is returned from FilterFreeze and is used to iterate over the raw logs and unpacked data for Freeze events raised by the Bt contract.
type BtFreezeIterator struct {
	Event *BtFreeze // Event containing the contract specifics and raw log

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
func (it *BtFreezeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BtFreeze)
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
		it.Event = new(BtFreeze)
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
func (it *BtFreezeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BtFreezeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BtFreeze represents a Freeze event raised by the Bt contract.
type BtFreeze struct {
	From  common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterFreeze is a free log retrieval operation binding the contract event 0xf97a274face0b5517365ad396b1fdba6f68bd3135ef603e44272adba3af5a1e0.
//
// Solidity: event Freeze(address indexed from, uint256 value)
func (_Bt *BtFilterer) FilterFreeze(opts *bind.FilterOpts, from []common.Address) (*BtFreezeIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}

	logs, sub, err := _Bt.contract.FilterLogs(opts, "Freeze", fromRule)
	if err != nil {
		return nil, err
	}
	return &BtFreezeIterator{contract: _Bt.contract, event: "Freeze", logs: logs, sub: sub}, nil
}

// WatchFreeze is a free log subscription operation binding the contract event 0xf97a274face0b5517365ad396b1fdba6f68bd3135ef603e44272adba3af5a1e0.
//
// Solidity: event Freeze(address indexed from, uint256 value)
func (_Bt *BtFilterer) WatchFreeze(opts *bind.WatchOpts, sink chan<- *BtFreeze, from []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}

	logs, sub, err := _Bt.contract.WatchLogs(opts, "Freeze", fromRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BtFreeze)
				if err := _Bt.contract.UnpackLog(event, "Freeze", log); err != nil {
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
func (_Bt *BtFilterer) ParseFreeze(log types.Log) (*BtFreeze, error) {
	event := new(BtFreeze)
	if err := _Bt.contract.UnpackLog(event, "Freeze", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BtMintedIterator is returned from FilterMinted and is used to iterate over the raw logs and unpacked data for Minted events raised by the Bt contract.
type BtMintedIterator struct {
	Event *BtMinted // Event containing the contract specifics and raw log

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
func (it *BtMintedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BtMinted)
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
		it.Event = new(BtMinted)
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
func (it *BtMintedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BtMintedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BtMinted represents a Minted event raised by the Bt contract.
type BtMinted struct {
	Operator common.Address
	To       common.Address
	Amount   *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterMinted is a free log retrieval operation binding the contract event 0x9d228d69b5fdb8d273a2336f8fb8612d039631024ea9bf09c424a9503aa078f0.
//
// Solidity: event Minted(address indexed operator, address indexed to, uint256 amount)
func (_Bt *BtFilterer) FilterMinted(opts *bind.FilterOpts, operator []common.Address, to []common.Address) (*BtMintedIterator, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Bt.contract.FilterLogs(opts, "Minted", operatorRule, toRule)
	if err != nil {
		return nil, err
	}
	return &BtMintedIterator{contract: _Bt.contract, event: "Minted", logs: logs, sub: sub}, nil
}

// WatchMinted is a free log subscription operation binding the contract event 0x9d228d69b5fdb8d273a2336f8fb8612d039631024ea9bf09c424a9503aa078f0.
//
// Solidity: event Minted(address indexed operator, address indexed to, uint256 amount)
func (_Bt *BtFilterer) WatchMinted(opts *bind.WatchOpts, sink chan<- *BtMinted, operator []common.Address, to []common.Address) (event.Subscription, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Bt.contract.WatchLogs(opts, "Minted", operatorRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BtMinted)
				if err := _Bt.contract.UnpackLog(event, "Minted", log); err != nil {
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
func (_Bt *BtFilterer) ParseMinted(log types.Log) (*BtMinted, error) {
	event := new(BtMinted)
	if err := _Bt.contract.UnpackLog(event, "Minted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BtTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the Bt contract.
type BtTransferIterator struct {
	Event *BtTransfer // Event containing the contract specifics and raw log

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
func (it *BtTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BtTransfer)
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
		it.Event = new(BtTransfer)
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
func (it *BtTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BtTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BtTransfer represents a Transfer event raised by the Bt contract.
type BtTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_Bt *BtFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*BtTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Bt.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &BtTransferIterator{contract: _Bt.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_Bt *BtFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *BtTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Bt.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BtTransfer)
				if err := _Bt.contract.UnpackLog(event, "Transfer", log); err != nil {
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
func (_Bt *BtFilterer) ParseTransfer(log types.Log) (*BtTransfer, error) {
	event := new(BtTransfer)
	if err := _Bt.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BtUnfreezeIterator is returned from FilterUnfreeze and is used to iterate over the raw logs and unpacked data for Unfreeze events raised by the Bt contract.
type BtUnfreezeIterator struct {
	Event *BtUnfreeze // Event containing the contract specifics and raw log

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
func (it *BtUnfreezeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BtUnfreeze)
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
		it.Event = new(BtUnfreeze)
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
func (it *BtUnfreezeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BtUnfreezeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BtUnfreeze represents a Unfreeze event raised by the Bt contract.
type BtUnfreeze struct {
	From  common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterUnfreeze is a free log retrieval operation binding the contract event 0x2cfce4af01bcb9d6cf6c84ee1b7c491100b8695368264146a94d71e10a63083f.
//
// Solidity: event Unfreeze(address indexed from, uint256 value)
func (_Bt *BtFilterer) FilterUnfreeze(opts *bind.FilterOpts, from []common.Address) (*BtUnfreezeIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}

	logs, sub, err := _Bt.contract.FilterLogs(opts, "Unfreeze", fromRule)
	if err != nil {
		return nil, err
	}
	return &BtUnfreezeIterator{contract: _Bt.contract, event: "Unfreeze", logs: logs, sub: sub}, nil
}

// WatchUnfreeze is a free log subscription operation binding the contract event 0x2cfce4af01bcb9d6cf6c84ee1b7c491100b8695368264146a94d71e10a63083f.
//
// Solidity: event Unfreeze(address indexed from, uint256 value)
func (_Bt *BtFilterer) WatchUnfreeze(opts *bind.WatchOpts, sink chan<- *BtUnfreeze, from []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}

	logs, sub, err := _Bt.contract.WatchLogs(opts, "Unfreeze", fromRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BtUnfreeze)
				if err := _Bt.contract.UnpackLog(event, "Unfreeze", log); err != nil {
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

// ParseUnfreeze is a log parse operation binding the contract event 0x2cfce4af01bcb9d6cf6c84ee1b7c491100b8695368264146a94d71e10a63083f.
//
// Solidity: event Unfreeze(address indexed from, uint256 value)
func (_Bt *BtFilterer) ParseUnfreeze(log types.Log) (*BtUnfreeze, error) {
	event := new(BtUnfreeze)
	if err := _Bt.contract.UnpackLog(event, "Unfreeze", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
