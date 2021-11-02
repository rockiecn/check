package provider

import (
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/rockiecn/check/internal/check"
	"github.com/rockiecn/check/internal/utils"
)

func TestVerifyOK(t *testing.T) {
	input := &check.Paycheck{
		Check: &check.Check{
			Value:        utils.String2BigInt("100000000000000000000"),
			TokenAddr:    common.HexToAddress("b213d01542d129806d664248a380db8b12059061"),
			Nonce:        6,
			FromAddr:     common.HexToAddress("Ab8483F64d9C6d1EcF9b849Ae677dD3315835cb2"),
			ToAddr:       common.HexToAddress("4B20993Bc481177ec7E8f571ceCaE8A9e22C02db"),
			OpAddr:       common.HexToAddress("5B38Da6a701c568545dCfcB03FcB875f56beddC4"),
			ContractAddr: common.HexToAddress("1c91347f2A44538ce62453BEBd9Aa907C662b4bD"),
			CheckSig:     utils.String2Byte("0e4f125c12d47a91508494d95e710476a7a0c97ed3ce9903ab3df77614de251156b9cbb50ab7bc73fea5ee287a8c1283b02a1eda5b10bc8022f25ea571f68a6801"),
		},
		PayValue:    utils.String2BigInt("1000000000000000000"),
		PaycheckSig: utils.String2Byte("b87d34cbb5ce832d8f3e6533fde6140d3e4562428eb0fa9e10dc1b29230a03401051d928f9a2f8ca0cf390e44449d7f83bf58e6003489d5d61ede2e2ad86990801"),
	}

	prov, err := NewProvider("cc6d63f85de8fef05446ebdd3c537c72152d0fc437fd7aa62b3019b79bd1fdd4")
	if err != nil {
		t.Error(err)
	}

	ok, err := prov.Verify(input, nil)
	if err != nil {
		t.Error(err)
	}
	if ok != true {
		t.Error("should be ok")
	}

}

/*
func TestSendTx(t *testing.T) {
	// use ethClient.BalanceAt() to get the old balance of provider
	// call testWithdraw, mine a block to enable the tx
	// check balance of provider again
	// calc:
	// 1. newBalance > oldBalance
	// 2. old+paycheck.payvalue + gaslimit(9000000) > new
	pro, err := NewProvider("cc6d63f85de8fef05446ebdd3c537c72152d0fc437fd7aa62b3019b79bd1fdd4")
	if err != nil {
		fmt.Println("new provider failed:", err)
		return
	}

	pc := &check.Paycheck{
		Check: &check.Check{
			Value:        utils.String2BigInt("10000000000000000000"),
			TokenAddr:    common.HexToAddress("0xb213d01542d129806d664248A380Db8B12059061"),
			Nonce:        0,
			FromAddr:     common.HexToAddress("0xAb8483F64d9C6d1EcF9b849Ae677dD3315835cb2"),
			ToAddr:       common.HexToAddress("0x4B20993Bc481177ec7E8f571ceCaE8A9e22C02db"),
			OpAddr:       common.HexToAddress("0x5B38Da6a701c568545dCfcB03FcB875f56beddC4"),
			ContractAddr: common.HexToAddress("0xb31BA5cDC07A2EaFAF77c95294fd4aE27D04E9CA"),
			CheckSig:     utils.String2Byte("456391994ed07bd03809b03a7afd9cdd4554c4a6c382289f7e4ea6c749afa7014937817cd32c0c8f983375dbae9959756e82f855538375d0bb99813a0770323500"),
		},
		PayValue:    utils.String2BigInt("1000000000000000000"),
		PaycheckSig: utils.String2Byte("2a8488c7a44fefe771b9fcc1c0ab801f95611906f12e58c4fe2e2ac9406dc97c1601c2ae6f328c4a18df90c1c33170872cab29e1ef970644b163d117a9819ac100"),
	}

	// view balance
	ethClient, err := utils.GetClient(pro.(*Provider).Host)
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
	plusGas = plusGas.Add(bal, utils.String2BigInt("9000000"))
	total := new(big.Int)
	total = total.Add(plusGas, pc.PayValue)
	// require: new < old + paycheck.payvalue + gaslimit(9000000)
	if newbal.Cmp(total) > 0 {
		t.Errorf("new balance should smaller than old balance + payvalue + gaslimit")
	}
}
*/
