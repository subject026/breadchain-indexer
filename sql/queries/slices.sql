-- name: CreateSlice :one

INSERT INTO slices (id, created_at, started_at)
VALUES ($1, $2, $3)
RETURNING *;

-- name: CreateSliceProject :one
INSERT INTO slice_projects (id, created_at, project_id, slice_id, value)
VALUES ($1, $2, $3, $4, $5)
RETURNING *;

-- name: GetSlices :many
SELECT * FROM slices JOIN slice_projects ON slices.id = slice_projects.slice_id
ORDER BY slices.created_at DESC;

-- name: GetLastSlice :one
SELECT * FROM slices
ORDER BY created_at DESC
LIMIT 1;
