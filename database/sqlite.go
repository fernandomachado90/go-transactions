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
		PRAGMA foreign_keys = ON;

		CREATE TABLE accounts(id INTEGER PRIMARY KEY AUTOINCREMENT, document_number TEXT NOT NULL);

		CREATE TABLE operations(id INTEGER PRIMARY KEY AUTOINCREMENT, 
								description TEXT NOT NULL,
								credit INTEGER NOT NULL);
		INSERT INTO operations(description, credit) VALUES ('COMPRA A VISTA', 0);
		INSERT INTO operations(description, credit) VALUES ('COMPRA PARCELADA', 0);
		INSERT INTO operations(description, credit) VALUES ('SAQUE', 0);
		INSERT INTO operations(description, credit) VALUES ('PAGAMENTO', 1);
		
		CREATE TABLE transactions(id INTEGER PRIMARY KEY AUTOINCREMENT, 
									account_id INTEGER NOT NULL,
									operation_id INTEGER NOT NULL,
									amount REAL NOT NULL,
									event_date TEXT NOT NULL,
									FOREIGN KEY(account_id) REFERENCES accounts(id),
									FOREIGN KEY(operation_id) REFERENCES operations(id));`
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

func (db *database) CreateTransaction(t core.Transaction) (core.Transaction, error) {
	query := "INSERT INTO transactions(account_id, operation_id, amount, event_date) VALUES (?, ?, ?, ?)"
	result, err := db.Exec(query, t.AccountID, t.OperationID, t.Amount, t.EventDate)
	if err != nil {
		return core.Transaction{}, err
	}

	lastId, _ := result.LastInsertId()
	t.ID = int(lastId)
	return t, nil
}

func (db *database) FindOperation(id int) (core.Operation, error) {
	query := "SELECT id, credit FROM operations WHERE id = ?"

	result := db.QueryRow(query, id)

	operation := core.Operation{}
	err := result.Scan(&operation.ID, &operation.Credit)
	if err != nil {
		return core.Operation{}, err
	}

	return operation, nil
}
