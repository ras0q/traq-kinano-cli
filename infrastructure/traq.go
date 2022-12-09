package infrastructure

import (
	"context"
	"fmt"

	"github.com/antihax/optional"
	"github.com/gofrs/uuid"
	"github.com/ras0q/traq-kinano-cli/interfaces/external"
	"github.com/ras0q/traq-kinano-cli/util/config"
	"github.com/sapphi-red/go-traq"
)

var (
	client = traq.NewAPIClient(traq.NewConfiguration())
	auth   = context.WithValue(context.Background(), traq.ContextAccessToken, config.Bot.Accesstoken)
)

type traqAPI struct {
	client *traq.APIClient
	auth   context.Context
}

func NewTraqAPI() external.TraqAPI {
	return &traqAPI{
		client: client,
		auth:   auth,
	}
}

func (t *traqAPI) PostMessage(channelID uuid.UUID, content string, embed bool) error {
	_, _, err := t.client.MessageApi.PostMessage(
		t.auth,
		channelID.String(),
		&traq.MessageApiPostMessageOpts{
			PostMessageRequest: optional.NewInterface(traq.PostMessageRequest{
				Content: content,
				Embed:   embed,
			}),
		},
	)
	if err != nil {
		return fmt.Errorf("failed to post message: %w", err)
	}

	return nil
}
