package infrastructure

import (
	"context"
	"fmt"
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

func NewServer() (*traqbot.BotServer, error) {
	// Setup ent client
	client, err := ent.Open("mysql", fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&collation=utf8mb4_general_ci",
		config.SQL.User,
		config.SQL.Pass,
		config.SQL.Host,
		config.SQL.Port,
		config.SQL.DBName,
	))
	if err != nil {
		log.Fatalf("Could not open database: %v", err)
	}
	defer client.Close()

	if err := client.Schema.Create(context.Background()); err != nil {
		return nil, err
	}

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
			cmds := injectCmds(context.Background(), client, pl)
			if err := cmds.Execute(args); err != nil {
				traq.MustPostMessage(pl.Message.ChannelID, err.Error())
			}
		}
	})

	traq.MustPostMessage(config.Traq.BotCh, "デプロイ完了やんね！:kinano.rotate:")

	return traqbot.NewBotServer(config.Bot.Verificationtoken, h), nil
}

// メッセージ先頭にメンションを含む場合はargsから除外する
func removeHeadMention(embed traqbot.EmbeddedInfoPayload, args []string) []string {
	if embed.Raw == args[0] && embed.ID == config.Bot.UserID {
		args = args[1:]
	}

	return args
}
