package main

import "container/list"

/*
 * @lc app=leetcode id=146 lang=golang
 *
 * [146] LRU Cache
 */

// @lc code=start
// type dlistNode struct {
//     prev  *dlistNode
//     next  *dlistNode
//     key   int
//     val   int
//     hnext *dlistNode
// }

type LRUCache struct {
	cap int
	l   *list.List
	m   map[int]*list.Element
}

type Pair struct {
	key   int
	value int
}

func Constructor(capacity int) LRUCache {
	return LRUCache{
		cap: capacity,
		l:   new(list.List),
		m:   make(map[int]*list.Element, capacity),
	}
}

func (this *LRUCache) Get(key int) int {
	if node, ok := this.m[key]; ok {
		val := node.Value.(*list.Element).Value.(Pair).value
		this.l.MoveToFront(node)
		return val
	}
	return -1
}

func (this *LRUCache) Put(key int, value int) {
	if node, ok := this.m[key]; ok {
		// val := node.Value.(*list.Element).Value.(Pair).value
		this.l.MoveToFront(node)
		node.Value.(*list.Element).Value = Pair{key: key, value: value}
	} else {
		if this.l.Len() == this.cap {
			idx := this.l.Back().Value.(*list.Element).Value.(Pair).key
			delete(this.m, idx)
			this.l.Remove(this.l.Back())
		}
		node := &list.Element{
			Value: Pair{
				key:   key,
				value: value,
			},
		}
		ptr := this.l.PushFront(node)
		this.m[key] = ptr
	}
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
// @lc code=end
