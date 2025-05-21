package api

import (
	"encoding/json"
	"time"

	"github.com/ryohidaka/go-zaim/models"
)

func ParseTransactionResponse(body []byte) (models.Transaction, error) {
	var raw struct {
		models.Transaction
		Money struct {
			models.TransactionMoney
			Modified models.ZaimTime `json:"modified"`
		} `json:"money"`
		User struct {
			models.TransactionUser
			DataModified models.ZaimTime `json:"data_modified"`
		} `json:"user"`
		Place struct {
			models.Place
			Modified models.ZaimTime `json:"modified"`
			Created  models.ZaimTime `json:"created"`
			Active   models.BoolInt  `json:"active"`
		} `json:"place"`
	}

	if err := json.Unmarshal(body, &raw); err != nil {
		return models.Transaction{}, err
	}

	t := raw.Transaction

	// Money フィールドの手動マッピング
	t.Money.ID = raw.Money.ID
	t.Money.Modified = raw.Money.Modified.Time

	// User フィールドの手動マッピング
	t.User.InputCount = raw.User.InputCount
	t.User.DataModified = raw.User.DataModified.Time

	// Place フィールドの手動マッピング
	t.Place = &raw.Place.Place
	t.Place.UserID = raw.Place.UserID
	t.Place.PlaceUID = raw.Place.PlaceUID
	t.Place.Service = raw.Place.Service
	t.Place.Name = raw.Place.Name
	t.Place.Mode = raw.Place.Mode
	t.Place.OriginalName = raw.Place.OriginalName
	t.Place.Modified = raw.Place.Modified.Time
	t.Place.Created = raw.Place.Created.Time
	t.Place.Count = raw.Place.Count
	t.Place.Active = bool(raw.Place.Active)
	t.Place.CategoryID = raw.Place.CategoryID
	t.Place.GenreID = raw.Place.GenreID
	t.Place.ID = raw.Place.ID

	return t, nil
}

// レシートID (UNIX エポック秒) を取得する
func GetReceiptID() *int64 {
	now := time.Now().Unix()

	return &now
}
