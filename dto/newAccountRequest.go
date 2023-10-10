package dto

import (
	"strings"

	"github.com/AshakaE/banking/errors"
)

type NewAccountRequest struct {
	CustomerId  string  `json:"customer_id"`
	AccountType string  `json:"account_type"`
	Amount      float64 `json:"amount"`
}

func (r NewAccountRequest) Validate() *errors.AppError {
	if r.Amount < 5000 {
		return errors.NewValidationError("To open a new account, you need a minimum balance of 5000")
	}
	if strings.ToLower(r.AccountType) != "saving" && strings.ToLower(r.AccountType) != "checking" {
		return errors.NewValidationError("Account type must be saving or checking")
	}
	return nil
}
