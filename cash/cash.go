// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package cash

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

// Check is an auto generated low-level Go binding around an user-defined struct.
type Check struct {
	Value        *big.Int
	TokenAddr    common.Address
	Nonce        uint64
	FromAddr     common.Address
	ToAddr       common.Address
	OpAddr       common.Address
	ContractAddr common.Address
	CheckSig     []byte
}

// Paycheck is an auto generated low-level Go binding around an user-defined struct.
type Paycheck struct {
	Check       Check
	PayValue    *big.Int
	PaycheckSig []byte
}

// CashMetaData contains all meta data concerning the Cash contract.
var CashMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"payable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"Paid\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"Received\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"node\",\"type\":\"address\"}],\"name\":\"get_nonce\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"get_owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"nodeNonce\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"tokenAddr\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"nonce\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"fromAddr\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"toAddr\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"opAddr\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"contractAddr\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"checkSig\",\"type\":\"bytes\"}],\"internalType\":\"structCheck\",\"name\":\"check\",\"type\":\"tuple\"},{\"internalType\":\"uint256\",\"name\":\"payValue\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"paycheckSig\",\"type\":\"bytes\"}],\"internalType\":\"structPaycheck\",\"name\":\"paycheck\",\"type\":\"tuple\"}],\"name\":\"withdraw\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
	Bin: "0x6080604052336000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550607b600160008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548167ffffffffffffffff021916908367ffffffffffffffff1602179055506112b7806100db6000396000f3fe6080604052600436106100435760003560e01c80630ac298dc146100885780631125b1df146100b35780631d728a0f146100e3578063e0d9045a1461012057610083565b36610083577f88a5966d370b9919b20f3e2c13ff65706f196a4e32cc2c12bf57088f885258743334604051610079929190610cf5565b60405180910390a1005b600080fd5b34801561009457600080fd5b5061009d61015d565b6040516100aa9190610cda565b60405180910390f35b6100cd60048036038101906100c89190610a4e565b610186565b6040516100da9190610d1e565b60405180910390f35b3480156100ef57600080fd5b5061010a60048036038101906101059190610a25565b610684565b6040516101179190610e3e565b60405180910390f35b34801561012c57600080fd5b5061014760048036038101906101429190610a25565b6106ab565b6040516101549190610e3e565b60405180910390f35b60008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905090565b60006001600083600001516080015173ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900467ffffffffffffffff1667ffffffffffffffff1682600001516040015167ffffffffffffffff161015610242576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161023990610d7e565b60405180910390fd5b81600001516000015182602001511115610291576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161028890610d9e565b60405180910390fd5b3373ffffffffffffffffffffffffffffffffffffffff1682600001516080015173ffffffffffffffffffffffffffffffffffffffff1614610307576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016102fe90610e1e565b60405180910390fd5b60008054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16826000015160a0015173ffffffffffffffffffffffffffffffffffffffff161461039d576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161039490610dfe565b60405180910390fd5b6000826000015160000151836000015160200151846000015160400151856000015160600151866000015160800151876000015160a00151886000015160c001516040516020016103f49796959493929190610c59565b6040516020818303038152906040529050600081805190602001209050600061042582866000015160e00151610708565b90508073ffffffffffffffffffffffffffffffffffffffff16856000015160a0015173ffffffffffffffffffffffffffffffffffffffff161461049d576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161049490610dbe565b60405180910390fd5b60008386602001516040516020016104b6929190610c31565b604051602081830303815290604052905060008180519060200120905060006104e3828960400151610708565b90508073ffffffffffffffffffffffffffffffffffffffff1688600001516060015173ffffffffffffffffffffffffffffffffffffffff161461055b576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161055290610dde565b60405180910390fd5b87600001516080015173ffffffffffffffffffffffffffffffffffffffff166108fc89602001519081150290604051600060405180830381858888f193505050501580156105ad573d6000803e3d6000fd5b507f737c69225d647e5994eab1a6c301bf6d9232beb2759ae1e27a8966b4732bc48988600001516080015189602001516040516105eb929190610cf5565b60405180910390a160018860000151604001516106089190610ed6565b600160008a600001516080015173ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548167ffffffffffffffff021916908367ffffffffffffffff16021790555060019650505050505050919050565b60016020528060005260406000206000915054906101000a900467ffffffffffffffff1681565b6000600160008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900467ffffffffffffffff169050919050565b6000604182511461071c576000905061080a565b60008060006020850151925060408501519150606085015160001a90507f7fffffffffffffffffffffffffffffff5d576e7357a4501ddfe92f46681b20a08260001c1115610770576000935050505061080a565b601b8160ff16101561078c57601b816107899190610f14565b90505b601b8160ff16141580156107a45750601c8160ff1614155b156107b5576000935050505061080a565b600186828585604051600081526020016040526040516107d89493929190610d39565b6020604051602081039080840390855afa1580156107fa573d6000803e3d6000fd5b5050506020604051035193505050505b92915050565b600061082361081e84610e7e565b610e59565b90508281526020810184848401111561083b57600080fd5b610846848285610fbe565b509392505050565b60008135905061085d8161123c565b92915050565b600082601f83011261087457600080fd5b8135610884848260208601610810565b91505092915050565b600061010082840312156108a057600080fd5b6108ab610100610e59565b905060006108bb848285016109fb565b60008301525060206108cf8482850161084e565b60208301525060406108e384828501610a10565b60408301525060606108f78482850161084e565b606083015250608061090b8482850161084e565b60808301525060a061091f8482850161084e565b60a08301525060c06109338482850161084e565b60c08301525060e082013567ffffffffffffffff81111561095357600080fd5b61095f84828501610863565b60e08301525092915050565b60006060828403121561097d57600080fd5b6109876060610e59565b9050600082013567ffffffffffffffff8111156109a357600080fd5b6109af8482850161088d565b60008301525060206109c3848285016109fb565b602083015250604082013567ffffffffffffffff8111156109e357600080fd5b6109ef84828501610863565b60408301525092915050565b600081359050610a0a81611253565b92915050565b600081359050610a1f8161126a565b92915050565b600060208284031215610a3757600080fd5b6000610a458482850161084e565b91505092915050565b600060208284031215610a6057600080fd5b600082013567ffffffffffffffff811115610a7a57600080fd5b610a868482850161096b565b91505092915050565b610a9881610f4b565b82525050565b610aaf610aaa82610f4b565b611031565b82525050565b610abe81610f5d565b82525050565b610acd81610f69565b82525050565b6000610ade82610eaf565b610ae88185610eba565b9350610af8818560208601610fcd565b80840191505092915050565b6000610b11601383610ec5565b9150610b1c826110fa565b602082019050919050565b6000610b34602a83610ec5565b9150610b3f82611123565b604082019050919050565b6000610b57601183610ec5565b9150610b6282611172565b602082019050919050565b6000610b7a601483610ec5565b9150610b858261119b565b602082019050919050565b6000610b9d602983610ec5565b9150610ba8826111c4565b604082019050919050565b6000610bc0601e83610ec5565b9150610bcb82611213565b602082019050919050565b610bdf81610f93565b82525050565b610bf6610bf182610f93565b611055565b82525050565b610c0581610f9d565b82525050565b610c1c610c1782610f9d565b61105f565b82525050565b610c2b81610fb1565b82525050565b6000610c3d8285610ad3565b9150610c498284610be5565b6020820191508190509392505050565b6000610c65828a610be5565b602082019150610c758289610a9e565b601482019150610c858288610c0b565b600882019150610c958287610a9e565b601482019150610ca58286610a9e565b601482019150610cb58285610a9e565b601482019150610cc58284610a9e565b60148201915081905098975050505050505050565b6000602082019050610cef6000830184610a8f565b92915050565b6000604082019050610d0a6000830185610a8f565b610d176020830184610bd6565b9392505050565b6000602082019050610d336000830184610ab5565b92915050565b6000608082019050610d4e6000830187610ac4565b610d5b6020830186610c22565b610d686040830185610ac4565b610d756060830184610ac4565b95945050505050565b60006020820190508181036000830152610d9781610b04565b9050919050565b60006020820190508181036000830152610db781610b27565b9050919050565b60006020820190508181036000830152610dd781610b4a565b9050919050565b60006020820190508181036000830152610df781610b6d565b9050919050565b60006020820190508181036000830152610e1781610b90565b9050919050565b60006020820190508181036000830152610e3781610bb3565b9050919050565b6000602082019050610e536000830184610bfc565b92915050565b6000610e63610e74565b9050610e6f8282611000565b919050565b6000604051905090565b600067ffffffffffffffff821115610e9957610e986110a0565b5b610ea2826110cf565b9050602081019050919050565b600081519050919050565b600081905092915050565b600082825260208201905092915050565b6000610ee182610f9d565b9150610eec83610f9d565b92508267ffffffffffffffff03821115610f0957610f08611071565b5b828201905092915050565b6000610f1f82610fb1565b9150610f2a83610fb1565b92508260ff03821115610f4057610f3f611071565b5b828201905092915050565b6000610f5682610f73565b9050919050565b60008115159050919050565b6000819050919050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000819050919050565b600067ffffffffffffffff82169050919050565b600060ff82169050919050565b82818337600083830152505050565b60005b83811015610feb578082015181840152602081019050610fd0565b83811115610ffa576000848401525b50505050565b611009826110cf565b810181811067ffffffffffffffff82111715611028576110276110a0565b5b80604052505050565b600061103c82611043565b9050919050565b600061104e826110ed565b9050919050565b6000819050919050565b600061106a826110e0565b9050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b6000601f19601f8301169050919050565b60008160c01b9050919050565b60008160601b9050919050565b7f636865636b2e6e6f6e636520746f6f206f6c6400000000000000000000000000600082015250565b7f70617976616c75652073686f756c64206e6f74206578636565642076616c756560008201527f206f6620636865636b2e00000000000000000000000000000000000000000000602082015250565b7f696c6c6567616c20636865636b20736967000000000000000000000000000000600082015250565b7f696c6c6567616c20706179636865636b20736967000000000000000000000000600082015250565b7f6f70657261746f722073686f756c64206265206f776e6572206f66207468697360008201527f20636f6e74726163740000000000000000000000000000000000000000000000602082015250565b7f63616c6c6572207368756f756c6420626520636865636b2e746f416464720000600082015250565b61124581610f4b565b811461125057600080fd5b50565b61125c81610f93565b811461126757600080fd5b50565b61127381610f9d565b811461127e57600080fd5b5056fea26469706673582212207449df33ff4dc5c7e89828c903c9f9871fc6971ad405b9f8073887bb97e2458164736f6c63430008020033",
}

