package seeker_example

import (
	"io"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSeek(t *testing.T) {
	t.Run("seek from start test", func(t *testing.T) {
		reader := strings.NewReader("1234567890")
		n1, err := reader.Seek(2, io.SeekStart)
		assert.Nil(t, err)
		assert.Equal(t, int64(2), n1)
		r1, _, _ := reader.ReadRune()
		assert.Equal(t, '3', r1)
	})
	t.Run("seek from start test", func(t *testing.T) {
		reader := strings.NewReader("1234567890")
		n1, err := reader.Seek(-2, io.SeekEnd)
		assert.Nil(t, err)
		assert.Equal(t, int64(8), n1)
		r1, _, _ := reader.ReadRune()
		assert.Equal(t, '9', r1)
	})
}
