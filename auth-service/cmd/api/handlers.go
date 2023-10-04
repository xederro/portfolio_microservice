package main

import (
	"github.com/xederro/portfolio/auth-service/data"
	"net/http"
)

func (a App) SignUp(w http.ResponseWriter, r *http.Request) {
	var m data.Model
	err := a.readJSON(w, r, &m)
	if err != nil {
		a.errorJSON(w, err)
		return
	}

	err = m.SignUp()
	if err != nil {
		a.errorJSON(w, err)
		return
	}

	a.writeJSON(w, http.StatusAccepted, jsonResponse{
		Error:   false,
		Message: "Signed Up",
	})
}

func (a *App) SignOut(w http.ResponseWriter, r *http.Request) {
	token, err := r.Cookie("Bearer")
	if err != nil {
		a.errorJSON(w, err)
		return
	}

	err = a.User.SignOut(token.Value)
	if err != nil {
		a.errorJSON(w, err)
		return
	}

	a.writeJSON(w, http.StatusAccepted, jsonResponse{
		Error:   false,
		Message: "Got Model",
	})
}

func (a App) SignIn(w http.ResponseWriter, r *http.Request) {
	var u data.Model
	err := a.readJSON(w, r, &u)
	if err != nil {
		a.errorJSON(w, err)
		return
	}

	t, err := u.SignIn()
	if err != nil {
		a.errorJSON(w, err)
		return
	}

	a.writeJSON(w, http.StatusAccepted, jsonResponse{
		Error:   false,
		Message: "Signed In",
		Data:    t,
	})
}

func (a App) GetOne(w http.ResponseWriter, r *http.Request) {
	token, err := r.Cookie("Bearer")
	if err != nil {
		a.errorJSON(w, err)
		return
	}

	up, err := a.User.GetOne(token.Value)
	if err != nil {
		a.errorJSON(w, err)
		return
	}

	a.writeJSON(w, http.StatusAccepted, jsonResponse{
		Error:   false,
		Message: "Got Model",
		Data:    up,
	})
}

func (a App) Update(w http.ResponseWriter, r *http.Request) {
	token, err := r.Cookie("Bearer")
	if err != nil {
		a.errorJSON(w, err)
		return
	}

	var m data.Model
	err = a.readJSON(w, r, &m)
	if err != nil {
		a.errorJSON(w, err)
		return
	}

	err = m.Update(token.Value, m.UpdateData)
	if err != nil {
		a.errorJSON(w, err)
		return
	}

	a.writeJSON(w, http.StatusAccepted, jsonResponse{
		Error:   false,
		Message: "Updated",
	})
}

func (a App) Delete(w http.ResponseWriter, r *http.Request) {
	var m data.Model
	err := a.readJSON(w, r, &m)
	if err != nil {
		a.errorJSON(w, err)
		return
	}

	err = m.Delete()
	if err != nil {
		a.errorJSON(w, err)
		return
	}

	a.writeJSON(w, http.StatusAccepted, jsonResponse{
		Error:   false,
		Message: "Deleted",
	})
}
