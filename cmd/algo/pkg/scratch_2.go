package main

import (
	"fmt"
	"slices"
)

func main() {
	ints := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20}
	fmt.Println(ArrayReverse(ints))
	//
	// // fmt.Println(sort.Reverse(z))
	slices.Reverse(ints)
	fmt.Println(ints)

	// s := []int{5, 2, 6, 3, 1, 4}
	//
	// slices.Reverse(s)
	//
	// fmt.Println(s)

	// fmt.Println(ArrayReverseX(ints))
}

func ArrayReverse(arr []int) []int {
	l := len(arr)
	if l <= 0 {
		return []int{}
	}
	// var res []int
	// for i := 0; i < l; i++ {
	// 	res = append(res, arr[i])
	// }
	for i, j := 0, l-1; i < j; i, j = i+1, j-1 {
		arr[i], arr[j] = arr[j], arr[i]
	}

	return arr
}

func ArrayReverseX(arr []int) []int {
	l := len(arr)
	if l <= 0 {
		return []int{}
	}
	var rev []int
	for _, n := range arr {
		rev = append([]int{n}, rev...)
	}
	return rev
}
