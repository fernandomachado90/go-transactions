package main

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/fernandomachado90/go-transactions/core"
	"github.com/fernandomachado90/go-transactions/database"
	"github.com/stretchr/testify/assert"
)

func TestHandleCreateAccount(t *testing.T) {
	// setup
	db, _ := database.NewDatabase()
	accountManager := core.NewAccountManager(db)
	server := API{accountManager}

	tests := map[string]func(*testing.T){
		"Should reach Create Account endpoint with success": func(t *testing.T) {
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
			assert.JSONEq(t, `{"ID":1,"DocumentNumber":"12345678900"}`, response)
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
