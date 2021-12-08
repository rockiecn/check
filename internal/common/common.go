// common functions for test
package common

import (
	"errors"
	"fmt"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/rockiecn/check/internal/check"
	"github.com/rockiecn/check/internal/utils"
	"github.com/rockiecn/check/operator"
	"github.com/rockiecn/check/provider"
	"github.com/rockiecn/check/user"
)

var (
	// a local account, with enough money in it
	//SenderSk = "503f38a9c967ed597e47fe25643985f032b072db8075426a92110f82df48dfcb"
	//
	SenderSk = "0a95533a110ee10bdaa902fed92e56f3f7709a532e22b5974c03c0251648a5d4"
	Token    = common.HexToAddress("0xb213d01542d129806d664248a380db8b12059061")
)

func InitOperator(odrDB string, chkDB string) (*operator.Operator, error) {
	// generate operator
	opSk, err := utils.GenerateSK()
	if err != nil {
		return nil, err
	}
	op, err := operator.New(opSk, odrDB, chkDB)
	if err != nil {
		return nil, err
	}
	Op := op.(*operator.Operator)

	// send 2 eth to operator
	fmt.Printf("sending 0.1 eth to operator: %v\n", Op.OpAddr)
	tx, err := utils.SendCoin(SenderSk, Op.OpAddr, utils.String2BigInt("100000000000000000"))
	if err != nil {
		return nil, err
	}
	err = utils.WaitForMiner(tx)
	if err != nil {
		return nil, err
	}

	fmt.Println("deploying contract with 0.08 eth")
	// operator deploy contract, with 0.08 eth
	tx, ctrAddr, err := op.Deploy(utils.String2BigInt("80000000000000000"))
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

func InitUser(pcDB string) (*user.User, error) {
	// generate user
	usrSk, err := utils.GenerateSK()
	if err != nil {
		return nil, err
	}
	usr, err := user.New(usrSk, pcDB)
	if err != nil {
		return nil, err
	}
	Usr := usr.(*user.User)

	return Usr, nil
}

func InitPro(pcDB string, btDB string) (*provider.Provider, error) {
	// provider0
	proSk, err := utils.GenerateSK()
	if err != nil {
		return nil, err
	}
	pro, err := provider.New(proSk, pcDB, btDB)
	if err != nil {
		return nil, err
	}

	Pro := pro.(*provider.Provider)

	// send 0.01 eth to provider
	fmt.Println("sending 0.001 eth to provider for tx")
	tx, err := utils.SendCoin(SenderSk, Pro.ProviderAddr, utils.String2BigInt("1000000000000000"))
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
	oid uint64,
	usr *user.User,
	op *operator.Operator,
	pro *provider.Provider,
	v string,
) error {
	odr := &operator.Order{
		ID:    oid,
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

	err := op.PutOrder(odr)
	if err != nil {
		return err
	}
	// operator create a check from order
	opChk, err := op.CreateCheck(oid)
	if err != nil {
		return err
	}
	err = op.PutCheck(oid, opChk)
	if err != nil {
		return err
	}
	err = op.StoreChk(oid, opChk)
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
	// store pc into db
	err = usr.Store(pchk)
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
	err = pro.Put(proPC)
	if err != nil {
		return 0, errors.New("put paycheck error")
	}

	// store paycheck into provider db
	err = pro.Store(proPC)
	if err != nil {
		return 0, errors.New("store paycheck into db error")
	}

	return userPC.Nonce, nil
}

func Withdraw(
	op *operator.Operator,
	pro *provider.Provider,
) (uint64, error) {
	// nonce 0 expected
	got, err := pro.PcStorer.GetNextPayable()
	if err != nil {
		return 0, fmt.Errorf("call pro.GetNextPayable failed: %v", err)
	}
	if got == nil {
		return 0, errors.New("nil paycheck got")
	}

	// query provider balance before withdraw
	b1, err := pro.QueryBalance()
	if err != nil {
		return 0, fmt.Errorf("call pro.QueryBalance failed: %v", err)
	}

	ctNonce, err := op.GetNonce(got.ToAddr)
	if err != nil {
		return 0, fmt.Errorf("call op.GetNonce failed: %v", err)
	}
	fmt.Println("nonce in contract:", ctNonce)

	fmt.Printf("withdrawing, nonce: %v\n", got.Nonce)
	tx, err := pro.Withdraw(got)
	if err != nil {
		return 0, fmt.Errorf("call pro.Withdraw failed: %v", err)
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
	realPV := new(big.Int).Add(delta, gasUsed)

	if utils.Debug {
		fmt.Println("   delta: ", delta)
		fmt.Println(" gasUsed: ", gasUsed)
		fmt.Println("  realPV: ", realPV.String())
		fmt.Println("payvalue: ", got.PayValue)
	}

	if realPV.Cmp(got.PayValue) != 0 {
		return 0, errors.New("withdrawed money not equal payvalue")
	}
	fmt.Println("OK- withdrawed money equal payvalue")

	return got.Nonce, nil
}
