package api

import (
	"encoding/json"

	"github.com/ryohidaka/go-zaim/models"
)

func ParseAccountResponse(body []byte) ([]models.Account, error) {
	var raw struct {
		Accounts []struct {
			models.Account
			Active   models.BoolInt  `json:"active"`
			Modified models.ZaimTime `json:"modified"`
		} `json:"accounts"`
	}

	if err := json.Unmarshal(body, &raw); err != nil {
		return nil, err
	}

	accounts := make([]models.Account, len(raw.Accounts))
	for i, a := range raw.Accounts {
		accounts[i] = a.Account
		accounts[i].Active = bool(a.Active)
		accounts[i].Modified = a.Modified.Time
	}

	return accounts, nil
}
