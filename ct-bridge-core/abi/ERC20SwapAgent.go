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

// ERC20SwapAgentMetaData contains all meta data concerning the ERC20SwapAgent contract.
var ERC20SwapAgentMetaData = &bind.MetaData{
	ABI: "[\n  {\n    \"anonymous\": false,\n    \"inputs\": [\n      {\n        \"indexed\": true,\n        \"internalType\": \"bytes32\",\n        \"name\": \"swapTxHash\",\n        \"type\": \"bytes32\"\n      },\n      {\n        \"indexed\": true,\n        \"internalType\": \"address\",\n        \"name\": \"tokenAddr\",\n        \"type\": \"address\"\n      },\n      {\n        \"indexed\": true,\n        \"internalType\": \"address\",\n        \"name\": \"recipient\",\n        \"type\": \"address\"\n      },\n      {\n        \"indexed\": false,\n        \"internalType\": \"uint256\",\n        \"name\": \"fromChainId\",\n        \"type\": \"uint256\"\n      },\n      {\n        \"indexed\": false,\n        \"internalType\": \"uint256\",\n        \"name\": \"amount\",\n        \"type\": \"uint256\"\n      }\n    ],\n    \"name\": \"BackwardSwapFilled\",\n    \"type\": \"event\"\n  },\n  {\n    \"anonymous\": false,\n    \"inputs\": [\n      {\n        \"indexed\": true,\n        \"internalType\": \"address\",\n        \"name\": \"mirroredTokenAddr\",\n        \"type\": \"address\"\n      },\n      {\n        \"indexed\": true,\n        \"internalType\": \"address\",\n        \"name\": \"sender\",\n        \"type\": \"address\"\n      },\n      {\n        \"indexed\": true,\n        \"internalType\": \"address\",\n        \"name\": \"recipient\",\n        \"type\": \"address\"\n      },\n      {\n        \"indexed\": false,\n        \"internalType\": \"uint256\",\n        \"name\": \"dstChainId\",\n        \"type\": \"uint256\"\n      },\n      {\n        \"indexed\": false,\n        \"internalType\": \"uint256\",\n        \"name\": \"amount\",\n        \"type\": \"uint256\"\n      },\n      {\n        \"indexed\": false,\n        \"internalType\": \"uint256\",\n        \"name\": \"feeAmount\",\n        \"type\": \"uint256\"\n      }\n    ],\n    \"name\": \"BackwardSwapStarted\",\n    \"type\": \"event\"\n  },\n  {\n    \"anonymous\": false,\n    \"inputs\": [\n      {\n        \"indexed\": true,\n        \"internalType\": \"address\",\n        \"name\": \"previousOwner\",\n        \"type\": \"address\"\n      },\n      {\n        \"indexed\": true,\n        \"internalType\": \"address\",\n        \"name\": \"newOwner\",\n        \"type\": \"address\"\n      }\n    ],\n    \"name\": \"OwnershipTransferred\",\n    \"type\": \"event\"\n  },\n  {\n    \"anonymous\": false,\n    \"inputs\": [\n      {\n        \"indexed\": true,\n        \"internalType\": \"bytes32\",\n        \"name\": \"swapTxHash\",\n        \"type\": \"bytes32\"\n      },\n      {\n        \"indexed\": true,\n        \"internalType\": \"address\",\n        \"name\": \"fromTokenAddr\",\n        \"type\": \"address\"\n      },\n      {\n        \"indexed\": true,\n        \"internalType\": \"address\",\n        \"name\": \"recipient\",\n        \"type\": \"address\"\n      },\n      {\n        \"indexed\": false,\n        \"internalType\": \"address\",\n        \"name\": \"mirroredTokenAddr\",\n        \"type\": \"address\"\n      },\n      {\n        \"indexed\": false,\n        \"internalType\": \"uint256\",\n        \"name\": \"fromChainId\",\n        \"type\": \"uint256\"\n      },\n      {\n        \"indexed\": false,\n        \"internalType\": \"uint256\",\n        \"name\": \"amount\",\n        \"type\": \"uint256\"\n      }\n    ],\n    \"name\": \"SwapFilled\",\n    \"type\": \"event\"\n  },\n  {\n    \"anonymous\": false,\n    \"inputs\": [\n      {\n        \"indexed\": true,\n        \"internalType\": \"bytes32\",\n        \"name\": \"registerTxHash\",\n        \"type\": \"bytes32\"\n      },\n      {\n        \"indexed\": true,\n        \"internalType\": \"address\",\n        \"name\": \"fromTokenAddr\",\n        \"type\": \"address\"\n      },\n      {\n        \"indexed\": true,\n        \"internalType\": \"address\",\n        \"name\": \"mirroredTokenAddr\",\n        \"type\": \"address\"\n      },\n      {\n        \"indexed\": false,\n        \"internalType\": \"uint256\",\n        \"name\": \"fromChainId\",\n        \"type\": \"uint256\"\n      }\n    ],\n    \"name\": \"SwapPairCreated\",\n    \"type\": \"event\"\n  },\n  {\n    \"anonymous\": false,\n    \"inputs\": [\n      {\n        \"indexed\": true,\n        \"internalType\": \"address\",\n        \"name\": \"sponsor\",\n        \"type\": \"address\"\n      },\n      {\n        \"indexed\": true,\n        \"internalType\": \"address\",\n        \"name\": \"tokenAddress\",\n        \"type\": \"address\"\n      },\n      {\n        \"indexed\": false,\n        \"internalType\": \"uint256\",\n        \"name\": \"toChainId\",\n        \"type\": \"uint256\"\n      },\n      {\n        \"indexed\": false,\n        \"internalType\": \"uint256\",\n        \"name\": \"feeAmount\",\n        \"type\": \"uint256\"\n      }\n    ],\n    \"name\": \"SwapPairRegister\",\n    \"type\": \"event\"\n  },\n  {\n    \"anonymous\": false,\n    \"inputs\": [\n      {\n        \"indexed\": true,\n        \"internalType\": \"address\",\n        \"name\": \"tokenAddr\",\n        \"type\": \"address\"\n      },\n      {\n        \"indexed\": true,\n        \"internalType\": \"address\",\n        \"name\": \"sender\",\n        \"type\": \"address\"\n      },\n      {\n        \"indexed\": true,\n        \"internalType\": \"address\",\n        \"name\": \"recipient\",\n        \"type\": \"address\"\n      },\n      {\n        \"indexed\": false,\n        \"internalType\": \"uint256\",\n        \"name\": \"dstChainId\",\n        \"type\": \"uint256\"\n      },\n      {\n        \"indexed\": false,\n        \"internalType\": \"uint256\",\n        \"name\": \"amount\",\n        \"type\": \"uint256\"\n      },\n      {\n        \"indexed\": false,\n        \"internalType\": \"uint256\",\n        \"name\": \"feeAmount\",\n        \"type\": \"uint256\"\n      }\n    ],\n    \"name\": \"SwapStarted\",\n    \"type\": \"event\"\n  },\n  {\n    \"inputs\": [\n      {\n        \"internalType\": \"bytes32\",\n        \"name\": \"registerTxHash\",\n        \"type\": \"bytes32\"\n      },\n      {\n        \"internalType\": \"address\",\n        \"name\": \"fromTokenAddr\",\n        \"type\": \"address\"\n      },\n      {\n        \"internalType\": \"uint256\",\n        \"name\": \"fromChainId\",\n        \"type\": \"uint256\"\n      },\n      {\n        \"internalType\": \"string\",\n        \"name\": \"_name\",\n        \"type\": \"string\"\n      },\n      {\n        \"internalType\": \"string\",\n        \"name\": \"_symbol\",\n        \"type\": \"string\"\n      },\n      {\n        \"internalType\": \"uint8\",\n        \"name\": \"_decimals\",\n        \"type\": \"uint8\"\n      }\n    ],\n    \"name\": \"createSwapPair\",\n    \"outputs\": [],\n    \"stateMutability\": \"nonpayable\",\n    \"type\": \"function\"\n  },\n  {\n    \"inputs\": [\n      {\n        \"internalType\": \"bytes32\",\n        \"name\": \"swapTxHash\",\n        \"type\": \"bytes32\"\n      },\n      {\n        \"internalType\": \"address\",\n        \"name\": \"fromTokenAddr\",\n        \"type\": \"address\"\n      },\n      {\n        \"internalType\": \"address\",\n        \"name\": \"recipient\",\n        \"type\": \"address\"\n      },\n      {\n        \"internalType\": \"uint256\",\n        \"name\": \"fromChainId\",\n        \"type\": \"uint256\"\n      },\n      {\n        \"internalType\": \"uint256\",\n        \"name\": \"amount\",\n        \"type\": \"uint256\"\n      }\n    ],\n    \"name\": \"fill\",\n    \"outputs\": [],\n    \"stateMutability\": \"nonpayable\",\n    \"type\": \"function\"\n  },\n  {\n    \"inputs\": [\n      {\n        \"internalType\": \"bytes32\",\n        \"name\": \"\",\n        \"type\": \"bytes32\"\n      }\n    ],\n    \"name\": \"filledSwap\",\n    \"outputs\": [\n      {\n        \"internalType\": \"bool\",\n        \"name\": \"\",\n        \"type\": \"bool\"\n      }\n    ],\n    \"stateMutability\": \"view\",\n    \"type\": \"function\"\n  },\n  {\n    \"inputs\": [],\n    \"name\": \"initialize\",\n    \"outputs\": [],\n    \"stateMutability\": \"nonpayable\",\n    \"type\": \"function\"\n  },\n  {\n    \"inputs\": [],\n    \"name\": \"owner\",\n    \"outputs\": [\n      {\n        \"internalType\": \"address\",\n        \"name\": \"\",\n        \"type\": \"address\"\n      }\n    ],\n    \"stateMutability\": \"view\",\n    \"type\": \"function\"\n  },\n  {\n    \"inputs\": [\n      {\n        \"internalType\": \"address\",\n        \"name\": \"tokenAddr\",\n        \"type\": \"address\"\n      },\n      {\n        \"internalType\": \"uint256\",\n        \"name\": \"chainId\",\n        \"type\": \"uint256\"\n      }\n    ],\n    \"name\": \"registerSwapPair\",\n    \"outputs\": [],\n    \"stateMutability\": \"payable\",\n    \"type\": \"function\"\n  },\n  {\n    \"inputs\": [\n      {\n        \"internalType\": \"uint256\",\n        \"name\": \"\",\n        \"type\": \"uint256\"\n      },\n      {\n        \"internalType\": \"address\",\n        \"name\": \"\",\n        \"type\": \"address\"\n      }\n    ],\n    \"name\": \"registeredToken\",\n    \"outputs\": [\n      {\n        \"internalType\": \"bool\",\n        \"name\": \"\",\n        \"type\": \"bool\"\n      }\n    ],\n    \"stateMutability\": \"view\",\n    \"type\": \"function\"\n  },\n  {\n    \"inputs\": [],\n    \"name\": \"renounceOwnership\",\n    \"outputs\": [],\n    \"stateMutability\": \"nonpayable\",\n    \"type\": \"function\"\n  },\n  {\n    \"inputs\": [\n      {\n        \"internalType\": \"address\",\n        \"name\": \"tokenAddr\",\n        \"type\": \"address\"\n      },\n      {\n        \"internalType\": \"address\",\n        \"name\": \"recipient\",\n        \"type\": \"address\"\n      },\n      {\n        \"internalType\": \"uint256\",\n        \"name\": \"amount\",\n        \"type\": \"uint256\"\n      },\n      {\n        \"internalType\": \"uint256\",\n        \"name\": \"dstChainId\",\n        \"type\": \"uint256\"\n      }\n    ],\n    \"name\": \"swap\",\n    \"outputs\": [],\n    \"stateMutability\": \"payable\",\n    \"type\": \"function\"\n  },\n  {\n    \"inputs\": [\n      {\n        \"internalType\": \"uint256\",\n        \"name\": \"\",\n        \"type\": \"uint256\"\n      },\n      {\n        \"internalType\": \"address\",\n        \"name\": \"\",\n        \"type\": \"address\"\n      }\n    ],\n    \"name\": \"swapMappingIncoming\",\n    \"outputs\": [\n      {\n        \"internalType\": \"address\",\n        \"name\": \"\",\n        \"type\": \"address\"\n      }\n    ],\n    \"stateMutability\": \"view\",\n    \"type\": \"function\"\n  },\n  {\n    \"inputs\": [\n      {\n        \"internalType\": \"uint256\",\n        \"name\": \"\",\n        \"type\": \"uint256\"\n      },\n      {\n        \"internalType\": \"address\",\n        \"name\": \"\",\n        \"type\": \"address\"\n      }\n    ],\n    \"name\": \"swapMappingOutgoing\",\n    \"outputs\": [\n      {\n        \"internalType\": \"address\",\n        \"name\": \"\",\n        \"type\": \"address\"\n      }\n    ],\n    \"stateMutability\": \"view\",\n    \"type\": \"function\"\n  },\n  {\n    \"inputs\": [\n      {\n        \"internalType\": \"address\",\n        \"name\": \"newOwner\",\n        \"type\": \"address\"\n      }\n    ],\n    \"name\": \"transferOwnership\",\n    \"outputs\": [],\n    \"stateMutability\": \"nonpayable\",\n    \"type\": \"function\"\n  }\n]",
}

