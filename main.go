package main

import (
	"io/ioutil"
	"log"
	"os"
	"strings"

	"github.com/Ras96/traq-kinano-cli/cmd"
	"github.com/Ras96/traq-kinano-cli/util/traq"
	traqbot "github.com/traPtitech/traq-bot"
)

var cmds = ls("./cmd")

func main() {
	h := traqbot.EventHandlers{}
	h.SetMessageCreatedHandler(func(pl *traqbot.MessageCreatedPayload) {
		if pl.Message.User.Bot {
			return
		}

		text := pl.Message.PlainText
		log.Println("INFO: Message created", text)

		args := strings.Fields(text)
		if _, ok := cmds[args[0]]; ok {
			if err := cmd.Execute(args); err != nil {
				traq.MustPostMessage(pl.Message.ChannelID, err.Error())
			}
		}
	})

	s := traqbot.NewBotServer(os.Getenv("BOT_VERIFICATION_TOKEN"), h)
	log.Fatal(s.ListenAndServe(":" + os.Getenv("APP_PORT")))
}

func ls(dir string) map[string]struct{} {
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		panic(err)
	}

	cmds := make(map[string]struct{})

	for _, file := range files {
		cmd := strings.Replace(file.Name(), ".go", "", 1)
		if cmd == "root" {
			continue
		}

		cmds[cmd] = struct{}{}
	}

	return cmds
}
