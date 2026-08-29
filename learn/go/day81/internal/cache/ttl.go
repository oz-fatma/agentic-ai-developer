package cache

import (
	"sync"
	"time"
)

type entry struct {
	value   string
	expires time.Time
}

// TTLCache is an in-memory key-value store with per-entry expiration.
type TTLCache struct {
	mu    sync.RWMutex
	items map[string]entry
	ttl   time.Duration
}

func NewTTLCache(ttl time.Duration) *TTLCache {
	return &TTLCache{
		items: make(map[string]entry),
		ttl:   ttl,
	}
}

func (c *TTLCache) Set(key, value string) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.items[key] = entry{value: value, expires: time.Now().Add(c.ttl)}
}

func (c *TTLCache) Get(key string) (string, bool) {
	c.mu.RLock()
	e, ok := c.items[key]
	c.mu.RUnlock()
	if !ok {
		return "", false
	}
	if time.Now().After(e.expires) {
		c.mu.Lock()
		delete(c.items, key)
		c.mu.Unlock()
		return "", false
	}
	return e.value, true
}

func (c *TTLCache) Len() int {
	c.mu.Lock()
	defer c.mu.Unlock()
	now := time.Now()
	n := 0
	for k, e := range c.items {
		if now.After(e.expires) {
			delete(c.items, k)
			continue
		}
		n++
	}
	return n
}
