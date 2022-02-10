// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package unchain

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

// UnchainMetaData contains all meta data concerning the Unchain contract.
var UnchainMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"ApprovalForAll\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"ids\",\"type\":\"uint256[]\"},{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"values\",\"type\":\"uint256[]\"}],\"name\":\"TransferBatch\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"token\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"tokenID\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"senderBalance\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"receiverBalance\",\"type\":\"uint256\"}],\"name\":\"TransferLog\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"TransferSingle\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"value\",\"type\":\"string\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"URI\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"accounts\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"ids\",\"type\":\"uint256[]\"}],\"name\":\"balanceOfBatch\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_name\",\"type\":\"string\"}],\"name\":\"createUniqueToken\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_name\",\"type\":\"string\"}],\"name\":\"getBalanceByName\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"isApprovedForAll\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_name\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"mint\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256[]\",\"name\":\"ids\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"amounts\",\"type\":\"uint256[]\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"safeBatchTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"setApprovalForAll\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"tokenNames\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"name\":\"tokens\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"tokensInCirculation\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_receiver\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"_name\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"transferTokens\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"uri\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// UnchainABI is the input ABI used to generate the binding from.
// Deprecated: Use UnchainMetaData.ABI instead.
var UnchainABI = UnchainMetaData.ABI

// Unchain is an auto generated Go binding around an Ethereum contract.
type Unchain struct {
	UnchainCaller     // Read-only binding to the contract
	UnchainTransactor // Write-only binding to the contract
	UnchainFilterer   // Log filterer for contract events
}

// UnchainCaller is an auto generated read-only Go binding around an Ethereum contract.
type UnchainCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// UnchainTransactor is an auto generated write-only Go binding around an Ethereum contract.
type UnchainTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// UnchainFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type UnchainFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// UnchainSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type UnchainSession struct {
	Contract     *Unchain          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// UnchainCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type UnchainCallerSession struct {
	Contract *UnchainCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// UnchainTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type UnchainTransactorSession struct {
	Contract     *UnchainTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// UnchainRaw is an auto generated low-level Go binding around an Ethereum contract.
type UnchainRaw struct {
	Contract *Unchain // Generic contract binding to access the raw methods on
}

// UnchainCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type UnchainCallerRaw struct {
	Contract *UnchainCaller // Generic read-only contract binding to access the raw methods on
}

// UnchainTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type UnchainTransactorRaw struct {
	Contract *UnchainTransactor // Generic write-only contract binding to access the raw methods on
}

