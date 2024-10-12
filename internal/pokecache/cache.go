package pokecache

import (
	"time"
)

type Cache struct {
	cache map[string]cacheEntry
	// mu    *sync.Mutex
}

type cacheEntry struct {
	createdAt time.Time
	val       []byte
}

func (c *Cache) Add(key string, val []byte) {

	c.cache[key] = cacheEntry{createdAt: time.Now().UTC(), val: val}
}

func (c *Cache) Get(key string) ([]byte, bool) {

	if val, ok := c.cache[key]; ok {
		return val.val, ok
	}

	return []byte{}, false

}

func (c *Cache) reap(interval time.Duration) {

	timeAgo := time.Now().UTC().Add(-interval)

	for key, val := range c.cache {
		if val.createdAt.Before(timeAgo) {
			delete(c.cache, key)
		}
	}

}

func (c *Cache) reapLoop(interval time.Duration) {
	ticker := time.NewTicker(interval)
	for range ticker.C {
		c.reap(interval)
	}
}

func NewCache(interval time.Duration) Cache {

	c := Cache{
		cache: make(map[string]cacheEntry),
	}
	go c.reapLoop(interval)
	return c

}
