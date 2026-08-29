package main

import (
	"fmt"

	"github.com/oz-fatma/agentic-ai-developer/learn/go/day95/internal/team"
)

func main() {
	fmt.Println("Day 95: team practices")
	fmt.Println("Run: go test ./day95/...")

	fmt.Print(team.ReleaseNotes("worker queue hardening"))
}
