package main

import (
	"fmt"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

// 扇出的编码模型比较简单，本文不多研究，我们提供一个扇入编程示例。
// 创建一个生成器函数 generate，通过 interval 参数控制消息生成频率。生成器返回消息 channel mc与停止 channel sc，停止 channel 用于停止生成器任务。
func main() {
	// create two sample message and stop channels
	mc1, sc1 := generate("message from generator 1", 200*time.Millisecond)
	mc2, sc2 := generate("message from generator 2", 300*time.Millisecond)

	// multiplex message channels
	mmc, wg1 := multiplex(mc1, mc2)

	// create errs channel for graceful shutdown
	errs := make(chan error)

	// wait for interrupt or terminate signal
	go func() {
		sc := make(chan os.Signal, 1)
		signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM)
		errs <- fmt.Errorf("%s signal received", <-sc)
	}()

	// wait for multiplexed messages
	wg2 := &sync.WaitGroup{}
	wg2.Add(1)
	go func() {
		defer wg2.Done()

		for m := range mmc {
			fmt.Println(m)
		}
	}()

	// wait for errors
	if err := <-errs; err != nil {
		fmt.Println(err.Error())
	}

	// stop generators
	stopGenerating(mc1, sc1)
	stopGenerating(mc2, sc2)
	wg1.Wait()

	// close multiplexed messages channel
	close(mmc)
	wg2.Wait()
}

func multiplex(mcs ...chan string) (chan string, *sync.WaitGroup) {
	mmc := make(chan string)
	wg := &sync.WaitGroup{}

	for _, mc := range mcs {
		wg.Add(1)

		go func(mc chan string, wg *sync.WaitGroup) {
			defer wg.Done()

			for m := range mc {
				mmc <- m
			}
		}(mc, wg)
	}

	return mmc, wg
}

func stopGenerating(mc chan string, sc chan struct{}) {
	sc <- struct{}{}

	close(mc)
}

func generate(message string, interval time.Duration) (chan string, chan struct{}) {
	mc := make(chan string)
	sc := make(chan struct{})

	go func() {
		defer func() {
			close(sc)
		}()

		for {
			select {
			case <-sc:
				return
			default:
				time.Sleep(interval)

				mc <- message
			}
		}
	}()

	return mc, sc
}
