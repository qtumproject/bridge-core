// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package tokenerc20

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

// Tokenerc20MetaData contains all meta data concerning the Tokenerc20 contract.
var Tokenerc20MetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"ticker\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"_signers\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"hash\",\"type\":\"string\"}],\"name\":\"HashAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"burn\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_hash\",\"type\":\"string\"}],\"name\":\"containsHash\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"subtractedValue\",\"type\":\"uint256\"}],\"name\":\"decreaseAllowance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getOwner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"addedValue\",\"type\":\"uint256\"}],\"name\":\"increaseAllowance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"_txHash\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"_receiver\",\"type\":\"address\"},{\"internalType\":\"bytes32[]\",\"name\":\"_r\",\"type\":\"bytes32[]\"},{\"internalType\":\"bytes32[]\",\"name\":\"_s\",\"type\":\"bytes32[]\"},{\"internalType\":\"uint8[]\",\"name\":\"_v\",\"type\":\"uint8[]\"}],\"name\":\"mint\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_receiver\",\"type\":\"address\"}],\"name\":\"mint\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"signersRep\",\"outputs\":[{\"internalType\":\"contractISignersRepository\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// Tokenerc20ABI is the input ABI used to generate the binding from.
// Deprecated: Use Tokenerc20MetaData.ABI instead.
var Tokenerc20ABI = Tokenerc20MetaData.ABI

// Tokenerc20 is an auto generated Go binding around an Ethereum contract.
type Tokenerc20 struct {
	Tokenerc20Caller     // Read-only binding to the contract
	Tokenerc20Transactor // Write-only binding to the contract
	Tokenerc20Filterer   // Log filterer for contract events
}

// Tokenerc20Caller is an auto generated read-only Go binding around an Ethereum contract.
type Tokenerc20Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Tokenerc20Transactor is an auto generated write-only Go binding around an Ethereum contract.
type Tokenerc20Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Tokenerc20Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type Tokenerc20Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Tokenerc20Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type Tokenerc20Session struct {
	Contract     *Tokenerc20       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// Tokenerc20CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type Tokenerc20CallerSession struct {
	Contract *Tokenerc20Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// Tokenerc20TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type Tokenerc20TransactorSession struct {
	Contract     *Tokenerc20Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// Tokenerc20Raw is an auto generated low-level Go binding around an Ethereum contract.
type Tokenerc20Raw struct {
	Contract *Tokenerc20 // Generic contract binding to access the raw methods on
}

// Tokenerc20CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type Tokenerc20CallerRaw struct {
	Contract *Tokenerc20Caller // Generic read-only contract binding to access the raw methods on
}

// Tokenerc20TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type Tokenerc20TransactorRaw struct {
	Contract *Tokenerc20Transactor // Generic write-only contract binding to access the raw methods on
}

