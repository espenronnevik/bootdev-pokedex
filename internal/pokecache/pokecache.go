package pokecache

import (
	"sync"
	"time"
)

type Cache struct {
	mu   sync.Mutex
	data map[string]cacheEntry
}

type cacheEntry struct {
	createdAt time.Time
	val       []byte
}

func NewCache(interval time.Duration) *Cache {
	cache := Cache{data: make(map[string]cacheEntry)}
	go cache.reapLoop(interval)
	return &cache
}

func (c *Cache) Add(key string, val []byte) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.data[key] = cacheEntry{createdAt: time.Now(), val: val}
}

func (c *Cache) Get(key string) ([]byte, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()
	v, exists := c.data[key]
	return v.val, exists
}

func (c *Cache) reapLoop(interval time.Duration) {
	ticker := time.NewTicker(interval)
	for {
		<-ticker.C
		c.mu.Lock()

		for k, v := range c.data {
			if time.Since(v.createdAt) > interval {
				delete(c.data, k)
			}
		}

		c.mu.Unlock()
	}
}
