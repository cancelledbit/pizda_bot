package main

import (
	"context"
	"errors"
	"github.com/cancelledbit/pizda_bot/stickers"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/joho/godotenv"
	"log"
	"os"
	"regexp"
	"time"
)

func main() {
	initEnv()
	bot := initBot()

	timeoutMap := make(map[string]int64)

	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	updateTimeout := func(ctx context.Context, tick <-chan time.Time) {
		for {
			select {
			case <-tick:
				clearTimeout(&timeoutMap)
			case <-ctx.Done():
				return
			}
		}
	}
	ticker := time.Tick(time.Second * 10)

	go updateTimeout(ctx, ticker)
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message == nil { // ignore any non-Message updates
			continue
		}

		if update.Message.IsCommand() { // ignore any non-command Messages
			continue
		}

		from := update.Message.From.String()
		if _, ok := timeoutMap[from]; ok {
			continue
		}

		if sticker, err := chooseSticker(update.Message.Text); err == nil {
			file := tgbotapi.FileID(sticker)
			msg := tgbotapi.NewSticker(update.Message.Chat.ID, file)
			msg.ReplyToMessageID = update.Message.MessageID
			timeoutMap[from] = time.Now().Unix()
			if _, err := bot.Send(msg); err != nil {
				log.Panic(err)
			}
		}
	}
}

func clearTimeout(timeoutMap *map[string]int64) {
	cTime := time.Now().Add(-(time.Second * 15)).Unix()
	for key, start := range *timeoutMap {
		if cTime > start {
			delete(*timeoutMap, key)
		}
	}
}

func initEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func initBot() (bot *tgbotapi.BotAPI) {
	bot, err := tgbotapi.NewBotAPI(os.Getenv("TELEGRAM_APITOKEN"))
	if err != nil {
		panic(err)
	}
	bot.Debug = true
	return
}

func chooseSticker(msg string) (string, error) {
	r, _ := regexp.Compile("^[Дд][Аа][!.?]{0,3}$")
	if r.MatchString(msg) {
		return stickers.Pizda, nil
	}
	r, _ = regexp.Compile("^[Нн][Ее][Тт][!.?]{0,3}$")
	if r.MatchString(msg) {
		return stickers.Minet, nil
	}
	r, _ = regexp.Compile("^[Яя][!.?]{0,3}$")
	if r.MatchString(msg) {
		return stickers.Golovka, nil
	}
	return "", errors.New("unknown")
}
