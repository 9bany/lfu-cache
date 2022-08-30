package main

import "fmt"

type Node struct {
	key   int
	value int
	prev  *Node
	next  *Node
}

func NewNode(key, value int) *Node {
	return &Node{
		key:   key,
		value: value,
		prev:  nil,
		next:  nil,
	}
}

type LFUCache struct {
	capacity int
	values   map[int]*Node
	head     *Node
	tail     *Node
}

func Constructor(capacity int) LFUCache {

	head := NewNode(-1, -1)
	tail := NewNode(-1, -1)
	head.next = tail
	tail.prev = head

	return LFUCache{
		capacity: capacity,
		head:     head,
		tail:     tail,
		values:   map[int]*Node{},
	}
}

func (cache *LFUCache) Get(key int) int {
	node := cache.values[key]

	if node == nil {
		return -1
	}

	fmt.Println("get -> ", key, node.value)
	node.prev.next = node.next
	node.next.prev = node.prev

	cache.moveToHead(node)

	return node.value
}

func (cache *LFUCache) Put(key int, value int) {
	fmt.Println("put -> ", key, value)
	node := cache.values[key]
	if node != nil {
		node.value = value
		node.prev.next = node.next
		node.next.prev = node.prev
		cache.moveToHead(node)
		return
	}

	if len(cache.values) == cache.capacity {
		fmt.Println("delete", cache.tail.prev.key)
		delete(cache.values, cache.tail.prev.key)
		cache.tail.prev = cache.tail.prev.prev
		cache.tail.prev.next = cache.tail
	}

	node = NewNode(key, value)
	cache.values[key] = node
	cache.moveToHead(node)
}

func (cache *LFUCache) moveToHead(node *Node) {
	node.prev = cache.head
	node.next = cache.head.next
	node.next.prev = node
	cache.head.next = node
}
