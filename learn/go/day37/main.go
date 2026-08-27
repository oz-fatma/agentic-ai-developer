package main

import (
	"fmt"
	"os"

	"github.com/oz-fatma/agentic-ai-developer/learn/go/day37/config"
)

func main() {
	fmt.Println("Day 37: env config struct with validation")

	os.Setenv("APP_PORT", "9090")
	os.Setenv("LOG_LEVEL", "debug")
	os.Setenv("ENABLE_DEBUG", "true")

	cfg, err := config.LoadFromEnv()
	if err != nil {
		fmt.Println("config error:", err)
		return
	}

	fmt.Printf("Listening on %s (log=%s, debug=%v)\n", cfg.Addr(), cfg.LogLevel, cfg.EnableDebug)
	fmt.Println("Run: go test ./day37/...")
}
