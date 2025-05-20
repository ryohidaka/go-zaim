package api

import (
	"encoding/json"

	"github.com/ryohidaka/go-zaim/models"
)

func ParseMeResponse(body []byte) (models.Me, error) {
	var raw struct {
		Me struct {
			models.Me
			ProfileModified models.ZaimTime `json:"profile_modified"`
			Active          models.BoolInt  `json:"active"`
			Created         models.ZaimTime `json:"created"`
		} `json:"me"`
	}

	if err := json.Unmarshal(body, &raw); err != nil {
		return models.Me{}, err
	}

	me := raw.Me.Me
	me.ProfileModified = raw.Me.ProfileModified.Time
	me.Active = bool(raw.Me.Active)
	me.Created = raw.Me.Created.Time

	return me, nil
}
