package bot

import (
	"fmt"
	"log"

	"github.com/FedorLap2006/disgolf"
	"github.com/yuk228/lynx/commands"
	"github.com/yuk228/lynx/config"
	"github.com/yuk228/lynx/handler"
)

type Bot struct {
	Instance *disgolf.Bot
}

func New() (*Bot, error) {
	bot, err := disgolf.New(config.GetBotToken())
	if err != nil {
		return nil, fmt.Errorf("failed to create bot: %w", err)
	}

	return &Bot{
		Instance: bot,
	}, nil
}

func (b *Bot) Setup() error {
	b.Instance.Router.Register(commands.PingPongCommand())
	b.Instance.Router.Register(commands.JoinVCCommand())

	b.Instance.AddHandler(handler.OnReady)
	b.Instance.AddHandler(b.Instance.Router.HandleInteraction)
	b.Instance.AddHandler(b.Instance.Router.MakeMessageHandler(&disgolf.MessageHandlerConfig{
		MentionPrefix: true,
	}))

	return nil
}

func (b *Bot) Start() error {
	err := b.Instance.Open()
	if err != nil {
		return fmt.Errorf("open exited with an error: %w", err)
	}
	err = b.Instance.Router.Sync(b.Instance.Session, "", "")
	if err != nil {
		return fmt.Errorf("cannot publish commands: %w", err)
	}

	log.Println("Bot started successfully")
	return nil
}


func (b *Bot) Close() error {
	return b.Instance.Close()
}

