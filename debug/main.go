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
	"github.com/traPtitech/traq-ws-bot/payload"
)

func newPayload(txt string) *payload.MessageCreated {
	return &payload.MessageCreated{
		Base: payload.Base{
			EventTime: time.Now(),
		},
		Message: payload.Message{
			ID: random.UUID().String(),
			User: payload.User{
				ID:          random.UUID().String(),
				Name:        "debug",
				DisplayName: "debug🔧",
				IconID:      random.UUID().String(),
				Bot:         false,
			},
			ChannelID: random.UUID().String(),
			Text:      txt,
			PlainText: txt,
			Embedded:  []payload.EmbeddedInfo{},
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
