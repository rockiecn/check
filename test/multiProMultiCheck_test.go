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

	fmt.Println("-> New User")

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

	fmt.Println("-> New 3 Providers ")
	// provider0
	proSk0, err := utils.GenerateSK()
	if err != nil {
		t.Fatal(err)
	}
	proAddr0, err := utils.SkToAddr(proSk0)
	if err != nil {
		t.Fatal(err)
	}
	pro0, err := provider.New(proSk0)
	if err != nil {
		t.Fatal(err)
	}

	// provider1
	proSk1, err := utils.GenerateSK()
	if err != nil {
		t.Fatal(err)
	}
	proAddr1, err := utils.SkToAddr(proSk1)
	if err != nil {
		t.Fatal(err)
	}
	pro1, err := provider.New(proSk1)
	if err != nil {
		t.Fatal(err)
	}

	// provider2
	proSk2, err := utils.GenerateSK()
	if err != nil {
		t.Fatal(err)
	}
	proAddr2, err := utils.SkToAddr(proSk2)
	if err != nil {
		t.Fatal(err)
	}
	pro2, err := provider.New(proSk2)
	if err != nil {
		t.Fatal(err)
	}

	fmt.Println("-> Order ID=0, value=0.3")

	// create  order for each provider
	odr0 := order.NewOdr(0,
		token,
		usrAddr,
		proAddr0,
		utils.String2BigInt("300000000000000000"), // order value: 0.3 eth
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
		proAddr1,
		utils.String2BigInt("300000000000000000"), // order value: 0.3 eth
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

	fmt.Println("-> Order ID=2, value=0.3")

	// create an order
	odr2 := order.NewOdr(2,
		token,
		usrAddr,
		proAddr2,
		utils.String2BigInt("300000000000000000"), // order value: 0.5 eth
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
	fmt.Println("-> pay 0.1 eth to provider0")
	// user generate a paycheck for paying to provider
	// store new paycheck into user pool
	// pay: 0.1 eth
	userPC, err := usr.Pay(proAddr0, utils.String2BigInt(("100000000000000000")))
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
	ok, err := pro0.Verify(proPC, utils.String2BigInt(("100000000000000000")))
	if err != nil {
		t.Fatal(err)
	}
	if !ok {
		t.Fatal("provider verify paycheck failed")
	}

	// provider store a paycheck into pool
	err = pro0.StorePaycheck(proPC)
	if err != nil {
		t.Fatal("store paycheck error")
	}

	fmt.Println("-> pay 0.2 eth to provider1")
	// pay: 0.2 eth
	userPC, err = usr.Pay(proAddr1, utils.String2BigInt(("200000000000000000")))
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
	ok, err = pro1.Verify(proPC, utils.String2BigInt(("200000000000000000")))
	if err != nil {
		t.Fatal(err)
	}
	if !ok {
		t.Fatal("provider verify paycheck failed")
	}

	// provider store a paycheck into pool
	err = pro1.StorePaycheck(proPC)
	if err != nil {
		t.Fatal(err)
	}

	fmt.Println("-> pay 0.3 eth to provider2")
	// user generate a paycheck for paying to provider
	// store new paycheck into user pool
	// pay: 0.3 eth
	userPC, err = usr.Pay(proAddr2, utils.String2BigInt(("300000000000000000")))
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
	ok, err = pro2.Verify(proPC, utils.String2BigInt(("300000000000000000")))
	if err != nil {
		t.Fatal(err)
	}
	if !ok {
		t.Fatal("provider verify paycheck failed")
	}

	// provider store a paycheck into pool
	err = pro2.StorePaycheck(proPC)
	if err != nil {
		t.Error("store paycheck error")
	}

	fmt.Println("<< Withdraw >>")
	fmt.Println("-> provider0 withdraw")
	// nonce 0 expected
	got, err := pro0.GetNextPayable()
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
	tx, err = utils.SendCoin(SenderSk, proAddr0, utils.String2BigInt("1000000000000000000"))
	if err != nil {
		t.Fatal(err)
	}
	utils.WaitForMiner(tx)

	// query provider balance before withdraw
	b1, err := pro0.QueryBalance()
	if err != nil {
		t.Fatal(err)
	}

	n, err := op.GetNonce(got.Check.ToAddr)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println("nonce in contract:", n)

	fmt.Printf("-> withdraw with paycheck nonce %v\n", got.Check.Nonce)
	tx, err = pro0.Withdraw(got)
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
	b2, err := pro0.QueryBalance()
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
	got, err = pro1.GetNextPayable()
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
	tx, err = utils.SendCoin(SenderSk, proAddr1, utils.String2BigInt("1000000000000000000"))
	if err != nil {
		t.Fatal(err)
	}
	utils.WaitForMiner(tx)

	// query provider balance before withdraw
	b1, err = pro1.QueryBalance()
	if err != nil {
		t.Fatal(err)
	}

	n, err = op.GetNonce(got.Check.ToAddr)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println("nonce in contract:", n)

	fmt.Printf("-> withdraw with paycheck nonce %v\n", got.Check.Nonce)

	tx, err = pro1.Withdraw(got)
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
	b2, err = pro1.QueryBalance()
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
	got, err = pro2.GetNextPayable()
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
	tx, err = utils.SendCoin(SenderSk, proAddr2, utils.String2BigInt("1000000000000000000"))
	if err != nil {
		t.Fatal(err)
	}
	utils.WaitForMiner(tx)

	// query provider balance before withdraw
	b1, err = pro2.QueryBalance()
	if err != nil {
		t.Fatal(err)
	}

	n, err = op.GetNonce(got.Check.ToAddr)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println("nonce in contract:", n)

	fmt.Printf("-> withdraw with paycheck nonce %v\n", got.Check.Nonce)

	tx, err = pro2.Withdraw(got)
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
	b2, err = pro2.QueryBalance()
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
