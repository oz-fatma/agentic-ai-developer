package main

import (
	"fmt"
	"sync"

	"github.com/oz-fatma/agentic-ai-developer/learn/go/day33/internal/store"
)

func main() {
	fmt.Println("Day 33: thread-safe in-memory store with interface")

	s := store.NewMemoryStore()
	var wg sync.WaitGroup
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(n int) {
			defer wg.Done()
			s.Create(fmt.Sprintf("task-%d", n))
		}(i)
	}
	wg.Wait()

	fmt.Printf("Created %d tasks concurrently\n", len(s.List()))
	fmt.Println("Run: go test ./day33/...")
}
