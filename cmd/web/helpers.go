package main

import (
	"bytes"
	"fmt"
	"net/http"
	"time"
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

	// en lugar de ejecutar y enviar el archivo html directamente
	// lo mejor seria que los guardemos en un buffer
	// si estamos libres de errores, tomamos la pagina del buffer y la enviamos
	// al cliente
	buff := new(bytes.Buffer)
	err := template.ExecuteTemplate(buff, "base", data)

	if err != nil {
		app.serverError(w, r, err)
		return
	}

	w.WriteHeader(status)

	// tomamos lo que hay dentro del buffer
	// lo enviamos(escribimos) en el parametro "w"
	// el cual es el encargado de enviar respuestas al cliente
	buff.WriteTo(w)
}

func (app Application) newTemplateData(r *http.Request) TemplateData {

	return TemplateData{
		CurrentYear: time.Now().Year(),
	}
}
