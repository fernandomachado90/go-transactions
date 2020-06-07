package main

import (
	"flag"
	"log"
	"net/http"
)

var (
	port = flag.String("port", "8888", "Informs the port where the API will be made available.")
)

func main() {
	flag.Parse()
	address := ":" + *port
	log.Printf("starting server @%s", address)

	server := &API{}

	if err := http.ListenAndServe(address, server.Routes()); err != nil {
		log.Fatalf("failed to start server. %s", err)
	}
}
