package core

import (
	"errors"
	"math/rand"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestNewTransaction(t *testing.T) {
	tests := map[string]func(*testing.T){
		"Should initialize entity": func(t *testing.T) {
			//given
			id := 1
			accountID := 1
			operationID := 1
			amount := -0.20
			eventDate := time.Now()

			// when
			transaction := &Transaction{
				ID:          id,
				AccountID:   accountID,
				OperationID: operationID,
				Amount:      amount,
				EventDate:   eventDate,
			}

			// then
			assert.NotEmpty(t, transaction)
		},
	}

	for name, run := range tests {
		t.Run(name, func(t *testing.T) {
			run(t)
		})
	}
}

func TestCreateTransaction(t *testing.T) {
	tests := map[string]func(*testing.T){
		"Should create transaction successfully": func(t *testing.T) {
			//given
			input := Transaction{
				AccountID:   1,
				OperationID: 4,
				Amount:      123.45,
			}
			db := new(dbMock)
			db.On("CreateTransaction", input).Return(rand.Int(), nil)
			transactionManager := TransactionManager{db: db}

			// when
			output, err := transactionManager.Create(input)

			// then
			assert.NotEmpty(t, output.ID)
			assert.NotEmpty(t, output.EventDate)
			assert.Equal(t, input.AccountID, output.AccountID)
			assert.Equal(t, input.OperationID, output.OperationID)
			assert.Equal(t, input.Amount, output.Amount)
			assert.NoError(t, err)
		},
		"Should not create transaction and return error": func(t *testing.T) {
			//given
			input := Transaction{
				AccountID:   1,
				OperationID: 4,
				Amount:      123.45,
			}
			db := new(dbMock)
			db.On("CreateTransaction", input).Return(0, errors.New("error"))
			transactionManager := TransactionManager{db: db}

			// when
			output, err := transactionManager.Create(input)

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

func (m *dbMock) CreateTransaction(transaction Transaction) (Transaction, error) {
	args := m.Called(transaction)
	id := args.Int(0)
	err := args.Error(1)

	if err != nil {
		return Transaction{}, err
	}

	transaction.ID = id
	transaction.EventDate = time.Now()
	return transaction, nil
}
