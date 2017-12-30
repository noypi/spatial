package geo

import (
	"bytes"
	"encoding/gob"

	"github.com/golang/geo/s2"
	. "github.com/noypi/spatial/common"
)

func init() {
	gob.Register(Region{})
	gob.Register(s2.Loop{})
}

type Region struct {
	Loop *s2.Loop
	V    interface{}
}

type RegionItem struct {
	*Region
	Item
}

func (this *Region) Marshal() (bb []byte, err error) {
	buf := new(bytes.Buffer)
	if err = gob.NewDecoder(buf).Decode(this); nil != err {
		return
	}
	bb = buf.Bytes()
	return
}

func UnmarshalItem(bb []byte) (item *Region, err error) {
	item = new(Region)
	buf := bytes.NewBuffer(bb)
	if err = gob.NewDecoder(buf).Decode(item); nil != err {
		return
	}
	return
}
