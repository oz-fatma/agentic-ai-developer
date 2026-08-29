package main

import (
	"fmt"

	"github.com/oz-fatma/agentic-ai-developer/learn/go/day97/internal/hardening"
)

func main() {
	fmt.Println("Day 97: hardening checklist")
	for category, items := range hardening.ByCategory() {
		fmt.Println(category + ":")
		for _, item := range items {
			fmt.Println("  -", item)
		}
	}
}
