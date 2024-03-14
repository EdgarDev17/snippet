package main

import (
	"fmt"
	"html/template"
	"path/filepath"
	"time"

	"snippetbox.edgardev.net/internal/models"
)

type TemplateData struct {
	Snippet     models.Snippet
	Snippets    []models.Snippet
	CurrentYear int
}

func HumanDateFormat(date time.Time) string {
	return date.Format("02 Jan 2005 at 15:00")
}

// Initialize a template.FuncMap object and store it in a global variable. This is
// essentially a string-keyed map which acts as a lookup between the names of our
// custom template functions and the functions themselves.
var functions = template.FuncMap{"HumanDateFormat": HumanDateFormat}

// la función se encarga de guardar los templates html en la cache
func newTemplateCache() (map[string]*template.Template, error) {

	// Initialize a new map to act as the cache.
	// key -> string
	// value -> Template
	cache := map[string]*template.Template{}

	// Use the filepath.Glob() function to get a slice of all filepaths that
	// match the pattern "./ui/html/pages/*.tmpl". This will essentially gives
	// us a slice of all the filepaths for our application 'page' templates
	// like: [ui/html/pages/home.tmpl ui/html/pages/view.tmpl]
	pages, err := filepath.Glob("./ui/html/pages/*.html")

	if err != nil {
		return nil, err
	}

	// Loop through the page filepaths one-by-one.
	for _, page := range pages {
		// Extract the file name (like 'home.tmpl') from the full filepath
		// and assign it to the name variable.
		name := filepath.Base(page)

		// The template.FuncMap must be registered with the template set before you
		// call the ParseFiles() method. This means we have to use template.New() to
		// create an empty template set, use the Funcs() method to register the
		// template.FuncMap, and then parse the file as normal.
		ts, err := template.New(name).Funcs(functions).ParseFiles("./ui/html/base.html")
		if err != nil {
			return nil, err
		}

		ts, err = ts.ParseGlob("./ui/html/partials/*.html")
		if err != nil {
			return nil, err
		}

		// Call ParseFiles() *on this template set* to add the
		ts, err = ts.ParseFiles(page)
		if err != nil {
			return nil, err
		}
		// Add the template set to the map, using the name of the page
		// (like 'home.html') as the key.
		cache[name] = ts
	}
	// Return the map.
	fmt.Println("this is how cache looks like: ", cache)
	return cache, nil
}
