package main

import (
	"fmt"
	"time"
)

func main() {
	ch1 := make(chan string)
	ch2 := make(chan string)

	go func() {
		time.Sleep(50 * time.Millisecond)
		ch1 <- "fast"
	}()

	go func() {
		time.Sleep(200 * time.Millisecond)
		ch2 <- "slow"
	}()

	select {
	case msg := <-ch1:
		fmt.Println("received:", msg)
	case msg := <-ch2:
		fmt.Println("received:", msg)
	case <-time.After(100 * time.Millisecond):
		fmt.Println("timeout")
	}

	done := make(chan struct{})
	go func() {
		time.Sleep(30 * time.Millisecond)
		close(done)
	}()

	select {
	case <-done:
		fmt.Println("work cancelled")
	case <-time.After(100 * time.Millisecond):
		fmt.Println("timeout waiting for cancel")
	}

	jobs := make(chan int, 2)
	select {
	case jobs <- 1:
		fmt.Println("accepted job")
	default:
		fmt.Println("backpressure: dropped job")
	}
}
