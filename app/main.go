package main

import (
	"context"
	"database/sql"
	"log"
	"math/rand"
	"os"
	"regexp"
	"strconv"
	"time"

	"github.com/cancelledbit/pizda_bot/app/commands"
	"github.com/cancelledbit/pizda_bot/app/repository"
	"github.com/cancelledbit/pizda_bot/app/stat"
	"github.com/cancelledbit/pizda_bot/app/stickers"
	_ "github.com/go-sql-driver/mysql"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/joho/godotenv"
)

var chatIDs = [...]int64{
	-1001169383931, // https://t.me/pr2ch
}

func main() {
	initEnv()
	bot := initBot()

	updateMyCommands(bot, chatIDs[:])

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
	cmdHandler := commands.GetHandler(db, bot)
	for update := range updates {
		if update.Message == nil { // ignore any non-Message updates
			continue
		}

		from := update.Message.From.String()
		handleSpecialChatEvents(update, db)
		if _, ok := timeoutMap[from]; ok {
			continue
		}

		limitTs := time.Now().Add(-2 * time.Minute)
		if update.Message.Time().Before(limitTs) {
			continue
		}

		if update.Message.IsCommand() {
			timeoutMap[from] = time.Now().Unix()
			cmdHandler.Handle(update.Message)
			continue
		}

		if sticker, err := stickers.GetStickerBy(update.Message.Text); err == nil {
			statHandler.PushStat(update.Message, sticker.Name)
			file := tgbotapi.FileID(sticker.ID)
			msg := tgbotapi.NewSticker(update.Message.Chat.ID, file)
			msg.ReplyToMessageID = update.Message.MessageID
			if !isShouldReply() {
				continue
			}
			timeoutMap[from] = time.Now().Unix()

			if _, err := bot.Send(msg); err != nil {
				log.Println(err)
			}
		}
	}
}

func handleSpecialChatEvents(update tgbotapi.Update, db *sql.DB) {
	pattern, err := commands.GetWisdomUserPattern(update.Message.From.ID)
	if err != nil {
		return
	}
	if rgx, err := regexp.Compile(pattern); err == nil {
		log.Println("compiled")
		if rgx.MatchString(update.Message.Text) {
			log.Println("matched")
			r := repository.UserWisdomRepository(context.Background(), db)
			_, _ = r.Create(
				&repository.WisdomPhrase{Text: update.Message.Text,
					AuthorId: strconv.FormatInt(update.Message.From.ID, 10),
				},
			)
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

func updateMyCommands(bot *tgbotapi.BotAPI, chatIDs []int64) {
	for _, chatID := range chatIDs {
		scope := tgbotapi.NewBotCommandScopeChat(chatID)

		Cmds := []tgbotapi.BotCommand{
			{
				Command:     "/gpt",
				Description: "Generate Wisdom (OpenAI GPT-4 Turbo) 🗿🗿",
			},
			{
				Command:     "/hz",
				Description: "Sends Hagz0r Wisdom",
			},
			{
				Command:     "/eu",
				Description: "Sends Wild Billy (EU) Wisdom",
			},
			{
				Command:     "/dk",
				Description: "Sends Double Keeper Wisdom",
			},
			{
				Command:     "/anton",
				Description: "Sends Anton Wisdom",
			},
		}

		deleteAllCommands := tgbotapi.NewDeleteMyCommands() // This deletes commands from all chats !
		bot.Request(deleteAllCommands)

		config := tgbotapi.NewSetMyCommandsWithScope(scope, Cmds...)
		bot.Request(config)
	}

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

func isShouldReply() bool {
	chance := 5
	if os.Getenv("CHANCE") != "" {
		if chanceEnv, err := strconv.Atoi(os.Getenv("CHANCE")); err == nil {
			chance = chanceEnv
		} else {
			log.Printf("chance not set using default %d", chance)
		}
	}
	return rand.Intn(chance) == 1
}