// ERC20SwapAgentABI is the input ABI used to generate the binding from.
// Deprecated: Use ERC20SwapAgentMetaData.ABI instead.
var ERC20SwapAgentABI = ERC20SwapAgentMetaData.ABI

// ERC20SwapAgent is an auto generated Go binding around an Ethereum contract.
type ERC20SwapAgent struct {
	ERC20SwapAgentCaller     // Read-only binding to the contract
	ERC20SwapAgentTransactor // Write-only binding to the contract
	ERC20SwapAgentFilterer   // Log filterer for contract events
}

// ERC20SwapAgentCaller is an auto generated read-only Go binding around an Ethereum contract.
type ERC20SwapAgentCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC20SwapAgentTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ERC20SwapAgentTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC20SwapAgentFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ERC20SwapAgentFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC20SwapAgentSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ERC20SwapAgentSession struct {
	Contract     *ERC20SwapAgent   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ERC20SwapAgentCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ERC20SwapAgentCallerSession struct {
	Contract *ERC20SwapAgentCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// ERC20SwapAgentTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ERC20SwapAgentTransactorSession struct {
	Contract     *ERC20SwapAgentTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// ERC20SwapAgentRaw is an auto generated low-level Go binding around an Ethereum contract.
type ERC20SwapAgentRaw struct {
	Contract *ERC20SwapAgent // Generic contract binding to access the raw methods on
}

// ERC20SwapAgentCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ERC20SwapAgentCallerRaw struct {
	Contract *ERC20SwapAgentCaller // Generic read-only contract binding to access the raw methods on
}

// ERC20SwapAgentTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ERC20SwapAgentTransactorRaw struct {
	Contract *ERC20SwapAgentTransactor // Generic write-only contract binding to access the raw methods on
}

