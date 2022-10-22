package helper

import (
	"log"

	"github.com/joho/godotenv"
)

func LoadEnv() {
	envErr := godotenv.Load(".env")

	if envErr != nil {
		log.Fatal("Error loading .env file")
	}
}
