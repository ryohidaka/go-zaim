package api

import (
	"encoding/json"

	"github.com/ryohidaka/go-zaim/models"
)

func ParseGenreResponse(body []byte) ([]models.Genre, error) {
	var raw struct {
		Genres []struct {
			models.Genre
			Active   models.BoolInt  `json:"active"`
			Modified models.ZaimTime `json:"modified"`
		} `json:"genres"`
	}

	if err := json.Unmarshal(body, &raw); err != nil {
		return nil, err
	}

	genres := make([]models.Genre, len(raw.Genres))
	for i, g := range raw.Genres {
		genres[i] = g.Genre
		genres[i].Active = bool(g.Active)
		genres[i].Modified = g.Modified.Time
	}

	return genres, nil
}

func ParseDefaultGenreResponse(body []byte) ([]models.DefaultGenre, error) {
	var raw struct {
		Genres []models.DefaultGenre `json:"genres"`
	}

	if err := json.Unmarshal(body, &raw); err != nil {
		return nil, err
	}

	return raw.Genres, nil
}
