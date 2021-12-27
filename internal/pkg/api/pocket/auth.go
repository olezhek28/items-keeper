package pocket

import (
	"context"
	"fmt"
)

func (c *client) generateAuthorizationLink(ctx context.Context, chatID int64) (string, error) {
	redirectURL := c.generateRedirectURL(chatID)

	requestToken, err := c.pocketClient.GetRequestToken(ctx, redirectURL)
	if err != nil {
		return "", err
	}

	return c.pocketClient.GetAuthorizationURL(requestToken, c.redirectURL)
}

func (c *client) generateRedirectURL(chatID int64) string {
	return fmt.Sprintf("%s?chat_id=%d", c.redirectURL, chatID)
}
