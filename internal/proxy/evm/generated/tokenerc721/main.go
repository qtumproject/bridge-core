// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package tokenerc721

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

// Tokenerc721MetaData contains all meta data concerning the Tokenerc721 contract.
var Tokenerc721MetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"string\",\"name\":\"tokenName\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"tokenSymbol\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"contractURI\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"_signersRep\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"approved\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"ApprovalForAll\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"hash\",\"type\":\"string\"}],\"name\":\"HashAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_hash\",\"type\":\"string\"}],\"name\":\"containsHash\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"contractURI\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"getApproved\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"isApprovedForAll\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_txHash\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_tokenURI\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"_receiver\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_tokenID\",\"type\":\"uint256\"},{\"internalType\":\"bytes32[]\",\"name\":\"_r\",\"type\":\"bytes32[]\"},{\"internalType\":\"bytes32[]\",\"name\":\"_s\",\"type\":\"bytes32[]\"},{\"internalType\":\"uint8[]\",\"name\":\"_v\",\"type\":\"uint8[]\"}],\"name\":\"mint\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_txHash\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_tokenURI\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"_tokenID\",\"type\":\"uint256\"},{\"internalType\":\"bytes32[]\",\"name\":\"_r\",\"type\":\"bytes32[]\"},{\"internalType\":\"bytes32[]\",\"name\":\"_s\",\"type\":\"bytes32[]\"},{\"internalType\":\"uint8[]\",\"name\":\"_v\",\"type\":\"uint8[]\"}],\"name\":\"mint\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"ownerOf\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"_data\",\"type\":\"bytes\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"setApprovalForAll\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_uri\",\"type\":\"string\"}],\"name\":\"setContractURI\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_signersRep\",\"type\":\"address\"}],\"name\":\"setSignersRep\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"signersRep\",\"outputs\":[{\"internalType\":\"contractISignersRepository\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"tokenByIndex\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"tokenOfOwnerByIndex\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"tokenURI\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenID\",\"type\":\"uint256\"}],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// Tokenerc721ABI is the input ABI used to generate the binding from.
// Deprecated: Use Tokenerc721MetaData.ABI instead.
var Tokenerc721ABI = Tokenerc721MetaData.ABI

// Tokenerc721 is an auto generated Go binding around an Ethereum contract.
type Tokenerc721 struct {
	Tokenerc721Caller     // Read-only binding to the contract
	Tokenerc721Transactor // Write-only binding to the contract
	Tokenerc721Filterer   // Log filterer for contract events
}

// Tokenerc721Caller is an auto generated read-only Go binding around an Ethereum contract.
type Tokenerc721Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Tokenerc721Transactor is an auto generated write-only Go binding around an Ethereum contract.
type Tokenerc721Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Tokenerc721Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type Tokenerc721Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Tokenerc721Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type Tokenerc721Session struct {
	Contract     *Tokenerc721      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// Tokenerc721CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type Tokenerc721CallerSession struct {
	Contract *Tokenerc721Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// Tokenerc721TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type Tokenerc721TransactorSession struct {
	Contract     *Tokenerc721Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// Tokenerc721Raw is an auto generated low-level Go binding around an Ethereum contract.
type Tokenerc721Raw struct {
	Contract *Tokenerc721 // Generic contract binding to access the raw methods on
}

// Tokenerc721CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type Tokenerc721CallerRaw struct {
	Contract *Tokenerc721Caller // Generic read-only contract binding to access the raw methods on
}

// Tokenerc721TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type Tokenerc721TransactorRaw struct {
	Contract *Tokenerc721Transactor // Generic write-only contract binding to access the raw methods on
}

