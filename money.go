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

	// API レスポンスをパースするための一時構造体を定義する
	var raw struct {
		Money []struct {
			models.GroupedMoney
			Date models.ZaimTime `json:"date"` // 日付
			Data []struct {
				models.GroupedMoneyData
				Active  models.BoolInt  `json:"active"`  // 有効フラグ
				Created models.ZaimTime `json:"created"` // 作成日時
			} `json:"data"`
		} `json:"money"`
	}

	// レスポンスを JSON から構造体に変換する
	if err := json.Unmarshal(body, &raw); err != nil {
		return nil, err
	}

	// 最終的な結果を作成する
	money := make([]models.GroupedMoney, len(raw.Money))

	for i, m := range raw.Money {
		// 各レシートのデータを処理する
		items := make([]models.GroupedMoneyData, len(m.Data))
		for j, d := range m.Data {
			items[j] = models.GroupedMoneyData{
				ID:         d.ID,
				CategoryID: d.CategoryID,
				GenreID:    d.GenreID,
				Amount:     d.Amount,
				Comment:    d.Comment,
				Active:     bool(d.Active),
				Created:    d.Created.Time,
				Name:       d.Name,
				ReceiptID:  d.ReceiptID,
			}
		}

		// グループ化された入出金データを格納する
		money[i] = models.GroupedMoney{
			Amount:        m.Amount,
			ToAccountID:   m.ToAccountID,
			FromAccountID: m.FromAccountID,
			Date:          m.Date.Time,
			ReceiptID:     m.ReceiptID,
			Mode:          m.Mode,
			PlaceUID:      m.PlaceUID,
			CategoryID:    m.CategoryID,
			GenreID:       m.GenreID,
			CurrencyCode:  m.CurrencyCode,
			Place:         m.Place,
			Data:          items,
		}
	}

	return money, nil
}
