package main

import (
	"fmt"
	"time"

	"github.com/oz-fatma/agentic-ai-developer/learn/go/day83/internal/pubsub"
)

func main() {
	fmt.Println("Day 83: pub/sub with channels")

	b := pubsub.NewBroker()
	orders := b.Subscribe("orders")
	alerts := b.Subscribe("orders")

	go func() {
		for msg := range orders {
			fmt.Println("orders subscriber:", msg)
		}
	}()

	b.Publish("orders", "order-101 created")
	b.Publish("orders", "order-102 shipped")

	select {
	case msg := <-alerts:
		fmt.Println("alerts subscriber:", msg)
	case <-time.After(100 * time.Millisecond):
	}

	b.Close("orders")
	fmt.Println("broker closed topic orders")
}
