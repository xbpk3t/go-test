package slice

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func TestMap(t *testing.T) {
	sMap := sync.Map{}

	fmt.Println("=====================Store=======================")

	fmt.Println("store: key=1, value=jay")
	sMap.Store(1, "jay")
	fmt.Println("store: key=2, value=lhj")
	sMap.Store(2, "lhj")
	fmt.Println("store: key=2, value=lhj-hhh")
	sMap.Store(2, "lhj-hhh")

	fmt.Println("=====================LoadOrStore=======================")

	result, succeed := sMap.LoadOrStore(1, "jays")
	fmt.Printf("key=1 store: value=jays; result: load_succeed=%t, value=%v\n", succeed, result)

	result, succeed = sMap.LoadOrStore(2, "lhj-s")
	fmt.Printf("key=2 store: value=lhj-s; result: load_succeed=%t, value=%v\n", succeed, result)

	result, succeed = sMap.LoadOrStore(3, "jay-lhj")
	fmt.Printf("key=3 store: value=jay-lhj; result: load_succeed=%t, value=%v\n", succeed, result)

	fmt.Println("=====================Load key:1~4=======================")

	if val, ok := sMap.Load(1); ok {
		fmt.Printf("key=1, value=%v\n", val)
	} else {
		fmt.Printf("key=1, load failed\n")
	}

	if val, ok := sMap.Load(2); ok {
		fmt.Printf("key=2, value=%v\n", val)
	} else {
		fmt.Printf("key=2, load failed\n")
	}

	if val, ok := sMap.Load(3); ok {
		fmt.Printf("key=3, value=%v\n", val)
	} else {
		fmt.Printf("key=3, load failed\n")
	}

	if val, ok := sMap.Load(4); ok {
		fmt.Printf("key=4, value=%v\n", val)
	} else {
		fmt.Printf("key=4, load failed\n")
	}

	fmt.Println("======================Delete======================")

	fmt.Println("delete key=1")
	sMap.Delete(1)

	fmt.Println("=====================Range=======================")
	sMap.Range(func(key, value interface{}) bool {
		fmt.Printf("key=%v value=%v\n", key, value)
		return true
	})
}

// [一口气搞懂 Go sync.map 所有知识点](https://mp.weixin.qq.com/s/8aufz1IzElaYR43ccuwMyA#tocbar-12kei0h)
func TestSync(t *testing.T) {
	m := sync.Map{}
	m.Store(1, 1)
	go do(m)
	go do(m)

	time.Sleep(1 * time.Second)
	fmt.Println(m.Load(1))
}

func do(m sync.Map) {
	i := 0
	for i < 10000 {
		m.Store(1, 1)
		i++
	}
}

// 并发读写，会直接报错
func TestSyncMock(t *testing.T) {
	m := make(map[int]int)
	// 读
	go func() {
		for {
			_ = m[1]
		}
	}()
	// 写
	go func() {
		for {
			m[2] = 2
		}
	}()
	select {}
}

// [orcaman/concurrent-map: a thread-safe concurrent map for go](https://github.com/orcaman/concurrent-map)
// 更小粒度的sync.Map
// 用map+读写锁解决并发写问题，在map的数据很大时，一把锁会导致大并发的客户端争抢锁。针对这种问题，我们在内部使用多个锁，每个区间共享一把锁，这样就减少了数据共享一把锁带来的性能影响
