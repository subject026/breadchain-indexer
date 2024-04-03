-- name: GetTestUsers :many
SELECT * FROM test_users JOIN users ON test_users.user_id = users.id;