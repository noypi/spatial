package spatial

import (
	"math"
)

type Range struct {
	Min, Max float64
}

func (a Range) Compare(b Range) int {
	cmp := a.Min - b.Min
	if IsZero(cmp) {
		cmp = a.Max - b.Max
	}

	if IsZero(cmp) {
		return 0
	} else if cmp < 0 {
		return -1
	}
	return 1
}

func (a Range) IsLessOrEqual(b Range) bool {
	cmp := a.Compare(b)
	if cmp < 0 {
		return true
	}
	return 0 == cmp
}

func (a *Range) MaximizeIfZeroMax() {
	if IsZero(a.Max) {
		a.Max = float64(math.MaxFloat64)
	}
}
