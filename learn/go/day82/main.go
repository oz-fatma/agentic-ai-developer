package main

import (
	"fmt"

	"github.com/oz-fatma/agentic-ai-developer/learn/go/day82/internal/cache"
)

func main() {
	fmt.Println("Day 82: LRU cache")

	c := cache.NewLRUCache(3)
	for _, kv := range [][2]string{{"a", "1"}, {"b", "2"}, {"c", "3"}, {"d", "4"}} {
		c.Set(kv[0], kv[1])
	}

	if v, ok := c.Get("a"); ok {
		fmt.Println("a still present ->", v)
	} else {
		fmt.Println("a evicted (LRU)")
	}

	if v, ok := c.Get("d"); ok {
		fmt.Println("d present ->", v)
	}
	fmt.Println("size:", c.Len())
}
