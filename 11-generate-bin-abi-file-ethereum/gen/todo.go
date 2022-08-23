// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package todo

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

// TodoTask is an auto generated low-level Go binding around an user-defined struct.
type TodoTask struct {
	Content string
	Status  bool
}

// TodoMetaData contains all meta data concerning the Todo contract.
var TodoMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_content\",\"type\":\"string\"}],\"name\":\"add\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_id\",\"type\":\"uint256\"}],\"name\":\"get\",\"outputs\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"content\",\"type\":\"string\"},{\"internalType\":\"bool\",\"name\":\"status\",\"type\":\"bool\"}],\"internalType\":\"structTodo.Task\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"list\",\"outputs\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"content\",\"type\":\"string\"},{\"internalType\":\"bool\",\"name\":\"status\",\"type\":\"bool\"}],\"internalType\":\"structTodo.Task[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_id\",\"type\":\"uint256\"}],\"name\":\"remove\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_id\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"_content\",\"type\":\"string\"}],\"name\":\"update\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b50336000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550610e96806100606000396000f3fe608060405234801561001057600080fd5b50600436106100625760003560e01c80630f560cd7146100675780634cc82215146100855780638da5cb5b146100a15780639507d39a146100bf578063b0c8f9dc146100ef578063f745630f1461010b575b600080fd5b61006f610127565b60405161007c9190610825565b60405180910390f35b61009f600480360381019061009a9190610891565b61028b565b005b6100a961032f565b6040516100b691906108ff565b60405180910390f35b6100d960048036038101906100d49190610891565b610353565b6040516100e69190610957565b60405180910390f35b61010960048036038101906101049190610aae565b610494565b005b61012560048036038101906101209190610af7565b61056a565b005b60603373ffffffffffffffffffffffffffffffffffffffff1660008054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff161461018157600080fd5b6001805480602002602001604051908101604052809291908181526020016000905b8282101561028257838290600052602060002090600202016040518060400160405290816000820180546101d690610b82565b80601f016020809104026020016040519081016040528092919081815260200182805461020290610b82565b801561024f5780601f106102245761010080835404028352916020019161024f565b820191906000526020600020905b81548152906001019060200180831161023257829003601f168201915b505050505081526020016001820160009054906101000a900460ff161515151581525050815260200190600101906101a3565b50505050905090565b3373ffffffffffffffffffffffffffffffffffffffff1660008054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16146102e357600080fd5b600181815481106102f7576102f6610bb3565b5b90600052602060002090600202016000808201600061031691906105f9565b6001820160006101000a81549060ff0219169055505050565b60008054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b61035b610639565b3373ffffffffffffffffffffffffffffffffffffffff1660008054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16146103b357600080fd5b600182815481106103c7576103c6610bb3565b5b90600052602060002090600202016040518060400160405290816000820180546103f090610b82565b80601f016020809104026020016040519081016040528092919081815260200182805461041c90610b82565b80156104695780601f1061043e57610100808354040283529160200191610469565b820191906000526020600020905b81548152906001019060200180831161044c57829003601f168201915b505050505081526020016001820160009054906101000a900460ff1615151515815250509050919050565b3373ffffffffffffffffffffffffffffffffffffffff1660008054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16146104ec57600080fd5b6001604051806040016040528083815260200160001515815250908060018154018082558091505060019003906000526020600020906002020160009091909190915060008201518160000190816105449190610d8e565b5060208201518160010160006101000a81548160ff021916908315150217905550505050565b3373ffffffffffffffffffffffffffffffffffffffff1660008054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16146105c257600080fd5b80600183815481106105d7576105d6610bb3565b5b906000526020600020906002020160000190816105f49190610d8e565b505050565b50805461060590610b82565b6000825580601f106106175750610636565b601f0160209004906000526020600020908101906106359190610655565b5b50565b6040518060400160405280606081526020016000151581525090565b5b8082111561066e576000816000905550600101610656565b5090565b600081519050919050565b600082825260208201905092915050565b6000819050602082019050919050565b600081519050919050565b600082825260208201905092915050565b60005b838110156106d85780820151818401526020810190506106bd565b838111156106e7576000848401525b50505050565b6000601f19601f8301169050919050565b60006107098261069e565b61071381856106a9565b93506107238185602086016106ba565b61072c816106ed565b840191505092915050565b60008115159050919050565b61074c81610737565b82525050565b6000604083016000830151848203600086015261076f82826106fe565b91505060208301516107846020860182610743565b508091505092915050565b600061079b8383610752565b905092915050565b6000602082019050919050565b60006107bb82610672565b6107c5818561067d565b9350836020820285016107d78561068e565b8060005b8581101561081357848403895281516107f4858261078f565b94506107ff836107a3565b925060208a019950506001810190506107db565b50829750879550505050505092915050565b6000602082019050818103600083015261083f81846107b0565b905092915050565b6000604051905090565b600080fd5b600080fd5b6000819050919050565b61086e8161085b565b811461087957600080fd5b50565b60008135905061088b81610865565b92915050565b6000602082840312156108a7576108a6610851565b5b60006108b58482850161087c565b91505092915050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b60006108e9826108be565b9050919050565b6108f9816108de565b82525050565b600060208201905061091460008301846108f0565b92915050565b6000604083016000830151848203600086015261093782826106fe565b915050602083015161094c6020860182610743565b508091505092915050565b60006020820190508181036000830152610971818461091a565b905092915050565b600080fd5b600080fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b6109bb826106ed565b810181811067ffffffffffffffff821117156109da576109d9610983565b5b80604052505050565b60006109ed610847565b90506109f982826109b2565b919050565b600067ffffffffffffffff821115610a1957610a18610983565b5b610a22826106ed565b9050602081019050919050565b82818337600083830152505050565b6000610a51610a4c846109fe565b6109e3565b905082815260208101848484011115610a6d57610a6c61097e565b5b610a78848285610a2f565b509392505050565b600082601f830112610a9557610a94610979565b5b8135610aa5848260208601610a3e565b91505092915050565b600060208284031215610ac457610ac3610851565b5b600082013567ffffffffffffffff811115610ae257610ae1610856565b5b610aee84828501610a80565b91505092915050565b60008060408385031215610b0e57610b0d610851565b5b6000610b1c8582860161087c565b925050602083013567ffffffffffffffff811115610b3d57610b3c610856565b5b610b4985828601610a80565b9150509250929050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b60006002820490506001821680610b9a57607f821691505b602082108103610bad57610bac610b53565b5b50919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b60008190508160005260206000209050919050565b60006020601f8301049050919050565b600082821b905092915050565b600060088302610c447fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82610c07565b610c4e8683610c07565b95508019841693508086168417925050509392505050565b6000819050919050565b6000610c8b610c86610c818461085b565b610c66565b61085b565b9050919050565b6000819050919050565b610ca583610c70565b610cb9610cb182610c92565b848454610c14565b825550505050565b600090565b610cce610cc1565b610cd9818484610c9c565b505050565b5b81811015610cfd57610cf2600082610cc6565b600181019050610cdf565b5050565b601f821115610d4257610d1381610be2565b610d1c84610bf7565b81016020851015610d2b578190505b610d3f610d3785610bf7565b830182610cde565b50505b505050565b600082821c905092915050565b6000610d6560001984600802610d47565b1980831691505092915050565b6000610d7e8383610d54565b9150826002028217905092915050565b610d978261069e565b67ffffffffffffffff811115610db057610daf610983565b5b610dba8254610b82565b610dc5828285610d01565b600060209050601f831160018114610df85760008415610de6578287015190505b610df08582610d72565b865550610e58565b601f198416610e0686610be2565b60005b82811015610e2e57848901518255600182019150602085019450602081019050610e09565b86831015610e4b5784890151610e47601f891682610d54565b8355505b6001600288020188555050505b50505050505056fea2646970667358221220f44f3cf3a38aba9973fb4bd0a1ffae86a4b5202f971034dc129f4cedeaa3917d64736f6c634300080f0033",
}

