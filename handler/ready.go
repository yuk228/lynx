package handler

import (
	"log"

	"github.com/bwmarrin/discordgo"
)

func OnReady(s *discordgo.Session, r *discordgo.Ready) {
	user := r.User
	log.Printf("Bot is up. Logged in as: %s#%s", user.Username, user.Discriminator)
	log.Printf("BotID: %s", user.ID)
	log.Printf("%d guilds connected ", len(r.Guilds))
}

