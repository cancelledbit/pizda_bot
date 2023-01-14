package main

import (
	"context"
	"database/sql"
	"github.com/cancelledbit/pizda_bot/app/repository"
	"github.com/cancelledbit/pizda_bot/app/stat"
	"github.com/cancelledbit/pizda_bot/app/stickers"
	_ "github.com/go-sql-driver/mysql"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/joho/godotenv"
	"log"
	"os"
	"strconv"
	"time"
)

func main() {
	initEnv()
	bot := initBot()

	throttlingTimeout := getThrottlingTimeout()

	timeoutMap := make(map[string]int64)

	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	updateTimeout := func(ctx context.Context, tick <-chan time.Time) {
		for {
			select {
			case <-tick:
				clearTimeout(timeoutMap, throttlingTimeout)
			case <-ctx.Done():
				return
			}
		}
	}

	ticker := time.Tick(time.Second * time.Duration(throttlingTimeout/2))

	go updateTimeout(ctx, ticker)

	db := repository.GetDbPool()
	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {
			log.Println(err)
		}
	}(db)
	statHandler := stat.NewStatHandler(db)
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

		if sticker, err := stickers.GetStickerBy(update.Message.Text); err == nil {
			statHandler.PushStat(update.Message, sticker.Name)
			file := tgbotapi.FileID(sticker.ID)
			msg := tgbotapi.NewSticker(update.Message.Chat.ID, file)
			msg.ReplyToMessageID = update.Message.MessageID
			timeoutMap[from] = time.Now().Unix()
			log.Println(msg)
			//if _, err := bot.Send(msg); err != nil {
			//	log.Println(err)
			//}
		}
	}
}

func clearTimeout(timeoutMap map[string]int64, timeout int) {
	cTime := time.Now().Add(-(time.Second * time.Duration(timeout))).Unix()
	for key, start := range timeoutMap {
		if cTime > start {
			delete(timeoutMap, key)
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

func getThrottlingTimeout() int {
	throttlingTimeout := 30
	if os.Getenv("THROTTLING") != "" {
		if throttlingEnv, err := strconv.Atoi(os.Getenv("THROTTLING")); err == nil {
			if throttlingEnv/2 != 0 {
				throttlingTimeout = throttlingEnv
			} else {
				log.Println("CANT USE VALUE LESS THAN 1 AS THROTTLING VALUE")
			}
		}
	}
	return throttlingTimeout
}
