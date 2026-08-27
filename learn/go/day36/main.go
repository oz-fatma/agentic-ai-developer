package main

import (
	"context"
	"fmt"
	"time"
)

func slowWork(ctx context.Context, name string, delay time.Duration) error {
	select {
	case <-time.After(delay):
		fmt.Printf("%s finished\n", name)
		return nil
	case <-ctx.Done():
		return ctx.Err()
	}
}

func main() {
	fmt.Println("Day 36: context.Context cancel/timeout")

	ctx, cancel := context.WithCancel(context.Background())
	go func() {
		time.Sleep(50 * time.Millisecond)
		cancel()
	}()

	if err := slowWork(ctx, "cancelled job", 200*time.Millisecond); err != nil {
		fmt.Println("cancel demo:", err)
	}

	timeoutCtx, timeoutCancel := context.WithTimeout(context.Background(), 80*time.Millisecond)
	defer timeoutCancel()

	if err := slowWork(timeoutCtx, "timeout job", 200*time.Millisecond); err != nil {
		fmt.Println("timeout demo:", err)
	}

	deadlineCtx, deadlineCancel := context.WithDeadline(context.Background(), time.Now().Add(100*time.Millisecond))
	defer deadlineCancel()

	if err := slowWork(deadlineCtx, "deadline job", 30*time.Millisecond); err != nil {
		fmt.Println("deadline demo:", err)
	} else {
		fmt.Println("deadline job completed in time")
	}
}
