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
		app.notFound(w)
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
		app.serverErr(w, err)
		return
	}
	// use ExecuteTemplate() to write the content of the "base" template
	// as response body
	err = ts.ExecuteTemplate(w, "base", nil)
	if err != nil {
		app.serverErr(w, err)
	}
}

func (app *application) snippetView(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil || id < 1 {
		app.notFound(w)
		return
	}
	fmt.Fprintf(w, "display a specific snippet with ID %d", id)
}

func (app *application) snippetCreate(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.Header().Set("Allow", http.MethodPost)
		app.clientErr(w, http.StatusMethodNotAllowed)
		return
	}

	// create some variables holding dummy data
	title := "0 small"
	content := "using swear words\nis a natural\npart of every human"
	expires := 7

	id, err := app.snippets.Insert(title, content, expires)
	if err != nil {
		app.serverErr(w, err)
		return
	}
	// redirect user to relevant page for snippet
	http.Redirect(w, r, fmt.Sprintf("/snippet/view?id=%d", id), http.StatusSeeOther)
}
