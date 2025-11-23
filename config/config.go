package config

import (
	"fmt"
	"log"
	"os"

	dotenv "github.com/joho/godotenv"
)

var (
	BotToken    string
	VoiceBoxURL string
)

func init() {
	err := dotenv.Load()
	if err != nil {
		log.Fatal(fmt.Errorf("cannot load .env: %w", err))
	}

	BotToken = os.Getenv("TOKEN")
	VoiceBoxURL = os.Getenv("VOICEBOX_URL")
}
