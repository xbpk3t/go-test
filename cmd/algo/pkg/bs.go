package main

import "fmt"

func binarySearch(needle int, haystack []int) bool {
	low := 0
	high := len(haystack) - 1
	for low <= high {
		median := (low + high) / 2
		if haystack[median] < needle {
			low = median + 1
		} else {
			high = median - 1
		}
	}
	if low == len(haystack) || haystack[low] != needle {
		return false
	}
	return true
}

func main() {
	items := []int{1, 2, 9, 20, 31, 45, 63, 70, 100}
	fmt.Println(binarySearch(63, items))

	bsc(2, []int{})
}

func bsc(needle int, haystack []int) int {
	low := 0
	high := len(haystack) - 1
	count := 0

	for low <= high {
		median := (low + high) / 2
		if haystack[median] < needle {
			low = median + 1
		} else if haystack[median] > needle {
			high = median - 1
		} else {
			// Found a match, count occurrences
			count++
			// Count occurrences to the left
			for i := median - 1; i >= low; i-- {
				if haystack[i] == needle {
					count++
				} else {
					break
				}
			}
			// Count occurrences to the right
			for i := median + 1; i <= high; i++ {
				if haystack[i] == needle {
					count++
				} else {
					break
				}
			}
			break
		}
	}
	return count

}
