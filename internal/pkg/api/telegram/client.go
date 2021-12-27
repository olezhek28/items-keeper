package telegram

//go:generate mockgen --build_flags=--mod=mod -destination=mocks/mock_telegram_client.go -package=mocks . ITelegramClient

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
	tgBot *tgBotAPI.BotAPI
}

func NewClient(tgClient *tgBotAPI.BotAPI) ITelegramClient {
	return &client{tgBot: tgClient}
}

func (c *client) Start() error {
	log.Printf("Authorized on account %s", c.tgBot.Self.UserName)

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

	return c.tgBot.GetUpdatesChan(u)
}
