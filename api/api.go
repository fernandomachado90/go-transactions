package main

import (
	"bytes"
	"encoding/json"
	"net/http"

	"github.com/fernandomachado90/go-transactions/core"
	"github.com/go-chi/chi"
)

type API struct {
	accountManager     *core.AccountManager
	transactionManager *core.TransactionManager
}

func (api *API) Routes() *chi.Mux {
	mux := chi.NewRouter()

	mux.Get("/healthcheck", api.handleHealthCheck())
	mux.Post("/accounts", api.handleCreateAccount())
	mux.Get("/accounts/{id}", api.handleFindAccount())
	mux.Post("/transactions", api.handleCreateTransaction())

	return mux
}

func (api *API) respond(w http.ResponseWriter, r *http.Request, status int, data interface{}) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(status)

	if data == nil {
		return
	}
	buffer := &bytes.Buffer{}
	encoder := json.NewEncoder(buffer)
	if err := encoder.Encode(data); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	_, _ = w.Write(buffer.Bytes())
}
