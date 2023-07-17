package commands

import (
	"database/sql"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"regexp"
)

type dbCmd struct {
	Name string
	DB   *sql.DB
	Bot  *tgbotapi.BotAPI
}

func (c dbCmd) Match(cmd *tgbotapi.Message) bool {
	r, e := regexp.Compile(c.Name)
	if e != nil {
		return false
	}
	return r.MatchString(cmd.Command())
}
