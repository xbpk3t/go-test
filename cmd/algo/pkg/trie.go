package main

import "fmt"

func main() {
	trie := NewTrie()
	// trie.Put("apple")

	for _, kw := range []string{"apple", "banana"} {
		trie.Put(kw)
	}

	fmt.Println(trie.Search("apple"))    // true
	fmt.Println(trie.Search("applexxx")) // false
	fmt.Println(trie.Search("app"))      // false
}

type Trie struct {
	root *TrieNode
}

type TrieNode struct {
	children map[rune]*TrieNode
	isEnd    bool
}

func NewTrie() *Trie {
	return &Trie{root: &TrieNode{
		children: make(map[rune]*TrieNode),
		isEnd:    false,
	}}
}

func (t Trie) Put(word string) {
	node := t.root
	for _, ch := range word {
		if node.children[ch] == nil {
			node.children[ch] = &TrieNode{
				children: make(map[rune]*TrieNode),
				isEnd:    false,
			}
		}
		node = node.children[ch]
	}
	node.isEnd = true
}

func (t Trie) Search(word string) bool {
	node := t.root
	for _, ch := range word {
		if node.children[ch] == nil {
			return false
		}
		node = node.children[ch]
	}
	return node.isEnd
}

func (t Trie) StartWith(word string) bool {
	node := t.root
	for _, ch := range word {
		if node.children[ch] == nil {
			return false
		}
		node = node.children[ch]
	}
	return true
}
