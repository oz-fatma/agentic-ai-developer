package main

import (
	"errors"
	"fmt"
	"time"

	"github.com/oz-fatma/agentic-ai-developer/learn/go/day74/internal/breaker"
)

func main() {
	fmt.Println("Day 74: circuit breaker pattern")

	cb := breaker.New(2, 200*time.Millisecond)
	fail := errors.New("upstream down")

	for i := 1; i <= 4; i++ {
		err := cb.Call(func() error {
			if cb.State() == breaker.HalfOpen {
				return nil
			}
			return fail
		})
		fmt.Printf("call %d: state=%v err=%v\n", i, cb.State(), err)
	}
}
