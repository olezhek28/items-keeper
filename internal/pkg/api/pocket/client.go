package pocket

import pocketSDK "github.com/zhashkevych/go-pocket-sdk"

type IPocketClient interface {
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
