package infrastructure

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"net/textproto"
	"os"
	"regexp"
	"sync"
	"time"

	"github.com/Ras96/traq-kinano-cli/interfaces/external"
	"github.com/Ras96/traq-kinano-cli/util/config"
	"github.com/antihax/optional"
	"github.com/gofrs/uuid"
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

func (t *traqAPI) SearchMessages(opts external.SearchMessagesOpts) (int, []string, error) {
	res, _, _ := client.MessageApi.SearchMessages(
		auth,
		&traq.MessageApiSearchMessagesOpts{
			After:  opts.After,
			Before: opts.Before,
			Bot:    opts.Bot,
			Limit:  opts.Limit,
			Offset: opts.Offset,
		},
	)

	msgs := make([]string, len(res.Hits))
	for i, hit := range res.Hits {
		msgs[i] = hit.Content
	}

	return int(res.TotalHits), msgs, nil
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

func (t *traqAPI) PostFile(channelID uuid.UUID, file *os.File) (uuid.UUID, error) {
	// NOTE: go-traqがcontent-typeをapplication/octet-streamにしてしまうので自前でAPIを叩く
	var b bytes.Buffer
	mw := multipart.NewWriter(&b)

	mh := make(textproto.MIMEHeader)
	mh.Set("Content-Type", "image/png")
	mh.Set("Content-Disposition", fmt.Sprintf(`form-data; name="file"; filename="%s"`, file.Name()))

	pw, err := mw.CreatePart(mh)
	if err != nil {
		return uuid.Nil, fmt.Errorf("failed to create part: %w", err)
	}

	if _, err := io.Copy(pw, file); err != nil {
		return uuid.Nil, fmt.Errorf("failed to copy file: %w", err)
	}

	contentType := mw.FormDataContentType()
	mw.Close()

	req, err := http.NewRequest(
		"POST",
		fmt.Sprintf("https://q.trap.jp/api/v3/files?channelId=%s", channelID.String()),
		&b,
	)
	if err != nil {
		return uuid.Nil, fmt.Errorf("Error creating request: %w", err)
	}

	req.Header.Set("Content-Type", contentType)
	req.Header.Set("Authorization", "Bearer "+config.Bot.Accesstoken)

	client := new(http.Client)
	res, err := client.Do(req)
	if err != nil {
		return uuid.Nil, fmt.Errorf("Error sending request: %w", err)
	}
	defer res.Body.Close()

	if res.StatusCode >= 300 {
		b, _ := io.ReadAll(res.Body)

		return uuid.Nil, fmt.Errorf("Error creating file: %s %s", res.Status, string(b))
	}

	var traqFile traq.FileInfo
	if err := json.NewDecoder(res.Body).Decode(&traqFile); err != nil {
		return uuid.Nil, fmt.Errorf("Error decoding response: %w", err)
	}

	return uuid.FromStringOrNil(traqFile.Id), nil
}

func getTraqDailyMsgs() ([]string, error) {
	var (
		now    = time.Now().UTC()
		after  = optional.NewTime(time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, time.UTC))
		before = optional.NewTime(time.Date(now.Year(), now.Month(), now.Day(), 23, 59, 59, 0, time.UTC))
		limit  = optional.NewInt32(100)
		bot    = optional.NewBool(false)
		r      = regexp.MustCompile(`!\{.+\}|https?:\/\/.+(\s|$)`)
		msgs   = make([]string, 0, 5000)
	)

	searchFunc := func(offset int) int {
		res, _, _ := client.MessageApi.SearchMessages(
			auth,
			&traq.MessageApiSearchMessagesOpts{
				Before: before,
				After:  after,
				Limit:  limit,
				Offset: optional.NewInt32(int32(offset * 100)),
				Bot:    bot,
			},
		)

		for _, msg := range res.Hits {
			if msg.Content != "" {
				plain := r.ReplaceAllString(msg.Content, "")
				msgs = append(msgs, plain)
			}
		}

		return int(res.TotalHits)
	}

	// 総メッセージ数を取得するために1かい先にAPIを叩く
	totalHits := searchFunc(0)

	num := totalHits / 100 // 並列で回す数
	wg := sync.WaitGroup{}
	wg.Add(num)
	for i := 0; i < num; i++ {
		go func(i int) {
			defer wg.Done()
			searchFunc(i)
		}(i)
	}
	wg.Wait()

	return msgs, nil
}
