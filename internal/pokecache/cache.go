package pokecache

import (
	"sync"
	"time"
)

type Cache struct {
	cache map[string]cacheEntry
	mu    sync.Mutex
}

type cacheEntry struct {
	createdAt time.Time
	val       []byte
}

func (c *Cache) Add(key string, val []byte) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.cache[key] = cacheEntry{time.Now(), val}
}

func (c *Cache) Get(key string) ([]byte, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()

	if val, ok := c.cache[key]; ok {
		return val.val, true
	}

	return []byte{}, false

}

func (c *Cache) reapLoop(interval time.Duration) {
	c.mu.Lock()
	defer c.mu.Unlock()

	ticker := time.NewTicker(interval)

	// go func() {
	// 	for ticker.C {
	// 		for _, entry  := range c.cache {
	// 			if entry.createdAt.
	// 		}
	// 	}
	// }()

}

func NewCache(interval time.Duration) {

	cache := Cache{}

	reapLoop(interval)

}
