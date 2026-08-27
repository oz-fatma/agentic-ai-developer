package main

import (
	"fmt"

	"github.com/oz-fatma/agentic-ai-developer/learn/go/day22/stringutil"
)

func main() {
	fmt.Println("Day 22: table-driven tests with t.Run")
	fmt.Println("Reverse:", stringutil.Reverse("hello"))
	fmt.Println("TitleCase:", stringutil.TitleCase("golang"))
	fmt.Println("Truncate:", stringutil.Truncate("hello world", 5))
	fmt.Println("Run: go test ./day22/...")
}
