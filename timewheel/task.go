package timewheel

import (
	"fmt"
	"time"
)

type doTask func() interface{}

// 每个slot链表中的task
type timeWheelTask struct {
	// 在slot中的索引位置
	id int64
	// 所属slot
	slotIdx int
	// 任务执行间隔
	interval time.Duration
	// 延迟指定圈后执行
	cycles int64
	// 执行任务的内容
	do doTask
	// 传递任务执行结果
	resChan chan interface{}
	// 任务重复执行次数
	repeat int64
}

// 创建新任务
func newTask(interval time.Duration, repeat int64, do func() interface{}) *timeWheelTask {
	return &timeWheelTask{
		interval: interval,
		cycles:   cycle(interval),
		repeat:   repeat,
		do:       do,
		resChan:  make(chan interface{}, 1),
	}
}

func cycle(interval time.Duration) (n int64) {
	n = 1 + int64(interval)/cycleCost
	return
}

func (twt *timeWheelTask) String() string {
	return fmt.Sprintf("[slot]: %d [interval]: %.feishu [repeat]: %d [cycle]: %dth [idx]:%d ",
		twt.slotIdx, twt.interval.Seconds(), twt.repeat, twt.cycles, twt.id)
}
