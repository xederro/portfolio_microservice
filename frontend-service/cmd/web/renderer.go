package main

import (
	"fmt"
	"html/template"
	"net/http"
)

var layout = []string{
	"./cmd/web/templates/layout/main.layout.gohtml",
	"./cmd/web/templates/partials/header.partial.gohtml",
	"./cmd/web/templates/partials/head.partial.gohtml",
}

type Data struct {
	PageName         string
	SiteAddress      string
	PageSpecificData any
}

func (a App) page(w http.ResponseWriter, page string, data *Data, isHTMX bool) {

	var templateData []string
	templateData = append(templateData, fmt.Sprintf("./cmd/web/templates/pages/%s", page))

	if !isHTMX {
		for _, x := range layout {
			templateData = append(templateData, x)
		}
	} else {
		templateData = append(templateData, fmt.Sprintf("./cmd/web/templates/layout/htmx.layout.gohtml"))
	}

	tmpl, err := template.ParseFiles(templateData...)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := tmpl.Execute(w, data); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func (a App) error(w http.ResponseWriter, code int, isHTMX bool) {
	var templateData []string
	templateData = append(templateData, fmt.Sprintf("./cmd/web/templates/pages/error.page.gohtml"))

	if !isHTMX {
		for _, x := range layout {
			templateData = append(templateData, x)
		}
	} else {
		templateData = append(templateData, fmt.Sprintf("./cmd/web/templates/layout/htmx.layout.gohtml"))
	}

	tmpl, err := template.ParseFiles(templateData...)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(code)
	if err := tmpl.Execute(w, code); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
