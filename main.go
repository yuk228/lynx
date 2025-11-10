package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/yuk228/lynx/commands"
	"github.com/yuk228/lynx/handler"

	"github.com/bwmarrin/discordgo"
	"github.com/joho/godotenv"
)

func loadEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Failed to load env file %v", err)
	}
	log.Println(".env loaded")
}

func main() {
	loadEnv()

	discord, err := discordgo.New("Bot " + os.Getenv("TOKEN"))
	if err != nil {
		log.Fatalf("Failed to login %v", err)
	}
	discord.Identify.Intents = discordgo.MakeIntent(discordgo.IntentsAll)
	discord.AddHandler(handler.OnMessageCreate)
	discord.AddHandler(handler.OnInteractionCreate)

	err = discord.Open()
	if err != nil {
		log.Fatalf("Failed to start session %v", err)
	}

	err = commands.RegisterAll(discord, "")
	if err != nil {
		log.Printf("Warning: Some commands may not have been registered: %v", err)
	}

	log.Println("Bot is running")
	defer discord.Close()

	stopBot := make(chan os.Signal, 1)
	signal.Notify(stopBot, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-stopBot
}
