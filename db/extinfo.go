package spatial

import (
	"bytes"
)

func (this *Spatial1D) SetExtInfo(category uint8, id []byte, v interface{}) error {
	vbb, err := serializeRaw(v)
	if nil != err {
		return err
	}

	buf := new(bytes.Buffer)
	buf.WriteByte(cPrefixExtInfo)
	buf.Write(id)
	buf.WriteByte(category)

	wrtr, err := this.extinfo.Writer()
	if nil != err {
		return err
	}
	batch := wrtr.NewBatch()
	batch.Set(buf.Bytes(), vbb)
	return wrtr.ExecuteBatch(batch)
}

func (this *Spatial1D) GetExtInfo(category uint8, id []byte) (v interface{}, err error) {
	rdr, err := this.extinfo.Reader()
	if nil != err {
		return
	}
	buf := new(bytes.Buffer)
	buf.WriteByte(cPrefixExtInfo)
	buf.Write(id)
	buf.WriteByte(category)

	bb, err := rdr.Get(buf.Bytes())
	if nil != err {
		return
	}

	return deserializeRaw(bb)
}

func (this *Spatial2D) SetExtInfo(category uint8, id []byte, v interface{}) error {
	return this.x.SetExtInfo(category, id, v)
}

func (this *Spatial2D) GetExtInfo(category uint8, id []byte) (v interface{}, err error) {
	return this.x.GetExtInfo(category, id)
}

func (this *Spatial3D) SetExtInfo(category uint8, id []byte, v interface{}) error {
	return this.xy.SetExtInfo(category, id, v)
}

func (this *Spatial3D) GetExtInfo(category uint8, id []byte) (v interface{}, err error) {
	return this.xy.GetExtInfo(category, id)
}
