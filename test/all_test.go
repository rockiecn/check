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

	fmt.Println("<< PHASE 1: Initialize >>")
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
	// send 2 eth to operator
	tx, err := utils.SendCoin(senderSk, opAddr, utils.String2BigInt("2000000000000000000"))
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

	fmt.Println("<< PHASE 2: New Order >>")

	// create an order
	token := common.HexToAddress("0xb213d01542d129806d664248a380db8b12059061")
	odr := order.NewOdr(0,
		token,
		usrAddr,
		proAddr,
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

	fmt.Println("<< PHASE 3: Order to Check >>")
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

	fmt.Println("<< PHASE 4: user pay 0.1 eth >>")
	// user generate a paycheck for paying to provider
	// store new paycheck into user pool
	// pay: 0.1 eth
	userPC, err := usr.Pay(proAddr, utils.String2BigInt(("100000000000000000")))
	if err != nil {
		t.Error(err)
	}

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

	fmt.Println("<< PHASE 5: user pay 0.2 eth >>")
	// pay: 0.2 eth
	userPC, err = usr.Pay(proAddr, utils.String2BigInt(("200000000000000000")))
	if err != nil {
		t.Error(err)
	}

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

	fmt.Println("<< PHASE 6: user pay 0.3 eth >>")
	// pay: 0.3 eth
	userPC, err = usr.Pay(proAddr, utils.String2BigInt(("300000000000000000")))
	if err != nil {
		t.Error(err)
	}

	// simulate provider receive paycheck from user
	proPC = new(check.Paycheck)
	*proPC = *userPC

	// provider verify received paycheck
	// datavalue: 0.2 eth
	ok, err = pro.Verify(proPC, utils.String2BigInt(("300000000000000000")))
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

	fmt.Println("<< PHASE 7: provider withdraw >>")

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
}
