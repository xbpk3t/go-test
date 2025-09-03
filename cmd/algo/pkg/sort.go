package main

func main() {

}

// 评价：最直观，也是最基础的比较排序算法，没啥使用场景，时间复杂度为 O(n^2)
// 使用双循环来进行排序。外部循环控制所有的回合，内部循环代表每一轮的冒泡处理，先进行元素比较，再进行元素交换
// 具体如下：

// 比较相邻的元素。如果第一个比第二个大，就交换他们两个。
// 对每一对相邻元素作同样的工作，从开始第一对到结尾的最后一对。这步做完后，最后的元素会是最大的数。
// 针对所有的元素重复以上的步骤，除了最后一个。
// 持续每次对越来越少的元素重复上面的步骤，直到没有任何一对数字需要比较。
func bubbleSort(arr []int) []int {
	length := len(arr)
	for i := 0; i < length; i++ {
		for j := 0; j < length-1-i; j++ {
			// 判断数组大小，并颠倒位置
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
	return arr
}

// 冒泡+flag
// 如果在本轮排序中，元素有交换，则说明数列无序；如果没有元素交换，说明数列已然有序，直接跳出大循环。
func bubbleSort(nums []int) []int {
	if len(nums) <= 1 {
		return nums
	}

	// 冒泡排序核心实现代码
	for i := 0; i < len(nums); i++ {
		// highlight-next-line
		flag := false
		for j := 0; j < len(nums)-i-1; j++ {
			if nums[j] > nums[j+1] {
				nums[j], nums[j+1] = nums[j+1], nums[j]
				// highlight-next-line
				flag = true
			}
		}
		// highlight-start
		if !flag {
			break
		}
		// highlight-end
	}

	return nums
}

// 选择排序
// 评价：与冒泡一样，也很直观，时间复杂度也是 O(n^2)。但是因为不占用额外的内存空间，所以尤其适合数据量较小的场景。但是需要注意选择排序是不稳定的排序算法。
// 选择排序每次会从未排序区间中找到最小的元素，将其放到已排序区间的末尾。这样一来，当遍历完未排序区间，就意味着已经完成整个序列的排序了
func selectionSort(nums []int) {
	if len(nums) <= 1 {
		return
	}
	// 已排序区间初始化为空，未排序区间初始化待排序切片
	for i := 0; i < len(nums); i++ {
		// 未排序区间最小值初始化为第一个元素
		min := i
		// 从未排序区间第二个元素开始遍历，直到找到最小值
		for j := i + 1; j < len(nums); j++ {
			if nums[j] < nums[min] {
				min = j
			}
		}
		// 将最小值与未排序区间第一个元素互换位置（等价于放到已排序区间最后一个位置）
		if min != i {
			nums[i], nums[min] = nums[min], nums[i]
		}
	}
}
