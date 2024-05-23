package models

import (
	"database/sql"
	"time"
)

// Model of a snippet in the database
type Snippet struct {
	ID      string
	Title   string
	Content string
	Created time.Time
	Expires time.Time
}

// SnippetModel is a wrapper for a db connection pool.
type SnippetModel struct {
	DB *sql.DB
}

// Insert will create a new Snippet record.
func (m *SnippetModel) Insert(title, content string, expires int) (int, error) {
	return 0, nil
}

// Get will retrieve a Snippet record by ID or return an error if none exists.
func (m *SnippetModel) Get(id string) (Snippet, error) {
	return Snippet{}, nil
}

// Latest will return the 10 most recently created snippets.
func (m *SnippetModel) Latest() ([]Snippet, error) {
	return nil, nil
}
