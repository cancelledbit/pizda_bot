package stat

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/cancelledbit/pizda_bot/app/repository"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
	"strconv"
)

type Handler struct {
	db *sql.DB
}

func (h *Handler) PushStat(message *tgbotapi.Message, reply string) {
	channelsRepo := repository.NewMysqlChannelRepository(context.Background(), h.db)
	channel := &repository.Channel{
		ChannelId:   strconv.FormatInt(message.Chat.ID, 10),
		ChannelName: message.Chat.Title,
		Type:        message.Chat.Type,
	}
	_, err := channelsRepo.Create(channel)
	if err != nil {
		log.Println(err)
		return
	}
	phrasesRepo := repository.NewMysqlPhrasesRepository(context.Background(), h.db)
	phrase := &repository.Phrase{
		SenderChannelId: strconv.FormatInt(message.Chat.ID, 10),
		SenderId:        message.From.UserName,
		SenderName:      fmt.Sprintf("%s %s", message.From.FirstName, message.From.LastName),
		PhraseText:      message.Text,
		Reply:           reply,
	}
	_, err = phrasesRepo.Create(phrase)
	if err != nil {
		log.Println(err)
		return
	}
}

func NewStatHandler(db *sql.DB) *Handler {
	return &Handler{
		db: db,
	}
}
