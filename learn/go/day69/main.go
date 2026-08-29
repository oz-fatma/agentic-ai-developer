package main

import (
	"fmt"

	"github.com/oz-fatma/agentic-ai-developer/learn/go/day69/internal/validator"
)

func main() {
	fmt.Println("Day 69: testing quality patterns (table-driven, subtests, t.Helper)")
	fmt.Println("Run: go test ./day69/...")

	valid := validator.Signup{Email: "demo@example.com", Password: "secret123", Age: 25}
	if err := validator.ValidateSignup(valid); err != nil {
		panic(err)
	}
	fmt.Println("Valid signup accepted")

	invalid := validator.Signup{Email: "bad", Password: "x", Age: 10}
	if err := validator.ValidateSignup(invalid); err != nil {
		fmt.Println("Expected validation errors:", err)
	}
}
