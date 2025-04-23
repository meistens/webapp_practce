package main

import (
	"log"
	"net/http"
)

// defne a home handler function which writes a byte slice containing
// "some words" as the response body
func home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("some words"))
}

func main() {
	// use the http.NewServerMux() func to init a new server
	// then register the home func as the handler for the "/" url pattern
	mux := http.NewServeMux()
	mux.HandleFunc("/", home)

	// use the http.ListenAndServe() func to start a new web server
	// pass in 2 params: TCP to listen on and ServeMux just created
	// if listenandserve returns an error, log.fatal() to log error msg and exit
	// Note-> any err returned by listennserve is non-nil
	log.Println("starting on :4000")
	err := http.ListenAndServe(":4000", mux)
	log.Fatal(err)
}
