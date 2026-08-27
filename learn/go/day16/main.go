package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup

	for i := 1; i <= 3; i++ {
		wg.Add(1)
		go func(n int) {
			defer wg.Done()
			fmt.Println("worker", n)
		}(i)
	}

	wg.Wait()

	var unsafeCounter int
	for i := 0; i < 1000; i++ {
		go func() {
			unsafeCounter++
		}()
	}
	fmt.Println("unsafe counter (race possible):", unsafeCounter)
}
