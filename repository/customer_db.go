package repository

import "github.com/jmoiron/sqlx"

type customerRepositoryDB struct {
	db *sqlx.DB
}

func NewCustomerRepositoryDB(db *sqlx.DB) CustomerRepository {
	return customerRepositoryDB{db: db}
}

func (r customerRepositoryDB) GetAll() ([]Customer, error) {
	customers := []Customer{}

	query := "SELECT customer_id, name, date_of_birth, city, zipcode, status FROM customers"
	if err := r.db.Select(&customers, query); err != nil {
		return nil, err
	}

	return customers, nil
}

func (r customerRepositoryDB) GetById(id int) (*Customer, error) {
	customer := Customer{}

	query := "SELECT customer_id, name, date_of_birth, city, zipcode, status FROM customers WHERE customer_id=?"
	if err := r.db.Get(&customer, query, id); err != nil {
		return nil, err
	}

	return &customer, nil
}
