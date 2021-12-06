package traq

import (
	"context"
	"errors"
	"fmt"
	"log"

	"github.com/Ras96/traq-kinano-cli/util/config"
	"github.com/antihax/optional"
	traqapi "github.com/sapphi-red/go-traq"
)

var (
	client = traqapi.NewAPIClient(traqapi.NewConfiguration())
	auth   = context.WithValue(context.Background(), traqapi.ContextAccessToken, config.Bot.Accesstoken)

	BotCh = config.Traq.BotCh

	stampMap map[string]string
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
		if errors.Is(err, errors.New("404 Not Found")) {
			log.Println(err)
		}

		log.Println(content)
	}
}

func AddStamp(messageID string, stampName string) error {
	sid, ok := stampMap[stampName]
	if !ok {
		return errors.New("stamp not found")
	}

	_, err := client.StampApi.AddMessageStamp(
		auth,
		messageID,
		sid,
		&traqapi.StampApiAddMessageStampOpts{},
	)
	if err != nil {
		return fmt.Errorf("failed to add stamp to message: %w", err)
	}

	return nil
}

func MustAddStamp(messageID string, stampName string) {
	if err := AddStamp(messageID, stampName); err != nil {
		log.Println(err)
	}
}

func init() {
	stamps, _, err := client.StampApi.GetStamps(
		auth,
		&traqapi.StampApiGetStampsOpts{
			IncludeUnicode: optional.NewBool(true),
		},
	)
	if err != nil {
		log.Fatal(err)
	}

	stampMap = make(map[string]string)
	for _, stamp := range stamps {
		stampMap[stamp.Name] = stamp.Id
	}
}
