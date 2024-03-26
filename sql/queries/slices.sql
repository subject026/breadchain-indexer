-- name: CreateSlice :one

INSERT INTO slices (id, created_at, started_at)
VALUES ($1, $2, $3)
RETURNING *;

-- name: CreateSliceProject :one
INSERT INTO slice_projects (id, created_at, slice_id, value)
VALUES ($1, $2, $3, $4)
RETURNING *;

-- name: GetLastSlice :one
SELECT * FROM slices
ORDER BY created_at DESC
LIMIT 1;