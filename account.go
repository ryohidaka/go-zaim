package zaim

import (
	"github.com/ryohidaka/go-zaim/internal/api"
	"github.com/ryohidaka/go-zaim/models"
)

// FetchAccounts は口座一覧を取得する
func (c *Client) FetchAccounts() ([]models.Account, error) {
	body, err := c.get("home/account", nil, true)
	if err != nil {
		return nil, err
	}

	return api.ParseAccountResponse(body)
}

// FetchDefaultAccounts はデフォルトの口座一覧を取得する
func (c *Client) FetchDefaultAccounts() ([]models.DefaultAccount, error) {
	body, err := c.get("account", nil, false)
	if err != nil {
		return nil, err
	}

	return api.ParseDefaultAccountResponse(body)
}
