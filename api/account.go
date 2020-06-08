package main

import (
	"encoding/json"
	"net/http"

	"github.com/fernandomachado90/go-transactions/core"
)

type payload struct {
	ID             int    `json:"account_id"`
	DocumentNumber string `json:"document_number"`
}

func (api *API) handleCreateAccount() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var err error
		defer func() {
			if err != nil {
				api.respond(w, r, http.StatusInternalServerError, "An error occurred when creating the account.")
			}
		}()

		req := new(payload)
		err = json.NewDecoder(r.Body).Decode(req)
		if err != nil {
			return
		}

		request := core.Account{
			DocumentNumber: req.DocumentNumber,
		}
		account, err := api.accountManager.Create(request)
		if err != nil {
			return
		}

		response := payload{
			ID:             account.ID,
			DocumentNumber: account.DocumentNumber,
		}
		api.respond(w, r, http.StatusCreated, response)
	}
}
