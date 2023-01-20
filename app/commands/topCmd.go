package commands

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/cancelledbit/pizda_bot/app/repository"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
	"strconv"
	"time"
)

type TopCmd struct {
	dbCmd
}

func (c TopCmd) Execute(cmd *tgbotapi.Message) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()
	r := repository.NewMysqlPhrasesRepository(ctx, c.DB)

	top, err := r.GetTop(strconv.FormatInt(cmd.Chat.ID, 10), 5)
	if err != nil {
		log.Println(err)
		return
	}
	text := ""
	for _, t := range top {
		text = text + fmt.Sprintf("%s: %d \n", t.Name, t.Count)
	}
	reply := tgbotapi.NewMessage(cmd.Chat.ID, text)
	reply.ReplyToMessageID = cmd.MessageID
	_, err = c.Bot.Send(reply)
	if err != nil {
		log.Println(err)
	}
}

func getTopCmd(db *sql.DB, bot *tgbotapi.BotAPI) TopCmd {
	return TopCmd{
		dbCmd{
			Name: "top",
			DB:   db,
			Bot:  bot,
		},
	}
}
