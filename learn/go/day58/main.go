package main

import (
	"fmt"

	"github.com/oz-fatma/agentic-ai-developer/learn/go/day58/internal/app"
)

func main() {
	fmt.Println("Day 58: dependency injection")
	fmt.Println("Run: go test ./day58/...")

	g := app.NewApp("World")
	fmt.Println(g.Run())
}
