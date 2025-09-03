package main

import "fmt"

// 快排适合用来处理数组中没有重复元素的情况
// 如果有重复元素则需要“三向切分”（也就是判断“相等情况”）
func main() {
	arr := []int{8, 5, 1, 0, 9, 2, 3, 6, 7, 7, 6}
	ints := sort(arr)
	fmt.Println(ints)
}

func sort(arr []int) []int {
	count := len(arr)
	if count <= 1 {
		return arr
	}
	pivot := arr[0]
	var lt []int
	var gt []int
	for i := 1; i < count; i++ {
		v := arr[i]
		if v < pivot {
			lt = append(lt, v)
		}
		if v > pivot {
			gt = append(gt, v)
		}
	}
	lt = sort(lt)
	gt = sort(gt)

	return append(append(lt, pivot), gt...)
}

// quicksort 三项切分
// 在前面 qs 的基础上实现，非常好懂
// func quicksort(arr []int) []int {
// 	count := len(arr)
// 	if count <= 1 {
// 		return arr
// 	}
//
// 	pivot := arr[0]
// 	var lt []int
// 	var eq []int
// 	var gt []int
//
// 	for i := 0; i < count; i++ {
// 		v := arr[i]
// 		if v < pivot {
// 			lt = append(lt, v)
// 		} else if v > pivot {
// 			gt = append(gt, v)
// 		} else {
// 			eq = append(eq, v)
// 		}
// 	}
//
// 	lt = quicksort(lt)
// 	gt = quicksort(gt)
//
// 	return append(append(lt, eq...), gt...)
// }

// 三向切分
// 划分方式：
//
// 第一个实现使用了三个数组（left、equal 和 right）来分别存储小于、等于和大于基准值的元素。
//
// 第二个实现使用了两个指针（lt 和 gt）来进行划分，通过交换元素的位置来将小于、等于和大于基准值的元素分别放在数组的左侧、中间和右侧。
//
// ---
//
// 性能比较：
//
// 在平均情况下，第一个实现的性能可能会略好于第二个实现。这是因为第一个实现在划分过程中只需要遍历一次整个数组，而第二个实现需要在每一次划分中进行多次元素交换。
//
// 然而，在最坏情况下，第一个实现可能会导致不平衡的划分，从而影响性能。而第二个实现使用了三向切分，可以更好地处理重复元素，避免不平衡划分。
//
// ---
//
// 总体而言，这两个实现在性能上可能有一些差异，具体取决于输入数据的特性。第一个实现可能在大部分情况下表现较好，但在处理大量重复元素的情况下，第二个实现可能更具优势。选择哪个实现取决于您的具体需求和数据特性。
func quicksort(arr []int) {
	if len(arr) <= 1 {
		return
	}

	lt, gt := 0, len(arr)-1
	pivot := arr[0]
	i := 1

	for i <= gt {
		if arr[i] < pivot {
			arr[i], arr[lt] = arr[lt], arr[i]
			lt++
			i++
		} else if arr[i] > pivot {
			arr[i], arr[gt] = arr[gt], arr[i]
			gt--
		} else {
			i++
		}
	}

	quicksort(arr[:lt])
	quicksort(arr[gt+1:])
}
