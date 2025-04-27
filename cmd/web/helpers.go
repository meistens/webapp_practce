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