// CashABI is the input ABI used to generate the binding from.
// Deprecated: Use CashMetaData.ABI instead.
var CashABI = CashMetaData.ABI

// CashBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use CashMetaData.Bin instead.
var CashBin = CashMetaData.Bin

// DeployCash deploys a new Ethereum contract, binding an instance of Cash to it.
func DeployCash(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Cash, error) {
	parsed, err := CashMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(CashBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Cash{CashCaller: CashCaller{contract: contract}, CashTransactor: CashTransactor{contract: contract}, CashFilterer: CashFilterer{contract: contract}}, nil
}

// Cash is an auto generated Go binding around an Ethereum contract.
type Cash struct {
	CashCaller     // Read-only binding to the contract
	CashTransactor // Write-only binding to the contract
	CashFilterer   // Log filterer for contract events
}

// CashCaller is an auto generated read-only Go binding around an Ethereum contract.
type CashCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CashTransactor is an auto generated write-only Go binding around an Ethereum contract.
type CashTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CashFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type CashFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CashSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type CashSession struct {
	Contract     *Cash             // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// CashCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type CashCallerSession struct {
	Contract *CashCaller   // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// CashTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type CashTransactorSession struct {
	Contract     *CashTransactor   // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// CashRaw is an auto generated low-level Go binding around an Ethereum contract.
type CashRaw struct {
	Contract *Cash // Generic contract binding to access the raw methods on
}

// CashCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type CashCallerRaw struct {
	Contract *CashCaller // Generic read-only contract binding to access the raw methods on
}

// CashTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type CashTransactorRaw struct {
	Contract *CashTransactor // Generic write-only contract binding to access the raw methods on
}

