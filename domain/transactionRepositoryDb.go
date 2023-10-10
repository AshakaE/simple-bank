package domain

import (
	"strconv"

	"github.com/AshakaE/banking/errors"
	"github.com/AshakaE/banking/logger"
	"github.com/jmoiron/sqlx"
)

type TransactionRepositoryDb struct {
	client *sqlx.DB
}

func (db AccountRepositoryDb) Withdraw(a Account) (*Account, *errors.AppError) {
	
	sqlInsert := "INSERT INTO accounts (customer_id, opening_date, account_type, amount, status) values ($1, $2, $3, $4, $5) RETURNING customer_id"

	result, err := db.client.Exec(sqlInsert, a.CustomerId, a.OpeningDate, a.AccountType, a.Amount, a.Status)
	if err != nil {
		logger.Error("Error creating account" + err.Error())
		return nil, errors.NewUnexpectedError("Something went wrong creating account")
	}
	id, err := result.RowsAffected()
	if err != nil {
		logger.Error("Error getting last ID for new account" + err.Error())
		return nil, errors.NewUnexpectedError("Something went wrong creating account")
	}

	a.AccountId = strconv.FormatInt(id, 10)

	return &a, nil
}

func (db AccountRepositoryDb) Deposit(a Account) (*Account, *errors.AppError) {
	
	sqlInsert := "INSERT INTO accounts (customer_id, opening_date, account_type, amount, status) values ($1, $2, $3, $4, $5) RETURNING customer_id"

	result, err := db.client.Exec(sqlInsert, a.CustomerId, a.OpeningDate, a.AccountType, a.Amount, a.Status)
	if err != nil {
		logger.Error("Error creating account" + err.Error())
		return nil, errors.NewUnexpectedError("Something went wrong creating account")
	}
	id, err := result.RowsAffected()
	if err != nil {
		logger.Error("Error getting last ID for new account" + err.Error())
		return nil, errors.NewUnexpectedError("Something went wrong creating account")
	}

	a.AccountId = strconv.FormatInt(id, 10)

	return &a, nil
}

// func (db AccountRepositoryDb) CheckBalance(id string) (*Account, *errors.AppError) {
// 	sqlInsert := "SELECT amount from account WHERE account_id = $1"

// 	var a Account

// 	err := db.client.Get(&a, sqlInsert, id)
// 	if err != nil {
// 		return nil, errors.NewNotFoundError("Account not found")
// 	} else {
// 		logger.Error("Error while querying account table" + err.Error())
// 		return nil, errors.NewUnexpectedError("Unexpected database error")
// 	}

// 	return &a, nil
// }

func NewTransactionRepositoryDb(db *sqlx.DB) TransactionRepositoryDb {
	return TransactionRepositoryDb{db}
}
