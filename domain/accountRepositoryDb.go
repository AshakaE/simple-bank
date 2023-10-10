package domain

import (
	"strconv"

	"github.com/AshakaE/banking/errors"
	"github.com/AshakaE/banking/logger"
	"github.com/jmoiron/sqlx"
)

type AccountRepositoryDb struct {
	client *sqlx.DB
}

func (db AccountRepositoryDb) Save(a Account) (*Account, *errors.AppError) {
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


func (d AccountRepositoryDb) SaveTransaction(t Transaction) (*Transaction, *errors.AppError) {
	// starting the database transaction block
	tx, err := d.client.Begin()
	if err != nil {
		logger.Error("Error while starting a new transaction for bank account transaction: " + err.Error())
		return nil, errors.NewUnexpectedError("Unexpected database error")
	}
	// fmt.Println("On to the next transaction 1")
	// inserting bank account transaction
	result, _ := tx.Exec(`INSERT INTO transactions (account_id, amount, transaction_type, transaction_date) 
											values ($1, $2, $3, $4)`, t.AccountId, t.Amount, t.TransactionType, t.TransactionDate)

	// updating account balance
	if t.IsWithdraw() {
		_, err = tx.Exec(`UPDATE accounts SET amount = amount - $1 where account_id = $2`, t.Amount, t.AccountId)
		} else {
		_, err = tx.Exec(`UPDATE accounts SET amount = amount + $1 where account_id = $2`, t.Amount, t.AccountId)
	}

	// in case of error Rollback, and changes from both the tables will be reverted
	if err != nil {
		tx.Rollback()
		logger.Error("Error while saving transaction: " + err.Error())
		return nil, errors.NewUnexpectedError("Unexpected database error")
	}
	// commit the transaction when all is good
	err = tx.Commit()
	if err != nil {
		tx.Rollback()
		logger.Error("Error while commiting transaction for bank account: " + err.Error())
		return nil, errors.NewUnexpectedError("Unexpected database error")
	}
	// getting the last transaction ID from the transaction table
	transactionId, err := result.RowsAffected()
	if err != nil {
		logger.Error("Error while getting the last transaction id: " + err.Error())
		return nil, errors.NewUnexpectedError("Unexpected database error")
	}

	// Getting the latest account information from the accounts table
	account, appErr := d.FindBy(t.AccountId)
	if appErr != nil {
		return nil, appErr
	}
	t.TransactionId = strconv.FormatInt(transactionId, 10)

	// updating the transaction struct with the latest balance
	t.Amount = account.Amount
	return &t, nil
}


func (d AccountRepositoryDb) FindBy(accountId string) (*Account, *errors.AppError) {
	sqlGetAccount := "SELECT account_id, customer_id, opening_date, account_type, amount from accounts where account_id = $1"
	var account Account
	err := d.client.Get(&account, sqlGetAccount, accountId)
	if err != nil {
		logger.Error("Error while fetching account information: " + err.Error())
		return nil, errors.NewUnexpectedError("Unexpected database error")
	}
	return &account, nil
}



func NewAccountRepositoryDb(db *sqlx.DB) AccountRepositoryDb {
	return AccountRepositoryDb{db}
}
