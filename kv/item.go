package spatial

import (
	. "github.com/noypi/spatial/common"
	"github.com/rs/xid"
)

type _Item struct {
	V      interface{}
	Ranges []Range
	err    error
	id     xid.ID
}

func (this _Item) Error() error {
	return this.err
}

func (this _Item) Range(n int) Range {
	return this.Ranges[n]
}

func (this _Item) Value() interface{} {
	return this.V
}

func (this _Item) ID() string {
	return this.id.String()
}
