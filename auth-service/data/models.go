package data

import (
	"context"
	"fmt"
	"github.com/nedpals/supabase-go"
	"log"
	"os"
)

var db *supabase.Client

func NewUser() Model {
	NewDB := supabase.CreateClient(os.Getenv("SupabaseUrl"), os.Getenv("SupabaseKey"))

	if NewDB == nil {
		log.Panic("Cant connect to Database")
	}

	db = NewDB

	return Model{}
}

func (m Model) SignUp() error {
	ctx := context.TODO()
	_, err := db.Auth.SignUp(ctx, supabase.UserCredentials{
		Email:    m.Credentials.Email,
		Password: m.Credentials.Password,
	})
	if err != nil {
		return err
	}

	return nil
}

func (m Model) SignOut(token string) error {
	err := db.Auth.SignOut(context.Background(), token)
	if err != nil {
		return err
	}

	return nil
}

func (m Model) SignIn() (string, error) {
	ctx := context.TODO()
	user, err := db.Auth.SignIn(ctx, supabase.UserCredentials{
		Email:    m.Credentials.Email,
		Password: m.Credentials.Password,
	})
	if err != nil {
		return "", err
	}

	return user.AccessToken, nil
}

func (m Model) GetAll() ([]*supabase.User, error) {
	var results []*supabase.User
	err := db.DB.From("users").Select("*").Execute(&results)
	if err != nil {
		return nil, err
	}

	return results, nil
}

func (m Model) GetByEmail(email string) (*supabase.User, error) {
	var results supabase.User
	err := db.DB.From("users").Select("*").Single().Eq("email", email).Execute(&results)
	if err != nil {
		return nil, err
	}

	return &results, nil
}

func (m Model) GetOne(token string) (*supabase.User, error) {
	user, err := db.Auth.User(context.Background(), token)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (m Model) Update(token string, updateData map[string]interface{}) error {
	_, err := db.Auth.UpdateUser(context.TODO(), token, updateData)
	if err != nil {
		return err
	}

	return nil
}

func (m Model) Delete() error {
	var results map[string]interface{}
	err := db.DB.From("users").Delete().Eq("id", m.Auth.User.ID).Execute(&results)
	if err != nil {
		return err
	}

	fmt.Println(results)
	return nil
}

func (m Model) DeleteByID(id string) error {
	var results map[string]interface{}
	err := db.DB.From("users").Delete().Eq("id", id).Execute(&results)
	if err != nil {
		return err
	}

	fmt.Println(results)
	return nil
}

func (m Model) ResetPassword(email string) error {
	err := db.Auth.ResetPasswordForEmail(context.Background(), email)
	if err != nil {
		return err
	}
	return nil
}

func (m Model) GetUUID(token string) (string, error) {
	ctx := context.TODO()
	user, err := db.Auth.User(ctx, token)
	if err != nil {
		return "", err
	}

	return user.ID, nil
}
