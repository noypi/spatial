package spatial

func (this *Spatial3D) AddRange(x, y, z Range, v interface{}) {
	this.xy.AddRange(x, y, v)
	this.z.AddRange(z, v)
}

func (this *Spatial3D) Contains(x, y, z float64) *Enum {
	return this.ContainsRange(Range{x, x}, Range{y, y}, Range{z, z})
}

func (this *Spatial3D) ContainsRange(x, y, z Range) *Enum {
	return this.searchRange(x, y, z, this.xy.ContainsRange, this.z.ContainsRange)
}

func (this *Spatial3D) WithinRange(x, y, z Range) *Enum {
	return this.searchRange(x, y, z, this.xy.WithinRange, this.z.WithinRange)
}

func (this *Spatial3D) searchRange(x, y, z Range, fnXY _search2Dfunc, fnZ _search1Dfunc) *Enum {
	oEnum := &Enum{ch: make(chan *_Item, 0)}

	go func() {
		m2 := map[string]struct{}{}
		e := fnXY(x, y)
		for {
			v2, has := e.Next()
			if !has {
				break
			}
			m2[v2.id.String()] = struct{}{}
		}

		e = fnZ(z.Min, z.Max)
		for {
			v3, has := e.Next()
			if !has {
				break
			}
			if _, has := m2[v3.id.String()]; has {
				oEnum.ch <- v3
			}
		}
		oEnum.Close()
	}()

	return oEnum
}
