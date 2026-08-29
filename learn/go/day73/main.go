package main

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/oz-fatma/agentic-ai-developer/learn/go/day73/internal/retry"
)

func main() {
	fmt.Println("Day 73: retry with exponential backoff")

	attempts := 0
	err := retry.Do(context.Background(), retry.Config{
		MaxAttempts: 4,
		BaseDelay:   50 * time.Millisecond,
		MaxDelay:    200 * time.Millisecond,
	}, func() error {
		attempts++
		fmt.Println("attempt", attempts)
		if attempts < 3 {
			return errors.New("transient failure")
		}
		return nil
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("Succeeded after", attempts, "attempts")
}
