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
	ABI: "[{\"inputs\":[],\"stateMutability\":\"payable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"Paid\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"Received\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"deposit\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getBalance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"node\",\"type\":\"address\"}],\"name\":\"getNonce\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getOwner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"nodeNonce\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"node\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"nonce\",\"type\":\"uint64\"}],\"name\":\"setNonce\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"tokenAddr\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"nonce\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"fromAddr\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"toAddr\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"opAddr\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"contractAddr\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"checkSig\",\"type\":\"bytes\"}],\"internalType\":\"structCheck\",\"name\":\"check\",\"type\":\"tuple\"},{\"internalType\":\"uint256\",\"name\":\"payValue\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"paycheckSig\",\"type\":\"bytes\"}],\"internalType\":\"structPaycheck\",\"name\":\"paycheck\",\"type\":\"tuple\"}],\"name\":\"withdraw\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
	Bin: "0x6080604052336000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550611351806100536000396000f3fe6080604052600436106100745760003560e01c80632d0335ab1161004e5780632d0335ab14610151578063893d20e81461018e578063d0e30db0146101b9578063f8e18b57146101c3576100b4565b80631125b1df146100b957806312065fe0146100e95780631d728a0f14610114576100b4565b366100b4577f88a5966d370b9919b20f3e2c13ff65706f196a4e32cc2c12bf57088f8852587433346040516100aa929190610db2565b60405180910390a1005b600080fd5b6100d360048036038101906100ce9190610b0b565b6101ec565b6040516100e09190610ddb565b60405180910390f35b3480156100f557600080fd5b506100fe610668565b60405161010b9190610efb565b60405180910390f35b34801561012057600080fd5b5061013b60048036038101906101369190610aa6565b610670565b6040516101489190610f16565b60405180910390f35b34801561015d57600080fd5b5061017860048036038101906101739190610aa6565b610697565b6040516101859190610f16565b60405180910390f35b34801561019a57600080fd5b506101a36106f4565b6040516101b09190610d97565b60405180910390f35b6101c161071d565b005b3480156101cf57600080fd5b506101ea60048036038101906101e59190610acf565b61071f565b005b60006001600083600001516080015173ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900467ffffffffffffffff1667ffffffffffffffff1682600001516040015167ffffffffffffffff1610156102a8576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161029f90610e3b565b60405180910390fd5b816000015160000151826020015111156102f7576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016102ee90610e5b565b60405180910390fd5b3373ffffffffffffffffffffffffffffffffffffffff1682600001516080015173ffffffffffffffffffffffffffffffffffffffff161461036d576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161036490610edb565b60405180910390fd5b60008054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16826000015160a0015173ffffffffffffffffffffffffffffffffffffffff1614610403576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016103fa90610ebb565b60405180910390fd5b6000826000015160000151836000015160200151846000015160400151856000015160600151866000015160800151876000015160a00151886000015160c0015160405160200161045a9796959493929190610d16565b6040516020818303038152906040529050600081805190602001209050600061048b82866000015160e00151610789565b90508073ffffffffffffffffffffffffffffffffffffffff16856000015160a0015173ffffffffffffffffffffffffffffffffffffffff1614610503576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016104fa90610e7b565b60405180910390fd5b600083866020015160405160200161051c929190610cee565b60405160208183030381529060405290506000818051906020012090506000610549828960400151610789565b90508073ffffffffffffffffffffffffffffffffffffffff1688600001516060015173ffffffffffffffffffffffffffffffffffffffff16146105c1576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016105b890610e9b565b60405180910390fd5b87600001516080015173ffffffffffffffffffffffffffffffffffffffff166108fc89602001519081150290604051600060405180830381858888f19350505050158015610613573d6000803e3d6000fd5b507f737c69225d647e5994eab1a6c301bf6d9232beb2759ae1e27a8966b4732bc4898860000151608001518960200151604051610651929190610db2565b60405180910390a160019650505050505050919050565b600047905090565b60016020528060005260406000206000915054906101000a900467ffffffffffffffff1681565b6000600160008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900467ffffffffffffffff169050919050565b60008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905090565b565b80600160008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548167ffffffffffffffff021916908367ffffffffffffffff1602179055505050565b6000604182511461079d576000905061088b565b60008060006020850151925060408501519150606085015160001a90507f7fffffffffffffffffffffffffffffff5d576e7357a4501ddfe92f46681b20a08260001c11156107f1576000935050505061088b565b601b8160ff16101561080d57601b8161080a9190610fae565b90505b601b8160ff16141580156108255750601c8160ff1614155b15610836576000935050505061088b565b600186828585604051600081526020016040526040516108599493929190610df6565b6020604051602081039080840390855afa15801561087b573d6000803e3d6000fd5b5050506020604051035193505050505b92915050565b60006108a461089f84610f56565b610f31565b9050828152602081018484840111156108bc57600080fd5b6108c7848285611058565b509392505050565b6000813590506108de816112d6565b92915050565b600082601f8301126108f557600080fd5b8135610905848260208601610891565b91505092915050565b6000610100828403121561092157600080fd5b61092c610100610f31565b9050600061093c84828501610a7c565b6000830152506020610950848285016108cf565b602083015250604061096484828501610a91565b6040830152506060610978848285016108cf565b606083015250608061098c848285016108cf565b60808301525060a06109a0848285016108cf565b60a08301525060c06109b4848285016108cf565b60c08301525060e082013567ffffffffffffffff8111156109d457600080fd5b6109e0848285016108e4565b60e08301525092915050565b6000606082840312156109fe57600080fd5b610a086060610f31565b9050600082013567ffffffffffffffff811115610a2457600080fd5b610a308482850161090e565b6000830152506020610a4484828501610a7c565b602083015250604082013567ffffffffffffffff811115610a6457600080fd5b610a70848285016108e4565b60408301525092915050565b600081359050610a8b816112ed565b92915050565b600081359050610aa081611304565b92915050565b600060208284031215610ab857600080fd5b6000610ac6848285016108cf565b91505092915050565b60008060408385031215610ae257600080fd5b6000610af0858286016108cf565b9250506020610b0185828601610a91565b9150509250929050565b600060208284031215610b1d57600080fd5b600082013567ffffffffffffffff811115610b3757600080fd5b610b43848285016109ec565b91505092915050565b610b5581610fe5565b82525050565b610b6c610b6782610fe5565b6110cb565b82525050565b610b7b81610ff7565b82525050565b610b8a81611003565b82525050565b6000610b9b82610f87565b610ba58185610f92565b9350610bb5818560208601611067565b80840191505092915050565b6000610bce601383610f9d565b9150610bd982611194565b602082019050919050565b6000610bf1602a83610f9d565b9150610bfc826111bd565b604082019050919050565b6000610c14601183610f9d565b9150610c1f8261120c565b602082019050919050565b6000610c37601483610f9d565b9150610c4282611235565b602082019050919050565b6000610c5a602983610f9d565b9150610c658261125e565b604082019050919050565b6000610c7d601e83610f9d565b9150610c88826112ad565b602082019050919050565b610c9c8161102d565b82525050565b610cb3610cae8261102d565b6110ef565b82525050565b610cc281611037565b82525050565b610cd9610cd482611037565b6110f9565b82525050565b610ce88161104b565b82525050565b6000610cfa8285610b90565b9150610d068284610ca2565b6020820191508190509392505050565b6000610d22828a610ca2565b602082019150610d328289610b5b565b601482019150610d428288610cc8565b600882019150610d528287610b5b565b601482019150610d628286610b5b565b601482019150610d728285610b5b565b601482019150610d828284610b5b565b60148201915081905098975050505050505050565b6000602082019050610dac6000830184610b4c565b92915050565b6000604082019050610dc76000830185610b4c565b610dd46020830184610c93565b9392505050565b6000602082019050610df06000830184610b72565b92915050565b6000608082019050610e0b6000830187610b81565b610e186020830186610cdf565b610e256040830185610b81565b610e326060830184610b81565b95945050505050565b60006020820190508181036000830152610e5481610bc1565b9050919050565b60006020820190508181036000830152610e7481610be4565b9050919050565b60006020820190508181036000830152610e9481610c07565b9050919050565b60006020820190508181036000830152610eb481610c2a565b9050919050565b60006020820190508181036000830152610ed481610c4d565b9050919050565b60006020820190508181036000830152610ef481610c70565b9050919050565b6000602082019050610f106000830184610c93565b92915050565b6000602082019050610f2b6000830184610cb9565b92915050565b6000610f3b610f4c565b9050610f47828261109a565b919050565b6000604051905090565b600067ffffffffffffffff821115610f7157610f7061113a565b5b610f7a82611169565b9050602081019050919050565b600081519050919050565b600081905092915050565b600082825260208201905092915050565b6000610fb98261104b565b9150610fc48361104b565b92508260ff03821115610fda57610fd961110b565b5b828201905092915050565b6000610ff08261100d565b9050919050565b60008115159050919050565b6000819050919050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000819050919050565b600067ffffffffffffffff82169050919050565b600060ff82169050919050565b82818337600083830152505050565b60005b8381101561108557808201518184015260208101905061106a565b83811115611094576000848401525b50505050565b6110a382611169565b810181811067ffffffffffffffff821117156110c2576110c161113a565b5b80604052505050565b60006110d6826110dd565b9050919050565b60006110e882611187565b9050919050565b6000819050919050565b60006111048261117a565b9050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b6000601f19601f8301169050919050565b60008160c01b9050919050565b60008160601b9050919050565b7f636865636b2e6e6f6e636520746f6f206f6c6400000000000000000000000000600082015250565b7f70617976616c75652073686f756c64206e6f74206578636565642076616c756560008201527f206f6620636865636b2e00000000000000000000000000000000000000000000602082015250565b7f696c6c6567616c20636865636b20736967000000000000000000000000000000600082015250565b7f696c6c6567616c20706179636865636b20736967000000000000000000000000600082015250565b7f6f70657261746f722073686f756c64206265206f776e6572206f66207468697360008201527f20636f6e74726163740000000000000000000000000000000000000000000000602082015250565b7f63616c6c6572207368756f756c6420626520636865636b2e746f416464720000600082015250565b6112df81610fe5565b81146112ea57600080fd5b50565b6112f68161102d565b811461130157600080fd5b50565b61130d81611037565b811461131857600080fd5b5056fea26469706673582212209e43e8473784ea1f219050e06a9138ee6375e3c36366e7633faec1c88e3aadb664736f6c63430008020033",
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

