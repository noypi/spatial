package spatial

type Spatial3D struct {
	xy *Spatial2D
	z  *Spatial1D
}

func New3D() *Spatial3D {
	o := new(Spatial3D)
	o.xy, o.z = New2D(), New1D()
	return o
}
