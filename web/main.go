package main

import (
	"context"
	"database/sql"
	"github.com/cancelledbit/pizda_bot/app/repository"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"log"
)

func main() {
	initEnv()
	db := repository.GetDbPool()
	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {
			log.Panic(err)
		}
	}(db)

	app := fiber.New()

	app.Get("/channels", func(ctx *fiber.Ctx) error {
		channelRepo := repository.NewMysqlChannelRepository(context.Background(), db)
		channels, err := channelRepo.GetByOffset(0, 100)
		if err != nil {
			return err
		}
		return ctx.JSON(channels)
	})
	app.Get("/phrases", func(ctx *fiber.Ctx) error {
		phrasesRepository := repository.NewMysqlPhrasesRepository(context.Background(), db)
		phrases, err := phrasesRepository.GetByOffset(0, 100)
		if err != nil {
			return err
		}
		err = ctx.JSON(phrases)
		ctx.Set("content-type", "application/json; charset=utf-8")
		return err
	})
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	err := app.Listen(":3000")
	if err != nil {
		log.Panic(err)
	}
}

func initEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}
