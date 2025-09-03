package main

import (
	"fmt"
	"math/rand"
)

const (
	maxLevel = 16
	p        = 0.5
)

// 从这两个 struct 可以看到，skiplist 的本质就是链表
type node struct {
	key   int
	value int
	next  []*node
}

func newNode(key, value, level int) *node {
	return &node{key: key, value: value, next: make([]*node, level)}
}

type skipList struct {
	head *node
}

func newSkipList() *skipList {
	return &skipList{
		head: newNode(0, 0, maxLevel),
	}
}

// 随机 level
func (r *skipList) randomLevel() int {
	level := 1
	for rand.Float32() < p && level < maxLevel {
		level++
	}
	return level
}

func (r *skipList) insert(key, value int) {
	update := make([]*node, maxLevel)
	cur := r.head
	for i := maxLevel - 1; i >= 0; i-- {
		for cur.next[i] != nil && cur.next[i].key < key {
			cur = cur.next[i]
		}
		update[i] = cur
	}
	level := r.randomLevel()
	newNode := newNode(key, value, level)
	for i := 0; i < level; i++ {
		newNode.next[i] = update[i].next[i]
		update[i].next[i] = newNode
	}
}

func (r *skipList) search(key int) *node {
	cur := r.head
	for i := maxLevel - 1; i >= 0; i-- {
		for cur.next[i] != nil && cur.next[i].key < key {
			cur = cur.next[i]
		}
	}
	if cur.next[0] != nil && cur.next[0].key == key {
		return cur.next[0]
	}
	return nil
}

func main() {
	sl := newSkipList()
	sl.insert(3, 5)
	sl.insert(1, 10)
	sl.insert(2, 7)

	node := sl.search(2)
	if node != nil {
		fmt.Println("Found:", node.value)
	} else {
		fmt.Println("Not Found")
	}
}
