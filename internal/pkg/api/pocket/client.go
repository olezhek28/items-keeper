package pocket

//go:generate mockgen --build_flags=--mod=mod -destination=mocks/mock_pocket_client.go -package=mocks . IPocketClient

import (
	"context"

	pocketSDK "github.com/zhashkevych/go-pocket-sdk"
)

type IPocketClient interface {
	GetRequestToken(ctx context.Context, chatID int64) (string, error)
	GetAuthorizationLink(requestToken string) (string, error)
}

type client struct {
	pocketSDKClient *pocketSDK.Client
	redirectURL     string
}

func NewClient(pocketSDKClient *pocketSDK.Client, redirectURL string) IPocketClient {
	return &client{
		pocketSDKClient: pocketSDKClient,
		redirectURL:     redirectURL,
	}
}
