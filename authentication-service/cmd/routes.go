package cmd

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func (app *Config) Routes() http.Handler {
	mux := chi.NewRouter()

	mux.Use(middleware.Heartbeat("/ping"))
	mux.Post("/authenticate", app.Authenticate)
	mux.Get("/users", app.GetAllUsersHandler)

	return mux
}
