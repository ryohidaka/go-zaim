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

	p := zaim.ZaimParams{
		ConsumerKey:    "dummy-key",
		ConsumerSecret: "dummy-secret",
		Token:          "dummy-token",
		TokenSecret:    "dummy-secret",
	}

	c := zaim.NewZaimClient(p)

	t.Run("正常系: ジャンル一覧を取得できる", func(t *testing.T) {
		url := zaim.BaseURL + "home/genre"
		err := testutil.MockResponseFromFile(url, "genre")
		assert.NoError(t, err)

		// ジャンル一覧を取得する
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
			assert.Equal(t, expected[i].ID, g.ID, "ID が一致しません")
			assert.Equal(t, expected[i].Name, g.Name, "Name が一致しません")
			assert.Equal(t, expected[i].Sort, g.Sort, "Sort が一致しません")
			assert.Equal(t, expected[i].Active, g.Active, "Active が一致しません")
			assert.Equal(t, expected[i].CategoryID, g.CategoryID, "CategoryID が一致しません")
			assert.Equal(t, expected[i].ParentGenreID, g.ParentGenreID, "ParentGenreID が一致しません")
			assert.Equal(t, expected[i].Modified, g.Modified, "Modified が一致しません")
		}
	})

	t.Run("異常系: サーバーエラー時にエラーを返却する", func(t *testing.T) {
		url := zaim.BaseURL + "home/genre"
		httpmock.RegisterResponder("GET", url, httpmock.NewStringResponder(500, `Internal Server Error`))

		// ジャンル一覧を取得する
		genres, err := c.FetchGenres()

		// レスポンスの確認
		assert.Error(t, err)
		assert.Empty(t, genres)
	})
}
