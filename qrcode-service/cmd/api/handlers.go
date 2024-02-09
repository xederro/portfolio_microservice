package main

import (
	"context"
	"encoding/base64"
	"errors"
	"fmt"
	"github.com/skip2/go-qrcode"
	qrcoderpc "github.com/xederro/portfolio/qrcode-service/cmd/qrcode"
)

func (a App) GetQRCode(ctx context.Context, q *qrcoderpc.QRCodeRequest) (*qrcoderpc.QRCodeResponse, error) {
	if q.Size > 1024 {
		q.Size = 1024
	} else if q.Size < 100 {
		q.Size = 100
	}

	if len(q.Link) == 0 {
		return nil, errors.New("No Link")
	}

	png, err := qrcode.Encode(q.Link, qrcode.Medium, int(q.Size))
	if err != nil {
		return nil, errors.New("There was an error while encoding QRCode")
	}

	return &qrcoderpc.QRCodeResponse{PNG: fmt.Sprintf("data:image/png;base64,%s", base64.StdEncoding.EncodeToString(png))}, nil
}
