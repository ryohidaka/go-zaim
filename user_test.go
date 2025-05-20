package zaim_test

import (
	"log"
	"os"
	"testing"

	"github.com/jarcoal/httpmock"
	"github.com/ryohidaka/go-zaim"
	"github.com/ryohidaka/go-zaim/testutil"
	"github.com/stretchr/testify/assert"
)

func ExampleClient_FetchMe() {
	// クライアント初期化
	p := zaim.ZaimParams{
		ConsumerKey:    os.Getenv("ZAIM_CONSUMER_KEY"),
		ConsumerSecret: os.Getenv("ZAIM_CONSUMER_SECRET"),
		Token:          os.Getenv("ZAIM_TOKEN"),
		TokenSecret:    os.Getenv("ZAIM_TOKEN_SECRET"),
	}

	c := zaim.NewZaimClient(p)

	// ユーザー情報を取得
	me, err := c.FetchMe()
	if err != nil {
		log.Fatal(err)
	}

	println(me.Name)
}

func TestFetchMe(t *testing.T) {
	// モックのHTTPサーバーを有効化
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	p := zaim.ZaimParams{
		ConsumerKey:    "dummy-key",
		ConsumerSecret: "dummy-secret",
		Token:          "dummy-token",
		TokenSecret:    "dummy-secret",
	}
	c := zaim.NewZaimClient(p)

	t.Run("正常系: ユーザー情報を正しく取得できる", func(t *testing.T) {
		// 正常なレスポンスを返すモックを設定
		err := testutil.MockResponseFromFile(zaim.BaseURL+"home/user/verify", "me")
		assert.NoError(t, err)

		// ユーザー情報を取得
		me, err := c.FetchMe()

		// レスポンスの確認
		assert.NoError(t, err)
		assert.Equal(t, int(me.ID), 10000000)
		assert.Equal(t, me.Name, "MyName")
		assert.Equal(t, me.Active, true)
	})

	t.Run("異常系: サーバーエラー時にエラーを返却する", func(t *testing.T) {
		// エラーを返すモックに差し替え
		httpmock.RegisterResponder("GET", zaim.BaseURL+"home/user/verify",
			httpmock.NewStringResponder(500, `Internal Server Error`))

		// ユーザー情報を取得
		me, err := c.FetchMe()

		// レスポンスの確認
		assert.Error(t, err)
		assert.Empty(t, me)
	})
}
