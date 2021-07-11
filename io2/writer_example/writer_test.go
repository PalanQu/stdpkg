package writer_example

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestWriter(t *testing.T) {
	t.Run("test1", func(t *testing.T) {
		// writer.Write方法输入一个bytes，然后将这个bytes写入writer中，返回写入的长度和错误
		writer := bytes.Buffer{}
		n1, err := writer.Write([]byte{1})
		assert.Nil(t, err)
		assert.Equal(t, 1, n1)
		assert.Equal(t, []byte{1}, writer.Bytes())
	})
}
