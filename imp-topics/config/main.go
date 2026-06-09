package config

import (
	"fmt"
	"os"
	"sync"
)

var (
	globalConfig map[string]string
	once         sync.Once
)

// Init loads config once (thread-safe)
func Init() {
	once.Do(func() {
		globalConfig = map[string]string{
			"APP_ENV":   getEnv("APP_ENV", "dev"),
			"DB_HOST":   getEnv("DB_HOST", "localhost"),
			"DB_PORT":   getEnv("DB_PORT", "5432"),
			"LOG_LEVEL": getEnv("LOG_LEVEL", "info"),
		}
	})
}

// Get returns a config value
func Get(key string) string {
	Init()
	return globalConfig[key]
}

// helper
func getEnv(key, fallback string) string {
	if val := os.Getenv(key); val != "" {
		return val
	}
	return fallback
}

func main() {
	fmt.Println(Get("APP_ENV"))
	fmt.Println(Get("DB_HOST"))
}
