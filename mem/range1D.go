package spatial

import (
	"math"

	. "github.com/noypi/spatial/common"
)

func (this *Spatial1D) AddRange(r Range, v interface{}) {
	var item Item
	if o, ok := v.(Item); ok {
		item = o
	} else {
		item = &_valuewrap{v: v}
	}
	r.MaximizeIfZeroMax()
	this.m.Put(&r, item)
	this.mrev.Put(&r, item)
}

func (this *Spatial1D) Contains(x float64) Enum {
	return this.ContainsRange(x, x)
}

func (this *Spatial1D) ContainsRange(min, max float64) Enum {
	if IsLessOrEqual(max, min) {
		max = min - Epsilonx10
	}

	oEnum := &_Enum{ch: make(chan Item, 0)}
	go func() {
		this.mrev.EachFrom(&Range{min + Epsilonx10, max - Epsilonx10}, func(k, v interface{}) bool {
			k1 := k.(*Range)
			bValid := IsLessOrEqual(k1.Min, min) && IsLessOrEqual(max, k1.Max)
			if !bValid {
				oEnum.Close()
				return false
			}

			oEnum.ch <- v.(Item)
			return true
		})
		oEnum.Close()
	}()
	return oEnum
}

func (this *Spatial1D) WithinRange(min, max float64) Enum {
	oEnum := &_Enum{ch: make(chan Item, 0)}
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

			oEnum.ch <- v.(Item)
			return true
		})
		oEnum.Close()
	}()

	return oEnum
}
