package main

import (
	"fmt"
	"time"

	"github.com/oz-fatma/agentic-ai-developer/learn/go/day85/internal/service"
)

func main() {
	fmt.Println("Day 85: cache + queue practice")
	fmt.Println("Run: go test ./day85/...")

	svc := service.NewCachedWorkerService(2, time.Minute)
	defer svc.Close()

	for i := 0; i < 2; i++ {
		v, cached, err := svc.Fetch("report:weekly")
		if err != nil {
			panic(err)
		}
		fmt.Printf("fetch %d -> %q cached=%v\n", i+1, v, cached)
	}
}
