package domain

import (
	"database/sql"

	"github.com/AshakaE/banking/errors"
	"github.com/AshakaE/banking/logger"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type CustomerRepositoryDb struct {
	client *sqlx.DB
}

func (d CustomerRepositoryDb) FindAll(status string) ([]Customer, *errors.AppError) {
	var err error
	customers := make([]Customer, 0)

	if status == "" {
		findAllSql := "select customer_id, name, city, zipcode, date_of_birth, status from customers"
		err = d.client.Select(&customers, findAllSql)
	} else {
		findAllSql := "select customer_id, name, city, zipcode, date_of_birth, status from customers where status = $1"
		err = d.client.Select(&customers, findAllSql, status)
	}

	if err != nil {
		logger.Error("Error while querying customer table" + err.Error())
		return nil, errors.NewUnexpectedError("Something went wrong")
	}

	return customers, nil
}

func (d CustomerRepositoryDb) FindById(id string) (*Customer, *errors.AppError) {
	findById := "select customer_id, name, city, zipcode, date_of_birth, status from customers where customer_id = $1"

	var c Customer
	err := d.client.Get(&c, findById, id)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.NewNotFoundError("Customer not found")
		} else {
			logger.Error("Error while querying customer table" + err.Error())
			return nil, errors.NewUnexpectedError("Unexpected database error")
		}
	}

	return &c, nil
}

func NewCustomerRepositoryDb(db *sqlx.DB) CustomerRepositoryDb {
	return CustomerRepositoryDb{db}
}
