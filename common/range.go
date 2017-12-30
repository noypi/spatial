package spatial

import (
	"math"
)

type Range struct {
	Min, Max uint64
}

func (a Range) Compare(b Range) int {
	cmp := cmpUint64(a.Min, b.Min)
	if 0 == cmp {
		cmp = cmpUint64(a.Max, b.Max)
	}

	return cmp
}

func (a Range) IsLessOrEqual(b Range) bool {
	cmp := a.Compare(b)
	if cmp < 0 {
		return true
	}
	return 0 == cmp
}

func (a *Range) MaximizeIfZeroMax() {
	if 0 == a.Max {
		a.Max = math.MaxUint64
	}
}
