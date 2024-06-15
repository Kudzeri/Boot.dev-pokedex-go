package pokecache

import (
	"sync"
	"time"
)

type Cache struct {
	Cache map[string]cacheEntry
	mux   *sync.Mutex
}

type cacheEntry struct {
	Val       []byte
	createdAt time.Time
}

func NewCache(interval time.Duration) Cache {
	c := Cache{
		Cache: make(map[string]cacheEntry),
		mux:   &sync.Mutex{},
	}
	go c.ReapLoop(interval)
	return c
}

func (c *Cache) Add(key string, val []byte) {
	c.mux.Lock()
	defer c.mux.Unlock()
	c.Cache[key] = cacheEntry{
		Val:       val,
		createdAt: time.Now().UTC(),
	}
}

func (c Cache) Get(key string) ([]byte, bool) {
	c.mux.Lock()
	defer c.mux.Unlock()
	cacheE, ok := c.Cache[key]
	return cacheE.Val, ok
}

func (c Cache) ReapLoop(interval time.Duration) {
	ticker := time.NewTicker(interval)
	for range ticker.C {
		c.Reap(interval)
	}
}

func (c Cache) Reap(interval time.Duration) {
	c.mux.Lock()
	defer c.mux.Unlock()
	expirationTime := time.Now().UTC().Add(-interval)
	for k, v := range c.Cache {
		if v.createdAt.Before(expirationTime) {
			delete(c.Cache, k)
		}
	}
}
