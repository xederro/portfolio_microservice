package main

import (
	"fmt"
	"github.com/angelofallars/htmx-go"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
	"github.com/xederro/portfolio/frontend-service/cmd/web/data"
	"net/http"
)

func (a App) routes() http.Handler {
	mux := chi.NewRouter()
	fs := http.FileServer(http.Dir("./cmd/web/public"))

	mux.Use(cors.Handler(cors.Options{
		AllowedOrigins:     []string{"https://*", "http://*"},
		AllowedMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:     []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:     []string{"Link"},
		AllowCredentials:   true,
		MaxAge:             300,
		AllowOriginFunc:    nil,
		OptionsPassthrough: false,
		Debug:              false,
	}))

	mux.Get("/", func(w http.ResponseWriter, r *http.Request) {
		d := data.MainPageData{}
		err := d.Populate()
		if err != nil {
			a.error(w, http.StatusInternalServerError, htmx.IsHTMX(r))
			fmt.Println(err)
		}

		a.page(w, "main.page.gohtml", &Data{
			PageName: "Home",
			//SiteAddress:      os.Getenv("PageAddress"),
			SiteAddress:      "xederro.pl",
			PageSpecificData: d,
		}, htmx.IsHTMX(r))
	})

	mux.Handle("/public/*", http.StripPrefix("/public/", fs))

	mux.NotFound(func(w http.ResponseWriter, r *http.Request) {
		a.error(w, http.StatusNotFound, htmx.IsHTMX(r))
	})

	return mux
}
