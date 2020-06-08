package main

import (
	"net/http"
)

func (api *API) handleHealthCheck() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		api.respond(w, r, http.StatusOK, nil)
	}
}
