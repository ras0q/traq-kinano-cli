package main

import (
	"log"
	"os"

	"github.com/Ras96/traq-kinano-cli/infrastructure"
)

func main() {
	s := infrastructure.NewServer()
	log.Fatal(s.ListenAndServe(":" + os.Getenv("APP_PORT")))
}
