package main

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
	"net/http"
)

func (a App) routes() http.Handler {
	mux := chi.NewRouter()

	mux.Use(cors.Handler(cors.Options{
		AllowedOrigins:     []string{"https://*", "http://*"},
		AllowedMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:     []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token", "hx-current-url", "hx-request"},
		ExposedHeaders:     []string{"Link"},
		AllowCredentials:   true,
		MaxAge:             300,
		AllowOriginFunc:    nil,
		OptionsPassthrough: false,
		Debug:              false,
	}))

	mux.Post("/", a.GetQRCode)
	mux.Options("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNoContent)
	})

	return mux
}
