package main

import (
	"errors"
	"fmt"

	"html/template"
	// "log"
	"net/http"
	"strconv"

	// importamos todos los modelos
	"snippetbox.edgardev.net/internal/models"
)

func (app *Application) home(w http.ResponseWriter, r *http.Request) {

	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	snippets, err := app.snippets.Lastest()

	if err != nil {
		app.logger.Error(err.Error())
	}

	for _, currentSnippet := range snippets {
		fmt.Fprintf(w, "%+v\n", currentSnippet)
	}

	// 👇🏽 el archivo base debe ir primero
	// files := []string{"./ui/html/base.tmpl", "./ui/html/home.tmpl", "./ui/html/nav.html"}
	// templateSet, err := template.ParseFiles(files...) // 👈🏽 esto "..." es un varidic argument

	// if err != nil {
	// 	log.Println(err)
	// 	http.Error(w, "Oops internal error", http.StatusInternalServerError)
	// 	return
	// }

	// err = templateSet.ExecuteTemplate(w, "base", nil)

	// if err != nil {
	// 	log.Print(err)
	// 	http.Error(w, "Opps internal error", http.StatusInternalServerError)
	// }
}

func (app *Application) snippetView(w http.ResponseWriter, r *http.Request) {

	// strconv.Atoi() convierte un string a un numero int
	id, err := strconv.Atoi(r.URL.Query().Get("id"))

	if err != nil || id < 1 {
		http.NotFound(w, r)
		return
	}

	snippet, err := app.snippets.GetById(id)

	if err != nil {
		if errors.Is(err, models.ErrNoRecord) {
			http.NotFound(w, r)
		} else {
			app.logger.Error(err.Error())
		}
		return
	}

	// estos son los archivos estaticos que se enviaran, agrega todos los que se deben enviar
	// al cliente, desde esta ruta llamada snippetView

	files := []string{
		"./ui/html/base.html",
		"./ui/html/nav.html",
		"./ui/html/pages/view.html",
	}

	templateFiles, err := template.ParseFiles(files...)

	if err != nil {
		app.serverError(w, r, err)
		return
	}

	// toma todos los archivos estaticos (html o tmpl)
	// y los convierte en uno solo, para luego enviarlos
	// al cliente
	err = templateFiles.ExecuteTemplate(w, "base", snippet)

	if err != nil {
		app.serverError(w, r, err)
	}

	// Write the snippet data as a plain-text HTTP response body.
	// fmt.Fprintf(w, "%+v", snippet)

}

func (app *Application) snippetCreate(w http.ResponseWriter, r *http.Request) {

	if r.Method != "POST" {
		w.Header().Set("Allow", "POST")
		http.Error(w, "Method not Allowed", http.StatusMethodNotAllowed)
		return
	}

	title := "titulo de prueba"
	content := "este contenido es de prueba"
	expires := 8

	id, err := app.snippets.Insert(title, content, expires)

	if err != nil {
		app.logger.Error(err.Error())
		return
	}

	http.Redirect(w, r, fmt.Sprintf("/snippet/view?id=%d", id), http.StatusSeeOther)
}
