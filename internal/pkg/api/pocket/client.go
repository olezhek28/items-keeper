package pocket

//go:generate mockgen --build_flags=--mod=mod -destination=mocks/mock_pocket_client.go -package=mocks . IPocketClient

import (
	"context"
	"fmt"

	pocketSDK "github.com/zhashkevych/go-pocket-sdk"
)

type IPocketClient interface {
	Add(ctx context.Context, accessToken, URL string) error
	Authorize(ctx context.Context, requestToken string) (*pocketSDK.AuthorizeResponse, error)
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

func (c *client) Add(ctx context.Context, accessToken, URL string) error {
	return c.pocketSDKClient.Add(ctx, pocketSDK.AddInput{
		URL:         URL,
		AccessToken: accessToken,
	})
}

func (c *client) Authorize(ctx context.Context, requestToken string) (*pocketSDK.AuthorizeResponse, error) {
	return c.pocketSDKClient.Authorize(ctx, requestToken)
}

func (c *client) GetRequestToken(ctx context.Context, chatID int64) (string, error) {
	c.redirectURL = c.generateRedirectURL(chatID)

	requestToken, err := c.pocketSDKClient.GetRequestToken(ctx, c.redirectURL)
	if err != nil {
		return "", err
	}

	return requestToken, err
}

func (c *client) GetAuthorizationLink(requestToken string) (string, error) {
	return c.pocketSDKClient.GetAuthorizationURL(requestToken, c.redirectURL)
}

func (c *client) generateRedirectURL(chatID int64) string {
	return fmt.Sprintf("%s?chat_id=%d", c.redirectURL, chatID)
}
