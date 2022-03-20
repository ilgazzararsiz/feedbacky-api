package appconfig

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

func EnvMongoURI() string {
	err := godotenv.Load()
	if err != nil {
		log.Println("Error loading .env file")
	}

	return os.Getenv("MONGOURI")
}
