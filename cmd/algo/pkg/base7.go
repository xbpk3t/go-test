package main

import "fmt"

func convertToBase7(num int) string {
	// 处理负数的情况
	if num < 0 {
		return "-" + convertToBase7(-num)
	}

	// 处理零的情况
	if num == 0 {
		return "0"
	}

	result := ""
	for num != 0 {
		// 取余数
		remainder := num % 7
		// 将余数转换为字符并拼接到结果中
		result = string(rune('0'+remainder)) + result
		// 除以7，继续进行下一位的转换
		num /= 7
	}

	return result
}

func main() {
	num := 100
	base7 := convertToBase7(num)
	fmt.Println("10进制数", num, "转换为7进制数为", base7)

	fmt.Println(-100, "====>", convertToBase7(-100))
	fmt.Println(-1, "====>", convertToBase7(-1))
}
