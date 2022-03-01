/*
 * @lc app=leetcode.cn id=146 lang=golang
 *
 * [146] LRU 缓存
 */
package leetcode

// @lc code=start
import "container/list"

type entry struct {
	key   int
	value int
}

type LRUCache struct {
	capacity int
	ll       *list.List
	cache    map[int]*list.Element
}

func Constructor(capacity int) LRUCache {
	return LRUCache{
		capacity: capacity,
		ll:       list.New(),
		cache:    make(map[int]*list.Element),
	}
}

func (this *LRUCache) Get(key int) int {
	if v, ok := this.cache[key]; ok {
		this.ll.MoveToFront(v)
		return v.Value.(*entry).value
	}
	return -1
}

func (this *LRUCache) Put(key int, value int) {
	if v, ok := this.cache[key]; ok {
		v.Value.(*entry).value = value
		this.ll.MoveToFront(v)
		return
	}

	el := this.ll.PushFront(&entry{key: key, value: value})
	this.cache[key] = el
	if this.ll.Len() > this.capacity {
		lastElement := this.ll.Back()
		this.ll.Remove(lastElement)
		delete(this.cache, lastElement.Value.(*entry).key)
	}
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
// @lc code=end
