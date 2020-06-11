package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"

	"github.com/fernandomachado90/go-transactions/core"
	"github.com/fernandomachado90/go-transactions/database"
)

const address = ":8080"

func newServer() *http.Server {
	log.Printf("Starting database")
	db, err := database.NewSQLite()
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
	server.RegisterOnShutdown(func() {
		log.Printf("Closing database connection")
		_ = db.Close()
	})
	return server
}

func main() {
	server := newServer()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	go func() {
		<-quit
		log.Printf("Shutting down server connection")
		_ = server.Shutdown(context.Background())
	}()

	if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatalf("Failed to start server. %s", err)
	}

	log.Printf("Application ended successfully")
}
