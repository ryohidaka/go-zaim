package zaim

import (
	"github.com/ryohidaka/go-zaim/internal/api"
	"github.com/ryohidaka/go-zaim/models"
)

// FetchCategories はカテゴリ一覧を取得する
func (c *Client) FetchCategories() ([]models.Category, error) {
	body, err := c.get("home/category", nil, true)
	if err != nil {
		return nil, err
	}

	return api.ParseCategoryResponse(body)
}

// FetchDefaultCategories はデフォルトのカテゴリ一覧を取得する
func (c *Client) FetchDefaultCategories() ([]models.DefaultCategory, error) {
	body, err := c.get("category", nil, false)
	if err != nil {
		return nil, err
	}

	return api.ParseDefaultCategoryResponse(body)
}
