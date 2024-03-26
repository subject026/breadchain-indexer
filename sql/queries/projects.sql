-- name: CreateProject :one
INSERT INTO projects (id, created_at, updated_at, name, wallet_address)
VALUES ($1, $2, $3, $4, $5
)
RETURNING *;

-- name: GetProjects :many
SELECT * FROM projects;

-- name: GetProjectByAddress :one
SELECT * FROM projects WHERE wallet_address = $1;