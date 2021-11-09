package test

import (
	"fmt"
	"math/big"
	"testing"
	"time"

	"github.com/rockiecn/check/internal/check"
	"github.com/rockiecn/check/internal/order"
	"github.com/rockiecn/check/internal/utils"
	"github.com/rockiecn/check/operator"
	"github.com/rockiecn/check/provider"
	"github.com/rockiecn/check/user"
)

// a single provider involved
// a single check is used
// multi pay actions by user, and 1 withdraw action by provider
func TestSingleProMultiCheck(t *testing.T) {

	fmt.Println("<< Initialize >>")
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

	// send 2 eth to operator
	tx, err := utils.SendCoin(SenderSk, opAddr, utils.String2BigInt("2000000000000000000"))
	if err != nil {
		t.Error(err)
	}
	utils.WaitForMiner(tx)

	fmt.Println("now deploy contract:")
	// operator deploy contract, with 1.8 eth
	tx, ctrAddr, err := op.Deploy(utils.String2BigInt("1800000000000000000"))
	if err != nil {
		t.Error(err)
	}
	utils.WaitForMiner(tx)

	// set contract address for operator
	op.SetCtrAddr(ctrAddr)

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

	fmt.Println("<< Create Provider >>")
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

	fmt.Println("<< New Order, ID:0 >>")

	// create an order
	odr0 := order.NewOdr(0,
		token,
		usrAddr,
		proAddr,
		utils.String2BigInt("500000000000000000"), // order value: 0.5 eth
		time.Now(),
		"jack",
		"123123123",
		"asdf@asdf.com",
		0,
		nil,
	)
	if odr0 == nil {
		t.Error("create order failed")
	}

	// operator store order into pool
	err = op.StoreOrder(odr0)
	if err != nil {
		t.Error(err)
	}
	// operator get an order by id
	odr0, err = op.GetOrder(0)
	if err != nil {
		t.Error(err)
	}
	if odr0 == nil {
		t.Error("get order failed")
	}

	fmt.Println("<< Order to Check >>")
	// operator create a check from order
	opChk0, err := op.CreateCheck(0)
	if err != nil {
		t.Error(err)
	}

	// simulate user receive check from operator
	usrChk0 := new(check.Check)
	*usrChk0 = *opChk0
	// user store check into pool
	err = usr.StoreCheck(usrChk0)
	if err != nil {
		t.Error(err)
	}

	fmt.Println("<< New Order, ID:1 >>")

	// create an order
	odr1 := order.NewOdr(1,
		token,
		usrAddr,
		proAddr,
		utils.String2BigInt("500000000000000000"), // order value: 0.5 eth
		time.Now(),
		"jack",
		"123123123",
		"asdf@asdf.com",
		0,
		nil,
	)
	if odr1 == nil {
		t.Error("create order failed")
	}

	// operator store order into pool
	err = op.StoreOrder(odr1)
	if err != nil {
		t.Error(err)
	}
	// operator get an order by id
	odr1, err = op.GetOrder(1)
	if err != nil {
		t.Error(err)
	}
	if odr1 == nil {
		t.Error("get order failed")
	}

	fmt.Println("<< Order to Check >>")
	// operator create a check from order
	opChk1, err := op.CreateCheck(1)
	if err != nil {
		t.Error(err)
	}

	// simulate user receive check from operator
	usrChk1 := new(check.Check)
	*usrChk1 = *opChk1
	// user store check into pool
	err = usr.StoreCheck(usrChk1)
	if err != nil {
		t.Error(err)
	}

	fmt.Println("<< New Order, ID:2 >>")

	// create an order
	odr2 := order.NewOdr(2,
		token,
		usrAddr,
		proAddr,
		utils.String2BigInt("500000000000000000"), // order value: 0.5 eth
		time.Now(),
		"jack",
		"123123123",
		"asdf@asdf.com",
		0,
		nil,
	)
	if odr2 == nil {
		t.Error("create order failed")
	}

	// operator store order into pool
	err = op.StoreOrder(odr2)
	if err != nil {
		t.Error(err)
	}
	// operator get an order by id
	odr2, err = op.GetOrder(2)
	if err != nil {
		t.Error(err)
	}
	if odr2 == nil {
		t.Error("get order failed")
	}

	fmt.Println("<< Order to Check >>")
	// operator create a check from order
	opChk2, err := op.CreateCheck(2)
	if err != nil {
		t.Error(err)
	}

	// simulate user receive check from operator
	usrChk2 := new(check.Check)
	*usrChk2 = *opChk2
	// user store check into pool
	err = usr.StoreCheck(usrChk2)
	if err != nil {
		t.Error(err)
	}

	fmt.Println("<< Pay >>")
	fmt.Println("<< pay 0.1 eth: nonce 0 should be enough >>")
	// user generate a paycheck for paying to provider
	// store new paycheck into user pool
	// pay: 0.1 eth
	userPC, err := usr.Pay(proAddr, utils.String2BigInt(("100000000000000000")))
	if err != nil {
		t.Error(err)
	}

	if userPC == nil || userPC.Check.Nonce != 0 {
		t.Fatal("test pay 0.1 failed")
	}
	fmt.Println("payer nonce: ", userPC.Check.Nonce)

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

	// provider store a paycheck into pool
	err = pro.StorePaycheck(proPC)
	if err != nil {
		t.Error("store paycheck error")
	}

	fmt.Println("<< pay 0.2 eth: nonce 0 should be enough >>")
	// pay: 0.2 eth
	userPC, err = usr.Pay(proAddr, utils.String2BigInt(("200000000000000000")))
	if err != nil {
		t.Error(err)
	}
	if userPC == nil || userPC.Check.Nonce != 0 {
		t.Fatal("test pay 0.2 failed")
	}
	fmt.Println("payer nonce: ", userPC.Check.Nonce)

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

	// provider store a paycheck into pool
	err = pro.StorePaycheck(proPC)
	if err != nil {
		t.Error(err)
	}

	fmt.Println("<< pay 0.4 eth: nonce 0 is not enough, nonce 1 should be used >>")
	// pay: 0.4 eth
	userPC, err = usr.Pay(proAddr, utils.String2BigInt(("400000000000000000")))
	if err != nil {
		t.Error(err)
	}
	if userPC == nil || userPC.Check.Nonce != 1 {
		t.Fatal("test pay 0.4 failed")
	}
	fmt.Println("payer nonce: ", userPC.Check.Nonce)

	// simulate provider receive paycheck from user
	proPC = new(check.Paycheck)
	*proPC = *userPC

	// provider verify received paycheck
	// datavalue: 0.4 eth
	ok, err = pro.Verify(proPC, utils.String2BigInt(("400000000000000000")))
	if err != nil {
		t.Error(err)
	}
	if !ok {
		t.Error("provider verify paycheck failed")
	}

	// provider store a paycheck into pool
	err = pro.StorePaycheck(proPC)
	if err != nil {
		t.Error(err)
	}

	fmt.Println("<< pay 0.2 eth: nonce 0 should be enough again(0.2 remained) >>")
	// user generate a paycheck for paying to provider
	// store new paycheck into user pool
	// pay: 0.3 eth
	userPC, err = usr.Pay(proAddr, utils.String2BigInt(("200000000000000000")))
	if err != nil {
		t.Error(err)
	}
	if userPC == nil || userPC.Check.Nonce != 0 {
		t.Fatal("test pay 0.2 failed")
	}
	fmt.Println("payer nonce: ", userPC.Check.Nonce)

	// simulate provider receive paycheck from user
	proPC = new(check.Paycheck)
	*proPC = *userPC

	// provider verify received paycheck
	// datavalue: 0.3 eth
	ok, err = pro.Verify(proPC, utils.String2BigInt(("200000000000000000")))
	if err != nil {
		t.Error(err)
	}
	if !ok {
		t.Fatal("provider verify paycheck failed")
	}

	// provider store a paycheck into pool
	err = pro.StorePaycheck(proPC)
	if err != nil {
		t.Error("store paycheck error")
	}

	fmt.Println("<< pay 0.6 eth: no check is enough(0.5 max), nil paycheck expected >>")
	// pay: 0.6 eth
	userPC, err = usr.Pay(proAddr, utils.String2BigInt(("600000000000000000")))
	if err != nil {
		fmt.Println(err)
	}
	if userPC != nil {
		t.Fatal("nil paycheck should expected")
	}

	fmt.Println("<< Withdraw >>")
	fmt.Println("<< first withdraw: nonce 0 expected >>")
	// nonce 0 expected
	got, err := pro.GetNextPayable()
	if err != nil {
		t.Fatal(err)
	}
	if got != nil {
		fmt.Println("nonce: ", got.Check.Nonce)
	}
	if got == nil || got.Check.Nonce != 0 {
		t.Fatal("nonce 0 expected")
	}

	// send 1 eth to provider
	fmt.Println("now send 1 eth to provider")
	tx, err = utils.SendCoin(SenderSk, proAddr, utils.String2BigInt("1000000000000000000"))
	if err != nil {
		t.Error(err)
	}
	utils.WaitForMiner(tx)

	// query provider balance before withdraw
	b1, err := pro.QueryBalance()
	if err != nil {
		t.Error(err)
	}

	n, err := op.GetNonce(got.Check.ToAddr)
	if err != nil {
		t.Error(err)
	}
	fmt.Println("nonce in contract:", n)

	fmt.Println("now withdraw:")
	tx, err = pro.Withdraw(got)
	if err != nil {
		t.Error(err)
	}
	utils.WaitForMiner(tx)

	fmt.Println("tx hash:", tx.Hash())

	gasUsed, err := utils.GetGasUsed(tx)
	if err != nil {
		t.Error(err)
	}

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
	fmt.Println("payvalue:", got.PayValue)
	fmt.Println("gasUsed:", gasUsed)
	fmt.Println("balance after withdraw b2:", b2)
	// need add used gas for withdraw tx
	delta := new(big.Int).Sub(b2, b1)
	delta.Add(delta, gasUsed)

	if delta.Cmp(got.PayValue) != 0 {
		t.Error("withdrawed money not equal payvalue")
	} else {
		fmt.Println("withdrawed money equal payvalue")
	}

	fmt.Println("<< second withdraw: nonce 1 expected >>")
	// nonce 1 expected
	got, err = pro.GetNextPayable()
	if err != nil {
		t.Fatal(err)
	}
	if got != nil {
		fmt.Println("nonce: ", got.Check.Nonce)
	} else {
		fmt.Println("nil paycheck")
	}

	if got == nil || got.Check.Nonce != 1 {
		t.Fatal("nonce 1 expected")
	}

	// send 1 eth to provider
	fmt.Println("now send 1 eth to provider")
	tx, err = utils.SendCoin(SenderSk, proAddr, utils.String2BigInt("1000000000000000000"))
	if err != nil {
		t.Error(err)
	}
	utils.WaitForMiner(tx)

	// query provider balance before withdraw
	b1, err = pro.QueryBalance()
	if err != nil {
		t.Error(err)
	}

	n, err = op.GetNonce(got.Check.ToAddr)
	if err != nil {
		t.Error(err)
	}
	fmt.Println("nonce in contract:", n)

	fmt.Println("now withdraw:")
	tx, err = pro.Withdraw(got)
	if err != nil {
		t.Error(err)
	}
	utils.WaitForMiner(tx)

	fmt.Println("tx hash:", tx.Hash())

	gasUsed, err = utils.GetGasUsed(tx)
	if err != nil {
		t.Error(err)
	}

	// wait for a block be mined
	err = utils.WaitForMiner(tx)
	if err != nil {
		t.Error(err)
	}
	// query provider balance after withdraw
	b2, err = pro.QueryBalance()
	if err != nil {
		t.Error(err)
	}
	fmt.Println("balance before withdraw b1:", b1)
	fmt.Println("payvalue:", got.PayValue)
	fmt.Println("gasUsed:", gasUsed)
	fmt.Println("balance after withdraw b2:", b2)
	// need add used gas for withdraw tx
	delta = new(big.Int).Sub(b2, b1)
	delta.Add(delta, gasUsed)

	if delta.Cmp(got.PayValue) != 0 {
		t.Error("withdrawed money not equal payvalue")
	} else {
		fmt.Println("withdrawed money equal payvalue")
	}

}
