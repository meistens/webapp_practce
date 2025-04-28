package main

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"github.com/meistens/snippetbox/internal/models"
)

func (app *application) home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		app.notFound(w)
		return
	}
	snippets, err := app.snippets.Latest()
	if err != nil {
		app.serverErr(w, err)
		return
	}
	// Use the new render helper.
	app.render(w, http.StatusOK, "home.tmpl", &templateData{
		Snippets: snippets,
	})
}
func (app *application) snippetView(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil || id < 1 {
		app.notFound(w)
		return
	}
	snippet, err := app.snippets.Get(id)
	if err != nil {
		if errors.Is(err, models.ErrNoRecord) {
			app.notFound(w)
		} else {
			app.serverErr(w, err)
		}
		return
	}
	// Use the new render helper.
	app.render(w, http.StatusOK, "view.tmpl", &templateData{
		Snippet: snippet,
	})
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
