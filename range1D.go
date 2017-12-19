package spatial

import (
	"math"
)

func (this *Spatial1D) AddRange(r Range, v interface{}) {
	r.maximizeIfZeroMax()
	this.m.Put(&r, v)
}

func (this *Spatial1D) Get(min, max float64) *Enum {
	oEnum := &Enum{ch: make(chan interface{}, 0)}
	if max <= 0 {
		max = float64(math.MaxFloat64)
	}
	go func() {
		this.m.EachFrom(&Range{min - Epsilonx10, max}, func(k, v interface{}) bool {
			k1 := k.(*Range)
			bValid := IsLessOrEqual(k1.Max, max)
			if !bValid {
				oEnum.Close()
				return false
			}

			oEnum.ch <- v
			return true
		})
		oEnum.Close()
	}()

	return oEnum
}
