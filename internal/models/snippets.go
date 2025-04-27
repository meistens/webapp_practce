package models

import (
	"database/sql"
	"time"
)

// define a snippet that holds the data for an individual snippet
type Snippet struct {
	ID      int
	Title   string
	Content string
	Created time.Time
	Expires time.Time
}

// define a snippetModel type which wraps a sql.DB pool
type snippetModel struct {
	DB *sql.DB
}

// insert new snippet to db
func (m *snippetModel) Insert(title string, content string, expires int) (int, error) {
	return 0, nil
}

// This will return a specific snippet based on its id.
func (m *snippetModel) Get(id int) (*Snippet, error) {
	return nil, nil
}

// This will return the 10 most recently created snippets.
func (m *snippetModel) Latest() ([]*Snippet, error) {
	return nil, nil
}
