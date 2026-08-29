package main

import (
	"fmt"
	"os"
	"os/exec"
)

func main() {
	fmt.Println("Day 91: Makefile targets")
	fmt.Println("Available: make run test fmt vet bench (see day91/Makefile)")

	makePath, err := exec.LookPath("make")
	if err != nil {
		fmt.Println("make not found in PATH; open day91/Makefile manually")
		return
	}
	out, err := exec.Command(makePath, "-f", "day91/Makefile", "-n", "test").CombinedOutput()
	if err != nil {
		fmt.Println("dry-run test target failed:", err)
		return
	}
	fmt.Println("dry-run `make test`:")
	fmt.Print(string(out))
	fmt.Println("cwd:", mustWd())
}

func mustWd() string {
	wd, err := os.Getwd()
	if err != nil {
		return "."
	}
	return wd
}
