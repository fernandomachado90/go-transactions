package core

import (
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
