package slice

import (
	"testing"

	"github.com/bmizerany/assert"
)

var ss = []int{2, 4, 6, 10, 20, 60, 80, 100, 200, 500, 1000}

// 把指定位置的元素剪切掉
func TestCut(t *testing.T) {
	i := 2
	j := 5
	ss = append(ss[:i], ss[j:]...)

	ret := []int{2, 4, 60, 80, 100, 200, 500, 1000}
	assert.Equal(t, ret, ss)
}

func TestDel(t *testing.T) {
	i := 2

	ss1 := append(ss[:i], ss[i+1:]...)
	ss2 := ss[:i+copy(ss[i:], ss[i+1:])]

	ret := []int{2, 4, 20, 60, 80, 100, 200, 500, 1000}
	t.Log(ss1)
	t.Log(ss2)
	assert.Equal(t, ret, ss1)
	assert.Equal(t, ret, ss2)
}
