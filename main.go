package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"lynx/handler"

	"github.com/bwmarrin/discordgo"
	"github.com/joho/godotenv"
)

func loadenv() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Failed to load env file %v", err)
	}
	log.Println(".env loaded")
}

func main() {
	loadenv()

	discord, err := discordgo.New("Bot " + os.Getenv("TOKEN"))
	if err != nil {
		log.Fatalf("Failed to login %v", err)
		return
	}
	discord.AddHandler(handler.OnMessageCreate)

	err = discord.Open()
	if err != nil {
		log.Fatalf("Failed to start session %v", err)
		return
	}

	defer discord.Close()
	stopBot := make(chan os.Signal, 1)
	signal.Notify(stopBot, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-stopBot
}