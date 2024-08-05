// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: query-building.sql

package jets

import (
	"context"
)

const countPilots = `-- name: CountPilots :one
SELECT COUNT(*) FROM pilots
`

func (q *Queries) CountPilots(ctx context.Context) (int64, error) {
	row := q.db.QueryRow(ctx, countPilots)
	var count int64
	err := row.Scan(&count)
	return count, err
}

const deletePilot = `-- name: DeletePilot :exec
DELETE FROM pilots WHERE id = $1
`

func (q *Queries) DeletePilot(ctx context.Context, id int32) error {
	_, err := q.db.Exec(ctx, deletePilot, id)
	return err
}

const listPilots = `-- name: ListPilots :many
SELECT id, name FROM pilots LIMIT 5
`

func (q *Queries) ListPilots(ctx context.Context) ([]Pilot, error) {
	rows, err := q.db.Query(ctx, listPilots)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Pilot
	for rows.Next() {
		var i Pilot
		if err := rows.Scan(&i.ID, &i.Name); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}
