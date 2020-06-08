package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/fernandomachado90/go-transactions/core"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestHandleCreateTransaction(t *testing.T) {
	// setup
	db := new(dbMock)
	transactionManager := core.NewTransactionManager(db)
	server := API{transactionManager: transactionManager}

	tests := map[string]func(*testing.T){
		"Should reach Create Transaction endpoint successfully": func(t *testing.T) {
			// mock
			db.On("FindOperation", 4).Return(true, nil)
			db.On("CreateTransaction", mock.AnythingOfType("Transaction")).
				Return(1, nil)

			// given
			body := strings.NewReader(`{ "account_id": 2, "operation_type_id": 4, "amount": 123.45 }`)
			request := httptest.NewRequest(http.MethodPost, "/transactions", body)
			recorder := httptest.NewRecorder()

			// when
			server.handleCreateTransaction()(recorder, request)

			// then
			status := recorder.Result().StatusCode
			response := new(transactionJSON)
			_ = json.Unmarshal(recorder.Body.Bytes(), response)
			assert.Equal(t, http.StatusCreated, status)
			assert.Equal(t, 1, response.ID)
			assert.Equal(t, 2, response.AccountID)
			assert.Equal(t, 4, response.OperationID)
			assert.Equal(t, 123.45, response.Amount)
			assert.NotEmpty(t, response.EventDate)
		},
		"Should reach Create Transaction endpoint without the required body and receive an error": func(t *testing.T) {
			// given
			body := strings.NewReader(`{"wrong_info": "here"}`)
			request := httptest.NewRequest(http.MethodPost, "/transactions", body)
			recorder := httptest.NewRecorder()

			// when
			server.handleCreateTransaction()(recorder, request)

			// then
			status := recorder.Result().StatusCode
			response := recorder.Body.String()
			assert.Equal(t, http.StatusInternalServerError, status)
			assert.Equal(t, "\"An error occurred when creating the transaction.\"\n", response)
		},
	}

	for name, run := range tests {
		t.Run(name, func(t *testing.T) {
			run(t)
		})
	}
}
