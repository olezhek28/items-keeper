package telegram

import (
	"log"

	tgBotAPI "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

const (
	timeout = 60
	offset  = 0
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
	log.Printf("Authorized on account %s", c.bot.Self.UserName)

	updates := c.initUpdatesChannel()
	c.handleUpdates(updates)

	return nil
}

func (c *client) handleUpdates(updates tgBotAPI.UpdatesChannel) {
	for update := range updates {
		// ignore any non-message updates
		if update.Message == nil {
			continue
		}

		c.handleMessage(update.Message)
	}
}

func (c *client) handleMessage(message *tgBotAPI.Message) {
	log.Printf("[%s] %s", message.From.UserName, message.Text)

	msg := tgBotAPI.NewMessage(message.Chat.ID, message.Text)
	msg.ReplyToMessageID = message.MessageID

	c.bot.Send(msg)
}

func (c *client) initUpdatesChannel() tgBotAPI.UpdatesChannel {
	u := tgBotAPI.NewUpdate(offset)
	u.Timeout = timeout

	return c.bot.GetUpdatesChan(u)
}
