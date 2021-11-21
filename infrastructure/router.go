package infrastructure

import (
	"log"
	"os"
	"strings"

	"github.com/Ras96/traq-kinano-cli/util/traq"
	traqbot "github.com/traPtitech/traq-bot"
)

type Handlers struct{}

func NewServer(cmdNames map[string]struct{}) *traqbot.BotServer {
	h := traqbot.EventHandlers{}
	h.SetMessageCreatedHandler(func(pl *traqbot.MessageCreatedPayload) {
		if pl.Message.User.Bot {
			return
		}

		text := pl.Message.PlainText
		log.Println("[INFO]Message created: ", text)

		args := strings.Fields(text)
		args = removeHeadMention(pl.Message.Embedded[0], args)

		if _, ok := cmdNames[args[0]]; ok {
			cmds := injectCmds(pl)
			if err := cmds.Execute(args); err != nil {
				traq.MustPostMessage(pl.Message.ChannelID, err.Error())
			}
		}
	})

	traq.MustPostMessage(traq.GpsTimesRasBot, "デプロイ完了やんね！:kinano.rotate:")

	return traqbot.NewBotServer(os.Getenv("BOT_VERIFICATION_TOKEN"), h)
}

// メッセージ先頭にメンションを含む場合はargsから除外する
func removeHeadMention(embed traqbot.EmbeddedInfoPayload, args []string) []string {
	if embed.Raw == args[0] && embed.ID == os.Getenv("BOT_USER_ID") {
		args = args[1:]
	}

	return args
}
