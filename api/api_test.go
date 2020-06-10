package main

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/go-chi/chi"
	"github.com/stretchr/testify/assert"
)

func TestAPI_404(t *testing.T) {
	server := newServer()
	go server.ListenAndServe()

	// given
	request, _ := http.NewRequest(http.MethodGet, "/unknown-url", nil)
	recorder := httptest.NewRecorder()

	// when
	server.Handler.ServeHTTP(recorder, request)

	// then
	status := recorder.Result().StatusCode
	response := recorder.Body.String()
	assert.Equal(t, http.StatusNotFound, status)
	assert.Equal(t, "404 page not found\n", response)
}

func TestAPI_GET_HealthCheck(t *testing.T) {
	server := newServer()
	go server.ListenAndServe()

	// given
	request, _ := http.NewRequest(http.MethodGet, "/healthcheck", nil)
	recorder := httptest.NewRecorder()

	// when
	server.Handler.ServeHTTP(recorder, request)

	// then
	status := recorder.Result().StatusCode
	response := recorder.Body.String()
	assert.Equal(t, http.StatusOK, status)
	assert.Empty(t, response)
}

func TestAPI_POST_Accounts(t *testing.T) {
	server := newServer()
	go server.ListenAndServe()

	// given
	body := strings.NewReader(`{"document_number": "12345678900"}`)
	request := httptest.NewRequest(http.MethodPost, "/accounts", body)
	recorder := httptest.NewRecorder()

	// when
	server.Handler.ServeHTTP(recorder, request)

	// then
	status := recorder.Result().StatusCode
	response := recorder.Body.String()
	assert.Equal(t, http.StatusCreated, status)
	assert.JSONEq(t, `{"account_id":1, "document_number":"12345678900"}`, response)
}

func TestAPI_GET_Accounts(t *testing.T) {
	server := newServer()
	go server.ListenAndServe()

	// scenario
	body := strings.NewReader(`{"document_number": "12345678900"}`)
	request := httptest.NewRequest(http.MethodPost, "/accounts", body)
	server.Handler.ServeHTTP(httptest.NewRecorder(), request)

	// given
	request = httptest.NewRequest(http.MethodGet, "/accounts/1", nil)
	recorder := httptest.NewRecorder()
	params := chi.RouteParams{Keys: []string{"id"}, Values: []string{"1"}}

	// when
	server.Handler.ServeHTTP(recorder, withRouteParams(request, params))

	// then
	status := recorder.Result().StatusCode
	response := recorder.Body.String()
	assert.Equal(t, http.StatusFound, status)
	assert.JSONEq(t, `{"account_id":1, "document_number":"12345678900"}`, response)
}

func TestAPI_POST_Transactions(t *testing.T) {
	server := newServer()
	go server.ListenAndServe()

	// scenario
	body := strings.NewReader(`{"document_number": "12345678900"}`)
	request := httptest.NewRequest(http.MethodPost, "/accounts", body)
	server.Handler.ServeHTTP(httptest.NewRecorder(), request)

	// given
	body = strings.NewReader(`{ "account_id": 1, "operation_type_id": 4, "amount": 123.45 }`)
	request = httptest.NewRequest(http.MethodPost, "/transactions", body)
	recorder := httptest.NewRecorder()

	// when
	server.Handler.ServeHTTP(recorder, request)

	// then
	status := recorder.Result().StatusCode
	response := new(transactionJSON)
	_ = json.Unmarshal(recorder.Body.Bytes(), response)
	assert.Equal(t, http.StatusCreated, status)
	assert.Equal(t, 1, response.ID)
	assert.Equal(t, 1, response.AccountID)
	assert.Equal(t, 4, response.OperationID)
	assert.Equal(t, 123.45, response.Amount)
	assert.NotEmpty(t, response.EventDate)
}

func withRouteParams(r *http.Request, urlParams chi.RouteParams) *http.Request {
	return r.WithContext(context.WithValue(r.Context(), chi.RouteCtxKey, &chi.Context{URLParams: urlParams}))
}
