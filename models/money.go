package models

import "time"

type Money struct {
	ID            uint64    `json:"id"`   // unique input id
	Mode          Mode      `json:"mode"` // income or payment or transfer
	UserID        uint64    `json:"user_id"`
	Date          time.Time `json:"date"`
	CategoryID    uint16    `json:"category_id"`
	GenreID       uint16    `json:"genre_id"`
	ToAccountID   uint64    `json:"to_account_id"`
	FromAccountID uint64    `json:"from_account_id"`
	Amount        int32     `json:"amount"`
	Comment       string    `json:"comment"`
	Active        bool      `json:"active"`
	Name          string    `json:"name"`
	ReceiptID     uint64    `json:"receipt_id"`
	Place         string    `json:"place"`
	PlaceUID      string    `json:"place_uid"`
	Created       time.Time `json:"created"`
	CurrencyCode  string    `json:"currency_code"`
}

type GroupedMoney struct {
	Amount        int32              `json:"amount"`
	ToAccountID   uint64             `json:"to_account_id"`
	FromAccountID uint64             `json:"from_account_id"`
	Date          time.Time          `json:"date"`
	ReceiptID     uint64             `json:"receipt_id"`
	Mode          Mode               `json:"mode"`
	PlaceUID      string             `json:"place_uid"`
	CategoryID    uint16             `json:"category_id"`
	GenreID       uint16             `json:"genre_id"`
	CurrencyCode  string             `json:"currency_code"`
	Place         string             `json:"place"`
	Data          []GroupedMoneyData `json:"data"`
}

type GroupedMoneyData struct {
	ID         uint64    `json:"id"`
	CategoryID uint16    `json:"category_id"`
	GenreID    uint16    `json:"genre_id"`
	Amount     int32     `json:"amount"`
	Comment    string    `json:"comment"`
	Active     bool      `json:"active"`
	Created    time.Time `json:"created"`
	Name       string    `json:"name"`
	ReceiptID  uint64    `json:"receipt_id"`
}

type Mode string

const (
	Income   Mode = "income"
	Payment  Mode = "payment"
	Transfer Mode = "transfer"
)

type Order string

const (
	Date Order = "date"
	ID   Order = "id"
)
