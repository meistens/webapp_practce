package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
)

// defne a home handler function which writes a byte slice containing
// "some words" as the response body
func home(w http.ResponseWriter, r *http.Request) {
	// check if the current request URL path exactly matches "/"
	// if it doesn't, use http.NotFound() to send a 404
	// return from the handler for normal ops
	// if not returned from the handler, keep executing and also
	// write the "some words"
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	w.Write([]byte("some words"))
}

// add a snippetCreate and snippetView handler function
func snippetCreate(w http.ResponseWriter, r *http.Request) {
	// refactor (check vc for the diff)
	if r.Method != http.MethodPost {
		w.Header().Set("Allow", http.MethodPost)
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}
	w.Write([]byte("create some stuff..."))
}
func snippetView(w http.ResponseWriter, r *http.Request) {
	// extract the value of the id params from the query string
	// and try to convert it to an integer using strconv.Atoi()
	// if it cannot be converted to an int, or value is less
	// than 1, 404
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil || id < 1 {
		http.NotFound(w, r)
		return
	}
	fmt.Fprintf(w, "display a specific snippet with ID %d", id)
}

func main() {
	// use the http.NewServerMux() func to init a new server
	// then register the home func as the handler for the "/" url pattern
	mux := http.NewServeMux()
	mux.HandleFunc("/", home)

	// register the two new handler funcs and corresponding URL patterns with
	// the servemux, same way as the first one
	mux.HandleFunc("/snippet/view", snippetView)
	mux.HandleFunc("/snippet/create", snippetCreate)

	// use the http.ListenAndServe() func to start a new web server
	// pass in 2 params: TCP to listen on and ServeMux just created
	// if listenandserve returns an error, log.fatal() to log error msg and exit
	// Note-> any err returned by listennserve is non-nil
	log.Println("starting on :4000")
	err := http.ListenAndServe(":4000", mux)
	log.Fatal(err)
}
