package osm

import (
	"encoding/binary"
)

func idFromInt(n int64) []byte {
	bb := make([]byte, 8)
	binary.BigEndian.PutUint64(bb, uint64(n))
	return bb
}

func idFromBytes(bb []byte) int64 {
	return int64(binary.BigEndian.Uint64(bb))
}

func idWithPrefix(prefix byte, n int64) []byte {
	return append([]byte{prefix}, idFromInt(n)...)
}
