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

func main() {
	flag.Parse()
	address := ":" + *port
	log.Printf("Starting server @%s", address)

	db, err := database.NewDatabase()
	defer func() {
		_ = db.Close()
	}()

	if err != nil {
		log.Fatalf("Failed to initalize Database. %s", err)
	}

	accountManager := core.NewAccountManager(db)
	server := API{accountManager}

	if err := http.ListenAndServe(address, server.Routes()); err != nil {
		log.Fatalf("Failed to start server. %s", err)
	}
}
