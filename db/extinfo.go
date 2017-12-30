package spatial

import (
	"bytes"
)

func (this *Spatial1D) SetExtInfo(category uint8, id []byte, v interface{}) error {
	vbb, err := serializeRaw(v)
	if nil != err {
		return err
	}

	buf := bytes.NewBuffer([]byte{cPrefixExtInfo})
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

func (this *Spatial2D) SetExtInfo(category uint8, id []byte, v interface{}) error {
	return this.x.SetExtInfo(category, id, v)
}

func (this *Spatial3D) SetExtInfo(category uint8, id []byte, v interface{}) error {
	return this.xy.SetExtInfo(category, id, v)
}
