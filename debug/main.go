//go:build debug
// +build debug

package main

import (
	"context"
	"flag"
	"log"
	"strings"
	"time"

	"github.com/Ras96/traq-kinano-cli/infrastructure"
	"github.com/Ras96/traq-kinano-cli/util/random"
	traqbot "github.com/traPtitech/traq-bot"
)

func newPayload(txt string) *traqbot.MessageCreatedPayload {
	return &traqbot.MessageCreatedPayload{
		BasePayload: traqbot.BasePayload{
			EventTime: time.Now(),
		},
		Message: traqbot.MessagePayload{
			ID: random.UUID().String(),
			User: traqbot.UserPayload{
				ID:          random.UUID().String(),
				Name:        "debug",
				DisplayName: "debug🔧",
				IconID:      random.UUID().String(),
				Bot:         false,
			},
			ChannelID: random.UUID().String(),
			Text:      txt,
			PlainText: txt,
			Embedded:  []traqbot.EmbeddedInfoPayload{},
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
	}
}

func main() {
	flag.Parse()
	args := flag.Args()
	pl := newPayload(strings.Join(args, " "))

	entClient, err := infrastructure.NewEntClient()
	if err == nil {
		defer entClient.Close()
	} else {
		log.Printf("[WARN]failed to create ent client: %v\n", err)
	}

	cmds := infrastructure.InjectCmds(context.Background(), entClient, pl)
	if err := cmds.Execute(args); err != nil {
		log.Println(err.Error())
	}
}
