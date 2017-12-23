package spatial

import (
	. "github.com/noypi/spatial/common"
)

func (this *Spatial2D) AddRange(x, y Range, v interface{}) {
	vwrap := &_valuewrap{v: v}
	this.x.AddRange(x, vwrap)
	this.y.AddRange(y, vwrap)
}

func (this *Spatial2D) Contains(x, y float64) Enum {
	return this.ContainsRange(Range{x, x}, Range{y, y})
}

func (this *Spatial2D) ContainsRange(x, y Range) Enum {
	return this.searchRange(x, y, this.x.ContainsRange, this.y.ContainsRange)
}

func (this *Spatial2D) WithinRange(x, y Range) Enum {
	return this.searchRange(x, y, this.x.WithinRange, this.y.WithinRange)
}

func (this *Spatial2D) searchRange(x, y Range, fnX, fnY _search1Dfunc) Enum {
	oEnum := &_Enum{ch: make(chan Item, 0)}

	go func() {
		m1 := map[*_valuewrap]struct{}{}
		e := fnX(x.Min, x.Max)
		for {
			v1, has := e.Next()
			if !has {
				break
			}
			m1[v1.(*_valuewrap)] = struct{}{}
		}
		e = fnY(y.Min, y.Max)
		for {
			v2, has := e.Next()
			if !has {
				break
			}
			if _, has := m1[v2.(*_valuewrap)]; has {
				oEnum.ch <- v2
			}
		}
		oEnum.Close()
	}()

	return oEnum
}
