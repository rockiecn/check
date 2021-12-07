package test

import (
	"errors"
	"fmt"
	"testing"

	"github.com/rockiecn/check/internal/check"
	"github.com/rockiecn/check/internal/common"
	"github.com/rockiecn/check/internal/utils"
)

// A single provider involved
// multi checks used(with the same provider)
// multi pay actions by user
// multi withdraw actions by provider
//
// Process:
// 1. create operator, user, provider
// 2. operator create 3 orders, each has 0.5 eth
// 3. operator create 3 checks from order
// 4. user store checks into pool
// 5. user pay 0.1 eth with check 0
// 6. user pay 0.2 eth with check 0
// 7. user pay 0.4 eth with check 1 (check 0 not enough money, only 0.2 remained)
// 8. user pay 0.2 eth with check 0 again(0.2 remained in check 0)
// 9. user pay 0.6 eth with failed(no check can pay, all value is 0.5)
// 10. provider first withdraw, check 0 should be selected
// 11. provider second withdraw, check 1 should be selected
// 12. user pay 0.1 eth with check 2, provider received it(payvalue=0.1)
// 13. user pay 0.2 eth with check 2 without provider receive it
// 14. provider third withdraw check 2 before latest paycheck is received(now payvalue=0.1)
// 15. then provider received this delayed paycheck and
// 	   verify should be failed(nonce in contract already changed by third withdraw)
func TestSingleProMultiCheck(t *testing.T) {

	// init roles
	fmt.Println("-> Init Operator")
	op, err := common.InitOperator("./op/order.db", "./op/check.db")
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println("-> Init Provider")
	pro, err := common.InitPro("./pro/pc.db", "./pro/bt.db")
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println("-> Init User")
	usr, err := common.InitUser("./usr/pc.db")
	if err != nil {
		t.Fatal(err)
	}

	op.ChkStorer.Clear()
	op.OdrStorer.Clear()
	usr.PcStorer.Clear()
	pro.BtStorer.Clear()
	pro.PcStorer.Clear()

	// create 3 orders
	fmt.Println("-> Init 3 orders")
	err = common.InitOrder(0, usr, op, pro, "500000000000000000")
	if err != nil {
		t.Fatal(err)
	}
	err = common.InitOrder(1, usr, op, pro, "500000000000000000")
	if err != nil {
		t.Fatal(err)
	}
	err = common.InitOrder(2, usr, op, pro, "500000000000000000")
	if err != nil {
		t.Fatal(err)
	}

	// pay
	fmt.Println("-> pay 0.1 eth: check with nonce 0 should be enough")
	n, err := common.Pay(usr, pro, "100000000000000000")
	if err != nil {
		t.Fatal(err)
	}
	if n != 0 {
		t.Fatalf("nonce %v picked, but should be 0", n)
	}

	fmt.Println("-> pay 0.2 eth: check with nonce 0 should be enough")
	n, err = common.Pay(usr, pro, "200000000000000000")
	if err != nil {
		t.Fatal(err)
	}
	if n != 0 {
		t.Fatalf("nonce %v picked, but should be 0", n)
	}

	fmt.Println("-> pay 0.4 eth: nonce 0 is not enough, nonce 1 should be used")
	n, err = common.Pay(usr, pro, "400000000000000000")
	if err != nil {
		t.Fatal(err)
	}
	if n != 1 {
		t.Fatalf("nonce %v picked, but should be 1", n)
	}

	fmt.Println("-> pay 0.2 eth: nonce 0 should be enough again(0.2 remained)")
	n, err = common.Pay(usr, pro, "200000000000000000")
	if err != nil {
		t.Fatal(err)
	}
	if n != 0 {
		t.Fatalf("nonce %v picked, but should be 0", n)
	}

	fmt.Println("-> pay 0.6 eth: no check is enough(0.5 max), nil paycheck expected")
	_, err = common.Pay(usr, pro, "600000000000000000")
	if err.Error() != "user: usable paycheck not found" {
		t.Fatal(errors.New("no paycheck should be found with enough money"))
	}

	// withdraw
	fmt.Println("-> withdraw with nonce 0")
	n, err = common.Withdraw(op, pro)
	if err != nil {
		t.Fatal(err)
	}
	if n != 0 {
		t.Fatalf("nonce %v picked, but should be 0", n)
	}

	fmt.Println("-> withdraw with nonce 1")
	n, err = common.Withdraw(op, pro)
	if err != nil {
		t.Fatal(err)
	}
	if n != 1 {
		t.Fatalf("nonce %v picked, but should be 1", n)
	}

	//Pay at withdraw
	fmt.Println("-> pay 0.1 eth: check with nonce 2 should be selected now")
	n, err = common.Pay(usr, pro, "100000000000000000")
	if err != nil {
		t.Fatal(err)
	}
	if n != 2 {
		t.Fatalf("nonce %v picked, but should be 2", n)
	}

	fmt.Println("Now test pay at withdraw:")
	fmt.Println("-> pay 0.2 eth: nonce 2 should be selected")
	// pay: 0.2 eth
	userPC, err := usr.Pay(pro.ProviderAddr, utils.String2BigInt(("200000000000000000")))
	if err != nil {
		t.Fatal(err)
	}
	if userPC == nil || userPC.Check.Nonce != 2 {
		t.Fatal("test pay 0.2 failed")
	} else {
		fmt.Println("OK- payer nonce: ", userPC.Check.Nonce)
	}

	fmt.Println("-> Paycheck created, but provider not received yet.")
	fmt.Println("-> Now provider withdrawed this paycheck before latest paycheck received and verified.")

	fmt.Println("-> third withdraw(before verify): nonce 2 expected")
	n, err = common.Withdraw(op, pro)
	if err != nil {
		t.Fatal(err)
	}
	if n != 2 {
		t.Fatalf("nonce %v picked, but should be 2", n)
	}

	// simulate provider receive paycheck from user
	proPC := new(check.Paycheck)
	*proPC = *userPC

	fmt.Println("now verify the latest paycheck, should fail with nonce error.")
	_, err = pro.Verify(proPC, utils.String2BigInt(("200000000000000000")))
	if err.Error() != "nonce should not less than contract nonce" {
		t.Fatal("nonce error should be detected.")
	}
	fmt.Println("OK- verify failed with error nonce")
}
