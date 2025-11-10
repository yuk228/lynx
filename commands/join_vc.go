package commands

import "github.com/bwmarrin/discordgo"

func init() {
	Register(JoinVCCommand())
}

// join vc
func JoinVCCommand() *Command {
	return &Command{
		Name:        "join",
		Description: "join vc channel",
		Options:     []*discordgo.ApplicationCommandOption{},
		Executor: func(s *discordgo.Session, i *discordgo.InteractionCreate) {

			vs, err := s.State.VoiceState(i.GuildID, i.Member.User.ID)
			if err != nil {
				s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
					Type: discordgo.InteractionResponseChannelMessageWithSource,
					Data: &discordgo.InteractionResponseData{
						Content: "you are not in vc channel",
					},
				})
				return
			}

			_, err = s.ChannelVoiceJoin(i.GuildID, vs.ChannelID, false, true)
			if err != nil {
				s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
					Type: discordgo.InteractionResponseChannelMessageWithSource,
					Data: &discordgo.InteractionResponseData{
						Content: "failed to join voice channel",
					},
				})
				return
			}

			s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
				Type: discordgo.InteractionResponseChannelMessageWithSource,
				Data: &discordgo.InteractionResponseData{
					Content: "joined voice channel",
				},
			})
		},
	}
}
