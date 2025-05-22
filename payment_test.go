package zaim_test

import (
	"testing"
	"time"

	"github.com/jarcoal/httpmock"
	"github.com/ryohidaka/go-zaim"
	"github.com/ryohidaka/go-zaim/models"
	"github.com/ryohidaka/go-zaim/testutil"
	"github.com/stretchr/testify/assert"
)

func TestCreatePayment(t *testing.T) {
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
	comment := "test-comment"
	name := "test-name"
	place := "test-place"

	params := zaim.CreatePaymentParams{
		CategoryID:    7,
		GenreID:       10101,
		Amount:        0,
		Date:          now,
		FromAccountID: &fromAccountID,
		Comment:       &comment,
		Name:          &name,
		Place:         &place,
	}

	t.Run("正常系: 支払情報を正しく登録できる", func(t *testing.T) {
		// 正常なレスポンスを返すモックを設定
		err := testutil.MockResponseFromFile("POST", zaim.BaseURL+"home/money/payment", "transaction")
		assert.NoError(t, err)

		// 支払情報を登録
		res, err := c.CreatePayment(params)

		// レスポンスの確認
		assert.NoError(t, err)

		m := res.Money
		assert.Equal(t, int(m.ID), 11820767)
		assert.Equal(t, m.Modified, time.Date(2013, 7, 8, 21, 4, 54, 0, time.UTC))

		u := res.User
		assert.Equal(t, u.DataModified, time.Date(2013, 7, 8, 21, 4, 56, 0, time.UTC))
		assert.Equal(t, int(u.InputCount), 12)

		p := res.Place
		assert.Equal(t, p.Mode, models.Payment)
		assert.Equal(t, p.CategoryID, params.CategoryID)
		assert.Equal(t, p.GenreID, params.GenreID)
	})

	t.Run("異常系: サーバーエラー時にエラーを返却する", func(t *testing.T) {
		// エラーを返すモックに差し替え
		httpmock.RegisterResponder("POST", zaim.BaseURL+"home/money/payment",
			httpmock.NewStringResponder(500, `Internal Server Error`))

		// 支払情報を登録
		res, err := c.CreatePayment(params)

		// レスポンスの確認
		assert.Error(t, err)
		assert.Empty(t, res)
	})
}
