package provider

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/rockiecn/check/check"
)

type Provider struct {
	ProviderSK string

	//
	History map[string]*check.PayCheck // keyHash -> key, paycheck, key: "operator:xxx, provider:xxx, nonce:xxx"
}

type IProvider interface {
	NewProvider(sk string) (*Provider, error)
	VerifyPayCheck(paycheck check.PayCheck, sig []byte, userAddr common.Address) (bool, error)
	WithDraw(paycheque *check.PayCheck) error
}
