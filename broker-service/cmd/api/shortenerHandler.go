package main

import (
	"errors"
	"github.com/go-chi/chi/v5"
	"github.com/xederro/portfolio/broker-service/cmd/shortener"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"net/http"
)

func (a App) GetAll(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("token")
	if err != nil {
		a.errorJSON(w, errors.New("There was an error while getting cookie"), http.StatusNoContent)
		return
	}

	q := shortener.ShortenerGetAllRequest{Token: cookie.Value}

	var opts []grpc.DialOption

	opts = append(opts, grpc.WithTransportCredentials(insecure.NewCredentials()))

	conn, err := grpc.Dial("shortener:8000", opts...)
	if err != nil {
		a.errorJSON(w, errors.New("There was an error while establishing connection"), http.StatusInternalServerError)
		return
	}
	defer conn.Close()

	client := shortener.NewShortenerServiceClient(conn)

	response, err := client.GetAll(context.TODO(), &q)
	if err != nil {
		//a.errorJSON(w, errors.New("There was an error"), http.StatusInternalServerError)
		a.errorJSON(w, err, http.StatusInternalServerError)
		return
	}

	a.writeJSON(w, http.StatusAccepted, response)
}

func (a App) GetOne(w http.ResponseWriter, r *http.Request) {
	q := shortener.ShortenerGetOneRequest{Short: chi.URLParam(r, "short")}

	var opts []grpc.DialOption

	opts = append(opts, grpc.WithTransportCredentials(insecure.NewCredentials()))

	conn, err := grpc.Dial("shortener:8000", opts...)
	if err != nil {
		a.errorJSON(w, errors.New("There was an error while establishing connection"), http.StatusInternalServerError)
		return
	}
	defer conn.Close()

	client := shortener.NewShortenerServiceClient(conn)

	response, err := client.GetOne(context.TODO(), &q)
	if err != nil {
		//a.errorJSON(w, errors.New("There was an error"), http.StatusInternalServerError)
		a.errorJSON(w, err, http.StatusInternalServerError)
		return
	}

	a.writeJSON(w, http.StatusAccepted, response)
}

func (a App) Insert(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("token")
	if err != nil {
		a.errorJSON(w, errors.New("There was an error while getting cookie"), http.StatusNoContent)
		return
	}

	q := shortener.ShortenerInsertRequest{}

	err = a.readJSON(w, r, &q)
	if err != nil || len(q.Long) <= 0 {
		a.errorJSON(w, errors.New("there was no link"), http.StatusNoContent)
		return
	}

	q.Token = cookie.Value

	var opts []grpc.DialOption

	opts = append(opts, grpc.WithTransportCredentials(insecure.NewCredentials()))

	conn, err := grpc.Dial("shortener:8000", opts...)
	if err != nil {
		a.errorJSON(w, errors.New("There was an error while establishing connection"), http.StatusInternalServerError)
		return
	}
	defer conn.Close()

	client := shortener.NewShortenerServiceClient(conn)

	response, err := client.Insert(context.TODO(), &q)
	if err != nil {
		//a.errorJSON(w, errors.New("There was an error"), http.StatusInternalServerError)
		a.errorJSON(w, err, http.StatusInternalServerError)
		return
	}

	a.writeJSON(w, http.StatusAccepted, response)
}

func (a App) Delete(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("token")
	if err != nil {
		a.errorJSON(w, errors.New("There was an error while getting cookie"), http.StatusNoContent)
		return
	}

	q := shortener.ShortenerDeleteRequest{Token: cookie.Value, Short: chi.URLParam(r, "short")}

	var opts []grpc.DialOption

	opts = append(opts, grpc.WithTransportCredentials(insecure.NewCredentials()))

	conn, err := grpc.Dial("shortener:8000", opts...)
	if err != nil {
		a.errorJSON(w, errors.New("There was an error while establishing connection"), http.StatusInternalServerError)
		return
	}
	defer conn.Close()

	client := shortener.NewShortenerServiceClient(conn)

	response, err := client.Delete(context.TODO(), &q)
	if err != nil {
		//a.errorJSON(w, errors.New("There was an error"), http.StatusInternalServerError)
		a.errorJSON(w, err, http.StatusInternalServerError)
		return
	}

	a.writeJSON(w, http.StatusAccepted, response)
}
