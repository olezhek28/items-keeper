package items_keeper

import (
	"context"
	"fmt"
	"net/url"

	tgBotAPI "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/olezhek28/items-keeper/internal/app/repository"
)

const (
	commandStart = "start"
)

const (
	replyStartTemplate     = "Привет! Чтобы сохранять ссылки в своем Pocket аккаунте, для начала тебе необходимо дать мне на это доступ. Для этого переходи по ссылке:\n%s"
	replayAlreadyAuthorize = "Ты уже авторизирован, брат! можешь пользоваться. Кидай ссылку, а я её сохраню."
	replaySuccessSaveLink  = "Твоя ссылка успешно сохранена!"
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
	msg := tgBotAPI.NewMessage(message.Chat.ID, replaySuccessSaveLink)

	_, err := url.ParseRequestURI(message.Text)
	if err != nil {
		return errorInvalidLink
	}

	accessToken, err := s.getAccessToken(message.Chat.ID)
	if err != nil {
		return errorUnauthorized
	}

	err = s.pocketClient.Add(context.Background(), accessToken, message.Text)
	if err != nil {
		return errorFailedSaveLink
	}

	return s.tgClient.Send(msg)
}

func (s *ItemsKeeperService) handleCommandStart(message *tgBotAPI.Message) error {
	var msg tgBotAPI.MessageConfig

	_, err := s.getAccessToken(message.Chat.ID)
	if err != nil {
		return s.initAuthorizationProcess(message)
	}

	msg = tgBotAPI.NewMessage(message.Chat.ID, replayAlreadyAuthorize)

	return s.tgClient.Send(msg)
}

func (s *ItemsKeeperService) handleUnknownCommand(message *tgBotAPI.Message) error {
	msg := tgBotAPI.NewMessage(message.Chat.ID, "Я не знаю такую команду")

	return s.tgClient.Send(msg)
}

func (s *ItemsKeeperService) generateAuthorizationLink(ctx context.Context, chatID int64) (string, error) {
	requestToken, err := s.pocketClient.GetRequestToken(ctx, chatID)
	if err != nil {
		return "", err
	}

	err = s.tokensRepository.Save(chatID, requestToken, repository.BucketRequestTokens)
	if err != nil {
		return "", err
	}

	return s.pocketClient.GetAuthorizationLink(requestToken, chatID)
}

func (s *ItemsKeeperService) getAccessToken(chatID int64) (string, error) {
	return s.tokensRepository.Get(chatID, repository.BucketAccessTokens)
}

func (s *ItemsKeeperService) initAuthorizationProcess(message *tgBotAPI.Message) error {
	ctx := context.Background()
	authLink, err := s.generateAuthorizationLink(ctx, message.Chat.ID)
	if err != nil {
		return err
	}

	msg := tgBotAPI.NewMessage(message.Chat.ID, fmt.Sprintf(replyStartTemplate, authLink))

	return s.tgClient.Send(msg)
}
