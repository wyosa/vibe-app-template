package config

import (
	"os"
	"testing"
)

func TestLoad(t *testing.T) {
	t.Run("missing DATABASE_URL", func(t *testing.T) {
		os.Unsetenv("DATABASE_URL")
		_, err := Load()
		if err == nil {
			t.Fatal("expected error when DATABASE_URL is not set")
		}
	})

	t.Run("valid config", func(t *testing.T) {
		t.Setenv("DATABASE_URL", "postgres://user:pass@localhost:5432/db?sslmode=disable")
		t.Setenv("HTTP_PORT", "9090")
		t.Setenv("APP_ENV", "test")

		cfg, err := Load()
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if cfg.DatabaseURL != "postgres://user:pass@localhost:5432/db?sslmode=disable" {
			t.Errorf("unexpected DatabaseURL: %s", cfg.DatabaseURL)
		}
		if cfg.HTTPPort != 9090 {
			t.Errorf("unexpected HTTPPort: %d", cfg.HTTPPort)
		}
		if cfg.Env != "test" {
			t.Errorf("unexpected Env: %s", cfg.Env)
		}
	})

	t.Run("defaults", func(t *testing.T) {
		t.Setenv("DATABASE_URL", "postgres://user:pass@localhost:5432/db?sslmode=disable")
		os.Unsetenv("HTTP_PORT")
		os.Unsetenv("APP_ENV")

		cfg, err := Load()
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if cfg.HTTPPort != 8080 {
			t.Errorf("expected default HTTPPort 8080, got %d", cfg.HTTPPort)
		}
		if cfg.Env != "development" {
			t.Errorf("expected default Env development, got %s", cfg.Env)
		}
	})
}
