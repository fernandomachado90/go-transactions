package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHandleHealthCheck(t *testing.T) {
	// given
	server := &API{}
	request := httptest.NewRequest(http.MethodGet, "/healthcheck", nil)
	recorder := httptest.NewRecorder()

	// when
	server.handleHealthCheck()(recorder, request)

	// then
	response := recorder.Result()
	assert.Equal(t, http.StatusOK, response.StatusCode)
}
