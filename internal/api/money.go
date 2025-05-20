package api

import (
	"encoding/json"

	"github.com/ryohidaka/go-zaim/models"
)

func ParseMoneyResponse(body []byte) ([]models.Money, error) {
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

func ParseGroupedMoneyResponse(body []byte) ([]models.GroupedMoney, error) {
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
