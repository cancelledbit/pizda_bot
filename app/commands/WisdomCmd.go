package commands

import (
	"context"
	"errors"
	"fmt"
	"github.com/cancelledbit/pizda_bot/app/repository"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
	"time"
)

type WisdomCmd struct {
	dbCmd
	AuthorId string
	UserName string
}

func (c WisdomCmd) Execute(cmd *tgbotapi.Message) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()
	r := repository.UserWisdomRepository(ctx, c.DB)

	euPhrase, err := r.Get(c.AuthorId)
	if err != nil {
		log.Println(err)
		return
	}
	text := fmt.Sprintf("%s %s: \n%s \n", c.UserName, "wisdom", euPhrase.Text)
	reply := tgbotapi.NewMessage(cmd.Chat.ID, text)
	reply.ReplyToMessageID = cmd.MessageID
	_, err = c.Bot.Send(reply)
	if err != nil {
		log.Println(err)
	}
}

func GetWisdomUserPattern(userId int64) (string, error) {
	config := map[int64]string{
		5167519420: "/(\\sпис[яею])|(\\sпоп[аук])|(износ)|(\\sвон[яю])|(\\sнож[ек])|(\\sслад)|(\\sхагз)|(\\sдево[нч])|(черк)|(лиза)|(\\sкончи)|(\\sжоп)/u",
		5865654725: "/\\s((sир[аоу])|(секс)|(муж)|(п[еи]зд)|(сос))/iu",
	}
	pattern, ok := config[userId]
	if ok {
		return pattern, nil
	}
	return "", errors.New("pattern not found")
}
