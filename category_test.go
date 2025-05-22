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

func ExampleClient_FetchCategories() {
	// クライアント初期化
	p := zaim.ZaimParams{
		ConsumerKey:    os.Getenv("ZAIM_CONSUMER_KEY"),
		ConsumerSecret: os.Getenv("ZAIM_CONSUMER_SECRET"),
		Token:          os.Getenv("ZAIM_TOKEN"),
		TokenSecret:    os.Getenv("ZAIM_TOKEN_SECRET"),
	}

	c := zaim.NewZaimClient(p)

	// カテゴリ一覧を取得する
	categories, err := c.FetchCategories()
	if err != nil {
		log.Fatal(err)
	}

	for _, c := range categories {
		println(c.Name)
	}
}

func TestFetchCategories(t *testing.T) {
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

	t.Run("正常系: カテゴリ一覧を取得できる", func(t *testing.T) {
		url := zaim.BaseURL + "home/category"
		err := testutil.MockResponseFromFile("GET", url, "category")
		assert.NoError(t, err)

		// カテゴリ一覧を取得する
		categories, err := c.FetchCategories()

		// レスポンスの確認
		assert.NoError(t, err)
		assert.Equal(t, len(categories), 2)

		expected := []models.Category{
			{
				ID:               12093,
				Name:             "Food",
				Mode:             models.Payment,
				Sort:             1,
				ParentCategoryID: 101,
				Active:           true,
				Modified:         time.Date(2013, 1, 1, 0, 0, 0, 0, time.UTC),
			},
			{
				ID:               12094,
				Name:             "Daily good",
				Mode:             models.Payment,
				Sort:             2,
				ParentCategoryID: 102,
				Active:           true,
				Modified:         time.Date(2013, 1, 1, 0, 0, 0, 0, time.UTC),
			},
		}

		for i, c := range categories {
			assert.Equal(t, expected[i].ID, c.ID, "ID が一致しません")
			assert.Equal(t, expected[i].Name, c.Name, "Name が一致しません")
			assert.Equal(t, expected[i].Mode, c.Mode, "Mode が一致しません")
			assert.Equal(t, expected[i].Sort, c.Sort, "Sort が一致しません")
			assert.Equal(t, expected[i].ParentCategoryID, c.ParentCategoryID, "ParentCategoryID が一致しません")
			assert.Equal(t, expected[i].Active, c.Active, "Active が一致しません")
			assert.Equal(t, expected[i].Modified, c.Modified, "Modified が一致しません")
		}
	})

	t.Run("異常系: サーバーエラー時にエラーを返却する", func(t *testing.T) {
		url := zaim.BaseURL + "home/category"
		httpmock.RegisterResponder("GET", url, httpmock.NewStringResponder(500, `Internal Server Error`))

		// カテゴリ一覧を取得する
		categories, err := c.FetchCategories()

		// レスポンスの確認
		assert.Error(t, err)
		assert.Empty(t, categories)
	})
}
