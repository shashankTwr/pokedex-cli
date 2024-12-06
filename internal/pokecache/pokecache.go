package pokecache

import (
	"sync"
	"time"
)

type Cache struct{
	cache map[string]cacheEntry
	mu *sync.Mutex
}

type cacheEntry struct {
	createdAt time.Time
	val []byte
}



func NewCache(interval time.Duration) Cache {
	c := Cache{
		cache: make(map[string]cacheEntry),
		mu : &sync.Mutex{},
	}

	go c.reapLoop(interval)
	return c
}




func(c *Cache) Add(key string, val []byte){
	c.mu.Lock()
	defer c.mu.Unlock() 
	c.cache[key] = cacheEntry{
		createdAt: time.Now().Local().In(time.FixedZone("Asia/Kolkata", 5.5*60*60)),
		val: val,
	}
}

func(c *Cache) Get(key string) ([]byte, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()
	cEntry, ok := c.cache[key]
	return cEntry.val, ok
}

func (c *Cache) reapLoop(interval time.Duration) {
	ticker := time.NewTicker(interval)
	for range ticker.C {
		c.reap(time.Now().Local().In(time.FixedZone("Asia/Kolkata", 5.5*60*60)), interval)
	}
	
}

func (c *Cache) reap(now time.Time, last time.Duration){
	c.mu.Lock()
	defer c.mu.Unlock()
	for key, entry := range c.cache {
		if entry.createdAt.Local().In(time.FixedZone("Asia/Kolkata", 5.5*60*60)).Before(now.Add(-last))  {
			delete(c.cache, key)
		}
	}
}