// NewTokenerc20 creates a new instance of Tokenerc20, bound to a specific deployed contract.
func NewTokenerc20(address common.Address, backend bind.ContractBackend) (*Tokenerc20, error) {
	contract, err := bindTokenerc20(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Tokenerc20{Tokenerc20Caller: Tokenerc20Caller{contract: contract}, Tokenerc20Transactor: Tokenerc20Transactor{contract: contract}, Tokenerc20Filterer: Tokenerc20Filterer{contract: contract}}, nil
}

// NewTokenerc20Caller creates a new read-only instance of Tokenerc20, bound to a specific deployed contract.
func NewTokenerc20Caller(address common.Address, caller bind.ContractCaller) (*Tokenerc20Caller, error) {
	contract, err := bindTokenerc20(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &Tokenerc20Caller{contract: contract}, nil
}

// NewTokenerc20Transactor creates a new write-only instance of Tokenerc20, bound to a specific deployed contract.
func NewTokenerc20Transactor(address common.Address, transactor bind.ContractTransactor) (*Tokenerc20Transactor, error) {
	contract, err := bindTokenerc20(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &Tokenerc20Transactor{contract: contract}, nil
}

// NewTokenerc20Filterer creates a new log filterer instance of Tokenerc20, bound to a specific deployed contract.
func NewTokenerc20Filterer(address common.Address, filterer bind.ContractFilterer) (*Tokenerc20Filterer, error) {
	contract, err := bindTokenerc20(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &Tokenerc20Filterer{contract: contract}, nil
}

// bindTokenerc20 binds a generic wrapper to an already deployed contract.
func bindTokenerc20(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(Tokenerc20ABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Tokenerc20 *Tokenerc20Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Tokenerc20.Contract.Tokenerc20Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Tokenerc20 *Tokenerc20Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Tokenerc20.Contract.Tokenerc20Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Tokenerc20 *Tokenerc20Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Tokenerc20.Contract.Tokenerc20Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Tokenerc20 *Tokenerc20CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Tokenerc20.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Tokenerc20 *Tokenerc20TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Tokenerc20.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Tokenerc20 *Tokenerc20TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Tokenerc20.Contract.contract.Transact(opts, method, params...)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_Tokenerc20 *Tokenerc20Caller) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Tokenerc20.contract.Call(opts, &out, "allowance", owner, spender)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_Tokenerc20 *Tokenerc20Session) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _Tokenerc20.Contract.Allowance(&_Tokenerc20.CallOpts, owner, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_Tokenerc20 *Tokenerc20CallerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _Tokenerc20.Contract.Allowance(&_Tokenerc20.CallOpts, owner, spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_Tokenerc20 *Tokenerc20Caller) BalanceOf(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Tokenerc20.contract.Call(opts, &out, "balanceOf", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_Tokenerc20 *Tokenerc20Session) BalanceOf(account common.Address) (*big.Int, error) {
	return _Tokenerc20.Contract.BalanceOf(&_Tokenerc20.CallOpts, account)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_Tokenerc20 *Tokenerc20CallerSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _Tokenerc20.Contract.BalanceOf(&_Tokenerc20.CallOpts, account)
}

// ContainsHash is a free data retrieval call binding the contract method 0x6492f57a.
//
// Solidity: function containsHash(string _hash) view returns(bool)
func (_Tokenerc20 *Tokenerc20Caller) ContainsHash(opts *bind.CallOpts, _hash string) (bool, error) {
	var out []interface{}
	err := _Tokenerc20.contract.Call(opts, &out, "containsHash", _hash)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// ContainsHash is a free data retrieval call binding the contract method 0x6492f57a.
//
// Solidity: function containsHash(string _hash) view returns(bool)
func (_Tokenerc20 *Tokenerc20Session) ContainsHash(_hash string) (bool, error) {
	return _Tokenerc20.Contract.ContainsHash(&_Tokenerc20.CallOpts, _hash)
}

// ContainsHash is a free data retrieval call binding the contract method 0x6492f57a.
//
// Solidity: function containsHash(string _hash) view returns(bool)
func (_Tokenerc20 *Tokenerc20CallerSession) ContainsHash(_hash string) (bool, error) {
	return _Tokenerc20.Contract.ContainsHash(&_Tokenerc20.CallOpts, _hash)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_Tokenerc20 *Tokenerc20Caller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _Tokenerc20.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_Tokenerc20 *Tokenerc20Session) Decimals() (uint8, error) {
	return _Tokenerc20.Contract.Decimals(&_Tokenerc20.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_Tokenerc20 *Tokenerc20CallerSession) Decimals() (uint8, error) {
	return _Tokenerc20.Contract.Decimals(&_Tokenerc20.CallOpts)
}

// GetOwner is a free data retrieval call binding the contract method 0x893d20e8.
//
// Solidity: function getOwner() view returns(address)
func (_Tokenerc20 *Tokenerc20Caller) GetOwner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Tokenerc20.contract.Call(opts, &out, "getOwner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetOwner is a free data retrieval call binding the contract method 0x893d20e8.
//
// Solidity: function getOwner() view returns(address)
func (_Tokenerc20 *Tokenerc20Session) GetOwner() (common.Address, error) {
	return _Tokenerc20.Contract.GetOwner(&_Tokenerc20.CallOpts)
}

// GetOwner is a free data retrieval call binding the contract method 0x893d20e8.
//
// Solidity: function getOwner() view returns(address)
func (_Tokenerc20 *Tokenerc20CallerSession) GetOwner() (common.Address, error) {
	return _Tokenerc20.Contract.GetOwner(&_Tokenerc20.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Tokenerc20 *Tokenerc20Caller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Tokenerc20.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Tokenerc20 *Tokenerc20Session) Name() (string, error) {
	return _Tokenerc20.Contract.Name(&_Tokenerc20.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Tokenerc20 *Tokenerc20CallerSession) Name() (string, error) {
	return _Tokenerc20.Contract.Name(&_Tokenerc20.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Tokenerc20 *Tokenerc20Caller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Tokenerc20.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Tokenerc20 *Tokenerc20Session) Owner() (common.Address, error) {
	return _Tokenerc20.Contract.Owner(&_Tokenerc20.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Tokenerc20 *Tokenerc20CallerSession) Owner() (common.Address, error) {
	return _Tokenerc20.Contract.Owner(&_Tokenerc20.CallOpts)
}

// SignersRep is a free data retrieval call binding the contract method 0x897b73cd.
//
// Solidity: function signersRep() view returns(address)
func (_Tokenerc20 *Tokenerc20Caller) SignersRep(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Tokenerc20.contract.Call(opts, &out, "signersRep")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// SignersRep is a free data retrieval call binding the contract method 0x897b73cd.
//
// Solidity: function signersRep() view returns(address)
func (_Tokenerc20 *Tokenerc20Session) SignersRep() (common.Address, error) {
	return _Tokenerc20.Contract.SignersRep(&_Tokenerc20.CallOpts)
}

// SignersRep is a free data retrieval call binding the contract method 0x897b73cd.
//
// Solidity: function signersRep() view returns(address)
func (_Tokenerc20 *Tokenerc20CallerSession) SignersRep() (common.Address, error) {
	return _Tokenerc20.Contract.SignersRep(&_Tokenerc20.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Tokenerc20 *Tokenerc20Caller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Tokenerc20.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Tokenerc20 *Tokenerc20Session) Symbol() (string, error) {
	return _Tokenerc20.Contract.Symbol(&_Tokenerc20.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Tokenerc20 *Tokenerc20CallerSession) Symbol() (string, error) {
	return _Tokenerc20.Contract.Symbol(&_Tokenerc20.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Tokenerc20 *Tokenerc20Caller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Tokenerc20.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Tokenerc20 *Tokenerc20Session) TotalSupply() (*big.Int, error) {
	return _Tokenerc20.Contract.TotalSupply(&_Tokenerc20.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Tokenerc20 *Tokenerc20CallerSession) TotalSupply() (*big.Int, error) {
	return _Tokenerc20.Contract.TotalSupply(&_Tokenerc20.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_Tokenerc20 *Tokenerc20Transactor) Approve(opts *bind.TransactOpts, spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Tokenerc20.contract.Transact(opts, "approve", spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_Tokenerc20 *Tokenerc20Session) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Tokenerc20.Contract.Approve(&_Tokenerc20.TransactOpts, spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_Tokenerc20 *Tokenerc20TransactorSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Tokenerc20.Contract.Approve(&_Tokenerc20.TransactOpts, spender, amount)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(uint256 _amount) returns(bool)
func (_Tokenerc20 *Tokenerc20Transactor) Burn(opts *bind.TransactOpts, _amount *big.Int) (*types.Transaction, error) {
	return _Tokenerc20.contract.Transact(opts, "burn", _amount)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(uint256 _amount) returns(bool)
func (_Tokenerc20 *Tokenerc20Session) Burn(_amount *big.Int) (*types.Transaction, error) {
	return _Tokenerc20.Contract.Burn(&_Tokenerc20.TransactOpts, _amount)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(uint256 _amount) returns(bool)
func (_Tokenerc20 *Tokenerc20TransactorSession) Burn(_amount *big.Int) (*types.Transaction, error) {
	return _Tokenerc20.Contract.Burn(&_Tokenerc20.TransactOpts, _amount)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_Tokenerc20 *Tokenerc20Transactor) DecreaseAllowance(opts *bind.TransactOpts, spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _Tokenerc20.contract.Transact(opts, "decreaseAllowance", spender, subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_Tokenerc20 *Tokenerc20Session) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _Tokenerc20.Contract.DecreaseAllowance(&_Tokenerc20.TransactOpts, spender, subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_Tokenerc20 *Tokenerc20TransactorSession) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _Tokenerc20.Contract.DecreaseAllowance(&_Tokenerc20.TransactOpts, spender, subtractedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_Tokenerc20 *Tokenerc20Transactor) IncreaseAllowance(opts *bind.TransactOpts, spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _Tokenerc20.contract.Transact(opts, "increaseAllowance", spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_Tokenerc20 *Tokenerc20Session) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _Tokenerc20.Contract.IncreaseAllowance(&_Tokenerc20.TransactOpts, spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_Tokenerc20 *Tokenerc20TransactorSession) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _Tokenerc20.Contract.IncreaseAllowance(&_Tokenerc20.TransactOpts, spender, addedValue)
}

// Mint is a paid mutator transaction binding the contract method 0x32a09b10.
//
// Solidity: function mint(uint256 _amount, string _txHash, address _receiver, bytes32[] _r, bytes32[] _s, uint8[] _v) returns(bool)
func (_Tokenerc20 *Tokenerc20Transactor) Mint(opts *bind.TransactOpts, _amount *big.Int, _txHash string, _receiver common.Address, _r [][32]byte, _s [][32]byte, _v []uint8) (*types.Transaction, error) {
	return _Tokenerc20.contract.Transact(opts, "mint", _amount, _txHash, _receiver, _r, _s, _v)
}

// Mint is a paid mutator transaction binding the contract method 0x32a09b10.
//
// Solidity: function mint(uint256 _amount, string _txHash, address _receiver, bytes32[] _r, bytes32[] _s, uint8[] _v) returns(bool)
func (_Tokenerc20 *Tokenerc20Session) Mint(_amount *big.Int, _txHash string, _receiver common.Address, _r [][32]byte, _s [][32]byte, _v []uint8) (*types.Transaction, error) {
	return _Tokenerc20.Contract.Mint(&_Tokenerc20.TransactOpts, _amount, _txHash, _receiver, _r, _s, _v)
}

// Mint is a paid mutator transaction binding the contract method 0x32a09b10.
//
// Solidity: function mint(uint256 _amount, string _txHash, address _receiver, bytes32[] _r, bytes32[] _s, uint8[] _v) returns(bool)
func (_Tokenerc20 *Tokenerc20TransactorSession) Mint(_amount *big.Int, _txHash string, _receiver common.Address, _r [][32]byte, _s [][32]byte, _v []uint8) (*types.Transaction, error) {
	return _Tokenerc20.Contract.Mint(&_Tokenerc20.TransactOpts, _amount, _txHash, _receiver, _r, _s, _v)
}

// Mint0 is a paid mutator transaction binding the contract method 0x94bf804d.
//
// Solidity: function mint(uint256 _amount, address _receiver) returns(bool)
func (_Tokenerc20 *Tokenerc20Transactor) Mint0(opts *bind.TransactOpts, _amount *big.Int, _receiver common.Address) (*types.Transaction, error) {
	return _Tokenerc20.contract.Transact(opts, "mint0", _amount, _receiver)
}

// Mint0 is a paid mutator transaction binding the contract method 0x94bf804d.
//
// Solidity: function mint(uint256 _amount, address _receiver) returns(bool)
func (_Tokenerc20 *Tokenerc20Session) Mint0(_amount *big.Int, _receiver common.Address) (*types.Transaction, error) {
	return _Tokenerc20.Contract.Mint0(&_Tokenerc20.TransactOpts, _amount, _receiver)
}

// Mint0 is a paid mutator transaction binding the contract method 0x94bf804d.
//
// Solidity: function mint(uint256 _amount, address _receiver) returns(bool)
func (_Tokenerc20 *Tokenerc20TransactorSession) Mint0(_amount *big.Int, _receiver common.Address) (*types.Transaction, error) {
	return _Tokenerc20.Contract.Mint0(&_Tokenerc20.TransactOpts, _amount, _receiver)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Tokenerc20 *Tokenerc20Transactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Tokenerc20.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Tokenerc20 *Tokenerc20Session) RenounceOwnership() (*types.Transaction, error) {
	return _Tokenerc20.Contract.RenounceOwnership(&_Tokenerc20.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Tokenerc20 *Tokenerc20TransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Tokenerc20.Contract.RenounceOwnership(&_Tokenerc20.TransactOpts)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_Tokenerc20 *Tokenerc20Transactor) Transfer(opts *bind.TransactOpts, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Tokenerc20.contract.Transact(opts, "transfer", recipient, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_Tokenerc20 *Tokenerc20Session) Transfer(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Tokenerc20.Contract.Transfer(&_Tokenerc20.TransactOpts, recipient, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_Tokenerc20 *Tokenerc20TransactorSession) Transfer(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Tokenerc20.Contract.Transfer(&_Tokenerc20.TransactOpts, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_Tokenerc20 *Tokenerc20Transactor) TransferFrom(opts *bind.TransactOpts, sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Tokenerc20.contract.Transact(opts, "transferFrom", sender, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_Tokenerc20 *Tokenerc20Session) TransferFrom(sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Tokenerc20.Contract.TransferFrom(&_Tokenerc20.TransactOpts, sender, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_Tokenerc20 *Tokenerc20TransactorSession) TransferFrom(sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Tokenerc20.Contract.TransferFrom(&_Tokenerc20.TransactOpts, sender, recipient, amount)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Tokenerc20 *Tokenerc20Transactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Tokenerc20.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Tokenerc20 *Tokenerc20Session) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Tokenerc20.Contract.TransferOwnership(&_Tokenerc20.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Tokenerc20 *Tokenerc20TransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Tokenerc20.Contract.TransferOwnership(&_Tokenerc20.TransactOpts, newOwner)
}

// Tokenerc20ApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the Tokenerc20 contract.
type Tokenerc20ApprovalIterator struct {
	Event *Tokenerc20Approval // Event containing the contract specifics and raw log

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
func (it *Tokenerc20ApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Tokenerc20Approval)
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
		it.Event = new(Tokenerc20Approval)
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
func (it *Tokenerc20ApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Tokenerc20ApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Tokenerc20Approval represents a Approval event raised by the Tokenerc20 contract.
type Tokenerc20Approval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_Tokenerc20 *Tokenerc20Filterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*Tokenerc20ApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _Tokenerc20.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &Tokenerc20ApprovalIterator{contract: _Tokenerc20.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_Tokenerc20 *Tokenerc20Filterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *Tokenerc20Approval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _Tokenerc20.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Tokenerc20Approval)
				if err := _Tokenerc20.contract.UnpackLog(event, "Approval", log); err != nil {
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

// ParseApproval is a log parse operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_Tokenerc20 *Tokenerc20Filterer) ParseApproval(log types.Log) (*Tokenerc20Approval, error) {
	event := new(Tokenerc20Approval)
	if err := _Tokenerc20.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Tokenerc20HashAddedIterator is returned from FilterHashAdded and is used to iterate over the raw logs and unpacked data for HashAdded events raised by the Tokenerc20 contract.
type Tokenerc20HashAddedIterator struct {
	Event *Tokenerc20HashAdded // Event containing the contract specifics and raw log

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
func (it *Tokenerc20HashAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Tokenerc20HashAdded)
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
		it.Event = new(Tokenerc20HashAdded)
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
func (it *Tokenerc20HashAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Tokenerc20HashAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Tokenerc20HashAdded represents a HashAdded event raised by the Tokenerc20 contract.
type Tokenerc20HashAdded struct {
	Hash string
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterHashAdded is a free log retrieval operation binding the contract event 0xe0c9f5e6f5abddac86dac0e02afc9f3fda7b7fc6d9454a13c51fcb28621e1e5f.
//
// Solidity: event HashAdded(string hash)
func (_Tokenerc20 *Tokenerc20Filterer) FilterHashAdded(opts *bind.FilterOpts) (*Tokenerc20HashAddedIterator, error) {

	logs, sub, err := _Tokenerc20.contract.FilterLogs(opts, "HashAdded")
	if err != nil {
		return nil, err
	}
	return &Tokenerc20HashAddedIterator{contract: _Tokenerc20.contract, event: "HashAdded", logs: logs, sub: sub}, nil
}

// WatchHashAdded is a free log subscription operation binding the contract event 0xe0c9f5e6f5abddac86dac0e02afc9f3fda7b7fc6d9454a13c51fcb28621e1e5f.
//
// Solidity: event HashAdded(string hash)
func (_Tokenerc20 *Tokenerc20Filterer) WatchHashAdded(opts *bind.WatchOpts, sink chan<- *Tokenerc20HashAdded) (event.Subscription, error) {

	logs, sub, err := _Tokenerc20.contract.WatchLogs(opts, "HashAdded")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Tokenerc20HashAdded)
				if err := _Tokenerc20.contract.UnpackLog(event, "HashAdded", log); err != nil {
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

// ParseHashAdded is a log parse operation binding the contract event 0xe0c9f5e6f5abddac86dac0e02afc9f3fda7b7fc6d9454a13c51fcb28621e1e5f.
//
// Solidity: event HashAdded(string hash)
func (_Tokenerc20 *Tokenerc20Filterer) ParseHashAdded(log types.Log) (*Tokenerc20HashAdded, error) {
	event := new(Tokenerc20HashAdded)
	if err := _Tokenerc20.contract.UnpackLog(event, "HashAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Tokenerc20OwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Tokenerc20 contract.
type Tokenerc20OwnershipTransferredIterator struct {
	Event *Tokenerc20OwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *Tokenerc20OwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Tokenerc20OwnershipTransferred)
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
		it.Event = new(Tokenerc20OwnershipTransferred)
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
func (it *Tokenerc20OwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Tokenerc20OwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Tokenerc20OwnershipTransferred represents a OwnershipTransferred event raised by the Tokenerc20 contract.
type Tokenerc20OwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Tokenerc20 *Tokenerc20Filterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*Tokenerc20OwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Tokenerc20.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &Tokenerc20OwnershipTransferredIterator{contract: _Tokenerc20.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Tokenerc20 *Tokenerc20Filterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *Tokenerc20OwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Tokenerc20.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Tokenerc20OwnershipTransferred)
				if err := _Tokenerc20.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

// ParseOwnershipTransferred is a log parse operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Tokenerc20 *Tokenerc20Filterer) ParseOwnershipTransferred(log types.Log) (*Tokenerc20OwnershipTransferred, error) {
	event := new(Tokenerc20OwnershipTransferred)
	if err := _Tokenerc20.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Tokenerc20TransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the Tokenerc20 contract.
type Tokenerc20TransferIterator struct {
	Event *Tokenerc20Transfer // Event containing the contract specifics and raw log

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
func (it *Tokenerc20TransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Tokenerc20Transfer)
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
		it.Event = new(Tokenerc20Transfer)
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
func (it *Tokenerc20TransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Tokenerc20TransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Tokenerc20Transfer represents a Transfer event raised by the Tokenerc20 contract.
type Tokenerc20Transfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_Tokenerc20 *Tokenerc20Filterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*Tokenerc20TransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Tokenerc20.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &Tokenerc20TransferIterator{contract: _Tokenerc20.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_Tokenerc20 *Tokenerc20Filterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *Tokenerc20Transfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Tokenerc20.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Tokenerc20Transfer)
				if err := _Tokenerc20.contract.UnpackLog(event, "Transfer", log); err != nil {
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
func (_Tokenerc20 *Tokenerc20Filterer) ParseTransfer(log types.Log) (*Tokenerc20Transfer, error) {
	event := new(Tokenerc20Transfer)
	if err := _Tokenerc20.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
