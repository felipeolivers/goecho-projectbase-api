package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func loadEnvironments() {
	if len(os.Getenv("APP_NAME")) < 1 {
		err := godotenv.Load("dev.env")
		if err != nil {
			log.Fatal("Error loading .env file")
		}
	}
}
