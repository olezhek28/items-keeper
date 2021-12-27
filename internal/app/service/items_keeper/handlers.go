package items_keeper

import (
	"log"

	tgBotAPI "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

const (
	commandStart = "start"
)

const (
	replyStartTemplate = "Привет! Чтобы сохранять ссылки в своем Pocket аккаунте, для начала тебе необходимо дать мне на это доступ. Для этого переходи по ссылке:\n%s"
)

func (s *ItemsKeeperService) handleCommand(message *tgBotAPI.Message) error {
	switch message.Command() {
	case commandStart:
		return s.handleCommandStart(message)
	default:
		return s.handleUnknownCommand(message)
	}
}

func (s *ItemsKeeperService) handleMessage(message *tgBotAPI.Message) error {
	log.Printf("[%s] %s", message.From.UserName, message.Text)

	msg := tgBotAPI.NewMessage(message.Chat.ID, message.Text)

	return s.tgClient.Send(msg)
}

func (s *ItemsKeeperService) handleCommandStart(message *tgBotAPI.Message) error {
	msg := tgBotAPI.NewMessage(message.Chat.ID, replyStartTemplate)

	return s.tgClient.Send(msg)
}

func (s *ItemsKeeperService) handleUnknownCommand(message *tgBotAPI.Message) error {
	msg := tgBotAPI.NewMessage(message.Chat.ID, "Я не знаю такую команду")

	return s.tgClient.Send(msg)
}
