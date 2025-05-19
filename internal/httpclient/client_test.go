package httpclient_test

import (
	"io"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"

	"github.com/ryohidaka/go-zaim/internal/httpclient"
)

// テスト用のHTTPサーバーを作成し、正常レスポンスを返す
func TestDoRequest_GetMethod_Success(t *testing.T) {
	// テストサーバー作成
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			t.Errorf("期待したHTTPメソッド: GET, 実際: %s", r.Method)
		}
		if r.URL.RawQuery != "foo=bar" {
			t.Errorf("期待したクエリパラメータ: foo=bar, 実際: %s", r.URL.RawQuery)
		}
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte("成功"))
	}))
	defer ts.Close()

	params := url.Values{}
	params.Set("foo", "bar")

	client := ts.Client()

	body, err := httpclient.DoRequest(client, http.MethodGet, ts.URL, "/", params)
	if err != nil {
		t.Fatalf("エラー発生: %v", err)
	}
	if string(body) != "成功" {
		t.Errorf("期待したレスポンスボディ: 成功, 実際: %s", string(body))
	}
}

// POSTメソッドでのリクエスト送信と正常レスポンスを検証
func TestDoRequest_PostMethod_Success(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			t.Errorf("期待したHTTPメソッド: POST, 実際: %s", r.Method)
		}
		ct := r.Header.Get("Content-Type")
		if !strings.Contains(ct, "application/x-www-form-urlencoded") {
			t.Errorf("期待したContent-Typeに含む: application/x-www-form-urlencoded, 実際: %s", ct)
		}
		bodyBytes, _ := io.ReadAll(r.Body)
		if string(bodyBytes) != "foo=bar" {
			t.Errorf("期待したリクエストボディ: foo=bar, 実際: %s", string(bodyBytes))
		}
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte("成功POST"))
	}))
	defer ts.Close()

	params := url.Values{}
	params.Set("foo", "bar")

	client := ts.Client()

	body, err := httpclient.DoRequest(client, http.MethodPost, ts.URL, "/", params)
	if err != nil {
		t.Fatalf("エラー発生: %v", err)
	}
	if string(body) != "成功POST" {
		t.Errorf("期待したレスポンスボディ: 成功POST, 実際: %s", string(body))
	}
}

// HTTPステータスエラー時にエラーが返ることを確認
func TestDoRequest_StatusError(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusBadRequest)
		_, _ = w.Write([]byte("エラーメッセージ"))
	}))
	defer ts.Close()

	client := ts.Client()

	_, err := httpclient.DoRequest(client, http.MethodGet, ts.URL, "/", nil)
	if err == nil {
		t.Fatal("エラーが返ることを期待していましたがnilでした")
	}
	if !strings.Contains(err.Error(), "HTTPステータスエラー 400") {
		t.Errorf("期待したエラーメッセージにHTTPステータスコードを含む: %v", err)
	}
}

// mapping=1 がparamsに自動付与されることを確認
func TestDoRequest_AppendsMapping(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		values := r.URL.Query()
		if values.Get("mapping") != "1" {
			t.Errorf("mapping パラメータが付与されていません: %s", r.URL.RawQuery)
		}
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte("OK"))
	}))
	defer ts.Close()

	params := url.Values{}
	params.Set("foo", "bar") // paramsがnilでないことを保証

	client := ts.Client()
	body, err := httpclient.DoRequest(client, http.MethodGet, ts.URL, "/", params)
	if err != nil {
		t.Fatalf("エラー発生: %v", err)
	}
	if string(body) != "OK" {
		t.Errorf("期待したレスポンスボディ: OK, 実際: %s", string(body))
	}
}
