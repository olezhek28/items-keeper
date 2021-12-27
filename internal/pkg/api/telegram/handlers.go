package telegram

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

func (c *client) handleCommand(message *tgBotAPI.Message) error {
	switch message.Command() {
	case commandStart:
		return c.handleCommandStart(message)
	default:
		return c.handleUnknownCommand(message)
	}
}

func (c *client) handleMessage(message *tgBotAPI.Message) {
	log.Printf("[%s] %s", message.From.UserName, message.Text)

	msg := tgBotAPI.NewMessage(message.Chat.ID, message.Text)

	c.tgBot.Send(msg)
}

func (c *client) handleCommandStart(message *tgBotAPI.Message) error {
	msg := tgBotAPI.NewMessage(message.Chat.ID, replyStartTemplate)

	_, err := c.tgBot.Send(msg)
	return err
}

func (c *client) handleUnknownCommand(message *tgBotAPI.Message) error {
	msg := tgBotAPI.NewMessage(message.Chat.ID, "Я не знаю такую команду")

	_, err := c.tgBot.Send(msg)
	return err
}
