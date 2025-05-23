package zaim

import (
	"github.com/ryohidaka/go-zaim/internal/api"
	"github.com/ryohidaka/go-zaim/models"
)

// FetchCurrency は通貨一覧を取得する
func (c *Client) FetchCurrency() ([]models.Currency, error) {
	body, err := c.get("currency", nil, false)
	if err != nil {
		return nil, err
	}

	return api.ParseCurrencyResponse(body)
}
