package infrastructure

import (
	"context"
	"log"
	"strings"

	"github.com/Ras96/traq-kinano-cli/cmd"
	"github.com/Ras96/traq-kinano-cli/ent"
	"github.com/Ras96/traq-kinano-cli/interfaces/external"
	"github.com/Ras96/traq-kinano-cli/util/config"
	"github.com/gofrs/uuid"

	// mysql driver
	_ "github.com/go-sql-driver/mysql"
	traqbot "github.com/traPtitech/traq-ws-bot"
	"github.com/traPtitech/traq-ws-bot/payload"
)

func NewWsBot(client *ent.Client, q external.TraqAPI) (*traqbot.Bot, error) {
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
			cmds := InjectCmds(context.Background(), client, pl, q)
			cmds.Execute(args)
		}
	})

	if err := NewTraqAPI().PostMessage(
		uuid.FromStringOrNil(config.Traq.BotCh),
		"デプロイ完了!",
		true,
	); err != nil {
		return nil, err
	}

	return b, nil
}

// メッセージ先頭にメンションを含む場合はargsから除外する
func removeHeadMention(embed payload.EmbeddedInfo, args []string) []string {
	if embed.Raw == args[0] && embed.ID == config.Bot.UserID {
		args = args[1:]
	}

	return args
}
