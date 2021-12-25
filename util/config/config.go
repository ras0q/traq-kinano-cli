package config

import (
	"log"
	"os"
	"strconv"
)

type appCfg struct {
	Port int
}

type botCfg struct {
	ID                string
	UserID            string
	Accesstoken       string
	Verificationtoken string
}

type sqlCfg struct {
	User   string
	Pass   string
	Host   string
	Port   int
	DBName string
}

type traqCfg struct {
	HomeCh string
	BotCh  string
}

var (
	App  appCfg
	Bot  botCfg
	SQL  sqlCfg
	Traq traqCfg
)

func init() {
	App.Port = mustAtoi(mustGetenv("APP_PORT"))

	Bot.ID = mustGetenv("BOT_ID")
	Bot.UserID = mustGetenv("BOT_USER_ID")
	Bot.Accesstoken = mustGetenv("BOT_ACCESS_TOKEN")
	Bot.Verificationtoken = mustGetenv("BOT_VERIFICATION_TOKEN")

	SQL.User = mustGetenv("MARIADB_USERNAME")
	SQL.Pass = mustGetenv("MARIADB_PASSWORD")
	SQL.Host = mustGetenv("MARIADB_HOSTNAME")
	SQL.Port = mustAtoi(mustGetenv("MARIADB_PORT"))
	SQL.DBName = mustGetenv("MARIADB_DATABASE")

	Traq.HomeCh = mustGetenv("TRAQ_HOME_CHANNEL_ID")
	Traq.BotCh = mustGetenv("TRAQ_BOT_CHANNEL_ID")
}

func mustGetenv(name string) string {
	env := os.Getenv(name)
	if env == "" {
		log.Println("env " + name + " is not set")
	}

	return env
}

func mustAtoi(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		log.Println("atoi failed: " + err.Error())
	}

	return i
}
