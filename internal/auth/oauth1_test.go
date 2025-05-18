package auth_test

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/ryohidaka/go-zaim/internal/auth"
)

// NewOAuth1Clientが返すクライアントでリクエストすると
// AuthorizationヘッダーにOAuth1署名が含まれているかテストする
func TestNewOAuth1Client_AuthorizationHeader(t *testing.T) {
	params := auth.OAuth1Params{
		ConsumerKey:    "key",
		ConsumerSecret: "secret",
		Token:          "token",
		TokenSecret:    "tokensecret",
	}

	client := auth.NewOAuth1Client(params)
	if client == nil {
		t.Fatal("NewOAuth1Clientがnilを返しました")
	}

	// OAuth1署名付きリクエストを受け取るモックサーバーを用意
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			t.Error("Authorizationヘッダーがありません")
		}
		if !strings.HasPrefix(authHeader, "OAuth ") {
			t.Errorf("AuthorizationヘッダーがOAuth署名ではありません: %s", authHeader)
		}
		w.WriteHeader(http.StatusOK)
	}))
	defer server.Close()

	// モックサーバーへGETリクエスト
	resp, err := client.Get(server.URL)
	if err != nil {
		t.Fatalf("リクエスト送信に失敗しました: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		t.Errorf("ステータスコードがOKではありません: got %d", resp.StatusCode)
	}
}
