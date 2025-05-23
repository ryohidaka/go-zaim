package zaim_test

import (
	"testing"
	"time"

	"github.com/jarcoal/httpmock"
	"github.com/ryohidaka/go-zaim"
	"github.com/ryohidaka/go-zaim/testutil"
	"github.com/stretchr/testify/assert"
)

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
