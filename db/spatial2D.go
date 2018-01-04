package spatial

type Spatial2D struct {
	x, y *Spatial1D
}

func New2D(opts ...Options) (o *Spatial2D, err error) {
	o = new(Spatial2D)
	opts = append(opts, optEnableExtInfo{true})
	if o.x, err = New1D(opts...); nil != err {
		return
	}

	// disable extinfo
	opts[len(opts)-1] = optEnableExtInfo{false}
	if o.y, err = New1D(opts...); nil != err {
		return
	}
	o.y.xyzOffset = 1
	return
}

func (this *Spatial2D) Close() {
	this.x.Close()
	this.y.Close()
}
