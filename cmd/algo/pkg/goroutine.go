package main

import (
	"fmt"
	"math/rand"
)

// 写代码实现两个 goroutine，其中一个产生随机数并写入到 go channel 中，另外一个从 channel 中读取数字并打印到标准输出。最终输出五个随机数。
func main() {
	ch := make(chan int)
	done := make(chan bool)
	go func() {
		for {
			select {
			case ch <- rand.Intn(5): // Create and send random number into channel
			case <-done: // If receive signal on done channel - Return
				return
			default:
			}
		}
	}()

	go func() {
		for i := 0; i < 5; i++ {
			fmt.Println("Rand Number = ", <-ch) // Print number received on standard output
		}
		done <- true // Send Terminate Signal and return
		return
	}()
	<-done // Exit Main when Terminate Signal received
}
