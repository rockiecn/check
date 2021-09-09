package provider

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/rockiecn/check/check"
)

type Provider struct {
	ProviderSK string

	PayChecks map[string]*check.PayCheck // (operator,nonce) to paycheck
}

type IProvider interface {
	NewProvider(sk string) (*Provider, error)
	VerifyPayCheck(paycheck check.PayCheck, sig []byte, userAddr common.Address) (bool, error)
	WithDraw(paycheque *check.PayCheck) error
}
