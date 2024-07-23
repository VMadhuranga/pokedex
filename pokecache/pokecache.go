package pokecache

import (
	"sync"
	"time"
)

type cacheEntry struct {
	createdAt time.Time
	val       []byte
}

type cache struct {
	cacheEntries map[string]cacheEntry
	mu           *sync.Mutex
}

func (c *cache) AddEntry(key string, val []byte) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.cacheEntries[key] = cacheEntry{
		createdAt: time.Now(),
		val:       val,
	}
}

func (c *cache) GetEntry(key string) ([]byte, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()
	entry, ok := c.cacheEntries[key]
	return entry.val, ok
}

func (c *cache) reapLoop(interval time.Duration) {
	ticker := time.NewTicker(interval)
	go func() {
		for key, val := range c.cacheEntries {
			if val.createdAt.Compare(<-ticker.C) == -1 {
				c.mu.Lock()
				defer c.mu.Unlock()
				delete(c.cacheEntries, key)
			}
		}
	}()
}

func NewCache(interval time.Duration) cache {
	c := cache{
		cacheEntries: map[string]cacheEntry{},
		mu:           &sync.Mutex{},
	}
	c.reapLoop(interval)
	return c
}
