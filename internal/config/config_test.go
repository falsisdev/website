package config

import "testing"

func TestLoadUsesEnvironmentOverrides(t *testing.T) {
	t.Setenv("PORT", "9090")
	t.Setenv("APP_ENV", "production")

	cfg, err := Load()
	if err != nil {
		t.Fatalf("Load() returned an unexpected error: %v", err)
	}

	if cfg.Port != "9090" {
		t.Fatalf("expected port 9090, got %q", cfg.Port)
	}

	if cfg.Environment != "production" {
		t.Fatalf("expected environment production, got %q", cfg.Environment)
	}
}

func TestLoadDefaultsToDevelopmentValues(t *testing.T) {
	cfg, err := Load()
	if err != nil {
		t.Fatalf("Load() returned an unexpected error: %v", err)
	}

	if cfg.Port == "" {
		t.Fatal("expected default port to be set")
	}

	if cfg.Environment == "" {
		t.Fatal("expected default environment to be set")
	}
}
