package main

import (
	"fmt"
	"time"

	"github.com/oz-fatma/agentic-ai-developer/learn/go/day81/internal/cache"
)

func main() {
	fmt.Println("Day 81: in-memory cache with TTL")

	c := cache.NewTTLCache(50 * time.Millisecond)
	c.Set("user:1", "alice")
	c.Set("user:2", "bob")

	if v, ok := c.Get("user:1"); ok {
		fmt.Println("hit user:1 ->", v)
	}

	time.Sleep(60 * time.Millisecond)
	if _, ok := c.Get("user:1"); !ok {
		fmt.Println("user:1 expired after TTL")
	}

	if v, ok := c.Get("user:2"); ok {
		fmt.Println("user:2 still valid ->", v)
	}
	fmt.Println("active entries:", c.Len())
}
