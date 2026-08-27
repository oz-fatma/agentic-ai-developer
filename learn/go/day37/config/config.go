package config

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Config struct {
	Port        int
	Host        string
	LogLevel    string
	EnableDebug bool
}

func LoadFromEnv() (Config, error) {
	portStr := envOrDefault("APP_PORT", "8080")
	port, err := strconv.Atoi(portStr)
	if err != nil || port <= 0 || port > 65535 {
		return Config{}, fmt.Errorf("invalid APP_PORT %q", portStr)
	}

	host := envOrDefault("APP_HOST", "localhost")
	if strings.TrimSpace(host) == "" {
		return Config{}, fmt.Errorf("APP_HOST must not be empty")
	}

	logLevel := envOrDefault("LOG_LEVEL", "info")
	if !isValidLogLevel(logLevel) {
		return Config{}, fmt.Errorf("invalid LOG_LEVEL %q", logLevel)
	}

	debugStr := envOrDefault("ENABLE_DEBUG", "false")
	enableDebug, err := strconv.ParseBool(debugStr)
	if err != nil {
		return Config{}, fmt.Errorf("invalid ENABLE_DEBUG %q", debugStr)
	}

	return Config{
		Port:        port,
		Host:        host,
		LogLevel:    logLevel,
		EnableDebug: enableDebug,
	}, nil
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
