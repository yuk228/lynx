package commands

import (
	"fmt"

	"github.com/FedorLap2006/disgolf"
	"github.com/bwmarrin/discordgo"

	"github.com/yuk228/lynx/handler"
)

func PingPongCommand() *disgolf.Command {
	return &disgolf.Command{
		Name:        "test",
		Description: "test",
		Type:        discordgo.ChatApplicationCommand,
		Options: []*discordgo.ApplicationCommandOption{
			{
				Type:        discordgo.ApplicationCommandOptionString,
				Name:        "message",
				Description: "msg",
				Required:    true,
			},
			{
				Type:        discordgo.ApplicationCommandOptionString,
				Name:        "speaker",
				Description: "speaker number",
				Required:    true,
			},
		},
		Handler: disgolf.HandlerFunc(func(ctx *disgolf.Ctx) {
			options := ctx.Interaction.ApplicationCommandData().Options
			message := "test"
			speaker := "1"

			for _, option := range options {
				switch option.Name {
				case "message":
					message = option.StringValue()
				case "speaker":
					speaker = option.StringValue()
				}
			}
			responseMessage := fmt.Sprintf("speaker, message : %s%s", speaker, message)

			bytes, err := handler.GetBinary(message, speaker)
			if err != nil {
				_ = RespondErrorMessage(ctx, err)
				return
			}
			path, err := handler.ToWav(bytes, "1111")
			fmt.Println("bianry size, file_path", len(bytes), path)

			if err != nil {
				_ = RespondErrorMessage(ctx, err)
				return
			}

			_ = ctx.Respond(&discordgo.InteractionResponse{
				Type: discordgo.InteractionResponseChannelMessageWithSource,
				Data: &discordgo.InteractionResponseData{
					Content: responseMessage,
				},
			})
		}),

		MessageHandler: disgolf.MessageHandlerFunc(func(ctx *disgolf.MessageCtx) {
			_, _ = ctx.Reply("Hi, I'm a bot built on Disgolf library", true)
		}),

		Middlewares: []disgolf.Handler{
			disgolf.HandlerFunc(func(ctx *disgolf.Ctx) {
				fmt.Println("Middleware worked!")
				ctx.Next()
			}),
		},

		MessageMiddlewares: []disgolf.MessageHandler{
			disgolf.MessageHandlerFunc(func(ctx *disgolf.MessageCtx) {
				fmt.Println("Message middleware worked!", ctx.Arguments)
				ctx.Next()
			}),
		},
	}
}
