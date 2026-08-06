package config

import (
	"fmt"
)

type Config struct {
	HTTPPort    int64
	DatabaseURL string
	Env         string
}

func Load() (*Config, error) {
	databaseURL, err := getStringFromEnv("DATABASE_URL")
	if err != nil {
		return nil, fmt.Errorf("config: DATABASE_URL: %w", err)
	}

	httpPort, err := getIntValueFromEnv("HTTP_PORT", 8080)
	if err != nil {
		return nil, fmt.Errorf("config: HTTP_PORT: %w", err)
	}

	env := getStringFromEnvOrDefault("APP_ENV", "development")

	return &Config{
		HTTPPort:    httpPort,
		DatabaseURL: databaseURL,
		Env:         env,
	}, nil
}
