package commands

import (
	"database/sql"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type dbCmd struct {
	Name string
	DB   *sql.DB
	Bot  *tgbotapi.BotAPI
}

func (c dbCmd) Match(cmd *tgbotapi.Message) bool {
	return c.Name == cmd.Command()
}
