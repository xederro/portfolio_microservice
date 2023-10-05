package data

import (
	"context"
	"fmt"
	"github.com/nedpals/supabase-go"
	"log"
	"os"
)

var db *supabase.Client

func NewShortener() Shortener {
	NewDB := supabase.CreateClient(os.Getenv("SupabaseUrl"), os.Getenv("SupabaseKey"))

	if NewDB == nil {
		log.Panic("Cant connect to Database")
	}

	db = NewDB

	return Shortener{}
}

func (s Shortener) GetOne(short string) (*Shortener, error) {
	var results Shortener
	err := db.DB.
		From("shortener").
		Select("long").
		Limit(1).
		Single().
		Eq("short", short).
		Execute(&results)
	if err != nil {
		return nil, err
	}

	return &results, nil
}

func (s Shortener) GetAll(uuid string) (*[]Shortener, error) {
	var results []Shortener
	err := db.DB.
		From("shortener").
		Select("long").
		Limit(1).
		Eq("user_id", uuid).
		Execute(&results)
	if err != nil {
		return nil, err
	}

	return &results, nil
}

func (s Shortener) Insert(token string) error {
	var m map[string]interface{}

	db.DB.Headers().Set("Authorization", fmt.Sprintf("Bearer %s", token))
	err := db.DB.From("shortener").Insert(s).Execute(&m)
	if err != nil {
		return err
	}

	return nil
}

func (s Shortener) Delete(short string, token string) error {
	db.DB.Headers().Set("Authorization", fmt.Sprintf("Bearer %s", token))
	var m map[string]interface{}
	err := db.DB.From("shortener").Delete().Eq("short", short).Execute(&m)
	if err != nil {
		return err
	}

	return nil
}

func GetUUID(token string) (string, error) {
	ctx := context.TODO()
	user, err := db.Auth.User(ctx, token)
	if err != nil {
		return "", err
	}

	return user.ID, nil
}
