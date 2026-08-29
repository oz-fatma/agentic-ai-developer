package main

import (
	"fmt"

	"github.com/oz-fatma/agentic-ai-developer/learn/go/day72/internal/metrics"
)

func main() {
	fmt.Println("Day 72: metrics counter demo")

	requests := metrics.NewCounter("http_requests_total")
	errors := metrics.NewCounter("http_errors_total")

	for i := 0; i < 5; i++ {
		requests.Inc()
	}
	errors.Inc()

	fmt.Println(requests)
	fmt.Println(errors)
	fmt.Printf("error rate: %.0f%%\n", float64(errors.Value())/float64(requests.Value())*100)
}
