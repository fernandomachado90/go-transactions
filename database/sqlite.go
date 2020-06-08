package database

import (
	"database/sql"

	"github.com/fernandomachado90/go-transactions/core"
	_ "github.com/mattn/go-sqlite3"
)

type database struct {
	*sql.DB
}

func NewDatabase() (*database, error) {
	db, err := sql.Open("sqlite3", ":memory:")
	if err != nil {
		return nil, err
	}

	schema := `
		CREATE TABLE accounts(id INTEGER PRIMARY KEY AUTOINCREMENT, document_number CHAR(13) NOT NULL);

		CREATE TABLE operations(id INTEGER PRIMARY KEY AUTOINCREMENT, description CHAR(42) NOT NULL);
		INSERT INTO operations(description) VALUES ('COMPRA A VISTA');
		INSERT INTO operations(description) VALUES ('COMPRA PARCELADA');
		INSERT INTO operations(description) VALUES ('SAQUE');
		INSERT INTO operations(description) VALUES ('PAGAMENTO');`
	_, err = db.Exec(schema)
	if err != nil {
		return nil, err
	}

	return &database{db}, nil
}

func (db *database) CreateAccount(account core.Account) (core.Account, error) {
	query := "INSERT INTO accounts(document_number) VALUES (?)"
	result, err := db.Exec(query, account.DocumentNumber)
	if err != nil {
		return core.Account{}, err
	}
	lastId, _ := result.LastInsertId()

	account.ID = int(lastId)
	return account, nil
}

func (db *database) FindAccount(id int) (core.Account, error) {
	query := "SELECT id, document_number FROM accounts WHERE id = ?"

	result := db.QueryRow(query, id)

	account := core.Account{}
	err := result.Scan(&account.ID, &account.DocumentNumber)
	if err != nil {
		return core.Account{}, err
	}

	return account, nil
}

func (db *database) CreateTransaction(transaction core.Transaction) (core.Transaction, error) {
	return core.Transaction{}, nil
}
