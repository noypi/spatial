package spatial

import (
	"fmt"

	. "github.com/noypi/spatial/common"
)

type _valuewrap struct {
	v      interface{}
	err    error
	ranges []Range
}

func (this _valuewrap) Error() error {
	return this.err
}

func (this _valuewrap) Value() interface{} {
	return this.v
}

func (this *_valuewrap) ID() string {
	return fmt.Sprintf("%p", this)
}

func (this _valuewrap) Range(n int) Range {
	return this.ranges[n]
}

func (this *_valuewrap) Delete() {
	panic("TODO")
}

func (this *_valuewrap) Set(v interface{}) error {
	panic("TODO")
}
