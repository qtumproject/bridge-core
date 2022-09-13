// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package bridge

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
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_signersRep\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"string\",\"name\":\"receiver\",\"type\":\"string\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"DepositERC20\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"hash\",\"type\":\"string\"}],\"name\":\"HashAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"Received\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_hash\",\"type\":\"string\"}],\"name\":\"containsHash\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"_receiver\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"depositERC20\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"name\":\"hashes\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"name\":\"onERC1155BatchReceived\",\"outputs\":[{\"internalType\":\"bytes4\",\"name\":\"\",\"type\":\"bytes4\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"name\":\"onERC1155Received\",\"outputs\":[{\"internalType\":\"bytes4\",\"name\":\"\",\"type\":\"bytes4\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_signersRep\",\"type\":\"address\"}],\"name\":\"setSignersRep\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"signersRep\",\"outputs\":[{\"internalType\":\"contractISignersRepository\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"_txHash\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_tokenURI\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"_nativeChainAddress\",\"type\":\"address\"},{\"internalType\":\"uint128\",\"name\":\"_tokenId\",\"type\":\"uint128\"},{\"internalType\":\"bytes32[]\",\"name\":\"_r\",\"type\":\"bytes32[]\"},{\"internalType\":\"bytes32[]\",\"name\":\"_s\",\"type\":\"bytes32[]\"},{\"internalType\":\"uint8[]\",\"name\":\"_v\",\"type\":\"uint8[]\"}],\"name\":\"withdrawERC1155\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_receiver\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"withdrawERC20\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"_txHash\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes32[]\",\"name\":\"_r\",\"type\":\"bytes32[]\"},{\"internalType\":\"bytes32[]\",\"name\":\"_s\",\"type\":\"bytes32[]\"},{\"internalType\":\"uint8[]\",\"name\":\"_v\",\"type\":\"uint8[]\"}],\"name\":\"withdrawERC20\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"_txHash\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"_tokenID\",\"type\":\"uint256\"},{\"internalType\":\"bytes32[]\",\"name\":\"_r\",\"type\":\"bytes32[]\"},{\"internalType\":\"bytes32[]\",\"name\":\"_s\",\"type\":\"bytes32[]\"},{\"internalType\":\"uint8[]\",\"name\":\"_v\",\"type\":\"uint8[]\"}],\"name\":\"withdrawERC721\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_receiver\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"withdrawERC721\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_receiver\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"withdrawNative\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_txHash\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes32[]\",\"name\":\"_r\",\"type\":\"bytes32[]\"},{\"internalType\":\"bytes32[]\",\"name\":\"_s\",\"type\":\"bytes32[]\"},{\"internalType\":\"uint8[]\",\"name\":\"_v\",\"type\":\"uint8[]\"}],\"name\":\"withdrawNative\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
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

