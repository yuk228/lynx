package handler

import (
	"log"

	"lynx/commands"

	"github.com/bwmarrin/discordgo"
)

func OnMessageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
	log.Println(m.Content)
}


func OnInteractionCreate(s *discordgo.Session, i *discordgo.InteractionCreate) {

	commandName := i.ApplicationCommandData().Name
	
	cmd, ok := commands.Get(commandName)
	if !ok {
		log.Printf("Unknown command: %s", commandName)
		return
	}

	cmd.Executor(s, i)
}