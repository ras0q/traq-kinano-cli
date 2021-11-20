package main

import (
	"log"
	"os"

	"github.com/Ras96/traq-kinano-cli/infrastructure"
	"github.com/Ras96/traq-kinano-cli/util/dir"
)

func main() {
	cmdNames := dir.Ls("./cmd")
	s := infrastructure.NewServer(cmdNames)

	log.Fatal(s.ListenAndServe(":" + os.Getenv("APP_PORT")))
}
