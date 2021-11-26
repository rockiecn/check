// common functions for test
package common

import (
	"errors"
	"fmt"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/rockiecn/check/internal/check"
	"github.com/rockiecn/check/internal/odrmgr"
	"github.com/rockiecn/check/internal/utils"
	"github.com/rockiecn/check/operator"
	"github.com/rockiecn/check/provider"
	"github.com/rockiecn/check/user"
)

var (
	// a local account, with enough money in it
	SenderSk = "503f38a9c967ed597e47fe25643985f032b072db8075426a92110f82df48dfcb"
	Token    = common.HexToAddress("0xb213d01542d129806d664248a380db8b12059061")
)

func InitOperator() (*operator.Operator, error) {
	// generate operator
	opSk, err := utils.GenerateSK()
	if err != nil {
		return nil, err
	}
	op, err := operator.New(opSk)
	if err != nil {
		return nil, err
	}
	Op := op.(*operator.Operator)

	// send 2 eth to operator
	fmt.Println("sending coin to operator")
	tx, err := utils.SendCoin(SenderSk, Op.OpAddr, utils.String2BigInt("2000000000000000000"))
	if err != nil {
		return nil, err
	}
	err = utils.WaitForMiner(tx)
	if err != nil {
		return nil, err
	}

	fmt.Println("deploying contract")
	// operator deploy contract, with 1.8 eth
	tx, ctrAddr, err := op.Deploy(utils.String2BigInt("1800000000000000000"))
	if err != nil {
		return nil, err
	}
	err = utils.WaitForMiner(tx)
	if err != nil {
		return nil, err
	}

	// set contract address for operator
	Op.SetCtrAddr(ctrAddr)

	return Op, nil
}

func InitUser() (*user.User, error) {
	// generate user
	usrSk, err := utils.GenerateSK()
	if err != nil {
		return nil, err
	}
	usr, err := user.New(usrSk)
	if err != nil {
		return nil, err
	}
	Usr := usr.(*user.User)

	return Usr, nil
}

func InitPro() (*provider.Provider, error) {
	// provider0
	proSk, err := utils.GenerateSK()
	if err != nil {
		return nil, err
	}
	pro, err := provider.New(proSk)
	if err != nil {
		return nil, err
	}

	Pro := pro.(*provider.Provider)

	// send 0.1 eth to provider
	fmt.Println("sending coin to provider for tx")
	tx, err := utils.SendCoin(SenderSk, Pro.ProviderAddr, utils.String2BigInt("100000000000000000"))
	if err != nil {
		return nil, err
	}
	err = utils.WaitForMiner(tx)
	if err != nil {
		return nil, err
	}

	return Pro, nil
}

func InitOrder(
	id uint64,
	usr *user.User,
	op *operator.Operator,
	pro *provider.Provider,
	v string,
) error {
	odr := &odrmgr.Order{
		ID:    id,
		Token: Token,
		Value: utils.String2BigInt(v), // order value: 0.3 eth
		From:  usr.UserAddr,
		To:    pro.ProviderAddr,
		Time:  time.Now().Unix(),
		Name:  "jack",
		Tel:   "123123123",
		Email: "asdf@asdf.com",
		State: 0,
	}
	if odr == nil {
		return errors.New("create order failed")
	}

	err := op.OM.PutOrder(odr)
	if err != nil {
		return err
	}
	// operator create a check from order
	opChk, err := op.CreateCheck(id)
	if err != nil {
		return err
	}
	// simulate user receive check from operator
	usrChk := new(check.Check)
	*usrChk = *opChk

	// generate paycheck from check
	pchk, err := usr.GenPchk(usrChk)
	if err != nil {
		return err
	}
	// user put paycheck into pool
	err = usr.Put(pchk)
	if err != nil {
		return err
	}

	return nil
}

func Pay(
	usr *user.User,
	pro *provider.Provider,
	v string,
) (uint64, error) {
	// user generate a paycheck for paying to provider
	// store new paycheck into user pool
	userPC, err := usr.Pay(pro.ProviderAddr, utils.String2BigInt(v))
	if err != nil {
		return 0, err
	}

	if userPC == nil {
		return 0, errors.New("generate nil paycheck")
	}

	// simulate provider receive paycheck from user
	proPC := new(check.Paycheck)
	*proPC = *userPC

	// provider verify received paycheck
	// datavalue: 0.1 eth
	ok, err := pro.Verify(proPC, utils.String2BigInt(v))
	if err != nil {
		return 0, err
	}
	if !ok {
		return 0, errors.New("provider verify paycheck failed")
	}

	// provider store a paycheck into pool
	err = pro.StorePaycheck(proPC)
	if err != nil {
		return 0, errors.New("store paycheck error")
	}

	return userPC.Nonce, nil
}

func Withdraw(
	op *operator.Operator,
	pro *provider.Provider,
) (uint64, error) {
	// nonce 0 expected
	got, err := pro.GetNextPayable()
	if err != nil {
		return 0, err
	}
	if got == nil {
		return 0, errors.New("nil paycheck got")
	}

	// query provider balance before withdraw
	b1, err := pro.QueryBalance()
	if err != nil {
		return 0, err
	}

	ctNonce, err := op.GetNonce(got.ToAddr)
	if err != nil {
		return 0, err
	}
	fmt.Println("nonce in contract:", ctNonce)

	fmt.Printf("withdrawing, nonce: %v\n", got.Nonce)
	tx, err := pro.Withdraw(got)
	if err != nil {
		return 0, err
	}
	err = utils.WaitForMiner(tx)
	if err != nil {
		return 0, err
	}

	gasUsed, err := utils.GetGasUsed(tx)
	if err != nil {
		return 0, err
	}

	// query provider balance after withdraw
	b2, err := pro.QueryBalance()
	if err != nil {
		return 0, err
	}

	// need add used gas for withdraw tx
	delta := new(big.Int).Sub(b2, b1)
	newPV := big.NewInt(0).Add(delta, gasUsed)

	if newPV.Cmp(got.PayValue) != 0 {
		return 0, errors.New("withdrawed money not equal payvalue")
	}
	fmt.Println("OK- withdrawed money equal payvalue")

	return got.Nonce, nil
}
