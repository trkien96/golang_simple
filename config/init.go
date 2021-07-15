package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func Initial() map[string]string {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	return map[string]string{
		"APP_URL":  os.Getenv("APP_URL"),
		"APP_PORT": os.Getenv("APP_PORT"),
	}
}
