package main

import (
	"fmt"

	"github.com/oz-fatma/agentic-ai-developer/learn/go/day96/internal/capstone"
)

func main() {
	fmt.Println("Day 96: capstone plan struct")

	plan := capstone.DefaultPlan()
	fmt.Println("title:", plan.Title)
	fmt.Println("milestones:", len(plan.Milestones))
	for _, m := range plan.Milestones {
		fmt.Printf("- [%s] %s: %s\n", m.Phase, m.Name, m.Description)
	}
	fmt.Println("completed:", plan.CompletedCount(), "/", len(plan.Milestones))
}
