package spatial

type Spatial2D struct {
	x *Spatial1D
	y *Spatial1D
}

func New2D() *Spatial2D {
	o := new(Spatial2D)
	o.x, o.y = New1D(), New1D()
	return o
}
