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

type OpinionCmd struct {
	dbCmd
	Members map[string][2]string
}

func (c OpinionCmd) Execute(cmd *tgbotapi.Message) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()
	r := repository.UserWisdomRepository(ctx, c.DB)
	if cmd.ReplyToMessage == nil {
		reply := tgbotapi.NewMessage(cmd.Chat.ID, "Эта команда должна быть ответом на чье либо сообщение")
		reply.ReplyToMessageID = cmd.MessageID
		c.Bot.Send(reply)
	}

	text := "Мнение авторитетов 2ch/pr по данному вороосу:\n"

	for id, data := range c.Members {
		name := data[0]
		customWhere := data[1]
		phrases, err := r.Get(id, 1, customWhere)
		if err != nil {
			log.Println(err)
			return
		}
		if len(phrases) == 0 {
			continue
		}
		text += fmt.Sprintf("\t%s: %s\n", name, phrases[0].Text)
	}
	reply := tgbotapi.NewMessage(cmd.Chat.ID, text)
	reply.ReplyToMessageID = cmd.ReplyToMessage.MessageID
	_, err := c.Bot.Send(reply)
	if err != nil {
		log.Println(err)
	}
}

func getOpinionCmd(db *sql.DB, bot *tgbotapi.BotAPI) OpinionCmd {
	return OpinionCmd{
		dbCmd: dbCmd{
			Name: "opinion",
			DB:   db,
			Bot:  bot,
		},
		Members: map[string][2]string{
			"5167519420": {"EU", ""},
			"5865654725": {"hagz0r", ""},
			"5655245858": {"Anton", ""},
			"418587687":  {"Double Keeper", " AND LENGTH(text) < 80"},
		},
	}
}
