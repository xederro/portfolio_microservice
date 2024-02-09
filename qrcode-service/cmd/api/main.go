package main

import (
	"fmt"
	"github.com/xederro/portfolio/qrcode-service/cmd/qrcode"
	"google.golang.org/grpc"
	"log"
	"net"
)

const port = 8000

type App struct {
	qrcode.UnimplementedQRCodeServiceServer
}

func main() {
	lis, err := net.Listen("tcp", fmt.Sprintf("qrcode:%d", port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	var opts []grpc.ServerOption
	grpcServer := grpc.NewServer(opts...)

	qrcode.RegisterQRCodeServiceServer(grpcServer, App{})
	err = grpcServer.Serve(lis)
	if err != nil {
		log.Fatalln(err)
	}
}
