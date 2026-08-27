package main

import (
	"fmt"

	"github.com/oz-fatma/agentic-ai-developer/learn/go/day21/mathutil"
)

func main() {
	fmt.Println("Day 21: unit tests")
	fmt.Println("Add(2, 3) =", mathutil.Add(2, 3))
	fmt.Println("Multiply(4, 5) =", mathutil.Multiply(4, 5))
	fmt.Println("IsEven(4) =", mathutil.IsEven(4))
	fmt.Println("Abs(-7) =", mathutil.Abs(-7))
	fmt.Println("Run: go test ./day21/...")
}
