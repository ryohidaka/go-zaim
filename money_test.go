package zaim_test

import (
	"log"
	"os"
	"testing"
	"time"

	"github.com/jarcoal/httpmock"
	"github.com/ryohidaka/go-zaim"
	"github.com/ryohidaka/go-zaim/models"
	"github.com/ryohidaka/go-zaim/testutil"
	"github.com/stretchr/testify/assert"
)

func ExampleClient_FetchMoney() {
	// クライアント初期化
	p := zaim.ZaimParams{
		ConsumerKey:    os.Getenv("ZAIM_CONSUMER_KEY"),
		ConsumerSecret: os.Getenv("ZAIM_CONSUMER_SECRET"),
		Token:          os.Getenv("ZAIM_TOKEN"),
		TokenSecret:    os.Getenv("ZAIM_TOKEN_SECRET"),
	}

	c := zaim.NewZaimClient(p)

	// 入出金履歴を取得
	money, err := c.FetchMoney()
	if err != nil {
		log.Fatal(err)
	}

	for _, m := range money {
		println(m.Name)
	}
}

func ExampleClient_FetchGroupedMoney() {
	// クライアント初期化
	p := zaim.ZaimParams{
		ConsumerKey:    os.Getenv("ZAIM_CONSUMER_KEY"),
		ConsumerSecret: os.Getenv("ZAIM_CONSUMER_SECRET"),
		Token:          os.Getenv("ZAIM_TOKEN"),
		TokenSecret:    os.Getenv("ZAIM_TOKEN_SECRET"),
	}

	c := zaim.NewZaimClient(p)

	// 入出金履歴を取得
	money, err := c.FetchGroupedMoney()
	if err != nil {
		log.Fatal(err)
	}

	for _, m := range money {
		println(m.ReceiptID)
	}
}

func TestFetchMoney(t *testing.T) {
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

	t.Run("正常系: 入出金履歴を取得できる", func(t *testing.T) {
		// 正常レスポンスを登録
		url := zaim.BaseURL + "home/money?category_id=101&end_date=2024-01-31&genre_id=5&limit=50&mapping=1&mode=payment&order=id&page=2&start_date=2024-01-01"
		err := testutil.MockResponseFromFile("GET", url, "money")
		assert.NoError(t, err)

		// 入出金履歴を取得
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
			assert.Equal(t, expected[i].ID, m.ID, "ID が一致しません")
			assert.Equal(t, expected[i].Mode, m.Mode, "Mode が一致しません")
			assert.Equal(t, expected[i].Active, m.Active, "Active が一致しません")
			assert.Equal(t, expected[i].Date, m.Date, "Date が一致しません")
			assert.Equal(t, expected[i].Created, m.Created, "Created が一致しません")
		}
	})

	t.Run("異常系: サーバーエラー時にエラーを返却する", func(t *testing.T) {
		// エラーを返すモックを設定
		url := zaim.BaseURL + "home/money"
		httpmock.RegisterResponder("GET", url, httpmock.NewStringResponder(500, `Internal Server Error`))

		money, err := c.FetchMoney()

		assert.Error(t, err)
		assert.Empty(t, money)
	})
}

func TestFetchGroupedMoney(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	p := zaim.ZaimParams{
		ConsumerKey:    "dummy-key",
		ConsumerSecret: "dummy-secret",
		Token:          "dummy-token",
		TokenSecret:    "dummy-secret",
	}
	c := zaim.NewZaimClient(p)

	t.Run("正常系: グループ化された入出金履歴を取得できる", func(t *testing.T) {
		url := zaim.BaseURL + "home/money?group_by=receipt_id&mapping=1"
		err := testutil.MockResponseFromFile("GET", url, "money-grouped")
		assert.NoError(t, err)

		// 入出金履歴を取得
		money, err := c.FetchGroupedMoney()

		// レスポンスの確認
		assert.NoError(t, err)

		expected := []models.GroupedMoney{
			{
				Amount:        1800,
				FromAccountID: 34555,
				Date:          time.Date(2011, 11, 7, 0, 0, 0, 0, time.UTC),
				Data: []models.GroupedMoneyData{
					{
						ID:      381,
						Created: time.Date(2011, 11, 7, 1, 10, 50, 0, time.UTC),
						Active:  true,
					},
					{
						ID:      382,
						Created: time.Date(2011, 11, 7, 1, 12, 0, 0, time.UTC),
						Active:  true,
					},
				},
			},
		}

		assert.Equal(t, expected[0].Amount, money[0].Amount, "Amount が一致しません")
		assert.Equal(t, expected[0].FromAccountID, money[0].FromAccountID, "FromAccountID が一致しません")
		assert.Equal(t, expected[0].Date, money[0].Date, "Date が一致しません")

		for i, m := range money[0].Data {
			assert.Equal(t, expected[0].Data[i].ID, m.ID, "ID が一致しません")
			assert.Equal(t, expected[0].Data[i].Created, m.Created, "Created が一致しません")
			assert.Equal(t, expected[0].Data[i].Active, m.Active, "Active が一致しません")
		}
	})

	t.Run("異常系: サーバーエラー時にエラーを返却する", func(t *testing.T) {
		url := zaim.BaseURL + "home/money?group_by=receipt_id&mapping=1"
		httpmock.RegisterResponder("GET", url, httpmock.NewStringResponder(500, `Internal Server Error`))

		money, err := c.FetchGroupedMoney()

		assert.Error(t, err)
		assert.Empty(t, money)
	})
}
