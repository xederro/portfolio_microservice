package data

import "time"

type Weather struct {
	Timestamp time.Time `json:"timestamp"`
	Location  string    `json:"location"`
	Temp      float32   `json:"temp"`
	Hum       float32   `json:"hum"`
	Press     float32   `json:"press"`
	Light     float32   `json:"light"`
}

type IWeather interface {
	GetLast() (*Weather, error)
	GetTimeWindow(start time.Time, stop time.Time, span uint8) (*[]Weather, error)
	Insert(cred string) error
}
