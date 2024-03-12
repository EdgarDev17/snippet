package main

import (
	"fmt"
	"html/template"
	"path/filepath"

	"snippetbox.edgardev.net/internal/models"
)

type TemplateData struct {
	Snippet     models.Snippet
	Snippets    []models.Snippet
	CurrentYear int
}

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

		ts, err := template.ParseFiles("./ui/html/base.html")
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
