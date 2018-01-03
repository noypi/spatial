package osm

import (
	"log"
	"sync/atomic"

	"github.com/golang/geo/s2"
	"github.com/thomersch/gosmparse"
)

func (this *OsmParser) ReadNode(n gosmparse.Node) {
	this.tCurrent = gosmparse.NodeType
	cnt := atomic.AddUint64(&this.nNodeCnt, 1)
	if cnt < this.SkipNodes {
		return
	}
	if (cnt % 500000) == 0 {
		log.Println("node cnt=", cnt)
	}
	latlng := s2.LatLngFromDegrees(n.Lat, n.Lon)
	this.addTmpItemLatLng('n', n.ID, latlng)

}
