package tumblgo

import "github.com/mrjones/oauth"

var (
	// OAuthServiceProvider oauth service provider for tumblr
	OAuthServiceProvider oauth.ServiceProvider
)

func init() {
	OAuthServiceProvider = oauth.ServiceProvider{
		RequestTokenUrl:   "https://www.tumblr.com/oauth/request_token",
		AuthorizeTokenUrl: "https://www.tumblr.com/oauth/authorize",
		AccessTokenUrl:    "https://www.tumblr.com/oauth/access_token",
	}
}

// OAuth oauthのクライアント
type OAuth struct {
	consumer     *oauth.Consumer
	requestToken *oauth.RequestToken
	loginURL     string
}

// getRequestTokenAndURL .
func (o OAuth) getRequestTokenAndURL() error {
	var err error
	o.requestToken, o.loginURL, err = o.consumer.GetRequestTokenAndUrl("oob")
	if err != nil {
		return err
	}

	return nil
}

// NewOAuth oauth client
func NewOAuth(key, secret string) (*OAuth, error) {
	o := OAuth{}
	o.consumer = oauth.NewConsumer(key, secret, OAuthServiceProvider)

	if err := o.getRequestTokenAndURL(); err != nil {
		return &o, err
	}

	return &o, nil
}
