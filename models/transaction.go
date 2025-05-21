package models

import "time"

type Transaction struct {
	Money   TransactionMoney `json:"money"`
	User    TransactionUser  `json:"user"`
	Banners *[]banner        `json:"banners"`
	Stamps  *stamps          `json:"stamps"`
	Place   *Place           `json:"place"`
}

type TransactionMoney struct {
	ID       uint64    `json:"id"`
	Modified time.Time `json:"modified"`
}

type TransactionUser struct {
	DataModified time.Time `json:"data_modified"`
	InputCount   uint16    `json:"input_count"`
}

type Place struct {
	UserID       uint64    `json:"user_id"`
	PlaceUID     string    `json:"place_uid"`
	Service      string    `json:"service"`
	Name         string    `json:"test"`
	Mode         Mode      `json:"mode"`
	OriginalName string    `json:"original_name"`
	Modified     time.Time `json:"modified"`
	Created      time.Time `json:"created"`
	Count        uint16    `json:"count"`
	Active       bool      `json:"active"`
	CategoryID   uint16    `json:"category_id"`
	GenreID      uint16    `json:"genre_id"`
	ID           uint64    `json:"id"`
}

// 構造が不明
type banner = any
type stamps = any
