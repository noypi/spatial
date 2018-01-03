package osm_test

import (
	"testing"

	"github.com/noypi/spatial/db"
	"github.com/noypi/spatial/osm"
	assertpkg "github.com/stretchr/testify/assert"
	"github.com/thomersch/gosmparse"
)

func TestOsm(t *testing.T) {
	assert := assertpkg.New(t)

	o, err := osm.New(spatial.OptKVDir{"/tmp"})
	assert.Nil(err)

	parser := o.AsParser(10)
	defer parser.Cleanup()

	for i := 0; i < 10; i++ {
		parser.ReadNode(gosmparse.Node{
			ID:  1001 + int64(i),
			Lat: 4001.12 + float64(i),
			Lon: 501.23 + float64(i),
		})
	}

	parser.ReadWay(gosmparse.Way{
		ID:      2001,
		NodeIDs: []int64{1001, 1003, 1004},
	})

	parser.ReadRelation(gosmparse.Relation{
		ID: 3001,
		Members: []gosmparse.RelationMember{
			{ID: 1001, Type: gosmparse.NodeType},
			{ID: 2001, Type: gosmparse.RelationType},
		},
		Tags: map[string]string{
			"Key01": "value01",
			"Key02": "value02",
			"Key03": "value03",
		},
	})

	results, err := parser.Search("value01")
	assert.Nil(err)
	assert.Equal(1, results.Hits.Len())
	assert.Equal("3001", results.Hits[0].ID)

	v, err := parser.GetInfo(3001)
	assert.Nil(err)
	m, ok := v.(map[string]string)
	assert.True(ok)
	assert.Equal("value01", m["Key01"])

}
