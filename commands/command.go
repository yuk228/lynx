package commands

import "github.com/bwmarrin/discordgo"

type Command struct {
	Name        string
	Description string
	Options     []*discordgo.ApplicationCommandOption
	AppCommand  *discordgo.ApplicationCommand
	Executor    func(*discordgo.Session, *discordgo.InteractionCreate)
}

func (c *Command) ToApplicationCommand() *discordgo.ApplicationCommand {
	return &discordgo.ApplicationCommand{
		Name:        c.Name,
		Description: c.Description,
		Options:     c.Options,
	}
}

func (c *Command) AddApplicationCommand(appCmd *discordgo.ApplicationCommand) {
	c.AppCommand = appCmd
}
