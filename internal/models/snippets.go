package models

import (
	"database/sql"
	"errors"
	"time"
)

// Define a Snippet type to hold the data for an individual snippet. Notice how
// the fields of the struct correspond to the fields in our MySQL snippets
// table?
type Snippet struct {
	ID      int
	Title   string
	Content string
	Created time.Time
	Expires time.Time
}

// Define a SnippetModel type which wraps a sql.DB connection pool.
type SnippetModel struct {
	DB *sql.DB
}

// This will insert a new snippet into the database.
func (m *SnippetModel) Insert(title string, content string, expires int) (int, error) {
	// write sql statement
	stmt := `INSERT INTO snippets (title, content, created, expires) VALUES(?, ?, UTC_TIMESTAMP(), DATE_ADD(UTC_TIMESTAMP(), INTERVAL ? DAY))`

	// use Exec() to execute connection pool to execute statement
	result, err := m.DB.Exec(stmt, title, content, expires)
	if err != nil {
		return 0, err
	}
	// use lastinsertid() to get id of newly inserted record
	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}
	// id returned has type int64, convert it to an int type
	return int(id), nil
}

// shortened version of previous block
func (m *SnippetModel) Get(id int) (*Snippet, error) {
	s := &Snippet{}
	err := m.DB.QueryRow("SELECT ...", id).Scan(&s.ID, &s.Title, &s.Content, &s.Created, &s.Expires)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, ErrNoRecord
		} else {
			return nil, err
		}
	}
	return s, nil
}

// This will return the 10 most recently created snippets.
func (m *SnippetModel) Latest() ([]*Snippet, error) {
	stmt := `SELECT id, title, content, created, expires FROM snippets WHERE expires > UTC_TIMESTAMP() ORDER BY id DESC LIMIT 10`

	rows, err := m.DB.Query(stmt)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	// init an empty slice to hold the snippet structs
	snippets := []*Snippet{}

	// iterate through the db
	for rows.Next() {
		// pointer to a new zeroed snippet struct
		s := &Snippet{}
		// rows.Scan() to copy values from each field in the row to the new snippet object created
		err = rows.Scan(&s.ID, &s.Title, &s.Content, &s.Created, &s.Expires)
		if err != nil {
			return nil, err
		}
		// append to slice of snippets
		snippets = append(snippets, s)
	}
	// call rows.Err() when rows.Next loop completes
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return snippets, nil
}
