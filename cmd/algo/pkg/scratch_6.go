package main

import "sync"

// 实现阻塞读且并发安全的 map
// GO ⾥⾯ MAP 如何实现 key 不存在 get 操作等待 直到 key 存在或者超时，保证并发安全，且需要实现以下接⼝：
// 看到阻塞协程第⼀个想到的就是 channel，题⽬中要求并发安全，那么必须⽤锁，还要实现多个 goroutine 读的
// 时候如果值不存在则阻塞，直到写⼊值，那么每个键值需要有⼀个阻塞 goroutine 的 channel。
func main() {

}

type Map struct {
	c   map[string]*entry
	rmx *sync.RWMutex
}

type entry struct {
	ch      chan struct{}
	value   interface{}
	isExist bool
}

func (m *Map) Out(key string, val interface{}) {
	m.rmx.Lock()
	defer m.rmx.Unlock()
	item, ok := m.c[key]
	if !ok {
		m.c[key] = &entry{
			value:   val,
			isExist: true,
		}
		return
	}
	item.value = val
	if !item.isExist {
		if item.ch != nil {
			close(item.ch)
			item.ch = nil
		}
	}
	return
}
