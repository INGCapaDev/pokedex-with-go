package pokecache

import (
	"sync"
	"time"
)

type CacheEntry struct {
	createdAt time.Time
	value     []byte
}

type Cache struct {
	cache map[string]CacheEntry
	mutex *sync.Mutex
}
