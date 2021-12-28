package items_keeper

import (
	"fmt"

	tgBotAPI "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func (s *ItemsKeeperService) Process() error {
	updates, err := s.tgClient.Start()
	if err != nil {
		return fmt.Errorf("failed to starting telegram client: %s", err.Error())
	}

	s.handleUpdates(updates)

	return nil
}

func (s *ItemsKeeperService) handleUpdates(updates tgBotAPI.UpdatesChannel) {
	for update := range updates {
		// ignore any non-message updates
		if update.Message == nil {
			continue
		}

		if update.Message.IsCommand() {
			err := s.handleCommand(update.Message)
			if err != nil {
				s.handleError(update.Message.Chat.ID, err)
			}

			continue
		}

		err := s.handleMessage(update.Message)
		if err != nil {
			s.handleError(update.Message.Chat.ID, err)
		}
	}
}
