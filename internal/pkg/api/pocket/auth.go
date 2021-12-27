package pocket

import (
	"context"
	"fmt"
)

func (c *client) GetRequestToken(ctx context.Context, chatID int64) (string, error) {
	redirectURL := c.generateRedirectURL(chatID)

	requestToken, err := c.pocketSDKClient.GetRequestToken(ctx, redirectURL)
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
