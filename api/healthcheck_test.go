package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHandleHealthCheck(t *testing.T) {
	tests := map[string]func(*testing.T){
		"Should reach healthcheck endpoint successfully": func(t *testing.T) {
			// given
			server := API{}
			request := httptest.NewRequest(http.MethodGet, "/healthcheck", nil)
			recorder := httptest.NewRecorder()

			// when
			server.handleHealthCheck()(recorder, request)

			// then
			status := recorder.Result().StatusCode
			response := recorder.Body.String()
			assert.Equal(t, http.StatusOK, status)
			assert.Empty(t, response)
		},
	}

	for name, run := range tests {
		t.Run(name, func(t *testing.T) {
			run(t)
		})
	}
}
