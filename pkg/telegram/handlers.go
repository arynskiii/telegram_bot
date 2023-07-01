package telegram

import (
	"context"
	"net/url"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/zhashkevych/go-pocket-sdk"
)

const (
	commandStart           = "start"
	ReplyStartTemplate     = "Привет! Чтобы сохранить ссылки в своем Pocket аккаунте,для начала тебе необходимо дать мне на это доступ. Для этого переходи по ссылке:\n%s"
	replyAlreadyAuthorized = "Ты уже авторизирован. Присылай ссылку, а я ее сохраню!"
)

func (b *Bot) handleCommand(message *tgbotapi.Message) error {

	switch message.Command() {
	case commandStart:
		return b.HandleStartCommand(message)
	default:
		return b.HandleUnknownCommand(message)
	}
}

func (b *Bot) handleMessage(message *tgbotapi.Message) error {
	_, err := url.ParseRequestURI(message.Text)
	if err != nil {
		return errInvalidURL
	}
	accessToken, err := b.GetAccessToken(message.Chat.ID)
	if err != nil {
		return errUnauthorized

	}
	if err := b.pocketClient.Add(context.Background(), pocket.AddInput{
		AccessToken: accessToken,
		URL:         message.Text,
	}); err != nil {
		return errUnableToSave
	}
	msg := tgbotapi.NewMessage(message.Chat.ID, "Ссылка успешно сохранена")

	_, err = b.bot.Send(msg)
	return err
}

func (b *Bot) HandleStartCommand(message *tgbotapi.Message) error {
	_, err := b.GetAccessToken(message.Chat.ID)
	if err != nil {
		return b.initAuthorizationProcess(message)
	}
	msg := tgbotapi.NewMessage(message.Chat.ID, replyAlreadyAuthorized)
	b.bot.Send(msg)
	return err

}

func (b *Bot) HandleUnknownCommand(message *tgbotapi.Message) error {
	msg := tgbotapi.NewMessage(message.Chat.ID, "Я не знаю такой команды!")
	_, err := b.bot.Send(msg)
	return err
}
