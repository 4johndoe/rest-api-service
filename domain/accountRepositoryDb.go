package domain

import (
	"banking/errs"
	"github.com/jmoiron/sqlx"
	"strconv"
)

type AccountRepositoryDb struct {
	client *sqlx.DB
}

func (d AccountRepositoryDb) Save(a Account) (*Account, *errs.AppError) {
	var id int64
	sqlInsert := "INSERT INTO accounts (customer_id, opening_date, account_type, amount, status) values ($1, $2, $3, $4, $5) returning account_id"

	row := d.client.QueryRow(sqlInsert, a.CustomerId, a.OpeningDate, a.AccountType, a.Amount, a.Status)
	if err := row.Scan(&id); err != nil {
		return nil, errs.NewUnexpectedError(err.Error())
	}

	a.AccountId = strconv.FormatInt(id, 10)
	return &a, nil
}

func NewAccountRepositoryDb(dbClient *sqlx.DB) AccountRepositoryDb {
	return AccountRepositoryDb{dbClient}
}
