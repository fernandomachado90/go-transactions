package main

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/go-chi/chi"
	"github.com/stretchr/testify/assert"
)

func TestIntegration(t *testing.T) {
	server := newServer()
	go server.ListenAndServe()

	tests := map[string]func(*testing.T){
		"Should not reach unknown endpoint": func(t *testing.T) {
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
		},
		"Should reach healthcheck endpoint": func(t *testing.T) {
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
		},
		"Should reach Create Account endpoint": func(t *testing.T) {
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
		},
		"Should reach Find Account endpoint with success": func(t *testing.T) {
			// given
			request := httptest.NewRequest(http.MethodGet, "/accounts/1", nil)
			params := chi.RouteParams{Keys: []string{"id"}, Values: []string{"1"}}
			recorder := httptest.NewRecorder()

			// when
			server.Handler.ServeHTTP(recorder, withRouteParams(request, params))

			// then
			status := recorder.Result().StatusCode
			response := recorder.Body.String()
			assert.Equal(t, http.StatusFound, status)
			assert.JSONEq(t, `{"account_id":1, "document_number":"12345678900"}`, response)
		},
	}

	for name, run := range tests {
		t.Run(name, func(t *testing.T) {
			run(t)
		})
	}
}
