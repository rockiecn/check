package operator

import (
	"testing"

	"github.com/rockiecn/check/internal/utils"
)

func TestRestoreOp(t *testing.T) {
	// generate operator
	opSk, err := utils.GenerateSK()
	if err != nil {
		t.Fatal(err)
	}
	op, err := New(opSk)
	if err != nil {
		t.Fatal(err)
	}
	Op, ok := op.(*Operator)
	if !ok {
		t.Fatal("new operator assertion failed")
	}

	// restore order and check into pool
	err = Op.RestoreChk()
	if err != nil {
		t.Fatal(err)
	}
	err = Op.RestoreOrder()
	if err != nil {
		t.Fatal(err)
	}

	// show data in pool
	Op.ShowChk()
	Op.ShowOdr()
}
