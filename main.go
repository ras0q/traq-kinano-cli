package main

import (
	"github.com/Ras96/traq-kinano-cli/infrastructure"
)

func main() {
	infrastructure.SetupCron()

	entClient, err := infrastructure.NewEntClient()
	if err != nil {
		panic("Error creating client: " + err.Error())
	}
	defer entClient.Close()

	bot, err := infrastructure.NewWsBot(entClient)
	if err != nil {
		panic("Error creating server: " + err.Error())
	}

	panic(bot.Start())
}
