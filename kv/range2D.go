package spatial

func (this *Spatial2D) AddRange(x, y Range, v interface{}) {
	item := NewItem(v)
	this.x.AddRange(x, item)
	this.y.AddRange(y, item)
}

func (this *Spatial2D) Contains(x, y float64) *Enum {
	return this.ContainsRange(Range{x, x}, Range{y, y})
}

func (this *Spatial2D) ContainsRange(x, y Range) *Enum {
	return this.searchRange(x, y, this.x.ContainsRange, this.y.ContainsRange)
}

func (this *Spatial2D) WithinRange(x, y Range) *Enum {
	return this.searchRange(x, y, this.x.WithinRange, this.y.WithinRange)
}

func (this *Spatial2D) searchRange(x, y Range, fnX, fnY _search1Dfunc) *Enum {
	oEnum := &Enum{ch: make(chan *_Item, 0)}

	go func() {
		m1 := map[string]struct{}{}
		e := fnX(x.Min, x.Max)
		for {
			v1, has := e.Next()
			if !has {
				break
			}
			m1[v1.id.String()] = struct{}{}
		}
		e = fnY(y.Min, y.Max)
		for {
			v2, has := e.Next()
			if !has {
				break
			}
			if _, has := m1[v2.id.String()]; has {
				oEnum.ch <- v2
			}
		}
		oEnum.Close()
	}()

	return oEnum
}
