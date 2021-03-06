package test

import (
	"fmt"
	"testing"

	"github.com/rockiecn/check/internal/common"
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
	op, err := common.InitOperator("./op/order.db", "./op/check.db")
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println("-> Init User")
	usr, err := common.InitUser("./usr/pc.db")
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println("-> Init 3 Providers ")
	pro0, err := common.InitPro("./pro0/pc.db", "./pro0/bt.db")
	if err != nil {
		t.Fatal(err)
	}
	pro1, err := common.InitPro("./pro1/pc.db", "./pro1/bt.db")
	if err != nil {
		t.Fatal(err)
	}
	pro2, err := common.InitPro("./pro2/pc.db", "./pro2/bt.db")
	if err != nil {
		t.Fatal(err)
	}

	op.ChkStorer.Clear()
	op.OdrStorer.Clear()
	usr.PcStorer.Clear()
	pro0.BtStorer.Clear()
	pro0.PcStorer.Clear()
	pro1.BtStorer.Clear()
	pro1.PcStorer.Clear()
	pro2.BtStorer.Clear()
	pro2.PcStorer.Clear()

	// init 3 orders
	fmt.Println("-> Init 3 orders, value: 0.003")
	err = common.InitOrder(0, usr, op, pro0, "3000000000000000")
	if err != nil {
		t.Fatal(err)
	}
	err = common.InitOrder(1, usr, op, pro1, "3000000000000000")
	if err != nil {
		t.Fatal(err)
	}
	err = common.InitOrder(2, usr, op, pro2, "3000000000000000")
	if err != nil {
		t.Fatal(err)
	}

	// pay
	fmt.Println("-> pay 0.001 eth to provider0 with nonce 0")
	n, err := common.Pay(usr, pro0, "1000000000000000")
	if err != nil {
		t.Fatal(err)
	}
	if n != 0 {
		t.Fatalf("nonce %v picked, but should be 0", n)
	}

	fmt.Println("-> pay 0.002 eth to provider1 with nonce 0")
	n, err = common.Pay(usr, pro1, "2000000000000000")
	if err != nil {
		t.Fatal(err)
	}
	if n != 0 {
		t.Fatalf("nonce %v picked, but should be 0", n)
	}

	fmt.Println("-> pay 0.003 eth to provider2 with nonce 0")
	n, err = common.Pay(usr, pro2, "3000000000000000")
	if err != nil {
		t.Fatal(err)
	}
	if n != 0 {
		t.Fatalf("nonce %v picked, but should be 0", n)
	}

	// withdraw
	fmt.Println("-> provider0 withdraw")
	n, err = common.Withdraw(op, pro0)
	if err != nil {
		t.Fatal(err)
	}
	if n != 0 {
		t.Fatalf("nonce %v picked, but should be 0", n)
	}

	fmt.Println("-> provider1 withdraw")
	n, err = common.Withdraw(op, pro1)
	if err != nil {
		t.Fatal(err)
	}
	if n != 0 {
		t.Fatalf("nonce %v picked, but should be 0", n)
	}

	fmt.Println("-> provider2 withdraw")
	n, err = common.Withdraw(op, pro2)
	if err != nil {
		t.Fatal(err)
	}
	if n != 0 {
		t.Fatalf("nonce %v picked, but should be 0", n)
	}
}
