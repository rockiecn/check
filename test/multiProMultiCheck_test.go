package test

import (
	"fmt"
	"math/big"
	"testing"
	"time"

	"github.com/rockiecn/check/internal/check"
	"github.com/rockiecn/check/internal/order"
	"github.com/rockiecn/check/internal/utils"
)

// 3 providers involved
// 3 checks used
// 3 pay actions by user
// 3 withdraw actions by provider
//
// Process:
// 1. create operator, user
// 2. create 3 providers
// 3. operator create 3 orders pay to each provider, all has 0.3 eth
// 4. operator transfer orders to checks and send to user
// 5. user store checks into pool
// 6. user pay provider0 0.1 eth with check0
// 7. user pay provider1 0.2 eth with check1
// 8. user pay provider2 0.3 eth with check2
// 9. provider0 withdraw, should receive 0.1 eth
// 10. provider1 withdraw, should receive 0.2 eth
// 11. provider2 withdraw, should receive 0.3 eth
func TestMultiProMultiCheck(t *testing.T) {

	fmt.Println("<< Initialize >>")

	fmt.Println("-> Init Operator")
	Op := InitOperator(t)

	fmt.Println("-> Init User")
	Usr := InitUser(t)

	fmt.Println("-> Init 3 Providers ")
	Pro0 := InitPro(t)
	Pro1 := InitPro(t)
	Pro2 := InitPro(t)

	fmt.Println("-> Init Order ID=0, value=0.3")
	odr0 := order.NewOdr(0,
		token,
		Usr.UserAddr,
		Pro0.ProviderAddr,
		utils.String2BigInt("300000000000000000"), // order value: 0.3 eth
		time.Now(),
		"jack",
		"123123123",
		"asdf@asdf.com",
		0,
		nil,
	)
	if odr0 == nil {
		t.Fatal("create order failed")
	}
	err := Op.StoreOrder(odr0)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println("-> Order to Check")
	// operator create a check from order
	opChk0, err := Op.CreateCheck(0)
	if err != nil {
		t.Fatal(err)
	}
	// simulate user receive check from operator
	usrChk0 := new(check.Check)
	*usrChk0 = *opChk0
	// user store check into pool
	err = Usr.StoreCheck(usrChk0)
	if err != nil {
		t.Fatal(err)
	}

	fmt.Println("-> Init Order ID=1, value=0.3")
	odr1 := order.NewOdr(1,
		token,
		Usr.UserAddr,
		Pro1.ProviderAddr,
		utils.String2BigInt("300000000000000000"), // order value: 0.3 eth
		time.Now(),
		"jack",
		"123123123",
		"asdf@asdf.com",
		0,
		nil,
	)
	if odr1 == nil {
		t.Fatal("create order failed")
	}
	err = Op.StoreOrder(odr1)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println("-> Order to Check")
	// operator create a check from order
	opChk1, err := Op.CreateCheck(1)
	if err != nil {
		t.Fatal(err)
	}
	// simulate user receive check from operator
	usrChk1 := new(check.Check)
	*usrChk1 = *opChk1
	// user store check into pool
	err = Usr.StoreCheck(usrChk1)
	if err != nil {
		t.Fatal(err)
	}

	fmt.Println("-> Init Order ID=2, value=0.3")
	odr2 := order.NewOdr(2,
		token,
		Usr.UserAddr,
		Pro2.ProviderAddr,
		utils.String2BigInt("300000000000000000"), // order value: 0.3 eth
		time.Now(),
		"jack",
		"123123123",
		"asdf@asdf.com",
		0,
		nil,
	)
	if odr2 == nil {
		t.Fatal("create order failed")
	}
	err = Op.StoreOrder(odr2)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println("-> Order to Check")
	// operator create a check from order
	opChk2, err := Op.CreateCheck(2)
	if err != nil {
		t.Fatal(err)
	}
	// simulate user receive check from operator
	usrChk2 := new(check.Check)
	*usrChk2 = *opChk2
	// user store check into pool
	err = Usr.StoreCheck(usrChk2)
	if err != nil {
		t.Fatal(err)
	}

	fmt.Println("<< Pay >>")
	fmt.Println("-> pay 0.1 eth to provider0")
	// user generate a paycheck for paying to provider
	// store new paycheck into user pool
	// pay: 0.1 eth
	userPC, err := Usr.Pay(Pro0.ProviderAddr, utils.String2BigInt(("100000000000000000")))
	if err != nil {
		t.Fatal(err)
	}

	if userPC == nil || userPC.Check.Nonce != 0 {
		t.Fatal("test pay 0.1 failed")
	} else {
		fmt.Println("OK- paycheck nonce: ", userPC.Check.Nonce)
	}

	// simulate provider receive paycheck from user
	proPC := new(check.Paycheck)
	*proPC = *userPC

	// provider verify received paycheck
	// datavalue: 0.1 eth
	ok, err := Pro0.Verify(proPC, utils.String2BigInt(("100000000000000000")))
	if err != nil {
		t.Fatal(err)
	}
	if !ok {
		t.Fatal("provider verify paycheck failed")
	}

	// provider store a paycheck into pool
	err = Pro0.StorePaycheck(proPC)
	if err != nil {
		t.Fatal("store paycheck error")
	}

	fmt.Println("-> pay 0.2 eth to provider1")
	// pay: 0.2 eth
	userPC, err = Usr.Pay(Pro1.ProviderAddr, utils.String2BigInt(("200000000000000000")))
	if err != nil {
		t.Fatal(err)
	}
	if userPC == nil || userPC.Check.Nonce != 0 {
		t.Fatal("test pay 0.2 failed")
	} else {
		fmt.Println("OK- paycheck1 nonce: ", userPC.Check.Nonce)
	}

	// simulate provider receive paycheck from user
	proPC = new(check.Paycheck)
	*proPC = *userPC

	// provider verify received paycheck
	// datavalue: 0.2 eth
	ok, err = Pro1.Verify(proPC, utils.String2BigInt(("200000000000000000")))
	if err != nil {
		t.Fatal(err)
	}
	if !ok {
		t.Fatal("provider verify paycheck failed")
	}

	// provider store a paycheck into pool
	err = Pro1.StorePaycheck(proPC)
	if err != nil {
		t.Fatal(err)
	}

	fmt.Println("-> pay 0.3 eth to provider2")
	// user generate a paycheck for paying to provider
	// store new paycheck into user pool
	// pay: 0.3 eth
	userPC, err = Usr.Pay(Pro2.ProviderAddr, utils.String2BigInt(("300000000000000000")))
	if err != nil {
		t.Fatal(err)
	}
	if userPC == nil || userPC.Check.Nonce != 0 {
		t.Fatal("test pay 0.3 failed")
	} else {
		fmt.Println("OK- paycheck2 nonce: ", userPC.Check.Nonce)
	}

	// simulate provider receive paycheck from user
	proPC = new(check.Paycheck)
	*proPC = *userPC

	// provider verify received paycheck
	// datavalue: 0.2 eth
	ok, err = Pro2.Verify(proPC, utils.String2BigInt(("300000000000000000")))
	if err != nil {
		t.Fatal(err)
	}
	if !ok {
		t.Fatal("provider verify paycheck failed")
	}

	// provider store a paycheck into pool
	err = Pro2.StorePaycheck(proPC)
	if err != nil {
		t.Error("store paycheck error")
	}

	fmt.Println("<< Withdraw >>")
	fmt.Println("-> provider0 withdraw")
	// nonce 0 expected
	got, err := Pro0.GetNextPayable()
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
	tx, err := utils.SendCoin(SenderSk, Pro0.ProviderAddr, utils.String2BigInt("1000000000000000000"))
	if err != nil {
		t.Fatal(err)
	}
	utils.WaitForMiner(tx)

	// query provider balance before withdraw
	b1, err := Pro0.QueryBalance()
	if err != nil {
		t.Fatal(err)
	}

	n, err := Op.GetNonce(got.Check.ToAddr)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println("nonce in contract:", n)

	fmt.Printf("-> withdraw with paycheck nonce %v\n", got.Check.Nonce)
	tx, err = Pro0.Withdraw(got)
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
	b2, err := Pro0.QueryBalance()
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

	fmt.Println("-> provider1 withdraw")
	// nonce 1 expected
	got, err = Pro1.GetNextPayable()
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
	fmt.Println("-> send 1 eth to provider")
	tx, err = utils.SendCoin(SenderSk, Pro1.ProviderAddr, utils.String2BigInt("1000000000000000000"))
	if err != nil {
		t.Fatal(err)
	}
	utils.WaitForMiner(tx)

	// query provider balance before withdraw
	b1, err = Pro1.QueryBalance()
	if err != nil {
		t.Fatal(err)
	}

	n, err = Op.GetNonce(got.Check.ToAddr)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println("nonce in contract:", n)

	fmt.Printf("-> withdraw with paycheck nonce %v\n", got.Check.Nonce)

	tx, err = Pro1.Withdraw(got)
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
	b2, err = Pro1.QueryBalance()
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

	fmt.Println("-> provider2 withdraw")
	// nonce 1 expected
	got, err = Pro2.GetNextPayable()
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
	fmt.Println("-> send 1 eth to provider")
	tx, err = utils.SendCoin(SenderSk, Pro2.ProviderAddr, utils.String2BigInt("1000000000000000000"))
	if err != nil {
		t.Fatal(err)
	}
	utils.WaitForMiner(tx)

	// query provider balance before withdraw
	b1, err = Pro2.QueryBalance()
	if err != nil {
		t.Fatal(err)
	}

	n, err = Op.GetNonce(got.Check.ToAddr)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println("nonce in contract:", n)

	fmt.Printf("-> withdraw with paycheck nonce %v\n", got.Check.Nonce)

	tx, err = Pro2.Withdraw(got)
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
	b2, err = Pro2.QueryBalance()
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
}
