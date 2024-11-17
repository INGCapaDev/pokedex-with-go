package pokecache

import (
	"sync"
	"time"
)

func NewCache(interval time.Duration) Cache {
	c := Cache{
		cache: map[string]CacheEntry{},
		mutex: &sync.Mutex{},
	}

	go c.reapLoop(interval)

	return c
}

func (c *Cache) Add(key string, value []byte) {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	c.cache[key] = CacheEntry{
		createdAt: time.Now(),
		value:     value,
	}
}

func (c *Cache) Get(key string) ([]byte, bool) {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	entry, ok := c.cache[key]
	return entry.value, ok
}

func (c *Cache) reapLoop(interval time.Duration) {
	ticker := time.NewTicker(interval)

	for range ticker.C {
		c.reap(time.Now().UTC(), interval)
	}
}

func (c *Cache) reap(now time.Time, lifetime time.Duration) {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	for key, entry := range c.cache {
		if entry.createdAt.Before(now.Add(-lifetime)) {
			delete(c.cache, key)
		}
	}
}
