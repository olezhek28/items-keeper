package telegram

import (
	"fmt"
	"log"

	tgBotAPI "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

const commandStart = "start"

func (c *client) handleCommand(message *tgBotAPI.Message) error {
	msg := tgBotAPI.NewMessage(message.Chat.ID, "Я не знаю такую команду")

	switch message.Command() {
	case commandStart:
		msg.Text = fmt.Sprintf("Привет, %s", message.From.FirstName)
	default:
	}

	_, err := c.bot.Send(msg)
	return err
}

func (c *client) handleMessage(message *tgBotAPI.Message) {
	log.Printf("[%s] %s", message.From.UserName, message.Text)

	msg := tgBotAPI.NewMessage(message.Chat.ID, message.Text)

	c.bot.Send(msg)
}
