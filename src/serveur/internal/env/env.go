package env

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func NewEnv() {
	if err := godotenv.Load("internal/env/.env"); err != nil {
		log.Println("No .env found, used env variable")
	}
}

func GetEnvVariable(key, fallback string) string {
	if val := os.Getenv(key); val != "" {
		return val
	}
	return fallback
}
