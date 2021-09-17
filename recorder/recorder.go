package recorder

import (
	"errors"

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

	rec := &Recorder{
		Entrys: make(map[Key]interface{}),
	}
	return rec
}

func (rec *Recorder) Record(entry interface{}) error {

	if c, ok := entry.(*check.Check); ok {
		key := Key{
			Operator: c.OpAddr,
			Provider: c.ToAddr,
			Nonce:    c.Nonce,
		}
		rec.Entrys[key] = c
		return nil
	}

	if pc, ok := entry.(*check.Paycheck); ok {
		key := Key{
			Operator: pc.Check.OpAddr,
			Provider: pc.Check.ToAddr,
			Nonce:    pc.Check.Nonce,
		}
		rec.Entrys[key] = pc
		return nil
	}

	return errors.New("entry type error, must be check or paycheck")

}
