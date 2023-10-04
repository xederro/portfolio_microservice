package main

import (
	"fmt"
	"github.com/xederro/portfolio/weather-service/data"
	"log"
	"net/http"
)

const webPort = "8002"

type App struct {
	Weather data.Weather
}

func main() {
	app := App{
		Weather: data.NewWeather(),
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
