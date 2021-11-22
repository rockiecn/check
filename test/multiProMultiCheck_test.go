package test

import (
	"fmt"
	"testing"

	"github.com/rockiecn/check/test/common"
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

	// init roles
	fmt.Println("-> Init Operator")
	op, err := common.InitOperator()
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println("-> Init User")
	usr, err := common.InitUser()
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println("-> Init 3 Providers ")
	pro0, err := common.InitPro()
	if err != nil {
		t.Fatal(err)
	}
	pro1, err := common.InitPro()
	if err != nil {
		t.Fatal(err)
	}
	pro2, err := common.InitPro()
	if err != nil {
		t.Fatal(err)
	}

	// init 3 orders
	fmt.Println("-> Init 3 orders")
	err = common.InitOrder(0, usr, op, pro0, "300000000000000000")
	if err != nil {
		t.Fatal(err)
	}
	err = common.InitOrder(1, usr, op, pro1, "300000000000000000")
	if err != nil {
		t.Fatal(err)
	}
	err = common.InitOrder(2, usr, op, pro2, "300000000000000000")
	if err != nil {
		t.Fatal(err)
	}

	// pay
	fmt.Println("-> pay 0.1 eth to provider0 with nonce 0")
	err = common.Pay(usr, pro0, "100000000000000000", 0)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println("-> pay 0.2 eth to provider1 with nonce 0")
	err = common.Pay(usr, pro1, "200000000000000000", 0)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println("-> pay 0.3 eth to provider2 with nonce 0")
	err = common.Pay(usr, pro2, "300000000000000000", 0)
	if err != nil {
		t.Fatal(err)
	}

	// withdraw
	fmt.Println("-> provider0 withdraw")
	err = common.Withdraw(op, pro0, 0)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println("-> provider1 withdraw")
	err = common.Withdraw(op, pro1, 0)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println("-> provider2 withdraw")
	err = common.Withdraw(op, pro2, 0)
	if err != nil {
		t.Fatal(err)
	}
}
