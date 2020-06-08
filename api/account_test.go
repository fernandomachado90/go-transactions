package main

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/fernandomachado90/go-transactions/core"
	"github.com/stretchr/testify/assert"
)

func TestHandleCreateAccount(t *testing.T) {
	// setup
	db := new(dbMock)
	accountManager := core.NewAccountManager(db)
	server := API{accountManager}

	tests := map[string]func(*testing.T){
		"Should reach Create Account endpoint with success": func(t *testing.T) {
			// mock
			input := core.Account{
				DocumentNumber: "12345678900",
			}
			db.On("CreateAccount", input).Return(1, nil)

			// given
			body := strings.NewReader(`{"document_number": "12345678900"}`)
			request := httptest.NewRequest(http.MethodPost, "/accounts", body)
			recorder := httptest.NewRecorder()

			// when
			server.handleCreateAccount()(recorder, request)

			// then
			status := recorder.Result().StatusCode
			response := recorder.Body.String()
			assert.Equal(t, http.StatusCreated, status)
			assert.JSONEq(t, `{"account_id":1, "document_number":"12345678900"}`, response)
		},
		"Should reach Create Account endpoint without the required body and receive an error": func(t *testing.T) {
			// given
			body := strings.NewReader(`{"wrong_info": "here"}`)
			request := httptest.NewRequest(http.MethodPost, "/accounts", body)
			recorder := httptest.NewRecorder()

			// when
			server.handleCreateAccount()(recorder, request)

			// then
			status := recorder.Result().StatusCode
			response := recorder.Body.String()
			assert.Equal(t, http.StatusInternalServerError, status)
			assert.Equal(t, "\"An error occurred when creating the account.\"\n", response)
		},
	}

	for name, run := range tests {
		t.Run(name, func(t *testing.T) {
			run(t)
		})
	}
}

func (m *dbMock) CreateAccount(account core.Account) (core.Account, error) {
	args := m.Called(account)
	id := args.Int(0)
	err := args.Error(1)

	if err != nil {
		return core.Account{}, err
	}

	account.ID = id
	return account, nil
}

func (m *dbMock) FindAccount(id int) (core.Account, error) {
	args := m.Called(id)
	err := args.Error(0)

	if err != nil {
		return core.Account{}, err
	}

	return core.Account{ID: id}, nil
}
