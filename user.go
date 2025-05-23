package zaim

import (
	"github.com/ryohidaka/go-zaim/internal/api"
	"github.com/ryohidaka/go-zaim/models"
)

// FetchMe は認証ユーザーの情報を取得する
func (c *Client) FetchMe() (models.Me, error) {
	body, err := c.get("home/user/verify", nil, true)
	if err != nil {
		return models.Me{}, err
	}

	return api.ParseMeResponse(body)
}
