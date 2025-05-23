package models

type Currency struct {
	CurrencyCode string `json:"currency_code"`
	Unit         string `json:"unit"`
	Name         string `json:"name"`
	Point        uint8  `json:"point"`
}
