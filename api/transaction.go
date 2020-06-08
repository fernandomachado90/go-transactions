package main

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/fernandomachado90/go-transactions/core"
)

type transactionJSON struct {
	ID          int       `json:"transaction_id"`
	AccountID   int       `json:"account_id"`
	OperationID int       `json:"operation_type_id"`
	Amount      float64   `json:"amount"`
	EventDate   time.Time `json:"event_data"`
}

func (api *API) handleCreateTransaction() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var err error
		defer func() {
			if err != nil {
				log.Println(err)
				api.respond(w, r, http.StatusInternalServerError, "An error occurred when creating the transaction.")
			}
		}()

		payload := new(transactionJSON)
		err = json.NewDecoder(r.Body).Decode(payload)
		if err != nil {
			return
		}

		request := core.Transaction{
			AccountID:   payload.AccountID,
			OperationID: payload.OperationID,
			Amount:      payload.Amount,
		}
		transaction, err := api.transactionManager.Create(request)
		if err != nil {
			return
		}

		response := transactionJSON{
			ID:          transaction.ID,
			AccountID:   transaction.AccountID,
			OperationID: transaction.OperationID,
			Amount:      transaction.Amount,
			EventDate:   transaction.EventDate,
		}
		api.respond(w, r, http.StatusCreated, response)
	}
}
