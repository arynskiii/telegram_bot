package telegram

import (
	"log"
	"telbot/pkg/repository"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/zhashkevych/go-pocket-sdk"
)

type Bot struct {
	bot             *tgbotapi.BotAPI
	pocketClient    *pocket.Client
	redirectURL     string
	tokenRepository repository.TokenRepository
}

func NewBot(bot *tgbotapi.BotAPI, tokenRepository repository.TokenRepository, pocketClient *pocket.Client, redirectURL string) *Bot {
	return &Bot{
		bot:             bot,
		pocketClient:    pocketClient,
		redirectURL:     redirectURL,
		tokenRepository: tokenRepository,
	}
}

func (b *Bot) Start() error {
	log.Printf("Authorized on account %s", b.bot.Self.UserName)
	updates, err := b.initUpdatesChannel()
	if err != nil {
		log.Fatal(err)
	}
	b.handleUpdates(updates)

	return nil
}

func (b *Bot) handleUpdates(updates tgbotapi.UpdatesChannel) {
	for update := range updates {
		if update.Message == nil {
			continue
		}
		if update.Message.IsCommand() {
			if err := b.handleCommand(update.Message); err != nil {
				b.handleError(update.Message.Chat.ID, err)
			}

			continue
		}
		if err := b.handleMessage(update.Message); err != nil {
			b.handleError(update.Message.Chat.ID, err)
		}
	}

}

func (b *Bot) initUpdatesChannel() (tgbotapi.UpdatesChannel, error) {
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	return b.bot.GetUpdatesChan(u), nil

}
