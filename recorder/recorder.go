package recorder

import (
	"errors"
	"reflect"

	"github.com/ethereum/go-ethereum/common"
	"github.com/rockiecn/check/check"
)

type Key struct {
	Operator common.Address
	Provider common.Address
	Nonce    uint64
}

// recorder for paycheck
type Recorder struct {
	Entrys map[Key]interface{}
}

func New() *Recorder {

	r := &Recorder{
		Entrys: make(map[Key]interface{}),
	}

	return r
}

// put a check/paycheck into Entrys
func (r *Recorder) Record(entry interface{}) error {

	t := reflect.TypeOf(entry)

	switch t.String() {
	case "*check.Check":
		if c, ok := entry.(*check.Check); ok {
			key := Key{
				Operator: c.OpAddr,
				Provider: c.ToAddr,
				Nonce:    c.Nonce,
			}
			r.Entrys[key] = c
			return nil
		}
		return errors.New("check type assertion failed")
	case "*check.Paycheck":
		if pc, ok := entry.(*check.Paycheck); ok {
			key := Key{
				Operator: pc.Check.OpAddr,
				Provider: pc.Check.ToAddr,
				Nonce:    pc.Check.Nonce,
			}
			r.Entrys[key] = pc
			return nil
		}
		return errors.New("paycheck type assertion failed")
	default:
		return errors.New("type of value must be Check or Paycheck")
	}
}

// check if a check/paycheck is valid to store
func (r *Recorder) IsValid(entry interface{}) (bool, error) {

	t := reflect.TypeOf(entry)

	switch t.String() {
	case "*check.Check":
		c := entry.(*check.Check)
		k := Key{
			Operator: c.OpAddr,
			Provider: c.ToAddr,
			Nonce:    c.Nonce,
		}
		v := r.Entrys[k]
		if v == nil {
			return true, nil // check not exist, ok to store
		} else {
			return false, nil // check already exist
		}
	case "*check.Paycheck":
		pc := entry.(*check.Paycheck)
		k := Key{
			Operator: pc.Check.OpAddr,
			Provider: pc.Check.ToAddr,
			Nonce:    pc.Check.Nonce,
		}
		v := r.Entrys[k]
		if v == nil {
			return true, nil // paycheck not exist, ok to store
		} else {
			if pc.PayValue.Cmp(v.(*check.Paycheck).PayValue) > 0 {
				return true, nil // payvalue increased, ok to store
			}
			return false, nil // invalid
		}
	default:
		return false, errors.New("type must be check/paycheck")
	}
}
