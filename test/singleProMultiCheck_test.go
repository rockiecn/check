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
// multi checks is used(with the same provider)
// multi pay actions by user, and 2 withdraw actions by provider
func TestSingleProMultiCheck(t *testing.T) {

	fmt.Println("<< Initialize >>")
	fmt.Println("-> New Operator")
	// generate operator
	opSk, err := utils.GenerateSK()
	if err != nil {
		t.Fatal(err)
	}
	opAddr, err := utils.SkToAddr(opSk)
	if err != nil {
		t.Fatal(err)
	}
	op, err := operator.New(opSk)
	if err != nil {
		t.Fatal(err)
	}

	// send 2 eth to operator
	fmt.Println("-> send some money to operator for deploy contract")

	// send 2 eth to operator
	tx, err := utils.SendCoin(SenderSk, opAddr, utils.String2BigInt("2000000000000000000"))
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
	op.SetCtrAddr(ctrAddr)

	// generate user
	usrSk, err := utils.GenerateSK()
	if err != nil {
		t.Fatal(err)
	}
	usrAddr, err := utils.SkToAddr(usrSk)
	if err != nil {
		t.Fatal(err)
	}
	usr, err := user.New(usrSk)
	if err != nil {
		t.Fatal(err)
	}

	fmt.Println("-> New Provider")
	// generate provider
	proSk, err := utils.GenerateSK()
	if err != nil {
		t.Fatal(err)
	}
	proAddr, err := utils.SkToAddr(proSk)
	if err != nil {
		t.Fatal(err)
	}
	pro, err := provider.New(proSk)
	if err != nil {
		t.Fatal(err)
	}

	fmt.Println("-> Order ID=0, value=0.5")

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
		t.Fatal("create order 0 failed")
	}

	fmt.Println("-> Operator Store Order")
	// operator store order into pool
	err = op.StoreOrder(odr0)
	if err != nil {
		t.Fatal(err)
	}
	// operator get an order by id
	odr0, err = op.GetOrder(0)
	if err != nil {
		t.Fatal(err)
	}
	if odr0 == nil {
		t.Fatal("get order failed")
	}

	fmt.Println("-> Order to Check")
	// operator create a check from order
	opChk0, err := op.CreateCheck(0)
	if err != nil {
		t.Fatal(err)
	}

	// simulate user receive check from operator
	usrChk0 := new(check.Check)
	*usrChk0 = *opChk0
	// user store check into pool
	err = usr.StoreCheck(usrChk0)
	if err != nil {
		t.Fatal(err)
	}

	fmt.Println("-> Order ID=1, value=0.5")

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
		t.Fatal("create order 1 failed")
	}

	// operator store order into pool
	err = op.StoreOrder(odr1)
	if err != nil {
		t.Fatal(err)
	}
	// operator get an order by id
	odr1, err = op.GetOrder(1)
	if err != nil {
		t.Fatal(err)
	}
	if odr1 == nil {
		t.Fatal("get order 1 failed")
	}

	fmt.Println("-> Order to Check")
	// operator create a check from order
	opChk1, err := op.CreateCheck(1)
	if err != nil {
		t.Fatal(err)
	}

	// simulate user receive check from operator
	usrChk1 := new(check.Check)
	*usrChk1 = *opChk1
	// user store check into pool
	err = usr.StoreCheck(usrChk1)
	if err != nil {
		t.Fatal(err)
	}

	fmt.Println("-> Order ID=2, value=0.5")

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
		t.Fatal("create order 2 failed")
	}

	// operator store order into pool
	err = op.StoreOrder(odr2)
	if err != nil {
		t.Fatal(err)
	}
	// operator get an order by id
	odr2, err = op.GetOrder(2)
	if err != nil {
		t.Fatal(err)
	}
	if odr2 == nil {
		t.Fatal("get order 2 failed")
	}

	fmt.Println("-> Order to Check")
	// operator create a check from order
	opChk2, err := op.CreateCheck(2)
	if err != nil {
		t.Fatal(err)
	}

	// simulate user receive check from operator
	usrChk2 := new(check.Check)
	*usrChk2 = *opChk2
	// user store check into pool
	err = usr.StoreCheck(usrChk2)
	if err != nil {
		t.Fatal(err)
	}

	fmt.Println("<< Pay >>")
	fmt.Println("-> pay 0.1 eth: check with nonce 0 should be enough")
	// user generate a paycheck for paying to provider
	// store new paycheck into user pool
	// pay: 0.1 eth
	userPC, err := usr.Pay(proAddr, utils.String2BigInt(("100000000000000000")))
	if err != nil {
		t.Fatal(err)
	}

	if userPC == nil || userPC.Check.Nonce != 0 {
		t.Fatal("test pay 0.1 failed")
	} else {
		fmt.Println("OK- payer nonce: ", userPC.Check.Nonce)
	}

	// simulate provider receive paycheck from user
	proPC := new(check.Paycheck)
	*proPC = *userPC

	// provider verify received paycheck
	// datavalue: 0.1 eth
	ok, err := pro.Verify(proPC, utils.String2BigInt(("100000000000000000")))
	if err != nil {
		t.Fatal(err)
	}
	if !ok {
		t.Fatal("provider verify paycheck failed")
	}

	// provider store a paycheck into pool
	err = pro.StorePaycheck(proPC)
	if err != nil {
		t.Fatal("store paycheck error")
	}

	fmt.Println("-> pay 0.2 eth: check with nonce 0 should be enough")
	// pay: 0.2 eth
	userPC, err = usr.Pay(proAddr, utils.String2BigInt(("200000000000000000")))
	if err != nil {
		t.Fatal(err)
	}
	if userPC == nil || userPC.Check.Nonce != 0 {
		t.Fatal("test pay 0.2 failed")
	} else {
		fmt.Println("OK- payer nonce: ", userPC.Check.Nonce)
	}

	// simulate provider receive paycheck from user
	proPC = new(check.Paycheck)
	*proPC = *userPC

	// provider verify received paycheck
	// datavalue: 0.2 eth
	ok, err = pro.Verify(proPC, utils.String2BigInt(("200000000000000000")))
	if err != nil {
		t.Fatal(err)
	}
	if !ok {
		t.Fatal("provider verify paycheck failed")
	}

	// provider store a paycheck into pool
	err = pro.StorePaycheck(proPC)
	if err != nil {
		t.Fatal(err)
	}

	fmt.Println("-> pay 0.4 eth: nonce 0 is not enough, nonce 1 should be used")
	// pay: 0.4 eth
	userPC, err = usr.Pay(proAddr, utils.String2BigInt(("400000000000000000")))
	if err != nil {
		t.Fatal(err)
	}
	if userPC == nil || userPC.Check.Nonce != 1 {
		t.Fatal("test pay 0.4 failed")
	} else {
		fmt.Println("OK- payer nonce: ", userPC.Check.Nonce)
	}

	// simulate provider receive paycheck from user
	proPC = new(check.Paycheck)
	*proPC = *userPC

	// provider verify received paycheck
	// datavalue: 0.4 eth
	ok, err = pro.Verify(proPC, utils.String2BigInt(("400000000000000000")))
	if err != nil {
		t.Fatal(err)
	}
	if !ok {
		t.Fatal("provider verify paycheck failed")
	}

	// provider store a paycheck into pool
	err = pro.StorePaycheck(proPC)
	if err != nil {
		t.Fatal(err)
	}

	fmt.Println("-> pay 0.2 eth: nonce 0 should be enough again(0.2 remained)")
	// user generate a paycheck for paying to provider
	// store new paycheck into user pool
	// pay: 0.2 eth
	userPC, err = usr.Pay(proAddr, utils.String2BigInt(("200000000000000000")))
	if err != nil {
		t.Fatal(err)
	}
	if userPC == nil || userPC.Check.Nonce != 0 {
		t.Fatal("test pay 0.2 failed")
	} else {
		fmt.Println("OK- payer nonce: ", userPC.Check.Nonce)
	}

	// simulate provider receive paycheck from user
	proPC = new(check.Paycheck)
	*proPC = *userPC

	// provider verify received paycheck
	// datavalue: 0.2 eth
	ok, err = pro.Verify(proPC, utils.String2BigInt(("200000000000000000")))
	if err != nil {
		t.Fatal(err)
	}
	if !ok {
		t.Fatal("provider verify paycheck failed")
	}

	// provider store a paycheck into pool
	err = pro.StorePaycheck(proPC)
	if err != nil {
		t.Error("store paycheck error")
	}

	fmt.Println("-> pay 0.6 eth: no check is enough(0.5 max), nil paycheck expected")
	// pay: 0.6 eth
	userPC, err = usr.Pay(proAddr, utils.String2BigInt(("600000000000000000")))
	if err != nil {
		fmt.Println(err)
	}
	if userPC != nil {
		t.Fatal("nil paycheck should expected")
	} else {
		fmt.Println("OK- nil check returned")
	}

	fmt.Println("<< Withdraw >>")
	fmt.Println("-> first withdraw: nonce 0 expected")
	// nonce 0 expected
	got, err := pro.GetNextPayable()
	if err != nil {
		t.Fatal(err)
	}
	if got == nil {
		t.Fatal("nil paycheck got")
	}
	if got.Check.Nonce != 0 {
		t.Fatalf("nonce=%v, nonce 0 expected\n", got.Check.Nonce)
	} else {
		fmt.Println("OK- withdrawed nonce: ", got.Check.Nonce)
	}

	// send 1 eth to provider
	fmt.Println("-> send 1 eth to provider for withdraw")
	tx, err = utils.SendCoin(SenderSk, proAddr, utils.String2BigInt("1000000000000000000"))
	if err != nil {
		t.Fatal(err)
	}
	utils.WaitForMiner(tx)

	// query provider balance before withdraw
	b1, err := pro.QueryBalance()
	if err != nil {
		t.Fatal(err)
	}

	n, err := op.GetNonce(got.Check.ToAddr)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println("nonce in contract:", n)

	fmt.Printf("-> withdraw with paycheck nonce %v\n", got.Check.Nonce)
	tx, err = pro.Withdraw(got)
	if err != nil {
		t.Fatal(err)
	}
	utils.WaitForMiner(tx)

	//fmt.Println("tx hash:", tx.Hash())

	gasUsed, err := utils.GetGasUsed(tx)
	if err != nil {
		t.Fatal(err)
	}

	// wait for a block be mined
	err = utils.WaitForMiner(tx)
	if err != nil {
		t.Fatal(err)
	}
	// query provider balance after withdraw
	b2, err := pro.QueryBalance()
	if err != nil {
		t.Fatal(err)
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

	fmt.Println("-> second withdraw: nonce 1 expected")
	// nonce 1 expected
	got, err = pro.GetNextPayable()
	if err != nil {
		t.Fatal(err)
	}
	if got == nil {
		t.Fatal("nil paycheck got")
	}
	if got.Check.Nonce != 1 {
		t.Fatalf("nonce=%v, nonce 1 expected\n", got.Check.Nonce)
	} else {
		fmt.Println("OK- withdrawed nonce: ", got.Check.Nonce)
	}

	// send 1 eth to provider
	fmt.Println("-> send 1 eth to provider")
	tx, err = utils.SendCoin(SenderSk, proAddr, utils.String2BigInt("1000000000000000000"))
	if err != nil {
		t.Fatal(err)
	}
	utils.WaitForMiner(tx)

	// query provider balance before withdraw
	b1, err = pro.QueryBalance()
	if err != nil {
		t.Fatal(err)
	}

	n, err = op.GetNonce(got.Check.ToAddr)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println("nonce in contract:", n)

	fmt.Printf("-> withdraw with paycheck nonce %v\n", got.Check.Nonce)

	tx, err = pro.Withdraw(got)
	if err != nil {
		t.Fatal(err)
	}
	utils.WaitForMiner(tx)

	//fmt.Println("tx hash:", tx.Hash())

	gasUsed, err = utils.GetGasUsed(tx)
	if err != nil {
		t.Fatal(err)
	}

	// wait for a block be mined
	err = utils.WaitForMiner(tx)
	if err != nil {
		t.Fatal(err)
	}
	// query provider balance after withdraw
	b2, err = pro.QueryBalance()
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println("balance before withdraw b1:", b1)
	fmt.Println("payvalue:", got.PayValue)
	fmt.Println("gasUsed:", gasUsed)
	fmt.Println("balance after withdraw b2:", b2)
	// need add used gas for withdraw tx
	delta = new(big.Int).Sub(b2, b1)
	delta.Add(delta, gasUsed)

	if delta.Cmp(got.PayValue) != 0 {
		t.Fatal("withdrawed money not equal payvalue")
	} else {
		fmt.Println("OK: withdrawed money equal payvalue")
	}

	fmt.Println("<< Pay at withdraw >>")
	fmt.Println("-> pay 0.1 eth: nonce 2 should be selected")
	// pay: 0.1 eth
	userPC, err = usr.Pay(proAddr, utils.String2BigInt(("100000000000000000")))
	if err != nil {
		t.Fatal(err)
	}
	if userPC == nil || userPC.Check.Nonce != 2 {
		t.Fatal("test pay 0.1 failed")
	} else {
		fmt.Println("OK- payer nonce: ", userPC.Check.Nonce)
	}

	// simulate provider receive paycheck from user
	proPC = new(check.Paycheck)
	*proPC = *userPC

	// provider verify received paycheck
	// datavalue: 0.1 eth
	ok, err = pro.Verify(proPC, utils.String2BigInt(("100000000000000000")))
	if err != nil {
		t.Fatal(err)
	}
	if !ok {
		t.Fatal("provider verify paycheck failed")
	}

	// provider store a paycheck into pool
	err = pro.StorePaycheck(proPC)
	if err != nil {
		t.Error("store paycheck error")
	}

	fmt.Println("-> pay 0.2 eth: nonce 2 should be selected")
	// pay: 0.2 eth
	userPC, err = usr.Pay(proAddr, utils.String2BigInt(("200000000000000000")))
	if err != nil {
		t.Fatal(err)
	}
	if userPC == nil || userPC.Check.Nonce != 2 {
		t.Fatal("test pay 0.2 failed")
	} else {
		fmt.Println("OK- payer nonce: ", userPC.Check.Nonce)
	}

	fmt.Println("paycheck created, but provider not received yet.")
	fmt.Println("and now provider withdrawed this paycheck before latest paycheck received and verified.")

	fmt.Println("-> third withdraw(before verify): nonce 2 expected")
	// nonce 1 expected
	got, err = pro.GetNextPayable()
	if err != nil {
		t.Fatal(err)
	}
	if got == nil {
		t.Fatal("nil paycheck got")
	}
	if got.Check.Nonce != 2 {
		t.Fatalf("nonce=%v, nonce 2 expected\n", got.Check.Nonce)
	} else {
		fmt.Println("OK- withdrawed nonce: ", got.Check.Nonce)
	}

	// send 1 eth to provider
	fmt.Println("-> send 1 eth to provider")
	tx, err = utils.SendCoin(SenderSk, proAddr, utils.String2BigInt("1000000000000000000"))
	if err != nil {
		t.Fatal(err)
	}
	utils.WaitForMiner(tx)

	// query provider balance before withdraw
	b1, err = pro.QueryBalance()
	if err != nil {
		t.Fatal(err)
	}

	n, err = op.GetNonce(got.Check.ToAddr)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println("nonce in contract:", n)

	fmt.Printf("-> withdraw with paycheck nonce %v\n", got.Check.Nonce)

	tx, err = pro.Withdraw(got)
	if err != nil {
		t.Fatal(err)
	}
	utils.WaitForMiner(tx)

	//fmt.Println("tx hash:", tx.Hash())

	gasUsed, err = utils.GetGasUsed(tx)
	if err != nil {
		t.Fatal(err)
	}

	// wait for a block be mined
	err = utils.WaitForMiner(tx)
	if err != nil {
		t.Fatal(err)
	}
	// query provider balance after withdraw
	b2, err = pro.QueryBalance()
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println("balance before withdraw b1:", b1)
	fmt.Println("payvalue:", got.PayValue)
	fmt.Println("gasUsed:", gasUsed)
	fmt.Println("balance after withdraw b2:", b2)
	// need add used gas for withdraw tx
	delta = new(big.Int).Sub(b2, b1)
	delta.Add(delta, gasUsed)

	if delta.Cmp(got.PayValue) != 0 {
		t.Fatal("withdrawed money not equal payvalue")
	} else {
		fmt.Println("OK- withdrawed money equal payvalue")
	}

	// simulate provider receive paycheck from user
	proPC = new(check.Paycheck)
	*proPC = *userPC

	fmt.Println("now verify the latest paycheck, should failed with nonce error.")
	// provider verify received paycheck
	// datavalue: 0.5 eth
	_, err = pro.Verify(proPC, utils.String2BigInt(("200000000000000000")))
	if err.Error() == "nonce should not less than contract nonce" {
		fmt.Println("OK- verify failed with error nonce")
	} else {
		t.Fatal("verify nonce should be detected.")
	}

}
