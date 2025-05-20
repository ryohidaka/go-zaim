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

func ExampleClient_FetchGenres() {
	// クライアント初期化
	p := zaim.ZaimParams{
		ConsumerKey:    os.Getenv("ZAIM_CONSUMER_KEY"),
		ConsumerSecret: os.Getenv("ZAIM_CONSUMER_SECRET"),
		Token:          os.Getenv("ZAIM_TOKEN"),
		TokenSecret:    os.Getenv("ZAIM_TOKEN_SECRET"),
	}

	c := zaim.NewZaimClient(p)

	// ジャンル一覧を取得する
	genres, err := c.FetchGenres()
	if err != nil {
		log.Fatal(err)
	}

	for _, g := range genres {
		println(g.Name)
	}
}

func TestFetchGenres(t *testing.T) {
	// モックのHTTPサーバーを有効化
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	// モックレスポンスを設定
	url := zaim.BaseURL + "home/genre"
	err := testutil.MockResponseFromFile(url, "genre")
	assert.NoError(t, err)

	p := zaim.ZaimParams{
		ConsumerKey:    "dummy-key",
		ConsumerSecret: "dummy-secret",
		Token:          "dummy-token",
		TokenSecret:    "dummy-secret",
	}

	// ジャンル一覧を取得する
	c := zaim.NewZaimClient(p)
	genres, err := c.FetchGenres()

	// レスポンスの確認
	assert.NoError(t, err)
	assert.Equal(t, len(genres), 2)

	expected := []models.Genre{
		{
			ID:            12093,
			Name:          "Geocery",
			Sort:          1,
			Active:        true,
			CategoryID:    101,
			ParentGenreID: 10101,
			Modified:      time.Date(2013, 1, 1, 0, 0, 0, 0, time.UTC),
		},
		{
			ID:            12094,
			Name:          "Tabacco",
			Sort:          1,
			Active:        true,
			CategoryID:    102,
			ParentGenreID: 10201,
			Modified:      time.Date(2013, 1, 1, 0, 0, 0, 0, time.UTC),
		},
	}

	for i, g := range genres {
		assert.Equal(t, expected[i].ID, g.ID, "IDが一致しません")
		assert.Equal(t, expected[i].Name, g.Name, "Nameが一致しません")
		assert.Equal(t, expected[i].Sort, g.Sort, "Sortが一致しません")
		assert.Equal(t, expected[i].Active, g.Active, "Activeが一致しません")
		assert.Equal(t, expected[i].CategoryID, g.CategoryID, "CategoryIDが一致しません")
		assert.Equal(t, expected[i].ParentGenreID, g.ParentGenreID, "ParentGenreIDが一致しません")
		assert.Equal(t, expected[i].Modified, g.Modified, "Modifiedが一致しません")
	}
}
