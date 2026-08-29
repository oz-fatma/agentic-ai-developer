package main

import (
	"fmt"

	"github.com/oz-fatma/agentic-ai-developer/learn/go/day90/internal/perf"
)

func main() {
	fmt.Println("Day 90: performance practice")
	fmt.Println("Run: go test ./day90/...")

	out := perf.Summarize([]string{"cache", "queue", "pprof", "builder"}, 2)
	fmt.Print(out)
}
