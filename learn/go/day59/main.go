package main

import (
	"fmt"
	"log"

	"github.com/oz-fatma/agentic-ai-developer/learn/go/day59/internal/domain"
	"github.com/oz-fatma/agentic-ai-developer/learn/go/day59/internal/service"
)

func main() {
	fmt.Println("Day 59: domain services")

	orders := service.NewOrderService()
	o, err := orders.Place(domain.Order{Customer: "Ada", Item: "book", Qty: 2, UnitPrice: 15})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Order placed: id=%d total=$%.2f\n", o.ID, o.Total)

	o2, err := orders.Place(domain.Order{Customer: "Bob", Item: "pen", Qty: 0, UnitPrice: 2})
	if err != nil {
		fmt.Println("Expected validation error:", err)
	}
	_ = o2
}
