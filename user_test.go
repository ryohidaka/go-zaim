package zaim_test

import (
	"log"
	"os"
	"testing"

	"github.com/jarcoal/httpmock"
	"github.com/ryohidaka/go-zaim"
	"github.com/ryohidaka/go-zaim/models"
	"github.com/ryohidaka/go-zaim/testutil"
	"github.com/stretchr/testify/assert"
)

func ExampleClient_FetchMe() {
	// クライアント初期化
	p := zaim.NewClientParams{
		ConsumerKey:    os.Getenv("ZAIM_CONSUMER_KEY"),
		ConsumerSecret: os.Getenv("ZAIM_CONSUMER_SECRET"),
		Token:          os.Getenv("ZAIM_TOKEN"),
		TokenSecret:    os.Getenv("ZAIM_TOKEN_SECRET"),
	}

	c := zaim.NewClient(p)

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

	// モックレスポンスを設定
	err := testutil.MockResponseFromFile(zaim.BaseURL+"home/user/verify", "me")
	assert.NoError(t, err)

	p := zaim.NewClientParams{
		ConsumerKey:    "dummy-key",
		ConsumerSecret: "dummy-secret",
		Token:          "dummy-token",
		TokenSecret:    "dummy-secret",
	}

	// ユーザー情報を取得
	c := zaim.NewClient(p)
	me, err := c.FetchMe()

	// レスポンスの確認
	assert.NoError(t, err)
	assert.Equal(t, int(me.ID), 10000000)
	assert.Equal(t, me.Name, "MyName")
	assert.Equal(t, me.Active, models.BoolInt(true))
}
