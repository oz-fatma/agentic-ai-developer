package main

import (
	"fmt"

	"github.com/oz-fatma/agentic-ai-developer/learn/go/day94/internal/api"
)

func main() {
	fmt.Println("Day 94: API docs comment patterns")
	fmt.Println("Run: go doc ./day94/internal/api")

	fmt.Println(api.Greeter("reviewer"))
	fmt.Println("See godoc-style comments on UserService and Greeter")
}
