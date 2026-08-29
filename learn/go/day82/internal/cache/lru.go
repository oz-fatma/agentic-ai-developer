package cache

import "container/list"

// LRUCache evicts the least recently used entry when capacity is exceeded.
type LRUCache struct {
	cap   int
	items map[string]*list.Element
	order *list.List
}

type pair struct {
	key   string
	value string
}

func NewLRUCache(capacity int) *LRUCache {
	return &LRUCache{
		cap:   capacity,
		items: make(map[string]*list.Element),
		order: list.New(),
	}
}

func (c *LRUCache) Get(key string) (string, bool) {
	el, ok := c.items[key]
	if !ok {
		return "", false
	}
	c.order.MoveToFront(el)
	return el.Value.(pair).value, true
}

func (c *LRUCache) Set(key, value string) {
	if el, ok := c.items[key]; ok {
		el.Value = pair{key: key, value: value}
		c.order.MoveToFront(el)
		return
	}
	el := c.order.PushFront(pair{key: key, value: value})
	c.items[key] = el
	if c.order.Len() > c.cap {
		back := c.order.Back()
		if back != nil {
			p := back.Value.(pair)
			delete(c.items, p.key)
			c.order.Remove(back)
		}
	}
}

func (c *LRUCache) Len() int {
	return len(c.items)
}
