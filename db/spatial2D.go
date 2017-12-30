package spatial

type Spatial2D struct {
	x, y *Spatial1D
}

func New2D() *Spatial2D {
	o := new(Spatial2D)
	o.x, o.y = New1D(), New1D()
	o.y.xyzOffset = 1
	return o
}
