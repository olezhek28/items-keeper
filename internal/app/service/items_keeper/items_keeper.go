package items_keeper

import (
	"github.com/olezhek28/items-keeper/internal/app/repository"
	"github.com/olezhek28/items-keeper/internal/pkg/api/pocket"
	"github.com/olezhek28/items-keeper/internal/pkg/api/telegram"
)

type ItemsKeeperService struct {
	tgClient     telegram.ITelegramClient
	pocketClient pocket.IPocketClient

	tokensRepository repository.ITokensRepository
}

func NewItemsKeeperService(
	tgClient telegram.ITelegramClient,
	pocketClient pocket.IPocketClient,

	tokensRepository repository.ITokensRepository,
) *ItemsKeeperService {
	return &ItemsKeeperService{
		tgClient:         tgClient,
		pocketClient:     pocketClient,
		tokensRepository: tokensRepository,
	}
}

func NewMockItemsKeeperService(deps ...interface{}) *ItemsKeeperService {
	ik := ItemsKeeperService{}

	for _, v := range deps {
		switch s := v.(type) {
		case telegram.ITelegramClient:
			ik.tgClient = s
		case pocket.IPocketClient:
			ik.pocketClient = s

		case repository.ITokensRepository:
			ik.tokensRepository = s
		}
	}

	return &ik
}
