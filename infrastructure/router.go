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
		log.Println("INFO: Message created", text)

		args := strings.Fields(text)
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
