package main

import "net/http"

func (app Application) routes() *http.ServeMux {

	mux := http.NewServeMux()

	// crear un servidor de archivos estaticos, esto permite enviar archivos estaticosvia HTTP
	fileServerHandler := http.FileServer(http.Dir("./ui/static/"))

	// creating a mux help you to create routers its like a router
	mux.HandleFunc("/", app.home)
	mux.HandleFunc("/view", app.snippetView)
	mux.HandleFunc("/create", app.snippetCreate)

	// Use the mux.Handle() function to register the file server as the handler for
	// all URL paths that start with "/static/"
	mux.Handle("/static/", http.StripPrefix("/static", fileServerHandler))

	return mux
}
