package config

import (
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
	BotCh string
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

	Traq.BotCh = mustGetenv("TRAQ_BOT_CHANNEL_ID")
}

func mustGetenv(name string) string {
	if env := os.Getenv(name); env == "" {
		panic("env " + name + " is not set")
	} else {
		return env
	}
}

func mustAtoi(s string) int {
	if i, err := strconv.Atoi(s); err != nil {
		panic("atoi failed: " + err.Error())
	} else {
		return i
	}
}
