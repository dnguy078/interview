package main

import (
	"container/list"
)

// capacity lru cache (limit on keys)
// once reaches capacity, you have evict least recently touched

// hashmap if key exist
// key 2, value: node{key, value}
// doubly link list  (list) to determine the least recently used

// [newest{key1: newest}, second_oldest,  oldest]
// [newest, oldest]
type LRUCache struct {
	capacity int
	l        *list.List
	m        map[int]*list.Element
}

func Constructor(capacity int) LRUCache {
	return LRUCache{
		capacity: capacity,
		m:        make(map[int]*list.Element),
		l:        list.New(),
	}
}

// check if it exist inside hashmap
// if it does, then return the value + push that value to the front list
func (this *LRUCache) Get(key int) int {
	if val, found := this.m[key]; found {
		this.l.MoveToFront(val)

		val, _ := val.Value.(ListValue)
		return val.value
	}

	return -1
}

type ListValue struct {
	key   int
	value int
}

// insert map if hasnt reached capacity
// if reach capacity, evict the least recently (delete the end of the list)
//  delete hashmap
// if hasnt reached cap, insert into hashmap and push onto the front of the list
// if it exist, you can take pointer and push onto the front of the list and update the hashmap
func (this *LRUCache) Put(key int, value int) {
	if val, found := this.m[key]; found {
		this.l.MoveToFront(val)
		val.Value = ListValue{key, value}
		return
	}

	if len(this.m) != this.capacity {
		// insert into hashmap and push the node onto the list
		element := this.l.PushFront(ListValue{key, value})
		this.m[key] = element
	} else {
		// evict least recently used and delete map
		leastRecentlyUsed := this.l.Back()
		this.l.Remove(leastRecentlyUsed)
		val, _ := leastRecentlyUsed.Value.(ListValue)
		delete(this.m, val.key)

		// push new element on
		element := this.l.PushFront(ListValue{key, value})
		this.m[key] = element
	}

}
