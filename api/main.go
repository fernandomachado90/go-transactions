package main

import (
	"flag"
	"log"
	"net/http"

	"github.com/fernandomachado90/go-transactions/core"
	"github.com/fernandomachado90/go-transactions/database"
)

var (
	port = flag.String("port", "8888", "Informs the port where the API will be made available.")
)

func newServer() *http.Server {
	flag.Parse()
	address := ":" + *port

	log.Printf("Starting database")
	db, err := database.NewDatabase()
	if err != nil {
		log.Fatalf("Failed to initalize Database. %s", err)
	}

	log.Printf("Starting server @%s", address)
	accountManager := core.NewAccountManager(db)
	api := API{accountManager}
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
