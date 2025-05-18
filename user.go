package zaim

import (
	"encoding/json"

	"github.com/ryohidaka/go-zaim/models"
)

// FetchMe は認証ユーザーの情報を取得する
func (c *Client) FetchMe() (models.Me, error) {
	body, err := c.get("home/user/verify", nil)
	if err != nil {
		return models.Me{}, err
	}

	var resp models.MeResponse
	if err := json.Unmarshal(body, &resp); err != nil {
		return models.Me{}, err
	}

	return resp.Me, nil
}
