package spatial_test

import (
	"testing"

	"github.com/noypi/spatial/db"
	assertpkg "github.com/stretchr/testify/assert"
)

func TestID(t *testing.T) {
	assert := assertpkg.New(t)

	var expected_id = []byte{1, 2, 3, 4}

	id := spatial.NewID(expected_id)
	id.SetLeftUint64(2442)
	id.SetRightUint64(4224)
	id.SetPrefix('b')

	assert.Equal(uint64(2442), id.LeftUint64())
	assert.Equal(uint64(4224), id.RightUint64())
	assert.Equal(byte('b'), id.Prefix())
	assert.Equal(expected_id, id.ID())
}

func TestIDReverse(t *testing.T) {
	assert := assertpkg.New(t)

	var expected_id = []byte{1, 2, 3, 4}

	id := spatial.NewID(expected_id)
	id.SetLeftUint64(2442)
	id.SetRightUint64(4224)
	id.SetPrefix('b')

	rev := id.Reverse('c')

	assert.Equal(uint64(2442), rev.RightUint64())
	assert.Equal(uint64(4224), rev.LeftUint64())
	assert.Equal(byte('c'), rev.Prefix())
	assert.Equal(id.ID(), rev.ID())
}
