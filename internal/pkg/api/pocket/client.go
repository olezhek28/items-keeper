package pocket

import pocketSDK "github.com/zhashkevych/go-pocket-sdk"

type IPocketClient interface {
}

type client struct {
	pocketClient *pocketSDK.Client
	redirectURL  string
}

func NewClient(pocketClient *pocketSDK.Client, redirectURL string) IPocketClient {
	return &client{
		pocketClient: pocketClient,
		redirectURL:  redirectURL,
	}
}
