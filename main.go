package main

import (
	"log"

	"github.com/Ras96/traq-kinano-cli/infrastructure"
)

func main() {
	entClient, err := infrastructure.NewEntClient()
	if err != nil {
		log.Fatal("Error creating client: ", err)
	}
	defer entClient.Close()

	bot, err := infrastructure.NewWsBot(entClient)
	if err != nil {
		log.Fatal("Error creating server: ", err)
	}

	log.Fatal(bot.Start())
}
