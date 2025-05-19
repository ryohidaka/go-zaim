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

func TestFetchMoney(t *testing.T) {
	// モックのHTTPサーバーを有効化
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	// モックレスポンスを設定
	url := zaim.BaseURL + "home/money?category_id=101&end_date=2024-01-31&genre_id=5&limit=50&mapping=1&mode=payment&order=id&page=2&start_date=2024-01-01"
	err := testutil.MockResponseFromFile(url, "money")
	assert.NoError(t, err)

	p := zaim.ZaimParams{
		ConsumerKey:    "dummy-key",
		ConsumerSecret: "dummy-secret",
		Token:          "dummy-token",
		TokenSecret:    "dummy-secret",
	}

	// 入出金履歴を取得
	c := zaim.NewZaimClient(p)

	params := zaim.FetchMoneyParams{
		CategoryID: 101,
		GenreID:    5,
		Mode:       models.Payment,
		Order:      models.ID,
		StartDate:  "2024-01-01",
		EndDate:    "2024-01-31",
		Page:       2,
		Limit:      50,
	}

	money, err := c.FetchMoney(params)

	// レスポンスの確認
	assert.NoError(t, err)
	assert.Equal(t, len(money), 2)

	expected := []models.Money{
		{
			ID:      381,
			Mode:    models.Income,
			Active:  true,
			Date:    time.Date(2011, 11, 7, 0, 0, 0, 0, time.UTC),
			Created: time.Date(2011, 11, 7, 1, 10, 50, 0, time.UTC),
		},
		{
			ID:      382,
			Mode:    models.Payment,
			Active:  true,
			Date:    time.Date(2011, 11, 7, 0, 0, 0, 0, time.UTC),
			Created: time.Date(2011, 11, 7, 1, 12, 0, 0, time.UTC),
		},
	}

	for i, m := range money {
		println(m.ID)
		assert.Equal(t, expected[i].ID, m.ID, "IDが一致しません")
		assert.Equal(t, expected[i].Mode, m.Mode, "Modeが一致しません")
		assert.Equal(t, expected[i].Active, m.Active, "Activeが一致しません")
		assert.Equal(t, expected[i].Date, m.Date, "Dateが一致しません")
		assert.Equal(t, expected[i].Created, m.Created, "Createdが一致しません")
	}
}
