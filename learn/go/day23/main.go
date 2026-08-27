package main

import (
	"fmt"

	"github.com/oz-fatma/agentic-ai-developer/learn/go/day23/sliceutil"
)

func main() {
	nums := []int{1, 2, 2, 3, 4, 4, 5}
	fmt.Println("Day 23: benchmarks and Example tests")
	fmt.Println("Sum:", sliceutil.Sum(nums))
	fmt.Println("Unique:", sliceutil.Unique(nums))
	fmt.Println("FilterEven:", sliceutil.FilterEven(nums))
	fmt.Println("Run: go test ./day23/...")
}
