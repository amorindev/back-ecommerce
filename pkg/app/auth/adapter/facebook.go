package adapter

type FacebookProvider struct {
	AppID     string
	AppSecret string
}

// ?oidc , donde lo uso?
func NewFacebookProvider(appID, appSecret string) *FacebookProvider{
	return &FacebookProvider{
		AppID: appID,
		AppSecret: appSecret,
	}
}