// NewERC20SwapAgent creates a new instance of ERC20SwapAgent, bound to a specific deployed contract.
func NewERC20SwapAgent(address common.Address, backend bind.ContractBackend) (*ERC20SwapAgent, error) {
	contract, err := bindERC20SwapAgent(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ERC20SwapAgent{ERC20SwapAgentCaller: ERC20SwapAgentCaller{contract: contract}, ERC20SwapAgentTransactor: ERC20SwapAgentTransactor{contract: contract}, ERC20SwapAgentFilterer: ERC20SwapAgentFilterer{contract: contract}}, nil
}

// NewERC20SwapAgentCaller creates a new read-only instance of ERC20SwapAgent, bound to a specific deployed contract.
func NewERC20SwapAgentCaller(address common.Address, caller bind.ContractCaller) (*ERC20SwapAgentCaller, error) {
	contract, err := bindERC20SwapAgent(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ERC20SwapAgentCaller{contract: contract}, nil
}

// NewERC20SwapAgentTransactor creates a new write-only instance of ERC20SwapAgent, bound to a specific deployed contract.
func NewERC20SwapAgentTransactor(address common.Address, transactor bind.ContractTransactor) (*ERC20SwapAgentTransactor, error) {
	contract, err := bindERC20SwapAgent(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ERC20SwapAgentTransactor{contract: contract}, nil
}

// NewERC20SwapAgentFilterer creates a new log filterer instance of ERC20SwapAgent, bound to a specific deployed contract.
func NewERC20SwapAgentFilterer(address common.Address, filterer bind.ContractFilterer) (*ERC20SwapAgentFilterer, error) {
	contract, err := bindERC20SwapAgent(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ERC20SwapAgentFilterer{contract: contract}, nil
}

// bindERC20SwapAgent binds a generic wrapper to an already deployed contract.
func bindERC20SwapAgent(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ERC20SwapAgentABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ERC20SwapAgent *ERC20SwapAgentRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ERC20SwapAgent.Contract.ERC20SwapAgentCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ERC20SwapAgent *ERC20SwapAgentRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ERC20SwapAgent.Contract.ERC20SwapAgentTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ERC20SwapAgent *ERC20SwapAgentRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ERC20SwapAgent.Contract.ERC20SwapAgentTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ERC20SwapAgent *ERC20SwapAgentCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ERC20SwapAgent.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ERC20SwapAgent *ERC20SwapAgentTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ERC20SwapAgent.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ERC20SwapAgent *ERC20SwapAgentTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ERC20SwapAgent.Contract.contract.Transact(opts, method, params...)
}

// FilledSwap is a free data retrieval call binding the contract method 0xa86894ca.
//
// Solidity: function filledSwap(bytes32 ) view returns(bool)
func (_ERC20SwapAgent *ERC20SwapAgentCaller) FilledSwap(opts *bind.CallOpts, arg0 [32]byte) (bool, error) {
	var out []interface{}
	err := _ERC20SwapAgent.contract.Call(opts, &out, "filledSwap", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// FilledSwap is a free data retrieval call binding the contract method 0xa86894ca.
//
// Solidity: function filledSwap(bytes32 ) view returns(bool)
func (_ERC20SwapAgent *ERC20SwapAgentSession) FilledSwap(arg0 [32]byte) (bool, error) {
	return _ERC20SwapAgent.Contract.FilledSwap(&_ERC20SwapAgent.CallOpts, arg0)
}

// FilledSwap is a free data retrieval call binding the contract method 0xa86894ca.
//
// Solidity: function filledSwap(bytes32 ) view returns(bool)
func (_ERC20SwapAgent *ERC20SwapAgentCallerSession) FilledSwap(arg0 [32]byte) (bool, error) {
	return _ERC20SwapAgent.Contract.FilledSwap(&_ERC20SwapAgent.CallOpts, arg0)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ERC20SwapAgent *ERC20SwapAgentCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ERC20SwapAgent.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ERC20SwapAgent *ERC20SwapAgentSession) Owner() (common.Address, error) {
	return _ERC20SwapAgent.Contract.Owner(&_ERC20SwapAgent.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ERC20SwapAgent *ERC20SwapAgentCallerSession) Owner() (common.Address, error) {
	return _ERC20SwapAgent.Contract.Owner(&_ERC20SwapAgent.CallOpts)
}

// RegisteredToken is a free data retrieval call binding the contract method 0x0b4f43c1.
//
// Solidity: function registeredToken(uint256 , address ) view returns(bool)
func (_ERC20SwapAgent *ERC20SwapAgentCaller) RegisteredToken(opts *bind.CallOpts, arg0 *big.Int, arg1 common.Address) (bool, error) {
	var out []interface{}
	err := _ERC20SwapAgent.contract.Call(opts, &out, "registeredToken", arg0, arg1)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// RegisteredToken is a free data retrieval call binding the contract method 0x0b4f43c1.
//
// Solidity: function registeredToken(uint256 , address ) view returns(bool)
func (_ERC20SwapAgent *ERC20SwapAgentSession) RegisteredToken(arg0 *big.Int, arg1 common.Address) (bool, error) {
	return _ERC20SwapAgent.Contract.RegisteredToken(&_ERC20SwapAgent.CallOpts, arg0, arg1)
}

// RegisteredToken is a free data retrieval call binding the contract method 0x0b4f43c1.
//
// Solidity: function registeredToken(uint256 , address ) view returns(bool)
func (_ERC20SwapAgent *ERC20SwapAgentCallerSession) RegisteredToken(arg0 *big.Int, arg1 common.Address) (bool, error) {
	return _ERC20SwapAgent.Contract.RegisteredToken(&_ERC20SwapAgent.CallOpts, arg0, arg1)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_ERC20SwapAgent *ERC20SwapAgentCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _ERC20SwapAgent.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_ERC20SwapAgent *ERC20SwapAgentSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _ERC20SwapAgent.Contract.SupportsInterface(&_ERC20SwapAgent.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_ERC20SwapAgent *ERC20SwapAgentCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _ERC20SwapAgent.Contract.SupportsInterface(&_ERC20SwapAgent.CallOpts, interfaceId)
}

// SwapMappingIncoming is a free data retrieval call binding the contract method 0xec686704.
//
// Solidity: function swapMappingIncoming(uint256 , address ) view returns(address)
func (_ERC20SwapAgent *ERC20SwapAgentCaller) SwapMappingIncoming(opts *bind.CallOpts, arg0 *big.Int, arg1 common.Address) (common.Address, error) {
	var out []interface{}
	err := _ERC20SwapAgent.contract.Call(opts, &out, "swapMappingIncoming", arg0, arg1)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// SwapMappingIncoming is a free data retrieval call binding the contract method 0xec686704.
//
// Solidity: function swapMappingIncoming(uint256 , address ) view returns(address)
func (_ERC20SwapAgent *ERC20SwapAgentSession) SwapMappingIncoming(arg0 *big.Int, arg1 common.Address) (common.Address, error) {
	return _ERC20SwapAgent.Contract.SwapMappingIncoming(&_ERC20SwapAgent.CallOpts, arg0, arg1)
}

// SwapMappingIncoming is a free data retrieval call binding the contract method 0xec686704.
//
// Solidity: function swapMappingIncoming(uint256 , address ) view returns(address)
func (_ERC20SwapAgent *ERC20SwapAgentCallerSession) SwapMappingIncoming(arg0 *big.Int, arg1 common.Address) (common.Address, error) {
	return _ERC20SwapAgent.Contract.SwapMappingIncoming(&_ERC20SwapAgent.CallOpts, arg0, arg1)
}

// SwapMappingOutgoing is a free data retrieval call binding the contract method 0x0d43d992.
//
// Solidity: function swapMappingOutgoing(uint256 , address ) view returns(address)
func (_ERC20SwapAgent *ERC20SwapAgentCaller) SwapMappingOutgoing(opts *bind.CallOpts, arg0 *big.Int, arg1 common.Address) (common.Address, error) {
	var out []interface{}
	err := _ERC20SwapAgent.contract.Call(opts, &out, "swapMappingOutgoing", arg0, arg1)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// SwapMappingOutgoing is a free data retrieval call binding the contract method 0x0d43d992.
//
// Solidity: function swapMappingOutgoing(uint256 , address ) view returns(address)
func (_ERC20SwapAgent *ERC20SwapAgentSession) SwapMappingOutgoing(arg0 *big.Int, arg1 common.Address) (common.Address, error) {
	return _ERC20SwapAgent.Contract.SwapMappingOutgoing(&_ERC20SwapAgent.CallOpts, arg0, arg1)
}

// SwapMappingOutgoing is a free data retrieval call binding the contract method 0x0d43d992.
//
// Solidity: function swapMappingOutgoing(uint256 , address ) view returns(address)
func (_ERC20SwapAgent *ERC20SwapAgentCallerSession) SwapMappingOutgoing(arg0 *big.Int, arg1 common.Address) (common.Address, error) {
	return _ERC20SwapAgent.Contract.SwapMappingOutgoing(&_ERC20SwapAgent.CallOpts, arg0, arg1)
}

// CreateSwapPair is a paid mutator transaction binding the contract method 0x4acbe1ca.
//
// Solidity: function createSwapPair(bytes32 registerTxHash, address fromTokenAddr, uint256 fromChainId, string uri) returns()
func (_ERC20SwapAgent *ERC20SwapAgentTransactor) CreateSwapPair(opts *bind.TransactOpts, registerTxHash [32]byte, fromTokenAddr common.Address, fromChainId *big.Int, name string, symbol string, decimals uint8) (*types.Transaction, error) {
	return _ERC20SwapAgent.contract.Transact(opts, "createSwapPair", registerTxHash, fromTokenAddr, fromChainId, name, symbol, decimals)
}

// CreateSwapPair is a paid mutator transaction binding the contract method 0x4acbe1ca.
//
// Solidity: function createSwapPair(bytes32 registerTxHash, address fromTokenAddr, uint256 fromChainId, string uri) returns()
func (_ERC20SwapAgent *ERC20SwapAgentSession) CreateSwapPair(registerTxHash [32]byte, fromTokenAddr common.Address, fromChainId *big.Int, name string, symbol string, decimals uint8) (*types.Transaction, error) {
	return _ERC20SwapAgent.Contract.CreateSwapPair(&_ERC20SwapAgent.TransactOpts, registerTxHash, fromTokenAddr, fromChainId, name, symbol, decimals)
}

// CreateSwapPair is a paid mutator transaction binding the contract method 0x4acbe1ca.
//
// Solidity: function createSwapPair(bytes32 registerTxHash, address fromTokenAddr, uint256 fromChainId, string uri) returns()
func (_ERC20SwapAgent *ERC20SwapAgentTransactorSession) CreateSwapPair(registerTxHash [32]byte, fromTokenAddr common.Address, fromChainId *big.Int, name string, symbol string, decimals uint8) (*types.Transaction, error) {
	return _ERC20SwapAgent.Contract.CreateSwapPair(&_ERC20SwapAgent.TransactOpts, registerTxHash, fromTokenAddr, fromChainId, name, symbol, decimals)
}

// Fill is a paid mutator transaction binding the contract method 0x04828122.
//
// Solidity: function fill(bytes32 swapTxHash, address fromTokenAddr, address recipient, uint256 fromChainId, uint256 amount) returns()
func (_ERC20SwapAgent *ERC20SwapAgentTransactor) Fill(opts *bind.TransactOpts, swapTxHash [32]byte, fromTokenAddr common.Address, recipient common.Address, fromChainId *big.Int, amount *big.Int) (*types.Transaction, error) {
	return _ERC20SwapAgent.contract.Transact(opts, "fill", swapTxHash, fromTokenAddr, recipient, fromChainId, amount)
}

// Fill is a paid mutator transaction binding the contract method 0x04828122.
//
// Solidity: function fill(bytes32 swapTxHash, address fromTokenAddr, address recipient, uint256 fromChainId, uint256 amount) returns()
func (_ERC20SwapAgent *ERC20SwapAgentSession) Fill(swapTxHash [32]byte, fromTokenAddr common.Address, recipient common.Address, fromChainId *big.Int, amount *big.Int) (*types.Transaction, error) {
	return _ERC20SwapAgent.Contract.Fill(&_ERC20SwapAgent.TransactOpts, swapTxHash, fromTokenAddr, recipient, fromChainId, amount)
}

// Fill is a paid mutator transaction binding the contract method 0x04828122.
//
// Solidity: function fill(bytes32 swapTxHash, address fromTokenAddr, address recipient, uint256 fromChainId, uint256 amount) returns()
func (_ERC20SwapAgent *ERC20SwapAgentTransactorSession) Fill(swapTxHash [32]byte, fromTokenAddr common.Address, recipient common.Address, fromChainId *big.Int, amount *big.Int) (*types.Transaction, error) {
	return _ERC20SwapAgent.Contract.Fill(&_ERC20SwapAgent.TransactOpts, swapTxHash, fromTokenAddr, recipient, fromChainId, amount)
}

// Initialize is a paid mutator transaction binding the contract method 0x8129fc1c.
//
// Solidity: function initialize() returns()
func (_ERC20SwapAgent *ERC20SwapAgentTransactor) Initialize(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ERC20SwapAgent.contract.Transact(opts, "initialize")
}

// Initialize is a paid mutator transaction binding the contract method 0x8129fc1c.
//
// Solidity: function initialize() returns()
func (_ERC20SwapAgent *ERC20SwapAgentSession) Initialize() (*types.Transaction, error) {
	return _ERC20SwapAgent.Contract.Initialize(&_ERC20SwapAgent.TransactOpts)
}

// Initialize is a paid mutator transaction binding the contract method 0x8129fc1c.
//
// Solidity: function initialize() returns()
func (_ERC20SwapAgent *ERC20SwapAgentTransactorSession) Initialize() (*types.Transaction, error) {
	return _ERC20SwapAgent.Contract.Initialize(&_ERC20SwapAgent.TransactOpts)
}

// OnERC20BatchReceived is a paid mutator transaction binding the contract method 0xbc197c81.
//
// Solidity: function onERC20BatchReceived(address , address , uint256[] , uint256[] , bytes ) returns(bytes4)
func (_ERC20SwapAgent *ERC20SwapAgentTransactor) OnERC20BatchReceived(opts *bind.TransactOpts, arg0 common.Address, arg1 common.Address, arg2 []*big.Int, arg3 []*big.Int, arg4 []byte) (*types.Transaction, error) {
	return _ERC20SwapAgent.contract.Transact(opts, "onERC20BatchReceived", arg0, arg1, arg2, arg3, arg4)
}

// OnERC20BatchReceived is a paid mutator transaction binding the contract method 0xbc197c81.
//
// Solidity: function onERC20BatchReceived(address , address , uint256[] , uint256[] , bytes ) returns(bytes4)
func (_ERC20SwapAgent *ERC20SwapAgentSession) OnERC20BatchReceived(arg0 common.Address, arg1 common.Address, arg2 []*big.Int, arg3 []*big.Int, arg4 []byte) (*types.Transaction, error) {
	return _ERC20SwapAgent.Contract.OnERC20BatchReceived(&_ERC20SwapAgent.TransactOpts, arg0, arg1, arg2, arg3, arg4)
}

// OnERC20BatchReceived is a paid mutator transaction binding the contract method 0xbc197c81.
//
// Solidity: function onERC20BatchReceived(address , address , uint256[] , uint256[] , bytes ) returns(bytes4)
func (_ERC20SwapAgent *ERC20SwapAgentTransactorSession) OnERC20BatchReceived(arg0 common.Address, arg1 common.Address, arg2 []*big.Int, arg3 []*big.Int, arg4 []byte) (*types.Transaction, error) {
	return _ERC20SwapAgent.Contract.OnERC20BatchReceived(&_ERC20SwapAgent.TransactOpts, arg0, arg1, arg2, arg3, arg4)
}

// OnERC20Received is a paid mutator transaction binding the contract method 0xf23a6e61.
//
// Solidity: function onERC20Received(address , address , uint256 , uint256 , bytes ) returns(bytes4)
func (_ERC20SwapAgent *ERC20SwapAgentTransactor) OnERC20Received(opts *bind.TransactOpts, arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 *big.Int, arg4 []byte) (*types.Transaction, error) {
	return _ERC20SwapAgent.contract.Transact(opts, "onERC20Received", arg0, arg1, arg2, arg3, arg4)
}

// OnERC20Received is a paid mutator transaction binding the contract method 0xf23a6e61.
//
// Solidity: function onERC20Received(address , address , uint256 , uint256 , bytes ) returns(bytes4)
func (_ERC20SwapAgent *ERC20SwapAgentSession) OnERC20Received(arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 *big.Int, arg4 []byte) (*types.Transaction, error) {
	return _ERC20SwapAgent.Contract.OnERC20Received(&_ERC20SwapAgent.TransactOpts, arg0, arg1, arg2, arg3, arg4)
}

// OnERC20Received is a paid mutator transaction binding the contract method 0xf23a6e61.
//
// Solidity: function onERC20Received(address , address , uint256 , uint256 , bytes ) returns(bytes4)
func (_ERC20SwapAgent *ERC20SwapAgentTransactorSession) OnERC20Received(arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 *big.Int, arg4 []byte) (*types.Transaction, error) {
	return _ERC20SwapAgent.Contract.OnERC20Received(&_ERC20SwapAgent.TransactOpts, arg0, arg1, arg2, arg3, arg4)
}

// RegisterSwapPair is a paid mutator transaction binding the contract method 0x45b1ab1b.
//
// Solidity: function registerSwapPair(address tokenAddr, uint256 chainId) payable returns()
func (_ERC20SwapAgent *ERC20SwapAgentTransactor) RegisterSwapPair(opts *bind.TransactOpts, tokenAddr common.Address, chainId *big.Int) (*types.Transaction, error) {
	return _ERC20SwapAgent.contract.Transact(opts, "registerSwapPair", tokenAddr, chainId)
}

// RegisterSwapPair is a paid mutator transaction binding the contract method 0x45b1ab1b.
//
// Solidity: function registerSwapPair(address tokenAddr, uint256 chainId) payable returns()
func (_ERC20SwapAgent *ERC20SwapAgentSession) RegisterSwapPair(tokenAddr common.Address, chainId *big.Int) (*types.Transaction, error) {
	return _ERC20SwapAgent.Contract.RegisterSwapPair(&_ERC20SwapAgent.TransactOpts, tokenAddr, chainId)
}

// RegisterSwapPair is a paid mutator transaction binding the contract method 0x45b1ab1b.
//
// Solidity: function registerSwapPair(address tokenAddr, uint256 chainId) payable returns()
func (_ERC20SwapAgent *ERC20SwapAgentTransactorSession) RegisterSwapPair(tokenAddr common.Address, chainId *big.Int) (*types.Transaction, error) {
	return _ERC20SwapAgent.Contract.RegisterSwapPair(&_ERC20SwapAgent.TransactOpts, tokenAddr, chainId)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_ERC20SwapAgent *ERC20SwapAgentTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ERC20SwapAgent.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_ERC20SwapAgent *ERC20SwapAgentSession) RenounceOwnership() (*types.Transaction, error) {
	return _ERC20SwapAgent.Contract.RenounceOwnership(&_ERC20SwapAgent.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_ERC20SwapAgent *ERC20SwapAgentTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _ERC20SwapAgent.Contract.RenounceOwnership(&_ERC20SwapAgent.TransactOpts)
}

// Swap is a paid mutator transaction binding the contract method 0xa180639a.
//
// Solidity: function swap(address tokenAddr, address recipient, uint256 amount, uint256 dstChainId) payable returns()
func (_ERC20SwapAgent *ERC20SwapAgentTransactor) Swap(opts *bind.TransactOpts, tokenAddr common.Address, recipient common.Address, amount *big.Int, dstChainId *big.Int) (*types.Transaction, error) {
	return _ERC20SwapAgent.contract.Transact(opts, "swap", tokenAddr, recipient, amount, dstChainId)
}

// Swap is a paid mutator transaction binding the contract method 0xa180639a.
//
// Solidity: function swap(address tokenAddr, address recipient, uint256 amount, uint256 dstChainId) payable returns()
func (_ERC20SwapAgent *ERC20SwapAgentSession) Swap(tokenAddr common.Address, recipient common.Address, amount *big.Int, dstChainId *big.Int) (*types.Transaction, error) {
	return _ERC20SwapAgent.Contract.Swap(&_ERC20SwapAgent.TransactOpts, tokenAddr, recipient, amount, dstChainId)
}

// Swap is a paid mutator transaction binding the contract method 0xa180639a.
//
// Solidity: function swap(address tokenAddr, address recipient, uint256 amount, uint256 dstChainId) payable returns()
func (_ERC20SwapAgent *ERC20SwapAgentTransactorSession) Swap(tokenAddr common.Address, recipient common.Address, amount *big.Int, dstChainId *big.Int) (*types.Transaction, error) {
	return _ERC20SwapAgent.Contract.Swap(&_ERC20SwapAgent.TransactOpts, tokenAddr, recipient, amount, dstChainId)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_ERC20SwapAgent *ERC20SwapAgentTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _ERC20SwapAgent.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_ERC20SwapAgent *ERC20SwapAgentSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _ERC20SwapAgent.Contract.TransferOwnership(&_ERC20SwapAgent.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_ERC20SwapAgent *ERC20SwapAgentTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _ERC20SwapAgent.Contract.TransferOwnership(&_ERC20SwapAgent.TransactOpts, newOwner)
}

// ERC20SwapAgentBackwardSwapFilledIterator is returned from FilterBackwardSwapFilled and is used to iterate over the raw logs and unpacked data for BackwardSwapFilled events raised by the ERC20SwapAgent contract.
type ERC20SwapAgentBackwardSwapFilledIterator struct {
	Event *ERC20SwapAgentBackwardSwapFilled // Event containing the contract specifics and raw log

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
func (it *ERC20SwapAgentBackwardSwapFilledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC20SwapAgentBackwardSwapFilled)
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
		it.Event = new(ERC20SwapAgentBackwardSwapFilled)
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
func (it *ERC20SwapAgentBackwardSwapFilledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC20SwapAgentBackwardSwapFilledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC20SwapAgentBackwardSwapFilled represents a BackwardSwapFilled event raised by the ERC20SwapAgent contract.
type ERC20SwapAgentBackwardSwapFilled struct {
	SwapTxHash  [32]byte
	TokenAddr   common.Address
	Recipient   common.Address
	FromChainId *big.Int
	Amount      *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterBackwardSwapFilled is a free log retrieval operation binding the contract event 0x58d9c075708eb3187538135716481600d99ff6eb833bdbac1fee1e54cebec1b3.
//
// Solidity: event BackwardSwapFilled(bytes32 indexed swapTxHash, address indexed tokenAddr, address indexed recipient, uint256 fromChainId, uint256 amount)
func (_ERC20SwapAgent *ERC20SwapAgentFilterer) FilterBackwardSwapFilled(opts *bind.FilterOpts, swapTxHash [][32]byte, tokenAddr []common.Address, recipient []common.Address) (*ERC20SwapAgentBackwardSwapFilledIterator, error) {

	var swapTxHashRule []interface{}
	for _, swapTxHashItem := range swapTxHash {
		swapTxHashRule = append(swapTxHashRule, swapTxHashItem)
	}
	var tokenAddrRule []interface{}
	for _, tokenAddrItem := range tokenAddr {
		tokenAddrRule = append(tokenAddrRule, tokenAddrItem)
	}
	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	logs, sub, err := _ERC20SwapAgent.contract.FilterLogs(opts, "BackwardSwapFilled", swapTxHashRule, tokenAddrRule, recipientRule)
	if err != nil {
		return nil, err
	}
	return &ERC20SwapAgentBackwardSwapFilledIterator{contract: _ERC20SwapAgent.contract, event: "BackwardSwapFilled", logs: logs, sub: sub}, nil
}

// WatchBackwardSwapFilled is a free log subscription operation binding the contract event 0x58d9c075708eb3187538135716481600d99ff6eb833bdbac1fee1e54cebec1b3.
//
// Solidity: event BackwardSwapFilled(bytes32 indexed swapTxHash, address indexed tokenAddr, address indexed recipient, uint256 fromChainId, uint256 amount)
func (_ERC20SwapAgent *ERC20SwapAgentFilterer) WatchBackwardSwapFilled(opts *bind.WatchOpts, sink chan<- *ERC20SwapAgentBackwardSwapFilled, swapTxHash [][32]byte, tokenAddr []common.Address, recipient []common.Address) (event.Subscription, error) {

	var swapTxHashRule []interface{}
	for _, swapTxHashItem := range swapTxHash {
		swapTxHashRule = append(swapTxHashRule, swapTxHashItem)
	}
	var tokenAddrRule []interface{}
	for _, tokenAddrItem := range tokenAddr {
		tokenAddrRule = append(tokenAddrRule, tokenAddrItem)
	}
	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	logs, sub, err := _ERC20SwapAgent.contract.WatchLogs(opts, "BackwardSwapFilled", swapTxHashRule, tokenAddrRule, recipientRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC20SwapAgentBackwardSwapFilled)
				if err := _ERC20SwapAgent.contract.UnpackLog(event, "BackwardSwapFilled", log); err != nil {
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

// ParseBackwardSwapFilled is a log parse operation binding the contract event 0x58d9c075708eb3187538135716481600d99ff6eb833bdbac1fee1e54cebec1b3.
//
// Solidity: event BackwardSwapFilled(bytes32 indexed swapTxHash, address indexed tokenAddr, address indexed recipient, uint256 fromChainId, uint256 amount)
func (_ERC20SwapAgent *ERC20SwapAgentFilterer) ParseBackwardSwapFilled(log types.Log) (*ERC20SwapAgentBackwardSwapFilled, error) {
	event := new(ERC20SwapAgentBackwardSwapFilled)
	if err := _ERC20SwapAgent.contract.UnpackLog(event, "BackwardSwapFilled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ERC20SwapAgentBackwardSwapStartedIterator is returned from FilterBackwardSwapStarted and is used to iterate over the raw logs and unpacked data for BackwardSwapStarted events raised by the ERC20SwapAgent contract.
type ERC20SwapAgentBackwardSwapStartedIterator struct {
	Event *ERC20SwapAgentBackwardSwapStarted // Event containing the contract specifics and raw log

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
func (it *ERC20SwapAgentBackwardSwapStartedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC20SwapAgentBackwardSwapStarted)
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
		it.Event = new(ERC20SwapAgentBackwardSwapStarted)
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
func (it *ERC20SwapAgentBackwardSwapStartedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC20SwapAgentBackwardSwapStartedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC20SwapAgentBackwardSwapStarted represents a BackwardSwapStarted event raised by the ERC20SwapAgent contract.
type ERC20SwapAgentBackwardSwapStarted struct {
	MirroredTokenAddr common.Address
	Sender            common.Address
	Recipient         common.Address
	DstChainId        *big.Int
	Amount            *big.Int
	FeeAmount         *big.Int
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterBackwardSwapStarted is a free log retrieval operation binding the contract event 0x5c317e3669ab4c20e7362c3bdc16700c64c83aa52a3abdd32e17eb1179e6706e.
//
// Solidity: event BackwardSwapStarted(address indexed mirroredTokenAddr, address indexed sender, address indexed recipient, uint256 dstChainId, uint256 amount, uint256 feeAmount)
func (_ERC20SwapAgent *ERC20SwapAgentFilterer) FilterBackwardSwapStarted(opts *bind.FilterOpts, mirroredTokenAddr []common.Address, sender []common.Address, recipient []common.Address) (*ERC20SwapAgentBackwardSwapStartedIterator, error) {

	var mirroredTokenAddrRule []interface{}
	for _, mirroredTokenAddrItem := range mirroredTokenAddr {
		mirroredTokenAddrRule = append(mirroredTokenAddrRule, mirroredTokenAddrItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	logs, sub, err := _ERC20SwapAgent.contract.FilterLogs(opts, "BackwardSwapStarted", mirroredTokenAddrRule, senderRule, recipientRule)
	if err != nil {
		return nil, err
	}
	return &ERC20SwapAgentBackwardSwapStartedIterator{contract: _ERC20SwapAgent.contract, event: "BackwardSwapStarted", logs: logs, sub: sub}, nil
}

// WatchBackwardSwapStarted is a free log subscription operation binding the contract event 0x5c317e3669ab4c20e7362c3bdc16700c64c83aa52a3abdd32e17eb1179e6706e.
//
// Solidity: event BackwardSwapStarted(address indexed mirroredTokenAddr, address indexed sender, address indexed recipient, uint256 dstChainId, uint256 amount, uint256 feeAmount)
func (_ERC20SwapAgent *ERC20SwapAgentFilterer) WatchBackwardSwapStarted(opts *bind.WatchOpts, sink chan<- *ERC20SwapAgentBackwardSwapStarted, mirroredTokenAddr []common.Address, sender []common.Address, recipient []common.Address) (event.Subscription, error) {

	var mirroredTokenAddrRule []interface{}
	for _, mirroredTokenAddrItem := range mirroredTokenAddr {
		mirroredTokenAddrRule = append(mirroredTokenAddrRule, mirroredTokenAddrItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	logs, sub, err := _ERC20SwapAgent.contract.WatchLogs(opts, "BackwardSwapStarted", mirroredTokenAddrRule, senderRule, recipientRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC20SwapAgentBackwardSwapStarted)
				if err := _ERC20SwapAgent.contract.UnpackLog(event, "BackwardSwapStarted", log); err != nil {
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

// ParseBackwardSwapStarted is a log parse operation binding the contract event 0x5c317e3669ab4c20e7362c3bdc16700c64c83aa52a3abdd32e17eb1179e6706e.
//
// Solidity: event BackwardSwapStarted(address indexed mirroredTokenAddr, address indexed sender, address indexed recipient, uint256 dstChainId, uint256 amount, uint256 feeAmount)
func (_ERC20SwapAgent *ERC20SwapAgentFilterer) ParseBackwardSwapStarted(log types.Log) (*ERC20SwapAgentBackwardSwapStarted, error) {
	event := new(ERC20SwapAgentBackwardSwapStarted)
	if err := _ERC20SwapAgent.contract.UnpackLog(event, "BackwardSwapStarted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ERC20SwapAgentOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the ERC20SwapAgent contract.
type ERC20SwapAgentOwnershipTransferredIterator struct {
	Event *ERC20SwapAgentOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *ERC20SwapAgentOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC20SwapAgentOwnershipTransferred)
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
		it.Event = new(ERC20SwapAgentOwnershipTransferred)
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
func (it *ERC20SwapAgentOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC20SwapAgentOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC20SwapAgentOwnershipTransferred represents a OwnershipTransferred event raised by the ERC20SwapAgent contract.
type ERC20SwapAgentOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_ERC20SwapAgent *ERC20SwapAgentFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*ERC20SwapAgentOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _ERC20SwapAgent.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &ERC20SwapAgentOwnershipTransferredIterator{contract: _ERC20SwapAgent.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_ERC20SwapAgent *ERC20SwapAgentFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *ERC20SwapAgentOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _ERC20SwapAgent.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC20SwapAgentOwnershipTransferred)
				if err := _ERC20SwapAgent.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_ERC20SwapAgent *ERC20SwapAgentFilterer) ParseOwnershipTransferred(log types.Log) (*ERC20SwapAgentOwnershipTransferred, error) {
	event := new(ERC20SwapAgentOwnershipTransferred)
	if err := _ERC20SwapAgent.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ERC20SwapAgentSwapFilledIterator is returned from FilterSwapFilled and is used to iterate over the raw logs and unpacked data for SwapFilled events raised by the ERC20SwapAgent contract.
type ERC20SwapAgentSwapFilledIterator struct {
	Event *ERC20SwapAgentSwapFilled // Event containing the contract specifics and raw log

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
func (it *ERC20SwapAgentSwapFilledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC20SwapAgentSwapFilled)
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
		it.Event = new(ERC20SwapAgentSwapFilled)
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
func (it *ERC20SwapAgentSwapFilledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC20SwapAgentSwapFilledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC20SwapAgentSwapFilled represents a SwapFilled event raised by the ERC20SwapAgent contract.
type ERC20SwapAgentSwapFilled struct {
	SwapTxHash        [32]byte
	FromTokenAddr     common.Address
	Recipient         common.Address
	MirroredTokenAddr common.Address
	FromChainId       *big.Int
	Amount            *big.Int
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterSwapFilled is a free log retrieval operation binding the contract event 0x295d1e2c3b0c279f7107336cf70913e43ee3b0e77dee0b1d04f74d733006d806.
//
// Solidity: event SwapFilled(bytes32 indexed swapTxHash, address indexed fromTokenAddr, address indexed recipient, address mirroredTokenAddr, uint256 fromChainId, uint256 amount)
func (_ERC20SwapAgent *ERC20SwapAgentFilterer) FilterSwapFilled(opts *bind.FilterOpts, swapTxHash [][32]byte, fromTokenAddr []common.Address, recipient []common.Address) (*ERC20SwapAgentSwapFilledIterator, error) {

	var swapTxHashRule []interface{}
	for _, swapTxHashItem := range swapTxHash {
		swapTxHashRule = append(swapTxHashRule, swapTxHashItem)
	}
	var fromTokenAddrRule []interface{}
	for _, fromTokenAddrItem := range fromTokenAddr {
		fromTokenAddrRule = append(fromTokenAddrRule, fromTokenAddrItem)
	}
	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	logs, sub, err := _ERC20SwapAgent.contract.FilterLogs(opts, "SwapFilled", swapTxHashRule, fromTokenAddrRule, recipientRule)

	if err != nil {
		return nil, err
	}
	return &ERC20SwapAgentSwapFilledIterator{contract: _ERC20SwapAgent.contract, event: "SwapFilled", logs: logs, sub: sub}, nil
}

// WatchSwapFilled is a free log subscription operation binding the contract event 0x295d1e2c3b0c279f7107336cf70913e43ee3b0e77dee0b1d04f74d733006d806.
//
// Solidity: event SwapFilled(bytes32 indexed swapTxHash, address indexed fromTokenAddr, address indexed recipient, address mirroredTokenAddr, uint256 fromChainId, uint256 amount)
func (_ERC20SwapAgent *ERC20SwapAgentFilterer) WatchSwapFilled(opts *bind.WatchOpts, sink chan<- *ERC20SwapAgentSwapFilled, swapTxHash [][32]byte, fromTokenAddr []common.Address, recipient []common.Address) (event.Subscription, error) {

	var swapTxHashRule []interface{}
	for _, swapTxHashItem := range swapTxHash {
		swapTxHashRule = append(swapTxHashRule, swapTxHashItem)
	}
	var fromTokenAddrRule []interface{}
	for _, fromTokenAddrItem := range fromTokenAddr {
		fromTokenAddrRule = append(fromTokenAddrRule, fromTokenAddrItem)
	}
	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	logs, sub, err := _ERC20SwapAgent.contract.WatchLogs(opts, "SwapFilled", swapTxHashRule, fromTokenAddrRule, recipientRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC20SwapAgentSwapFilled)
				if err := _ERC20SwapAgent.contract.UnpackLog(event, "SwapFilled", log); err != nil {
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

// ParseSwapFilled is a log parse operation binding the contract event 0x295d1e2c3b0c279f7107336cf70913e43ee3b0e77dee0b1d04f74d733006d806.
//
// Solidity: event SwapFilled(bytes32 indexed swapTxHash, address indexed fromTokenAddr, address indexed recipient, address mirroredTokenAddr, uint256 fromChainId, uint256 amount)
func (_ERC20SwapAgent *ERC20SwapAgentFilterer) ParseSwapFilled(log types.Log) (*ERC20SwapAgentSwapFilled, error) {
	event := new(ERC20SwapAgentSwapFilled)
	if err := _ERC20SwapAgent.contract.UnpackLog(event, "SwapFilled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ERC20SwapAgentSwapPairCreatedIterator is returned from FilterSwapPairCreated and is used to iterate over the raw logs and unpacked data for SwapPairCreated events raised by the ERC20SwapAgent contract.
type ERC20SwapAgentSwapPairCreatedIterator struct {
	Event *ERC20SwapAgentSwapPairCreated // Event containing the contract specifics and raw log

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
func (it *ERC20SwapAgentSwapPairCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC20SwapAgentSwapPairCreated)
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
		it.Event = new(ERC20SwapAgentSwapPairCreated)
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
func (it *ERC20SwapAgentSwapPairCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC20SwapAgentSwapPairCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC20SwapAgentSwapPairCreated represents a SwapPairCreated event raised by the ERC20SwapAgent contract.
type ERC20SwapAgentSwapPairCreated struct {
	RegisterTxHash    [32]byte
	FromTokenAddr     common.Address
	MirroredTokenAddr common.Address
	FromChainId       *big.Int
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterSwapPairCreated is a free log retrieval operation binding the contract event 0x621d9726b59bf7f3a9cbd292df8310172e6dec3a7279c906c7c22129f406708e.
//
// Solidity: event SwapPairCreated(bytes32 indexed registerTxHash, address indexed fromTokenAddr, address indexed mirroredTokenAddr, uint256 fromChainId)
func (_ERC20SwapAgent *ERC20SwapAgentFilterer) FilterSwapPairCreated(opts *bind.FilterOpts, registerTxHash [][32]byte, fromTokenAddr []common.Address, mirroredTokenAddr []common.Address) (*ERC20SwapAgentSwapPairCreatedIterator, error) {

	var registerTxHashRule []interface{}
	for _, registerTxHashItem := range registerTxHash {
		registerTxHashRule = append(registerTxHashRule, registerTxHashItem)
	}
	var fromTokenAddrRule []interface{}
	for _, fromTokenAddrItem := range fromTokenAddr {
		fromTokenAddrRule = append(fromTokenAddrRule, fromTokenAddrItem)
	}
	var mirroredTokenAddrRule []interface{}
	for _, mirroredTokenAddrItem := range mirroredTokenAddr {
		mirroredTokenAddrRule = append(mirroredTokenAddrRule, mirroredTokenAddrItem)
	}

	logs, sub, err := _ERC20SwapAgent.contract.FilterLogs(opts, "SwapPairCreated", registerTxHashRule, fromTokenAddrRule, mirroredTokenAddrRule)
	if err != nil {
		return nil, err
	}
	return &ERC20SwapAgentSwapPairCreatedIterator{contract: _ERC20SwapAgent.contract, event: "SwapPairCreated", logs: logs, sub: sub}, nil
}

// WatchSwapPairCreated is a free log subscription operation binding the contract event 0x621d9726b59bf7f3a9cbd292df8310172e6dec3a7279c906c7c22129f406708e.
//
// Solidity: event SwapPairCreated(bytes32 indexed registerTxHash, address indexed fromTokenAddr, address indexed mirroredTokenAddr, uint256 fromChainId)
func (_ERC20SwapAgent *ERC20SwapAgentFilterer) WatchSwapPairCreated(opts *bind.WatchOpts, sink chan<- *ERC20SwapAgentSwapPairCreated, registerTxHash [][32]byte, fromTokenAddr []common.Address, mirroredTokenAddr []common.Address) (event.Subscription, error) {

	var registerTxHashRule []interface{}
	for _, registerTxHashItem := range registerTxHash {
		registerTxHashRule = append(registerTxHashRule, registerTxHashItem)
	}
	var fromTokenAddrRule []interface{}
	for _, fromTokenAddrItem := range fromTokenAddr {
		fromTokenAddrRule = append(fromTokenAddrRule, fromTokenAddrItem)
	}
	var mirroredTokenAddrRule []interface{}
	for _, mirroredTokenAddrItem := range mirroredTokenAddr {
		mirroredTokenAddrRule = append(mirroredTokenAddrRule, mirroredTokenAddrItem)
	}

	logs, sub, err := _ERC20SwapAgent.contract.WatchLogs(opts, "SwapPairCreated", registerTxHashRule, fromTokenAddrRule, mirroredTokenAddrRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC20SwapAgentSwapPairCreated)
				if err := _ERC20SwapAgent.contract.UnpackLog(event, "SwapPairCreated", log); err != nil {
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

// ParseSwapPairCreated is a log parse operation binding the contract event 0x621d9726b59bf7f3a9cbd292df8310172e6dec3a7279c906c7c22129f406708e.
//
// Solidity: event SwapPairCreated(bytes32 indexed registerTxHash, address indexed fromTokenAddr, address indexed mirroredTokenAddr, uint256 fromChainId)
func (_ERC20SwapAgent *ERC20SwapAgentFilterer) ParseSwapPairCreated(log types.Log) (*ERC20SwapAgentSwapPairCreated, error) {
	event := new(ERC20SwapAgentSwapPairCreated)
	if err := _ERC20SwapAgent.contract.UnpackLog(event, "SwapPairCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ERC20SwapAgentSwapPairRegisterIterator is returned from FilterSwapPairRegister and is used to iterate over the raw logs and unpacked data for SwapPairRegister events raised by the ERC20SwapAgent contract.
type ERC20SwapAgentSwapPairRegisterIterator struct {
	Event *ERC20SwapAgentSwapPairRegister // Event containing the contract specifics and raw log

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
func (it *ERC20SwapAgentSwapPairRegisterIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC20SwapAgentSwapPairRegister)
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
		it.Event = new(ERC20SwapAgentSwapPairRegister)
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
func (it *ERC20SwapAgentSwapPairRegisterIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC20SwapAgentSwapPairRegisterIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC20SwapAgentSwapPairRegister represents a SwapPairRegister event raised by the ERC20SwapAgent contract.
type ERC20SwapAgentSwapPairRegister struct {
	Sponsor      common.Address
	TokenAddress common.Address
	ToChainId    *big.Int
	FeeAmount    *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterSwapPairRegister is a free log retrieval operation binding the contract event 0x6745f5c9e689dd3b252d145964ac6fbb294d72cedae5a4d52b1774728289e606.
//
// Solidity: event SwapPairRegister(address indexed sponsor, address indexed tokenAddress, uint256 toChainId, uint256 feeAmount)
func (_ERC20SwapAgent *ERC20SwapAgentFilterer) FilterSwapPairRegister(opts *bind.FilterOpts, sponsor []common.Address, tokenAddress []common.Address) (*ERC20SwapAgentSwapPairRegisterIterator, error) {

	var sponsorRule []interface{}
	for _, sponsorItem := range sponsor {
		sponsorRule = append(sponsorRule, sponsorItem)
	}
	var tokenAddressRule []interface{}
	for _, tokenAddressItem := range tokenAddress {
		tokenAddressRule = append(tokenAddressRule, tokenAddressItem)
	}

	logs, sub, err := _ERC20SwapAgent.contract.FilterLogs(opts, "SwapPairRegister", sponsorRule, tokenAddressRule)
	if err != nil {
		return nil, err
	}
	return &ERC20SwapAgentSwapPairRegisterIterator{contract: _ERC20SwapAgent.contract, event: "SwapPairRegister", logs: logs, sub: sub}, nil
}

// WatchSwapPairRegister is a free log subscription operation binding the contract event 0x6745f5c9e689dd3b252d145964ac6fbb294d72cedae5a4d52b1774728289e606.
//
// Solidity: event SwapPairRegister(address indexed sponsor, address indexed tokenAddress, uint256 toChainId, uint256 feeAmount)
func (_ERC20SwapAgent *ERC20SwapAgentFilterer) WatchSwapPairRegister(opts *bind.WatchOpts, sink chan<- *ERC20SwapAgentSwapPairRegister, sponsor []common.Address, tokenAddress []common.Address) (event.Subscription, error) {

	var sponsorRule []interface{}
	for _, sponsorItem := range sponsor {
		sponsorRule = append(sponsorRule, sponsorItem)
	}
	var tokenAddressRule []interface{}
	for _, tokenAddressItem := range tokenAddress {
		tokenAddressRule = append(tokenAddressRule, tokenAddressItem)
	}

	logs, sub, err := _ERC20SwapAgent.contract.WatchLogs(opts, "SwapPairRegister", sponsorRule, tokenAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC20SwapAgentSwapPairRegister)
				if err := _ERC20SwapAgent.contract.UnpackLog(event, "SwapPairRegister", log); err != nil {
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

// ParseSwapPairRegister is a log parse operation binding the contract event 0x6745f5c9e689dd3b252d145964ac6fbb294d72cedae5a4d52b1774728289e606.
//
// Solidity: event SwapPairRegister(address indexed sponsor, address indexed tokenAddress, uint256 toChainId, uint256 feeAmount)
func (_ERC20SwapAgent *ERC20SwapAgentFilterer) ParseSwapPairRegister(log types.Log) (*ERC20SwapAgentSwapPairRegister, error) {
	event := new(ERC20SwapAgentSwapPairRegister)
	if err := _ERC20SwapAgent.contract.UnpackLog(event, "SwapPairRegister", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ERC20SwapAgentSwapStartedIterator is returned from FilterSwapStarted and is used to iterate over the raw logs and unpacked data for SwapStarted events raised by the ERC20SwapAgent contract.
type ERC20SwapAgentSwapStartedIterator struct {
	Event *ERC20SwapAgentSwapStarted // Event containing the contract specifics and raw log

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
func (it *ERC20SwapAgentSwapStartedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC20SwapAgentSwapStarted)
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
		it.Event = new(ERC20SwapAgentSwapStarted)
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
func (it *ERC20SwapAgentSwapStartedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC20SwapAgentSwapStartedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC20SwapAgentSwapStarted represents a SwapStarted event raised by the ERC20SwapAgent contract.
type ERC20SwapAgentSwapStarted struct {
	TokenAddr  common.Address
	Sender     common.Address
	Recipient  common.Address
	DstChainId *big.Int
	Amount     *big.Int
	FeeAmount  *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterSwapStarted is a free log retrieval operation binding the contract event 0x074135076f5fc18420e0de96ce28c6bfe93048607465c707bdf3c10cf1e92d3f.
//
// Solidity: event SwapStarted(address indexed tokenAddr, address indexed sender, address indexed recipient, uint256 dstChainId, uint256 amount, uint256 feeAmount)
func (_ERC20SwapAgent *ERC20SwapAgentFilterer) FilterSwapStarted(opts *bind.FilterOpts, tokenAddr []common.Address, sender []common.Address, recipient []common.Address) (*ERC20SwapAgentSwapStartedIterator, error) {

	var tokenAddrRule []interface{}
	for _, tokenAddrItem := range tokenAddr {
		tokenAddrRule = append(tokenAddrRule, tokenAddrItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	logs, sub, err := _ERC20SwapAgent.contract.FilterLogs(opts, "SwapStarted", tokenAddrRule, senderRule, recipientRule)
	if err != nil {
		return nil, err
	}
	return &ERC20SwapAgentSwapStartedIterator{contract: _ERC20SwapAgent.contract, event: "SwapStarted", logs: logs, sub: sub}, nil
}

// WatchSwapStarted is a free log subscription operation binding the contract event 0x074135076f5fc18420e0de96ce28c6bfe93048607465c707bdf3c10cf1e92d3f.
//
// Solidity: event SwapStarted(address indexed tokenAddr, address indexed sender, address indexed recipient, uint256 dstChainId, uint256 amount, uint256 feeAmount)
func (_ERC20SwapAgent *ERC20SwapAgentFilterer) WatchSwapStarted(opts *bind.WatchOpts, sink chan<- *ERC20SwapAgentSwapStarted, tokenAddr []common.Address, sender []common.Address, recipient []common.Address) (event.Subscription, error) {

	var tokenAddrRule []interface{}
	for _, tokenAddrItem := range tokenAddr {
		tokenAddrRule = append(tokenAddrRule, tokenAddrItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	logs, sub, err := _ERC20SwapAgent.contract.WatchLogs(opts, "SwapStarted", tokenAddrRule, senderRule, recipientRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC20SwapAgentSwapStarted)
				if err := _ERC20SwapAgent.contract.UnpackLog(event, "SwapStarted", log); err != nil {
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

// ParseSwapStarted is a log parse operation binding the contract event 0x074135076f5fc18420e0de96ce28c6bfe93048607465c707bdf3c10cf1e92d3f.
//
// Solidity: event SwapStarted(address indexed tokenAddr, address indexed sender, address indexed recipient, uint256 dstChainId, uint256 amount, uint256 feeAmount)
func (_ERC20SwapAgent *ERC20SwapAgentFilterer) ParseSwapStarted(log types.Log) (*ERC20SwapAgentSwapStarted, error) {
	event := new(ERC20SwapAgentSwapStarted)
	if err := _ERC20SwapAgent.contract.UnpackLog(event, "SwapStarted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
