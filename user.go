package zaim

import (
	"encoding/json"

	"github.com/ryohidaka/go-zaim/models"
)

// FetchMe は認証ユーザーの情報を取得する
func (c *Client) FetchMe() (models.Me, error) {
	body, err := c.get("home/user/verify", nil)
	if err != nil {
		return models.Me{}, err
	}

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
