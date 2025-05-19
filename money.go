package zaim

import (
	"encoding/json"
	"net/url"

	"github.com/google/go-querystring/query"
	"github.com/ryohidaka/go-zaim/models"
)

type FetchMoneyParams struct {
	CategoryID uint16       `url:"category_id,omitempty"`
	GenreID    uint16       `url:"genre_id,omitempty"`
	Mode       models.Mode  `url:"mode,omitempty"`       // payment, income, transfer
	Order      models.Order `url:"order,omitempty"`      // id or date (default: date)
	StartDate  string       `url:"start_date,omitempty"` // YYYY-MM-DD
	EndDate    string       `url:"end_date,omitempty"`   // YYYY-MM-DD
	Page       uint8        `url:"page,omitempty"`       // default 1
	Limit      uint8        `url:"limit,omitempty"`      // default 20, max 100
}

// FetchMoney はユーザーの入出金履歴を取得する
func (c *Client) FetchMoney(opts ...FetchMoneyParams) ([]models.Money, error) {
	var v url.Values

	if len(opts) > 0 {
		values, err := query.Values(opts[0])
		if err != nil {
			return nil, err
		}
		v = values
	}

	body, err := c.get("home/money", v)
	if err != nil {
		return nil, err
	}

	var raw struct {
		Money []struct {
			models.Money
			Active  models.BoolInt  `json:"active"`
			Date    models.ZaimTime `json:"date"`
			Created models.ZaimTime `json:"created"`
		} `json:"money"`
	}

	if err := json.Unmarshal(body, &raw); err != nil {
		return nil, err
	}

	money := make([]models.Money, len(raw.Money))
	for i, m := range raw.Money {
		money[i] = m.Money
		money[i].Active = bool(m.Active)
		money[i].Date = m.Date.Time
		money[i].Created = m.Created.Time
	}

	return money, nil
}
