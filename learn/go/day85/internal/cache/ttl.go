package cache

import (
	"sync"
	"time"
)

type entry struct {
	value   string
	expires time.Time
}

type TTLCache struct {
	mu    sync.RWMutex
	items map[string]entry
	ttl   time.Duration
}

func NewTTLCache(ttl time.Duration) *TTLCache {
	return &TTLCache{items: make(map[string]entry), ttl: ttl}
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
	if !ok || time.Now().After(e.expires) {
		if ok {
			c.mu.Lock()
			delete(c.items, key)
			c.mu.Unlock()
		}
		return "", false
	}
	return e.value, true
}
