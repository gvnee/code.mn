package main

import (
	"flag"
	"log"
	"net/http"
	"os"
)

// application wide dependencies
type Application struct {
	errorLog *log.Logger
	infoLog  *log.Logger
}

type Config struct {
	Addr string
}

func main() {
	cfg := new(Config)
	flag.StringVar(&cfg.Addr, "addr", ":4000", "HTTP network address")
	flag.Parse()

	// could redirect logs for archival
	infoLog := log.New(os.Stdout, "[INFO]\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stderr, "[ERROR]\t", log.Ldate|log.Ltime|log.Lshortfile)

	app := &Application{
		errorLog: errorLog,
		infoLog:  infoLog,
	}

	// special kind of handler that passes requests on to other handlers
	mux := http.NewServeMux()

	// requests are handled concurrently!
	// HandleFunc: transforms given function to a handler and then registers it
	mux.HandleFunc("GET /", app.home)
	mux.HandleFunc("GET /blog/{id}", app.showBlog)
	mux.HandleFunc("POST /blog/create", app.createBlog)

	// to use custom error logger
	srv := &http.Server{
		Addr:     cfg.Addr,
		ErrorLog: errorLog,
		Handler:  mux,
	}

	infoLog.Printf("Starting server on %s", cfg.Addr)
	err := srv.ListenAndServe()
	errorLog.Fatal(err)
}
