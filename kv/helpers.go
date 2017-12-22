package spatial

import (
	"bytes"
	"encoding/binary"
	"encoding/gob"
	"math"

	"github.com/twinj/uuid"
)

const (
	xFF = 0xffffffffffffffff
	x80 = 0x8000000000000000
)

var (
	bbEndKeyRange        = fillarray(0xff, 17)
	bbEndKeyRangeReverse = fillarray(0xff, 17)
)

var Epsilon = math.Nextafter(1, 2) - 1
var Epsilonx10 = Epsilon * 10

func NewItem(v interface{}) *_Item {
	return &_Item{Value: v, ID: uuid.NewV4().String()}
}

func init() {
	bbEndKeyRange[0] = cPrefixRange
	bbEndKeyRangeReverse[0] = cPrefixRangeReverse
	gob.Register(_gobitem{})
	gob.Register(_Item{})
}

func IsZero(a float64) bool {
	return math.Abs(a) < Epsilon
}

func IsLessOrEqual(a, b float64) bool {
	if a < b {
		return true
	}

	return IsZero(a - b)
}

type _search1Dfunc func(x, y float64) *Enum
type _search2Dfunc func(x, y Range) *Enum

type _gobitem struct {
	V *_Item
}

type _Item struct {
	Error error
	Value interface{}
	ID    string
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

func SerializeFloat64(f float64) []byte {
	bb := make([]byte, 8)
	n := math.Float64bits(f)
	if x80 == (x80 & n) {
		n ^= xFF
	} else {
		n ^= x80
	}
	binary.BigEndian.PutUint64(bb, n)
	return bb
}

func DeserializeFloat64(bb []byte) float64 {
	n := binary.BigEndian.Uint64(bb)
	if x80 == (x80 & n) {
		n ^= x80
	} else {
		n ^= xFF
	}
	return math.Float64frombits(n)
}

func toKey(prefix byte, r Range) []byte {
	return toKeyf(prefix, r.Min, r.Max)
}

func toKeyReverse(prefix byte, r Range) []byte {
	return toKeyf(prefix, r.Max, r.Min)
}

func toKeyf(prefix byte, min, max float64) []byte {
	key := append([]byte{prefix}, SerializeFloat64(min)...)
	return append(key, SerializeFloat64(max)...)
}

func searchKey(prefix byte, f float64) []byte {
	return append([]byte{prefix}, SerializeFloat64(f)...)
}

func keyToRange(bb []byte) (r Range) {
	f1 := DeserializeFloat64(bb[1:9])
	f2 := DeserializeFloat64(bb[9:17])
	if cPrefixRange == bb[0] {
		r.Min, r.Max = f1, f2
	} else if cPrefixRangeReverse == bb[0] {
		r.Min, r.Max = f2, f1
	}
	return
}

func fillarray(fill byte, n int) []byte {
	bb := make([]byte, n)
	for i := 0; i < n; i++ {
		bb[i] = fill
	}
	return bb
}
