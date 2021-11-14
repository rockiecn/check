package test

import (
	"fmt"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/rockiecn/check/internal/utils"
	"github.com/rockiecn/check/operator"
	"github.com/rockiecn/check/provider"
	"github.com/rockiecn/check/user"
)

var (
	// a local account, with enough money in it
	SenderSk = "503f38a9c967ed597e47fe25643985f032b072db8075426a92110f82df48dfcb"
	token    = common.HexToAddress("0xb213d01542d129806d664248a380db8b12059061")
)

func InitOperator(t *testing.T) *operator.Operator {
	// generate operator
	opSk, err := utils.GenerateSK()
	if err != nil {
		t.Fatal(err)
	}
	op, err := operator.New(opSk)
	if err != nil {
		t.Fatal(err)
	}
	Op := op.(*operator.Operator)

	// send 2 eth to operator
	fmt.Println("-> send some money to operator for deploy contract")

	// send 2 eth to operator
	tx, err := utils.SendCoin(SenderSk, Op.OpAddr, utils.String2BigInt("2000000000000000000"))
	if err != nil {
		t.Fatal(err)
	}
	utils.WaitForMiner(tx)

	fmt.Println("-> deploy contract")
	// operator deploy contract, with 1.8 eth
	tx, ctrAddr, err := op.Deploy(utils.String2BigInt("1800000000000000000"))
	if err != nil {
		t.Fatal(err)
	}
	utils.WaitForMiner(tx)

	// set contract address for operator
	Op.SetCtrAddr(ctrAddr)

	return Op
}

func InitUser(t *testing.T) *user.User {
	// generate user
	usrSk, err := utils.GenerateSK()
	if err != nil {
		t.Fatal(err)
	}
	usr, err := user.New(usrSk)
	if err != nil {
		t.Fatal(err)
	}
	Usr := usr.(*user.User)

	return Usr
}

func InitPro(t *testing.T) *provider.Provider {
	// provider0
	proSk, err := utils.GenerateSK()
	if err != nil {
		t.Fatal(err)
	}
	pro, err := provider.New(proSk)
	if err != nil {
		t.Fatal(err)
	}

	Pro := pro.(*provider.Provider)
	return Pro
}
