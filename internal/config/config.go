package config

import (
	"fmt"
	"os"
)

type Config struct {
	Port             string
	Environment      string
	Host             string
	GitHubUsername   string
	GitHubToken      string
	SanityProjectID  string
	SanityDataset    string
	SanityAPIVersion string
	SanityToken      string
}

func Load() (Config, error) {
	cfg := Config{
		Port:             getEnv("PORT", "8080"),
		Environment:      getEnv("APP_ENV", "development"),
		Host:             getEnv("HOST", "0.0.0.0"),
		GitHubUsername:   getEnv("GITHUB_USERNAME", "falsisdev"),
		GitHubToken:      getEnv("GITHUB_TOKEN", ""),
		SanityProjectID:  getEnv("SANITY_PROJECT_ID", ""),
		SanityDataset:    getEnv("SANITY_DATASET", ""),
		SanityAPIVersion: getEnv("SANITY_API_VERSION", "2024-01-01"),
		SanityToken:      getEnv("SANITY_TOKEN", ""),
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
