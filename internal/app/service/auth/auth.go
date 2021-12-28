package auth

import (
	"net/http"

	"github.com/olezhek28/items-keeper/internal/app/repository"
	"github.com/olezhek28/items-keeper/internal/pkg/api/pocket"
)

type AuthService struct {
	pocketClient pocket.IPocketClient

	tokensRepository repository.ITokensRepository

	server      *http.Server
	redirectURL string
}

func NewAuthService(
	pocketClient pocket.IPocketClient,

	tokensRepository repository.ITokensRepository,

	redirectURL string,
) *AuthService {
	return &AuthService{
		pocketClient: pocketClient,

		tokensRepository: tokensRepository,

		redirectURL: redirectURL,
	}
}

func NewMockAuthService(deps ...interface{}) *AuthService {
	a := AuthService{}

	for _, v := range deps {
		switch s := v.(type) {
		case pocket.IPocketClient:
			a.pocketClient = s

		case repository.ITokensRepository:
			a.tokensRepository = s

		}
	}

	return &a
}
