package test

import (
	"fmt"
	"testing"

	"github.com/rockiecn/check/test/common"
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
	op, err := common.InitOperator()
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println("-> Init Provider")
	pro, err := common.InitPro()
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println("-> Init 3 Users")
	usr0, err := common.InitUser()
	if err != nil {
		t.Fatal(err)
	}
	usr1, err := common.InitUser()
	if err != nil {
		t.Fatal(err)
	}
	usr2, err := common.InitUser()
	if err != nil {
		t.Fatal(err)
	}

	// init 3 orders
	fmt.Println("-> Init 3 orders")
	err = common.InitOrder(0, usr0, op, pro, "300000000000000000")
	if err != nil {
		t.Fatal(err)
	}
	err = common.InitOrder(1, usr1, op, pro, "300000000000000000")
	if err != nil {
		t.Fatal(err)
	}
	err = common.InitOrder(2, usr2, op, pro, "300000000000000000")
	if err != nil {
		t.Fatal(err)
	}

	// pay
	fmt.Println("-> user0 pay 0.1 eth to provider with nonce 0")
	err = common.Pay(usr0, pro, "100000000000000000", 0)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println("-> user1 pay 0.2 eth to provider with nonce 1")
	err = common.Pay(usr1, pro, "200000000000000000", 1)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println("-> user2 pay 0.3 eth to provider with nonce 2")
	err = common.Pay(usr2, pro, "300000000000000000", 2)
	if err != nil {
		t.Fatal(err)
	}

	// withdraw
	fmt.Println("-> withdraw with nonce 0")
	err = common.Withdraw(op, pro, 0)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println("-> withdraw with nonce 1")
	err = common.Withdraw(op, pro, 1)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println("-> withdraw with nonce 2")
	err = common.Withdraw(op, pro, 2)
	if err != nil {
		t.Fatal(err)
	}
}