// NewUnchain creates a new instance of Unchain, bound to a specific deployed contract.
func NewUnchain(address common.Address, backend bind.ContractBackend) (*Unchain, error) {
	contract, err := bindUnchain(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Unchain{UnchainCaller: UnchainCaller{contract: contract}, UnchainTransactor: UnchainTransactor{contract: contract}, UnchainFilterer: UnchainFilterer{contract: contract}}, nil
}

// NewUnchainCaller creates a new read-only instance of Unchain, bound to a specific deployed contract.
func NewUnchainCaller(address common.Address, caller bind.ContractCaller) (*UnchainCaller, error) {
	contract, err := bindUnchain(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &UnchainCaller{contract: contract}, nil
}

// NewUnchainTransactor creates a new write-only instance of Unchain, bound to a specific deployed contract.
func NewUnchainTransactor(address common.Address, transactor bind.ContractTransactor) (*UnchainTransactor, error) {
	contract, err := bindUnchain(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &UnchainTransactor{contract: contract}, nil
}

// NewUnchainFilterer creates a new log filterer instance of Unchain, bound to a specific deployed contract.
func NewUnchainFilterer(address common.Address, filterer bind.ContractFilterer) (*UnchainFilterer, error) {
	contract, err := bindUnchain(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &UnchainFilterer{contract: contract}, nil
}

// bindUnchain binds a generic wrapper to an already deployed contract.
func bindUnchain(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(UnchainABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Unchain *UnchainRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Unchain.Contract.UnchainCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Unchain *UnchainRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Unchain.Contract.UnchainTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Unchain *UnchainRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Unchain.Contract.UnchainTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Unchain *UnchainCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Unchain.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Unchain *UnchainTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Unchain.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Unchain *UnchainTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Unchain.Contract.contract.Transact(opts, method, params...)
}

// BalanceOf is a free data retrieval call binding the contract method 0x00fdd58e.
//
// Solidity: function balanceOf(address account, uint256 id) view returns(uint256)
func (_Unchain *UnchainCaller) BalanceOf(opts *bind.CallOpts, account common.Address, id *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Unchain.contract.Call(opts, &out, "balanceOf", account, id)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x00fdd58e.
//
// Solidity: function balanceOf(address account, uint256 id) view returns(uint256)
func (_Unchain *UnchainSession) BalanceOf(account common.Address, id *big.Int) (*big.Int, error) {
	return _Unchain.Contract.BalanceOf(&_Unchain.CallOpts, account, id)
}

// BalanceOf is a free data retrieval call binding the contract method 0x00fdd58e.
//
// Solidity: function balanceOf(address account, uint256 id) view returns(uint256)
func (_Unchain *UnchainCallerSession) BalanceOf(account common.Address, id *big.Int) (*big.Int, error) {
	return _Unchain.Contract.BalanceOf(&_Unchain.CallOpts, account, id)
}

// BalanceOfBatch is a free data retrieval call binding the contract method 0x4e1273f4.
//
// Solidity: function balanceOfBatch(address[] accounts, uint256[] ids) view returns(uint256[])
func (_Unchain *UnchainCaller) BalanceOfBatch(opts *bind.CallOpts, accounts []common.Address, ids []*big.Int) ([]*big.Int, error) {
	var out []interface{}
	err := _Unchain.contract.Call(opts, &out, "balanceOfBatch", accounts, ids)

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// BalanceOfBatch is a free data retrieval call binding the contract method 0x4e1273f4.
//
// Solidity: function balanceOfBatch(address[] accounts, uint256[] ids) view returns(uint256[])
func (_Unchain *UnchainSession) BalanceOfBatch(accounts []common.Address, ids []*big.Int) ([]*big.Int, error) {
	return _Unchain.Contract.BalanceOfBatch(&_Unchain.CallOpts, accounts, ids)
}

// BalanceOfBatch is a free data retrieval call binding the contract method 0x4e1273f4.
//
// Solidity: function balanceOfBatch(address[] accounts, uint256[] ids) view returns(uint256[])
func (_Unchain *UnchainCallerSession) BalanceOfBatch(accounts []common.Address, ids []*big.Int) ([]*big.Int, error) {
	return _Unchain.Contract.BalanceOfBatch(&_Unchain.CallOpts, accounts, ids)
}

// GetBalanceByName is a free data retrieval call binding the contract method 0xca1ae01c.
//
// Solidity: function getBalanceByName(string _name) view returns(uint256)
func (_Unchain *UnchainCaller) GetBalanceByName(opts *bind.CallOpts, _name string) (*big.Int, error) {
	var out []interface{}
	err := _Unchain.contract.Call(opts, &out, "getBalanceByName", _name)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetBalanceByName is a free data retrieval call binding the contract method 0xca1ae01c.
//
// Solidity: function getBalanceByName(string _name) view returns(uint256)
func (_Unchain *UnchainSession) GetBalanceByName(_name string) (*big.Int, error) {
	return _Unchain.Contract.GetBalanceByName(&_Unchain.CallOpts, _name)
}

// GetBalanceByName is a free data retrieval call binding the contract method 0xca1ae01c.
//
// Solidity: function getBalanceByName(string _name) view returns(uint256)
func (_Unchain *UnchainCallerSession) GetBalanceByName(_name string) (*big.Int, error) {
	return _Unchain.Contract.GetBalanceByName(&_Unchain.CallOpts, _name)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address account, address operator) view returns(bool)
func (_Unchain *UnchainCaller) IsApprovedForAll(opts *bind.CallOpts, account common.Address, operator common.Address) (bool, error) {
	var out []interface{}
	err := _Unchain.contract.Call(opts, &out, "isApprovedForAll", account, operator)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address account, address operator) view returns(bool)
func (_Unchain *UnchainSession) IsApprovedForAll(account common.Address, operator common.Address) (bool, error) {
	return _Unchain.Contract.IsApprovedForAll(&_Unchain.CallOpts, account, operator)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address account, address operator) view returns(bool)
func (_Unchain *UnchainCallerSession) IsApprovedForAll(account common.Address, operator common.Address) (bool, error) {
	return _Unchain.Contract.IsApprovedForAll(&_Unchain.CallOpts, account, operator)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Unchain *UnchainCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _Unchain.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Unchain *UnchainSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _Unchain.Contract.SupportsInterface(&_Unchain.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Unchain *UnchainCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _Unchain.Contract.SupportsInterface(&_Unchain.CallOpts, interfaceId)
}

// TokenNames is a free data retrieval call binding the contract method 0x310495ab.
//
// Solidity: function tokenNames(uint256 ) view returns(string)
func (_Unchain *UnchainCaller) TokenNames(opts *bind.CallOpts, arg0 *big.Int) (string, error) {
	var out []interface{}
	err := _Unchain.contract.Call(opts, &out, "tokenNames", arg0)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// TokenNames is a free data retrieval call binding the contract method 0x310495ab.
//
// Solidity: function tokenNames(uint256 ) view returns(string)
func (_Unchain *UnchainSession) TokenNames(arg0 *big.Int) (string, error) {
	return _Unchain.Contract.TokenNames(&_Unchain.CallOpts, arg0)
}

// TokenNames is a free data retrieval call binding the contract method 0x310495ab.
//
// Solidity: function tokenNames(uint256 ) view returns(string)
func (_Unchain *UnchainCallerSession) TokenNames(arg0 *big.Int) (string, error) {
	return _Unchain.Contract.TokenNames(&_Unchain.CallOpts, arg0)
}

// Tokens is a free data retrieval call binding the contract method 0x04c2320b.
//
// Solidity: function tokens(string ) view returns(uint256)
func (_Unchain *UnchainCaller) Tokens(opts *bind.CallOpts, arg0 string) (*big.Int, error) {
	var out []interface{}
	err := _Unchain.contract.Call(opts, &out, "tokens", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Tokens is a free data retrieval call binding the contract method 0x04c2320b.
//
// Solidity: function tokens(string ) view returns(uint256)
func (_Unchain *UnchainSession) Tokens(arg0 string) (*big.Int, error) {
	return _Unchain.Contract.Tokens(&_Unchain.CallOpts, arg0)
}

// Tokens is a free data retrieval call binding the contract method 0x04c2320b.
//
// Solidity: function tokens(string ) view returns(uint256)
func (_Unchain *UnchainCallerSession) Tokens(arg0 string) (*big.Int, error) {
	return _Unchain.Contract.Tokens(&_Unchain.CallOpts, arg0)
}

// TokensInCirculation is a free data retrieval call binding the contract method 0xf5a0d6c0.
//
// Solidity: function tokensInCirculation() view returns(uint256)
func (_Unchain *UnchainCaller) TokensInCirculation(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Unchain.contract.Call(opts, &out, "tokensInCirculation")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TokensInCirculation is a free data retrieval call binding the contract method 0xf5a0d6c0.
//
// Solidity: function tokensInCirculation() view returns(uint256)
func (_Unchain *UnchainSession) TokensInCirculation() (*big.Int, error) {
	return _Unchain.Contract.TokensInCirculation(&_Unchain.CallOpts)
}

// TokensInCirculation is a free data retrieval call binding the contract method 0xf5a0d6c0.
//
// Solidity: function tokensInCirculation() view returns(uint256)
func (_Unchain *UnchainCallerSession) TokensInCirculation() (*big.Int, error) {
	return _Unchain.Contract.TokensInCirculation(&_Unchain.CallOpts)
}

// Uri is a free data retrieval call binding the contract method 0x0e89341c.
//
// Solidity: function uri(uint256 ) view returns(string)
func (_Unchain *UnchainCaller) Uri(opts *bind.CallOpts, arg0 *big.Int) (string, error) {
	var out []interface{}
	err := _Unchain.contract.Call(opts, &out, "uri", arg0)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Uri is a free data retrieval call binding the contract method 0x0e89341c.
//
// Solidity: function uri(uint256 ) view returns(string)
func (_Unchain *UnchainSession) Uri(arg0 *big.Int) (string, error) {
	return _Unchain.Contract.Uri(&_Unchain.CallOpts, arg0)
}

// Uri is a free data retrieval call binding the contract method 0x0e89341c.
//
// Solidity: function uri(uint256 ) view returns(string)
func (_Unchain *UnchainCallerSession) Uri(arg0 *big.Int) (string, error) {
	return _Unchain.Contract.Uri(&_Unchain.CallOpts, arg0)
}

// CreateUniqueToken is a paid mutator transaction binding the contract method 0x8653be85.
//
// Solidity: function createUniqueToken(string _name) returns(string)
func (_Unchain *UnchainTransactor) CreateUniqueToken(opts *bind.TransactOpts, _name string) (*types.Transaction, error) {
	return _Unchain.contract.Transact(opts, "createUniqueToken", _name)
}

// CreateUniqueToken is a paid mutator transaction binding the contract method 0x8653be85.
//
// Solidity: function createUniqueToken(string _name) returns(string)
func (_Unchain *UnchainSession) CreateUniqueToken(_name string) (*types.Transaction, error) {
	return _Unchain.Contract.CreateUniqueToken(&_Unchain.TransactOpts, _name)
}

// CreateUniqueToken is a paid mutator transaction binding the contract method 0x8653be85.
//
// Solidity: function createUniqueToken(string _name) returns(string)
func (_Unchain *UnchainTransactorSession) CreateUniqueToken(_name string) (*types.Transaction, error) {
	return _Unchain.Contract.CreateUniqueToken(&_Unchain.TransactOpts, _name)
}

// Mint is a paid mutator transaction binding the contract method 0x056b01ce.
//
// Solidity: function mint(string _name, uint256 _amount) returns()
func (_Unchain *UnchainTransactor) Mint(opts *bind.TransactOpts, _name string, _amount *big.Int) (*types.Transaction, error) {
	return _Unchain.contract.Transact(opts, "mint", _name, _amount)
}

// Mint is a paid mutator transaction binding the contract method 0x056b01ce.
//
// Solidity: function mint(string _name, uint256 _amount) returns()
func (_Unchain *UnchainSession) Mint(_name string, _amount *big.Int) (*types.Transaction, error) {
	return _Unchain.Contract.Mint(&_Unchain.TransactOpts, _name, _amount)
}

// Mint is a paid mutator transaction binding the contract method 0x056b01ce.
//
// Solidity: function mint(string _name, uint256 _amount) returns()
func (_Unchain *UnchainTransactorSession) Mint(_name string, _amount *big.Int) (*types.Transaction, error) {
	return _Unchain.Contract.Mint(&_Unchain.TransactOpts, _name, _amount)
}

// SafeBatchTransferFrom is a paid mutator transaction binding the contract method 0x2eb2c2d6.
//
// Solidity: function safeBatchTransferFrom(address from, address to, uint256[] ids, uint256[] amounts, bytes data) returns()
func (_Unchain *UnchainTransactor) SafeBatchTransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, ids []*big.Int, amounts []*big.Int, data []byte) (*types.Transaction, error) {
	return _Unchain.contract.Transact(opts, "safeBatchTransferFrom", from, to, ids, amounts, data)
}

// SafeBatchTransferFrom is a paid mutator transaction binding the contract method 0x2eb2c2d6.
//
// Solidity: function safeBatchTransferFrom(address from, address to, uint256[] ids, uint256[] amounts, bytes data) returns()
func (_Unchain *UnchainSession) SafeBatchTransferFrom(from common.Address, to common.Address, ids []*big.Int, amounts []*big.Int, data []byte) (*types.Transaction, error) {
	return _Unchain.Contract.SafeBatchTransferFrom(&_Unchain.TransactOpts, from, to, ids, amounts, data)
}

// SafeBatchTransferFrom is a paid mutator transaction binding the contract method 0x2eb2c2d6.
//
// Solidity: function safeBatchTransferFrom(address from, address to, uint256[] ids, uint256[] amounts, bytes data) returns()
func (_Unchain *UnchainTransactorSession) SafeBatchTransferFrom(from common.Address, to common.Address, ids []*big.Int, amounts []*big.Int, data []byte) (*types.Transaction, error) {
	return _Unchain.Contract.SafeBatchTransferFrom(&_Unchain.TransactOpts, from, to, ids, amounts, data)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0xf242432a.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 id, uint256 amount, bytes data) returns()
func (_Unchain *UnchainTransactor) SafeTransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, id *big.Int, amount *big.Int, data []byte) (*types.Transaction, error) {
	return _Unchain.contract.Transact(opts, "safeTransferFrom", from, to, id, amount, data)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0xf242432a.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 id, uint256 amount, bytes data) returns()
func (_Unchain *UnchainSession) SafeTransferFrom(from common.Address, to common.Address, id *big.Int, amount *big.Int, data []byte) (*types.Transaction, error) {
	return _Unchain.Contract.SafeTransferFrom(&_Unchain.TransactOpts, from, to, id, amount, data)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0xf242432a.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 id, uint256 amount, bytes data) returns()
func (_Unchain *UnchainTransactorSession) SafeTransferFrom(from common.Address, to common.Address, id *big.Int, amount *big.Int, data []byte) (*types.Transaction, error) {
	return _Unchain.Contract.SafeTransferFrom(&_Unchain.TransactOpts, from, to, id, amount, data)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_Unchain *UnchainTransactor) SetApprovalForAll(opts *bind.TransactOpts, operator common.Address, approved bool) (*types.Transaction, error) {
	return _Unchain.contract.Transact(opts, "setApprovalForAll", operator, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_Unchain *UnchainSession) SetApprovalForAll(operator common.Address, approved bool) (*types.Transaction, error) {
	return _Unchain.Contract.SetApprovalForAll(&_Unchain.TransactOpts, operator, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_Unchain *UnchainTransactorSession) SetApprovalForAll(operator common.Address, approved bool) (*types.Transaction, error) {
	return _Unchain.Contract.SetApprovalForAll(&_Unchain.TransactOpts, operator, approved)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address _newOwner) returns()
func (_Unchain *UnchainTransactor) TransferOwnership(opts *bind.TransactOpts, _newOwner common.Address) (*types.Transaction, error) {
	return _Unchain.contract.Transact(opts, "transferOwnership", _newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address _newOwner) returns()
func (_Unchain *UnchainSession) TransferOwnership(_newOwner common.Address) (*types.Transaction, error) {
	return _Unchain.Contract.TransferOwnership(&_Unchain.TransactOpts, _newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address _newOwner) returns()
func (_Unchain *UnchainTransactorSession) TransferOwnership(_newOwner common.Address) (*types.Transaction, error) {
	return _Unchain.Contract.TransferOwnership(&_Unchain.TransactOpts, _newOwner)
}

// TransferTokens is a paid mutator transaction binding the contract method 0xc00bb4b2.
//
// Solidity: function transferTokens(address _receiver, string _name, uint256 _amount) returns()
func (_Unchain *UnchainTransactor) TransferTokens(opts *bind.TransactOpts, _receiver common.Address, _name string, _amount *big.Int) (*types.Transaction, error) {
	return _Unchain.contract.Transact(opts, "transferTokens", _receiver, _name, _amount)
}

// TransferTokens is a paid mutator transaction binding the contract method 0xc00bb4b2.
//
// Solidity: function transferTokens(address _receiver, string _name, uint256 _amount) returns()
func (_Unchain *UnchainSession) TransferTokens(_receiver common.Address, _name string, _amount *big.Int) (*types.Transaction, error) {
	return _Unchain.Contract.TransferTokens(&_Unchain.TransactOpts, _receiver, _name, _amount)
}

// TransferTokens is a paid mutator transaction binding the contract method 0xc00bb4b2.
//
// Solidity: function transferTokens(address _receiver, string _name, uint256 _amount) returns()
func (_Unchain *UnchainTransactorSession) TransferTokens(_receiver common.Address, _name string, _amount *big.Int) (*types.Transaction, error) {
	return _Unchain.Contract.TransferTokens(&_Unchain.TransactOpts, _receiver, _name, _amount)
}

// UnchainApprovalForAllIterator is returned from FilterApprovalForAll and is used to iterate over the raw logs and unpacked data for ApprovalForAll events raised by the Unchain contract.
type UnchainApprovalForAllIterator struct {
	Event *UnchainApprovalForAll // Event containing the contract specifics and raw log

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
func (it *UnchainApprovalForAllIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(UnchainApprovalForAll)
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
		it.Event = new(UnchainApprovalForAll)
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
func (it *UnchainApprovalForAllIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *UnchainApprovalForAllIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// UnchainApprovalForAll represents a ApprovalForAll event raised by the Unchain contract.
type UnchainApprovalForAll struct {
	Account  common.Address
	Operator common.Address
	Approved bool
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApprovalForAll is a free log retrieval operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed account, address indexed operator, bool approved)
func (_Unchain *UnchainFilterer) FilterApprovalForAll(opts *bind.FilterOpts, account []common.Address, operator []common.Address) (*UnchainApprovalForAllIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _Unchain.contract.FilterLogs(opts, "ApprovalForAll", accountRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return &UnchainApprovalForAllIterator{contract: _Unchain.contract, event: "ApprovalForAll", logs: logs, sub: sub}, nil
}

// WatchApprovalForAll is a free log subscription operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed account, address indexed operator, bool approved)
func (_Unchain *UnchainFilterer) WatchApprovalForAll(opts *bind.WatchOpts, sink chan<- *UnchainApprovalForAll, account []common.Address, operator []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _Unchain.contract.WatchLogs(opts, "ApprovalForAll", accountRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(UnchainApprovalForAll)
				if err := _Unchain.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
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
// Solidity: event ApprovalForAll(address indexed account, address indexed operator, bool approved)
func (_Unchain *UnchainFilterer) ParseApprovalForAll(log types.Log) (*UnchainApprovalForAll, error) {
	event := new(UnchainApprovalForAll)
	if err := _Unchain.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// UnchainTransferBatchIterator is returned from FilterTransferBatch and is used to iterate over the raw logs and unpacked data for TransferBatch events raised by the Unchain contract.
type UnchainTransferBatchIterator struct {
	Event *UnchainTransferBatch // Event containing the contract specifics and raw log

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
func (it *UnchainTransferBatchIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(UnchainTransferBatch)
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
		it.Event = new(UnchainTransferBatch)
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
func (it *UnchainTransferBatchIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *UnchainTransferBatchIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// UnchainTransferBatch represents a TransferBatch event raised by the Unchain contract.
type UnchainTransferBatch struct {
	Operator common.Address
	From     common.Address
	To       common.Address
	Ids      []*big.Int
	Values   []*big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterTransferBatch is a free log retrieval operation binding the contract event 0x4a39dc06d4c0dbc64b70af90fd698a233a518aa5d07e595d983b8c0526c8f7fb.
//
// Solidity: event TransferBatch(address indexed operator, address indexed from, address indexed to, uint256[] ids, uint256[] values)
func (_Unchain *UnchainFilterer) FilterTransferBatch(opts *bind.FilterOpts, operator []common.Address, from []common.Address, to []common.Address) (*UnchainTransferBatchIterator, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}
	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Unchain.contract.FilterLogs(opts, "TransferBatch", operatorRule, fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &UnchainTransferBatchIterator{contract: _Unchain.contract, event: "TransferBatch", logs: logs, sub: sub}, nil
}

// WatchTransferBatch is a free log subscription operation binding the contract event 0x4a39dc06d4c0dbc64b70af90fd698a233a518aa5d07e595d983b8c0526c8f7fb.
//
// Solidity: event TransferBatch(address indexed operator, address indexed from, address indexed to, uint256[] ids, uint256[] values)
func (_Unchain *UnchainFilterer) WatchTransferBatch(opts *bind.WatchOpts, sink chan<- *UnchainTransferBatch, operator []common.Address, from []common.Address, to []common.Address) (event.Subscription, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}
	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Unchain.contract.WatchLogs(opts, "TransferBatch", operatorRule, fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(UnchainTransferBatch)
				if err := _Unchain.contract.UnpackLog(event, "TransferBatch", log); err != nil {
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

// ParseTransferBatch is a log parse operation binding the contract event 0x4a39dc06d4c0dbc64b70af90fd698a233a518aa5d07e595d983b8c0526c8f7fb.
//
// Solidity: event TransferBatch(address indexed operator, address indexed from, address indexed to, uint256[] ids, uint256[] values)
func (_Unchain *UnchainFilterer) ParseTransferBatch(log types.Log) (*UnchainTransferBatch, error) {
	event := new(UnchainTransferBatch)
	if err := _Unchain.contract.UnpackLog(event, "TransferBatch", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// UnchainTransferLogIterator is returned from FilterTransferLog and is used to iterate over the raw logs and unpacked data for TransferLog events raised by the Unchain contract.
type UnchainTransferLogIterator struct {
	Event *UnchainTransferLog // Event containing the contract specifics and raw log

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
func (it *UnchainTransferLogIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(UnchainTransferLog)
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
		it.Event = new(UnchainTransferLog)
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
func (it *UnchainTransferLogIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *UnchainTransferLogIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// UnchainTransferLog represents a TransferLog event raised by the Unchain contract.
type UnchainTransferLog struct {
	Sender          common.Address
	Receiver        common.Address
	Token           string
	TokenID         *big.Int
	Amount          *big.Int
	SenderBalance   *big.Int
	ReceiverBalance *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterTransferLog is a free log retrieval operation binding the contract event 0xbca92066e2ae354a2a873f30e4a671e173c4092b46d6b26eff8e6d313ebf6598.
//
// Solidity: event TransferLog(address indexed sender, address indexed receiver, string token, uint256 tokenID, uint256 amount, uint256 senderBalance, uint256 receiverBalance)
func (_Unchain *UnchainFilterer) FilterTransferLog(opts *bind.FilterOpts, sender []common.Address, receiver []common.Address) (*UnchainTransferLogIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var receiverRule []interface{}
	for _, receiverItem := range receiver {
		receiverRule = append(receiverRule, receiverItem)
	}

	logs, sub, err := _Unchain.contract.FilterLogs(opts, "TransferLog", senderRule, receiverRule)
	if err != nil {
		return nil, err
	}
	return &UnchainTransferLogIterator{contract: _Unchain.contract, event: "TransferLog", logs: logs, sub: sub}, nil
}

// WatchTransferLog is a free log subscription operation binding the contract event 0xbca92066e2ae354a2a873f30e4a671e173c4092b46d6b26eff8e6d313ebf6598.
//
// Solidity: event TransferLog(address indexed sender, address indexed receiver, string token, uint256 tokenID, uint256 amount, uint256 senderBalance, uint256 receiverBalance)
func (_Unchain *UnchainFilterer) WatchTransferLog(opts *bind.WatchOpts, sink chan<- *UnchainTransferLog, sender []common.Address, receiver []common.Address) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var receiverRule []interface{}
	for _, receiverItem := range receiver {
		receiverRule = append(receiverRule, receiverItem)
	}

	logs, sub, err := _Unchain.contract.WatchLogs(opts, "TransferLog", senderRule, receiverRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(UnchainTransferLog)
				if err := _Unchain.contract.UnpackLog(event, "TransferLog", log); err != nil {
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

// ParseTransferLog is a log parse operation binding the contract event 0xbca92066e2ae354a2a873f30e4a671e173c4092b46d6b26eff8e6d313ebf6598.
//
// Solidity: event TransferLog(address indexed sender, address indexed receiver, string token, uint256 tokenID, uint256 amount, uint256 senderBalance, uint256 receiverBalance)
func (_Unchain *UnchainFilterer) ParseTransferLog(log types.Log) (*UnchainTransferLog, error) {
	event := new(UnchainTransferLog)
	if err := _Unchain.contract.UnpackLog(event, "TransferLog", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// UnchainTransferSingleIterator is returned from FilterTransferSingle and is used to iterate over the raw logs and unpacked data for TransferSingle events raised by the Unchain contract.
type UnchainTransferSingleIterator struct {
	Event *UnchainTransferSingle // Event containing the contract specifics and raw log

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
func (it *UnchainTransferSingleIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(UnchainTransferSingle)
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
		it.Event = new(UnchainTransferSingle)
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
func (it *UnchainTransferSingleIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *UnchainTransferSingleIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// UnchainTransferSingle represents a TransferSingle event raised by the Unchain contract.
type UnchainTransferSingle struct {
	Operator common.Address
	From     common.Address
	To       common.Address
	Id       *big.Int
	Value    *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterTransferSingle is a free log retrieval operation binding the contract event 0xc3d58168c5ae7397731d063d5bbf3d657854427343f4c083240f7aacaa2d0f62.
//
// Solidity: event TransferSingle(address indexed operator, address indexed from, address indexed to, uint256 id, uint256 value)
func (_Unchain *UnchainFilterer) FilterTransferSingle(opts *bind.FilterOpts, operator []common.Address, from []common.Address, to []common.Address) (*UnchainTransferSingleIterator, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}
	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Unchain.contract.FilterLogs(opts, "TransferSingle", operatorRule, fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &UnchainTransferSingleIterator{contract: _Unchain.contract, event: "TransferSingle", logs: logs, sub: sub}, nil
}

// WatchTransferSingle is a free log subscription operation binding the contract event 0xc3d58168c5ae7397731d063d5bbf3d657854427343f4c083240f7aacaa2d0f62.
//
// Solidity: event TransferSingle(address indexed operator, address indexed from, address indexed to, uint256 id, uint256 value)
func (_Unchain *UnchainFilterer) WatchTransferSingle(opts *bind.WatchOpts, sink chan<- *UnchainTransferSingle, operator []common.Address, from []common.Address, to []common.Address) (event.Subscription, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}
	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Unchain.contract.WatchLogs(opts, "TransferSingle", operatorRule, fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(UnchainTransferSingle)
				if err := _Unchain.contract.UnpackLog(event, "TransferSingle", log); err != nil {
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

// ParseTransferSingle is a log parse operation binding the contract event 0xc3d58168c5ae7397731d063d5bbf3d657854427343f4c083240f7aacaa2d0f62.
//
// Solidity: event TransferSingle(address indexed operator, address indexed from, address indexed to, uint256 id, uint256 value)
func (_Unchain *UnchainFilterer) ParseTransferSingle(log types.Log) (*UnchainTransferSingle, error) {
	event := new(UnchainTransferSingle)
	if err := _Unchain.contract.UnpackLog(event, "TransferSingle", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// UnchainURIIterator is returned from FilterURI and is used to iterate over the raw logs and unpacked data for URI events raised by the Unchain contract.
type UnchainURIIterator struct {
	Event *UnchainURI // Event containing the contract specifics and raw log

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
func (it *UnchainURIIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(UnchainURI)
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
		it.Event = new(UnchainURI)
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
func (it *UnchainURIIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *UnchainURIIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// UnchainURI represents a URI event raised by the Unchain contract.
type UnchainURI struct {
	Value string
	Id    *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterURI is a free log retrieval operation binding the contract event 0x6bb7ff708619ba0610cba295a58592e0451dee2622938c8755667688daf3529b.
//
// Solidity: event URI(string value, uint256 indexed id)
func (_Unchain *UnchainFilterer) FilterURI(opts *bind.FilterOpts, id []*big.Int) (*UnchainURIIterator, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _Unchain.contract.FilterLogs(opts, "URI", idRule)
	if err != nil {
		return nil, err
	}
	return &UnchainURIIterator{contract: _Unchain.contract, event: "URI", logs: logs, sub: sub}, nil
}

// WatchURI is a free log subscription operation binding the contract event 0x6bb7ff708619ba0610cba295a58592e0451dee2622938c8755667688daf3529b.
//
// Solidity: event URI(string value, uint256 indexed id)
func (_Unchain *UnchainFilterer) WatchURI(opts *bind.WatchOpts, sink chan<- *UnchainURI, id []*big.Int) (event.Subscription, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _Unchain.contract.WatchLogs(opts, "URI", idRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(UnchainURI)
				if err := _Unchain.contract.UnpackLog(event, "URI", log); err != nil {
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

// ParseURI is a log parse operation binding the contract event 0x6bb7ff708619ba0610cba295a58592e0451dee2622938c8755667688daf3529b.
//
// Solidity: event URI(string value, uint256 indexed id)
func (_Unchain *UnchainFilterer) ParseURI(log types.Log) (*UnchainURI, error) {
	event := new(UnchainURI)
	if err := _Unchain.contract.UnpackLog(event, "URI", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
