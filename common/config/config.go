package config

import (
	"os"

	"github.com/joho/godotenv"
)

func LoadEnvFile(path string) error {
	if path == "" {
		path = ".env"
	}

	return godotenv.Load(path)
}

func GetEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}

	return fallback
}
