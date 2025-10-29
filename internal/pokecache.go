package pokecache

import (
	"sync"
	"time"
	"fmt"
)

type cacheEntry struct {
	createdAt	time.Time
	val			[]byte
}

type Cache struct {
	stored	map[string]cacheEntry
	mu		*sync.Mutex
}

func NewCache(interval time.Duration) Cache {
	fmt.Println("New Cache")
	c := Cache{
        cache: make(map[string]cacheEntry),
        mux:   &sync.Mutex{},
    }

	go c.reapLoop(interval)
	return c
}

func (*c Cache) Add(key string, val []byte) {
	fmt.Println("adding %s to cache", key)
	obj := cacheEntry{
		createdAt:  time.Now(),
		val:		val,
	}
	c.mu.Lock()
	c.cacheMap[key] = obj
	c.mu.Unlock()
}

func (*c Cache) Get(key string) ([]byte, bool) {
	fmt.Println("getting cache %s", key)
	c.mu.Lock()
	cmap, ok := c.cacheMap[key]
	c.mu.Unlock()
	if !ok {return nil, ok}
	return cmap.val, ok
}

func (*c Cache) reapLoop(interval time.Duration) {
	fmt.Println("Running ReapLoop")
	now := time.Now
	c.mu.Lock()
	for cmap := range c.cacheMap {
		time := now.Sub(c.cacheMap[cmap].createdAt)
		if time >= interval {
			delete(c.cacheMap, cmap)
			fmt.Println("Deleted %s from cache", cmap)
		}
	}
	c.mu.Unlock()
}