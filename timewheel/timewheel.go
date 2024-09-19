package timewheel

import (
	"fmt"
	"log"
	"sync"
	"sync/atomic"
	"time"
)

type TimeWheel struct {
	ticker *time.Ticker
	// 每次tick时长
	tickGap time.Duration
	// slot的数量
	slotNum int
	// 当前slot序号
	curSlot int
	// 槽数量
	slots []*timeWheelSlot
	// taskId -> taskPtr
	taskMap map[int64]*timeWheelNode
	// 自增id
	incrId int64
	// task缓冲channel
	taskChan chan *timeWheelTask
	// 读写锁
	lock sync.RWMutex
}

// 周期耗时
var cycleCost int64

func NewTimeWheel(tickGap time.Duration, slotNum int) *TimeWheel {
	tw := &TimeWheel{
		ticker:   time.NewTicker(tickGap),
		tickGap:  tickGap,
		slotNum:  slotNum,
		slots:    make([]*timeWheelSlot, 0, slotNum),
		taskMap:  make(map[int64]*timeWheelNode),
		taskChan: make(chan *timeWheelTask, 100),
		lock:     sync.RWMutex{},
	}

	cycleCost = int64(tw.tickGap * time.Duration(tw.slotNum))
	for i := 0; i < slotNum; i++ {
		tw.slots = append(tw.slots, newSlot(i))
	}
	go tw.turn()

	return tw
}

// 执行延时任务
func (tw *TimeWheel) After(timeout time.Duration, do doTask) (int64, chan interface{}) {
	if timeout < 0 {
		return -1, nil
	}

	t := newTask(timeout, 1, do)
	tw.locate(t, t.interval, false)
	tw.taskChan <- t
	return t.id, t.resChan
}

// 执行指定重试逻辑的重复任务
func (tw *TimeWheel) AfterPoints(timeoutUnit time.Duration, points []int64, do doTask) ([]int64, []chan interface{}) {
	if timeoutUnit < 0 || len(points) == 0 {
		return nil, nil
	}

	var tids []int64
	var resChs []chan interface{}
	for _, point := range points {
		timeout := timeoutUnit * time.Duration(point)
		tid, resCh := tw.After(timeout, do)
		tids = append(tids, tid)
		resChs = append(resChs, resCh)
	}

	return tids, resChs
}

// 执行重复任务
func (tw *TimeWheel) Repeat(interval time.Duration, repeatN int64, do doTask) ([]int64, chan interface{}) {
	if interval <= 0 || repeatN < 1 {
		return nil, nil
	}

	costSum := repeatN * int64(interval) // 全部任务耗时
	cycleSum := costSum / cycleCost      // 全部任务执行总圈数
	trip := cycleSum / cycle(interval)   // 每个任务多少圈才执行一次

	var tids []int64
	var resChs []chan interface{}
	if trip > 0 {
		gap := interval
		for step := int64(0); step < cycleCost; step += int64(interval) { // 每隔 interval 放置执行 trip 次的 task
			t := newTask(interval, trip, do)
			tw.locate(t, gap, false)
			tw.taskChan <- t
			gap += interval
			tids = append(tids, t.id)
			resChs = append(resChs, t.resChan)
		}
	}

	// 计算余下几个任务时需重头开始计算
	gap := time.Duration(0)
	remain := (costSum % cycleCost) / int64(interval)
	for i := 0; i < int(remain); i++ {
		t := newTask(interval, 1, do)
		t.cycles = trip + 1
		tw.locate(t, gap, true)
		tw.taskChan <- t
		gap += interval
		tids = append(tids, t.id)
		resChs = append(resChs, t.resChan)
	}

	allDone := make(chan interface{}, 1)
	go func(doneChs []chan interface{}) {
		for _, ch := range doneChs {
			for range ch {
			}
		}
		allDone <- nil // 等待全部子任务完成
	}(resChs)
	return tids, allDone
}

