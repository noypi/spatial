package osm

import (
	"archive/tar"
	"compress/bzip2"
	"encoding/gob"
	"fmt"
	"io"
	"log"

	"github.com/golang/geo/s2"
	"github.com/noypi/spatial/db"
	"github.com/thomersch/gosmparse"
)

func init() {
	gob.Register(s2.LatLng{})
	gob.Register([]s2.LatLng{})
}

func (this *OsmParser) ParsePbf(r io.Reader) error {
	err := this.useTempKV()
	if nil != err {
		return err
	}
	defer this.Cleanup()
	dec := gosmparse.NewDecoder(r)
	dec.QueueSize = 50
	err = dec.Parse(this)
	this.flushIndex()
	return err
}

func (this *OsmParser) ParseFromTarBz2(r io.Reader) error {
	return this.ParseFromTar(bzip2.NewReader(r))
}

func (this *OsmParser) ParseFromTar(r io.Reader) error {
	tarRdr := tar.NewReader(r)
	for {
		_, err := tarRdr.Next()
		if io.EOF == err {
			break
		}
		if nil != err {
			return err
		}

		this.ParsePbf(tarRdr)
	}

	return nil

}

func (this *OsmParser) addTmpItemLatLng(prefix byte, id int64, latlng s2.LatLng) {
	this.addTmpItem(prefix, id, &Item{
		LatLngs: []s2.LatLng{latlng},
	})
}

func (this *OsmParser) addTmpItem(prefix byte, id int64, item *Item) {
	bb, err := spatial.SerializeRaw(item)
	if nil != err {
		log.Println("addTmpItem err:", err)
		return
	}

	this.syncAdd.Lock()
	defer this.syncAdd.Unlock()
	this.tmpKvBatch.Set(idWithPrefix(prefix, id), bb)

	this.tmpKvBatchCnt++
	if this.batchsize < this.tmpKvBatchCnt {
		this.execBatch()
	}
}

func (this *OsmParser) execBatch() {
	this.FlushExt()
	this.flushIndex()
	fmt.Println("batch cnt=", this.tmpKvBatchCnt)
	this.tmpKvWrtr.ExecuteBatch(this.tmpKvBatch)
	this.tmpKvBatch.Reset()
	this.tmpKvBatchCnt = 0
}

func (this *OsmParser) getLatlngFromTmp(prefix byte, id int64) (v interface{}, err error) {
	rdr, err := this.tmpKv.Reader()
	if nil != err {
		return
	}
	defer rdr.Close()
	bb, err := rdr.Get(idWithPrefix(prefix, id))
	if nil != err {
		return
	}

	return spatial.DeserializeRaw(bb)
}

func (this *OsmParser) getLatlngFromTmpNodeMulti(ids []int64) (latlngs []s2.LatLng, err error) {
	rdr, err := this.tmpKv.Reader()
	if nil != err {
		return
	}
	defer rdr.Close()
	var keys [][]byte = make([][]byte, len(ids))
	for i, id := range ids {
		keys[i] = idWithPrefix('n', id)
	}
	bbarr, err := rdr.MultiGet(keys)
	if nil != err {
		err = fmt.Errorf("%v, getLatlngFromTmpNodeMulti MultiGet.", err)
		return
	}
	latlngs = make([]s2.LatLng, len(bbarr))
	for i, bb := range bbarr {
		if 0 == len(bb) {
			continue
		}
		v, err := spatial.DeserializeRaw(bb)
		if nil != err {
			err = fmt.Errorf("%v, getLatlngFromTmpNodeMulti DeserializeRaw. len bb=%d.", err, len(bb))
			return nil, err
		}
		latlngs[i] = v.(Item).LatLngs[0]
		//latlngs = append(latlngs, v.([]s2.LatLng)[0])
	}

	return
}

func (this *OsmParser) getLatlngFromTmpWay(wayID int64) (item Item, err error) {
	v, err := this.getLatlngFromTmp('w', wayID)
	if nil != err {
		return
	}
	item = v.(Item)
	return
}

func (this *OsmParser) getLatlngFromTmpNode(nodeID int64) (latlng s2.LatLng, err error) {
	v, err := this.getLatlngFromTmp('n', nodeID)
	if nil != err {
		return
	}
	latlng = v.(Item).LatLngs[0]
	return
}
