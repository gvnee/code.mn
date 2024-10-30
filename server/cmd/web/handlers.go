package main

import (
	"fmt"
	"net/http"
	"strconv"
)

func (app *Application) home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	w.Write([]byte("Hello"))
}

func (app *Application) showBlog(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil || id < 1 {
		http.NotFound(w, r)
		return
	}
	fmt.Fprintf(w, "Display a specific snippet with ID %d...", id)
}

func (app *Application) createBlog(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Create a new blog..."))
}
