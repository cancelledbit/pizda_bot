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

type MyStatCmd struct {
	name string
	db   *sql.DB
	bot  *tgbotapi.BotAPI
}

func (c MyStatCmd) Match(cmd *tgbotapi.Message) bool {
	return c.name == cmd.Command()
}

func (c MyStatCmd) Execute(cmd *tgbotapi.Message) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r := repository.NewMysqlPhrasesRepository(ctx, c.db)

	phrases, err := r.GetPhrasesByUserId(cmd.From.ID)
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
	_, err = c.bot.Send(reply)
	if err != nil {
		log.Println(err)
	}
}

func getMyStatCmd(db *sql.DB, bot *tgbotapi.BotAPI) MyStatCmd {
	return MyStatCmd{
		name: "stat",
		db:   db,
		bot:  bot,
	}
}
