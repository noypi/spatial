package spatial

import (
	. "github.com/noypi/spatial/common"
)

func (this *Spatial3D) Set(id []byte, x, y, z Range, v interface{}) error {
	if err := this.xy.Set(id, x, y, v); nil != err {
		return err
	}
	if err := this.z.Set(id, z, v); nil != err {
		return err
	}
	return nil
}

func (this *Spatial3D) Contains(x, y, z uint64) Enum {
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
		m2 := map[string]struct{}{}
		e := fnXY(x, y)
		for {
			v2, has := e.Next()
			if !has {
				break
			}
			m2[v2.(*_Item).ID()] = struct{}{}
		}

		e = fnZ(z.Min, z.Max)
		for {
			v3, has := e.Next()
			if !has {
				break
			}
			if _, has := m2[v3.(*_Item).ID()]; has {
				oEnum.ch <- v3
			}
		}
		oEnum.Close()
	}()

	return oEnum
}
