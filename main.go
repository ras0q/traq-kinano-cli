package main

import (
	"fmt"
	"log"

	"github.com/Ras96/traq-kinano-cli/infrastructure"
	"github.com/Ras96/traq-kinano-cli/util/config"
)

func main() {
	entClient, err := infrastructure.NewEntClient()
	if err != nil {
		log.Fatal("Error creating client: ", err)
	}
	defer entClient.Close()

	server := infrastructure.NewServer(entClient)

	log.Fatal(server.ListenAndServe(fmt.Sprintf(":%d", config.App.Port))) //nolint:gocritic
}
