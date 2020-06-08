package core

import (
	"errors"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
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
				AccountID:   2,
				OperationID: 4,
				Amount:      123.45,
			}
			db := new(dbMock)
			db.On("CreateTransaction", mock.AnythingOfType("Transaction")).
				Return(1, nil)
			transactionManager := NewTransactionManager(db)

			// when
			output, err := transactionManager.Create(input)

			// then
			assert.Equal(t, 1, output.ID)
			assert.Equal(t, input.AccountID, output.AccountID)
			assert.Equal(t, input.OperationID, output.OperationID)
			assert.Equal(t, input.Amount, output.Amount)
			assert.NotEmpty(t, output.EventDate)
			assert.NoError(t, err)
		},
		"Should not create account because of missing required attribute": func(t *testing.T) {
			//given
			input := Transaction{}
			db := new(dbMock)
			transactionManager := NewTransactionManager(db)

			// when
			output, err := transactionManager.Create(input)

			// then
			assert.Empty(t, output)
			assert.EqualError(t, err, "missing required attribute")
		},
		"Should not create transaction because of database error": func(t *testing.T) {
			//given
			input := Transaction{
				AccountID:   1,
				OperationID: 4,
				Amount:      123.45,
			}
			db := new(dbMock)
			db.On("CreateTransaction", mock.AnythingOfType("Transaction")).
				Return(0, errors.New("database error"))
			transactionManager := NewTransactionManager(db)

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
