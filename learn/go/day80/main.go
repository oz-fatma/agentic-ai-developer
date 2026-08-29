package main

import "fmt"

func main() {
	fmt.Println("Day 80: Deploy & CI/CD recap")
	fmt.Println()
	fmt.Println("=== Deployment checklist ===")
	fmt.Println("1. go test ./... && go vet ./...")
	fmt.Println("2. docker build -f day76/Dockerfile -t learn-go .")
	fmt.Println("3. docker compose -f day77/docker-compose.yml up -d")
	fmt.Println("4. make -f day78/Makefile run-all")
	fmt.Println("5. Push triggers day79/ci.yml on GitHub Actions")
	fmt.Println()
	fmt.Println("=== Observability stack ===")
	fmt.Println("- Structured logs: log/slog (day71)")
	fmt.Println("- Metrics counters: day72/internal/metrics")
	fmt.Println("- Retry/backoff: day73, circuit breaker: day74, timeouts: day75")
	fmt.Println()
	fmt.Println("=== gRPC stack ===")
	fmt.Println("- Proto: internal/greeterpb/greeter.proto")
	fmt.Println("- In-memory gRPC: bufconn (days 62-65)")
	fmt.Println()
	fmt.Println("Days 61-80 complete. Run: go run ./day80")
}
