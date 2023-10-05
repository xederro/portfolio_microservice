package main

import (
	"fmt"
	"github.com/xederro/portfolio/shortener-service/data"
	"log"
	"net/http"
)

const webPort = "8003"

type App struct {
	Shortener data.Shortener
}

func main() {
	app := App{
		Shortener: data.NewShortener(),
	}

	srv := &http.Server{
		Addr:    fmt.Sprintf(":%s", webPort),
		Handler: app.routes(),
	}

	err := srv.ListenAndServe()
	if err != nil {
		log.Panic(err)
	}
}
