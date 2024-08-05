// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: query.sql

package querytest

import (
	"context"
	"database/sql"
)

const getAuthorsWithBooksCount = `-- name: GetAuthorsWithBooksCount :many
SELECT id, name, bio, (
  SELECT COUNT(id) FROM books
  WHERE books.author_id = id
) AS books_count
FROM authors
`

type GetAuthorsWithBooksCountRow struct {
	ID         int64
	Name       string
	Bio        sql.NullString
	BooksCount int64
}

func (q *Queries) GetAuthorsWithBooksCount(ctx context.Context) ([]GetAuthorsWithBooksCountRow, error) {
	rows, err := q.db.QueryContext(ctx, getAuthorsWithBooksCount)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []GetAuthorsWithBooksCountRow
	for rows.Next() {
		var i GetAuthorsWithBooksCountRow
		if err := rows.Scan(
			&i.ID,
			&i.Name,
			&i.Bio,
			&i.BooksCount,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}
