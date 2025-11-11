package config

import (
	"fmt"
	"log"
	"os"

	dotenv "github.com/joho/godotenv"
)

func init() {
	err := dotenv.Load()
	if err != nil {
		log.Fatal(fmt.Errorf("cannot load .env: %w", err))
	}
}

func GetBotToken() string {
	return os.Getenv("TOKEN")
}
