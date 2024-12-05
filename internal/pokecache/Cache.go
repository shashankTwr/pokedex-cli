package pokecache

import (
	"sync"
	"time"
)

func NewCache(interval time.Duration) *Cache {
	c := Cache{}
	c.cacheMap = make(map[string]CacheEntry)
	c.mu = &sync.Mutex{}
	go c.reapLoop(interval)
	return &c
}


type Cache struct{
	cacheMap map[string]CacheEntry
	mu *sync.Mutex
}

func(c *Cache) Add(key string, val []byte){ 
	cE := CacheEntry{}
	cE.val = val
	cE.createdAt = time.Now()
	c.mu.Lock() // writing to cacheMap 
	c.cacheMap[key] = cE
	c.mu.Unlock()
}

func(c *Cache) Get(key string) ([]byte, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()
	cEntry, ok := c.cacheMap[key]
	if !ok {
		return nil, false
	}
	return cEntry.val, true
}

func (c *Cache) reapLoop(interval time.Duration) {
	for  {
		c.mu.Lock()
		for key, cE := range c.cacheMap {
			if time.Since(cE.createdAt) > interval {
				delete(c.cacheMap, key)
			}
		}
		c.mu.Unlock()
		time.Sleep(interval)
	}
	
}