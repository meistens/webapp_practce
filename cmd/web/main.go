package main

import (
	"log"
	"net/http"
)

func main() {
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

	log.Println("\nstarting server on :4000")
	err := http.ListenAndServe(":4000", mux)
	log.Fatal(err)
}
