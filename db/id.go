package spatial

type _ID []byte

const cIDRequiredLen = 17

func NewID(id []byte) _ID {
	bb := make([]byte, cIDRequiredLen+len(id))
	copy(bb[17:], id)
	return bb
}

func (id _ID) Reverse(prefix byte) _ID {
	rev := NewID(id.ID())
	rev.SetPrefix(prefix)
	rev.SetLeft(id.Right())
	rev.SetRight(id.Left())
	copy(rev[cIDRequiredLen:], id.ID())
	return rev
}

func (id _ID) Clone() _ID {
	out := make([]byte, len(id))
	copy(out[:], id[:])
	return out
}

func (id _ID) Prefix() byte {
	return id[0]
}

func (id _ID) Left() []byte {
	return id[1:9]
}

func (id _ID) Right() []byte {
	return id[9:17]
}

func (id _ID) LeftUint64() uint64 {
	return DeserializeUint64(id.Left())
}

func (id _ID) RightUint64() uint64 {
	return DeserializeUint64(id.Right())
}

func (id _ID) ID() []byte {
	return id[17:len(id)]
}

func (id *_ID) SetPrefix(b byte) {
	(*id)[0] = b
}

func (id *_ID) SetLeft(bb []byte) {
	copy((*id)[1:], bb)
}

func (id *_ID) SetRight(bb []byte) {
	copy((*id)[9:], bb)
}

func (id *_ID) SetLeftUint64(left uint64) {
	id.SetLeft(SerializeUint64(left))
}

func (id *_ID) SetRightUint64(right uint64) {
	id.SetRight(SerializeUint64(right))
}
