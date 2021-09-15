package recorder

import (
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/rockiecn/check/check"
)

// recorder for paycheck
type PRecorder struct {
	Data map[common.Address]map[common.Address]map[*big.Int]check.Paycheck
}

func NewPRecorder() *PRecorder {
	prec := new(PRecorder)
	return prec
}

// record a paycheck into data pool.
func (prec *PRecorder) Record(pchk *check.Paycheck) {

	op := pchk.Check.OpAddr
	pro := pchk.Check.ToAddr
	nonce := pchk.Check.Nonce

	if prec.Data == nil {
		opMap := make(map[common.Address]map[common.Address]map[*big.Int]check.Paycheck)
		prec.Data = opMap
	}

	if prec.Data[op] == nil {
		proMap := make(map[common.Address]map[*big.Int]check.Paycheck)
		prec.Data[op] = proMap
	}

	if prec.Data[op][pro] == nil {
		nonMap := make(map[*big.Int]check.Paycheck)
		prec.Data[op][pro] = nonMap
	}

	prec.Data[op][pro][nonce] = *pchk
}

func (rec *PRecorder) List() {
	for opaddr, proMap := range rec.Data {
		fmt.Println("---> operator:", opaddr)
		for proaddr, nonceMap := range proMap {
			fmt.Println("--> provider:", proaddr)
			for nonce, pc := range nonceMap {
				fmt.Println("->nonce:", nonce)
				fmt.Println("Value:", pc.Check.Value)
				fmt.Println("TokenAddr:", pc.Check.TokenAddr)
				fmt.Println("From:", pc.Check.FromAddr)
				fmt.Println("To:", pc.Check.ToAddr)
				fmt.Println("OperatorAddr:", pc.Check.OpAddr)
				fmt.Println("PayValue:", pc.PayValue)
			}
		}
		fmt.Println()
	}
}

// recorder for check
type CRecorder struct {
	Data map[common.Address]map[*big.Int]check.Check
}

func NewCRecorder() *CRecorder {
	crec := new(CRecorder)
	return crec
}

// record a paycheck into data pool.
func (crec *CRecorder) Record(chk *check.Check) {

	pro := chk.ToAddr
	nonce := chk.Nonce

	if crec.Data == nil {
		proNon := make(map[common.Address]map[*big.Int]check.Check)
		crec.Data = proNon
	}

	if crec.Data[pro] == nil {
		nonChk := make(map[*big.Int]check.Check)
		crec.Data[pro] = nonChk
	}

	crec.Data[pro][nonce] = *chk
}
