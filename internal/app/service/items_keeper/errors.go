package items_keeper

import (
	"errors"

	tgBotAPI "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

const (
	errorTextFailedSaveLink = "Не удалось сохранить ссылку:( Попробуй ещё разок!"
	errorTextInvalidLink    = "Твоя ссылка невалидная:( Попробуй исправить это, я верю в тебя!"
	errorTextUnauthorized   = "Ты ещё не авторизирован. Вызови команду /start"
	errorTextUnknownError   = "Произошла неизвестная ошибка."
)

var (
	errorFailedSaveLink = errors.New(errorTextFailedSaveLink)
	errorInvalidLink    = errors.New(errorTextInvalidLink)
	errorUnauthorized   = errors.New(errorTextUnauthorized)
)

func (s *ItemsKeeperService) handleError(chatID int64, err error) {
	msg := tgBotAPI.NewMessage(chatID, errorTextUnknownError)

	if err == errorFailedSaveLink || err == errorInvalidLink || err == errorUnauthorized {
		msg.Text = err.Error()
	}

	s.tgClient.Send(msg)
}
