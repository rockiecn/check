package provider

import (
	"context"
	"fmt"
	"math/big"
	"testing"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/rockiecn/check/check"
	comn "github.com/rockiecn/check/common"
)

func TestSendTx(t *testing.T) {
	// use ethClient.BalanceAt() to get the old balance of provider
	// call testWithdraw, mine a block to enable the tx
	// check balance of provider again
	// calc:
	// 1. newBalance > oldBalance
	// 2. old+paycheck.payvalue + gaslimit(9000000) > new
	pro, err := New("cc6d63f85de8fef05446ebdd3c537c72152d0fc437fd7aa62b3019b79bd1fdd4")
	if err != nil {
		fmt.Println("new provider failed:", err)
		return
	}

	pc := &check.Paycheck{
		Check: check.Check{
			Value:        comn.String2BigInt("10000000000000000000"),
			TokenAddr:    common.HexToAddress("0xb213d01542d129806d664248A380Db8B12059061"),
			Nonce:        0,
			FromAddr:     common.HexToAddress("0xAb8483F64d9C6d1EcF9b849Ae677dD3315835cb2"),
			ToAddr:       common.HexToAddress("0x4B20993Bc481177ec7E8f571ceCaE8A9e22C02db"),
			OpAddr:       common.HexToAddress("0x5B38Da6a701c568545dCfcB03FcB875f56beddC4"),
			ContractAddr: common.HexToAddress("0xb31BA5cDC07A2EaFAF77c95294fd4aE27D04E9CA"),
			CheckSig:     comn.String2Byte("456391994ed07bd03809b03a7afd9cdd4554c4a6c382289f7e4ea6c749afa7014937817cd32c0c8f983375dbae9959756e82f855538375d0bb99813a0770323500"),
		},
		PayValue:    comn.String2BigInt("1000000000000000000"),
		PaycheckSig: comn.String2Byte("2a8488c7a44fefe771b9fcc1c0ab801f95611906f12e58c4fe2e2ac9406dc97c1601c2ae6f328c4a18df90c1c33170872cab29e1ef970644b163d117a9819ac100"),
	}

	// view balance
	ethClient, err := comn.GetClient(pro.(*Provider).Host)
	if err != nil {
		return
	}
	defer ethClient.Close()

	// get old balance of provider
	bal, err := ethClient.BalanceAt(context.Background(), common.HexToAddress("0x4B20993Bc481177ec7E8f571ceCaE8A9e22C02db"), nil)
	if err != nil {
		fmt.Println("get balance failed")
	}
	fmt.Println("old balance of provider:", bal.String())

	// call withdraw
	tx, _ := pro.SendTx(pc)
	// deploy contract, wait for mining.
	for {
		txReceipt, _ := ethClient.TransactionReceipt(context.Background(), tx.Hash())
		// receipt ok
		if txReceipt != nil {
			break
		}
		fmt.Println("wait mining")
		time.Sleep(time.Duration(3) * time.Second)
	}

	newbal, err := ethClient.BalanceAt(context.Background(), common.HexToAddress("0x4B20993Bc481177ec7E8f571ceCaE8A9e22C02db"), nil)
	if err != nil {
		fmt.Println("get balance failed")
	}
	fmt.Println("new balance of provider:", newbal.String())

	// require: newBalance > oldBalance
	if newbal.Cmp(bal) < 0 {
		t.Errorf("new balance should larger than old balance")
	}

	plusGas := new(big.Int)
	plusGas = plusGas.Add(bal, comn.String2BigInt("9000000"))
	total := new(big.Int)
	total = total.Add(plusGas, pc.PayValue)
	// require: new < old+paycheck.payvalue + gaslimit(9000000)
	if newbal.Cmp(total) > 0 {
		t.Errorf("new balance should smaller than old balance + payvalue + gaslimit")
	}
}
