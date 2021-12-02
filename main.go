package main

import (
	"fmt"
	"log"

	"github.com/Ras96/traq-kinano-cli/infrastructure"
	"github.com/Ras96/traq-kinano-cli/util/config"
)

func main() {
	s, err := infrastructure.NewServer()
	if err != nil {
		log.Fatal(err)
	}

	log.Fatal(s.ListenAndServe(fmt.Sprintf(":%d", config.App.Port)))
}
