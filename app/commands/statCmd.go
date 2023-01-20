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

type StatCmd struct {
	dbCmd
}

func (c StatCmd) Execute(cmd *tgbotapi.Message) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()
	r := repository.NewMysqlPhrasesRepository(ctx, c.DB)

	phrases, err := r.GetPhrasesByUserId(cmd.From.UserName)
	if err != nil {
		log.Println(err)
		return
	}
	total := 0
	for _, phrase := range *phrases {
		if phrase.SenderChannelId == strconv.FormatInt(cmd.Chat.ID, 10) {
			total++
		}
	}
	reply := tgbotapi.NewMessage(cmd.Chat.ID, fmt.Sprintf("Всего срабатываний: %d", total))
	reply.ReplyToMessageID = cmd.MessageID
	_, err = c.Bot.Send(reply)
	if err != nil {
		log.Println(err)
	}
}

func getMyStatCmd(db *sql.DB, bot *tgbotapi.BotAPI) StatCmd {
	return StatCmd{
		dbCmd{
			Name: "stat",
			DB:   db,
			Bot:  bot,
		},
	}
}
