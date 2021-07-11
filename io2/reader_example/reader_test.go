package render_example

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReader(t *testing.T) {
	t.Run("api test", func(t *testing.T) {
		reader := strings.NewReader("0123456789")
		b := []byte{}
		n, err := reader.Read(b)
		assert.Equal(t, n, 0)
		assert.Nil(t, err)
		assert.Equal(t, b, []byte{})
	})
}
