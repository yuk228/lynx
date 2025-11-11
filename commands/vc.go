package commands

import (
	"github.com/FedorLap2006/disgolf"
	"github.com/bwmarrin/discordgo"
)

type VcData struct {
	connection *discordgo.VoiceConnection
}

func JoinVCCommand() *disgolf.Command {
	return &disgolf.Command{
		Name:        "join",
		Description: "join vc channel",
		Type:        discordgo.ChatApplicationCommand,
		Handler: disgolf.HandlerFunc(func(ctx *disgolf.Ctx) {
			s := ctx.Session
			guildID := ctx.Interaction.GuildID
			userID := ctx.Interaction.Member.User.ID

			vs, err := s.State.VoiceState(guildID, userID)
			if err != nil {
				_ = ctx.Respond(&discordgo.InteractionResponse{
					Type: discordgo.InteractionResponseChannelMessageWithSource,
					Data: &discordgo.InteractionResponseData{
						Content: "you are not in vc channel",
					},
				})
				return
			}

			_, err = s.ChannelVoiceJoin(guildID, vs.ChannelID, false, true)
			if err != nil {
				_ = ctx.Respond(&discordgo.InteractionResponse{
					Type: discordgo.InteractionResponseChannelMessageWithSource,
					Data: &discordgo.InteractionResponseData{
						Content: "failed to join voice channel",
					},
				})
				return
			}

			_ = ctx.Respond(&discordgo.InteractionResponse{
				Type: discordgo.InteractionResponseChannelMessageWithSource,
				Data: &discordgo.InteractionResponseData{
					Content: "joined voice channel",
				},
			})
		}),
	}
}

