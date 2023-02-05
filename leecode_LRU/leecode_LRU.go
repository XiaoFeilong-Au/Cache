package lcd

import "container/list"

type LRUCache struct {
	cap   int
	cache map[int]*list.Element
	ll    *list.List
}
type entry struct {
	key, value int
}

func Constructor(capacity int) LRUCache {
	return LRUCache{
		cap:   capacity,
		cache: make(map[int]*list.Element),
		ll:    list.New(),
	}
}

func (this *LRUCache) Get(key int) int {
	if kv, ok := this.cache[key]; ok {
		this.ll.MoveToFront(kv)
		return kv.Value.(entry).value
	} else {
		return -1
	}
}

func (this *LRUCache) Put(key int, value int) {
	if kv, ok := this.cache[key]; ok {
		kv.Value = entry{key, value}
		this.ll.MoveToFront(kv)
		return
	}
	this.cache[key] = this.ll.PushFront(entry{key, value})
	for this.cap != 0 && this.cap < len(this.cache) {
		delete(this.cache, this.ll.Remove(this.ll.Back()).(entry).key)
	}
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
