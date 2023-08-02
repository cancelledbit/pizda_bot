package commands

import (
	"context"
	"errors"
	"fmt"
	"github.com/cancelledbit/pizda_bot/app/repository"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
	"math"
	"os"
	"strconv"
	"time"
)

type WisdomCmd struct {
	dbCmd
	AuthorId string
	UserName string
}

func (c WisdomCmd) Execute(cmd *tgbotapi.Message) {
	count, err := strconv.Atoi(cmd.CommandArguments())
	if err != nil || count == 0 {
		count = 1
	}

	maxPhrases, err := strconv.Atoi(os.Getenv("MAX_PHRASES_GEN"))
	if err != nil {
		maxPhrases = 3
	}

	count = int(math.Min(float64(count), float64(maxPhrases)))

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()
	r := repository.UserWisdomRepository(ctx, c.DB)

	phrases, err := r.Get(c.AuthorId, count)
	if err != nil {
		log.Println(err)
		return
	}
	text := fmt.Sprintf("%s когда то сказал: \n", c.UserName)
	for _, p := range phrases {
		text += fmt.Sprintf("%s \n", p.Text)
	}

	reply := tgbotapi.NewMessage(cmd.Chat.ID, text)
	reply.ReplyToMessageID = cmd.MessageID
	_, err = c.Bot.Send(reply)
	if err != nil {
		log.Println(err)
	}
}

func GetWisdomUserPattern(userId int64) (string, error) {
	config := map[int64]string{
		5167519420: "/(\\sпис[яею])|(\\sпоп[аук])|(износ)|(накур)|(дудк)|(\\sвон[яю])|(\\sнож[ек])|(\\sслад)|(\\sхагз)|(\\sдево[нч])|(черк)|(лиза)|(\\sкончи)|(\\sжоп)/iu",
		5865654725: "/\\s(ир[аоу])|(секс)|(муж)|(п[еи]зд)|(сос)|(ваг)|(сис)|(пис)|(ху)/iu",
		409842850:  "/\\s(ебл)|(секс)|(муж)|(нарк)|(мото)|(дудк)|(трава)|(баб)|(гомо)|(геи)/iu",
		5655245858: "/\\s(монгол)|(коре)|(казах)|(п[еи]тон)|(сперм)|(секс)|(ганд)|(джанг)|(апи)|(ма[мт])|(ксан)|(клуб)|(бар)|(рейв)/iu",
	}
	pattern, ok := config[userId]
	if ok {
		return pattern, nil
	}
	return "", errors.New("pattern not found")
}
