package pokecache

import (
	"sync"
	"time"
)

type cacheEntry struct {
	createdAt	time.Time
	val			[]byte
}

type Cache struct {
	cacheMap	map[string]cacheEntry
	mu			*sync.Mutex
}

func NewCache(interval time.Duration) Cache {
	c := Cache{
        cacheMap: make(map[string]cacheEntry),
        mu:   &sync.Mutex{},
    }

	go c.reapLoop(interval)
	return c
}

func (c *Cache) Add(key string, val []byte) {
	c.mu.Lock()
	defer c.mu.Unlock()
	obj := cacheEntry{
		createdAt:  time.Now(),
		val:		val,
	}
	c.cacheMap[key] = obj
	
}

func (c *Cache) Get(key string) ([]byte, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()
	cmap, ok := c.cacheMap[key]
	return cmap.val, ok
}

func (c *Cache) reapLoop(interval time.Duration) {
	now := time.Now()
	c.mu.Lock()
	defer c.mu.Unlock()
	for cmap := range c.cacheMap {
		timeout := c.cacheMap[cmap].createdAt.Before(now.Add(-interval))
		if timeout {
			delete(c.cacheMap, cmap)
		}
	}
}