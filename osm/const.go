package osm

import (
	"github.com/thomersch/gosmparse"
)

type tCategory = gosmparse.MemberType

const (
	Node     = gosmparse.NodeType
	Way      = gosmparse.WayType
	Relation = gosmparse.RelationType
)

func toPrefix(t tCategory) byte {
	switch t {
	case gosmparse.NodeType:
		return 'n'
	case gosmparse.WayType:
		return 'w'
	case gosmparse.RelationType:
		return 'r'
	}

	return 0
}
