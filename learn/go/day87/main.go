package main

import (
	"fmt"
	"testing"

	"github.com/oz-fatma/agentic-ai-developer/learn/go/day87/internal/bench"
)

func main() {
	fmt.Println("Day 87: memory and allocation benchmarks")
	fmt.Println("Run: go test -bench=. ./day87/...")

	result := testing.Benchmark(func(b *testing.B) {
		for b.Loop() {
			_ = bench.Builder(256)
		}
	})
	fmt.Println("Builder(256):", result.String())
}
