package models

type MeResponse struct {
	Me Me `json:"me"`

	Requested
}

type Me struct {
	ID              uint64   `json:"id"`                // unique user id
	Login           string   `json:"login"`             // unique string for user login
	Name            string   `json:"name"`              // user name
	InputCount      uint16   `json:"input_count"`       // total number of inputs
	DayCount        uint16   `json:"day_count"`         // total number of days
	RepeatCount     uint16   `json:"repeat_count"`      // days continuous recording
	Day             uint8    `json:"day"`               // start date of the month
	Week            uint8    `json:"week"`              // first day of the week
	Month           uint8    `json:"month"`             // start date of the year
	CurrencyCode    string   `json:"currency_code"`     // default currency code
	ProfileImageURL string   `json:"profile_image_url"` // profile image url
	CoverImageURL   string   `json:"cover_image_url"`   // cover image url
	ProfileModified ZaimTime `json:"profile_modified"`  // modified
	Active          BoolInt  `json:"active"`            // active
	Created         ZaimTime `json:"created"`           // created
}
