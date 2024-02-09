package main

import (
	"errors"
	"fmt"

	// "html/template"
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

	// Write the snippet data as a plain-text HTTP response body.
	fmt.Fprintf(w, "%+v", snippet)

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
