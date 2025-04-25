package main

import (
	"flag"
	"log"
	"net/http"
)

func main() {
	// define a cmd-line flag with name addr, a default value of 4000
	// and some help text
	// parse command to... its obvious what it does...
	addr := flag.String("addr", ":4000", "http network address")
	flag.Parse()

	mux := http.NewServeMux()

	// create a file server which uses files from the static dir.
	fileServer := http.FileServer(http.Dir("./ui/static/"))
	// use the mux.Handle() to register the file server as the handler
	// for all url path that starts with "/static/"
	// for matching paths, strip the /static before the req reaches the fileserver
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))
	mux.HandleFunc("/", home)
	mux.HandleFunc("/snippet/view", snippetView)
	mux.HandleFunc("/snippet/create", snippetCreate)

	log.Printf("\nstarting server on %s", *addr)
	err := http.ListenAndServe(*addr, mux)
	log.Fatal(err)
}
