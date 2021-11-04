package test

import (
	"math/big"
	"testing"

	"github.com/rockiecn/check/internal/utils"
	"github.com/rockiecn/check/operator"
	"github.com/rockiecn/check/order"
	"github.com/rockiecn/check/provider"
	"github.com/rockiecn/check/user"
)

func TestAll(t *testing.T) {
	opSk := utils.Generate()
	ts := "token"
	op, err := operator.New(opSk, ts)
	if err != nil {
		t.Error(err)
	}

	usrSk := utils.Generate()
	usr, err := user.New(usrSk)
	if err != nil {
		t.Error(err)
	}

	proSk := utils.Generate()
	proAddr, err := utils.SkToAddr(proSk)
	if err != nil {
		t.Error(err)
	}
	pro, err := provider.New(proSk)
	if err != nil {
		t.Error(err)
	}

	// create new order manager
	om := new(order.OrderMgr)
	// add useraddr and proaddr to
	odr := new(order.Order)
	// store order into order manager
	om.PutOrder(odr)
	// get an order from order manager
	odr, err = op.QueryOrder(om, 0)
	if err != nil {
		t.Error(err)
	}
	if odr == nil {
		t.Error("query order error")
	}

	// generate a check from an order
	chk, err := op.NewCheck(0)
	if err != nil {
		t.Error(err)
	}

	// store a check into order manager
	err = usr.StoreCheck(om, chk)
	if err != nil {
		t.Error(err)
	}

	// generate a new paycheck for paying to provider
	pchk, err := usr.NewPaycheck(proAddr, utils.String2BigInt(("10")))
	if err != nil {
		t.Error(err)
	}

	// verify paycheck
	ok, err := pro.Verify(pchk, utils.String2BigInt(("10")))
	if err != nil {
		t.Error(err)
	}
	if !ok {
		t.Error("provider verify paycheck failed")
	}

	// get next payable paycheck from pool
	npchk, err := pro.GetNextPayable()
	if err != nil {
		t.Error(err)
	}

	// query balance before withdraw
	b1, err := op.QueryBalance()
	if err != nil {
		t.Error(err)
	}
	// call contract to withdraw a paycheck
	txHash, err := pro.Withdraw(npchk)
	if err != nil {
		t.Error(err)
	}
	// wait for a block be mined
	err = utils.WaitForMiner(txHash)
	if err != nil {
		t.Error(err)
	}
	// query balance after withdraw
	b2, err := op.QueryBalance()
	if err != nil {
		t.Error(err)
	}
	delta := new(big.Int).Sub(b2, b1)
	if delta.Cmp(pchk.PayValue) != 0 {
		t.Error("withdrawed money not equal payvalue")
	}

}
