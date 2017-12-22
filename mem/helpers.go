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
