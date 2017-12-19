package spatial

type Point2D struct {
	X, Y float64
}

func (this *Spatial2D) AddRange(x, y Range, v interface{}) {
	x.maximizeIfZeroMax()
	y.maximizeIfZeroMax()
	this.m.Put(&_item2D{x: x, y: y}, v)
}

func (this *Spatial2D) Get(x, y Range) *Enum {
	oEnum := &Enum{ch: make(chan interface{}, 0)}
	x.maximizeIfZeroMax()
	y.maximizeIfZeroMax()

	go func() {
		x.Min -= Epsilonx10
		y.Min -= Epsilonx10
		this.m.EachFrom(&_item2D{x, y}, func(k, v interface{}) bool {
			k1 := k.(*_item2D)
			bValid := IsLessOrEqual(k1.x.Max, x.Max) && IsLessOrEqual(k1.y.Max, y.Max)
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
