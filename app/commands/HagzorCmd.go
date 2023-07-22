package commands

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"time"

	"github.com/cancelledbit/pizda_bot/app/repository"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type HagzorCmd struct {
	dbCmd
}

func (c HagzorCmd) Execute(cmd *tgbotapi.Message) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()
	r := repository.NewMysqlEUPhrasesRepository(ctx, c.DB)

	hagzorPhrase, err := r.Get()
	if err != nil {
		log.Println(err)
		return
	}

	text := fmt.Sprintf("%s: \n %s \n", "hagz0r wisdom", hagzorPhrase.Text)
	reply := tgbotapi.NewMessage(cmd.Chat.ID, text)
	reply.ReplyToMessageID = cmd.MessageID
	_, err = c.Bot.Send(reply)
	if err != nil {
		log.Println(err)
	}
}

func getHagzorCmd(db *sql.DB, bot *tgbotapi.BotAPI) HagzorCmd {
	return HagzorCmd{
		dbCmd{
			Name: "hagz0r|hz|hagzor",
			DB:   db,
			Bot:  bot,
		},
	}
}
