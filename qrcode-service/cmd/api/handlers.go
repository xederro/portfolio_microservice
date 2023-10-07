package main

import (
	"encoding/base64"
	"errors"
	"fmt"
	qrcode "github.com/skip2/go-qrcode"
	"net/http"
)

type QRCodeDetails struct {
	Link string `json:"link"`
	Size uint   `json:"size"`
}

type QRCode struct {
	PNG string `json:"png"`
}

func (a App) GetQRCode(w http.ResponseWriter, r *http.Request) {
	var q QRCodeDetails
	err := a.readJSON(w, r, &q)
	if err != nil {
		a.errorJSON(w, err)
		return
	}

	if q.Size > 1024 {
		q.Size = 1024
	} else if q.Size < 100 {
		q.Size = 100
	}

	if len(q.Link) == 0 {
		a.errorJSON(w, errors.New("link is needed"))
		return
	}

	var png []byte
	png, err = qrcode.Encode(q.Link, qrcode.Medium, int(q.Size))
	if err != nil {
		a.errorJSON(w, err)
		return
	}

	a.writeJSON(w, http.StatusAccepted, QRCode{PNG: fmt.Sprintf("data:image/png;base64,%s", base64.StdEncoding.EncodeToString(png))})
}
