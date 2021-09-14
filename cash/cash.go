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
	Nonce        *big.Int
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
	ABI: "[{\"inputs\":[],\"stateMutability\":\"payable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"Paid\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"Received\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"Show256\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"Showaddr\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"name\":\"Showbytes\",\"type\":\"event\"},{\"inputs\":[{\"components\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"tokenAddr\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"fromAddr\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"toAddr\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"opAddr\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"contractAddr\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"checkSig\",\"type\":\"bytes\"}],\"internalType\":\"structCheck\",\"name\":\"check\",\"type\":\"tuple\"},{\"internalType\":\"uint256\",\"name\":\"payValue\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"paycheckSig\",\"type\":\"bytes\"}],\"internalType\":\"structPaycheck\",\"name\":\"paycheck\",\"type\":\"tuple\"}],\"name\":\"apply_check\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"node\",\"type\":\"address\"}],\"name\":\"get_node_nonce\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"get_owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"nodeNonce\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"x\",\"type\":\"uint256\"}],\"name\":\"toBytes1\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"b\",\"type\":\"bytes\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
	Bin: "0x6080604052336000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555061155a806100536000396000f3fe60806040526004361061004e5760003560e01c80630ac298dc146100935780631d728a0f146100be5780633cb754d8146100fb5780635f510b7b14610138578063c8193561146101755761008e565b3661008e577f88a5966d370b9919b20f3e2c13ff65706f196a4e32cc2c12bf57088f885258743334604051610084929190610f97565b60405180910390a1005b600080fd5b34801561009f57600080fd5b506100a86101a5565b6040516100b59190610f7c565b60405180910390f35b3480156100ca57600080fd5b506100e560048036038101906100e09190610c8b565b6101ce565b6040516100f29190611102565b60405180910390f35b34801561010757600080fd5b50610122600480360381019061011d9190610cf5565b6101e6565b60405161012f9190611020565b60405180910390f35b34801561014457600080fd5b5061015f600480360381019061015a9190610c8b565b610269565b60405161016c9190611102565b60405180910390f35b61018f600480360381019061018a9190610cb4565b6102b2565b60405161019c9190610fc0565b60405180910390f35b60008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905090565b60016020528060005260406000206000915090505481565b6060602067ffffffffffffffff811115610229577f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b6040519080825280601f01601f19166020018201604052801561025b5781602001600182028036833780820191505090505b509050816020820152919050565b6000600160008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020549050919050565b60007fcbe629aaa9d71fcba91cbad4676731b37fd435513517a19f0012fe8804a10afb8260000151604001516040516102eb9190611102565b60405180910390a17fcbe629aaa9d71fcba91cbad4676731b37fd435513517a19f0012fe8804a10afb82602001516040516103269190611102565b60405180910390a17f7876b0f7cc2519df0fdac81db62f5a4a05f5332c05ceb4d667b783eef87094a38260000151608001516040516103659190610f7c565b60405180910390a17f7876b0f7cc2519df0fdac81db62f5a4a05f5332c05ceb4d667b783eef87094a33360405161039c9190610f7c565b60405180910390a17f7876b0f7cc2519df0fdac81db62f5a4a05f5332c05ceb4d667b783eef87094a3826000015160a001516040516103db9190610f7c565b60405180910390a17f7876b0f7cc2519df0fdac81db62f5a4a05f5332c05ceb4d667b783eef87094a360008054906101000a900473ffffffffffffffffffffffffffffffffffffffff166040516104329190610f7c565b60405180910390a17f4daa2d63a96f5e545f8ff5b850799b444e2e8a26e69f59425ba735fa0528a268826000015160e001516040516104719190611020565b60405180910390a17f4daa2d63a96f5e545f8ff5b850799b444e2e8a26e69f59425ba735fa0528a26882604001516040516104ac9190611020565b60405180910390a16001600083600001516080015173ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020548260000151604001511015610546576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161053d90611042565b60405180910390fd5b81600001516000015182602001511115610595576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161058c90611062565b60405180910390fd5b3373ffffffffffffffffffffffffffffffffffffffff1682600001516080015173ffffffffffffffffffffffffffffffffffffffff161461060b576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610602906110e2565b60405180910390fd5b60008054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16826000015160a0015173ffffffffffffffffffffffffffffffffffffffff16146106a1576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610698906110c2565b60405180910390fd5b6000826000015160000151836000015160200151846000015160400151856000015160600151866000015160800151876000015160a00151886000015160c001516040516020016106f89796959493929190610efb565b6040516020818303038152906040529050600081805190602001209050600061072982866000015160e0015161096d565b90508073ffffffffffffffffffffffffffffffffffffffff16856000015160a0015173ffffffffffffffffffffffffffffffffffffffff16146107a1576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161079890611082565b60405180910390fd5b60008386602001516040516020016107ba929190610ed3565b604051602081830303815290604052905060008180519060200120905060006107e782896040015161096d565b90508073ffffffffffffffffffffffffffffffffffffffff1688600001516060015173ffffffffffffffffffffffffffffffffffffffff161461085f576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610856906110a2565b60405180910390fd5b87600001516080015173ffffffffffffffffffffffffffffffffffffffff166108fc89602001519081150290604051600060405180830381858888f193505050501580156108b1573d6000803e3d6000fd5b507f737c69225d647e5994eab1a6c301bf6d9232beb2759ae1e27a8966b4732bc48988600001516080015189602001516040516108ef929190610f97565b60405180910390a16109136001896000015160400151610a7590919063ffffffff16565b600160008a600001516080015173ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000208190555060019650505050505050919050565b600060418251146109815760009050610a6f565b60008060006020850151925060408501519150606085015160001a90507f7fffffffffffffffffffffffffffffff5d576e7357a4501ddfe92f46681b20a08260001c11156109d55760009350505050610a6f565b601b8160ff1610156109f157601b816109ee9190611201565b90505b601b8160ff1614158015610a095750601c8160ff1614155b15610a1a5760009350505050610a6f565b60018682858560405160008152602001604052604051610a3d9493929190610fdb565b6020604051602081039080840390855afa158015610a5f573d6000803e3d6000fd5b5050506020604051035193505050505b92915050565b60008183610a8391906111ab565b905092915050565b6000610a9e610a9984611142565b61111d565b905082815260208101848484011115610ab657600080fd5b610ac1848285611297565b509392505050565b600081359050610ad8816114f6565b92915050565b600082601f830112610aef57600080fd5b8135610aff848260208601610a8b565b91505092915050565b60006101008284031215610b1b57600080fd5b610b2661010061111d565b90506000610b3684828501610c76565b6000830152506020610b4a84828501610ac9565b6020830152506040610b5e84828501610c76565b6040830152506060610b7284828501610ac9565b6060830152506080610b8684828501610ac9565b60808301525060a0610b9a84828501610ac9565b60a08301525060c0610bae84828501610ac9565b60c08301525060e082013567ffffffffffffffff811115610bce57600080fd5b610bda84828501610ade565b60e08301525092915050565b600060608284031215610bf857600080fd5b610c02606061111d565b9050600082013567ffffffffffffffff811115610c1e57600080fd5b610c2a84828501610b08565b6000830152506020610c3e84828501610c76565b602083015250604082013567ffffffffffffffff811115610c5e57600080fd5b610c6a84828501610ade565b60408301525092915050565b600081359050610c858161150d565b92915050565b600060208284031215610c9d57600080fd5b6000610cab84828501610ac9565b91505092915050565b600060208284031215610cc657600080fd5b600082013567ffffffffffffffff811115610ce057600080fd5b610cec84828501610be6565b91505092915050565b600060208284031215610d0757600080fd5b6000610d1584828501610c76565b91505092915050565b610d2781611238565b82525050565b610d3e610d3982611238565b61130a565b82525050565b610d4d8161124a565b82525050565b610d5c81611256565b82525050565b6000610d6d82611173565b610d77818561117e565b9350610d878185602086016112a6565b610d9081611396565b840191505092915050565b6000610da682611173565b610db0818561118f565b9350610dc08185602086016112a6565b80840191505092915050565b6000610dd960138361119a565b9150610de4826113b4565b602082019050919050565b6000610dfc602a8361119a565b9150610e07826113dd565b604082019050919050565b6000610e1f60118361119a565b9150610e2a8261142c565b602082019050919050565b6000610e4260148361119a565b9150610e4d82611455565b602082019050919050565b6000610e6560298361119a565b9150610e708261147e565b604082019050919050565b6000610e88601e8361119a565b9150610e93826114cd565b602082019050919050565b610ea781611280565b82525050565b610ebe610eb982611280565b61132e565b82525050565b610ecd8161128a565b82525050565b6000610edf8285610d9b565b9150610eeb8284610ead565b6020820191508190509392505050565b6000610f07828a610ead565b602082019150610f178289610d2d565b601482019150610f278288610ead565b602082019150610f378287610d2d565b601482019150610f478286610d2d565b601482019150610f578285610d2d565b601482019150610f678284610d2d565b60148201915081905098975050505050505050565b6000602082019050610f916000830184610d1e565b92915050565b6000604082019050610fac6000830185610d1e565b610fb96020830184610e9e565b9392505050565b6000602082019050610fd56000830184610d44565b92915050565b6000608082019050610ff06000830187610d53565b610ffd6020830186610ec4565b61100a6040830185610d53565b6110176060830184610d53565b95945050505050565b6000602082019050818103600083015261103a8184610d62565b905092915050565b6000602082019050818103600083015261105b81610dcc565b9050919050565b6000602082019050818103600083015261107b81610def565b9050919050565b6000602082019050818103600083015261109b81610e12565b9050919050565b600060208201905081810360008301526110bb81610e35565b9050919050565b600060208201905081810360008301526110db81610e58565b9050919050565b600060208201905081810360008301526110fb81610e7b565b9050919050565b60006020820190506111176000830184610e9e565b92915050565b6000611127611138565b905061113382826112d9565b919050565b6000604051905090565b600067ffffffffffffffff82111561115d5761115c611367565b5b61116682611396565b9050602081019050919050565b600081519050919050565b600082825260208201905092915050565b600081905092915050565b600082825260208201905092915050565b60006111b682611280565b91506111c183611280565b9250827fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff038211156111f6576111f5611338565b5b828201905092915050565b600061120c8261128a565b91506112178361128a565b92508260ff0382111561122d5761122c611338565b5b828201905092915050565b600061124382611260565b9050919050565b60008115159050919050565b6000819050919050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000819050919050565b600060ff82169050919050565b82818337600083830152505050565b60005b838110156112c45780820151818401526020810190506112a9565b838111156112d3576000848401525b50505050565b6112e282611396565b810181811067ffffffffffffffff8211171561130157611300611367565b5b80604052505050565b60006113158261131c565b9050919050565b6000611327826113a7565b9050919050565b6000819050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b6000601f19601f8301169050919050565b60008160601b9050919050565b7f636865636b2e6e6f6e636520746f6f206f6c6400000000000000000000000000600082015250565b7f70617976616c75652073686f756c64206e6f74206578636565642076616c756560008201527f206f6620636865636b2e00000000000000000000000000000000000000000000602082015250565b7f696c6c6567616c20636865636b20736967000000000000000000000000000000600082015250565b7f696c6c6567616c20706179636865636b20736967000000000000000000000000600082015250565b7f6f70657261746f722073686f756c64206265206f776e6572206f66207468697360008201527f20636f6e74726163740000000000000000000000000000000000000000000000602082015250565b7f63616c6c6572207368756f756c6420626520636865636b2e746f416464720000600082015250565b6114ff81611238565b811461150a57600080fd5b50565b61151681611280565b811461152157600080fd5b5056fea26469706673582212201292c55ff7d580447ce2e542cbad725f9d4d06ebe8bff32b1f34565b4a69670764736f6c63430008020033",
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

