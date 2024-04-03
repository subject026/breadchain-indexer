-- name: CreateVote :one
INSERT INTO votes (id, created_at, user_id, project_id, value)
VALUES ($1, $2, $3, $4, $5
)
RETURNING *;

-- name: GetVotes :many
SELECT * FROM votes;

-- name: GetVotesInRange :many
SELECT votes.id, votes.created_at, votes.project_id, votes.value, users.wallet_address FROM votes
INNER JOIN users ON votes.user_id = users.id
WHERE votes.created_at > $1 AND votes.created_at < $2 ORDER BY votes.created_at DESC;

-- name: GetVotesInRangeForUser :many
SELECT * FROM votes WHERE user_id = $1 AND created_at > $2 AND created_at < $3 ORDER BY created_at DESC;
