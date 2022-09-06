package domain

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"log"
	"time"
)

type CustomerRepositoryDb struct {
	client *sql.DB
}

func (d CustomerRepositoryDb) FindAll() ([]Customer, error) {
	findALlSql := "select customer_id, name, city, zipcode, date_of_birth, status from customers"

	rows, err := d.client.Query(findALlSql)
	if err != nil {
		log.Println("Error while querying customer table " + err.Error())
		return nil, err
	}

	customers := make([]Customer, 0)
	for rows.Next() {
		var c Customer
		err := rows.Scan(&c.Id, &c.Name, &c.City, &c.Zipcode, &c.DateofBirth, &c.Status)
		if err != nil {
			log.Println("Error while scanning customers " + err.Error())
			return nil, err
		}
		customers = append(customers, c)
	}
	return customers, nil
}

func NewCustomerRepositoryDb() CustomerRepositoryDb {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		"localhost", 5434, "postgres", "postgres", "postgres")
	client, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}

	client.SetConnMaxLifetime(time.Minute * 3)
	client.SetMaxOpenConns(10)
	client.SetMaxIdleConns(10)
	return CustomerRepositoryDb{client}
}
