package main

import (
	"fmt"
	"net/http"
)

func (app *Application) serverError(w http.ResponseWriter, r *http.Request, err error) {

	var (
		method = r.Method
		uri    = r.URL.RequestURI()
	)

	app.logger.Error(err.Error(), "method", method, "url", uri)
	http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
}

func (app Application) render(w http.ResponseWriter, r *http.Request, status int, page string, data TemplateData) {

	template, ok := app.templateCache[page]

	if !ok {
		err := fmt.Errorf("the template %s does not exist", page)
		app.serverError(w, r, err)
		return
	}

	w.WriteHeader(status)

	err := template.ExecuteTemplate(w, "base", data)

	if err != nil {
		app.serverError(w, r, err)
	}
}
