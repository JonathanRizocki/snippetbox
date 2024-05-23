package models

import (
	"context"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
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
	DB *pgxpool.Pool
}

// Insert will create a new Snippet record.
func (m *SnippetModel) Insert(title, content string, expires int) (Snippet, error) {
	stmt := `INSERT INTO snippets (title, content, created, expires)
	VALUES(?, ?, timezone('utc', current_timestamp), 
	current_timestamp + interval '? days' )`

	var item Snippet

	err := m.DB.QueryRow(context.Background(), stmt, title, content, expires).Scan(&item)
	if err != nil {
		return Snippet{}, err
	}
	return item, err
}

// Get will retrieve a Snippet record by ID or return an error if none exists.
func (m *SnippetModel) Get(id string) (Snippet, error) {
	return Snippet{}, nil
}

// Latest will return the 10 most recently created snippets.
func (m *SnippetModel) Latest() ([]Snippet, error) {
	return nil, nil
}
