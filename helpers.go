package spatial

import (
	"math"
)

var Epsilon = math.Nextafter(1, 2) - 1
var Epsilonx10 = Epsilon * 10

func IsZero(a float64) bool {
	return math.Abs(a) < Epsilon
}

func IsLessOrEqual(a, b float64) bool {
	if a < b {
		return true
	}

	return IsZero(a - b)
}

type _valuewrap struct {
	v interface{}
}

type _search1Dfunc func(x, y float64) *Enum
type _search2Dfunc func(x, y Range) *Enum
