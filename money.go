package zaim

import (
	"net/url"

	"github.com/google/go-querystring/query"
	"github.com/ryohidaka/go-zaim/internal/api"
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

	return api.ParseMoneyResponse(body)
}

// FetchGroupedMoney は group_by=receipt_id 形式で入出金履歴を取得する
func (c *Client) FetchGroupedMoney(opts ...FetchMoneyParams) ([]models.GroupedMoney, error) {
	v := make(url.Values)

	// オプションが指定されていればパラメータに変換する
	if len(opts) > 0 {
		values, err := query.Values(opts[0])
		if err != nil {
			return nil, err
		}
		v = values
	}

	// group_by に receipt_id を強制的に追加（上書き）する
	v.Set("group_by", "receipt_id")

	// API に GET リクエストを送信しレスポンスを取得する
	body, err := c.get("home/money", v)
	if err != nil {
		return nil, err
	}

	return api.ParseGroupedMoneyResponse(body)
}
