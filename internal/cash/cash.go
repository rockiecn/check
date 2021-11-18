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

// BatchCheck is an auto generated low-level Go binding around an user-defined struct.
type BatchCheck struct {
	OpAddr     common.Address
	ToAddr     common.Address
	CtrAddr    common.Address
	TokenAddr  common.Address
	BatchValue *big.Int
	MinNonce   uint64
	MaxNonce   uint64
	BatchSig   []byte
}

// Check is an auto generated low-level Go binding around an user-defined struct.
type Check struct {
	Value     *big.Int
	TokenAddr common.Address
	Nonce     uint64
	FromAddr  common.Address
	ToAddr    common.Address
	OpAddr    common.Address
	CtrAddr   common.Address
	CheckSig  []byte
}

// Paycheck is an auto generated low-level Go binding around an user-defined struct.
type Paycheck struct {
	Check       Check
	PayValue    *big.Int
	PaycheckSig []byte
}

// CashMetaData contains all meta data concerning the Cash contract.
var CashMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"payable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"Paid\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"Pos\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"Received\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"deposit\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getBalance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"node\",\"type\":\"address\"}],\"name\":\"getNonce\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getOwner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"nodeNonce\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"node\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"nonce\",\"type\":\"uint64\"}],\"name\":\"setNonce\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"tokenAddr\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"nonce\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"fromAddr\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"toAddr\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"opAddr\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"ctrAddr\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"checkSig\",\"type\":\"bytes\"}],\"internalType\":\"structCheck\",\"name\":\"check\",\"type\":\"tuple\"},{\"internalType\":\"uint256\",\"name\":\"payValue\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"paycheckSig\",\"type\":\"bytes\"}],\"internalType\":\"structPaycheck\",\"name\":\"paycheck\",\"type\":\"tuple\"}],\"name\":\"withdraw\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"opAddr\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"toAddr\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"ctrAddr\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenAddr\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"batchValue\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"minNonce\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"maxNonce\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"batchSig\",\"type\":\"bytes\"}],\"internalType\":\"structBatchCheck\",\"name\":\"bc\",\"type\":\"tuple\"}],\"name\":\"withdrawBatch\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
	Bin: "0x6080604052336000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550611d90806100536000396000f3fe60806040526004361061007f5760003560e01c8063893d20e81161004e578063893d20e814610199578063d0e30db0146101c4578063d139e245146101ce578063f8e18b57146101fe576100bf565b80631125b1df146100c457806312065fe0146100f45780631d728a0f1461011f5780632d0335ab1461015c576100bf565b366100bf577f88a5966d370b9919b20f3e2c13ff65706f196a4e32cc2c12bf57088f8852587433346040516100b59291906115f7565b60405180910390a1005b600080fd5b6100de60048036038101906100d9919061121b565b610227565b6040516100eb9190611620565b60405180910390f35b34801561010057600080fd5b50610109610725565b6040516101169190611827565b60405180910390f35b34801561012b57600080fd5b5061014660048036038101906101419190611175565b61072d565b6040516101539190611842565b60405180910390f35b34801561016857600080fd5b50610183600480360381019061017e9190611175565b610754565b6040516101909190611842565b60405180910390f35b3480156101a557600080fd5b506101ae6107b1565b6040516101bb91906115dc565b60405180910390f35b6101cc6107da565b005b6101e860048036038101906101e391906111da565b6107dc565b6040516101f59190611620565b60405180910390f35b34801561020a57600080fd5b506102256004803603810190610220919061119e565b610d10565b005b60006001600083600001516080015173ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900467ffffffffffffffff1667ffffffffffffffff1682600001516040015167ffffffffffffffff1610156102e3576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016102da90611707565b60405180910390fd5b81600001516000015182602001511115610332576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161032990611727565b60405180910390fd5b3373ffffffffffffffffffffffffffffffffffffffff1682600001516080015173ffffffffffffffffffffffffffffffffffffffff16146103a8576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161039f90611807565b60405180910390fd5b60008054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16826000015160a0015173ffffffffffffffffffffffffffffffffffffffff161461043e576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610435906117c7565b60405180910390fd5b6000826000015160000151836000015160200151846000015160400151856000015160600151866000015160800151876000015160a00151886000015160c00151604051602001610495979695949392919061155b565b604051602081830303815290604052905060008180519060200120905060006104c682866000015160e00151610d7a565b90508073ffffffffffffffffffffffffffffffffffffffff16856000015160a0015173ffffffffffffffffffffffffffffffffffffffff161461053e576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161053590611747565b60405180910390fd5b6000838660200151604051602001610557929190611533565b60405160208183030381529060405290506000818051906020012090506000610584828960400151610d7a565b90508073ffffffffffffffffffffffffffffffffffffffff1688600001516060015173ffffffffffffffffffffffffffffffffffffffff16146105fc576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016105f390611767565b60405180910390fd5b87600001516080015173ffffffffffffffffffffffffffffffffffffffff166108fc89602001519081150290604051600060405180830381858888f1935050505015801561064e573d6000803e3d6000fd5b507f737c69225d647e5994eab1a6c301bf6d9232beb2759ae1e27a8966b4732bc489886000015160800151896020015160405161068c9291906115f7565b60405180910390a160018860000151604001516106a991906118da565b600160008a600001516080015173ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548167ffffffffffffffff021916908367ffffffffffffffff16021790555060019650505050505050919050565b600047905090565b60016020528060005260406000206000915054906101000a900467ffffffffffffffff1681565b6000600160008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900467ffffffffffffffff169050919050565b60008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905090565b565b600060016000836020015173ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900467ffffffffffffffff1667ffffffffffffffff168260a0015167ffffffffffffffff161015610890576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610887906117e7565b60405180910390fd5b7fd5c119d642a6b81e9029b7425495b7fdf2d4d5a9b0a513fb06e7ed8ab54d59b160016040516108c09190611680565b60405180910390a13073ffffffffffffffffffffffffffffffffffffffff16826040015173ffffffffffffffffffffffffffffffffffffffff161461093a576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610931906117a7565b60405180910390fd5b7fd5c119d642a6b81e9029b7425495b7fdf2d4d5a9b0a513fb06e7ed8ab54d59b1600260405161096a919061169b565b60405180910390a13373ffffffffffffffffffffffffffffffffffffffff16826020015173ffffffffffffffffffffffffffffffffffffffff16146109e4576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016109db90611807565b60405180910390fd5b7fd5c119d642a6b81e9029b7425495b7fdf2d4d5a9b0a513fb06e7ed8ab54d59b16003604051610a1491906116b6565b60405180910390a160008054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16826000015173ffffffffffffffffffffffffffffffffffffffff1614610aae576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610aa5906117c7565b60405180910390fd5b7fd5c119d642a6b81e9029b7425495b7fdf2d4d5a9b0a513fb06e7ed8ab54d59b16004604051610ade91906116d1565b60405180910390a16000826000015183602001518460400151856060015186608001518760a001518860c00151604051602001610b2197969594939291906114b2565b60405160208183030381529060405290506000818051906020012090506000610b4e828660e00151610d7a565b90508073ffffffffffffffffffffffffffffffffffffffff16856000015173ffffffffffffffffffffffffffffffffffffffff1614610bc2576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610bb990611787565b60405180910390fd5b7fd5c119d642a6b81e9029b7425495b7fdf2d4d5a9b0a513fb06e7ed8ab54d59b16005604051610bf291906116ec565b60405180910390a1846020015173ffffffffffffffffffffffffffffffffffffffff166108fc86608001519081150290604051600060405180830381858888f19350505050158015610c48573d6000803e3d6000fd5b507f737c69225d647e5994eab1a6c301bf6d9232beb2759ae1e27a8966b4732bc48985602001518660800151604051610c829291906115f7565b60405180910390a160018560c00151610c9b91906118da565b60016000876020015173ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548167ffffffffffffffff021916908367ffffffffffffffff16021790555060019350505050919050565b80600160008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548167ffffffffffffffff021916908367ffffffffffffffff1602179055505050565b60006041825114610d8e5760009050610e7c565b60008060006020850151925060408501519150606085015160001a90507f7fffffffffffffffffffffffffffffff5d576e7357a4501ddfe92f46681b20a08260001c1115610de25760009350505050610e7c565b601b8160ff161015610dfe57601b81610dfb9190611918565b90505b601b8160ff1614158015610e165750601c8160ff1614155b15610e275760009350505050610e7c565b60018682858560405160008152602001604052604051610e4a949392919061163b565b6020604051602081039080840390855afa158015610e6c573d6000803e3d6000fd5b5050506020604051035193505050505b92915050565b6000610e95610e9084611882565b61185d565b905082815260208101848484011115610ead57600080fd5b610eb8848285611a1c565b509392505050565b600081359050610ecf81611d15565b92915050565b600082601f830112610ee657600080fd5b8135610ef6848260208601610e82565b91505092915050565b60006101008284031215610f1257600080fd5b610f1d61010061185d565b90506000610f2d84828501610ec0565b6000830152506020610f4184828501610ec0565b6020830152506040610f5584828501610ec0565b6040830152506060610f6984828501610ec0565b6060830152506080610f7d8482850161114b565b60808301525060a0610f9184828501611160565b60a08301525060c0610fa584828501611160565b60c08301525060e082013567ffffffffffffffff811115610fc557600080fd5b610fd184828501610ed5565b60e08301525092915050565b60006101008284031215610ff057600080fd5b610ffb61010061185d565b9050600061100b8482850161114b565b600083015250602061101f84828501610ec0565b602083015250604061103384828501611160565b604083015250606061104784828501610ec0565b606083015250608061105b84828501610ec0565b60808301525060a061106f84828501610ec0565b60a08301525060c061108384828501610ec0565b60c08301525060e082013567ffffffffffffffff8111156110a357600080fd5b6110af84828501610ed5565b60e08301525092915050565b6000606082840312156110cd57600080fd5b6110d7606061185d565b9050600082013567ffffffffffffffff8111156110f357600080fd5b6110ff84828501610fdd565b60008301525060206111138482850161114b565b602083015250604082013567ffffffffffffffff81111561113357600080fd5b61113f84828501610ed5565b60408301525092915050565b60008135905061115a81611d2c565b92915050565b60008135905061116f81611d43565b92915050565b60006020828403121561118757600080fd5b600061119584828501610ec0565b91505092915050565b600080604083850312156111b157600080fd5b60006111bf85828601610ec0565b92505060206111d085828601611160565b9150509250929050565b6000602082840312156111ec57600080fd5b600082013567ffffffffffffffff81111561120657600080fd5b61121284828501610eff565b91505092915050565b60006020828403121561122d57600080fd5b600082013567ffffffffffffffff81111561124757600080fd5b611253848285016110bb565b91505092915050565b6112658161194f565b82525050565b61127c6112778261194f565b611a8f565b82525050565b61128b81611961565b82525050565b61129a8161196d565b82525050565b60006112ab826118b3565b6112b581856118be565b93506112c5818560208601611a2b565b80840191505092915050565b6112da816119c2565b82525050565b6112e9816119d4565b82525050565b6112f8816119e6565b82525050565b611307816119f8565b82525050565b61131681611a0a565b82525050565b60006113296013836118c9565b915061133482611b58565b602082019050919050565b600061134c602a836118c9565b915061135782611b81565b604082019050919050565b600061136f6011836118c9565b915061137a82611bd0565b602082019050919050565b60006113926014836118c9565b915061139d82611bf9565b602082019050919050565b60006113b5600e836118c9565b91506113c082611c22565b602082019050919050565b60006113d86016836118c9565b91506113e382611c4b565b602082019050919050565b60006113fb6029836118c9565b915061140682611c74565b604082019050919050565b600061141e6019836118c9565b915061142982611cc3565b602082019050919050565b6000611441601e836118c9565b915061144c82611cec565b602082019050919050565b61146081611997565b82525050565b61147761147282611997565b611ab3565b82525050565b611486816119a1565b82525050565b61149d611498826119a1565b611abd565b82525050565b6114ac816119b5565b82525050565b60006114be828a61126b565b6014820191506114ce828961126b565b6014820191506114de828861126b565b6014820191506114ee828761126b565b6014820191506114fe8286611466565b60208201915061150e828561148c565b60088201915061151e828461148c565b60088201915081905098975050505050505050565b600061153f82856112a0565b915061154b8284611466565b6020820191508190509392505050565b6000611567828a611466565b602082019150611577828961126b565b601482019150611587828861148c565b600882019150611597828761126b565b6014820191506115a7828661126b565b6014820191506115b7828561126b565b6014820191506115c7828461126b565b60148201915081905098975050505050505050565b60006020820190506115f1600083018461125c565b92915050565b600060408201905061160c600083018561125c565b6116196020830184611457565b9392505050565b60006020820190506116356000830184611282565b92915050565b60006080820190506116506000830187611291565b61165d60208301866114a3565b61166a6040830185611291565b6116776060830184611291565b95945050505050565b600060208201905061169560008301846112d1565b92915050565b60006020820190506116b060008301846112e0565b92915050565b60006020820190506116cb60008301846112ef565b92915050565b60006020820190506116e660008301846112fe565b92915050565b6000602082019050611701600083018461130d565b92915050565b600060208201905081810360008301526117208161131c565b9050919050565b600060208201905081810360008301526117408161133f565b9050919050565b6000602082019050818103600083015261176081611362565b9050919050565b6000602082019050818103600083015261178081611385565b9050919050565b600060208201905081810360008301526117a0816113a8565b9050919050565b600060208201905081810360008301526117c0816113cb565b9050919050565b600060208201905081810360008301526117e0816113ee565b9050919050565b6000602082019050818103600083015261180081611411565b9050919050565b6000602082019050818103600083015261182081611434565b9050919050565b600060208201905061183c6000830184611457565b92915050565b6000602082019050611857600083018461147d565b92915050565b6000611867611878565b90506118738282611a5e565b919050565b6000604051905090565b600067ffffffffffffffff82111561189d5761189c611afe565b5b6118a682611b2d565b9050602081019050919050565b600081519050919050565b600081905092915050565b600082825260208201905092915050565b60006118e5826119a1565b91506118f0836119a1565b92508267ffffffffffffffff0382111561190d5761190c611acf565b5b828201905092915050565b6000611923826119b5565b915061192e836119b5565b92508260ff0382111561194457611943611acf565b5b828201905092915050565b600061195a82611977565b9050919050565b60008115159050919050565b6000819050919050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000819050919050565b600067ffffffffffffffff82169050919050565b600060ff82169050919050565b60006119cd82611997565b9050919050565b60006119df82611997565b9050919050565b60006119f182611997565b9050919050565b6000611a0382611997565b9050919050565b6000611a1582611997565b9050919050565b82818337600083830152505050565b60005b83811015611a49578082015181840152602081019050611a2e565b83811115611a58576000848401525b50505050565b611a6782611b2d565b810181811067ffffffffffffffff82111715611a8657611a85611afe565b5b80604052505050565b6000611a9a82611aa1565b9050919050565b6000611aac82611b4b565b9050919050565b6000819050919050565b6000611ac882611b3e565b9050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b6000601f19601f8301169050919050565b60008160c01b9050919050565b60008160601b9050919050565b7f636865636b2e6e6f6e636520746f6f206f6c6400000000000000000000000000600082015250565b7f70617976616c75652073686f756c64206e6f74206578636565642076616c756560008201527f206f6620636865636b2e00000000000000000000000000000000000000000000602082015250565b7f696c6c6567616c20636865636b20736967000000000000000000000000000000600082015250565b7f696c6c6567616c20706179636865636b20736967000000000000000000000000600082015250565b7f696c6c6567616c20626320736967000000000000000000000000000000000000600082015250565b7f636f6e74726163742061646472657373206572726f7200000000000000000000600082015250565b7f6f70657261746f722073686f756c64206265206f776e6572206f66207468697360008201527f20636f6e74726163740000000000000000000000000000000000000000000000602082015250565b7f626174636820636865636b206e6f6e636520746f6f206f6c6400000000000000600082015250565b7f63616c6c6572207368756f756c6420626520636865636b2e746f416464720000600082015250565b611d1e8161194f565b8114611d2957600080fd5b50565b611d3581611997565b8114611d4057600080fd5b50565b611d4c816119a1565b8114611d5757600080fd5b5056fea2646970667358221220455041c3c7f0db4fd8d0af2160d60158135d02b6446db1e69ad8ff99c6f4b9ec64736f6c63430008020033",
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

