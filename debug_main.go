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
	"github.com/Ras96/traq-kinano-cli/util/config"
	"github.com/Ras96/traq-kinano-cli/util/random"
	"github.com/traPtitech/traq-ws-bot/payload"
)

func init() {
	config.Traq.HomeCh = config.Traq.BotDMCh
	config.Traq.BotCh = config.Traq.BotDMCh
}

func main() {
	flag.Parse()
	args := flag.Args()
	pl := newPayload(strings.Join(args, " "))

	q := infrastructure.NewTraqAPI()

	if err := infrastructure.PostWordcloudToTraq(q); err != nil {
		panic(err)
	}

	entClient, err := infrastructure.NewEntClient()
	if err == nil {
		defer entClient.Close()
	} else {
		log.Println("[WARN]failed to create ent client:", err.Error())
	}

	cmds := infrastructure.InjectCmds(context.Background(), entClient, pl, q)
	cmds.Execute(args)
}

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
			ChannelID: config.Traq.BotDMCh,
			Text:      txt,
			PlainText: txt,
			Embedded:  []payload.EmbeddedInfo{},
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
	}
}
