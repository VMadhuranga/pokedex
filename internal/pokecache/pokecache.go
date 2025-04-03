package pokecache

import (
	"sync"
	"time"
)

type cacheEntry struct {
	createdAt time.Time
	val       []byte
}

type Cache struct {
	entries map[string]cacheEntry
	mu      *sync.Mutex
}

func (c *Cache) Add(key string, val []byte) {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.entries[key] = cacheEntry{
		createdAt: time.Now(),
		val:       val,
	}
}

func (c *Cache) Get(key string) ([]byte, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()

	entry, ok := c.entries[key]
	if !ok {
		return []byte{}, false
	}

	return entry.val, true
}

func (c *Cache) reapLoop() {
	c.mu.Lock()
	defer c.mu.Unlock()

	for k := range c.entries {
		delete(c.entries, k)
	}
}

func NewCache(interval time.Duration) *Cache {
	c := Cache{
		entries: make(map[string]cacheEntry),
		mu:      &sync.Mutex{},
	}

	ticker := time.NewTicker(interval)
	go func() {
		for range ticker.C {
			c.reapLoop()
		}
	}()

	return &c
}
