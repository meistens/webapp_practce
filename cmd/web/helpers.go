package main

import (
	"fmt"
	"net/http"
	"runtime/debug"
)

// serverErr helper writes an error msg and stack trace to the errLog
// sends a generic 500 response to the user
func (app *application) serverErr(w http.ResponseWriter, err error) {
	trace := fmt.Sprintf("%s\n%s", err.Error(), debug.Stack())
	app.errorLog.Output(2, trace)

	http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
}

// clientErr helper sends a specific status code and corresponding description
// to the user
func (app *application) clientErr(w http.ResponseWriter, status int) {
	http.Error(w, http.StatusText(status), status)
}

// for constistency, implement a notFound helper
// which returns a 404
func (app *application) notFound(w http.ResponseWriter) {
	app.clientErr(w, http.StatusNotFound)
}

func (app *application) render(w http.ResponseWriter, status int, page string, data *templateData) {
	// Retrieve the appropriate template set from the cache based on the page
	// name (like 'home.tmpl'). If no entry exists in the cache with the
	// provided name, then create a new error and call the serverError() helper
	// method that we made earlier and return.
	ts, ok := app.templateCache[page]
	if !ok {
		err := fmt.Errorf("the template %s does not exist", page)
		app.serverErr(w, err)
		return
	}
	// Write out the provided HTTP status code ('200 OK', '400 Bad Request'
	// etc).
	w.WriteHeader(status)
	// Execute the template set and write the response body. Again, if there
	// is any error we call the the serverError() helper.
	err := ts.ExecuteTemplate(w, "base", data)
	if err != nil {
		app.serverErr(w, err)
	}
}
