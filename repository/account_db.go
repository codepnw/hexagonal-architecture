package repository

import "github.com/jmoiron/sqlx"

type accountRepositoryDB struct {
	db *sqlx.DB
}

func NewAccountRepository(db *sqlx.DB) AccountRepository {
	return accountRepositoryDB{db: db}
}

func (r accountRepositoryDB) Create(acc Account) (*Account, error) {
	query := `
		INSERT INTO accounts (customer_id, opening_date, account_type, amount, status)
		VALUES (?, ?, ?, ?, ?)
	`
	result, err := r.db.Exec(
		query,
		acc.CustomerID,
		acc.OpeningDate,
		acc.AccountType,
		acc.Amount,
		acc.Status,
	)
	if err != nil {
		return nil, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return nil, err
	}

	acc.AccountID = int(id)
	return &acc, nil
}

func (r accountRepositoryDB) GetAll(customerID int) ([]Account, error) {
	query := `
		SELECT account_id, customer_id, opening_date, account_type, amount, status 
		FROM accounts 
		WHERE customer_id=?
	`
	accounts := []Account{}

	if err := r.db.Select(&accounts, query, customerID); err != nil {
		return nil, err
	}
	return accounts, nil
}
