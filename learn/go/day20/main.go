package main

import (
	"fmt"
	"sync"
	"time"
)

func fetch(id int, delay time.Duration, results chan<- string, wg *sync.WaitGroup) {
	defer wg.Done()
	time.Sleep(delay)
	results <- fmt.Sprintf("result-%d", id)
}

func pipeline() {
	numbers := make(chan int)
	squares := make(chan int)

	go func() {
		for i := 1; i <= 5; i++ {
			numbers <- i
		}
		close(numbers)
	}()

	go func() {
		for n := range numbers {
			squares <- n * n
		}
		close(squares)
	}()

	fmt.Println("pipeline:")
	for sq := range squares {
		fmt.Println(sq)
	}
}

func fetchWithTimeout(id int, delay time.Duration) (string, error) {
	resultCh := make(chan string, 1)
	go func() {
		time.Sleep(delay)
		resultCh <- fmt.Sprintf("result-%d", id)
	}()

	select {
	case res := <-resultCh:
		return res, nil
	case <-time.After(100 * time.Millisecond):
		return "", fmt.Errorf("timeout fetching %d", id)
	}
}

func main() {
	results := make(chan string, 3)
	var wg sync.WaitGroup

	delays := []time.Duration{30 * time.Millisecond, 50 * time.Millisecond, 80 * time.Millisecond}
	for i, d := range delays {
		wg.Add(1)
		go fetch(i+1, d, results, &wg)
	}

	wg.Wait()
	close(results)

	fmt.Println("downloads:")
	for r := range results {
		fmt.Println(r)
	}

	pipeline()

	if res, err := fetchWithTimeout(99, 150*time.Millisecond); err != nil {
		fmt.Println("timeout fetch:", err)
	} else {
		fmt.Println("timeout fetch:", res)
	}
}
