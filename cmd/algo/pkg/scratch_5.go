package main

import (
	"fmt"
	"slices"
)

// How to remove nil from slice of interface?
// [go - Golang removing nil from slice of interface{} - Stack Overflow](https://stackoverflow.com/questions/40899548/golang-removing-nil-from-slice-of-interface)
func main() {
	things := []interface{}{
		nil,
		1,
		nil,
		"2",
		nil,
		3,
		nil,
	}

	things = slices.DeleteFunc(
		things,
		func(thing interface{}) bool {
			return thing == nil
		},
	)

	things = slices.Clip(things)

	fmt.Printf("%#v\n", things)

	s := []int{1, 2, 2, 3, 3, 4, 5, 2, 3}
	newSlice := slices.Compact(s)
	fmt.Println(newSlice)

	// slices, BinarySearch()/Func(), Clip(), Clone(), Compact()/Func(), Compare()/Func(), Contains()/Func(), Delete()/Func(), Equal()/Func(), Grow(), Index()/Func(), Insert(), IsSorted()/Func(), Max()/Func(), Min()/Func(), Replace(), Reverse(), Sort()/Func()
	fmt.Println(slices.Compare(s, []int{1, 2, 3}))

	fmt.Println(slices.BinarySearch(s, 2))
}