// TodoABI is the input ABI used to generate the binding from.
// Deprecated: Use TodoMetaData.ABI instead.
var TodoABI = TodoMetaData.ABI

// TodoBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use TodoMetaData.Bin instead.
var TodoBin = TodoMetaData.Bin

// DeployTodo deploys a new Ethereum contract, binding an instance of Todo to it.
func DeployTodo(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Todo, error) {
	parsed, err := TodoMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(TodoBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Todo{TodoCaller: TodoCaller{contract: contract}, TodoTransactor: TodoTransactor{contract: contract}, TodoFilterer: TodoFilterer{contract: contract}}, nil
}

// Todo is an auto generated Go binding around an Ethereum contract.
type Todo struct {
	TodoCaller     // Read-only binding to the contract
	TodoTransactor // Write-only binding to the contract
	TodoFilterer   // Log filterer for contract events
}

// TodoCaller is an auto generated read-only Go binding around an Ethereum contract.
type TodoCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TodoTransactor is an auto generated write-only Go binding around an Ethereum contract.
type TodoTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TodoFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type TodoFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TodoSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type TodoSession struct {
	Contract     *Todo             // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// TodoCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type TodoCallerSession struct {
	Contract *TodoCaller   // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// TodoTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type TodoTransactorSession struct {
	Contract     *TodoTransactor   // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// TodoRaw is an auto generated low-level Go binding around an Ethereum contract.
type TodoRaw struct {
	Contract *Todo // Generic contract binding to access the raw methods on
}

// TodoCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type TodoCallerRaw struct {
	Contract *TodoCaller // Generic read-only contract binding to access the raw methods on
}

// TodoTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type TodoTransactorRaw struct {
	Contract *TodoTransactor // Generic write-only contract binding to access the raw methods on
}

