package domain

import (
	"github.com/AshakaE/banking/dto"
	"github.com/AshakaE/banking/errors"
)

const dbTSLayout = "2006-01-02 15:04:05"

type Account struct {
	AccountId   string  `db:"account_id"`
	CustomerId  string  `db:"customer_id"`
	OpeningDate string  `db:"opening_date"`
	AccountType string  `db:"account_type"`
	Amount      float64 `db:"amount"`
	Status      string  `db:"status"`
}

func (a Account) ToNewAccountResponseDto() *dto.NewAccountResponse {
	return &dto.NewAccountResponse{AccountId: a.AccountId}
}

type AccountRepository interface {
	Save(account Account) (*Account, *errors.AppError)
	SaveTransaction(transaction Transaction) (*Transaction, *errors.AppError)
	FindBy(accountId string) (*Account, *errors.AppError)
}

func (a Account) CannnotWithdraw(amount float64) bool {
	return a.Amount < amount
}

func NewAccount(customerId, accountType string, amount float64) Account {
	return Account{
		CustomerId:  customerId,
		OpeningDate: dbTSLayout,
		AccountType: accountType,
		Amount:      amount,
		Status:      "1",
	}
}

