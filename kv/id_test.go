package spatial_test

import (
	"testing"

	"github.com/noypi/spatial/kv"
	assertpkg "github.com/stretchr/testify/assert"
)

func TestID(t *testing.T) {
	assert := assertpkg.New(t)

	id := spatial.NewID()
	id.SetLeftFloat64(24.42)
	id.SetRightFloat64(42.24)
	id.SetPrefix('b')

	assert.Equal(24.42, id.LeftFloat64())
	assert.Equal(42.24, id.RightFloat64())
	assert.Equal(byte('b'), id.Prefix())
	assert.Equal(12, len(id.XID()))
}

func TestIDReverse(t *testing.T) {
	assert := assertpkg.New(t)

	id := spatial.NewID()
	id.SetLeftFloat64(24.42)
	id.SetRightFloat64(42.24)
	id.SetPrefix('b')

	rev := id.Reverse('c')

	assert.Equal(24.42, rev.RightFloat64())
	assert.Equal(42.24, rev.LeftFloat64())
	assert.Equal(byte('c'), rev.Prefix())
	assert.Equal(id.XID(), rev.XID())
}
