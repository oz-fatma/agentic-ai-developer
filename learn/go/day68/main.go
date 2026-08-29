package main

import (
	"fmt"
	"os/exec"

	"github.com/oz-fatma/agentic-ai-developer/learn/go/day68/internal/mathutil"
)

func main() {
	fmt.Println("Day 68: test coverage demo")
	fmt.Println("Run: go test -cover ./day68/...")

	fmt.Println("Add(2,3) =", mathutil.Add(2, 3))
	fmt.Println("Divide(10,2) =", mustDiv(10, 2))

	out, err := exec.Command("go", "test", "-cover", "./day68/internal/mathutil").Output()
	if err != nil {
		fmt.Println("Coverage:", string(out))
	} else {
		fmt.Print(string(out))
	}
}

func mustDiv(a, b int) int {
	v, err := mathutil.Divide(a, b)
	if err != nil {
		panic(err)
	}
	return v
}
