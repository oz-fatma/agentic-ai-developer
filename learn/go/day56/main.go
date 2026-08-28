package main

import (
	"fmt"

	"github.com/oz-fatma/agentic-ai-developer/learn/go/day56/internal/app"
	"github.com/oz-fatma/agentic-ai-developer/learn/go/day56/internal/domain"
)

func main() {
	fmt.Println("Day 56: project layout")

	greeter := app.NewGreeter()
	msg := greeter.Greet(domain.User{Name: "Fatma", Role: "developer"})
	fmt.Println(msg)
	fmt.Println("Layout: day56/internal/{domain,app}")
}
