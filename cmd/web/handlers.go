package main

import (
	"errors"
	"fmt"

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

	data := app.newTemplateData(r)
	data.Snippets = snippets

	app.render(w, r, http.StatusOK, "home.html", data)
}

// se encarga de manejar la pagina donde se muestra un snippet individual
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

	data := app.newTemplateData(r)
	data.Snippet = snippet
	app.render(w, r, http.StatusOK, "view.html", data)

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
