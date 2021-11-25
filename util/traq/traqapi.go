package traq

import (
	"context"

	"github.com/Ras96/traq-kinano-cli/util/config"
	"github.com/antihax/optional"
	"github.com/pkg/errors"
	traqapi "github.com/sapphi-red/go-traq"
)

var (
	client = traqapi.NewAPIClient(traqapi.NewConfiguration())
	auth   = context.WithValue(context.Background(), traqapi.ContextAccessToken, config.Bot.Accesstoken)

	BotCh = config.Traq.BotCh
)

func PostMessage(channelID string, content string) (*traqapi.Message, error) {
	message, _, err := client.MessageApi.PostMessage(
		auth,
		channelID,
		&traqapi.MessageApiPostMessageOpts{
			PostMessageRequest: optional.NewInterface(traqapi.PostMessageRequest{
				Content: content,
				Embed:   true,
			}),
		},
	)
	if err != nil {
		return nil, errors.Wrap(err, "failed to post message")
	}

	return &message, nil
}

func MustPostMessage(channelID string, content string) {
	if _, err := PostMessage(channelID, content); err != nil {
		panic(err)
	}
}
