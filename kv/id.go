package spatial

import (
	"github.com/rs/xid"
)

type ID [29]byte

func NewID() ID {
	id := ID{}
	x := xid.New()
	copy(id[17:], x[:])
	return id
}

func (id ID) Reverse(prefix byte) ID {
	rev := ID{}
	rev.SetPrefix(prefix)
	rev.SetLeft(id.Right())
	rev.SetRight(id.Left())
	copy(rev[17:], id.XID())
	return rev
}

func (id ID) Clone() ID {
	out := ID{}
	copy(out[:], id[:])
	return out
}

func (id ID) Prefix() byte {
	return id[0]
}

func (id ID) Left() []byte {
	return id[1:9]
}

func (id ID) Right() []byte {
	return id[9:17]
}

func (id ID) LeftFloat64() float64 {
	return DeserializeFloat64(id.Left())
}

func (id ID) RightFloat64() float64 {
	return DeserializeFloat64(id.Right())
}

func (id ID) XID() []byte {
	return id[17:29]
}

func (id *ID) SetPrefix(b byte) {
	id[0] = b
}

func (id *ID) SetLeft(bb []byte) {
	copy(id[1:], bb)
}

func (id *ID) SetRight(bb []byte) {
	copy(id[9:], bb)
}

func (id *ID) SetLeftFloat64(left float64) {
	id.SetLeft(SerializeFloat64(left))
}

func (id *ID) SetRightFloat64(right float64) {
	id.SetRight(SerializeFloat64(right))
}
