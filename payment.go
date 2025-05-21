package zaim

import (
	"time"

	"github.com/ryohidaka/go-zaim/internal/api"
	"github.com/ryohidaka/go-zaim/models"
)

type CreatePaymentParams struct {
	CategoryID    uint16    `url:"category_id"`
	GenreID       uint16    `url:"genre_id"`
	Amount        int32     `url:"amount"`
	Date          time.Time `url:"date"`
	ReceiptID     *int64    `url:"receipt_id"`
	FromAccountID *uint64   `url:"from_account_id,omitempty"`
	Comment       *string   `url:"comment,omitempty"`
	Name          *string   `url:"name,omitempty"`
	Place         *string   `url:"place,omitempty"`
}

// CreatePayment は支払情報を登録する
func (c *Client) CreatePayment(p CreatePaymentParams) (models.Transaction, error) {
	// レシートIDを取得
	if p.ReceiptID == nil {
		p.ReceiptID = api.GetReceiptID()
	}

	params, err := api.BuildQueryParams(p)
	if err != nil {
		return models.Transaction{}, err
	}

	body, err := c.post("home/money/payment", params)
	if err != nil {
		return models.Transaction{}, err
	}

	return api.ParseTransactionResponse(body)
}
