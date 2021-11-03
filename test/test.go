package test

import (
	"testing"

	"github.com/rockiecn/check/internal/utils"
	"github.com/rockiecn/check/operator"
	"github.com/rockiecn/check/provider"
	"github.com/rockiecn/check/user"
)

func TestAll(t *testing.T) {
	sk := utils.Generate()
	ts := "token"
	op, err := operator.NewOperator(sk, ts)

	usersk := utils.Generate()
	userAddr := utils.KeyToAddr(usersk)
	us, err := user.NewUser(usersk)

	proSk := utils.Generate()
	proAddr := utils.KeyToAddr(prosk)
	pro, err := provider.NewProvider(prosk)

	or := new(order.Order) // add useradd and proaddr to

	op.QueryOrder(or)

	chk, err := op.GenCheck(ord)

	us.StoreCheck(chk)
	pchk := us.GenPaycheck(proAddr, 10)

	pro.Verify(pchk)
	npchk = pro.GetNextPayable()
	pro.Withdraw(npchk)

}
