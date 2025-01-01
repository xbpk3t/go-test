package concs

import (
	"fmt"
	"github.com/sourcegraph/conc"
	"sync"
)

type CyclicBarrier struct {
	count     int
	total     int
	mutex     sync.Mutex
	cond      *sync.Cond
	waitGroup conc.WaitGroup
}

func NewCyclicBarrier(total int) *CyclicBarrier {
	cb := &CyclicBarrier{total: total}
	cb.cond = sync.NewCond(&cb.mutex)
	return cb
}

func (cb *CyclicBarrier) Await() {
	cb.mutex.Lock()
	cb.count++
	if cb.count == cb.total {
		cb.count = 0
		cb.cond.Broadcast()
	} else {
		cb.cond.Wait()
	}
	cb.mutex.Unlock()
}

func UseCyclicBarrier() {
	cb := NewCyclicBarrier(3)

	for i := 0; i < 3; i++ {
		cb.waitGroup.Go(func() {
			fmt.Println("Task before barrier")
			cb.Await()
			fmt.Println("Task after barrier")
		})
	}

	cb.waitGroup.Wait()
}
