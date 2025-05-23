package zaim

import (
	"fmt"
	"time"

	"github.com/ryohidaka/go-zaim/internal/api"
	"github.com/ryohidaka/go-zaim/models"
)

type CreateIncomeParams struct {
	CategoryID  uint16    `url:"category_id"`
	Amount      int32     `url:"amount"`
	Date        time.Time `url:"date"`
	ToAccountID *uint64   `url:"to_account_id,omitempty"`
	Comment     *string   `url:"comment,omitempty"`
	Place       *string   `url:"place,omitempty"`
}

// CreateIncome は収入情報を登録する
func (c *Client) CreateIncome(p CreateIncomeParams) (models.Transaction, error) {
	params, err := api.BuildQueryParams(p)
	if err != nil {
		return models.Transaction{}, err
	}

	body, err := c.post("home/money/income", params)
	if err != nil {
		return models.Transaction{}, err
	}

	return api.ParseTransactionResponse(body)
}

type UpdateIncomeParams struct {
	Amount        int32     `url:"amount"`
	Date          time.Time `url:"date"`
	FromAccountID *uint64   `url:"from_account_id,omitempty"`
	CategoryID    uint16    `url:"category_id"`
	GenreID       uint16    `url:"genre_id"`
	Comment       *string   `url:"comment,omitempty"`
	Name          *string   `url:"name,omitempty"`
}

// UpdateIncome は収入情報を更新する
func (c *Client) UpdateIncome(id uint64, p UpdateIncomeParams) (models.Transaction, error) {
	params, err := api.BuildQueryParams(p)
	if err != nil {
		return models.Transaction{}, err
	}

	path := fmt.Sprintf("home/money/income/%d", id)
	body, err := c.put(path, params)
	if err != nil {
		return models.Transaction{}, err
	}

	return api.ParseTransactionResponse(body)
}

// DeleteIncome は収入情報を削除する
func (c *Client) DeleteIncome(id uint64) (models.Transaction, error) {
	path := fmt.Sprintf("home/money/income/%d", id)
	body, err := c.delete(path)
	if err != nil {
		return models.Transaction{}, err
	}

	return api.ParseTransactionResponse(body)
}
