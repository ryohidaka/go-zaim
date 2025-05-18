package auth

import (
	"net/http"

	"github.com/dghubble/oauth1"
)

type OAuth1Params struct {
	ConsumerKey    string
	ConsumerSecret string
	Token          string
	TokenSecret    string
}

// OAuth1 認証情報を使って認証済みの HTTP クライアントを作成する
func NewOAuth1Client(p OAuth1Params) *http.Client {
	c := oauth1.NewConfig(p.ConsumerKey, p.ConsumerSecret)
	t := oauth1.NewToken(p.Token, p.TokenSecret)
	return c.Client(oauth1.NoContext, t)
}
