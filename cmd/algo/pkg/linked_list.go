package main

import (
	"fmt"
)

func main() {
	ll := LinkedList{}
	ll.Insert(1)
	ll.Insert(2)
	ll.Insert(3)
	ll.Print()

	ll.Remove(2)
	ll.Print()
}

// 首先定义节点和链表结构体，没啥说的
// 需要注意的是，链表的元素为 head 节点，来作为端点
type Node struct {
	data interface{}
	next *Node
}

type LinkedList struct {
	head *Node
}

func (l *LinkedList) Len() int {
	cur := l.head
	count := 0
	for cur != nil {
		count++
		cur = cur.next
	}
	return count
}

func (l *LinkedList) Insert(data interface{}) {
	newNode := &Node{data: data}
	if l.head == nil {
		l.head = newNode
	} else {
		cur := l.head
		for cur.next != nil {
			cur = cur.next
		}
		cur.next = newNode
	}
}

func (l *LinkedList) Remove(data interface{}) {
	if l.head == nil {
		return
	}
	if l.head.data == data {
		l.head = l.head.next
		return
	}
	cur := l.head
	for cur.next != nil {
		if cur.next.data == data {
			cur.next = cur.next.next
			return
		}
		cur = cur.next
	}
}

func (l *LinkedList) Print() {
	if l.head == nil {
		fmt.Println("Empty list")
		return
	}
	cur := l.head
	for cur != nil {
		fmt.Println(cur.data)
		cur = cur.next
	}
}
