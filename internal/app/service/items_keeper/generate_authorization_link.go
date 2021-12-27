package items_keeper

import (
	"context"

	"github.com/olezhek28/items-keeper/internal/app/repository"
)

func (s *ItemsKeeperService) generateAuthorizationLink(ctx context.Context, chatID int64) (string, error) {
	requestToken, err := s.pocketClient.GetRequestToken(ctx, chatID)
	if err != nil {
		return "", err
	}

	err = s.tokensRepository.Save(chatID, requestToken, repository.BucketRequestTokens)
	if err != nil {
		return "", err
	}

	return s.pocketClient.GetAuthorizationLink(requestToken)
}
