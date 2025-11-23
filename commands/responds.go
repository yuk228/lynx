package commands

import (
	"log"

	"github.com/FedorLap2006/disgolf"
	"github.com/bwmarrin/discordgo"
)

func RespondErrorMessage(ctx *disgolf.Ctx, err error) error {
	log.Printf("error %v: ", err)
	embed := &discordgo.MessageEmbed{
		Author:      &discordgo.MessageEmbedAuthor{},
		Title:       "Error",
		Description: "```\n" + err.Error() + "\n```",
		Color:       0xFF0000,
	}
	return ctx.Respond(&discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Embeds: []*discordgo.MessageEmbed{embed},
		},
	})
}
