package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	DBURL          string
	JWTSecret      string
	AuthServiceURL string
}

func Load() *Config {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, using ENV vars")
	}

	return &Config{
		DBURL:          getEnv("DB_URL", "postgres://user:pass@localhost:5432/db?sslmode=disable"),
		JWTSecret:      getEnv("JWT_SECRET", "super-secret-key"),
		AuthServiceURL: getEnv("AUTH_SERVICE_URL", "http://localhost:8080"),
	}
}

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}
