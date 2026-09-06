package config

import (
	"fmt"
	"os"
)

type Config struct {
	Port        string
	Environment string
	Host        string
}

func Load() (Config, error) {
	cfg := Config{
		Port:        getEnv("PORT", "8080"),
		Environment: getEnv("APP_ENV", "development"),
		Host:        getEnv("HOST", "0.0.0.0"),
	}

	if cfg.Port == "" {
		return Config{}, fmt.Errorf("PORT cannot be empty")
	}
	if cfg.Environment == "" {
		return Config{}, fmt.Errorf("APP_ENV cannot be empty")
	}

	return cfg, nil
}

func getEnv(key, fallback string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return fallback
}
