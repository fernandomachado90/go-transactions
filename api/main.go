package main

import (
	"log"
	"net/http"

	"github.com/fernandomachado90/go-transactions/core"
	"github.com/fernandomachado90/go-transactions/database"
)

const address = ":8080"

func newServer() *http.Server {
	log.Printf("Starting database")
	db, err := database.NewDatabase()
	if err != nil {
		log.Fatalf("Failed to initalize Database. %s", err)
	}

	log.Printf("Starting server @%s", address)
	accountManager := core.NewAccountManager(db)
	transactionManager := core.NewTransactionManager(db)

	api := API{accountManager, transactionManager}
	server := &http.Server{
		Addr:    address,
		Handler: api.Routes(),
	}
	return server
}

func main() {
	server := newServer()
	if err := server.ListenAndServe(); err != nil {
		log.Fatalf("Failed to start server. %s", err)
	}
}
