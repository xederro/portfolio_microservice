package data

type Shortener struct {
	Short string `json:"short,omitempty"`
	Long  string `json:"long"`
}

type IWeather interface {
	GetOne(short string) (*Shortener, error)
	GetAll(uuid string) (*[]Shortener, error)
	Insert(token string) error
	Delete(short string, token string) error
}
