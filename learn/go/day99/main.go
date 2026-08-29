package main

import (
	"fmt"

	"github.com/oz-fatma/agentic-ai-developer/learn/go/day99/internal/ci"
)

func main() {
	fmt.Println("Day 99: CI pipeline summary")
	fmt.Println(ci.Summary())
	for i, step := range ci.Pipeline() {
		fmt.Printf("%2d. %-14s %s\n", i+1, step.Name+":", step.Command)
	}
}
