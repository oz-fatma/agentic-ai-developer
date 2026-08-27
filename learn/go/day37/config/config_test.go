package config

import "testing"

func TestLoadFromEnvDefaults(t *testing.T) {
	t.Setenv("APP_PORT", "")
	t.Setenv("APP_HOST", "")
	t.Setenv("LOG_LEVEL", "")
	t.Setenv("ENABLE_DEBUG", "")

	cfg, err := LoadFromEnv()
	if err != nil {
		t.Fatal(err)
	}
	if cfg.Port != 8080 || cfg.Host != "localhost" || cfg.LogLevel != "info" {
		t.Fatalf("unexpected defaults: %+v", cfg)
	}
}

func TestLoadFromEnvInvalidPort(t *testing.T) {
	t.Setenv("APP_PORT", "99999")
	_, err := LoadFromEnv()
	if err == nil {
		t.Fatal("expected error for invalid port")
	}
}

func TestLoadFromEnvCustom(t *testing.T) {
	t.Setenv("APP_PORT", "3000")
	t.Setenv("APP_HOST", "0.0.0.0")
	t.Setenv("LOG_LEVEL", "debug")
	t.Setenv("ENABLE_DEBUG", "true")

	cfg, err := LoadFromEnv()
	if err != nil {
		t.Fatal(err)
	}
	if cfg.Addr() != "0.0.0.0:3000" || !cfg.EnableDebug {
		t.Fatalf("unexpected config: %+v", cfg)
	}
}
