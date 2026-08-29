package main

import (
	"fmt"
	"time"

	"github.com/oz-fatma/agentic-ai-developer/learn/go/day88/internal/tune"
)

func main() {
	fmt.Println("Day 88: concurrency tuning")

	jobs := 40
	bestWorkers := 1
	var best time.Duration = 1<<63 - 1

	for workers := 1; workers <= 8; workers++ {
		elapsed, count := tune.Process(workers, jobs)
		fmt.Printf("workers=%d elapsed=%s processed=%d\n", workers, elapsed.Round(time.Millisecond), count)
		if elapsed < best {
			best = elapsed
			bestWorkers = workers
		}
	}
	fmt.Println("fastest config: workers =", bestWorkers)
}
