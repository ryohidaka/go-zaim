package zaim

import (
	"github.com/ryohidaka/go-zaim/internal/api"
	"github.com/ryohidaka/go-zaim/models"
)

// FetchGenres はジャンル一覧を取得する
func (c *Client) FetchGenres() ([]models.Genre, error) {
	body, err := c.get("home/genre", nil, true)
	if err != nil {
		return nil, err
	}

	return api.ParseGenreResponse(body)
}
