package main

import (
	"fmt"
	"github.com/xederro/portfolio/auth-service/cmd/auth"
	"github.com/xederro/portfolio/auth-service/data"
	"google.golang.org/grpc"
	"log"
	"net"
)

const port = "8000"

type App struct {
	User data.Model
}

func main() {
	app := App{
		User: data.NewUser(),
	}

	lis, err := net.Listen("tcp", fmt.Sprintf("auth:%s", port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	var opts []grpc.ServerOption
	grpcServer := grpc.NewServer(opts...)

	auth.RegisterAuthServiceServer(grpcServer, app.User)
	err = grpcServer.Serve(lis)
	if err != nil {
		log.Fatalln(err)
	}
}
