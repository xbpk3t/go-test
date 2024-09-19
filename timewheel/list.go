package timewheel

import "fmt"

type timeWheelNode struct {
	value      interface{}
	prev, next *timeWheelNode
}

func newNode(v interface{}) *timeWheelNode {
	return &timeWheelNode{value: v}
}

type timeWheelList struct {
	head, tail *timeWheelNode
	size       int
}

func newList() *timeWheelList {
	return new(timeWheelList)
}

func (twl *timeWheelList) push(v interface{}) *timeWheelNode {
	n := newNode(v)
	if twl.head == nil {
		twl.head, twl.tail = n, n
		twl.size++
		return n
	}

	n.prev = twl.tail
	n.next = nil
	twl.tail.next = n
	twl.tail = n
	twl.size++
	return n
}

func (twl *timeWheelList) remove(n *timeWheelNode) {
	if n == nil {
		return
	}
	prev, next := n.prev, n.next
	if prev == nil {
		twl.head = next
	} else {
		prev.next = next
	}

	if next == nil {
		twl.tail = prev
	} else {
		next.prev = prev
	}

	// 释放内存
	n = nil
	twl.size--
}

func (twl *timeWheelList) String() (s string) {
	s = fmt.Sprintf("[%d]: ", twl.size)
	for cur := twl.head; cur != nil; cur = cur.next {
		s += fmt.Sprintf("%v <-> ", cur.value)
	}
	s += "<nil>"
	return
}
