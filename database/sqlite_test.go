package database

import (
	"testing"
	"time"

	"github.com/fernandomachado90/go-transactions/core"
	"github.com/stretchr/testify/assert"
)

func TestNewDatabase(t *testing.T) {
	tests := map[string]func(*testing.T){
		"Should create a new database": func(t *testing.T) {
			// when
			db, err := NewDatabase()

			// then
			assert.NoError(t, err)
			assert.NotNil(t, db)
		},
	}

	for name, run := range tests {
		t.Run(name, func(t *testing.T) {
			run(t)
		})
	}
}

func TestCreateAccount(t *testing.T) {
	// setup
	db, _ := NewDatabase()

	tests := map[string]func(*testing.T){
		"Should create a new account": func(t *testing.T) {
			// given
			input := core.Account{
				DocumentNumber: "1234567890",
			}

			// when
			output, err := db.CreateAccount(input)

			// then
			assert.NoError(t, err)
			assert.NotEmpty(t, output.ID)
			assert.Equal(t, input.DocumentNumber, output.DocumentNumber)
		},
	}

	for name, run := range tests {
		t.Run(name, func(t *testing.T) {
			run(t)
		})
	}
}

func TestFindAccount(t *testing.T) {
	// setup
	db, _ := NewDatabase()

	tests := map[string]func(*testing.T){
		"Should find a account": func(t *testing.T) {
			// given
			input := core.Account{
				DocumentNumber: "1234567890",
			}
			account, _ := db.CreateAccount(input)

			// when
			output, err := db.FindAccount(account.ID)

			// then
			assert.NoError(t, err)
			assert.Equal(t, account.ID, output.ID)
			assert.Equal(t, input.DocumentNumber, output.DocumentNumber)
		},
		"Should not find a account": func(t *testing.T) {
			// when
			_, err := db.FindAccount(0)

			// then
			assert.EqualError(t, err, "sql: no rows in result set")
		},
	}

	for name, run := range tests {
		t.Run(name, func(t *testing.T) {
			run(t)
		})
	}
}

func TestCreateTransaction(t *testing.T) {
	// setup
	db, _ := NewDatabase()
	_, _ = db.CreateAccount(core.Account{
		DocumentNumber: "1234567890",
	})

	tests := map[string]func(*testing.T){
		"Should create a new transaction": func(t *testing.T) {
			// given
			input := core.Transaction{
				AccountID:   1,
				OperationID: 4,
				Amount:      123.45,
				EventDate:   time.Now(),
			}

			// when
			output, err := db.CreateTransaction(input)

			// then
			assert.NoError(t, err)
			assert.NotEmpty(t, output.ID)
			assert.Equal(t, input.AccountID, output.AccountID)
			assert.Equal(t, input.OperationID, output.OperationID)
			assert.Equal(t, input.Amount, output.Amount)
			assert.Equal(t, input.EventDate, output.EventDate)
		},
		"Should not create a new transaction due to foreign key constraint": func(t *testing.T) {
			// given
			input := core.Transaction{
				AccountID:   1,
				OperationID: -777,
				Amount:      123.45,
				EventDate:   time.Now(),
			}

			// when
			output, err := db.CreateTransaction(input)

			// then
			assert.EqualError(t, err, "FOREIGN KEY constraint failed")
			assert.Empty(t, output)
		},
	}

	for name, run := range tests {
		t.Run(name, func(t *testing.T) {
			run(t)
		})
	}
}
