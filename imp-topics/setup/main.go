package main

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	AppName string
	Port    string
	DBUrl   string
}

func LoadConfig() *Config {
	err := godotenv.Load()
	if err != nil {
		log.Println("No .env file found, using system env")
	}

	return &Config{
		AppName: getEnv("APP_NAME", "DefaultApp"),
		Port:    getEnv("PORT", "8800"),
		DBUrl:   getEnv("DB_URL", ""),
	}

}

func getEnv(key string, fallBack string) string {
	val, ok := os.LookupEnv(key)
	if ok {
		return val
	}
	return fallBack
}
func main() {
	cfg := LoadConfig()
	fmt.Println("App Name", cfg.AppName)
	fmt.Println("DB URL", cfg.DBUrl)
	fmt.Println("PORT", cfg.Port)
}
