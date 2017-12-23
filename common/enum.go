package spatial

type Item interface {
	Error() error
	Value() interface{}
	ID() string
	Range(n int) Range
	Delete()
	Set(v interface{}) error
}

type Enum interface {
	Next() (item Item, has bool)
	Close()
}
