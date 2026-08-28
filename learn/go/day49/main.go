package main

import (
	"fmt"
	"log"
	"sync"
	"time"

	"github.com/oz-fatma/agentic-ai-developer/learn/go/day49/internal/db"
)

func main() {
	fmt.Println("Day 49: connection pooling")

	conn, err := db.OpenPooled("file::memory:?cache=shared")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	if err := db.Migrate(conn); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Pool: maxOpen=%d maxIdle=%d\n", conn.Stats().MaxOpenConnections, 2)

	var wg sync.WaitGroup
	for i := 0; i < 8; i++ {
		wg.Add(1)
		go func(n int) {
			defer wg.Done()
			if err := db.InsertMetric(conn, fmt.Sprintf("worker-%d", n)); err != nil {
				log.Println("insert:", err)
			}
			time.Sleep(10 * time.Millisecond)
		}(i)
	}
	wg.Wait()

	stats := conn.Stats()
	fmt.Printf("Stats: open=%d inUse=%d idle=%d waitCount=%d\n",
		stats.OpenConnections, stats.InUse, stats.Idle, stats.WaitCount)

	count, err := db.CountMetrics(conn)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Metrics inserted:", count)
}
