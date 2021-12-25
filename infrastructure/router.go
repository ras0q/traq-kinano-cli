package infrastructure

import (
	"context"
	"fmt"
	"image/png"
	"log"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/Ras96/traq-kinano-cli/cmd"
	"github.com/Ras96/traq-kinano-cli/ent"
	"github.com/Ras96/traq-kinano-cli/util/config"
	"github.com/robfig/cron"

	// mysql driver
	_ "github.com/go-sql-driver/mysql"
	traqbot "github.com/traPtitech/traq-ws-bot"
	"github.com/traPtitech/traq-ws-bot/payload"
)

func SetupCron() {
	c := cron.NewWithLocation(time.FixedZone("Asia/Tokyo", 9*60*60))
	c.AddFunc("0 50 23 * *", func() {
		img, err := generateWordcloud()
		if err != nil {
			panic(fmt.Errorf("Error generating wordcloud: %w", err))
		}

		path, _ := filepath.Abs("./wordcloud.png")
		file, err := os.Create(path)
		if err != nil {
			panic(fmt.Errorf("Error creating wordcloud file: %w", err))
		}
		defer file.Close()

		if err := png.Encode(file, img); err != nil {
			panic(fmt.Errorf("Error encoding wordcloud: %w", err))
		}

		file.Seek(0, os.SEEK_SET)

		cid := config.Traq.BotCh
		fid, err := SendFile(file, cid)
		if err != nil {
			panic(fmt.Errorf("Error sending wordcloud: %w", err))
		}

		NewWriter().
			SetChannelID(cid).
			Write([]byte("https://q.trap.jp/files/" + fid))
	})

	c.Start()
}

func NewWsBot(client *ent.Client) (*traqbot.Bot, error) {
	w := NewWriter()

	b, err := traqbot.NewBot(&traqbot.Options{
		AccessToken:   config.Bot.Accesstoken,
		Origin:        "wss://q.trap.jp",
		AutoReconnect: true,
	})
	if err != nil {
		return nil, err
	}

	b.OnPing(func(pl *payload.Ping) {
		log.Println("ping from bot-console")
	})

	b.OnMessageCreated(func(pl *payload.MessageCreated) {
		if pl.Message.User.Bot {
			return
		}

		text := pl.Message.PlainText
		log.Println("[INFO]Message created: ", text)

		args := strings.Fields(text)
		if embeds := pl.Message.Embedded; len(embeds) > 0 {
			args = removeHeadMention(embeds[0], args)
		}

		if _, ok := cmd.CmdNames[args[0]]; ok {
			cmds := InjectCmds(context.Background(), client, pl, w)
			cmds.Execute(args)
		}
	})

	return b, nil
}

// メッセージ先頭にメンションを含む場合はargsから除外する
func removeHeadMention(embed payload.EmbeddedInfo, args []string) []string {
	if embed.Raw == args[0] && embed.ID == config.Bot.UserID {
		args = args[1:]
	}

	return args
}
