package main

import (
	"fmt"
	"github.com/xederro/portfolio/shortener-service/cmd/shortener"
	"google.golang.org/grpc"
	"log"
	"net"
)

const port = 8000

func main() {
	app := App{}

	lis, err := net.Listen("tcp", fmt.Sprintf("shortener:%d", port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	var opts []grpc.ServerOption
	grpcServer := grpc.NewServer(opts...)

	shortener.RegisterShortenerServiceServer(grpcServer, app)
	err = grpcServer.Serve(lis)
	if err != nil {
		log.Fatalln(err)
	}
}
