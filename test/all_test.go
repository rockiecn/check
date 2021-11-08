package test

import (
	"fmt"
	"math/big"
	"testing"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/rockiecn/check/internal/check"
	"github.com/rockiecn/check/internal/order"
	"github.com/rockiecn/check/internal/utils"
	"github.com/rockiecn/check/operator"
	"github.com/rockiecn/check/provider"
	"github.com/rockiecn/check/user"
)

func TestAll(t *testing.T) {

	// generate operator
	opSk, err := utils.GenerateSK()
	if err != nil {
		t.Error(err)
	}
	opAddr, err := utils.SkToAddr(opSk)
	if err != nil {
		t.Error(err)
	}
	op, err := operator.New(opSk)
	if err != nil {
		t.Error(err)
	}

	// send 2 eth to operator
	fmt.Println("send some money to operator for deploy contract")

	// sender: a local account, with enough money in it
	senderSk := "503f38a9c967ed597e47fe25643985f032b072db8075426a92110f82df48dfcb"
	tx, err := utils.SendCoin(senderSk, opAddr, utils.String2BigInt("2000000000000000000"))
	if err != nil {
		t.Error(err)
	}
	utils.WaitForMiner(tx)

	fmt.Println("now deploy contract:")
	// operator deploy contract, with 1 eth
	tx, ctrAddr, err := op.Deploy(utils.String2BigInt("1000000000000000000"))
	if err != nil {
		t.Error(err)
	}
	utils.WaitForMiner(tx)

	// set contract address for operator
	op.SetCtrAddr(ctrAddr)

	fmt.Println("1")
	// generate user
	usrSk, err := utils.GenerateSK()
	if err != nil {
		t.Error(err)
	}
	usrAddr, err := utils.SkToAddr(usrSk)
	if err != nil {
		t.Error(err)
	}
	usr, err := user.New(usrSk)
	if err != nil {
		t.Error(err)
	}

	fmt.Println("2")
	// generate provider
	proSk, err := utils.GenerateSK()
	if err != nil {
		t.Error(err)
	}
	proAddr, err := utils.SkToAddr(proSk)
	if err != nil {
		t.Error(err)
	}
	pro, err := provider.New(proSk)
	if err != nil {
		t.Error(err)
	}

	fmt.Println("3")

	// create an order
	token := common.HexToAddress("0xb213d01542d129806d664248a380db8b12059061")
	odr := order.NewOdr(0,
		token,
		usrAddr,
		proAddr,
		utils.String2BigInt("800000000000000000"), // value: 0.8 eth
		time.Now(),
		"jack",
		"123123123",
		"asdf@asdf.com",
		0,
		nil,
	)
	if odr == nil {
		t.Error("create order failed")
	}

	fmt.Println("4")
	// operator store order into pool
	err = op.StoreOrder(odr)
	if err != nil {
		t.Error(err)
	}
	// operator get an order by id
	odr, err = op.QueryOrder(0)
	if err != nil {
		t.Error(err)
	}
	if odr == nil {
		t.Error("query order error")
	}

	fmt.Println("5")
	// operator generate a check from order
	chk, err := op.NewCheck(0)
	if err != nil {
		t.Error(err)
	}

	fmt.Println("6")
	// user store a check into pool
	err = usr.StoreCheck(chk)
	if err != nil {
		t.Error(err)
	}

	fmt.Println("7 user generate paycheck No.1 for paying")
	// user generate a paycheck for paying to provider
	// payvalue: 0.1 eth
	userPC, err := usr.NewPaycheck(proAddr, utils.String2BigInt(("100000000000000000")))
	if err != nil {
		t.Error(err)
	}

	fmt.Println("8")

	// simulate provider receive paycheck from user
	proPC := new(check.Paycheck)
	*proPC = *userPC

	// provider verify received paycheck
	// datavalue: 0.1 eth
	ok, err := pro.Verify(proPC, utils.String2BigInt(("100000000000000000")))
	if err != nil {
		t.Error(err)
	}
	if !ok {
		t.Error("provider verify paycheck failed")
	}

	fmt.Println("9")

	// provider store a paycheck into pool
	err = pro.StorePaycheck(proPC)
	if err != nil {
		t.Error("store paycheck error")
	}

	fmt.Println("9.1 --> user generate paycheck No.2 for paying")
	// user generate a paycheck for paying to provider
	// store new paycheck into user pool
	// dataValue: 0.2 eth
	userPC, err = usr.NewPaycheck(proAddr, utils.String2BigInt(("200000000000000000")))
	if err != nil {
		t.Error(err)
	}

	fmt.Println("9.2")

	// simulate provider receive paycheck from user
	proPC = new(check.Paycheck)
	*proPC = *userPC

	// provider verify received paycheck
	// datavalue: 0.2 eth
	ok, err = pro.Verify(proPC, utils.String2BigInt(("200000000000000000")))
	if err != nil {
		t.Error(err)
	}
	if !ok {
		t.Error("provider verify paycheck failed")
	}

	fmt.Println("9.3")
	// provider store a paycheck into pool
	err = pro.StorePaycheck(proPC)
	if err != nil {
		t.Error(err)
	}

	fmt.Println("10 --> Provider withdraw with paycheck")
	// provider get next payable paycheck from pool
	npchk, err := pro.GetNextPayable()
	if err != nil {
		t.Error(err)
	}

	// send 1 eth to provider
	fmt.Println("now send 1 eth to provider")
	tx, err = utils.SendCoin(senderSk, proAddr, utils.String2BigInt("1000000000000000000"))
	if err != nil {
		t.Error(err)
	}
	utils.WaitForMiner(tx)

	fmt.Println("11")
	// query provider balance before withdraw
	b1, err := pro.QueryBalance()
	if err != nil {
		t.Error(err)
	}

	// call contract to withdraw a paycheck
	// fmt.Println("paycheck info:")
	// fmt.Println("contract:", npchk.Check.ContractAddr)
	// fmt.Println("from:", npchk.Check.FromAddr)
	// fmt.Println("to:", npchk.Check.ToAddr)
	// fmt.Println("operator:", npchk.Check.OpAddr)
	// fmt.Println("nonce:", npchk.Check.Nonce)
	// fmt.Println("token:", npchk.Check.TokenAddr)
	// fmt.Println("value:", npchk.Check.Value)
	// fmt.Printf("check sig:%x\n", npchk.Check.CheckSig)
	// fmt.Println("payvalue:", npchk.PayValue)
	// fmt.Printf("paycheck sig:%x\n", npchk.PaycheckSig)

	n, err := op.GetNonce(npchk.Check.ToAddr)
	if err != nil {
		t.Error(err)
	}
	fmt.Println("nonce in contract:", n)

	fmt.Println("now withdraw:")
	tx, err = pro.Withdraw(npchk)
	if err != nil {
		t.Error(err)
	}
	utils.WaitForMiner(tx)

	fmt.Println("tx hash:", tx.Hash())

	gasUsed, err := utils.GetGasUsed(tx)
	if err != nil {
		t.Error(err)
	}

	fmt.Println("12")
	// wait for a block be mined
	err = utils.WaitForMiner(tx)
	if err != nil {
		t.Error(err)
	}
	// query provider balance after withdraw
	b2, err := pro.QueryBalance()
	if err != nil {
		t.Error(err)
	}
	fmt.Println("balance before withdraw b1:", b1)
	fmt.Println("balance after withdraw b2:", b2)
	fmt.Println("gasUsed:", gasUsed)
	fmt.Println("payvalue in paycheck:", npchk.PayValue)
	fmt.Println("b2 = b1 + payvalue - gasUsed")
	// need add used gas for withdraw tx
	delta := new(big.Int).Sub(b2, b1)
	delta.Add(delta, gasUsed)

	if delta.Cmp(npchk.PayValue) != 0 {
		t.Error("withdrawed money not equal payvalue")
	}
}
