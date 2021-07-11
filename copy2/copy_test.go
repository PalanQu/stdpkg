package copy2

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCopy(t *testing.T) {
	t.Run("copy slice test", func(t *testing.T) {
		// copy方法在copy一个slice的时候，会看dst的长度，而不是容量，
		// 长度如果小于src的长度，那么截取部分src然后copy到dst上。
		slice := make([]byte, 2, 3)
		assert.Equal(t, 2, len(slice))
		assert.Equal(t, 3, cap(slice))
		copy(slice, []byte{1, 2, 3, 4})
		assert.Equal(t, []byte{1, 2}, slice)

		// 注意以下的代码生成的slice的长度和容量都是0，所以无论用什么copy到这个slice，
		// 结果都是空slice
		slice2 := []byte{}
		assert.Equal(t, 0, len(slice2))
		assert.Equal(t, 0, cap(slice2))
		copy(slice2, []byte{1, 2, 3, 4})
		assert.Equal(t, []byte{}, slice2)
	})
	t.Run("dst is shorter than src", func(t *testing.T) {
		// 当src的长度比较短的时候，会copy全部src到dst，而dst多余的部分不会覆盖
		slice := make([]byte, 3)
		copy(slice, []byte{1, 2})
		assert.Equal(t, slice, []byte{1, 2, 0})
	})
	t.Run("copy struct", func(t *testing.T) {
		// copy只是浅拷贝，struct的指针是相同的，如果这时候更改了slice中的student，
		// 那么slice2中的也会跟着一起改，反之亦然
		slice := make([]*student, 1)
		tom := student{"tom"}
		slice2 := []*student{&tom}
		copy(slice, slice2)
		assert.NotEqual(t, fmt.Sprintf("%p", slice), fmt.Sprintf("%p", slice2))
		assert.Equal(t,
			fmt.Sprintf("%p", slice[0]), fmt.Sprintf("%p", slice2[0]))
		newName := "jack"
		slice[0].name = newName
		assert.Equal(t, newName, slice2[0].name)
	})
}

type student struct {
	name string
}
