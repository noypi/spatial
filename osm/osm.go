// openstreetmap implementation
package osm

import (
	"os"

	"github.com/blevesearch/bleve"
	"github.com/noypi/spatial/db"
	"github.com/noypi/spatial/geo"
)

type Osm struct {
	*geo.SpatialGeo
	index bleve.Index
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
	return parser
}

func (this Osm) myDbPath() string {
	s := this.DbPath()
	if 0 == len(s) {
		s = "."
	}
	return s
}

func (this *Osm) GetInfo(id int64) (v interface{}, err error) {
	bbID := idFromInt(id)
	return this.GetExtInfo(uint8(Way), bbID)
}

func (this *Osm) openIndex() (err error) {
	fpath := this.myDbPath() + "/index.bleve"
	_, err = os.Stat(fpath)
	if os.IsNotExist(err) {
		indexMap := bleve.NewIndexMapping()
		if this.index, err = bleve.New(fpath, indexMap); nil != err {
			return
		}

	} else {
		if this.index, err = bleve.Open(fpath); nil != err {
			return
		}
	}

	return
}
