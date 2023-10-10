package domain

import (
	"github.com/AshakaE/banking/dto"
	"github.com/AshakaE/banking/errors"
)

const (
	WITHDRAWAL = "withdrawal"
	DEPOSIT    = "deposit"
)

type Transaction struct {
	TransactionType string  `db:"transaction_type"`
	TransactionDate string  `db:"transaction_date"`
	TransactionId   string  `db:"transaction_id"`
	AccountId       string  `db:"account_id"`
	Amount          float64 `db:"amount"`
}

func (t Transaction) IsWithdraw() bool {
	return t.TransactionType == WITHDRAWAL
}

func (r Transaction) Validate() *errors.AppError {
	if r.TransactionType != WITHDRAWAL && r.TransactionType != DEPOSIT {
		return errors.NewValidationError("Unknown transaction type")
	}
	if r.Amount < 0 {
		return errors.NewValidationError("Amount cannot be less than zero")
	}
	return nil
}

func (t Transaction) ToDto() dto.TransactionResponse {
	return dto.TransactionResponse{
		TransactionId:   t.TransactionId,
		AccountId:       t.AccountId,
		Amount:          t.Amount,
		TransactionType: t.TransactionType,
		TransactionDate: t.TransactionDate,
	}
}
