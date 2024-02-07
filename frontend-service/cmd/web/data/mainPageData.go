package data

import (
	"encoding/json"
	"os"
)

type MainPageData struct {
	Education []struct {
		Name   string   `json:"name"`
		Degree string   `json:"degree"`
		Major  string   `json:"major"`
		Year   string   `json:"year"`
		Place  string   `json:"place"`
		Bullet []string `json:"bullet,omitempty"`
	} `json:"education"`
	Projects []struct {
		ProjectName  string   `json:"projectName"`
		Lang         string   `json:"lang"`
		LangColor    string   `json:"langColor,omitempty"`
		Description  string   `json:"description"`
		Link         string   `json:"link"`
		Technologies []string `json:"technologies"`
	} `json:"projects"`
}

func (m *MainPageData) Populate() error {
	file, err := os.ReadFile("./cmd/web/data/mainPageData.json")
	if err != nil {
		return err
	}

	err = json.Unmarshal(file, m)
	if err != nil {
		return err
	}

	return nil
}
