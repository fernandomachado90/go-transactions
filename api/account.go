package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/fernandomachado90/go-transactions/core"
	"github.com/go-chi/chi"
)

type accountJSON struct {
	ID             int    `json:"account_id"`
	DocumentNumber string `json:"document_number"`
}

func (api *API) handleCreateAccount() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var err error
		defer func() {
			if err != nil {
				log.Println(err)
				api.respond(w, r, http.StatusInternalServerError, "An error occurred when creating the account.")
			}
		}()

		payload := new(accountJSON)
		err = json.NewDecoder(r.Body).Decode(payload)
		if err != nil {
			return
		}

		request := core.Account{
			DocumentNumber: payload.DocumentNumber,
		}
		account, err := api.accountManager.Create(request)
		if err != nil {
			return
		}

		response := accountJSON{
			ID:             account.ID,
			DocumentNumber: account.DocumentNumber,
		}
		api.respond(w, r, http.StatusCreated, response)
	}
}

func (api *API) handleFindAccount() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var err error
		defer func() {
			if err != nil {
				log.Println(err)
				api.respond(w, r, http.StatusInternalServerError, "An error occurred when finding the account.")
			}
		}()

		id, err := strconv.Atoi(chi.URLParam(r, "id"))
		if err != nil {
			return
		}

		account, err := api.accountManager.Find(id)
		if err != nil {
			return
		}

		response := accountJSON{
			ID:             account.ID,
			DocumentNumber: account.DocumentNumber,
		}
		api.respond(w, r, http.StatusFound, response)
	}
}
