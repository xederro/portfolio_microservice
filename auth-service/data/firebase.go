package data

import (
	"context"
	"firebase.google.com/go/v4"
	"firebase.google.com/go/v4/auth"
	RPCAuth "github.com/xederro/portfolio/auth-service/cmd/auth"
	"google.golang.org/api/option"
	"log"
	"os"
)

type Model struct {
	Auth auth.UserInfo `json:"auth,omitempty"`

	RPCAuth.UnimplementedAuthServiceServer
}

var app *firebase.App
var authClient *auth.Client

func NewUser() Model {
	opt := option.WithCredentialsFile(os.Getenv("FirebaseConfigPath"))
	config := &firebase.Config{ProjectID: os.Getenv("FirebaseProjectID")}
	newApp, err := firebase.NewApp(context.Background(), config, opt)
	if err != nil {
		log.Fatalf("error initializing app: %v\n", err)
	}
	user, err := newApp.Auth(context.Background())
	if err != nil {
		log.Fatalf("error initializing app: %v\n", err)
	}

	app = newApp
	authClient = user

	return Model{}
}

func (m Model) GetUUID(token string) (string, error) {
	t, err := authClient.VerifyIDToken(context.TODO(), token)
	if err != nil {
		return "", err
	}

	return t.UID, nil
}

func (m Model) CheckAuth(ctx context.Context, req *RPCAuth.AuthRequest) (*RPCAuth.AuthResponse, error) {
	_, err := authClient.VerifyIDToken(context.TODO(), req.Token)
	if err != nil {
		return &RPCAuth.AuthResponse{
			IsAuth: false,
		}, err
	}

	return &RPCAuth.AuthResponse{
		IsAuth: true,
	}, nil
}
