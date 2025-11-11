package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/yuk228/lynx/bot"
	_ "github.com/yuk228/lynx/config"
)

func main() {
	b, err := bot.New()
	if err != nil {
		log.Fatal(err)
	}

	err = b.Setup()

	if err != nil {
		log.Fatal(err)
	}

	err = b.Start()
	if err != nil {
		log.Fatal(err)
	}
	defer b.Close()

	stchan := make(chan os.Signal, 1)
	signal.Notify(stchan, syscall.SIGTERM, os.Interrupt, syscall.SIGSEGV)

end:
	for {
		select {
		case <-stchan:
			break end
		default:
		}
		time.Sleep(time.Second)
	}
}
