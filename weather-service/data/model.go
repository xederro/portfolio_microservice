package data

import (
	"errors"
	"github.com/nedpals/supabase-go"
	"log"
	"os"
	"time"
)

var db *supabase.Client

func NewWeather() Weather {
	NewDB := supabase.CreateClient(os.Getenv("SupabaseUrl"), os.Getenv("SupabaseKey"))

	if NewDB == nil {
		log.Panic("Cant connect to Database")
	}

	db = NewDB

	return Weather{}
}

func (w Weather) GetLast() (*Weather, error) {
	var results Weather
	err := db.DB.
		From("weather").
		Select("*").
		Limit(1).
		Single().
		Filter("order", "timestamp", "desc").
		Execute(&results)
	if err != nil {
		return nil, err
	}

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
