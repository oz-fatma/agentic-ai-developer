package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

var once sync.Once
var setupCount int

func setup() {
	setupCount++
	fmt.Println("expensive setup ran")
}

func getSetup() {
	once.Do(setup)
}

func main() {
	var wg sync.WaitGroup
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			getSetup()
		}()
	}
	wg.Wait()
	fmt.Println("setupCount:", setupCount)

	var mu sync.Mutex
	counter := 0
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			mu.Lock()
			counter++
			mu.Unlock()
		}()
	}
	wg.Wait()
	fmt.Println("mutex counter:", counter)

	var atomicCounter int64
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			atomic.AddInt64(&atomicCounter, 1)
		}()
	}
	wg.Wait()
	fmt.Println("atomic counter:", atomicCounter)
}
