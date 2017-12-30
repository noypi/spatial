package spatial

import (
	"bytes"
	"encoding/binary"
	"encoding/gob"
	"math"

	. "github.com/noypi/spatial/common"

	"github.com/noypi/kv"
	"github.com/rs/xid"
)

const (
	xFF = 0xffffffffffffffff
	x80 = 0x8000000000000000
)

var (
	bbEndKeyRange        = fillarray(0xff, 17)
	bbEndKeyRangeReverse = fillarray(0xff, 17)
)

func init() {
	bbEndKeyRange[0] = cPrefixRange
	bbEndKeyRangeReverse[0] = cPrefixRangeReverse
	gob.Register(_gobitem{})
	gob.Register(_Item{})
	gob.Register(_ID{})
	gob.Register(Range{})
}

type _search1Dfunc func(x, y uint64) Enum
type _search2Dfunc func(x, y Range) Enum

type _gobitem struct {
	V *_Item
}

func serializev(id []byte, v interface{}, r Range) (item *_Item, vbb []byte, err error) {
	if o, ok := v.(*_Item); ok {
		item = o.Clone()
	} else {
		item = NewItem(id, v, r)
	}
	vbb, err = GobSerialize(item)
	return
}

func setItemToBatch(batch kv.KVBatch, id _ID, vbb []byte) {
	batch.Set(id[:], vbb)
	rev := id.Reverse(cPrefixRangeReverse)
	batch.Set(rev[:], vbb)
}

func deleteItemToBatch(batch kv.KVBatch, id _ID) {
	batch.Delete(id[:])
	rev := id.Reverse(cPrefixRangeReverse)
	batch.Delete(rev[:])
}

func GobSerialize(v *_Item) ([]byte, error) {
	buf := new(bytes.Buffer)
	o := &_gobitem{v}
	err := gob.NewEncoder(buf).Encode(o)
	return buf.Bytes(), err
}

func GobDeserialize(bb []byte) (*_Item, error) {
	buf := bytes.NewBuffer(bb)
	o := new(_gobitem)
	if err := gob.NewDecoder(buf).Decode(o); nil != err {
		return nil, err
	}
	return o.V, nil
}

func FloatToSortableInt(f float64) uint64 {
	n := math.Float64bits(f)
	if x80 == (x80 & n) {
		n ^= xFF
	} else {
		n ^= x80
	}
	return n
}

func FloatFromSortableInt(n uint64) float64 {
	if x80 == (x80 & n) {
		n ^= x80
	} else {
		n ^= xFF
	}
	return math.Float64frombits(n)
}

func SerializeUint64(n uint64) []byte {
	bb := make([]byte, 8)
	binary.BigEndian.PutUint64(bb, n)
	return bb
}

func DeserializeUint64(bb []byte) uint64 {
	return binary.BigEndian.Uint64(bb)
}

func toKey(prefix byte, r Range, id xid.ID) []byte {
	return toKeyf(prefix, r.Min, r.Max, id)

}

func toKeyReverse(prefix byte, r Range, id xid.ID) []byte {
	return toKeyf(prefix, r.Max, r.Min, id)

}

func toKeyf(prefix byte, min, max uint64, id xid.ID) []byte {
	key := append([]byte{prefix}, SerializeUint64(min)...)
	bb := append(key, SerializeUint64(max)...)
	return append(bb, id[:]...)
}

func searchKey(prefix byte, n uint64) []byte {
	return append([]byte{prefix}, SerializeUint64(n)...)
}

func keyToRange(bb []byte) (r Range, id xid.ID) {
	f1 := DeserializeUint64(bb[1:9])
	f2 := DeserializeUint64(bb[9:17])
	if cPrefixRange == bb[0] {
		r.Min, r.Max = f1, f2
	} else if cPrefixRangeReverse == bb[0] {
		r.Min, r.Max = f2, f1
	}
	copy(id[:], bb[cIDRequiredLen:len(bb)])
	return
}

func fillarray(fill byte, n int) []byte {
	bb := make([]byte, n)
	for i := 0; i < n; i++ {
		bb[i] = fill
	}
	return bb
}