// NewCash creates a new instance of Cash, bound to a specific deployed contract.
func NewCash(address common.Address, backend bind.ContractBackend) (*Cash, error) {
	contract, err := bindCash(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Cash{CashCaller: CashCaller{contract: contract}, CashTransactor: CashTransactor{contract: contract}, CashFilterer: CashFilterer{contract: contract}}, nil
}

// NewCashCaller creates a new read-only instance of Cash, bound to a specific deployed contract.
func NewCashCaller(address common.Address, caller bind.ContractCaller) (*CashCaller, error) {
	contract, err := bindCash(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &CashCaller{contract: contract}, nil
}

// NewCashTransactor creates a new write-only instance of Cash, bound to a specific deployed contract.
func NewCashTransactor(address common.Address, transactor bind.ContractTransactor) (*CashTransactor, error) {
	contract, err := bindCash(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &CashTransactor{contract: contract}, nil
}

// NewCashFilterer creates a new log filterer instance of Cash, bound to a specific deployed contract.
func NewCashFilterer(address common.Address, filterer bind.ContractFilterer) (*CashFilterer, error) {
	contract, err := bindCash(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &CashFilterer{contract: contract}, nil
}

// bindCash binds a generic wrapper to an already deployed contract.
func bindCash(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(CashABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Cash *CashRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Cash.Contract.CashCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Cash *CashRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Cash.Contract.CashTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Cash *CashRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Cash.Contract.CashTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Cash *CashCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Cash.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Cash *CashTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Cash.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Cash *CashTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Cash.Contract.contract.Transact(opts, method, params...)
}

// GetNonce is a free data retrieval call binding the contract method 0xe0d9045a.
//
// Solidity: function get_nonce(address node) view returns(uint64)
func (_Cash *CashCaller) GetNonce(opts *bind.CallOpts, node common.Address) (uint64, error) {
	var out []interface{}
	err := _Cash.contract.Call(opts, &out, "get_nonce", node)

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// GetNonce is a free data retrieval call binding the contract method 0xe0d9045a.
//
// Solidity: function get_nonce(address node) view returns(uint64)
func (_Cash *CashSession) GetNonce(node common.Address) (uint64, error) {
	return _Cash.Contract.GetNonce(&_Cash.CallOpts, node)
}

// GetNonce is a free data retrieval call binding the contract method 0xe0d9045a.
//
// Solidity: function get_nonce(address node) view returns(uint64)
func (_Cash *CashCallerSession) GetNonce(node common.Address) (uint64, error) {
	return _Cash.Contract.GetNonce(&_Cash.CallOpts, node)
}

// GetOwner is a free data retrieval call binding the contract method 0x0ac298dc.
//
// Solidity: function get_owner() view returns(address)
func (_Cash *CashCaller) GetOwner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Cash.contract.Call(opts, &out, "get_owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetOwner is a free data retrieval call binding the contract method 0x0ac298dc.
//
// Solidity: function get_owner() view returns(address)
func (_Cash *CashSession) GetOwner() (common.Address, error) {
	return _Cash.Contract.GetOwner(&_Cash.CallOpts)
}

// GetOwner is a free data retrieval call binding the contract method 0x0ac298dc.
//
// Solidity: function get_owner() view returns(address)
func (_Cash *CashCallerSession) GetOwner() (common.Address, error) {
	return _Cash.Contract.GetOwner(&_Cash.CallOpts)
}

// NodeNonce is a free data retrieval call binding the contract method 0x1d728a0f.
//
// Solidity: function nodeNonce(address ) view returns(uint64)
func (_Cash *CashCaller) NodeNonce(opts *bind.CallOpts, arg0 common.Address) (uint64, error) {
	var out []interface{}
	err := _Cash.contract.Call(opts, &out, "nodeNonce", arg0)

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// NodeNonce is a free data retrieval call binding the contract method 0x1d728a0f.
//
// Solidity: function nodeNonce(address ) view returns(uint64)
func (_Cash *CashSession) NodeNonce(arg0 common.Address) (uint64, error) {
	return _Cash.Contract.NodeNonce(&_Cash.CallOpts, arg0)
}

// NodeNonce is a free data retrieval call binding the contract method 0x1d728a0f.
//
// Solidity: function nodeNonce(address ) view returns(uint64)
func (_Cash *CashCallerSession) NodeNonce(arg0 common.Address) (uint64, error) {
	return _Cash.Contract.NodeNonce(&_Cash.CallOpts, arg0)
}

// Withdraw is a paid mutator transaction binding the contract method 0x1125b1df.
//
// Solidity: function withdraw(((uint256,address,uint64,address,address,address,address,bytes),uint256,bytes) paycheck) payable returns(bool)
func (_Cash *CashTransactor) Withdraw(opts *bind.TransactOpts, paycheck Paycheck) (*types.Transaction, error) {
	return _Cash.contract.Transact(opts, "withdraw", paycheck)
}

// Withdraw is a paid mutator transaction binding the contract method 0x1125b1df.
//
// Solidity: function withdraw(((uint256,address,uint64,address,address,address,address,bytes),uint256,bytes) paycheck) payable returns(bool)
func (_Cash *CashSession) Withdraw(paycheck Paycheck) (*types.Transaction, error) {
	return _Cash.Contract.Withdraw(&_Cash.TransactOpts, paycheck)
}

// Withdraw is a paid mutator transaction binding the contract method 0x1125b1df.
//
// Solidity: function withdraw(((uint256,address,uint64,address,address,address,address,bytes),uint256,bytes) paycheck) payable returns(bool)
func (_Cash *CashTransactorSession) Withdraw(paycheck Paycheck) (*types.Transaction, error) {
	return _Cash.Contract.Withdraw(&_Cash.TransactOpts, paycheck)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Cash *CashTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Cash.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Cash *CashSession) Receive() (*types.Transaction, error) {
	return _Cash.Contract.Receive(&_Cash.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Cash *CashTransactorSession) Receive() (*types.Transaction, error) {
	return _Cash.Contract.Receive(&_Cash.TransactOpts)
}

// CashPaidIterator is returned from FilterPaid and is used to iterate over the raw logs and unpacked data for Paid events raised by the Cash contract.
type CashPaidIterator struct {
	Event *CashPaid // Event containing the contract specifics and raw log

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
func (it *CashPaidIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CashPaid)
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
		it.Event = new(CashPaid)
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
func (it *CashPaidIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CashPaidIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CashPaid represents a Paid event raised by the Cash contract.
type CashPaid struct {
	Arg0 common.Address
	Arg1 *big.Int
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterPaid is a free log retrieval operation binding the contract event 0x737c69225d647e5994eab1a6c301bf6d9232beb2759ae1e27a8966b4732bc489.
//
// Solidity: event Paid(address arg0, uint256 arg1)
func (_Cash *CashFilterer) FilterPaid(opts *bind.FilterOpts) (*CashPaidIterator, error) {

	logs, sub, err := _Cash.contract.FilterLogs(opts, "Paid")
	if err != nil {
		return nil, err
	}
	return &CashPaidIterator{contract: _Cash.contract, event: "Paid", logs: logs, sub: sub}, nil
}

// WatchPaid is a free log subscription operation binding the contract event 0x737c69225d647e5994eab1a6c301bf6d9232beb2759ae1e27a8966b4732bc489.
//
// Solidity: event Paid(address arg0, uint256 arg1)
func (_Cash *CashFilterer) WatchPaid(opts *bind.WatchOpts, sink chan<- *CashPaid) (event.Subscription, error) {

	logs, sub, err := _Cash.contract.WatchLogs(opts, "Paid")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CashPaid)
				if err := _Cash.contract.UnpackLog(event, "Paid", log); err != nil {
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

// ParsePaid is a log parse operation binding the contract event 0x737c69225d647e5994eab1a6c301bf6d9232beb2759ae1e27a8966b4732bc489.
//
// Solidity: event Paid(address arg0, uint256 arg1)
func (_Cash *CashFilterer) ParsePaid(log types.Log) (*CashPaid, error) {
	event := new(CashPaid)
	if err := _Cash.contract.UnpackLog(event, "Paid", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CashReceivedIterator is returned from FilterReceived and is used to iterate over the raw logs and unpacked data for Received events raised by the Cash contract.
type CashReceivedIterator struct {
	Event *CashReceived // Event containing the contract specifics and raw log

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
func (it *CashReceivedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CashReceived)
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
		it.Event = new(CashReceived)
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
func (it *CashReceivedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CashReceivedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CashReceived represents a Received event raised by the Cash contract.
type CashReceived struct {
	Arg0 common.Address
	Arg1 *big.Int
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterReceived is a free log retrieval operation binding the contract event 0x88a5966d370b9919b20f3e2c13ff65706f196a4e32cc2c12bf57088f88525874.
//
// Solidity: event Received(address arg0, uint256 arg1)
func (_Cash *CashFilterer) FilterReceived(opts *bind.FilterOpts) (*CashReceivedIterator, error) {

	logs, sub, err := _Cash.contract.FilterLogs(opts, "Received")
	if err != nil {
		return nil, err
	}
	return &CashReceivedIterator{contract: _Cash.contract, event: "Received", logs: logs, sub: sub}, nil
}

// WatchReceived is a free log subscription operation binding the contract event 0x88a5966d370b9919b20f3e2c13ff65706f196a4e32cc2c12bf57088f88525874.
//
// Solidity: event Received(address arg0, uint256 arg1)
func (_Cash *CashFilterer) WatchReceived(opts *bind.WatchOpts, sink chan<- *CashReceived) (event.Subscription, error) {

	logs, sub, err := _Cash.contract.WatchLogs(opts, "Received")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CashReceived)
				if err := _Cash.contract.UnpackLog(event, "Received", log); err != nil {
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
// Solidity: event Received(address arg0, uint256 arg1)
func (_Cash *CashFilterer) ParseReceived(log types.Log) (*CashReceived, error) {
	event := new(CashReceived)
	if err := _Cash.contract.UnpackLog(event, "Received", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
