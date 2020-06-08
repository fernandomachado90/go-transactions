package core

import (
	"errors"
	"math/rand"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewAccount(t *testing.T) {
	tests := map[string]func(*testing.T){
		"Should initialize entity": func(t *testing.T) {
			//given
			id := 1
			documentNumber := "1234567890"

			// when
			account := Account{
				ID:             id,
				DocumentNumber: documentNumber,
			}

			// then
			assert.NotEmpty(t, account)
		},
	}

	for name, run := range tests {
		t.Run(name, func(t *testing.T) {
			run(t)
		})
	}
}

func TestCreateAccount(t *testing.T) {
	tests := map[string]func(*testing.T){
		"Should create account successfully": func(t *testing.T) {
			//given
			input := Account{
				DocumentNumber: "1234567890",
			}
			db := new(dbMock)
			db.On("CreateAccount", input).Return(rand.Int(), nil)
			accountManager := NewAccountManager(db)

			// when
			output, err := accountManager.Create(input)

			// then
			assert.NotEmpty(t, output.ID)
			assert.Equal(t, input.DocumentNumber, output.DocumentNumber)
			assert.NoError(t, err)
		},
		"Should not create account because of missing required attribute": func(t *testing.T) {
			//given
			input := Account{}
			db := new(dbMock)
			accountManager := NewAccountManager(db)

			// when
			output, err := accountManager.Create(input)

			// then
			assert.Empty(t, output)
			assert.EqualError(t, err, "missing required attribute")
		},
		"Should not create account because of database error": func(t *testing.T) {
			//given
			input := Account{
				DocumentNumber: "1234567890",
			}
			db := new(dbMock)
			db.On("CreateAccount", input).Return(0, errors.New("database error"))
			accountManager := NewAccountManager(db)

			// when
			output, err := accountManager.Create(input)

			// then
			assert.Empty(t, output)
			assert.EqualError(t, err, "database error")
		},
	}

	for name, run := range tests {
		t.Run(name, func(t *testing.T) {
			run(t)
		})
	}
}

func TestFindAccount(t *testing.T) {
	tests := map[string]func(*testing.T){
		"Should find account successfully": func(t *testing.T) {
			//given
			input := rand.Int()
			db := new(dbMock)
			db.On("FindAccount", input).Return(nil)
			accountManager := NewAccountManager(db)

			// when
			output, err := accountManager.Find(input)

			// then
			assert.NotEmpty(t, output.ID)
			assert.NoError(t, err)
		},
		"Should not find account and return error": func(t *testing.T) {
			//given
			input := rand.Int()
			db := new(dbMock)
			db.On("FindAccount", input).Return(errors.New("error"))
			accountManager := NewAccountManager(db)

			// when
			output, err := accountManager.Find(input)

			// then
			assert.Empty(t, output)
			assert.Error(t, err)
		},
	}

	for name, run := range tests {
		t.Run(name, func(t *testing.T) {
			run(t)
		})
	}
}

func (m *dbMock) CreateAccount(account Account) (Account, error) {
	args := m.Called(account)
	id := args.Int(0)
	err := args.Error(1)

	if err != nil {
		return Account{}, err
	}

	account.ID = id
	return account, nil
}

func (m *dbMock) FindAccount(id int) (Account, error) {
	args := m.Called(id)
	err := args.Error(0)

	if err != nil {
		return Account{}, err
	}

	return Account{ID: id}, nil
}