// 更新任务
func (tw *TimeWheel) Update(tids []int64, interval time.Duration, repeatN int64, do doTask) ([]int64, chan interface{}) {
	if len(tids) == 0 || interval <= 0 || repeatN < 1 {
		return nil, nil
	}

	if repeatN == 1 {
		if !tw.Cancel(tids[0]) {
			// return nil, nil // 按需处理
		}
		newTid, resCh := tw.After(interval, do)
		return []int64{newTid}, resCh
	}

	// 重复任务需全部取消
	for _, tid := range tids {
		if !tw.Cancel(tid) {
			// return nil, nil // 按需处理
		}
	}
	return tw.Repeat(interval, repeatN, do)
}

// 取消任务
func (tw *TimeWheel) Cancel(tid int64) bool {
	tw.lock.Lock()
	defer tw.lock.Unlock()

	node, ok := tw.taskMap[tid]
	if !ok {
		return false // 任务已执行完毕或不存在
	}

	t := node.value.(*timeWheelTask)
	t.resChan <- nil
	close(t.resChan) // 避免资源泄漏

	slot := tw.slots[t.slotIdx]
	slot.tasks.remove(node)
	delete(tw.taskMap, tid)
	return true
}

// 接收 task 并定时运行 slot 中的任务
func (tw *TimeWheel) turn() {
	idx := 0
	for {
		select {
		case <-tw.ticker.C:
			idx %= tw.slotNum
			tw.lock.Lock()
			tw.curSlot = idx // 锁粒度要细，不要重叠
			tw.lock.Unlock()
			tw.handleSlotTasks(idx)
			idx++
		case t := <-tw.taskChan:
			tw.lock.Lock()
			// fmt.Println(t)
			slot := tw.slots[t.slotIdx]
			tw.taskMap[t.id] = slot.tasks.push(t)
			tw.lock.Unlock()
		}
	}
}

// 计算 task 所在 slot 的编号
func (tw *TimeWheel) locate(t *timeWheelTask, gap time.Duration, restart bool) {
	tw.lock.Lock()
	defer tw.lock.Unlock()
	if restart {
		t.slotIdx = tw.convSlotIdx(gap)
	} else {
		t.slotIdx = tw.curSlot + tw.convSlotIdx(gap)
	}
	t.id = tw.slot2Task(t.slotIdx)
}

// 执行指定 slot 中的所有任务
func (tw *TimeWheel) handleSlotTasks(idx int) {
	var expNodes []*timeWheelNode

	tw.lock.RLock()
	slot := tw.slots[idx]
	for node := slot.tasks.head; node != nil; node = node.next {
		task := node.value.(*timeWheelTask)
		task.cycles--
		if task.cycles > 0 {
			continue
		}
		// 重复任务恢复 cycle
		if task.repeat > 0 {
			task.cycles = cycle(task.interval)
			task.repeat--
		}

		// 不重复任务或重复任务最后一次执行都将移除
		if task.repeat == 0 {
			expNodes = append(expNodes, node)
		}
		go func() {
			defer func() {
				if err := recover(); err != nil {
					log.Printf("task exec paic: %v", err) // 出错暂只记录
				}
			}()

			var res interface{}
			if task.do != nil {
				res = task.do()
			}
			task.resChan <- res
			if task.repeat == 0 {
				close(task.resChan)
			}
		}()
	}
	tw.lock.RUnlock()

	tw.lock.Lock()
	for _, n := range expNodes {
		slot.tasks.remove(n)                            // 剔除过期任务
		delete(tw.taskMap, n.value.(*timeWheelTask).id) //
	}
	tw.lock.Unlock()
}

// 在指定 slot 中无重复生成新 task id
func (tw *TimeWheel) slot2Task(slotIdx int) int64 {
	return int64(slotIdx)<<32 + atomic.AddInt64(&tw.incrId, 1) // 保证去重优先
}

// 反向获取 task 所在的 slot
func (tw *TimeWheel) task2Slot(taskIdx int64) int {
	return int(taskIdx >> 32)
}

// 将指定间隔计算到指定的 slot 中
func (tw *TimeWheel) convSlotIdx(gap time.Duration) int {
	timeGap := gap % time.Duration(cycleCost)
	slotGap := int(timeGap / tw.tickGap)
	return int(slotGap % tw.slotNum)
}

func (tw *TimeWheel) String() (s string) {
	for _, slot := range tw.slots {
		if slot.tasks.size > 0 {
			s += fmt.Sprintf("[%v]\t", slot.tasks)
		}
	}
	return
}
