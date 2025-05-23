package api

import (
	"encoding/json"

	"github.com/ryohidaka/go-zaim/models"
)

func ParseCategoryResponse(body []byte) ([]models.Category, error) {
	var raw struct {
		Categories []struct {
			models.Category
			Active   models.BoolInt  `json:"active"`
			Modified models.ZaimTime `json:"modified"`
		} `json:"categories"`
	}

	if err := json.Unmarshal(body, &raw); err != nil {
		return nil, err
	}

	categories := make([]models.Category, len(raw.Categories))
	for i, c := range raw.Categories {
		categories[i] = c.Category
		categories[i].Active = bool(c.Active)
		categories[i].Modified = c.Modified.Time
	}

	return categories, nil
}

func ParseDefaultCategoryResponse(body []byte) ([]models.DefaultCategory, error) {
	var raw struct {
		Categories []models.DefaultCategory `json:"categories"`
	}

	if err := json.Unmarshal(body, &raw); err != nil {
		return nil, err
	}

	return raw.Categories, nil
}
