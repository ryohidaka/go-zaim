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

	// モックレスポンスを設定
	url := zaim.BaseURL + "home/category"
	err := testutil.MockResponseFromFile(url, "category")
	assert.NoError(t, err)

	p := zaim.ZaimParams{
		ConsumerKey:    "dummy-key",
		ConsumerSecret: "dummy-secret",
		Token:          "dummy-token",
		TokenSecret:    "dummy-secret",
	}

	// カテゴリ一覧を取得する
	c := zaim.NewZaimClient(p)
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
		assert.Equal(t, expected[i].ID, c.ID, "IDが一致しません")
		assert.Equal(t, expected[i].Name, c.Name, "Nameが一致しません")
		assert.Equal(t, expected[i].Mode, c.Mode, "Modeが一致しません")
		assert.Equal(t, expected[i].Sort, c.Sort, "Sortが一致しません")
		assert.Equal(t, expected[i].ParentCategoryID, c.ParentCategoryID, "ParentCategoryIDが一致しません")
		assert.Equal(t, expected[i].Active, c.Active, "Activeが一致しません")
		assert.Equal(t, expected[i].Modified, c.Modified, "Modifiedが一致しません")
	}
}