// WithdrawBatch is a paid mutator transaction binding the contract method 0xd139e245.
//
// Solidity: function withdrawBatch((address,address,address,address,uint256,uint64,uint64,bytes) bc) payable returns(bool)
func (_Cash *CashTransactor) WithdrawBatch(opts *bind.TransactOpts, bc BatchCheck) (*types.Transaction, error) {
	return _Cash.contract.Transact(opts, "withdrawBatch", bc)
}

// WithdrawBatch is a paid mutator transaction binding the contract method 0xd139e245.
//
// Solidity: function withdrawBatch((address,address,address,address,uint256,uint64,uint64,bytes) bc) payable returns(bool)
func (_Cash *CashSession) WithdrawBatch(bc BatchCheck) (*types.Transaction, error) {
	return _Cash.Contract.WithdrawBatch(&_Cash.TransactOpts, bc)
}

// WithdrawBatch is a paid mutator transaction binding the contract method 0xd139e245.
//
// Solidity: function withdrawBatch((address,address,address,address,uint256,uint64,uint64,bytes) bc) payable returns(bool)
func (_Cash *CashTransactorSession) WithdrawBatch(bc BatchCheck) (*types.Transaction, error) {
	return _Cash.Contract.WithdrawBatch(&_Cash.TransactOpts, bc)
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

// CashPosIterator is returned from FilterPos and is used to iterate over the raw logs and unpacked data for Pos events raised by the Cash contract.
type CashPosIterator struct {
	Event *CashPos // Event containing the contract specifics and raw log

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
func (it *CashPosIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CashPos)
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
		it.Event = new(CashPos)
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
func (it *CashPosIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CashPosIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CashPos represents a Pos event raised by the Cash contract.
type CashPos struct {
	Arg0 *big.Int
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterPos is a free log retrieval operation binding the contract event 0xd5c119d642a6b81e9029b7425495b7fdf2d4d5a9b0a513fb06e7ed8ab54d59b1.
//
// Solidity: event Pos(uint256 arg0)
func (_Cash *CashFilterer) FilterPos(opts *bind.FilterOpts) (*CashPosIterator, error) {

	logs, sub, err := _Cash.contract.FilterLogs(opts, "Pos")
	if err != nil {
		return nil, err
	}
	return &CashPosIterator{contract: _Cash.contract, event: "Pos", logs: logs, sub: sub}, nil
}

// WatchPos is a free log subscription operation binding the contract event 0xd5c119d642a6b81e9029b7425495b7fdf2d4d5a9b0a513fb06e7ed8ab54d59b1.
//
// Solidity: event Pos(uint256 arg0)
func (_Cash *CashFilterer) WatchPos(opts *bind.WatchOpts, sink chan<- *CashPos) (event.Subscription, error) {

	logs, sub, err := _Cash.contract.WatchLogs(opts, "Pos")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CashPos)
				if err := _Cash.contract.UnpackLog(event, "Pos", log); err != nil {
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

// ParsePos is a log parse operation binding the contract event 0xd5c119d642a6b81e9029b7425495b7fdf2d4d5a9b0a513fb06e7ed8ab54d59b1.
//
// Solidity: event Pos(uint256 arg0)
func (_Cash *CashFilterer) ParsePos(log types.Log) (*CashPos, error) {
	event := new(CashPos)
	if err := _Cash.contract.UnpackLog(event, "Pos", log); err != nil {
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
