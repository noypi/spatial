// openstreetmap implementation
package osm

import (
	"sync"

	"github.com/blevesearch/bleve"
	"github.com/golang/geo/s2"
	"github.com/noypi/spatial/db"
	"github.com/noypi/spatial/geo"
)

type Osm struct {
	*geo.SpatialGeo
	index     bleve.Index
	syncIndex sync.Mutex

	indexBatch *bleve.Batch
}

func New(opts ...spatial.Options) (o *Osm, err error) {
	o = new(Osm)
	if o.SpatialGeo, err = geo.NewGeo(opts...); nil != err {
		return
	}

	return
}

func (this *Osm) AsParser(batchsize uint32) *OsmParser {
	parser := &OsmParser{
		Osm:       this,
		batchsize: batchsize,
		dbpath:    this.myDbPath(),
	}
	parser.useTempKV()
	parser.workerPool.Start(200)
	return parser
}

func (this Osm) myDbPath() string {
	s := this.DbPath()
	if 0 == len(s) {
		s = "."
	}
	return s
}

func (this *Osm) Close() {
	this.FlushExt()
	this.SpatialGeo.Close()
	if nil != this.index {
		this.index.Close()
		this.index = nil
	}
}

func (this *Osm) GetWayInfo(id int64) (v interface{}, err error) {
	/*	latlngs, err := this.getLatlngs(Way, id)
		if nil != err {
			return
		}*/

	d, err := this.index.Document(toSearchId(Way, id))
	if nil != err {
		return
	}

	v = d
	return
}

func (this *Osm) getLatlngs(t tCategory, id int64) (latlngs []s2.LatLng, err error) {
	bbID := idFromInt(id)
	v, err := this.GetExtInfo(uint8(t), bbID)
	if nil != err {
		return
	}
	latlngs = v.(Item).LatLngs
	return
}
