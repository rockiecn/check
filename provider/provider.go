package provider

import (
	"github.com/rockiecn/check/check"
	"github.com/rockiecn/check/utils"
)

type Provider struct {
	ProviderSK   string
	ProviderAddr string

	//
	History map[string]*check.PayCheck // keyHash -> key, paycheck, key: "operator:xxx, provider:xxx, nonce:xxx"
}

type IProvider interface {
	NewProvider(sk string) (*Provider, error)
	VerifyPayCheck(paycheck check.PayCheck) (bool, error)
	RecordPayCheck(check *check.PayCheck) error
	WithDraw(paycheque *check.PayCheck) error
}

func NewProvider(sk string) *Provider {
	pro := new(Provider)
	pro.ProviderSK = sk
	pro.ProviderAddr = utils.KeyToAddr(sk)

	return pro
}

func (pro *Provider) VerifyPayCheck(paycheck check.PayCheck) (bool, error) {

	return true, nil

}
