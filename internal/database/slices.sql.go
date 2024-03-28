// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0
// source: slices.sql

package database

import (
	"context"
	"time"

	"github.com/google/uuid"
)

const createSlice = `-- name: CreateSlice :one

INSERT INTO slices (id, created_at, started_at)
VALUES ($1, $2, $3)
RETURNING id, created_at, started_at
`

type CreateSliceParams struct {
	ID        uuid.UUID
	CreatedAt time.Time
	StartedAt time.Time
}

func (q *Queries) CreateSlice(ctx context.Context, arg CreateSliceParams) (Slice, error) {
	row := q.db.QueryRowContext(ctx, createSlice, arg.ID, arg.CreatedAt, arg.StartedAt)
	var i Slice
	err := row.Scan(&i.ID, &i.CreatedAt, &i.StartedAt)
	return i, err
}

const createSliceProject = `-- name: CreateSliceProject :one
INSERT INTO slice_projects (id, created_at, project_id, slice_id, value)
VALUES ($1, $2, $3, $4, $5)
RETURNING id, created_at, project_id, slice_id, value
`

type CreateSliceProjectParams struct {
	ID        uuid.UUID
	CreatedAt time.Time
	ProjectID uuid.UUID
	SliceID   uuid.UUID
	Value     int32
}

func (q *Queries) CreateSliceProject(ctx context.Context, arg CreateSliceProjectParams) (SliceProject, error) {
	row := q.db.QueryRowContext(ctx, createSliceProject,
		arg.ID,
		arg.CreatedAt,
		arg.ProjectID,
		arg.SliceID,
		arg.Value,
	)
	var i SliceProject
	err := row.Scan(
		&i.ID,
		&i.CreatedAt,
		&i.ProjectID,
		&i.SliceID,
		&i.Value,
	)
	return i, err
}

const getLastSlice = `-- name: GetLastSlice :one
SELECT id, created_at, started_at FROM slices
ORDER BY created_at DESC
LIMIT 1
`

func (q *Queries) GetLastSlice(ctx context.Context) (Slice, error) {
	row := q.db.QueryRowContext(ctx, getLastSlice)
	var i Slice
	err := row.Scan(&i.ID, &i.CreatedAt, &i.StartedAt)
	return i, err
}

const getSlices = `-- name: GetSlices :many
SELECT slices.id, slices.created_at, started_at, slice_projects.id, slice_projects.created_at, project_id, slice_id, value FROM slices JOIN slice_projects ON slices.id = slice_projects.slice_id
ORDER BY slices.created_at DESC
`

type GetSlicesRow struct {
	ID          uuid.UUID
	CreatedAt   time.Time
	StartedAt   time.Time
	ID_2        uuid.UUID
	CreatedAt_2 time.Time
	ProjectID   uuid.UUID
	SliceID     uuid.UUID
	Value       int32
}

func (q *Queries) GetSlices(ctx context.Context) ([]GetSlicesRow, error) {
	rows, err := q.db.QueryContext(ctx, getSlices)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []GetSlicesRow
	for rows.Next() {
		var i GetSlicesRow
		if err := rows.Scan(
			&i.ID,
			&i.CreatedAt,
			&i.StartedAt,
			&i.ID_2,
			&i.CreatedAt_2,
			&i.ProjectID,
			&i.SliceID,
			&i.Value,
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
