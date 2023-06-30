package main

import (
	"log"
	"telbot/pkg/repository"
	"telbot/pkg/repository/boltdb"
	"telbot/pkg/server"
	"telbot/pkg/telegram"

	"github.com/boltdb/bolt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/zhashkevych/go-pocket-sdk"
)

func main() {
	bot, err := tgbotapi.NewBotAPI("6102466922:AAG7Z7hwoqXqat4MahnIX12dr2_7qffap6U")
	if err != nil {
		log.Fatal(err)
	}
	bot.Debug = true
	pocketClient, err := pocket.NewClient("107911-c15f221d9385259dbb63120")
	if err != nil {
		log.Fatal(err)
	}
	db, err := InitDb()
	if err != nil {
		log.Fatal(err)
	}
	tokenRepository := boltdb.NewTokenRepository(db)

	telegramBot := telegram.NewBot(bot, tokenRepository, pocketClient, "http://localhost/")
	authorizationServer := server.NewAuthorizationServer(pocketClient, tokenRepository, "https://t.me/poket_aryn_bot")

	go func() {
		if err := telegramBot.Start(); err != nil {
			log.Fatal(err)
		}

	}()
	if err := authorizationServer.Start(); err != nil {
		log.Fatal(err)

	}
}
func InitDb() (*bolt.DB, error) {
	db, err := bolt.Open("bot.db", 0600, nil)
	if err != nil {
		log.Fatal(err)
	}
	db.Update(func(tx *bolt.Tx) error {
		_, err := tx.CreateBucketIfNotExists([]byte(repository.AccessTokens))
		if err != nil {
			return err
		}
		_, err = tx.CreateBucketIfNotExists([]byte(repository.RequesTokens))
		if err != nil {
			return err
		}
		return nil
	})
	return db, nil
}
