package commands

import (
	"database/sql"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type CommandInterface interface {
	Match(cmd *tgbotapi.Message) bool
	Execute(cmd *tgbotapi.Message)
}

type CommandHandler struct {
	list []CommandInterface
}

func (h CommandHandler) Handle(cmd *tgbotapi.Message) {
	for _, command := range h.list {
		if command.Match(cmd) {
			command.Execute(cmd)
			return
		}
	}
}

func GetHandler(db *sql.DB, bot *tgbotapi.BotAPI) CommandHandler {
	return CommandHandler{
		list: []CommandInterface{
			getMyStatCmd(db, bot),
			getTopCmd(db, bot),
			getEUCmd(db, bot),
			getHagzorCmd(db, bot),
			getAntonCmd(db, bot),
		},
	}
}
