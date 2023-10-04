package main

import (
	"net/http"
	"time"
)

type timeWindow struct {
	Start time.Time `json:"start"`
	Stop  time.Time `json:"stop"`
	Span  uint8     `json:"span"`
}

func (a App) Last(w http.ResponseWriter, r *http.Request) {
	last, err := a.Weather.GetLast()
	if err != nil {
		a.errorJSON(w, err)
		return
	}

	a.writeJSON(w, http.StatusAccepted, last)
}

func (a App) Get(w http.ResponseWriter, r *http.Request) {
	var tw timeWindow
	err := a.readJSON(w, r, &tw)
	if err != nil {
		a.errorJSON(w, err)
		return
	}

	get, err := a.Weather.GetTimeWindow(tw.Start, tw.Stop, tw.Span)
	if err != nil {
		a.errorJSON(w, err)
		return
	}

	a.writeJSON(w, http.StatusAccepted, get)
}
