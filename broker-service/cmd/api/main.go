package main

import (
	"fmt"
	"log"
	"net/http"
)

const webPort = "3000"

type App struct {
}

func main() {
	app := App{}

	srv := &http.Server{
		Addr:    fmt.Sprintf(":%s", webPort),
		Handler: app.routes(),
	}

	err := srv.ListenAndServe()
	if err != nil {
		log.Panic(err)
	}
}
