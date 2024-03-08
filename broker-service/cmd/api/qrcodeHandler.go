package main

import (
	"context"
	"errors"
	"github.com/xederro/portfolio/broker-service/cmd/qrcode"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"net/http"
)

func (a App) GetQRCode(w http.ResponseWriter, r *http.Request) {
	var q qrcode.QRCodeRequest
	err := a.readJSON(w, r, &q)
	if err != nil {
		a.errorJSON(w, err)
		return
	}

	q.Size = 256

	if q.Size > 1024 {
		q.Size = 1024
	} else if q.Size < 100 {
		q.Size = 100
	}

	if len(q.Link) == 0 {
		a.errorJSON(w, errors.New("link is needed"))
		return
	}

	var opts []grpc.DialOption

	opts = append(opts, grpc.WithTransportCredentials(insecure.NewCredentials()))

	conn, err := grpc.Dial("qrcode:8000", opts...)
	if err != nil {
		a.errorJSON(w, errors.New("There was an error while establishing connection"), http.StatusInternalServerError)
		return
	}
	defer conn.Close()

	client := qrcode.NewQRCodeServiceClient(conn)

	response, err := client.GetQRCode(context.TODO(), &q)
	if err != nil {
		a.errorJSON(w, errors.New("There was an error"), http.StatusInternalServerError)
		return
	}

	a.writeJSON(w, http.StatusAccepted, response)
}
