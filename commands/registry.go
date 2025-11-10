package commands

import (
	"fmt"
	"log"

	"github.com/bwmarrin/discordgo"
)

var registry = make(map[string]*Command)

func Register(cmd *Command) {
	if cmd == nil {
		log.Println("attempted to register nil command")
		return
	}
	if cmd.Name == "" {
		log.Println("attempted to register command with empty name")
		return
	}
	if _, exists := registry[cmd.Name]; exists {
		log.Printf("%s is already registered, overwriting", cmd.Name)
	}
	registry[cmd.Name] = cmd
	log.Printf("Command '%s' registered", cmd.Name)
}

// Command名からコマンドを取得
func Get(name string) (*Command, bool) {
	cmd, ok := registry[name]
	return cmd, ok
}

// 登録された全てのコマンドをreturn
func GetAll() []*Command {
	cmds := make([]*Command, 0, len(registry))
	for _, cmd := range registry {
		cmds = append(cmds, cmd)
	}
	return cmds
}

// 登録された全てのコマンドを作成
func RegisterAll(session *discordgo.Session, guildID string) error {
	cmds := GetAll()
	if len(cmds) == 0 {
		return fmt.Errorf("no commands registered")
	}

	for _, cmd := range cmds {
		appCmd := cmd.ToApplicationCommand()
		createdCmd, err := session.ApplicationCommandCreate(session.State.User.ID, guildID, appCmd)
		if err != nil {
			log.Printf("Error registering command '%s': %v", cmd.Name, err)
		}
		cmd.AddApplicationCommand(createdCmd)
		log.Printf("Command '%s' registered to Discord (ID: %s)", cmd.Name, createdCmd.ID)
	}

	return nil
}

