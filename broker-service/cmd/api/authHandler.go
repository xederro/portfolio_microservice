package main

import (
	"context"
	"errors"
	"fmt"
	"github.com/xederro/broker-service/cmd/auth"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"net/http"
)

func (a App) CheckAuth(w http.ResponseWriter, r *http.Request) {
	var q auth.AuthRequest
	cookie, err := r.Cookie("token")
	if err != nil {
		a.errorJSON(w, errors.New("There was an error while getting cookie"), http.StatusNoContent)
		return
	}
	q.Token = cookie.Value

	fmt.Println(q.Token)

	var opts []grpc.DialOption

	opts = append(opts, grpc.WithTransportCredentials(insecure.NewCredentials()))

	conn, err := grpc.Dial("auth:8000", opts...)
	if err != nil {
		a.errorJSON(w, errors.New("There was an error while establishing connection"), http.StatusInternalServerError)
		return
	}
	defer conn.Close()

	client := auth.NewAuthServiceClient(conn)

	response, err := client.CheckAuth(context.TODO(), &q)
	if err != nil {
		//a.errorJSON(w, errors.New("There was an error"), http.StatusInternalServerError)
		a.errorJSON(w, err, http.StatusInternalServerError)
		return
	}

	a.writeJSON(w, http.StatusAccepted, response)
}
