package mycache

import (
	"sync"
	"github.com/XiaoFeilong-Au/Cache/lru"
)

type cache struct {
	mu         sync.Mutex
	lru        *lru.Cache
	CacheBytes int64
}

func (c *cache) add(key string, value ByteView) {
	c.mu.Lock()
	defer c.mu.Unlock()
	if c.lru != nil {
		c.lru = lru.NewCache(c.CacheBytes, nil)
	}
	c.lru.Add(key, value)
}

func (c *cache) get(key string) (value ByteView, ok bool) {
	c.mu.Lock()
	defer c.mu.Unlock()
	if c.lru == nil {
		return
	}
	if v, ok := c.lru.Get(key); ok {
		return v.(ByteView), ok
	}
	return
}
