package config

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Config struct {
	Port     int
	Host     string
	LogLevel string
}

func LoadFromEnv() (Config, error) {
	portStr := envOrDefault("APP_PORT", "8080")
	port, err := strconv.Atoi(portStr)
	if err != nil || port <= 0 || port > 65535 {
		return Config{}, fmt.Errorf("invalid APP_PORT %q", portStr)
	}

	host := envOrDefault("APP_HOST", "127.0.0.1")
	logLevel := envOrDefault("LOG_LEVEL", "info")
	if !isValidLogLevel(logLevel) {
		return Config{}, fmt.Errorf("invalid LOG_LEVEL %q", logLevel)
	}

	return Config{Port: port, Host: host, LogLevel: logLevel}, nil
}

func (c Config) Addr() string {
	return fmt.Sprintf("%s:%d", c.Host, c.Port)
}

func envOrDefault(key, fallback string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return fallback
}

func isValidLogLevel(level string) bool {
	switch strings.ToLower(level) {
	case "debug", "info", "warn", "error":
		return true
	default:
		return false
	}
}
