package config

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	DBPassword string
	DBUser     string
	DB         string
}

func New() *Config {
	_ = godotenv.Load()

	return &Config{
		DBPassword: getEnv("POSTGRES_PASSWORD", ""),
		DBUser:     "POSTGRES_USER",
		DB:         "POSTGRES_DB",
	}
}

func getEnv(key, defaultValue string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return defaultValue
}
