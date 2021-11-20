package main

import (
	"log"
	"os"
	"strings"

	"github.com/Ras96/traq-kinano-cli/cmd"
	traqbot "github.com/traPtitech/traq-bot"
)

func main() {
	h := traqbot.EventHandlers{}
	h.SetMessageCreatedHandler(func(payload *traqbot.MessageCreatedPayload) {
		args := strings.Fields(payload.Message.PlainText)
		cmd.Execute(args)
	})

	s := traqbot.NewBotServer(os.Getenv("BOT_VERIFICATION_TOKEN"), h)
	log.Fatal(s.ListenAndServe(":" + os.Getenv("APP_PORT")))
}
