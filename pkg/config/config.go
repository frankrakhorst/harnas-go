package config

import (
	"fmt"
	"os"
)

// Config holds runtime configuration.
type Config struct {
	LogLevel string
}

// Load reads config from environment with defaults.
func Load() (Config, error) {
	cfg := Config{LogLevel: "info"}
	if v := os.Getenv("HARNAS_LOG_LEVEL"); v != "" {
		cfg.LogLevel = v
	}
	if cfg.LogLevel == "" {
		return cfg, fmt.Errorf("invalid log level")
	}
	return cfg, nil
}
