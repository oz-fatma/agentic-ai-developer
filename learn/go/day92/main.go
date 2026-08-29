package main

import (
	"fmt"

	"github.com/oz-fatma/agentic-ai-developer/learn/go/day92/internal/version"
)

func main() {
	fmt.Println("Day 92: semver and version in main")
	fmt.Println("version:", version.Current)
	fmt.Println("major:", version.Current.Major)
}