// NewTodo creates a new instance of Todo, bound to a specific deployed contract.
func NewTodo(address common.Address, backend bind.ContractBackend) (*Todo, error) {
	contract, err := bindTodo(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Todo{TodoCaller: TodoCaller{contract: contract}, TodoTransactor: TodoTransactor{contract: contract}, TodoFilterer: TodoFilterer{contract: contract}}, nil
}

// NewTodoCaller creates a new read-only instance of Todo, bound to a specific deployed contract.
func NewTodoCaller(address common.Address, caller bind.ContractCaller) (*TodoCaller, error) {
	contract, err := bindTodo(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &TodoCaller{contract: contract}, nil
}

// NewTodoTransactor creates a new write-only instance of Todo, bound to a specific deployed contract.
func NewTodoTransactor(address common.Address, transactor bind.ContractTransactor) (*TodoTransactor, error) {
	contract, err := bindTodo(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &TodoTransactor{contract: contract}, nil
}

// NewTodoFilterer creates a new log filterer instance of Todo, bound to a specific deployed contract.
func NewTodoFilterer(address common.Address, filterer bind.ContractFilterer) (*TodoFilterer, error) {
	contract, err := bindTodo(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &TodoFilterer{contract: contract}, nil
}

// bindTodo binds a generic wrapper to an already deployed contract.
func bindTodo(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(TodoABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Todo *TodoRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Todo.Contract.TodoCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Todo *TodoRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Todo.Contract.TodoTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Todo *TodoRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Todo.Contract.TodoTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Todo *TodoCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Todo.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Todo *TodoTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Todo.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Todo *TodoTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Todo.Contract.contract.Transact(opts, method, params...)
}

// Get is a free data retrieval call binding the contract method 0x9507d39a.
//
// Solidity: function get(uint256 _id) view returns((string,bool))
func (_Todo *TodoCaller) Get(opts *bind.CallOpts, _id *big.Int) (TodoTask, error) {
	var out []interface{}
	err := _Todo.contract.Call(opts, &out, "get", _id)

	if err != nil {
		return *new(TodoTask), err
	}

	out0 := *abi.ConvertType(out[0], new(TodoTask)).(*TodoTask)

	return out0, err

}

// Get is a free data retrieval call binding the contract method 0x9507d39a.
//
// Solidity: function get(uint256 _id) view returns((string,bool))
func (_Todo *TodoSession) Get(_id *big.Int) (TodoTask, error) {
	return _Todo.Contract.Get(&_Todo.CallOpts, _id)
}

// Get is a free data retrieval call binding the contract method 0x9507d39a.
//
// Solidity: function get(uint256 _id) view returns((string,bool))
func (_Todo *TodoCallerSession) Get(_id *big.Int) (TodoTask, error) {
	return _Todo.Contract.Get(&_Todo.CallOpts, _id)
}

// List is a free data retrieval call binding the contract method 0x0f560cd7.
//
// Solidity: function list() view returns((string,bool)[])
func (_Todo *TodoCaller) List(opts *bind.CallOpts) ([]TodoTask, error) {
	var out []interface{}
	err := _Todo.contract.Call(opts, &out, "list")

	if err != nil {
		return *new([]TodoTask), err
	}

	out0 := *abi.ConvertType(out[0], new([]TodoTask)).(*[]TodoTask)

	return out0, err

}

// List is a free data retrieval call binding the contract method 0x0f560cd7.
//
// Solidity: function list() view returns((string,bool)[])
func (_Todo *TodoSession) List() ([]TodoTask, error) {
	return _Todo.Contract.List(&_Todo.CallOpts)
}

// List is a free data retrieval call binding the contract method 0x0f560cd7.
//
// Solidity: function list() view returns((string,bool)[])
func (_Todo *TodoCallerSession) List() ([]TodoTask, error) {
	return _Todo.Contract.List(&_Todo.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Todo *TodoCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Todo.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Todo *TodoSession) Owner() (common.Address, error) {
	return _Todo.Contract.Owner(&_Todo.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Todo *TodoCallerSession) Owner() (common.Address, error) {
	return _Todo.Contract.Owner(&_Todo.CallOpts)
}

// Add is a paid mutator transaction binding the contract method 0xb0c8f9dc.
//
// Solidity: function add(string _content) returns()
func (_Todo *TodoTransactor) Add(opts *bind.TransactOpts, _content string) (*types.Transaction, error) {
	return _Todo.contract.Transact(opts, "add", _content)
}

// Add is a paid mutator transaction binding the contract method 0xb0c8f9dc.
//
// Solidity: function add(string _content) returns()
func (_Todo *TodoSession) Add(_content string) (*types.Transaction, error) {
	return _Todo.Contract.Add(&_Todo.TransactOpts, _content)
}

// Add is a paid mutator transaction binding the contract method 0xb0c8f9dc.
//
// Solidity: function add(string _content) returns()
func (_Todo *TodoTransactorSession) Add(_content string) (*types.Transaction, error) {
	return _Todo.Contract.Add(&_Todo.TransactOpts, _content)
}

// Remove is a paid mutator transaction binding the contract method 0x4cc82215.
//
// Solidity: function remove(uint256 _id) returns()
func (_Todo *TodoTransactor) Remove(opts *bind.TransactOpts, _id *big.Int) (*types.Transaction, error) {
	return _Todo.contract.Transact(opts, "remove", _id)
}

// Remove is a paid mutator transaction binding the contract method 0x4cc82215.
//
// Solidity: function remove(uint256 _id) returns()
func (_Todo *TodoSession) Remove(_id *big.Int) (*types.Transaction, error) {
	return _Todo.Contract.Remove(&_Todo.TransactOpts, _id)
}

// Remove is a paid mutator transaction binding the contract method 0x4cc82215.
//
// Solidity: function remove(uint256 _id) returns()
func (_Todo *TodoTransactorSession) Remove(_id *big.Int) (*types.Transaction, error) {
	return _Todo.Contract.Remove(&_Todo.TransactOpts, _id)
}

// Update is a paid mutator transaction binding the contract method 0xf745630f.
//
// Solidity: function update(uint256 _id, string _content) returns()
func (_Todo *TodoTransactor) Update(opts *bind.TransactOpts, _id *big.Int, _content string) (*types.Transaction, error) {
	return _Todo.contract.Transact(opts, "update", _id, _content)
}

// Update is a paid mutator transaction binding the contract method 0xf745630f.
//
// Solidity: function update(uint256 _id, string _content) returns()
func (_Todo *TodoSession) Update(_id *big.Int, _content string) (*types.Transaction, error) {
	return _Todo.Contract.Update(&_Todo.TransactOpts, _id, _content)
}

// Update is a paid mutator transaction binding the contract method 0xf745630f.
//
// Solidity: function update(uint256 _id, string _content) returns()
func (_Todo *TodoTransactorSession) Update(_id *big.Int, _content string) (*types.Transaction, error) {
	return _Todo.Contract.Update(&_Todo.TransactOpts, _id, _content)
}
