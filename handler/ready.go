package handler

import (
	"log"

	"github.com/bwmarrin/discordgo"
)

func OnReady(s *discordgo.Session, r *discordgo.Ready) {
	log.Println("Bot is up!")
}

