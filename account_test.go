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

func TestFetchAccounts(t *testing.T) {
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

	t.Run("正常系: 口座一覧を取得できる", func(t *testing.T) {
		url := zaim.BaseURL + "home/account"
		err := testutil.MockResponseFromFile(url, "account")
		assert.NoError(t, err)

		// 口座一覧を取得する
		accounts, err := c.FetchAccounts()

		// レスポンスの確認
		assert.NoError(t, err)
		assert.Equal(t, len(accounts), 2)

		expected := []models.Account{
			{
				ID:              15497739,
				Name:            "Credit card",
				Modified:        time.Date(2022, 3, 15, 13, 39, 52, 0, time.UTC),
				Sort:            8,
				Active:          true,
				LocalID:         15497739,
				WebsiteID:       0,
				ParentAccountID: 0,
			},
			{
				ID:              16324163,
				Name:            "Wallet",
				Modified:        time.Date(2022, 11, 28, 15, 48, 5, 0, time.UTC),
				Sort:            9,
				Active:          true,
				LocalID:         16324163,
				WebsiteID:       0,
				ParentAccountID: 0,
			},
		}

		for i, a := range accounts {
			assert.Equal(t, expected[i].ID, a.ID, "ID が一致しません")
			assert.Equal(t, expected[i].Name, a.Name, "Name が一致しません")
			assert.Equal(t, expected[i].Modified, a.Modified, "Modified が一致しません")
			assert.Equal(t, expected[i].Sort, a.Sort, "Sort が一致しません")
			assert.Equal(t, expected[i].Active, a.Active, "Active が一致しません")
			assert.Equal(t, expected[i].LocalID, a.LocalID, "LocalID が一致しません")
			assert.Equal(t, expected[i].WebsiteID, a.WebsiteID, "WebsiteID が一致しません")
			assert.Equal(t, expected[i].ParentAccountID, a.ParentAccountID, "ParentAccountID が一致しません")
		}
	})

	t.Run("異常系: サーバーエラー時にエラーを返却する", func(t *testing.T) {
		url := zaim.BaseURL + "home/account"
		httpmock.RegisterResponder("GET", url, httpmock.NewStringResponder(500, `Internal Server Error`))

		// 口座一覧を取得する
		accounts, err := c.FetchAccounts()

		// レスポンスの確認
		assert.Error(t, err)
		assert.Empty(t, accounts)
	})
}
