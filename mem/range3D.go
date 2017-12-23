package spatial

import (
	. "github.com/noypi/spatial/common"
)

func (this *Spatial3D) AddRange(x, y, z Range, v interface{}) {
	vwrap := &_valuewrap{v: v}
	this.xy.AddRange(x, y, vwrap)
	this.z.AddRange(z, vwrap)
}

func (this *Spatial3D) Contains(x, y, z float64) Enum {
	return this.ContainsRange(Range{x, x}, Range{y, y}, Range{z, z})
}

func (this *Spatial3D) ContainsRange(x, y, z Range) Enum {
	return this.searchRange(x, y, z, this.xy.ContainsRange, this.z.ContainsRange)
}

func (this *Spatial3D) WithinRange(x, y, z Range) Enum {
	return this.searchRange(x, y, z, this.xy.WithinRange, this.z.WithinRange)
}

func (this *Spatial3D) searchRange(x, y, z Range, fnXY _search2Dfunc, fnZ _search1Dfunc) Enum {
	oEnum := &_Enum{ch: make(chan Item, 0)}

	go func() {
		m2 := map[*_valuewrap]struct{}{}
		e := fnXY(x, y)
		for {
			v2, has := e.Next()
			if !has {
				break
			}
			m2[v2.(*_valuewrap)] = struct{}{}
		}

		e = fnZ(z.Min, z.Max)
		for {
			v3, has := e.Next()
			if !has {
				break
			}
			if _, has := m2[v3.(*_valuewrap)]; has {
				oEnum.ch <- v3
			}
		}
		oEnum.Close()
	}()

	return oEnum
}
