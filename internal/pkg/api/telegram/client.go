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
	tgClient *tgBotAPI.BotAPI
}

func NewClient(tgClient *tgBotAPI.BotAPI) ITelegramClient {
	return &client{tgClient: tgClient}
}

func (c *client) Start() error {
	log.Printf("Authorized on account %s", c.tgClient.Self.UserName)

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

		if update.Message.IsCommand() {
			c.handleCommand(update.Message)
			continue
		}

		c.handleMessage(update.Message)
	}
}

func (c *client) initUpdatesChannel() tgBotAPI.UpdatesChannel {
	u := tgBotAPI.NewUpdate(offset)
	u.Timeout = timeout

	return c.tgClient.GetUpdatesChan(u)
}
