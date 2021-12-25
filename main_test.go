package main_test

import (
	"context"
	"flag"
	"log"
	"strings"
	"testing"
	"time"

	"github.com/Ras96/traq-kinano-cli/cmd"
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

type writer struct{}

func NewWriter() cmd.Writer {
	return &writer{}
}

func (w *writer) Write(p []byte) (int, error) {
	log.Println(string(p))
	return len(p), nil
}

func (w *writer) SetChannelID(channelID string) cmd.Writer {
	return w
}

func (w *writer) SetEmbed(embed bool) cmd.Writer {
	return w
}

func Test_main(t *testing.T) {
	flag.Parse()
	args := flag.Args()
	pl := newPayload(strings.Join(args, " "))

	if err := infrastructure.PostWordcloutToTraq(); err != nil {
		t.Fatal(err)
	}

	entClient, err := infrastructure.NewEntClient()
	if err == nil {
		defer entClient.Close()
	} else {
		log.Println("[WARN]failed to create ent client:", err.Error())
	}

	w := NewWriter()

	cmds := infrastructure.InjectCmds(context.Background(), entClient, pl, w)
	cmds.Execute(args)
}
