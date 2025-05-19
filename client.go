package zaim

import (
	"net/http"
	"net/url"

	"github.com/ryohidaka/go-zaim/internal/auth"
	"github.com/ryohidaka/go-zaim/internal/httpclient"
)

type Client struct {
	httpClient *http.Client
	endpoint   string
}

// OAuth1認証情報をまとめたパラメータ構造体
type NewClientParams struct {
	ConsumerKey    string
	ConsumerSecret string
	Token          string
	TokenSecret    string
}

// 認証情報を使って新しいAPIクライアントを作成する
//
// [Authorize with Oauth 1.0a]
//
// [Authorize with Oauth 1.0a]: https://dev.zaim.net/home/api/authorize
func NewZaimClient(params NewClientParams) *Client {
	httpClient := auth.NewOAuth1Client(auth.OAuth1Params(params))
	return &Client{
		httpClient: httpClient,
		endpoint:   BaseURL,
	}
}

// GETリクエストを送信する
func (c *Client) get(path string, params url.Values) ([]byte, error) {
	return httpclient.DoRequest(c.httpClient, http.MethodGet, c.endpoint, path, params)
}

// POSTリクエストを送信する
func (c *Client) post(path string, params url.Values) ([]byte, error) {
	return httpclient.DoRequest(c.httpClient, http.MethodPost, c.endpoint, path, params)
}

// PUTリクエストを送信する
func (c *Client) put(path string, params url.Values) ([]byte, error) {
	return httpclient.DoRequest(c.httpClient, http.MethodPut, c.endpoint, path, params)
}

// DELETEリクエストを送信する
func (c *Client) delete(path string, params url.Values) ([]byte, error) {
	return httpclient.DoRequest(c.httpClient, http.MethodDelete, c.endpoint, path, params)
}