// GetNodeNonce is a free data retrieval call binding the contract method 0x5f510b7b.
//
// Solidity: function get_node_nonce(address node) view returns(uint256)
func (_Cash *CashCaller) GetNodeNonce(opts *bind.CallOpts, node common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Cash.contract.Call(opts, &out, "get_node_nonce", node)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetNodeNonce is a free data retrieval call binding the contract method 0x5f510b7b.
//
// Solidity: function get_node_nonce(address node) view returns(uint256)
func (_Cash *CashSession) GetNodeNonce(node common.Address) (*big.Int, error) {
	return _Cash.Contract.GetNodeNonce(&_Cash.CallOpts, node)
}

// GetNodeNonce is a free data retrieval call binding the contract method 0x5f510b7b.
//
// Solidity: function get_node_nonce(address node) view returns(uint256)
func (_Cash *CashCallerSession) GetNodeNonce(node common.Address) (*big.Int, error) {
	return _Cash.Contract.GetNodeNonce(&_Cash.CallOpts, node)
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
// Solidity: function nodeNonce(address ) view returns(uint256)
func (_Cash *CashCaller) NodeNonce(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Cash.contract.Call(opts, &out, "nodeNonce", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// NodeNonce is a free data retrieval call binding the contract method 0x1d728a0f.
//
// Solidity: function nodeNonce(address ) view returns(uint256)
func (_Cash *CashSession) NodeNonce(arg0 common.Address) (*big.Int, error) {
	return _Cash.Contract.NodeNonce(&_Cash.CallOpts, arg0)
}

// NodeNonce is a free data retrieval call binding the contract method 0x1d728a0f.
//
// Solidity: function nodeNonce(address ) view returns(uint256)
func (_Cash *CashCallerSession) NodeNonce(arg0 common.Address) (*big.Int, error) {
	return _Cash.Contract.NodeNonce(&_Cash.CallOpts, arg0)
}

// ToBytes1 is a free data retrieval call binding the contract method 0x3cb754d8.
//
// Solidity: function toBytes1(uint256 x) pure returns(bytes b)
func (_Cash *CashCaller) ToBytes1(opts *bind.CallOpts, x *big.Int) ([]byte, error) {
	var out []interface{}
	err := _Cash.contract.Call(opts, &out, "toBytes1", x)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// ToBytes1 is a free data retrieval call binding the contract method 0x3cb754d8.
//
// Solidity: function toBytes1(uint256 x) pure returns(bytes b)
func (_Cash *CashSession) ToBytes1(x *big.Int) ([]byte, error) {
	return _Cash.Contract.ToBytes1(&_Cash.CallOpts, x)
}

// ToBytes1 is a free data retrieval call binding the contract method 0x3cb754d8.
//
// Solidity: function toBytes1(uint256 x) pure returns(bytes b)
func (_Cash *CashCallerSession) ToBytes1(x *big.Int) ([]byte, error) {
	return _Cash.Contract.ToBytes1(&_Cash.CallOpts, x)
}

// ApplyCheck is a paid mutator transaction binding the contract method 0xc8193561.
//
// Solidity: function apply_check(((uint256,address,uint256,address,address,address,address,bytes),uint256,bytes) paycheck) payable returns(bool)
func (_Cash *CashTransactor) ApplyCheck(opts *bind.TransactOpts, paycheck Paycheck) (*types.Transaction, error) {
	return _Cash.contract.Transact(opts, "apply_check", paycheck)
}

// ApplyCheck is a paid mutator transaction binding the contract method 0xc8193561.
//
// Solidity: function apply_check(((uint256,address,uint256,address,address,address,address,bytes),uint256,bytes) paycheck) payable returns(bool)
func (_Cash *CashSession) ApplyCheck(paycheck Paycheck) (*types.Transaction, error) {
	return _Cash.Contract.ApplyCheck(&_Cash.TransactOpts, paycheck)
}

// ApplyCheck is a paid mutator transaction binding the contract method 0xc8193561.
//
// Solidity: function apply_check(((uint256,address,uint256,address,address,address,address,bytes),uint256,bytes) paycheck) payable returns(bool)
func (_Cash *CashTransactorSession) ApplyCheck(paycheck Paycheck) (*types.Transaction, error) {
	return _Cash.Contract.ApplyCheck(&_Cash.TransactOpts, paycheck)
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

// CashShow256Iterator is returned from FilterShow256 and is used to iterate over the raw logs and unpacked data for Show256 events raised by the Cash contract.
type CashShow256Iterator struct {
	Event *CashShow256 // Event containing the contract specifics and raw log

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
func (it *CashShow256Iterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CashShow256)
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
		it.Event = new(CashShow256)
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
func (it *CashShow256Iterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CashShow256Iterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CashShow256 represents a Show256 event raised by the Cash contract.
type CashShow256 struct {
	Arg0 *big.Int
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterShow256 is a free log retrieval operation binding the contract event 0xcbe629aaa9d71fcba91cbad4676731b37fd435513517a19f0012fe8804a10afb.
//
// Solidity: event Show256(uint256 arg0)
func (_Cash *CashFilterer) FilterShow256(opts *bind.FilterOpts) (*CashShow256Iterator, error) {

	logs, sub, err := _Cash.contract.FilterLogs(opts, "Show256")
	if err != nil {
		return nil, err
	}
	return &CashShow256Iterator{contract: _Cash.contract, event: "Show256", logs: logs, sub: sub}, nil
}

// WatchShow256 is a free log subscription operation binding the contract event 0xcbe629aaa9d71fcba91cbad4676731b37fd435513517a19f0012fe8804a10afb.
//
// Solidity: event Show256(uint256 arg0)
func (_Cash *CashFilterer) WatchShow256(opts *bind.WatchOpts, sink chan<- *CashShow256) (event.Subscription, error) {

	logs, sub, err := _Cash.contract.WatchLogs(opts, "Show256")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CashShow256)
				if err := _Cash.contract.UnpackLog(event, "Show256", log); err != nil {
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

// ParseShow256 is a log parse operation binding the contract event 0xcbe629aaa9d71fcba91cbad4676731b37fd435513517a19f0012fe8804a10afb.
//
// Solidity: event Show256(uint256 arg0)
func (_Cash *CashFilterer) ParseShow256(log types.Log) (*CashShow256, error) {
	event := new(CashShow256)
	if err := _Cash.contract.UnpackLog(event, "Show256", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CashShowaddrIterator is returned from FilterShowaddr and is used to iterate over the raw logs and unpacked data for Showaddr events raised by the Cash contract.
type CashShowaddrIterator struct {
	Event *CashShowaddr // Event containing the contract specifics and raw log

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
func (it *CashShowaddrIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CashShowaddr)
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
		it.Event = new(CashShowaddr)
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
func (it *CashShowaddrIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CashShowaddrIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CashShowaddr represents a Showaddr event raised by the Cash contract.
type CashShowaddr struct {
	Arg0 common.Address
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterShowaddr is a free log retrieval operation binding the contract event 0x7876b0f7cc2519df0fdac81db62f5a4a05f5332c05ceb4d667b783eef87094a3.
//
// Solidity: event Showaddr(address arg0)
func (_Cash *CashFilterer) FilterShowaddr(opts *bind.FilterOpts) (*CashShowaddrIterator, error) {

	logs, sub, err := _Cash.contract.FilterLogs(opts, "Showaddr")
	if err != nil {
		return nil, err
	}
	return &CashShowaddrIterator{contract: _Cash.contract, event: "Showaddr", logs: logs, sub: sub}, nil
}

// WatchShowaddr is a free log subscription operation binding the contract event 0x7876b0f7cc2519df0fdac81db62f5a4a05f5332c05ceb4d667b783eef87094a3.
//
// Solidity: event Showaddr(address arg0)
func (_Cash *CashFilterer) WatchShowaddr(opts *bind.WatchOpts, sink chan<- *CashShowaddr) (event.Subscription, error) {

	logs, sub, err := _Cash.contract.WatchLogs(opts, "Showaddr")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CashShowaddr)
				if err := _Cash.contract.UnpackLog(event, "Showaddr", log); err != nil {
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

// ParseShowaddr is a log parse operation binding the contract event 0x7876b0f7cc2519df0fdac81db62f5a4a05f5332c05ceb4d667b783eef87094a3.
//
// Solidity: event Showaddr(address arg0)
func (_Cash *CashFilterer) ParseShowaddr(log types.Log) (*CashShowaddr, error) {
	event := new(CashShowaddr)
	if err := _Cash.contract.UnpackLog(event, "Showaddr", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CashShowbytesIterator is returned from FilterShowbytes and is used to iterate over the raw logs and unpacked data for Showbytes events raised by the Cash contract.
type CashShowbytesIterator struct {
	Event *CashShowbytes // Event containing the contract specifics and raw log

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
func (it *CashShowbytesIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CashShowbytes)
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
		it.Event = new(CashShowbytes)
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
func (it *CashShowbytesIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CashShowbytesIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CashShowbytes represents a Showbytes event raised by the Cash contract.
type CashShowbytes struct {
	Arg0 []byte
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterShowbytes is a free log retrieval operation binding the contract event 0x4daa2d63a96f5e545f8ff5b850799b444e2e8a26e69f59425ba735fa0528a268.
//
// Solidity: event Showbytes(bytes arg0)
func (_Cash *CashFilterer) FilterShowbytes(opts *bind.FilterOpts) (*CashShowbytesIterator, error) {

	logs, sub, err := _Cash.contract.FilterLogs(opts, "Showbytes")
	if err != nil {
		return nil, err
	}
	return &CashShowbytesIterator{contract: _Cash.contract, event: "Showbytes", logs: logs, sub: sub}, nil
}

// WatchShowbytes is a free log subscription operation binding the contract event 0x4daa2d63a96f5e545f8ff5b850799b444e2e8a26e69f59425ba735fa0528a268.
//
// Solidity: event Showbytes(bytes arg0)
func (_Cash *CashFilterer) WatchShowbytes(opts *bind.WatchOpts, sink chan<- *CashShowbytes) (event.Subscription, error) {

	logs, sub, err := _Cash.contract.WatchLogs(opts, "Showbytes")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CashShowbytes)
				if err := _Cash.contract.UnpackLog(event, "Showbytes", log); err != nil {
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

// ParseShowbytes is a log parse operation binding the contract event 0x4daa2d63a96f5e545f8ff5b850799b444e2e8a26e69f59425ba735fa0528a268.
//
// Solidity: event Showbytes(bytes arg0)
func (_Cash *CashFilterer) ParseShowbytes(log types.Log) (*CashShowbytes, error) {
	event := new(CashShowbytes)
	if err := _Cash.contract.UnpackLog(event, "Showbytes", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
