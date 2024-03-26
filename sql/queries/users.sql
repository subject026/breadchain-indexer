-- name: CreateUser :one
INSERT INTO users (id, created_at, wallet_address)
VALUES ($1, $2, $3)
RETURNING *;

-- name: GetUsers :many

SELECT * FROM users;

-- name: GetUserByAddress :one
SELECT * FROM users WHERE wallet_address = $1;