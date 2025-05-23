package models

import "time"

type Category struct {
	ID               uint16    `json:"id"`
	Name             string    `json:"name"`
	Mode             Mode      `json:"mode"`
	Sort             uint8     `json:"sort"`
	ParentCategoryID uint16    `json:"parent_category_id"`
	LocalID          uint16    `json:"local_id"`
	Active           bool      `json:"active"`
	Modified         time.Time `json:"modified"`
}

type DefaultCategory struct {
	ID   uint16 `json:"id"`
	Mode Mode   `json:"mode"`
	Name string `json:"name"`
}
