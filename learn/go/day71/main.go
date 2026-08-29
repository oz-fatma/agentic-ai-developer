package main

import (
	"fmt"
	"log/slog"
	"os"
)

func main() {
	fmt.Println("Day 71: structured logging with log/slog")

	logger := slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}))
	slog.SetDefault(logger)

	slog.Info("service started",
		"day", 71,
		"module", "observability",
	)
	slog.Warn("demo warning", "retries", 0)
	slog.Debug("debug detail", "key", "value") // below Info, not shown
}
