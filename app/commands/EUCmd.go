package commands

import (
	"database/sql"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type EUWisdomCmd struct {
	WisdomCmd
}

func getEUCmd(db *sql.DB, bot *tgbotapi.BotAPI) EUWisdomCmd {
	return EUWisdomCmd{
		WisdomCmd{
			dbCmd: dbCmd{
				Name: "EU|eu",
				DB:   db,
				Bot:  bot,
			},
			AuthorId: "5167519420",
			UserName: "EU",
		},
	}
}
