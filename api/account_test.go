package main

import (
	"context"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/fernandomachado90/go-transactions/core"
	"github.com/go-chi/chi"
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

func TestHandleFindAccount(t *testing.T) {
	// setup
	db := new(dbMock)
	accountManager := core.NewAccountManager(db)
	server := API{accountManager}

	tests := map[string]func(*testing.T){
		"Should reach Find Account endpoint with success": func(t *testing.T) {
			// mock
			db.On("FindAccount", 1).Return("12345678900", nil)

			// given
			request := httptest.NewRequest(http.MethodGet, "/accounts/1", nil)
			params := chi.RouteParams{Keys: []string{"id"}, Values: []string{"1"}}
			recorder := httptest.NewRecorder()

			// when
			server.handleFindAccount()(recorder, withRouteParams(request, params))

			// then
			status := recorder.Result().StatusCode
			response := recorder.Body.String()
			assert.Equal(t, http.StatusFound, status)
			assert.JSONEq(t, `{"account_id":1, "document_number":"12345678900"}`, response)
		},
		"Should reach Find Account endpoint without the required param and receive an error": func(t *testing.T) {
			// given
			request := httptest.NewRequest(http.MethodGet, "/accounts/not-an-id", nil)
			params := chi.RouteParams{Keys: []string{"id"}, Values: []string{"not-an-id"}}
			recorder := httptest.NewRecorder()

			// when
			server.handleFindAccount()(recorder, withRouteParams(request, params))

			// then
			status := recorder.Result().StatusCode
			response := recorder.Body.String()
			assert.Equal(t, http.StatusInternalServerError, status)
			assert.Equal(t, "\"An error occurred when finding the account.\"\n", response)
		},
	}

	for name, run := range tests {
		t.Run(name, func(t *testing.T) {
			run(t)
		})
	}
}

func withRouteParams(r *http.Request, urlParams chi.RouteParams) *http.Request {
	return r.WithContext(context.WithValue(r.Context(), chi.RouteCtxKey, &chi.Context{URLParams: urlParams}))
}
