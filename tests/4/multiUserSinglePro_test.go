package test

import (
	"fmt"
	"testing"

	"github.com/rockiecn/check/internal/common"
)

// 3 users involved
// 1 provider involved
// 3 checks used
// 3 pay actions by user
// 3 withdraw actions by provider
//
// Process:
// 1. create operator, provider
// 2. create 3 users
// 3. operator create 1 order for each user, all has 0.3 eth
// 4. operator transfer orders to checks and send to users
// 5. users store check into pool
// 6. user0 pay provider 0.1 eth
// 7. user1 pay provider 0.2 eth
// 8. user2 pay provider 0.3 eth
// 9. provider withdraw all 3 paychecks
//    should receive 0.1 eth, 0.2 eth, 0.3 eth respectively
func TestMultiUserSinglePro(t *testing.T) {
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
	fmt.Println("-> Init 3 Users")
	usr0, err := common.InitUser("./usr0/pc.db")
	if err != nil {
		t.Fatal(err)
	}
	usr1, err := common.InitUser("./usr1/pc.db")
	if err != nil {
		t.Fatal(err)
	}
	usr2, err := common.InitUser("./usr2/pc.db")
	if err != nil {
		t.Fatal(err)
	}

	op.ChkStorer.Clear()
	op.OdrStorer.Clear()
	usr0.PcStorer.Clear()
	usr1.PcStorer.Clear()
	usr2.PcStorer.Clear()
	pro.BtStorer.Clear()
	pro.PcStorer.Clear()

	// init 3 orders
	fmt.Println("-> Init 3 orders, valule: 0.003 eth")
	err = common.InitOrder(0, usr0, op, pro, "3000000000000000")
	if err != nil {
		t.Fatal(err)
	}
	err = common.InitOrder(1, usr1, op, pro, "3000000000000000")
	if err != nil {
		t.Fatal(err)
	}
	err = common.InitOrder(2, usr2, op, pro, "3000000000000000")
	if err != nil {
		t.Fatal(err)
	}

	// pay
	fmt.Println("-> user0 pay 0.001 eth to provider with nonce 0")
	n, err := common.Pay(usr0, pro, "1000000000000000")
	if err != nil {
		t.Fatal(err)
	}
	if n != 0 {
		t.Fatalf("nonce %v picked, but should be 0", n)
	}

	fmt.Println("-> user1 pay 0.002 eth to provider with nonce 1")
	n, err = common.Pay(usr1, pro, "2000000000000000")
	if err != nil {
		t.Fatal(err)
	}
	if n != 1 {
		t.Fatalf("nonce %v picked, but should be 1", n)
	}

	fmt.Println("-> user2 pay 0.003 eth to provider with nonce 2")
	n, err = common.Pay(usr2, pro, "3000000000000000")
	if err != nil {
		t.Fatal(err)
	}
	if n != 2 {
		t.Fatalf("nonce %v picked, but should be 2", n)
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

	fmt.Println("-> withdraw with nonce 2")
	n, err = common.Withdraw(op, pro)
	if err != nil {
		t.Fatal(err)
	}
	if n != 2 {
		t.Fatalf("nonce %v picked, but should be 2", n)
	}
}
