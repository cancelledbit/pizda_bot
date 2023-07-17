package commands

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/cancelledbit/pizda_bot/app/repository"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
	"time"
)

type EUCmd struct {
	dbCmd
}

func (c EUCmd) Execute(cmd *tgbotapi.Message) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()
	r := repository.NewMysqlEUPhrasesRepository(ctx, c.DB)

	euPhrase, err := r.Get()
	if err != nil {
		log.Println(err)
		return
	}
	text := fmt.Sprintf("%s: \n %s \n", "EU wisdom", euPhrase.Text)
	reply := tgbotapi.NewMessage(cmd.Chat.ID, text)
	reply.ReplyToMessageID = cmd.MessageID
	_, err = c.Bot.Send(reply)
	if err != nil {
		log.Println(err)
	}
}

func getEUCmd(db *sql.DB, bot *tgbotapi.BotAPI) EUCmd {
	return EUCmd{
		dbCmd{
			Name: "EU|eu",
			DB:   db,
			Bot:  bot,
		},
	}
}
