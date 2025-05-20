package models

import "time"

type Account struct {
	ID              uint32    `json:"id"`
	Name            string    `json:"name"`
	Modified        time.Time `json:"modified"`
	Sort            uint8     `json:"sort"`
	Active          bool      `json:"active"`
	LocalID         uint32    `json:"local_id"`
	WebsiteID       uint32    `json:"website_id"`
	ParentAccountID uint32    `json:"parent_account_id"`
}
