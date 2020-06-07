package database

import (
	"testing"

	"github.com/fernandomachado90/go-transactions/core"
	"github.com/stretchr/testify/assert"
)

func TestNewRepository(t *testing.T) {
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
		"Should find a new account": func(t *testing.T) {
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
	}

	for name, run := range tests {
		t.Run(name, func(t *testing.T) {
			run(t)
		})
	}
}
