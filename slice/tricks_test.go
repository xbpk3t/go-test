package slice

import (
	"fmt"
	"testing"
)

// 切片切分问题汇总demo

// [Go语言中切片操作的那些技巧](https://mp.weixin.qq.com/s?__biz=MzAxMTA4Njc0OQ==&mid=2651447062&idx=3&sn=c0fa26d14bf9de04723dd7559d0a0e25)
// [Go Slice Tricks Cheat Sheet](https://ueokande.github.io/go-slice-tricks/)
// [SliceTricks · golang/go Wiki](https://github.com/golang/go/wiki/SliceTricks)
// [Go 官方 Slice 教程图解版](https://mp.weixin.qq.com/s?__biz=MzAxMTA4Njc0OQ==&mid=2651445035&idx=3&sn=7800fa08fe3f83015635dc3f2febb6fd)

func TestSlice1(t *testing.T) {
	ints := f([]int{1, 2, 3})

	fmt.Println(ints)
}

var a []int

func f(b []int) []int {
	a = b[:2]
	return a
}

var c []int

// [Go 切片导致内存泄露，被坑两次了！](https://mp.weixin.qq.com/s?__biz=MzUxMDI4MDc1NA==&mid=2247492355&idx=1&sn=0e468b75394ba9778437b5c72e43c3ad)
func ff(b []int) []int {
	a = b[:2]

	c = append(c, b[:2]...)
	return a
}

// [Go 切片这道题，吵了一个下午！](https://mp.weixin.qq.com/s?__biz=MzI2MzEwNTY3OQ==&mid=2648983302&idx=1&sn=596997b1ee7eb5bf8cd41beb9f895d3b)
