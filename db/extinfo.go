package spatial

import (
	"bytes"
)

func (this *Spatial1D) SetExtInfo(category uint8, id []byte, v interface{}) error {
	vbb, err := SerializeRaw(v)
	if nil != err {
		return err
	}

	buf := new(bytes.Buffer)
	buf.WriteByte(cPrefixExtInfo)
	buf.Write(id)
	buf.WriteByte(category)

	return this.setExtInfo(buf.Bytes(), vbb)
}

func (this *Spatial1D) SetExtBatchSize(n uint) {
	this.syncExtBatch.Lock()
	this.extBatchCnt = n
	this.syncExtBatch.Unlock()
}

func (this *Spatial1D) setExtInfo(k, v []byte) (err error) {
	this.syncExtBatch.Lock()

	if nil == this.extwrtr {
		this.extwrtr, err = this.extinfo.Writer()
		if nil != err {
			return
		}
	}

	if nil == this.extbatch {
		this.extbatch = this.extwrtr.NewBatch()
	}
	this.extbatch.Set(k, v)
	this.extBatchCnt++

	this.syncExtBatch.Unlock()

	return this.FlushExt()
}

func (this *Spatial1D) FlushExt() (err error) {
	this.syncExtBatch.Lock()
	defer this.syncExtBatch.Unlock()

	if this.extBatchMaxCnt < this.extBatchCnt {
		if err = this.extwrtr.ExecuteBatch(this.extbatch); nil == err {
			this.extBatchCnt = 0
			this.extbatch.Reset()
		}
	}
	return err
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

	return DeserializeRaw(bb)
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