// NewTokenerc721 creates a new instance of Tokenerc721, bound to a specific deployed contract.
func NewTokenerc721(address common.Address, backend bind.ContractBackend) (*Tokenerc721, error) {
	contract, err := bindTokenerc721(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Tokenerc721{Tokenerc721Caller: Tokenerc721Caller{contract: contract}, Tokenerc721Transactor: Tokenerc721Transactor{contract: contract}, Tokenerc721Filterer: Tokenerc721Filterer{contract: contract}}, nil
}

// NewTokenerc721Caller creates a new read-only instance of Tokenerc721, bound to a specific deployed contract.
func NewTokenerc721Caller(address common.Address, caller bind.ContractCaller) (*Tokenerc721Caller, error) {
	contract, err := bindTokenerc721(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &Tokenerc721Caller{contract: contract}, nil
}

// NewTokenerc721Transactor creates a new write-only instance of Tokenerc721, bound to a specific deployed contract.
func NewTokenerc721Transactor(address common.Address, transactor bind.ContractTransactor) (*Tokenerc721Transactor, error) {
	contract, err := bindTokenerc721(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &Tokenerc721Transactor{contract: contract}, nil
}

// NewTokenerc721Filterer creates a new log filterer instance of Tokenerc721, bound to a specific deployed contract.
func NewTokenerc721Filterer(address common.Address, filterer bind.ContractFilterer) (*Tokenerc721Filterer, error) {
	contract, err := bindTokenerc721(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &Tokenerc721Filterer{contract: contract}, nil
}

// bindTokenerc721 binds a generic wrapper to an already deployed contract.
func bindTokenerc721(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(Tokenerc721ABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Tokenerc721 *Tokenerc721Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Tokenerc721.Contract.Tokenerc721Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Tokenerc721 *Tokenerc721Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Tokenerc721.Contract.Tokenerc721Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Tokenerc721 *Tokenerc721Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Tokenerc721.Contract.Tokenerc721Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Tokenerc721 *Tokenerc721CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Tokenerc721.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Tokenerc721 *Tokenerc721TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Tokenerc721.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Tokenerc721 *Tokenerc721TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Tokenerc721.Contract.contract.Transact(opts, method, params...)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_Tokenerc721 *Tokenerc721Caller) BalanceOf(opts *bind.CallOpts, owner common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Tokenerc721.contract.Call(opts, &out, "balanceOf", owner)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_Tokenerc721 *Tokenerc721Session) BalanceOf(owner common.Address) (*big.Int, error) {
	return _Tokenerc721.Contract.BalanceOf(&_Tokenerc721.CallOpts, owner)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_Tokenerc721 *Tokenerc721CallerSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _Tokenerc721.Contract.BalanceOf(&_Tokenerc721.CallOpts, owner)
}

// ContainsHash is a free data retrieval call binding the contract method 0x6492f57a.
//
// Solidity: function containsHash(string _hash) view returns(bool)
func (_Tokenerc721 *Tokenerc721Caller) ContainsHash(opts *bind.CallOpts, _hash string) (bool, error) {
	var out []interface{}
	err := _Tokenerc721.contract.Call(opts, &out, "containsHash", _hash)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// ContainsHash is a free data retrieval call binding the contract method 0x6492f57a.
//
// Solidity: function containsHash(string _hash) view returns(bool)
func (_Tokenerc721 *Tokenerc721Session) ContainsHash(_hash string) (bool, error) {
	return _Tokenerc721.Contract.ContainsHash(&_Tokenerc721.CallOpts, _hash)
}

// ContainsHash is a free data retrieval call binding the contract method 0x6492f57a.
//
// Solidity: function containsHash(string _hash) view returns(bool)
func (_Tokenerc721 *Tokenerc721CallerSession) ContainsHash(_hash string) (bool, error) {
	return _Tokenerc721.Contract.ContainsHash(&_Tokenerc721.CallOpts, _hash)
}

// ContractURI is a free data retrieval call binding the contract method 0xe8a3d485.
//
// Solidity: function contractURI() view returns(string)
func (_Tokenerc721 *Tokenerc721Caller) ContractURI(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Tokenerc721.contract.Call(opts, &out, "contractURI")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// ContractURI is a free data retrieval call binding the contract method 0xe8a3d485.
//
// Solidity: function contractURI() view returns(string)
func (_Tokenerc721 *Tokenerc721Session) ContractURI() (string, error) {
	return _Tokenerc721.Contract.ContractURI(&_Tokenerc721.CallOpts)
}

// ContractURI is a free data retrieval call binding the contract method 0xe8a3d485.
//
// Solidity: function contractURI() view returns(string)
func (_Tokenerc721 *Tokenerc721CallerSession) ContractURI() (string, error) {
	return _Tokenerc721.Contract.ContractURI(&_Tokenerc721.CallOpts)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_Tokenerc721 *Tokenerc721Caller) GetApproved(opts *bind.CallOpts, tokenId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Tokenerc721.contract.Call(opts, &out, "getApproved", tokenId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_Tokenerc721 *Tokenerc721Session) GetApproved(tokenId *big.Int) (common.Address, error) {
	return _Tokenerc721.Contract.GetApproved(&_Tokenerc721.CallOpts, tokenId)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_Tokenerc721 *Tokenerc721CallerSession) GetApproved(tokenId *big.Int) (common.Address, error) {
	return _Tokenerc721.Contract.GetApproved(&_Tokenerc721.CallOpts, tokenId)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_Tokenerc721 *Tokenerc721Caller) IsApprovedForAll(opts *bind.CallOpts, owner common.Address, operator common.Address) (bool, error) {
	var out []interface{}
	err := _Tokenerc721.contract.Call(opts, &out, "isApprovedForAll", owner, operator)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_Tokenerc721 *Tokenerc721Session) IsApprovedForAll(owner common.Address, operator common.Address) (bool, error) {
	return _Tokenerc721.Contract.IsApprovedForAll(&_Tokenerc721.CallOpts, owner, operator)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_Tokenerc721 *Tokenerc721CallerSession) IsApprovedForAll(owner common.Address, operator common.Address) (bool, error) {
	return _Tokenerc721.Contract.IsApprovedForAll(&_Tokenerc721.CallOpts, owner, operator)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Tokenerc721 *Tokenerc721Caller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Tokenerc721.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Tokenerc721 *Tokenerc721Session) Name() (string, error) {
	return _Tokenerc721.Contract.Name(&_Tokenerc721.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Tokenerc721 *Tokenerc721CallerSession) Name() (string, error) {
	return _Tokenerc721.Contract.Name(&_Tokenerc721.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Tokenerc721 *Tokenerc721Caller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Tokenerc721.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Tokenerc721 *Tokenerc721Session) Owner() (common.Address, error) {
	return _Tokenerc721.Contract.Owner(&_Tokenerc721.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Tokenerc721 *Tokenerc721CallerSession) Owner() (common.Address, error) {
	return _Tokenerc721.Contract.Owner(&_Tokenerc721.CallOpts)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_Tokenerc721 *Tokenerc721Caller) OwnerOf(opts *bind.CallOpts, tokenId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Tokenerc721.contract.Call(opts, &out, "ownerOf", tokenId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_Tokenerc721 *Tokenerc721Session) OwnerOf(tokenId *big.Int) (common.Address, error) {
	return _Tokenerc721.Contract.OwnerOf(&_Tokenerc721.CallOpts, tokenId)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_Tokenerc721 *Tokenerc721CallerSession) OwnerOf(tokenId *big.Int) (common.Address, error) {
	return _Tokenerc721.Contract.OwnerOf(&_Tokenerc721.CallOpts, tokenId)
}

// SignersRep is a free data retrieval call binding the contract method 0x897b73cd.
//
// Solidity: function signersRep() view returns(address)
func (_Tokenerc721 *Tokenerc721Caller) SignersRep(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Tokenerc721.contract.Call(opts, &out, "signersRep")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// SignersRep is a free data retrieval call binding the contract method 0x897b73cd.
//
// Solidity: function signersRep() view returns(address)
func (_Tokenerc721 *Tokenerc721Session) SignersRep() (common.Address, error) {
	return _Tokenerc721.Contract.SignersRep(&_Tokenerc721.CallOpts)
}

// SignersRep is a free data retrieval call binding the contract method 0x897b73cd.
//
// Solidity: function signersRep() view returns(address)
func (_Tokenerc721 *Tokenerc721CallerSession) SignersRep() (common.Address, error) {
	return _Tokenerc721.Contract.SignersRep(&_Tokenerc721.CallOpts)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Tokenerc721 *Tokenerc721Caller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _Tokenerc721.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Tokenerc721 *Tokenerc721Session) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _Tokenerc721.Contract.SupportsInterface(&_Tokenerc721.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Tokenerc721 *Tokenerc721CallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _Tokenerc721.Contract.SupportsInterface(&_Tokenerc721.CallOpts, interfaceId)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Tokenerc721 *Tokenerc721Caller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Tokenerc721.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Tokenerc721 *Tokenerc721Session) Symbol() (string, error) {
	return _Tokenerc721.Contract.Symbol(&_Tokenerc721.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Tokenerc721 *Tokenerc721CallerSession) Symbol() (string, error) {
	return _Tokenerc721.Contract.Symbol(&_Tokenerc721.CallOpts)
}

// TokenByIndex is a free data retrieval call binding the contract method 0x4f6ccce7.
//
// Solidity: function tokenByIndex(uint256 index) view returns(uint256)
func (_Tokenerc721 *Tokenerc721Caller) TokenByIndex(opts *bind.CallOpts, index *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Tokenerc721.contract.Call(opts, &out, "tokenByIndex", index)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TokenByIndex is a free data retrieval call binding the contract method 0x4f6ccce7.
//
// Solidity: function tokenByIndex(uint256 index) view returns(uint256)
func (_Tokenerc721 *Tokenerc721Session) TokenByIndex(index *big.Int) (*big.Int, error) {
	return _Tokenerc721.Contract.TokenByIndex(&_Tokenerc721.CallOpts, index)
}

// TokenByIndex is a free data retrieval call binding the contract method 0x4f6ccce7.
//
// Solidity: function tokenByIndex(uint256 index) view returns(uint256)
func (_Tokenerc721 *Tokenerc721CallerSession) TokenByIndex(index *big.Int) (*big.Int, error) {
	return _Tokenerc721.Contract.TokenByIndex(&_Tokenerc721.CallOpts, index)
}

// TokenOfOwnerByIndex is a free data retrieval call binding the contract method 0x2f745c59.
//
// Solidity: function tokenOfOwnerByIndex(address owner, uint256 index) view returns(uint256)
func (_Tokenerc721 *Tokenerc721Caller) TokenOfOwnerByIndex(opts *bind.CallOpts, owner common.Address, index *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Tokenerc721.contract.Call(opts, &out, "tokenOfOwnerByIndex", owner, index)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TokenOfOwnerByIndex is a free data retrieval call binding the contract method 0x2f745c59.
//
// Solidity: function tokenOfOwnerByIndex(address owner, uint256 index) view returns(uint256)
func (_Tokenerc721 *Tokenerc721Session) TokenOfOwnerByIndex(owner common.Address, index *big.Int) (*big.Int, error) {
	return _Tokenerc721.Contract.TokenOfOwnerByIndex(&_Tokenerc721.CallOpts, owner, index)
}

// TokenOfOwnerByIndex is a free data retrieval call binding the contract method 0x2f745c59.
//
// Solidity: function tokenOfOwnerByIndex(address owner, uint256 index) view returns(uint256)
func (_Tokenerc721 *Tokenerc721CallerSession) TokenOfOwnerByIndex(owner common.Address, index *big.Int) (*big.Int, error) {
	return _Tokenerc721.Contract.TokenOfOwnerByIndex(&_Tokenerc721.CallOpts, owner, index)
}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_Tokenerc721 *Tokenerc721Caller) TokenURI(opts *bind.CallOpts, tokenId *big.Int) (string, error) {
	var out []interface{}
	err := _Tokenerc721.contract.Call(opts, &out, "tokenURI", tokenId)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_Tokenerc721 *Tokenerc721Session) TokenURI(tokenId *big.Int) (string, error) {
	return _Tokenerc721.Contract.TokenURI(&_Tokenerc721.CallOpts, tokenId)
}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_Tokenerc721 *Tokenerc721CallerSession) TokenURI(tokenId *big.Int) (string, error) {
	return _Tokenerc721.Contract.TokenURI(&_Tokenerc721.CallOpts, tokenId)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Tokenerc721 *Tokenerc721Caller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Tokenerc721.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Tokenerc721 *Tokenerc721Session) TotalSupply() (*big.Int, error) {
	return _Tokenerc721.Contract.TotalSupply(&_Tokenerc721.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Tokenerc721 *Tokenerc721CallerSession) TotalSupply() (*big.Int, error) {
	return _Tokenerc721.Contract.TotalSupply(&_Tokenerc721.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_Tokenerc721 *Tokenerc721Transactor) Approve(opts *bind.TransactOpts, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Tokenerc721.contract.Transact(opts, "approve", to, tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_Tokenerc721 *Tokenerc721Session) Approve(to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Tokenerc721.Contract.Approve(&_Tokenerc721.TransactOpts, to, tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_Tokenerc721 *Tokenerc721TransactorSession) Approve(to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Tokenerc721.Contract.Approve(&_Tokenerc721.TransactOpts, to, tokenId)
}

// Mint is a paid mutator transaction binding the contract method 0xaaa11882.
//
// Solidity: function mint(string _txHash, string _tokenURI, address _receiver, uint256 _tokenID, bytes32[] _r, bytes32[] _s, uint8[] _v) returns(uint256)
func (_Tokenerc721 *Tokenerc721Transactor) Mint(opts *bind.TransactOpts, _txHash string, _tokenURI string, _receiver common.Address, _tokenID *big.Int, _r [][32]byte, _s [][32]byte, _v []uint8) (*types.Transaction, error) {
	return _Tokenerc721.contract.Transact(opts, "mint", _txHash, _tokenURI, _receiver, _tokenID, _r, _s, _v)
}

// Mint is a paid mutator transaction binding the contract method 0xaaa11882.
//
// Solidity: function mint(string _txHash, string _tokenURI, address _receiver, uint256 _tokenID, bytes32[] _r, bytes32[] _s, uint8[] _v) returns(uint256)
func (_Tokenerc721 *Tokenerc721Session) Mint(_txHash string, _tokenURI string, _receiver common.Address, _tokenID *big.Int, _r [][32]byte, _s [][32]byte, _v []uint8) (*types.Transaction, error) {
	return _Tokenerc721.Contract.Mint(&_Tokenerc721.TransactOpts, _txHash, _tokenURI, _receiver, _tokenID, _r, _s, _v)
}

// Mint is a paid mutator transaction binding the contract method 0xaaa11882.
//
// Solidity: function mint(string _txHash, string _tokenURI, address _receiver, uint256 _tokenID, bytes32[] _r, bytes32[] _s, uint8[] _v) returns(uint256)
func (_Tokenerc721 *Tokenerc721TransactorSession) Mint(_txHash string, _tokenURI string, _receiver common.Address, _tokenID *big.Int, _r [][32]byte, _s [][32]byte, _v []uint8) (*types.Transaction, error) {
	return _Tokenerc721.Contract.Mint(&_Tokenerc721.TransactOpts, _txHash, _tokenURI, _receiver, _tokenID, _r, _s, _v)
}

// Mint0 is a paid mutator transaction binding the contract method 0xe7308983.
//
// Solidity: function mint(string _txHash, string _tokenURI, uint256 _tokenID, bytes32[] _r, bytes32[] _s, uint8[] _v) returns(uint256)
func (_Tokenerc721 *Tokenerc721Transactor) Mint0(opts *bind.TransactOpts, _txHash string, _tokenURI string, _tokenID *big.Int, _r [][32]byte, _s [][32]byte, _v []uint8) (*types.Transaction, error) {
	return _Tokenerc721.contract.Transact(opts, "mint0", _txHash, _tokenURI, _tokenID, _r, _s, _v)
}

// Mint0 is a paid mutator transaction binding the contract method 0xe7308983.
//
// Solidity: function mint(string _txHash, string _tokenURI, uint256 _tokenID, bytes32[] _r, bytes32[] _s, uint8[] _v) returns(uint256)
func (_Tokenerc721 *Tokenerc721Session) Mint0(_txHash string, _tokenURI string, _tokenID *big.Int, _r [][32]byte, _s [][32]byte, _v []uint8) (*types.Transaction, error) {
	return _Tokenerc721.Contract.Mint0(&_Tokenerc721.TransactOpts, _txHash, _tokenURI, _tokenID, _r, _s, _v)
}

// Mint0 is a paid mutator transaction binding the contract method 0xe7308983.
//
// Solidity: function mint(string _txHash, string _tokenURI, uint256 _tokenID, bytes32[] _r, bytes32[] _s, uint8[] _v) returns(uint256)
func (_Tokenerc721 *Tokenerc721TransactorSession) Mint0(_txHash string, _tokenURI string, _tokenID *big.Int, _r [][32]byte, _s [][32]byte, _v []uint8) (*types.Transaction, error) {
	return _Tokenerc721.Contract.Mint0(&_Tokenerc721.TransactOpts, _txHash, _tokenURI, _tokenID, _r, _s, _v)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Tokenerc721 *Tokenerc721Transactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Tokenerc721.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Tokenerc721 *Tokenerc721Session) RenounceOwnership() (*types.Transaction, error) {
	return _Tokenerc721.Contract.RenounceOwnership(&_Tokenerc721.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Tokenerc721 *Tokenerc721TransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Tokenerc721.Contract.RenounceOwnership(&_Tokenerc721.TransactOpts)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_Tokenerc721 *Tokenerc721Transactor) SafeTransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Tokenerc721.contract.Transact(opts, "safeTransferFrom", from, to, tokenId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_Tokenerc721 *Tokenerc721Session) SafeTransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Tokenerc721.Contract.SafeTransferFrom(&_Tokenerc721.TransactOpts, from, to, tokenId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_Tokenerc721 *Tokenerc721TransactorSession) SafeTransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Tokenerc721.Contract.SafeTransferFrom(&_Tokenerc721.TransactOpts, from, to, tokenId)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes _data) returns()
func (_Tokenerc721 *Tokenerc721Transactor) SafeTransferFrom0(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int, _data []byte) (*types.Transaction, error) {
	return _Tokenerc721.contract.Transact(opts, "safeTransferFrom0", from, to, tokenId, _data)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes _data) returns()
func (_Tokenerc721 *Tokenerc721Session) SafeTransferFrom0(from common.Address, to common.Address, tokenId *big.Int, _data []byte) (*types.Transaction, error) {
	return _Tokenerc721.Contract.SafeTransferFrom0(&_Tokenerc721.TransactOpts, from, to, tokenId, _data)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes _data) returns()
func (_Tokenerc721 *Tokenerc721TransactorSession) SafeTransferFrom0(from common.Address, to common.Address, tokenId *big.Int, _data []byte) (*types.Transaction, error) {
	return _Tokenerc721.Contract.SafeTransferFrom0(&_Tokenerc721.TransactOpts, from, to, tokenId, _data)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_Tokenerc721 *Tokenerc721Transactor) SetApprovalForAll(opts *bind.TransactOpts, operator common.Address, approved bool) (*types.Transaction, error) {
	return _Tokenerc721.contract.Transact(opts, "setApprovalForAll", operator, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_Tokenerc721 *Tokenerc721Session) SetApprovalForAll(operator common.Address, approved bool) (*types.Transaction, error) {
	return _Tokenerc721.Contract.SetApprovalForAll(&_Tokenerc721.TransactOpts, operator, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_Tokenerc721 *Tokenerc721TransactorSession) SetApprovalForAll(operator common.Address, approved bool) (*types.Transaction, error) {
	return _Tokenerc721.Contract.SetApprovalForAll(&_Tokenerc721.TransactOpts, operator, approved)
}

// SetContractURI is a paid mutator transaction binding the contract method 0x938e3d7b.
//
// Solidity: function setContractURI(string _uri) returns()
func (_Tokenerc721 *Tokenerc721Transactor) SetContractURI(opts *bind.TransactOpts, _uri string) (*types.Transaction, error) {
	return _Tokenerc721.contract.Transact(opts, "setContractURI", _uri)
}

// SetContractURI is a paid mutator transaction binding the contract method 0x938e3d7b.
//
// Solidity: function setContractURI(string _uri) returns()
func (_Tokenerc721 *Tokenerc721Session) SetContractURI(_uri string) (*types.Transaction, error) {
	return _Tokenerc721.Contract.SetContractURI(&_Tokenerc721.TransactOpts, _uri)
}

// SetContractURI is a paid mutator transaction binding the contract method 0x938e3d7b.
//
// Solidity: function setContractURI(string _uri) returns()
func (_Tokenerc721 *Tokenerc721TransactorSession) SetContractURI(_uri string) (*types.Transaction, error) {
	return _Tokenerc721.Contract.SetContractURI(&_Tokenerc721.TransactOpts, _uri)
}

// SetSignersRep is a paid mutator transaction binding the contract method 0xc3e345f9.
//
// Solidity: function setSignersRep(address _signersRep) returns()
func (_Tokenerc721 *Tokenerc721Transactor) SetSignersRep(opts *bind.TransactOpts, _signersRep common.Address) (*types.Transaction, error) {
	return _Tokenerc721.contract.Transact(opts, "setSignersRep", _signersRep)
}

// SetSignersRep is a paid mutator transaction binding the contract method 0xc3e345f9.
//
// Solidity: function setSignersRep(address _signersRep) returns()
func (_Tokenerc721 *Tokenerc721Session) SetSignersRep(_signersRep common.Address) (*types.Transaction, error) {
	return _Tokenerc721.Contract.SetSignersRep(&_Tokenerc721.TransactOpts, _signersRep)
}

// SetSignersRep is a paid mutator transaction binding the contract method 0xc3e345f9.
//
// Solidity: function setSignersRep(address _signersRep) returns()
func (_Tokenerc721 *Tokenerc721TransactorSession) SetSignersRep(_signersRep common.Address) (*types.Transaction, error) {
	return _Tokenerc721.Contract.SetSignersRep(&_Tokenerc721.TransactOpts, _signersRep)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_Tokenerc721 *Tokenerc721Transactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Tokenerc721.contract.Transact(opts, "transferFrom", from, to, tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_Tokenerc721 *Tokenerc721Session) TransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Tokenerc721.Contract.TransferFrom(&_Tokenerc721.TransactOpts, from, to, tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_Tokenerc721 *Tokenerc721TransactorSession) TransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Tokenerc721.Contract.TransferFrom(&_Tokenerc721.TransactOpts, from, to, tokenId)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Tokenerc721 *Tokenerc721Transactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Tokenerc721.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Tokenerc721 *Tokenerc721Session) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Tokenerc721.Contract.TransferOwnership(&_Tokenerc721.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Tokenerc721 *Tokenerc721TransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Tokenerc721.Contract.TransferOwnership(&_Tokenerc721.TransactOpts, newOwner)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 tokenID) returns()
func (_Tokenerc721 *Tokenerc721Transactor) Withdraw(opts *bind.TransactOpts, tokenID *big.Int) (*types.Transaction, error) {
	return _Tokenerc721.contract.Transact(opts, "withdraw", tokenID)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 tokenID) returns()
func (_Tokenerc721 *Tokenerc721Session) Withdraw(tokenID *big.Int) (*types.Transaction, error) {
	return _Tokenerc721.Contract.Withdraw(&_Tokenerc721.TransactOpts, tokenID)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 tokenID) returns()
func (_Tokenerc721 *Tokenerc721TransactorSession) Withdraw(tokenID *big.Int) (*types.Transaction, error) {
	return _Tokenerc721.Contract.Withdraw(&_Tokenerc721.TransactOpts, tokenID)
}

// Tokenerc721ApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the Tokenerc721 contract.
type Tokenerc721ApprovalIterator struct {
	Event *Tokenerc721Approval // Event containing the contract specifics and raw log

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
func (it *Tokenerc721ApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Tokenerc721Approval)
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
		it.Event = new(Tokenerc721Approval)
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
func (it *Tokenerc721ApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Tokenerc721ApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Tokenerc721Approval represents a Approval event raised by the Tokenerc721 contract.
type Tokenerc721Approval struct {
	Owner    common.Address
	Approved common.Address
	TokenId  *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_Tokenerc721 *Tokenerc721Filterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, approved []common.Address, tokenId []*big.Int) (*Tokenerc721ApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var approvedRule []interface{}
	for _, approvedItem := range approved {
		approvedRule = append(approvedRule, approvedItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _Tokenerc721.contract.FilterLogs(opts, "Approval", ownerRule, approvedRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &Tokenerc721ApprovalIterator{contract: _Tokenerc721.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_Tokenerc721 *Tokenerc721Filterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *Tokenerc721Approval, owner []common.Address, approved []common.Address, tokenId []*big.Int) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var approvedRule []interface{}
	for _, approvedItem := range approved {
		approvedRule = append(approvedRule, approvedItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _Tokenerc721.contract.WatchLogs(opts, "Approval", ownerRule, approvedRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Tokenerc721Approval)
				if err := _Tokenerc721.contract.UnpackLog(event, "Approval", log); err != nil {
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
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_Tokenerc721 *Tokenerc721Filterer) ParseApproval(log types.Log) (*Tokenerc721Approval, error) {
	event := new(Tokenerc721Approval)
	if err := _Tokenerc721.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Tokenerc721ApprovalForAllIterator is returned from FilterApprovalForAll and is used to iterate over the raw logs and unpacked data for ApprovalForAll events raised by the Tokenerc721 contract.
type Tokenerc721ApprovalForAllIterator struct {
	Event *Tokenerc721ApprovalForAll // Event containing the contract specifics and raw log

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
func (it *Tokenerc721ApprovalForAllIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Tokenerc721ApprovalForAll)
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
		it.Event = new(Tokenerc721ApprovalForAll)
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
func (it *Tokenerc721ApprovalForAllIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Tokenerc721ApprovalForAllIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Tokenerc721ApprovalForAll represents a ApprovalForAll event raised by the Tokenerc721 contract.
type Tokenerc721ApprovalForAll struct {
	Owner    common.Address
	Operator common.Address
	Approved bool
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApprovalForAll is a free log retrieval operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_Tokenerc721 *Tokenerc721Filterer) FilterApprovalForAll(opts *bind.FilterOpts, owner []common.Address, operator []common.Address) (*Tokenerc721ApprovalForAllIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _Tokenerc721.contract.FilterLogs(opts, "ApprovalForAll", ownerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return &Tokenerc721ApprovalForAllIterator{contract: _Tokenerc721.contract, event: "ApprovalForAll", logs: logs, sub: sub}, nil
}

// WatchApprovalForAll is a free log subscription operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_Tokenerc721 *Tokenerc721Filterer) WatchApprovalForAll(opts *bind.WatchOpts, sink chan<- *Tokenerc721ApprovalForAll, owner []common.Address, operator []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _Tokenerc721.contract.WatchLogs(opts, "ApprovalForAll", ownerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Tokenerc721ApprovalForAll)
				if err := _Tokenerc721.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
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

// ParseApprovalForAll is a log parse operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_Tokenerc721 *Tokenerc721Filterer) ParseApprovalForAll(log types.Log) (*Tokenerc721ApprovalForAll, error) {
	event := new(Tokenerc721ApprovalForAll)
	if err := _Tokenerc721.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Tokenerc721HashAddedIterator is returned from FilterHashAdded and is used to iterate over the raw logs and unpacked data for HashAdded events raised by the Tokenerc721 contract.
type Tokenerc721HashAddedIterator struct {
	Event *Tokenerc721HashAdded // Event containing the contract specifics and raw log

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
func (it *Tokenerc721HashAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Tokenerc721HashAdded)
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
		it.Event = new(Tokenerc721HashAdded)
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
func (it *Tokenerc721HashAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Tokenerc721HashAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Tokenerc721HashAdded represents a HashAdded event raised by the Tokenerc721 contract.
type Tokenerc721HashAdded struct {
	Hash string
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterHashAdded is a free log retrieval operation binding the contract event 0xe0c9f5e6f5abddac86dac0e02afc9f3fda7b7fc6d9454a13c51fcb28621e1e5f.
//
// Solidity: event HashAdded(string hash)
func (_Tokenerc721 *Tokenerc721Filterer) FilterHashAdded(opts *bind.FilterOpts) (*Tokenerc721HashAddedIterator, error) {

	logs, sub, err := _Tokenerc721.contract.FilterLogs(opts, "HashAdded")
	if err != nil {
		return nil, err
	}
	return &Tokenerc721HashAddedIterator{contract: _Tokenerc721.contract, event: "HashAdded", logs: logs, sub: sub}, nil
}

// WatchHashAdded is a free log subscription operation binding the contract event 0xe0c9f5e6f5abddac86dac0e02afc9f3fda7b7fc6d9454a13c51fcb28621e1e5f.
//
// Solidity: event HashAdded(string hash)
func (_Tokenerc721 *Tokenerc721Filterer) WatchHashAdded(opts *bind.WatchOpts, sink chan<- *Tokenerc721HashAdded) (event.Subscription, error) {

	logs, sub, err := _Tokenerc721.contract.WatchLogs(opts, "HashAdded")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Tokenerc721HashAdded)
				if err := _Tokenerc721.contract.UnpackLog(event, "HashAdded", log); err != nil {
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
func (_Tokenerc721 *Tokenerc721Filterer) ParseHashAdded(log types.Log) (*Tokenerc721HashAdded, error) {
	event := new(Tokenerc721HashAdded)
	if err := _Tokenerc721.contract.UnpackLog(event, "HashAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Tokenerc721OwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Tokenerc721 contract.
type Tokenerc721OwnershipTransferredIterator struct {
	Event *Tokenerc721OwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *Tokenerc721OwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Tokenerc721OwnershipTransferred)
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
		it.Event = new(Tokenerc721OwnershipTransferred)
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
func (it *Tokenerc721OwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Tokenerc721OwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Tokenerc721OwnershipTransferred represents a OwnershipTransferred event raised by the Tokenerc721 contract.
type Tokenerc721OwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Tokenerc721 *Tokenerc721Filterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*Tokenerc721OwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Tokenerc721.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &Tokenerc721OwnershipTransferredIterator{contract: _Tokenerc721.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Tokenerc721 *Tokenerc721Filterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *Tokenerc721OwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Tokenerc721.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Tokenerc721OwnershipTransferred)
				if err := _Tokenerc721.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_Tokenerc721 *Tokenerc721Filterer) ParseOwnershipTransferred(log types.Log) (*Tokenerc721OwnershipTransferred, error) {
	event := new(Tokenerc721OwnershipTransferred)
	if err := _Tokenerc721.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Tokenerc721TransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the Tokenerc721 contract.
type Tokenerc721TransferIterator struct {
	Event *Tokenerc721Transfer // Event containing the contract specifics and raw log

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
func (it *Tokenerc721TransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Tokenerc721Transfer)
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
		it.Event = new(Tokenerc721Transfer)
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
func (it *Tokenerc721TransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Tokenerc721TransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Tokenerc721Transfer represents a Transfer event raised by the Tokenerc721 contract.
type Tokenerc721Transfer struct {
	From    common.Address
	To      common.Address
	TokenId *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_Tokenerc721 *Tokenerc721Filterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address, tokenId []*big.Int) (*Tokenerc721TransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _Tokenerc721.contract.FilterLogs(opts, "Transfer", fromRule, toRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &Tokenerc721TransferIterator{contract: _Tokenerc721.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_Tokenerc721 *Tokenerc721Filterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *Tokenerc721Transfer, from []common.Address, to []common.Address, tokenId []*big.Int) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _Tokenerc721.contract.WatchLogs(opts, "Transfer", fromRule, toRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Tokenerc721Transfer)
				if err := _Tokenerc721.contract.UnpackLog(event, "Transfer", log); err != nil {
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
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_Tokenerc721 *Tokenerc721Filterer) ParseTransfer(log types.Log) (*Tokenerc721Transfer, error) {
	event := new(Tokenerc721Transfer)
	if err := _Tokenerc721.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
