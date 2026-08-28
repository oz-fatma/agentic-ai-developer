package main

import (
	"fmt"

	"github.com/oz-fatma/agentic-ai-developer/learn/go/day60/internal/calc"
)

func main() {
	fmt.Println("Day 60: refactor practice")
	fmt.Println("Run: go test ./day60/...")

	fmt.Println("Sum(1..5) =", calc.SumRange(1, 5))
	fmt.Println("Avg(10,20,30) =", calc.Average(10, 20, 30))
	fmt.Println("IsPrime(17) =", calc.IsPrime(17))
}
