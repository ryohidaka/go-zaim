package api

import (
	"encoding/json"

	"github.com/ryohidaka/go-zaim/models"
)

func ParseCurrencyResponse(body []byte) ([]models.Currency, error) {
	var raw struct {
		Currencies []models.Currency `json:"currencies"`
	}

	if err := json.Unmarshal(body, &raw); err != nil {
		return nil, err
	}

	return raw.Currencies, nil
}
