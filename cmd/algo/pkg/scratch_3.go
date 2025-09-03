package main

import "fmt"

// [说两个 Go 循环里的坑，第二个在JS里也有](https://mp.weixin.qq.com/s?__biz=MzUzNTY5MzU2MA==&mid=2247496998&idx=1&sn=b8fd8d6f8f89c53928d5b4353c42f2ca)
// [Golang中的for-range趟坑](https://mp.weixin.qq.com/s?__biz=MzAxMTA4Njc0OQ==&mid=2651449964&idx=3&sn=ee3d552e54577d2c24846ce4abd12bb0)
// [Go 陷阱之 for 循环迭代变量 | Go 技术论坛](https://learnku.com/articles/26861)

// [随笔：Golang 循环变量引用问题以及官方语义修复-腾讯云开发者社区-腾讯云](https://cloud.tencent.com/developer/article/2240620)
// golang1.22 之后，之前“for循环时的变量问题”解决了
func main() {

	var prints []func()
	for _, v := range []int{1, 2, 3} {
		prints = append(prints, func() { fmt.Println(v) })
	}
	for _, print := range prints {
		print()
	}
}
