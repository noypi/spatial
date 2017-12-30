package spatial

import (
	. "github.com/noypi/spatial/common"
)

func (this *Spatial2D) Set(id []byte, x, y Range, v interface{}) error {
	item := NewItem(id, v, x, y)
	if err := this.x.Set(id, x, item); nil != err {
		return err
	}
	if err := this.y.Set(id, y, item); nil != err {
		return err
	}

	return nil
}

func (this *Spatial2D) Contains(x, y uint64) Enum {
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
		m1 := map[string]struct{}{}
		e := fnX(x.Min, x.Max)
		for {
			v1, has := e.Next()
			if !has {
				break
			}
			m1[v1.(*_Item).ID()] = struct{}{}
		}
		e = fnY(y.Min, y.Max)
		for {
			v2, has := e.Next()
			if !has {
				break
			}
			if _, has := m1[v2.(*_Item).ID()]; has {
				oEnum.ch <- v2
			}
		}
		oEnum.Close()
	}()

	return oEnum
}
