package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

type Config struct {
	App      string   `json:"app"`
	Version  int      `json:"version"`
	Required []string `json:"required"`
}

type Summary struct {
	TotalWords    int            `json:"total_words"`
	UniqueWords   int            `json:"unique_words"`
	TopWords      map[string]int `json:"top_words"`
	ConfigApp     string         `json:"config_app"`
	ConfigVersion int            `json:"config_version"`
}

func loadConfig(path string) (Config, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return Config{}, fmt.Errorf("read config: %w", err)
	}
	var cfg Config
	if err := json.Unmarshal(data, &cfg); err != nil {
		return Config{}, fmt.Errorf("parse config: %w", err)
	}
	if cfg.App == "" || cfg.Version == 0 {
		return Config{}, fmt.Errorf("invalid config: missing required fields")
	}
	return cfg, nil
}

func wordCounts(text string) map[string]int {
	counts := make(map[string]int)
	for _, word := range strings.Fields(text) {
		counts[strings.ToLower(word)]++
	}
	return counts
}

func main() {
	cfg, err := loadConfig("day15/config.json")
	if err != nil {
		fmt.Println("config error:", err)
		return
	}

	text, err := os.ReadFile("day15/sample.txt")
	if err != nil {
		fmt.Println("read error:", err)
		return
	}

	counts := wordCounts(string(text))
	total := 0
	for _, n := range counts {
		total += n
	}

	summary := Summary{
		TotalWords:    total,
		UniqueWords:   len(counts),
		TopWords:      counts,
		ConfigApp:     cfg.App,
		ConfigVersion: cfg.Version,
	}

	out, err := json.MarshalIndent(summary, "", "  ")
	if err != nil {
		fmt.Println("marshal error:", err)
		return
	}

	if err := os.WriteFile("day15/summary.json", out, 0644); err != nil {
		fmt.Println("write error:", err)
		return
	}

	fmt.Println(string(out))
}
