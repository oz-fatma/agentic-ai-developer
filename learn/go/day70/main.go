package main

import (
	"fmt"
	"os/exec"
)

func main() {
	fmt.Println("Day 70: advanced testing recap")
	fmt.Println("Run: go test ./day70/...")
	fmt.Println("Run: go test -bench=. ./day70/...")

	out, _ := exec.Command("go", "test", "./day70/...").CombinedOutput()
	fmt.Print(string(out))
}
