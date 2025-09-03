package main

import "sync"

func main() {

}

// [Go语言 | 从并发模式看channel使用技巧 | 飞雪无情的博客](https://www.flysnow.org/2020/08/04/golang-goroutine-channel)

// 1、前提是从10个数据里读取任意5个
// 2、初始化的chan大小是10，但是通过for循环只存放了5个true
// 3、然后对chan循环读取数据，如果是true就开启go协程获取数据，如果是false就终止这次循环
// 4、当前在这之前还会判断下是否已经成功获取了5个，如果是的话，直接跳出整个for循环
// 5、通过readerIndex每次尝试获取一个数据，如果成功赛一个false到chan中，如果失败则塞个true
// 6、这样不成功的readerIndex不再尝试读取，失败了就通过true标记尝试读取下一个readerIndex
// 7、通过chan这种巧妙的方式不断循环，直到成功读取5个，或者把10个数据都读一遍为止
// 8、最终再基于是否成功读取到5个数据，做最终的判断，是返回成功数据，还是错误

// Read reads from readers in parallel. Returns p.dataBlocks number of bufs.
func (p *parallelReader) Read(dst [][]byte) ([][]byte, error) {
	newBuf := dst
	// 省略不太相关代码
	var newBufLK sync.RWMutex

	// 省略无关
	// channel开始创建，要发挥作用了。这里记住几个数字：
	// readTriggerCh大小是10，p.dataBlocks大小是5
	readTriggerCh := make(chan bool, len(p.readers))
	for i := 0; i < p.dataBlocks; i++ {
		// Setup read triggers for p.dataBlocks number of reads so that it reads in parallel.
		readTriggerCh <- true
	}

	healRequired := int32(0) // Atomic bool flag.
	readerIndex := 0
	var wg sync.WaitGroup
	// readTrigger 为 true, 意味着需要用disk.ReadAt() 读取下一个数据
	// readTrigger 为 false, 意味着读取成功了，不再需要读取
	for readTrigger := range readTriggerCh {
		newBufLK.RLock()
		canDecode := p.canDecode(newBuf)
		newBufLK.RUnlock()
		// 判断是否有5个成功的，如果有，退出for循环
		if canDecode {
			break
		}
		// 读取次数上限，不能大于10
		if readerIndex == len(p.readers) {
			break
		}
		// 成功了，退出本次读取
		if !readTrigger {
			continue
		}
		wg.Add(1)
		// 并发读取数据
		go func(i int) {
			defer wg.Done()
			// 省略不太相关代码
			_, err := rr.ReadAt(p.buf[bufIdx], p.offset)
			if err != nil {
				// 省略不太相关代码
				// 失败了，标记为true，触发下一个读取.
				readTriggerCh <- true
				return
			}
			newBufLK.Lock()
			newBuf[bufIdx] = p.buf[bufIdx]
			newBufLK.Unlock()
			// 成功了，标记为false，不再读取
			readTriggerCh <- false
		}(readerIndex)
		// 控制次数，同时用来作为索引获取和存储数据
		readerIndex++
	}
	wg.Wait()

	// 最终结果判断，如果OK了就正确返回，如果有失败的，返回error信息。
	if p.canDecode(newBuf) {
		p.offset += p.shardSize
		if healRequired != 0 {
			return newBuf, errHealRequired
		}
		return newBuf, nil
	}

	return nil, errErasureReadQuorum
}
