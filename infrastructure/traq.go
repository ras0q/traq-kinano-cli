package infrastructure

import (
	"context"

	"github.com/Ras96/traq-kinano-cli/cmd"
	"github.com/Ras96/traq-kinano-cli/util/config"
	"github.com/antihax/optional"
	"github.com/sapphi-red/go-traq"
)

var (
	client = traq.NewAPIClient(traq.NewConfiguration())
	auth   = context.WithValue(context.Background(), traq.ContextAccessToken, config.Bot.Accesstoken)
)

type writer struct {
	channelID string
	embed     bool // Default: true
}

func NewWriter(accessToken string) cmd.Writer {
	return &writer{
		channelID: "",
		embed:     true,
	}
}

func (w *writer) SetChannelID(channelID string) cmd.Writer {
	w.channelID = channelID

	return w
}

func (w *writer) SetEmbed(embed bool) cmd.Writer {
	w.embed = embed

	return w
}

// Implement io.Writer interface
func (w *writer) Write(p []byte) (int, error) {
	_, _, err := client.MessageApi.PostMessage(
		auth,
		w.channelID,
		&traq.MessageApiPostMessageOpts{
			PostMessageRequest: optional.NewInterface(traq.PostMessageRequest{
				Content: string(p),
				Embed:   w.embed,
			}),
		},
	)
	if err != nil {
		return 0, err
	}

	return len(p), nil
}
