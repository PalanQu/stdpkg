package render_example

import (
	"fmt"
	"io"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReader(t *testing.T) {
	t.Run("api test", func(t *testing.T) {
		// 生成一个字符串的reader，此reader的Read方法就是从init时候传入的字符串中读取字符，
		// Read方法会返回两个参数，一个是此次Read读了多少字符，另一个是在读的过程中有没有错误
		// 边界情况是，
		// 		1. 正好读取完字符串，没有错误返回，继续读取会返回错误，同时改变的字符长度为0
		//		2. 当第一次没有读取完字符串，第二次不仅读取完字符串，还超出了长度，
		//		那么不会返回错误，改变的字符串长度就是剩下没有读取完的长度
		reader := strings.NewReader("0123456789")
		b := make([]byte, 3)
		n, err := reader.Read(b)
		assert.Equal(t, 3, n)
		assert.Nil(t, err)
		assert.Equal(t, []byte{'0', '1', '2'}, b)

		b2 := make([]byte, 2)
		n2, err := reader.Read(b2)
		assert.Equal(t, 2, n2)
		assert.Nil(t, err)
		assert.Equal(t, []byte{'3', '4'}, b2)

		b3 := make([]byte, 5)
		n3, err := reader.Read(b3)
		assert.Equal(t, 5, n3)
		assert.Nil(t, err)
		assert.Equal(t, []byte{'5', '6', '7', '8', '9'}, b3)

		b4 := make([]byte, 2)
		n4, err := reader.Read(b4)
		assert.Equal(t, 0, n4)
		assert.EqualError(t, err, io.EOF.Error())
		assert.Equal(t, []byte{0, 0}, b4)

		// 边界情况2
		reader.Reset("0123456789")

		b5 := make([]byte, 9)
		n5, err := reader.Read(b5)
		assert.Equal(t, 9, n5)
		assert.Nil(t, err)
		assert.Equal(t, 9, len(b5))

		b6 := make([]byte, 2)
		n6, err := reader.Read(b6)
		assert.Equal(t, 1, n6)
		assert.Nil(t, err)
		fmt.Println(string(b6))
		assert.Equal(t, []byte{'9', 0}, b6)
	})
}
