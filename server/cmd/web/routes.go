package main

import "net/http"

func (app *Application) routes() *http.ServeMux {
	// special kind of handler that passes requests on to other handlers
	mux := http.NewServeMux()

	// requests are handled concurrently!
	// HandleFunc: transforms given function to a handler and then registers it
	mux.HandleFunc("GET /", app.home)
	mux.HandleFunc("GET /blog/{id}", app.showBlog)
	mux.HandleFunc("POST /blog/create", app.createBlog)

	return mux
}
