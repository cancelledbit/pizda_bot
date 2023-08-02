package commands

import (
	"database/sql"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type AntonCmd struct {
	WisdomCmd
}

func getAntonCmd(db *sql.DB, bot *tgbotapi.BotAPI) AntonCmd {
	return AntonCmd{
		WisdomCmd{
			dbCmd: dbCmd{
				Name: "anton|python|pt",
				DB:   db,
				Bot:  bot,
			},
			AuthorId: "5655245858",
			UserName: "Anton",
		},
	}
}
