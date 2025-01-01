package concs

import (
	"fmt"
	"github.com/sourcegraph/conc"
	"github.com/sourcegraph/conc/iter"
	"github.com/sourcegraph/conc/panics"
	"github.com/sourcegraph/conc/pool"
	"github.com/sourcegraph/conc/stream"
	"time"
)

// 并发处理一个切片：使用 iter.Map 函数对切片进行并发映射
func concMap(input []int, f func(*int) int) []int {
	return iter.Map(input, f)
}

// 使用 stream.Stream 处理有序的并发任务
func mapStream(in chan int, out chan int, f func(int) int) {
	s := stream.New().WithMaxGoroutines(10)
	for elem := range in {
		elem := elem
		s.Go(func() stream.Callback {
			res := f(elem)
			return func() { out <- res }
		})
	}
	s.Wait()
}

func startTheThing(wg *conc.WaitGroup) {
	wg.Go(func() {
		fmt.Println("Starting Thing...")
	})
}

// 使用 panics.Catcher 捕获 goroutine 中的 panic
func ExampleCatcher() {
	var pc panics.Catcher
	i := 0
	pc.Try(func() { i += 1 })
	pc.Try(func() { panic("abort!") })
	pc.Try(func() { i += 1 })

	rc := pc.Recovered()

	fmt.Println(i)
	fmt.Println(rc.Value.(string))
	// Output:
	// 2
	// abort!
}

// 用conc实现多线程轮流打印
func p() {
	const (
		numGoroutines = 4
		totalPrints   = 100
	)

	// 创建一个channel用于控制打印顺序
	ch := make(chan int, 1)
	wg := conc.NewWaitGroup()

	// 创建4个协程
	for i := 0; i < numGoroutines; i++ {
		id := i + 1
		wg.Go(func() {
			for j := id; j <= totalPrints; j += numGoroutines {
				// 等待轮到自己的回合
				current := <-ch
				if current%numGoroutines+1 == id {
					fmt.Println(id)
					time.Sleep(time.Second)
					// 通知下一个协程
					ch <- current + 1
				}
			}
		})
	}

	// 启动第一次打印
	ch <- 0

	// 等待所有协程完成
	wg.Wait()
}

// 并发处理流中的每个元素：使用 pool.New() 和 WithMaxGoroutines
func process(stream chan int) {
	p := pool.New().WithMaxGoroutines(10)
	for elem := range stream {
		elem := elem
		p.Go(func() {
			// handle(elem)
			fmt.Println("Processing ", elem)
		})
	}
	p.Wait()
}

func xxx() {
	p := pool.New().WithErrors()
	p.Go(func() error {
		// 任务逻辑
		return nil
	})
	if err := p.Wait(); err != nil {
		fmt.Println("Error:", err)
	}
}

// 使用 ResultPool 收集任务结果
func zzz() {
	var rp pool.ResultPool[int]
	rp.Go(func() int {
		return 42
	})
	results := rp.Wait()
	fmt.Println(results) // 输出: [42]
}
