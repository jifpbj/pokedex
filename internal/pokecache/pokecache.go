// Package pokecache is a simple in-memory cache for pokedex data
package pokecache

import (
	"sync"
	"time"
)

type Cache struct {
	x        map[string]cacheEntry
	mutex    sync.Mutex
	interval time.Duration
}

type cacheEntry struct {
	createdAt time.Time
	val       []byte
}

func NewCache(interval time.Duration) *Cache {
	c := Cache{
		make(map[string]cacheEntry),
		sync.Mutex{},
		interval,
	}

	go c.reapLoop(interval)

	return &c
}

func (c *Cache) Add(key string, val []byte) {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	c.x[key] = cacheEntry{
		createdAt: time.Now().UTC(),
		val:       val,
	}
}

func (c *Cache) Get(key string) ([]byte, bool) {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	entry, ok := c.x[key]
	if !ok {
		return nil, false
	}
	return entry.val, true
}

func (c *Cache) reapLoop(interval time.Duration) {
	ticker := time.NewTicker(interval)
	defer ticker.Stop()
	for range ticker.C {
		c.mutex.Lock()
		now := time.Now()
		for key, entry := range c.x {
			if now.Sub(entry.createdAt) > c.interval {
				delete(c.x, key)
			}
		}
		c.mutex.Unlock()
	}
}
