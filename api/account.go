package main

import (
	"encoding/json"
	"net/http"

	"github.com/fernandomachado90/go-transactions/core"
)

func (api *API) handleCreateAccount() http.HandlerFunc {
	type request struct {
		DocumentNumber string `json:"document_number"`
	}

	return func(w http.ResponseWriter, r *http.Request) {
		var err error
		defer func() {
			if err != nil {
				api.respond(w, r, http.StatusInternalServerError, "An error occurred when creating the account.")
			}
		}()

		req := new(request)
		err = json.NewDecoder(r.Body).Decode(req)
		if err != nil {
			return
		}

		account := core.Account{
			DocumentNumber: req.DocumentNumber,
		}
		account, err = api.accountManager.Create(account)
		if err != nil {
			return
		}

		api.respond(w, r, http.StatusCreated, account)
	}
}
