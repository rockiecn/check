package test

import (
	"fmt"
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
func TestMultiProMultiCheck(t *testing.T) {

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

	fmt.Println("<< Create Provider 0 >>")
	// generate provider
	proSk0, err := utils.GenerateSK()
	if err != nil {
		t.Error(err)
	}
	proAddr0, err := utils.SkToAddr(proSk0)
	if err != nil {
		t.Error(err)
	}
	pro0, err := provider.New(proSk0)
	if err != nil {
		t.Error(err)
	}

	fmt.Println("<< Create Provider 1 >>")
	// generate provider
	proSk1, err := utils.GenerateSK()
	if err != nil {
		t.Error(err)
	}
	proAddr1, err := utils.SkToAddr(proSk1)
	if err != nil {
		t.Error(err)
	}
	pro1, err := provider.New(proSk1)
	if err != nil {
		t.Error(err)
	}

	fmt.Println("<< Create Provider 2 >>")
	// generate provider
	proSk2, err := utils.GenerateSK()
	if err != nil {
		t.Error(err)
	}
	proAddr2, err := utils.SkToAddr(proSk2)
	if err != nil {
		t.Error(err)
	}
	pro2, err := provider.New(proSk2)
	if err != nil {
		t.Error(err)
	}

	fmt.Println("<< New Order, ID:0 >>")

	// create an order
	odr := order.NewOdr(0,
		token,
		usrAddr,
		proAddr0,
		utils.String2BigInt("1500000000000000000"), // order value: 1.5 eth
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

	// operator store order into pool
	err = op.StoreOrder(odr)
	if err != nil {
		t.Error(err)
	}
	// operator get an order by id
	odr, err = op.GetOrder(0)
	if err != nil {
		t.Error(err)
	}
	if odr == nil {
		t.Error("get order failed")
	}

	fmt.Println("<< Order to Check >>")
	// operator create a check from order
	opChk, err := op.CreateCheck(0)
	if err != nil {
		t.Error(err)
	}

	// simulate user receive check from operator
	usrChk := new(check.Check)
	*usrChk = *opChk
	// user store check into pool
	err = usr.StoreCheck(usrChk)
	if err != nil {
		t.Error(err)
	}

	fmt.Println("<< New Order, ID:1 >>")

	// create an order
	odr = order.NewOdr(1,
		token,
		usrAddr,
		proAddr1,
		utils.String2BigInt("1500000000000000000"), // order value: 1.5 eth
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

	// operator store order into pool
	err = op.StoreOrder(odr)
	if err != nil {
		t.Error(err)
	}
	// operator get an order by id
	odr, err = op.GetOrder(1)
	if err != nil {
		t.Error(err)
	}
	if odr == nil {
		t.Error("get order failed")
	}

	fmt.Println("<< Order to Check >>")
	// operator create a check from order
	opChk, err = op.CreateCheck(1)
	if err != nil {
		t.Error(err)
	}

	// simulate user receive check from operator
	usrChk = new(check.Check)
	*usrChk = *opChk
	// user store check into pool
	err = usr.StoreCheck(usrChk)
	if err != nil {
		t.Error(err)
	}

	fmt.Println("<< New Order, ID:2 >>")

	// create an order
	odr = order.NewOdr(2,
		token,
		usrAddr,
		proAddr2,
		utils.String2BigInt("1500000000000000000"), // order value: 1.5 eth
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

	// operator store order into pool
	err = op.StoreOrder(odr)
	if err != nil {
		t.Error(err)
	}
	// operator get an order by id
	odr, err = op.GetOrder(2)
	if err != nil {
		t.Error(err)
	}
	if odr == nil {
		t.Error("get order failed")
	}

	fmt.Println("<< Order to Check >>")
	// operator create a check from order
	opChk, err = op.CreateCheck(2)
	if err != nil {
		t.Error(err)
	}

	// simulate user receive check from operator
	usrChk = new(check.Check)
	*usrChk = *opChk
	// user store check into pool
	err = usr.StoreCheck(usrChk)
	if err != nil {
		t.Error(err)
	}

	fmt.Println("<< Pay >>")
	fmt.Println("<< pay 0.1 eth to pro0>>")
	// user generate a paycheck for paying to provider
	// store new paycheck into user pool
	// pay: 0.1 eth
	userPC, err := usr.Pay(proAddr0, utils.String2BigInt(("100000000000000000")))
	if err != nil {
		t.Error(err)
	}

	// simulate provider receive paycheck from user
	proPC := new(check.Paycheck)
	*proPC = *userPC

	// provider verify received paycheck
	// datavalue: 0.1 eth
	ok, err := pro0.Verify(proPC, utils.String2BigInt(("100000000000000000")))
	if err != nil {
		t.Error(err)
	}
	if !ok {
		t.Error("provider verify paycheck failed")
	}

	// provider store a paycheck into pool
	err = pro0.StorePaycheck(proPC)
	if err != nil {
		t.Error("store paycheck error")
	}

	fmt.Println("<< pay 0.2 eth to pro 1 >>")
	// pay: 0.2 eth
	userPC, err = usr.Pay(proAddr1, utils.String2BigInt(("200000000000000000")))
	if err != nil {
		t.Error(err)
	}

	// simulate provider receive paycheck from user
	proPC = new(check.Paycheck)
	*proPC = *userPC

	// provider verify received paycheck
	// datavalue: 0.2 eth
	ok, err = pro1.Verify(proPC, utils.String2BigInt(("200000000000000000")))
	if err != nil {
		t.Error(err)
	}
	if !ok {
		t.Error("provider verify paycheck failed")
	}

	// provider store a paycheck into pool
	err = pro1.StorePaycheck(proPC)
	if err != nil {
		t.Error(err)
	}

	fmt.Println("<< pay 0.3 eth to pro 2 >>")
	// pay: 0.3 eth
	userPC, err = usr.Pay(proAddr2, utils.String2BigInt(("300000000000000000")))
	if err != nil {
		t.Error(err)
	}

	// simulate provider receive paycheck from user
	proPC = new(check.Paycheck)
	*proPC = *userPC

	// provider verify received paycheck
	// datavalue: 0.2 eth
	ok, err = pro2.Verify(proPC, utils.String2BigInt(("300000000000000000")))
	if err != nil {
		t.Error(err)
	}
	if !ok {
		t.Error("provider verify paycheck failed")
	}

	// provider store a paycheck into pool
	err = pro2.StorePaycheck(proPC)
	if err != nil {
		t.Error(err)
	}

	fmt.Println("<< Withdraw >>")

	/*
		// provider get next payable paycheck from pool
		npchk, err := pro.GetNextPayable()
		if err != nil {
			t.Error(err)
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
		fmt.Println("payvalue:", npchk.PayValue)
		fmt.Println("gasUsed:", gasUsed)
		fmt.Println("balance after withdraw b2:", b2)
		// need add used gas for withdraw tx
		delta := new(big.Int).Sub(b2, b1)
		delta.Add(delta, gasUsed)

		if delta.Cmp(npchk.PayValue) != 0 {
			t.Error("withdrawed money not equal payvalue")
		} else {
			fmt.Println("withdrawed money equal payvalue")
		}

	*/
}
