package main

import (
	"github.com/ras0q/traq-kinano-cli/infrastructure"
)

func main() {
	traqAPI := infrastructure.NewTraqAPI()

	entClient, err := infrastructure.NewEntClient()
	if err != nil {
		panic("Error creating client: " + err.Error())
	}
	defer entClient.Close()

	bot, err := infrastructure.NewWsBot(entClient, traqAPI)
	if err != nil {
		panic("Error creating server: " + err.Error())
	}

	panic(bot.Start())
}