// GetBalance is a free data retrieval call binding the contract method 0x12065fe0.
//
// Solidity: function getBalance() view returns(uint256)
func (_Cash *CashCaller) GetBalance(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Cash.contract.Call(opts, &out, "getBalance")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetBalance is a free data retrieval call binding the contract method 0x12065fe0.
//
// Solidity: function getBalance() view returns(uint256)
func (_Cash *CashSession) GetBalance() (*big.Int, error) {
	return _Cash.Contract.GetBalance(&_Cash.CallOpts)
}

// GetBalance is a free data retrieval call binding the contract method 0x12065fe0.
//
// Solidity: function getBalance() view returns(uint256)
func (_Cash *CashCallerSession) GetBalance() (*big.Int, error) {
	return _Cash.Contract.GetBalance(&_Cash.CallOpts)
}

// GetNonce is a free data retrieval call binding the contract method 0x2d0335ab.
//
// Solidity: function getNonce(address node) view returns(uint64)
func (_Cash *CashCaller) GetNonce(opts *bind.CallOpts, node common.Address) (uint64, error) {
	var out []interface{}
	err := _Cash.contract.Call(opts, &out, "getNonce", node)

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// GetNonce is a free data retrieval call binding the contract method 0x2d0335ab.
//
// Solidity: function getNonce(address node) view returns(uint64)
func (_Cash *CashSession) GetNonce(node common.Address) (uint64, error) {
	return _Cash.Contract.GetNonce(&_Cash.CallOpts, node)
}

// GetNonce is a free data retrieval call binding the contract method 0x2d0335ab.
//
// Solidity: function getNonce(address node) view returns(uint64)
func (_Cash *CashCallerSession) GetNonce(node common.Address) (uint64, error) {
	return _Cash.Contract.GetNonce(&_Cash.CallOpts, node)
}

// GetOwner is a free data retrieval call binding the contract method 0x893d20e8.
//
// Solidity: function getOwner() view returns(address)
func (_Cash *CashCaller) GetOwner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Cash.contract.Call(opts, &out, "getOwner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetOwner is a free data retrieval call binding the contract method 0x893d20e8.
//
// Solidity: function getOwner() view returns(address)
func (_Cash *CashSession) GetOwner() (common.Address, error) {
	return _Cash.Contract.GetOwner(&_Cash.CallOpts)
}

// GetOwner is a free data retrieval call binding the contract method 0x893d20e8.
//
// Solidity: function getOwner() view returns(address)
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

// Deposit is a paid mutator transaction binding the contract method 0xd0e30db0.
//
// Solidity: function deposit() payable returns()
func (_Cash *CashTransactor) Deposit(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Cash.contract.Transact(opts, "deposit")
}

// Deposit is a paid mutator transaction binding the contract method 0xd0e30db0.
//
// Solidity: function deposit() payable returns()
func (_Cash *CashSession) Deposit() (*types.Transaction, error) {
	return _Cash.Contract.Deposit(&_Cash.TransactOpts)
}

// Deposit is a paid mutator transaction binding the contract method 0xd0e30db0.
//
// Solidity: function deposit() payable returns()
func (_Cash *CashTransactorSession) Deposit() (*types.Transaction, error) {
	return _Cash.Contract.Deposit(&_Cash.TransactOpts)
}

// SetNonce is a paid mutator transaction binding the contract method 0xf8e18b57.
//
// Solidity: function setNonce(address node, uint64 nonce) returns()
func (_Cash *CashTransactor) SetNonce(opts *bind.TransactOpts, node common.Address, nonce uint64) (*types.Transaction, error) {
	return _Cash.contract.Transact(opts, "setNonce", node, nonce)
}

// SetNonce is a paid mutator transaction binding the contract method 0xf8e18b57.
//
// Solidity: function setNonce(address node, uint64 nonce) returns()
func (_Cash *CashSession) SetNonce(node common.Address, nonce uint64) (*types.Transaction, error) {
	return _Cash.Contract.SetNonce(&_Cash.TransactOpts, node, nonce)
}

// SetNonce is a paid mutator transaction binding the contract method 0xf8e18b57.
//
// Solidity: function setNonce(address node, uint64 nonce) returns()
func (_Cash *CashTransactorSession) SetNonce(node common.Address, nonce uint64) (*types.Transaction, error) {
	return _Cash.Contract.SetNonce(&_Cash.TransactOpts, node, nonce)
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
