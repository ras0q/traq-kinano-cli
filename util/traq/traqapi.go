package traq

import (
	"context"
	"fmt"

	"github.com/Ras96/traq-kinano-cli/util/config"
	"github.com/antihax/optional"
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
		return nil, fmt.Errorf("failed to post message: %w", err)
	}

	return &message, nil
}

func MustPostMessage(channelID string, content string) {
	if _, err := PostMessage(channelID, content); err != nil {
		panic(err)
	}
}
