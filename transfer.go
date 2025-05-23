package zaim

import (
	"time"

	"github.com/ryohidaka/go-zaim/internal/api"
	"github.com/ryohidaka/go-zaim/models"
)

type CreateTransferParams struct {
	Amount        int32     `url:"amount"`
	Date          time.Time `url:"date"`
	FromAccountID uint64    `url:"from_account_id,omitempty"`
	ToAccountID   uint64    `url:"to_account_id,omitempty"`
	Comment       *string   `url:"comment,omitempty"`
}

// CreateTransfer は振替情報を登録する
func (c *Client) CreateTransfer(p CreateTransferParams) (models.Transaction, error) {
	params, err := api.BuildQueryParams(p)
	if err != nil {
		return models.Transaction{}, err
	}

	body, err := c.post("home/money/transfer", params)
	if err != nil {
		return models.Transaction{}, err
	}

	return api.ParseTransactionResponse(body)
}