// ContainsHash is a free data retrieval call binding the contract method 0x6492f57a.
//
// Solidity: function containsHash(string _hash) view returns(bool)
func (_Bridge *BridgeCaller) ContainsHash(opts *bind.CallOpts, _hash string) (bool, error) {
	var out []interface{}
	err := _Bridge.contract.Call(opts, &out, "containsHash", _hash)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// ContainsHash is a free data retrieval call binding the contract method 0x6492f57a.
//
// Solidity: function containsHash(string _hash) view returns(bool)
func (_Bridge *BridgeSession) ContainsHash(_hash string) (bool, error) {
	return _Bridge.Contract.ContainsHash(&_Bridge.CallOpts, _hash)
}

// ContainsHash is a free data retrieval call binding the contract method 0x6492f57a.
//
// Solidity: function containsHash(string _hash) view returns(bool)
func (_Bridge *BridgeCallerSession) ContainsHash(_hash string) (bool, error) {
	return _Bridge.Contract.ContainsHash(&_Bridge.CallOpts, _hash)
}

// Hashes is a free data retrieval call binding the contract method 0x685e0dc7.
//
// Solidity: function hashes(string ) view returns(bool)
func (_Bridge *BridgeCaller) Hashes(opts *bind.CallOpts, arg0 string) (bool, error) {
	var out []interface{}
	err := _Bridge.contract.Call(opts, &out, "hashes", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Hashes is a free data retrieval call binding the contract method 0x685e0dc7.
//
// Solidity: function hashes(string ) view returns(bool)
func (_Bridge *BridgeSession) Hashes(arg0 string) (bool, error) {
	return _Bridge.Contract.Hashes(&_Bridge.CallOpts, arg0)
}

// Hashes is a free data retrieval call binding the contract method 0x685e0dc7.
//
// Solidity: function hashes(string ) view returns(bool)
func (_Bridge *BridgeCallerSession) Hashes(arg0 string) (bool, error) {
	return _Bridge.Contract.Hashes(&_Bridge.CallOpts, arg0)
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

// SignersRep is a free data retrieval call binding the contract method 0x897b73cd.
//
// Solidity: function signersRep() view returns(address)
func (_Bridge *BridgeCaller) SignersRep(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Bridge.contract.Call(opts, &out, "signersRep")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// SignersRep is a free data retrieval call binding the contract method 0x897b73cd.
//
// Solidity: function signersRep() view returns(address)
func (_Bridge *BridgeSession) SignersRep() (common.Address, error) {
	return _Bridge.Contract.SignersRep(&_Bridge.CallOpts)
}

// SignersRep is a free data retrieval call binding the contract method 0x897b73cd.
//
// Solidity: function signersRep() view returns(address)
func (_Bridge *BridgeCallerSession) SignersRep() (common.Address, error) {
	return _Bridge.Contract.SignersRep(&_Bridge.CallOpts)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Bridge *BridgeCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _Bridge.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Bridge *BridgeSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _Bridge.Contract.SupportsInterface(&_Bridge.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Bridge *BridgeCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _Bridge.Contract.SupportsInterface(&_Bridge.CallOpts, interfaceId)
}

// DepositERC20 is a paid mutator transaction binding the contract method 0x7288f11a.
//
// Solidity: function depositERC20(address _token, string _receiver, uint256 _amount) returns()
func (_Bridge *BridgeTransactor) DepositERC20(opts *bind.TransactOpts, _token common.Address, _receiver string, _amount *big.Int) (*types.Transaction, error) {
	return _Bridge.contract.Transact(opts, "depositERC20", _token, _receiver, _amount)
}

// DepositERC20 is a paid mutator transaction binding the contract method 0x7288f11a.
//
// Solidity: function depositERC20(address _token, string _receiver, uint256 _amount) returns()
func (_Bridge *BridgeSession) DepositERC20(_token common.Address, _receiver string, _amount *big.Int) (*types.Transaction, error) {
	return _Bridge.Contract.DepositERC20(&_Bridge.TransactOpts, _token, _receiver, _amount)
}

// DepositERC20 is a paid mutator transaction binding the contract method 0x7288f11a.
//
// Solidity: function depositERC20(address _token, string _receiver, uint256 _amount) returns()
func (_Bridge *BridgeTransactorSession) DepositERC20(_token common.Address, _receiver string, _amount *big.Int) (*types.Transaction, error) {
	return _Bridge.Contract.DepositERC20(&_Bridge.TransactOpts, _token, _receiver, _amount)
}

// OnERC1155BatchReceived is a paid mutator transaction binding the contract method 0xbc197c81.
//
// Solidity: function onERC1155BatchReceived(address , address , uint256[] , uint256[] , bytes ) returns(bytes4)
func (_Bridge *BridgeTransactor) OnERC1155BatchReceived(opts *bind.TransactOpts, arg0 common.Address, arg1 common.Address, arg2 []*big.Int, arg3 []*big.Int, arg4 []byte) (*types.Transaction, error) {
	return _Bridge.contract.Transact(opts, "onERC1155BatchReceived", arg0, arg1, arg2, arg3, arg4)
}

// OnERC1155BatchReceived is a paid mutator transaction binding the contract method 0xbc197c81.
//
// Solidity: function onERC1155BatchReceived(address , address , uint256[] , uint256[] , bytes ) returns(bytes4)
func (_Bridge *BridgeSession) OnERC1155BatchReceived(arg0 common.Address, arg1 common.Address, arg2 []*big.Int, arg3 []*big.Int, arg4 []byte) (*types.Transaction, error) {
	return _Bridge.Contract.OnERC1155BatchReceived(&_Bridge.TransactOpts, arg0, arg1, arg2, arg3, arg4)
}

// OnERC1155BatchReceived is a paid mutator transaction binding the contract method 0xbc197c81.
//
// Solidity: function onERC1155BatchReceived(address , address , uint256[] , uint256[] , bytes ) returns(bytes4)
func (_Bridge *BridgeTransactorSession) OnERC1155BatchReceived(arg0 common.Address, arg1 common.Address, arg2 []*big.Int, arg3 []*big.Int, arg4 []byte) (*types.Transaction, error) {
	return _Bridge.Contract.OnERC1155BatchReceived(&_Bridge.TransactOpts, arg0, arg1, arg2, arg3, arg4)
}

// OnERC1155Received is a paid mutator transaction binding the contract method 0xf23a6e61.
//
// Solidity: function onERC1155Received(address , address , uint256 , uint256 , bytes ) returns(bytes4)
func (_Bridge *BridgeTransactor) OnERC1155Received(opts *bind.TransactOpts, arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 *big.Int, arg4 []byte) (*types.Transaction, error) {
	return _Bridge.contract.Transact(opts, "onERC1155Received", arg0, arg1, arg2, arg3, arg4)
}

// OnERC1155Received is a paid mutator transaction binding the contract method 0xf23a6e61.
//
// Solidity: function onERC1155Received(address , address , uint256 , uint256 , bytes ) returns(bytes4)
func (_Bridge *BridgeSession) OnERC1155Received(arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 *big.Int, arg4 []byte) (*types.Transaction, error) {
	return _Bridge.Contract.OnERC1155Received(&_Bridge.TransactOpts, arg0, arg1, arg2, arg3, arg4)
}

// OnERC1155Received is a paid mutator transaction binding the contract method 0xf23a6e61.
//
// Solidity: function onERC1155Received(address , address , uint256 , uint256 , bytes ) returns(bytes4)
func (_Bridge *BridgeTransactorSession) OnERC1155Received(arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 *big.Int, arg4 []byte) (*types.Transaction, error) {
	return _Bridge.Contract.OnERC1155Received(&_Bridge.TransactOpts, arg0, arg1, arg2, arg3, arg4)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Bridge *BridgeTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Bridge.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Bridge *BridgeSession) RenounceOwnership() (*types.Transaction, error) {
	return _Bridge.Contract.RenounceOwnership(&_Bridge.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Bridge *BridgeTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Bridge.Contract.RenounceOwnership(&_Bridge.TransactOpts)
}

// SetSignersRep is a paid mutator transaction binding the contract method 0xc3e345f9.
//
// Solidity: function setSignersRep(address _signersRep) returns()
func (_Bridge *BridgeTransactor) SetSignersRep(opts *bind.TransactOpts, _signersRep common.Address) (*types.Transaction, error) {
	return _Bridge.contract.Transact(opts, "setSignersRep", _signersRep)
}

// SetSignersRep is a paid mutator transaction binding the contract method 0xc3e345f9.
//
// Solidity: function setSignersRep(address _signersRep) returns()
func (_Bridge *BridgeSession) SetSignersRep(_signersRep common.Address) (*types.Transaction, error) {
	return _Bridge.Contract.SetSignersRep(&_Bridge.TransactOpts, _signersRep)
}

// SetSignersRep is a paid mutator transaction binding the contract method 0xc3e345f9.
//
// Solidity: function setSignersRep(address _signersRep) returns()
func (_Bridge *BridgeTransactorSession) SetSignersRep(_signersRep common.Address) (*types.Transaction, error) {
	return _Bridge.Contract.SetSignersRep(&_Bridge.TransactOpts, _signersRep)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Bridge *BridgeTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Bridge.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Bridge *BridgeSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Bridge.Contract.TransferOwnership(&_Bridge.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Bridge *BridgeTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Bridge.Contract.TransferOwnership(&_Bridge.TransactOpts, newOwner)
}

// WithdrawERC1155 is a paid mutator transaction binding the contract method 0x7a355e0e.
//
// Solidity: function withdrawERC1155(address _token, string _txHash, string _tokenURI, address _nativeChainAddress, uint128 _tokenId, bytes32[] _r, bytes32[] _s, uint8[] _v) returns()
func (_Bridge *BridgeTransactor) WithdrawERC1155(opts *bind.TransactOpts, _token common.Address, _txHash string, _tokenURI string, _nativeChainAddress common.Address, _tokenId *big.Int, _r [][32]byte, _s [][32]byte, _v []uint8) (*types.Transaction, error) {
	return _Bridge.contract.Transact(opts, "withdrawERC1155", _token, _txHash, _tokenURI, _nativeChainAddress, _tokenId, _r, _s, _v)
}

// WithdrawERC1155 is a paid mutator transaction binding the contract method 0x7a355e0e.
//
// Solidity: function withdrawERC1155(address _token, string _txHash, string _tokenURI, address _nativeChainAddress, uint128 _tokenId, bytes32[] _r, bytes32[] _s, uint8[] _v) returns()
func (_Bridge *BridgeSession) WithdrawERC1155(_token common.Address, _txHash string, _tokenURI string, _nativeChainAddress common.Address, _tokenId *big.Int, _r [][32]byte, _s [][32]byte, _v []uint8) (*types.Transaction, error) {
	return _Bridge.Contract.WithdrawERC1155(&_Bridge.TransactOpts, _token, _txHash, _tokenURI, _nativeChainAddress, _tokenId, _r, _s, _v)
}

// WithdrawERC1155 is a paid mutator transaction binding the contract method 0x7a355e0e.
//
// Solidity: function withdrawERC1155(address _token, string _txHash, string _tokenURI, address _nativeChainAddress, uint128 _tokenId, bytes32[] _r, bytes32[] _s, uint8[] _v) returns()
func (_Bridge *BridgeTransactorSession) WithdrawERC1155(_token common.Address, _txHash string, _tokenURI string, _nativeChainAddress common.Address, _tokenId *big.Int, _r [][32]byte, _s [][32]byte, _v []uint8) (*types.Transaction, error) {
	return _Bridge.Contract.WithdrawERC1155(&_Bridge.TransactOpts, _token, _txHash, _tokenURI, _nativeChainAddress, _tokenId, _r, _s, _v)
}

// WithdrawERC20 is a paid mutator transaction binding the contract method 0x44004cc1.
//
// Solidity: function withdrawERC20(address _token, address _receiver, uint256 _amount) returns()
func (_Bridge *BridgeTransactor) WithdrawERC20(opts *bind.TransactOpts, _token common.Address, _receiver common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Bridge.contract.Transact(opts, "withdrawERC20", _token, _receiver, _amount)
}

// WithdrawERC20 is a paid mutator transaction binding the contract method 0x44004cc1.
//
// Solidity: function withdrawERC20(address _token, address _receiver, uint256 _amount) returns()
func (_Bridge *BridgeSession) WithdrawERC20(_token common.Address, _receiver common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Bridge.Contract.WithdrawERC20(&_Bridge.TransactOpts, _token, _receiver, _amount)
}

// WithdrawERC20 is a paid mutator transaction binding the contract method 0x44004cc1.
//
// Solidity: function withdrawERC20(address _token, address _receiver, uint256 _amount) returns()
func (_Bridge *BridgeTransactorSession) WithdrawERC20(_token common.Address, _receiver common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Bridge.Contract.WithdrawERC20(&_Bridge.TransactOpts, _token, _receiver, _amount)
}

// WithdrawERC200 is a paid mutator transaction binding the contract method 0xf7f668ef.
//
// Solidity: function withdrawERC20(address _token, string _txHash, uint256 _amount, bytes32[] _r, bytes32[] _s, uint8[] _v) returns()
func (_Bridge *BridgeTransactor) WithdrawERC200(opts *bind.TransactOpts, _token common.Address, _txHash string, _amount *big.Int, _r [][32]byte, _s [][32]byte, _v []uint8) (*types.Transaction, error) {
	return _Bridge.contract.Transact(opts, "withdrawERC200", _token, _txHash, _amount, _r, _s, _v)
}

// WithdrawERC200 is a paid mutator transaction binding the contract method 0xf7f668ef.
//
// Solidity: function withdrawERC20(address _token, string _txHash, uint256 _amount, bytes32[] _r, bytes32[] _s, uint8[] _v) returns()
func (_Bridge *BridgeSession) WithdrawERC200(_token common.Address, _txHash string, _amount *big.Int, _r [][32]byte, _s [][32]byte, _v []uint8) (*types.Transaction, error) {
	return _Bridge.Contract.WithdrawERC200(&_Bridge.TransactOpts, _token, _txHash, _amount, _r, _s, _v)
}

// WithdrawERC200 is a paid mutator transaction binding the contract method 0xf7f668ef.
//
// Solidity: function withdrawERC20(address _token, string _txHash, uint256 _amount, bytes32[] _r, bytes32[] _s, uint8[] _v) returns()
func (_Bridge *BridgeTransactorSession) WithdrawERC200(_token common.Address, _txHash string, _amount *big.Int, _r [][32]byte, _s [][32]byte, _v []uint8) (*types.Transaction, error) {
	return _Bridge.Contract.WithdrawERC200(&_Bridge.TransactOpts, _token, _txHash, _amount, _r, _s, _v)
}

// WithdrawERC721 is a paid mutator transaction binding the contract method 0x177b5521.
//
// Solidity: function withdrawERC721(address _token, string _txHash, uint256 _tokenID, bytes32[] _r, bytes32[] _s, uint8[] _v) returns()
func (_Bridge *BridgeTransactor) WithdrawERC721(opts *bind.TransactOpts, _token common.Address, _txHash string, _tokenID *big.Int, _r [][32]byte, _s [][32]byte, _v []uint8) (*types.Transaction, error) {
	return _Bridge.contract.Transact(opts, "withdrawERC721", _token, _txHash, _tokenID, _r, _s, _v)
}

// WithdrawERC721 is a paid mutator transaction binding the contract method 0x177b5521.
//
// Solidity: function withdrawERC721(address _token, string _txHash, uint256 _tokenID, bytes32[] _r, bytes32[] _s, uint8[] _v) returns()
func (_Bridge *BridgeSession) WithdrawERC721(_token common.Address, _txHash string, _tokenID *big.Int, _r [][32]byte, _s [][32]byte, _v []uint8) (*types.Transaction, error) {
	return _Bridge.Contract.WithdrawERC721(&_Bridge.TransactOpts, _token, _txHash, _tokenID, _r, _s, _v)
}

// WithdrawERC721 is a paid mutator transaction binding the contract method 0x177b5521.
//
// Solidity: function withdrawERC721(address _token, string _txHash, uint256 _tokenID, bytes32[] _r, bytes32[] _s, uint8[] _v) returns()
func (_Bridge *BridgeTransactorSession) WithdrawERC721(_token common.Address, _txHash string, _tokenID *big.Int, _r [][32]byte, _s [][32]byte, _v []uint8) (*types.Transaction, error) {
	return _Bridge.Contract.WithdrawERC721(&_Bridge.TransactOpts, _token, _txHash, _tokenID, _r, _s, _v)
}

// WithdrawERC7210 is a paid mutator transaction binding the contract method 0x4025feb2.
//
// Solidity: function withdrawERC721(address _token, address _receiver, uint256 _amount) returns()
func (_Bridge *BridgeTransactor) WithdrawERC7210(opts *bind.TransactOpts, _token common.Address, _receiver common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Bridge.contract.Transact(opts, "withdrawERC7210", _token, _receiver, _amount)
}

// WithdrawERC7210 is a paid mutator transaction binding the contract method 0x4025feb2.
//
// Solidity: function withdrawERC721(address _token, address _receiver, uint256 _amount) returns()
func (_Bridge *BridgeSession) WithdrawERC7210(_token common.Address, _receiver common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Bridge.Contract.WithdrawERC7210(&_Bridge.TransactOpts, _token, _receiver, _amount)
}

// WithdrawERC7210 is a paid mutator transaction binding the contract method 0x4025feb2.
//
// Solidity: function withdrawERC721(address _token, address _receiver, uint256 _amount) returns()
func (_Bridge *BridgeTransactorSession) WithdrawERC7210(_token common.Address, _receiver common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Bridge.Contract.WithdrawERC7210(&_Bridge.TransactOpts, _token, _receiver, _amount)
}

// WithdrawNative is a paid mutator transaction binding the contract method 0x07b18bde.
//
// Solidity: function withdrawNative(address _receiver, uint256 _amount) returns()
func (_Bridge *BridgeTransactor) WithdrawNative(opts *bind.TransactOpts, _receiver common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Bridge.contract.Transact(opts, "withdrawNative", _receiver, _amount)
}

// WithdrawNative is a paid mutator transaction binding the contract method 0x07b18bde.
//
// Solidity: function withdrawNative(address _receiver, uint256 _amount) returns()
func (_Bridge *BridgeSession) WithdrawNative(_receiver common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Bridge.Contract.WithdrawNative(&_Bridge.TransactOpts, _receiver, _amount)
}

// WithdrawNative is a paid mutator transaction binding the contract method 0x07b18bde.
//
// Solidity: function withdrawNative(address _receiver, uint256 _amount) returns()
func (_Bridge *BridgeTransactorSession) WithdrawNative(_receiver common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Bridge.Contract.WithdrawNative(&_Bridge.TransactOpts, _receiver, _amount)
}

// WithdrawNative0 is a paid mutator transaction binding the contract method 0xd7d07f73.
//
// Solidity: function withdrawNative(string _txHash, uint256 _amount, bytes32[] _r, bytes32[] _s, uint8[] _v) returns()
func (_Bridge *BridgeTransactor) WithdrawNative0(opts *bind.TransactOpts, _txHash string, _amount *big.Int, _r [][32]byte, _s [][32]byte, _v []uint8) (*types.Transaction, error) {
	return _Bridge.contract.Transact(opts, "withdrawNative0", _txHash, _amount, _r, _s, _v)
}

// WithdrawNative0 is a paid mutator transaction binding the contract method 0xd7d07f73.
//
// Solidity: function withdrawNative(string _txHash, uint256 _amount, bytes32[] _r, bytes32[] _s, uint8[] _v) returns()
func (_Bridge *BridgeSession) WithdrawNative0(_txHash string, _amount *big.Int, _r [][32]byte, _s [][32]byte, _v []uint8) (*types.Transaction, error) {
	return _Bridge.Contract.WithdrawNative0(&_Bridge.TransactOpts, _txHash, _amount, _r, _s, _v)
}

// WithdrawNative0 is a paid mutator transaction binding the contract method 0xd7d07f73.
//
// Solidity: function withdrawNative(string _txHash, uint256 _amount, bytes32[] _r, bytes32[] _s, uint8[] _v) returns()
func (_Bridge *BridgeTransactorSession) WithdrawNative0(_txHash string, _amount *big.Int, _r [][32]byte, _s [][32]byte, _v []uint8) (*types.Transaction, error) {
	return _Bridge.Contract.WithdrawNative0(&_Bridge.TransactOpts, _txHash, _amount, _r, _s, _v)
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

// BridgeDepositERC20Iterator is returned from FilterDepositERC20 and is used to iterate over the raw logs and unpacked data for DepositERC20 events raised by the Bridge contract.
type BridgeDepositERC20Iterator struct {
	Event *BridgeDepositERC20 // Event containing the contract specifics and raw log

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
func (it *BridgeDepositERC20Iterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BridgeDepositERC20)
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
		it.Event = new(BridgeDepositERC20)
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
func (it *BridgeDepositERC20Iterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BridgeDepositERC20Iterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BridgeDepositERC20 represents a DepositERC20 event raised by the Bridge contract.
type BridgeDepositERC20 struct {
	Receiver common.Hash
	Token    common.Address
	Amount   *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterDepositERC20 is a free log retrieval operation binding the contract event 0x520e128d3d305b7cfe6d8293bbca1d10efa7c64379fdf0d4998362f3545dc646.
//
// Solidity: event DepositERC20(string indexed receiver, address indexed token, uint256 amount)
func (_Bridge *BridgeFilterer) FilterDepositERC20(opts *bind.FilterOpts, receiver []string, token []common.Address) (*BridgeDepositERC20Iterator, error) {

	var receiverRule []interface{}
	for _, receiverItem := range receiver {
		receiverRule = append(receiverRule, receiverItem)
	}
	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _Bridge.contract.FilterLogs(opts, "DepositERC20", receiverRule, tokenRule)
	if err != nil {
		return nil, err
	}
	return &BridgeDepositERC20Iterator{contract: _Bridge.contract, event: "DepositERC20", logs: logs, sub: sub}, nil
}

// WatchDepositERC20 is a free log subscription operation binding the contract event 0x520e128d3d305b7cfe6d8293bbca1d10efa7c64379fdf0d4998362f3545dc646.
//
// Solidity: event DepositERC20(string indexed receiver, address indexed token, uint256 amount)
func (_Bridge *BridgeFilterer) WatchDepositERC20(opts *bind.WatchOpts, sink chan<- *BridgeDepositERC20, receiver []string, token []common.Address) (event.Subscription, error) {

	var receiverRule []interface{}
	for _, receiverItem := range receiver {
		receiverRule = append(receiverRule, receiverItem)
	}
	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _Bridge.contract.WatchLogs(opts, "DepositERC20", receiverRule, tokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BridgeDepositERC20)
				if err := _Bridge.contract.UnpackLog(event, "DepositERC20", log); err != nil {
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

// ParseDepositERC20 is a log parse operation binding the contract event 0x520e128d3d305b7cfe6d8293bbca1d10efa7c64379fdf0d4998362f3545dc646.
//
// Solidity: event DepositERC20(string indexed receiver, address indexed token, uint256 amount)
func (_Bridge *BridgeFilterer) ParseDepositERC20(log types.Log) (*BridgeDepositERC20, error) {
	event := new(BridgeDepositERC20)
	if err := _Bridge.contract.UnpackLog(event, "DepositERC20", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BridgeHashAddedIterator is returned from FilterHashAdded and is used to iterate over the raw logs and unpacked data for HashAdded events raised by the Bridge contract.
type BridgeHashAddedIterator struct {
	Event *BridgeHashAdded // Event containing the contract specifics and raw log

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
func (it *BridgeHashAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BridgeHashAdded)
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
		it.Event = new(BridgeHashAdded)
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
func (it *BridgeHashAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BridgeHashAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BridgeHashAdded represents a HashAdded event raised by the Bridge contract.
type BridgeHashAdded struct {
	Hash string
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterHashAdded is a free log retrieval operation binding the contract event 0xe0c9f5e6f5abddac86dac0e02afc9f3fda7b7fc6d9454a13c51fcb28621e1e5f.
//
// Solidity: event HashAdded(string hash)
func (_Bridge *BridgeFilterer) FilterHashAdded(opts *bind.FilterOpts) (*BridgeHashAddedIterator, error) {

	logs, sub, err := _Bridge.contract.FilterLogs(opts, "HashAdded")
	if err != nil {
		return nil, err
	}
	return &BridgeHashAddedIterator{contract: _Bridge.contract, event: "HashAdded", logs: logs, sub: sub}, nil
}

// WatchHashAdded is a free log subscription operation binding the contract event 0xe0c9f5e6f5abddac86dac0e02afc9f3fda7b7fc6d9454a13c51fcb28621e1e5f.
//
// Solidity: event HashAdded(string hash)
func (_Bridge *BridgeFilterer) WatchHashAdded(opts *bind.WatchOpts, sink chan<- *BridgeHashAdded) (event.Subscription, error) {

	logs, sub, err := _Bridge.contract.WatchLogs(opts, "HashAdded")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BridgeHashAdded)
				if err := _Bridge.contract.UnpackLog(event, "HashAdded", log); err != nil {
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
func (_Bridge *BridgeFilterer) ParseHashAdded(log types.Log) (*BridgeHashAdded, error) {
	event := new(BridgeHashAdded)
	if err := _Bridge.contract.UnpackLog(event, "HashAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BridgeOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Bridge contract.
type BridgeOwnershipTransferredIterator struct {
	Event *BridgeOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *BridgeOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BridgeOwnershipTransferred)
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
		it.Event = new(BridgeOwnershipTransferred)
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
func (it *BridgeOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BridgeOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BridgeOwnershipTransferred represents a OwnershipTransferred event raised by the Bridge contract.
type BridgeOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Bridge *BridgeFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*BridgeOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Bridge.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &BridgeOwnershipTransferredIterator{contract: _Bridge.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Bridge *BridgeFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *BridgeOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Bridge.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BridgeOwnershipTransferred)
				if err := _Bridge.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_Bridge *BridgeFilterer) ParseOwnershipTransferred(log types.Log) (*BridgeOwnershipTransferred, error) {
	event := new(BridgeOwnershipTransferred)
	if err := _Bridge.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BridgeReceivedIterator is returned from FilterReceived and is used to iterate over the raw logs and unpacked data for Received events raised by the Bridge contract.
type BridgeReceivedIterator struct {
	Event *BridgeReceived // Event containing the contract specifics and raw log

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
func (it *BridgeReceivedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BridgeReceived)
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
		it.Event = new(BridgeReceived)
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
func (it *BridgeReceivedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BridgeReceivedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BridgeReceived represents a Received event raised by the Bridge contract.
type BridgeReceived struct {
	Arg0 common.Address
	Arg1 *big.Int
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterReceived is a free log retrieval operation binding the contract event 0x88a5966d370b9919b20f3e2c13ff65706f196a4e32cc2c12bf57088f88525874.
//
// Solidity: event Received(address indexed arg0, uint256 arg1)
func (_Bridge *BridgeFilterer) FilterReceived(opts *bind.FilterOpts, arg0 []common.Address) (*BridgeReceivedIterator, error) {

	var arg0Rule []interface{}
	for _, arg0Item := range arg0 {
		arg0Rule = append(arg0Rule, arg0Item)
	}

	logs, sub, err := _Bridge.contract.FilterLogs(opts, "Received", arg0Rule)
	if err != nil {
		return nil, err
	}
	return &BridgeReceivedIterator{contract: _Bridge.contract, event: "Received", logs: logs, sub: sub}, nil
}

// WatchReceived is a free log subscription operation binding the contract event 0x88a5966d370b9919b20f3e2c13ff65706f196a4e32cc2c12bf57088f88525874.
//
// Solidity: event Received(address indexed arg0, uint256 arg1)
func (_Bridge *BridgeFilterer) WatchReceived(opts *bind.WatchOpts, sink chan<- *BridgeReceived, arg0 []common.Address) (event.Subscription, error) {

	var arg0Rule []interface{}
	for _, arg0Item := range arg0 {
		arg0Rule = append(arg0Rule, arg0Item)
	}

	logs, sub, err := _Bridge.contract.WatchLogs(opts, "Received", arg0Rule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BridgeReceived)
				if err := _Bridge.contract.UnpackLog(event, "Received", log); err != nil {
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

// ParseReceived is a log parse operation binding the contract event 0x88a5966d370b9919b20f3e2c13ff65706f196a4e32cc2c12bf57088f88525874.
//
// Solidity: event Received(address indexed arg0, uint256 arg1)
func (_Bridge *BridgeFilterer) ParseReceived(log types.Log) (*BridgeReceived, error) {
	event := new(BridgeReceived)
	if err := _Bridge.contract.UnpackLog(event, "Received", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
