package infrastructure

import (
	"context"
	"log"
	"strings"

	"github.com/Ras96/traq-kinano-cli/cmd"
	"github.com/Ras96/traq-kinano-cli/ent"
	"github.com/Ras96/traq-kinano-cli/util/config"
	"github.com/Ras96/traq-kinano-cli/util/traq"

	// mysql driver
	_ "github.com/go-sql-driver/mysql"
	traqbot "github.com/traPtitech/traq-bot"
)

type Handlers struct{}

func NewServer(client *ent.Client) *traqbot.BotServer {
	// Setup traQ EventHandlers
	h := traqbot.EventHandlers{}
	h.SetMessageCreatedHandler(func(pl *traqbot.MessageCreatedPayload) {
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
			cmds := InjectCmds(context.Background(), client, pl)
			if err := cmds.Execute(args); err != nil {
				traq.MustPostMessage(pl.Message.ChannelID, err.Error())
			}
		}
	})

	traq.MustPostMessage(config.Traq.BotCh, "デプロイ完了やんね！:kinano.rotate:")

	return traqbot.NewBotServer(config.Bot.Verificationtoken, h)
}

// メッセージ先頭にメンションを含む場合はargsから除外する
func removeHeadMention(embed traqbot.EmbeddedInfoPayload, args []string) []string {
	if embed.Raw == args[0] && embed.ID == config.Bot.UserID {
		args = args[1:]
	}

	return args
}
