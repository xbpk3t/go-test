package main

import (
	"fmt"
	"github.com/bits-and-blooms/bitset"
)

// 设置哈希数组默认大小为 16
const DefaultSize = 16

// 设置种子，保证不同哈希函数有不同的计算方式
var seeds = []uint{7, 11, 13, 31, 37, 61}

// 布隆过滤器结构，包括二进制数组和多个哈希函数
type BloomFilter struct {
	// 使用第三方库
	set *bitset.BitSet
	// 指定长度为 6
	hashFuncs [6]func(seed uint, value string) uint
}

// 构造一个布隆过滤器，包括数组和哈希函数的初始化
func NewBloomFilter() *BloomFilter {
	bf := new(BloomFilter)
	bf.set = bitset.New(DefaultSize)

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

// 添加元素
func (b *BloomFilter) add(value string) {
	for i, f := range b.hashFuncs {
		// 将哈希函数计算结果对应的数组位置 1
		b.set.Set(f(seeds[i], value))
	}
}

// 判断元素是否存在
func (b *BloomFilter) contains(value string) bool {
	// 调用每个哈希函数，并且判断数组对应位是否为 1
	// 如果不为 1，直接返回 false，表明一定不存在
	for i, f := range b.hashFuncs {
		// result = result && b.set.Test(f(seeds[i], value))
		if !b.set.Test(f(seeds[i], value)) {
			return false
		}
	}
	return true
}
func main() {
	filter := NewBloomFilter()
	filter.add("asd")
	fmt.Println(filter.contains("asd"))
	fmt.Println(filter.contains("2222"))
	fmt.Println(filter.contains("155343"))
}
