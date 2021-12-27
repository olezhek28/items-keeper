package telegram

import (
	"log"

	tgBotAPI "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type ITelegramClient interface {
	Start() error
}

type client struct {
	bot *tgBotAPI.BotAPI
}

func NewClient(bot *tgBotAPI.BotAPI) ITelegramClient {
	return &client{bot: bot}
}

func (c *client) Start() error {
	c.bot.Debug = true

	log.Printf("Authorized on account %s", c.bot.Self.UserName)

	u := tgBotAPI.NewUpdate(0)
	u.Timeout = 60

	updates := c.bot.GetUpdatesChan(u)

	for update := range updates {
		// If we got a message
		if update.Message != nil {
			log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)

			msg := tgBotAPI.NewMessage(update.Message.Chat.ID, update.Message.Text)
			msg.ReplyToMessageID = update.Message.MessageID

			c.bot.Send(msg)
		}
	}

	return nil
}
