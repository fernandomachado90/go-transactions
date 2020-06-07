package core

import (
	"errors"
	"math/rand"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
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
			accountManager := AccountManager{db: db}

			// when
			output, err := accountManager.Create(input)

			// then
			assert.NotEmpty(t, output.ID)
			assert.NoError(t, err)
		},
		"Should not create account and return error": func(t *testing.T) {
			//given
			input := Account{
				DocumentNumber: "1234567890",
			}
			db := new(dbMock)
			db.On("CreateAccount", input).Return(0, errors.New("error"))
			accountManager := AccountManager{db: db}

			// when
			output, err := accountManager.Create(input)

			// then
			assert.Empty(t, output.ID)
			assert.Error(t, err)
		},
	}

	for name, run := range tests {
		t.Run(name, func(t *testing.T) {
			run(t)
		})
	}
}

type dbMock struct {
	mock.Mock
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
