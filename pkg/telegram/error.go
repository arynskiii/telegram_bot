package telegram

import (
	"errors"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

var (
	errInvalidURL   = errors.New("url is invalid")
	errUnauthorized = errors.New("user is not authorized")
	errUnableToSave = errors.New("unable to save")
)

func (b *Bot) handleError(chatId int64, err error) {
	msg := tgbotapi.NewMessage(chatId, "Произошла неизвестная ошибка")
	switch err {
	case errInvalidURL:
		msg.Text = "Это невалидная ссылка!"
	case errUnauthorized:
		msg.Text = "Ты не авторизирован. Используй команду /start"
	case errUnableToSave:
		msg.Text = "Увы не удалось сохранить ссылку. Попробуй еще раз позже!"
	default:
		b.bot.Send(msg)
	}

}
