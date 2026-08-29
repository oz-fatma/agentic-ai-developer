package main

import (
	"fmt"
	"log"

	"github.com/oz-fatma/agentic-ai-developer/learn/go/day67/internal/fixtures"
)

func main() {
	fmt.Println("Day 67: test fixtures from testdata/")
	fmt.Println("Run: go test ./day67/internal/fixtures/...")

	users, err := fixtures.LoadUsersFromDir("day67/testdata")
	if err != nil {
		log.Fatal(err)
	}
	for _, u := range users {
		fmt.Printf("  %s (%s)\n", u.Name, u.Role)
	}
}
