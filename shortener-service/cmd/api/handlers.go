package main

import (
	"github.com/go-chi/chi/v5"
	"github.com/xederro/portfolio/shortener-service/data"
	"net/http"
)

func (a App) GetAll(w http.ResponseWriter, r *http.Request) {
	token, err := r.Cookie("Bearer")
	if err != nil {
		a.errorJSON(w, err)
		return
	}

	uuid, err := data.GetUUID(token.Value)
	if err != nil {
		a.errorJSON(w, err, http.StatusUnauthorized)
		return
	}

	all, err := a.Shortener.GetAll(uuid)
	if err != nil {
		a.errorJSON(w, err)
		return
	}

	a.writeJSON(w, http.StatusAccepted, all)
}

func (a App) GetOne(w http.ResponseWriter, r *http.Request) {
	short := chi.URLParam(r, "short")

	one, err := a.Shortener.GetOne(short)
	if err != nil {
		a.errorJSON(w, err)
		return
	}

	a.writeJSON(w, http.StatusAccepted, one)
}

func (a App) Insert(w http.ResponseWriter, r *http.Request) {
	token, err := r.Cookie("Bearer")
	if err != nil {
		a.errorJSON(w, err)
		return
	}

	var s data.Shortener
	err = a.readJSON(w, r, &s)
	if err != nil {
		a.errorJSON(w, err)
		return
	}

	err = s.Insert(token.Value)
	if err != nil {
		a.errorJSON(w, err)
		return
	}

	a.writeJSON(w, http.StatusAccepted, nil)
}

func (a App) Delete(w http.ResponseWriter, r *http.Request) {
	short := chi.URLParam(r, "short")

	token, err := r.Cookie("Bearer")
	if err != nil {
		a.errorJSON(w, err)
		return
	}

	err = a.Shortener.Delete(short, token.Value)
	if err != nil {
		a.errorJSON(w, err)
		return
	}

	a.writeJSON(w, http.StatusAccepted, nil)
}
