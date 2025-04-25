package main

import (
	"flag"
	"log"
	"net/http"
	"os"
)

// dependency injection, might be different if using imported packages
// basically, struct holds app-wide dependencies to be used
type application struct {
	errLog  *log.Logger
	infoLog *log.Logger
}

func main() {
	// define a cmd-line flag with name addr, a default value of 4000
	// and some help text
	// parse command to... its obvious what it does...
	addr := flag.String("addr", ":4000", "http network address")
	flag.Parse()

	// log.New() for creating a logger for wrting info. msgs
	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	// same as infoLog, but stderr is the dest.
	errLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	// init. a new instance of the application struct
	app := &application{
		errLog:  errLog,
		infoLog: infoLog,
	}

	// swap the route declarations to use the application struct method as the
	// handler function
	mux := http.NewServeMux()

	// create a file server which uses files from the static dir.
	fileServer := http.FileServer(http.Dir("./ui/static/"))
	// use the mux.Handle() to register the file server as the handler
	// for all url path that starts with "/static/"
	// for matching paths, strip the /static before the req reaches the fileserver
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))
	mux.HandleFunc("/", app.home)
	mux.HandleFunc("/snippet/view", app.snippetView)
	mux.HandleFunc("/snippet/create", app.snippetCreate)

	// init. a new http.Server struct
	// setup addr and handler so the server uses the same network as before
	// set errLog so server uses custom error logger
	srv := &http.Server{
		Addr:     *addr,
		ErrorLog: errLog,
		Handler:  mux,
	}

	infoLog.Printf("\nstarting server on %s", *addr)
	err := srv.ListenAndServe()
	errLog.Fatal(err)
}
