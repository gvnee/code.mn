package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
)

func home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	w.Write([]byte("Hello"))
}

func showBlog(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil || id < 1 {
		http.NotFound(w, r)
		return
	}
	fmt.Fprintf(w, "Display a specific snippet with ID %d...", id)
}

func createBlog(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Create a new blog..."))
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /", home)
	mux.HandleFunc("GET /blog/{id}", showBlog)
	mux.HandleFunc("POST /blog/create", createBlog)

	log.Println("Starting server on :4000")
	err := http.ListenAndServe(":4000", mux)
	log.Fatal(err)
}
