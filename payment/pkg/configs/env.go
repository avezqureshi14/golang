package configs

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	DB_URL string
	PORT   string
	SECRET string
}

func mustGetEnv(key string) string {
	val, exists := os.LookupEnv(key)
	if !exists || val == "" {
		log.Fatalf("❌ Environment variable %s is required but not set", key)
	}
	return val
}

func Load() Config {
	// Load .env (only for local dev)
	err := godotenv.Load()
	if err != nil {
		log.Println("⚠️ No .env file found, relying on system env")
	}

	return Config{
		DB_URL: mustGetEnv("DB_URL"),
		PORT:   mustGetEnv("PORT"),
		SECRET: mustGetEnv("SECRET"),
	}
}