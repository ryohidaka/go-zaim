package zaim_test

import (
	"log"
	"os"
	"testing"
	"time"

	"github.com/jarcoal/httpmock"
	"github.com/ryohidaka/go-zaim"
	"github.com/ryohidaka/go-zaim/testutil"
	"github.com/stretchr/testify/assert"
)

func ExampleClient_CreateTransfer() {
	// クライアント初期化
	p := zaim.ZaimParams{
		ConsumerKey:    os.Getenv("ZAIM_CONSUMER_KEY"),
		ConsumerSecret: os.Getenv("ZAIM_CONSUMER_SECRET"),
		Token:          os.Getenv("ZAIM_TOKEN"),
		TokenSecret:    os.Getenv("ZAIM_TOKEN_SECRET"),
	}

	c := zaim.NewZaimClient(p)

	params := zaim.CreateTransferParams{
		Amount:        1,
		Date:          time.Now(),
		FromAccountID: 1,
		ToAccountID:   2,
	}

	// 振替情報を登録
	res, err := c.CreateTransfer(params)
	if err != nil {
		log.Fatal(err)
	}

	println(res.Money.ID)
	println(res.Money.Modified.String())
}

func ExampleClient_UpdateTransfer() {
	// クライアント初期化
	p := zaim.ZaimParams{
		ConsumerKey:    os.Getenv("ZAIM_CONSUMER_KEY"),
		ConsumerSecret: os.Getenv("ZAIM_CONSUMER_SECRET"),
		Token:          os.Getenv("ZAIM_TOKEN"),
		TokenSecret:    os.Getenv("ZAIM_TOKEN_SECRET"),
	}

	c := zaim.NewZaimClient(p)

	params := zaim.UpdateTransferParams{
		Amount:        1,
		Date:          time.Now(),
		FromAccountID: 1,
		ToAccountID:   2,
	}

	// 振替情報を更新
	res, err := c.UpdateTransfer(381, params)
	if err != nil {
		log.Fatal(err)
	}

	println(res.Money.ID)
	println(res.Money.Modified.String())
}

func TestCreateTransfer(t *testing.T) {
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

	now := time.Now()
	fromAccountID := uint64(1)
	toAccountID := uint64(2)
	comment := "test-comment"

	params := zaim.CreateTransferParams{
		Amount:        0,
		Date:          now,
		FromAccountID: fromAccountID,
		ToAccountID:   toAccountID,
		Comment:       &comment,
	}

	t.Run("正常系: 振替情報を正しく登録できる", func(t *testing.T) {
		// 正常なレスポンスを返すモックを設定
		err := testutil.MockResponseFromFile("POST", zaim.BaseURL+"home/money/transfer", "transfer")
		assert.NoError(t, err)

		// 振替情報を登録
		res, err := c.CreateTransfer(params)

		// レスポンスの確認
		assert.NoError(t, err)

		m := res.Money
		assert.Equal(t, int(m.ID), 11820767)
		assert.Equal(t, m.Modified, time.Date(2013, 7, 8, 21, 4, 54, 0, time.UTC))

		u := res.User
		assert.Equal(t, u.DataModified, time.Date(2013, 7, 8, 21, 4, 56, 0, time.UTC))
		assert.Equal(t, int(u.InputCount), 12)

	})

	t.Run("異常系: サーバーエラー時にエラーを返却する", func(t *testing.T) {
		// エラーを返すモックに差し替え
		httpmock.RegisterResponder("POST", zaim.BaseURL+"home/money/transfer",
			httpmock.NewStringResponder(500, `Internal Server Error`))

		// 振替情報を登録
		res, err := c.CreateTransfer(params)

		// レスポンスの確認
		assert.Error(t, err)
		assert.Empty(t, res)
	})
}

func TestUpdateTransfer(t *testing.T) {
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

	now := time.Now()
	fromAccountID := uint64(1)
	toAccountID := uint64(1)
	comment := "test-comment"

	params := zaim.UpdateTransferParams{
		Amount:        0,
		Date:          now,
		FromAccountID: fromAccountID,
		ToAccountID:   toAccountID,
		Comment:       &comment,
	}

	t.Run("正常系: 振替情報を正しく更新できる", func(t *testing.T) {
		// 正常なレスポンスを返すモックを設定
		err := testutil.MockResponseFromFile("PUT", zaim.BaseURL+"home/money/transfer/11820767", "transfer")
		assert.NoError(t, err)

		// 振替情報を更新
		res, err := c.UpdateTransfer(11820767, params)

		// レスポンスの確認
		assert.NoError(t, err)

		m := res.Money
		assert.Equal(t, int(m.ID), 11820767)
		assert.Equal(t, m.Modified, time.Date(2013, 7, 8, 21, 4, 54, 0, time.UTC))

		u := res.User
		assert.Equal(t, u.DataModified, time.Date(2013, 7, 8, 21, 4, 56, 0, time.UTC))
		assert.Equal(t, int(u.InputCount), 12)
	})

	t.Run("異常系: サーバーエラー時にエラーを返却する", func(t *testing.T) {
		// エラーを返すモックに差し替え
		httpmock.RegisterResponder("PUT", zaim.BaseURL+"home/money/transfer/11820767",
			httpmock.NewStringResponder(500, `Internal Server Error`))

		// 振替情報を更新
		res, err := c.UpdateTransfer(11820767, params)

		// レスポンスの確認
		assert.Error(t, err)
		assert.Empty(t, res)
	})
}

func TestDeleteTransfer(t *testing.T) {
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

	t.Run("正常系: 振替情報を正しく削除できる", func(t *testing.T) {
		// 正常なレスポンスを返すモックを設定
		err := testutil.MockResponseFromFile("DELETE", zaim.BaseURL+"home/money/transfer/11820767", "transfer")
		assert.NoError(t, err)

		// 振替情報を削除
		res, err := c.DeleteTransfer(11820767)

		// レスポンスの確認
		assert.NoError(t, err)

		m := res.Money
		assert.Equal(t, int(m.ID), 11820767)
		assert.Equal(t, m.Modified, time.Date(2013, 7, 8, 21, 4, 54, 0, time.UTC))

		u := res.User
		assert.Equal(t, u.DataModified, time.Date(2013, 7, 8, 21, 4, 56, 0, time.UTC))
		assert.Equal(t, int(u.InputCount), 12)

	})

	t.Run("異常系: サーバーエラー時にエラーを返却する", func(t *testing.T) {
		// エラーを返すモックに差し替え
		httpmock.RegisterResponder("DELETE", zaim.BaseURL+"home/money/transfer/11820767",
			httpmock.NewStringResponder(500, `Internal Server Error`))

		// 振替情報を削除
		res, err := c.DeleteTransfer(11820767)

		// レスポンスの確認
		assert.Error(t, err)
		assert.Empty(t, res)
	})
}
