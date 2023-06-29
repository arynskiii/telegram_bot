package main

import (
	"log"
	"telbot/pkg/telegram"

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

	telegramBot := telegram.NewBot(bot, pocketClient, "https://www.google.com/")
	if err := telegramBot.Start(); err != nil {
		log.Fatal(err)
	}

}
