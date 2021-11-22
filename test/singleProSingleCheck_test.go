package test

import (
	"fmt"
	"testing"

	"github.com/rockiecn/check/internal/common"
)

// a single provider involved
// a single check is used
// multi pay actions by user, and 1 withdraw action by provider
func TestSingleProSingleCheck(t *testing.T) {

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
	fmt.Println("-> Init User")
	usr, err := common.InitUser()
	if err != nil {
		t.Fatal(err)
	}

	fmt.Println("-> Init order")
	err = common.InitOrder(0, usr, op, pro, "1500000000000000000")
	if err != nil {
		t.Fatal(err)
	}

	// pay
	fmt.Println("-> user pay 0.1 eth to provider with nonce 0")
	err = common.Pay(usr, pro, "100000000000000000", 0)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println("-> user pay 0.2 eth to provider with nonce 0")
	err = common.Pay(usr, pro, "200000000000000000", 0)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println("-> user pay 0.3 eth to provider with nonce 0")
	err = common.Pay(usr, pro, "300000000000000000", 0)
	if err != nil {
		t.Fatal(err)
	}

	// withdraw
	fmt.Println("-> withdraw with nonce 0")
	err = common.Withdraw(op, pro, 0)
	if err != nil {
		t.Fatal(err)
	}
}
