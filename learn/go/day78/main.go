package main

import (
	"fmt"
	"os/exec"
)

func main() {
	fmt.Println("Day 78: Makefile automation")
	fmt.Println("Run: make -f day78/Makefile help")

	cmd := exec.Command("make", "-f", "day78/Makefile", "help")
	out, err := cmd.Output()
	if err != nil {
		fmt.Println("make help:", err)
		return
	}
	fmt.Print(string(out))
}
