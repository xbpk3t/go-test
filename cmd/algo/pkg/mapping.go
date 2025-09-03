package main

import (
	"fmt"
)

// 假设我们有一个数字到字母表的映射:
// 1-> ['a', 'b', 'c]
// 2-> ['d’, 'e’]
// 3-> ['f', 'g', "h]
// 实现一个函数,对于给定的一串数字,例如"1". "233",返回一个包含所有可能的组合的字符串列表
//
// 对于固定长度且长度较小的可以用多个for生成组成，对于本题这种不固定长度的要用递归来生成
func main() {
	ret := letterConbination("3122")
	fmt.Println(ret)
}

var numMap = map[string][]string{
	"1": {"a", "b", "c"},
	"2": {"d", "e"},
	"3": {"f", "g", "h"},
}

func letterConbination(nums string) (ret []string) {

	c := string(nums[0])
	letterMap := numMap[c]

	retconb := make([]string, 0)
	if len(nums) > 1 {
		retconb = letterConbination(nums[1:])
	}

	for _, letter := range letterMap {

		if len(nums) > 1 {
			for _, conb := range retconb {
				a := letter + conb
				ret = append(ret, a)
			}
		} else {
			ret = append(ret, letter)
		}

	}
	return
}
