package zaim

import (
	"net/http"
	"net/url"

	"github.com/ryohidaka/go-zaim/internal/auth"
	"github.com/ryohidaka/go-zaim/internal/httpclient"
)

type Client struct {
	oAuth1Client *http.Client
	httpClient   *http.Client
	endpoint     string
}

// OAuth1認証情報をまとめたパラメータ構造体
type ZaimParams struct {
	ConsumerKey    string
	ConsumerSecret string
	Token          string
	TokenSecret    string
}

// NewZaimClient は、Zaim API と通信するための新しいクライアントを作成する。
// 認証情報が提供されている場合、OAuth 1.0a を用いた認証付きクライアントを生成する。
//
// [Authorize with Oauth 1.0a]
//
// [Authorize with Oauth 1.0a]: https://dev.zaim.net/home/api/authorize
func NewZaimClient(params ...ZaimParams) *Client {
	var oAuth1Client *http.Client

	if len(params) > 0 {
		oAuth1Client = auth.NewOAuth1Client(auth.OAuth1Params(params[0]))
	}

	return &Client{
		oAuth1Client: oAuth1Client,
		httpClient:   &http.Client{},
		endpoint:     BaseURL,
	}
}

// GETリクエストを送信する
func (c *Client) get(path string, params url.Values, useAuth bool) ([]byte, error) {
	client := c.httpClient
	if useAuth {
		client = c.oAuth1Client
	}
	return httpclient.DoRequest(client, http.MethodGet, c.endpoint, path, params)
}

// POSTリクエストを送信する
func (c *Client) post(path string, params url.Values) ([]byte, error) {
	return httpclient.DoRequest(c.oAuth1Client, http.MethodPost, c.endpoint, path, params)
}

// PUTリクエストを送信する
func (c *Client) put(path string, params url.Values) ([]byte, error) {
	return httpclient.DoRequest(c.oAuth1Client, http.MethodPut, c.endpoint, path, params)
}

// DELETEリクエストを送信する
func (c *Client) delete(path string) ([]byte, error) {
	return httpclient.DoRequest(c.oAuth1Client, http.MethodDelete, c.endpoint, path, nil)
}
