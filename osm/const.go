package osm

type tCategory uint8

const (
	Node tCategory = iota
	Way
	Relation
)
