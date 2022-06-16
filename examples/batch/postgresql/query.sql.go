// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.14.0
// source: query.sql

package batch

import (
	"context"
)

const createAuthor = `-- name: CreateAuthor :one
INSERT INTO authors (name) VALUES ($1)
RETURNING author_id, name
`

func (q *Queries) CreateAuthor(ctx context.Context, name string) (Author, error) {
	row := q.db.QueryRow(ctx, createAuthor, name)
	var i Author
	err := row.Scan(&i.AuthorID, &i.Name)
	return i, err
}

const getAuthor = `-- name: GetAuthor :one
SELECT author_id, name FROM authors
WHERE author_id = $1
`

func (q *Queries) GetAuthor(ctx context.Context, authorID int32) (Author, error) {
	row := q.db.QueryRow(ctx, getAuthor, authorID)
	var i Author
	err := row.Scan(&i.AuthorID, &i.Name)
	return i, err
}
