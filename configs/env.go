package configs

import (
	"os"
	"log"
	"github.com/joho/godotenv"
)

func EnvMongoURI() (string, string) {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	return os.Getenv("MONGOURI"), os.Getenv("VIURI")
}