package config

import (
	"os"
	"strconv"
)

type Config struct {
	App AppConfig
}

type AppConfig struct {
	Name string
	Env  string
	Port int
}

func Load() Config {
	return Config{
		App: AppConfig{
			Name: getEnv("APP_NAME", "media-engine"),
			Env:  getEnv("APP_ENV", "development"),
			Port: getEnvInt("APP_PORT", 8080),
		},
	}
}

func getEnv(key string, fallback string) string {
	value := os.Getenv(key)

	if value == "" {
		return fallback
	}

	return value
}

func getEnvInt(key string, fallback int) int {
	value := os.Getenv(key)

	if value == "" {
		return fallback
	}

	parsed, err := strconv.Atoi(value)

	if err != nil {
		return fallback
	}

	return parsed
}
