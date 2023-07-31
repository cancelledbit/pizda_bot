package commands

import (
	"database/sql"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type HagzorCmd struct {
	WisdomCmd
}

func getHagzorCmd(db *sql.DB, bot *tgbotapi.BotAPI) EUWisdomCmd {
	return EUWisdomCmd{
		WisdomCmd{
			dbCmd: dbCmd{
				Name: "hz",
				DB:   db,
				Bot:  bot,
			},
			AuthorId: "5865654725",
			UserName: "hagz0r",
		},
	}
}
