package data

import (
	"context"
	"firebase.google.com/go/v4"
	"firebase.google.com/go/v4/auth"
	"google.golang.org/api/option"
	"log"
	"os"
)

var app *firebase.App
var authClient *auth.Client

func NewUser() Model {
	opt := option.WithCredentialsFile(os.Getenv("FirebaseConfigPath"))
	config := &firebase.Config{ProjectID: os.Getenv("FirebaseProjectID")}
	newApp, err := firebase.NewApp(context.Background(), config, opt)
	if err != nil {
		log.Fatalf("error initializing app: %v\n", err)
	}
	user, err := app.Auth(context.Background())
	if err != nil {
		log.Fatalf("error initializing app: %v\n", err)
	}

	app = newApp
	authClient = user

	return Model{}
}

func (m Model) SignUp() error {
	ctx := context.TODO()
	params := (&auth.UserToCreate{}).
		Email(m.Credentials.Email).
		EmailVerified(false).
		Password(m.Credentials.Password).
		Disabled(false)
	_, err := authClient.CreateUser(ctx, params)
	if err != nil {
		return err
	}
	return nil
}

func (m Model) Update(token string) error {
	uuid, err := m.GetUUID(token)
	if err != nil {
		return err
	}

	params := (&auth.UserToUpdate{}).
		Password(m.Credentials.Password).
		DisplayName(m.Auth.DisplayName)
	_, err = authClient.UpdateUser(context.TODO(), uuid, params)
	if err != nil {
		return err
	}

	return nil
}

func (m Model) Delete(token string) error {
	uuid, err := m.GetUUID(token)
	if err != nil {
		return err
	}

	err = authClient.DeleteUser(context.TODO(), uuid)
	if err != nil {
		return err
	}

	return nil
}

func (m Model) GetUUID(token string) (string, error) {
	t, err := authClient.VerifyIDToken(context.TODO(), token)
	if err != nil {
		return "", err
	}

	return t.UID, nil
}
