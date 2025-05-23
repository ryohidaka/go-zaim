package zaim_test

import (
	"testing"

	"github.com/jarcoal/httpmock"
	"github.com/ryohidaka/go-zaim"
	"github.com/ryohidaka/go-zaim/models"
	"github.com/ryohidaka/go-zaim/testutil"
	"github.com/stretchr/testify/assert"
)

func TestFetchCurrency(t *testing.T) {
	// モックのHTTPサーバーを有効化
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	c := zaim.NewZaimClient()

	t.Run("正常系: 通貨一覧を取得できる", func(t *testing.T) {
		url := zaim.BaseURL + "currency"
		err := testutil.MockResponseFromFile("GET", url, "currency")
		assert.NoError(t, err)

		// 通貨一覧を取得する
		currencies, err := c.FetchCurrency()

		// レスポンスの確認
		assert.NoError(t, err)
		assert.Equal(t, len(currencies), 2)

		expected := []models.Currency{
			{
				CurrencyCode: "AUD",
				Unit:         "$",
				Name:         "Australian dollar",
				Point:        2,
			},
			{
				CurrencyCode: "JPY",
				Unit:         "￥",
				Name:         "Japanese YEN",
				Point:        0,
			},
		}

		for i, c := range currencies {
			assert.Equal(t, expected[i].CurrencyCode, c.CurrencyCode, "CurrencyCode が一致しません")
			assert.Equal(t, expected[i].Unit, c.Unit, "Unit が一致しません")
			assert.Equal(t, expected[i].Name, c.Name, "Name が一致しません")
			assert.Equal(t, expected[i].Point, c.Point, "Point が一致しません")
		}
	})

	t.Run("異常系: サーバーエラー時にエラーを返却する", func(t *testing.T) {
		url := zaim.BaseURL + "currency"
		httpmock.RegisterResponder("GET", url, httpmock.NewStringResponder(500, `Internal Server Error`))

		// 通貨一覧を取得する
		currencies, err := c.FetchCurrency()

		// レスポンスの確認
		assert.Error(t, err)
		assert.Empty(t, currencies)
	})
}
