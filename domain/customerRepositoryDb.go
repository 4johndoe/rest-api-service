package domain

import (
	"banking/errs"
	"banking/logger"
	"database/sql"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type CustomerRepositoryDb struct {
	client *sqlx.DB
}

func (d CustomerRepositoryDb) FindAll() ([]Customer, error) {
	var err error
	var status string
	customers := make([]Customer, 0)

	if status == "" {
		findALlSql := "select customer_id, name, city, zipcode, date_of_birth, status from customers"
		err = d.client.Select(&customers, findALlSql)
	} else {
		findALlSql := "select customer_id, name, city, zipcode, date_of_birth, status from customers where status = $1"
		err = d.client.Select(&customers, findALlSql, status)
	}

	if err != nil {
		logger.Error("Error while querying customer table " + err.Error())
		return nil, err
	}

	return customers, nil
}

func (d CustomerRepositoryDb) ById(id string) (*Customer, *errs.AppError) {
	customerSql := "select customer_id, name, city, zipcode, date_of_birth, status from customers where customer_id = $1"

	var c Customer
	err := d.client.Get(&c, customerSql, id)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errs.NewNotFoundError("Customer not found")
		} else {
			logger.Error("Error while scanning customer " + err.Error())
			return nil, errs.NewUnexpectedError("Unexpected database error")
		}
	}
	return &c, nil
}

func NewCustomerRepositoryDb(dbClient *sqlx.DB) CustomerRepositoryDb {
	return CustomerRepositoryDb{dbClient}
}
