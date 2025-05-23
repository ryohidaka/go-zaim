package models

import "time"

type Genre struct {
	ID            uint16    `json:"id"`
	CategoryID    uint16    `json:"category_id"`
	Name          string    `json:"name"`
	Sort          uint8     `json:"sort"`
	Active        bool      `json:"active"`
	Modified      time.Time `json:"modified"`
	ParentGenreID uint16    `json:"parent_genre_id"`
	LocalID       uint16    `json:"local_id"`
}

type DefaultGenre struct {
	ID         uint16 `json:"id"`
	CategoryID uint16 `json:"category_id"`
	Name       string `json:"name"`
}
