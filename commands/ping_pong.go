package commands

import (
	"fmt"

	"github.com/FedorLap2006/disgolf"
	"github.com/bwmarrin/discordgo"
)


func PingPongCommand() *disgolf.Command {
	return &disgolf.Command{
		Name:        "ping_pong",
		Description: "Ping it!",
		Type:        discordgo.ChatApplicationCommand,
		Handler: disgolf.HandlerFunc(func(ctx *disgolf.Ctx) {
			_ = ctx.Respond(&discordgo.InteractionResponse{
				Type: discordgo.InteractionResponseChannelMessageWithSource,
				Data: &discordgo.InteractionResponseData{
					Content: "Hi, I'm a bot built on Disgolf library.",
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

