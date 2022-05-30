package config

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

func Config(key string) string {
	err := godotenv.Load()

	if err != nil {
		log.Fatal("can not load .env file.")
	}

	return os.Getenv(key)
}
