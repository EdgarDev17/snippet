package main

import "net/http"

func (app Application) routes() http.Handler {

	// al instanciar un mux, este nos permite crear rutas y mas.
	mux := http.NewServeMux()

	// crear un servidor de archivos estaticos, esto permite enviar archivos estaticos via HTTP
	fileServerHandler := http.FileServer(http.Dir("./ui/static/"))

	// creating a mux help you to create routers its like a router
	mux.HandleFunc("/", app.home)
	mux.HandleFunc("/snippet/view", app.snippetView)
	mux.HandleFunc("/create", app.snippetCreate)

	// Use the mux.Handle() function to register the file server as the handler for
	// all URL paths that start with "/static/"
	mux.Handle("/static/", http.StripPrefix("/static", fileServerHandler))

	// Cubrimos nuestra cadena(chain) existente con el nuevo middleware logginRequest
	return app.logginRequest(secureHandlers(mux))
}
