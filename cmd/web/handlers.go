package main

import (
	"fmt"
	"html/template"
	"net/http"
	"strconv"
)

// modify func. sign. of home handler so it is defined as a method
// against *application
func (app *application) home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	// init a slice containing the paths to the two files
	// base template should be the first
	files := []string{
		"./ui/html/base.tmpl",
		"./ui/html/partials/nav.tmpl",
		"./ui/html/pages/home.tmpl",
	}
	// use template.ParseFiles() to read template into template
	// set
	// if error, log the detailed error msg and use the http.Error()
	// to send a generic 500 to user
	ts, err := template.ParseFiles(files...)
	if err != nil {
		// now that home handler func. is now a method against application
		// it can access its fields, including the error logger
		app.errLog.Println(err.Error())
		http.Error(w, "internal server err", 500)
		return
	}
	// use ExecuteTemplate() to write the content of the "base" template
	// as rresponse body
	err = ts.ExecuteTemplate(w, "base", nil)
	if err != nil {
		app.errLog.Println(err.Error())
		http.Error(w, "internal error", 500)
	}
}

func (app *application) snippetView(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil || id < 1 {
		http.NotFound(w, r)
		return
	}
	fmt.Fprintf(w, "display a specific snippet with ID %d", id)
}

func (app *application) snippetCreate(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.Header().Set("Allow", http.MethodPost)
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}
	w.Write([]byte("creating something..."))
}
