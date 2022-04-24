package util

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

// GoDotEnvVariable returns the value of the environment variable
func GoDotEnvVariable(key string) string {
	err := godotenv.Load("environment.env")
	if err != nil {
		log.Fatal("error loading environment variables")
	}
	return os.Getenv(key)
}
