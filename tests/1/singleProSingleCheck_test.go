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

	// init roles with db
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

	fmt.Println("contract address:", op.CtrAddr)

	op.ChkStorer.Clear()
	op.OdrStorer.Clear()
	usr.PcStorer.Clear()
	pro.BtStorer.Clear()
	pro.PcStorer.Clear()

	fmt.Println("-> Init order, value: 0.15 eth")
	err = common.InitOrder(0, usr, op, pro, "150000000000000000")
	if err != nil {
		t.Fatal(err)
	}

	// pay
	fmt.Println("-> user pay 0.01 eth to provider with nonce 0")
	n, err := common.Pay(usr, pro, "10000000000000000")
	if err != nil {
		t.Fatal(err)
	}
	if n != 0 {
		t.Fatalf("nonce %v picked, but should be 0", n)
	}

	fmt.Println("-> user pay 0.02 eth to provider with nonce 0")
	n, err = common.Pay(usr, pro, "20000000000000000")
	if err != nil {
		t.Fatal(err)
	}
	if n != 0 {
		t.Fatalf("nonce %v picked, but should be 0", n)
	}

	fmt.Println("-> user pay 0.03 eth to provider with nonce 0")
	n, err = common.Pay(usr, pro, "30000000000000000")
	if err != nil {
		t.Fatal(err)
	}
	if n != 0 {
		t.Fatalf("nonce %v picked, but should be 0", n)
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
}
