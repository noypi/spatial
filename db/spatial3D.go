package spatial

type Spatial3D struct {
	xy *Spatial2D
	z  *Spatial1D
}

func New3D(opts ...Options) (o *Spatial3D, err error) {
	o = new(Spatial3D)

	opts = append(opts, optEnableExtInfo{true})
	if o.xy, err = New2D(); nil != err {
		return
	}

	// disable extinfo
	opts[len(opts)-1] = optEnableExtInfo{false}
	if o.z, err = New1D(); nil != err {
		return
	}
	o.z.xyzOffset = 2
	return
}
