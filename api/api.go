package main

import "github.com/go-chi/chi"

type API struct {
}

func (api *API) Routes() *chi.Mux {
	mux := chi.NewRouter()

	mux.Get("/healthcheck", api.handleHealthCheck())

	return mux
}
