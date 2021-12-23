package infrastructure

import (
	"context"
	"os"
	"regexp"
	"sync"
	"time"

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

func SendFile(file *os.File, channelID string) error {
	_, _, err := client.FileApi.PostFile(
		auth,
		file,
		channelID,
	)
	if err != nil {
		return err
	}

	return nil
}

func getTraqDailyMsgs() ([]string, error) {
	time.FixedZone("Asia/Tokyo", 9*60*60)

	var (
		now    = time.Now()
		after  = optional.NewTime(time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, now.Location()))
		before = optional.NewTime(time.Date(now.Year(), now.Month(), now.Day(), 23, 59, 59, 59, now.Location()))
		limit  = optional.NewInt32(100)
		bot    = optional.NewBool(false)
		hasURL = optional.NewBool(false)
	)

	res, _, err := client.MessageApi.SearchMessages(
		auth,
		&traq.MessageApiSearchMessagesOpts{
			Before: before,
			After:  after,
			Limit:  limit,
			Bot:    bot,
			HasURL: hasURL,
		},
	)
	if err != nil {
		return nil, err
	}

	msgs := make([]string, 0, res.TotalHits)
	for _, msg := range res.Hits {
		if msg.Content != "" {
			msgs = append(msgs, msg.Content)
		}
	}

	num := int(res.TotalHits) / 100 // 並列で回す数
	wg := sync.WaitGroup{}
	wg.Add(num)
	for i := 0; i < num; i++ {
		go func(i int) {
			defer wg.Done()
			res, _, _ := client.MessageApi.SearchMessages(
				auth,
				&traq.MessageApiSearchMessagesOpts{
					Before: before,
					After:  after,
					Limit:  limit,
					Offset: optional.NewInt32(int32(i * 100)),
					Bot:    bot,
					HasURL: hasURL,
				},
			)

			for _, msg := range res.Hits {
				if msg.Content != "" {
					plain := regexp.MustCompile("!\\{.+\\}").ReplaceAllString(msg.Content, "")
					msgs = append(msgs, plain)
				}
			}
		}(i)
	}
	wg.Wait()

	return msgs, nil
}
