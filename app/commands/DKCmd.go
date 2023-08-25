package commands

import (
	"database/sql"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type DKCmd struct {
	WisdomCmd
}

func getDKCmd(db *sql.DB, bot *tgbotapi.BotAPI) DKCmd {
	return DKCmd{
		WisdomCmd{
			dbCmd: dbCmd{
				Name: "DK|dk|alien",
				DB:   db,
				Bot:  bot,
			},
			AuthorId: "418587687",
			UserName: "Double Keeper",
		},
	}
}
