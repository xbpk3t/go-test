package main

import "fmt"

// 先需要考虑，怎么构造 filter struct
// 建议使用 uint
type Filter struct {
	hashFuncs []func(seed uint, value string) uint
	bitset    *Bitset
}

type Bitset struct {
	// 用于存储数据的 slice
	data []uint
	// 用于存储数据的长度
	length uint
}

// 设置哈希数组默认大小为 16
const DefaultSize = 16

func NewFilter() *Filter {
	bf := new(Filter)
	bf.bitset = &Bitset{length: DefaultSize, data: make([]uint, DefaultSize)}
	for i := 0; i < len(bf.hashFuncs); i++ {
		bf.hashFuncs[i] = createHash()
	}
	return bf
}

// 构造 6 个哈希函数，每个哈希函数有参数 seed 保证计算方式的不同
func createHash() func(seed uint, value string) uint {
	return func(seed uint, value string) uint {
		var result uint = 0
		for i := 0; i < len(value); i++ {
			result = result*seed + uint(value[i])
		}
		// length = 2^n 时，X % length = X & (length - 1)
		return result & (DefaultSize - 1)
	}
}

func (bf Filter) put(b string) {
	for i, f := range bf.hashFuncs {
		// 将哈希函数计算结果对应的数组位置 1
		// 如果 bitset 有 Set() 方法，可以直接调用
		// 我这里直接复制
		bf.bitset.data[f(uint(i), b)] = 1
	}
}

func (bf Filter) contains(b string) bool {
	// 调用每个哈希函数，并且判断数组对应位是否为 1
	// 如果不为 1，直接返回 false，表明一定不存在
	for i, f := range bf.hashFuncs {
		if bf.bitset.data[f(uint(i), b)] != 1 {
			return false
		}
	}
	return true
}

func main() {
	bf := NewFilter()
	bf.put("hello")
	fmt.Println(bf.contains("hello"))
	fmt.Println(bf.contains("world"))
	fmt.Println(bf.contains("hello world"))
}
