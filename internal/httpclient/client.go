package httpclient

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
)

// DoRequest 指定のHTTPクライアントでリクエストを実行し、レスポンスボディを返す
func DoRequest(httpClient *http.Client, method, endpoint, path string, params url.Values) ([]byte, error) {
	fullURL := endpoint + path

	req, err := buildRequest(method, fullURL, params)
	if err != nil {
		return nil, fmt.Errorf("リクエストの作成に失敗しました: %w", err)
	}

	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("HTTPリクエストの送信に失敗しました: %w", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("レスポンスの読み取りに失敗しました: %w", err)
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("HTTPステータスエラー %d: %s", resp.StatusCode, string(body))
	}

	return body, nil
}

// isGetRequest GETメソッドかつパラメータが存在するか判定
func isGetRequest(method string, params url.Values) bool {
	return method == http.MethodGet && params != nil && len(params) > 0
}

// buildRequest HTTPリクエストを生成する
func buildRequest(method, fullURL string, params url.Values) (*http.Request, error) {
	if isGetRequest(method, params) {
		fullURL += "?" + params.Encode()
		return http.NewRequest(method, fullURL, nil)
	}

	body := strings.NewReader(params.Encode())
	req, err := http.NewRequest(method, fullURL, body)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	return req, nil
}
