package main

import (
	"fmt"
	"github.com/xederro/auth-service/data"
	"log"
	"net/http"
)

const webPort = "8001"

type App struct {
	User data.IModel
}

func main() {
	app := App{
		User: data.NewUser(),
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
