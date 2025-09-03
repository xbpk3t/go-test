package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(strings.TrimRight("123aabc", "abc"))
	fmt.Println(strings.TrimSuffix("123aabc", "abc"))
	fmt.Println(strings.Trim("123aabc", "abc"))
	fmt.Println(strings.CutSuffix("123aabc", "abc"))
}
