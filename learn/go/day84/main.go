package main

import (
	"context"
	"fmt"
	"sync/atomic"

	"github.com/oz-fatma/agentic-ai-developer/learn/go/day84/internal/queue"
)

func main() {
	fmt.Println("Day 84: worker queue")

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	wq := queue.NewWorkerQueue(3, 10)
	wq.Start(ctx)

	var processed atomic.Int32
	for i := 1; i <= 6; i++ {
		n := i
		wq.Submit(func() {
			processed.Add(1)
			fmt.Printf("processed job %d\n", n)
		})
	}

	wq.CloseAndWait()
	fmt.Println("total processed:", processed.Load())
	fmt.Println("workers:", wq.Workers())
}
