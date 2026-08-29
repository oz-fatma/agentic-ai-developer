package main

import (
	"fmt"

	"github.com/oz-fatma/agentic-ai-developer/learn/go/day93/internal/checklist"
)

func main() {
	fmt.Println("Day 93: code review checklist")
	for i, item := range checklist.Items {
		fmt.Printf("%2d. %s\n", i+1, item)
	}
}
