package data

import (
	"cloud.google.com/go/firestore"
	"context"
	"errors"
	"firebase.google.com/go/v4"
	"fmt"
	"google.golang.org/api/iterator"
	"google.golang.org/api/option"
	"log"
	"os"
	"time"
)

var db *firestore.Client

func NewWeather() Weather {
	opt := option.WithCredentialsFile(os.Getenv("FirebaseConfigPath"))
	config := &firebase.Config{ProjectID: os.Getenv("FirebaseProjectID")}
	newApp, err := firebase.NewApp(context.Background(), config, opt)
	if err != nil {
		log.Fatalf("error initializing app: %v\n", err)
	}

	database, err := newApp.Firestore(context.Background())
	if err != nil {
		log.Fatalf("error initializing database: %v\n", err)
	}
	db = database
	defer func(db *firestore.Client) {
		err := db.Close()
		if err != nil {

		}
	}(db)

	return Weather{}
}

func (w Weather) GetLast() (*Weather, error) {
	var results Weather
	docs := db.Collection("/weather").
		Limit(1).
		Select("*").
		OrderBy("timestamp", firestore.Desc).
		Documents(context.TODO())

	for {
		doc, err := docs.Next()
		if errors.Is(err, iterator.Done) {
			break
		}
		if err != nil {
			return nil, err
		}
		fmt.Println(doc.Data())
	}

	docs.Stop()

	return &results, nil
}

func (w Weather) GetTimeWindow(start time.Time, stop time.Time, span uint8) (*[]Weather, error) {
	var results []Weather
	err := db.DB.
		From("weather").
		Select("*").
		Filter("order", "timestamp", "desc").
		Eq("location", "KSW").
		Lte("timestamp", start.Format("2006-01-02T15:04:05")).
		Gte("timestamp", stop.Format("2006-01-02T15:04:05")).
		Execute(results)
	if err != nil {
		return nil, err
	}

	return &results, nil
}

func (w Weather) Insert(cred string) error {
	if cred == os.Getenv("WeatherKey") {
		var m []Weather
		err := db.DB.From("weather").Insert(w).Execute(&m)
		if err != nil {
			return err
		}

		return nil
	}
	return errors.New("unauthenticated Insert")
}
