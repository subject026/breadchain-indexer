// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0
// source: test_users.sql

package database

import (
	"context"
	"time"

	"github.com/google/uuid"
)

const getTestUsers = `-- name: GetTestUsers :many
SELECT user_id, id, created_at, wallet_address FROM test_users JOIN users ON test_users.user_id = users.id
`

type GetTestUsersRow struct {
	UserID        uuid.UUID
	ID            uuid.UUID
	CreatedAt     time.Time
	WalletAddress string
}

func (q *Queries) GetTestUsers(ctx context.Context) ([]GetTestUsersRow, error) {
	rows, err := q.db.QueryContext(ctx, getTestUsers)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []GetTestUsersRow
	for rows.Next() {
		var i GetTestUsersRow
		if err := rows.Scan(
			&i.UserID,
			&i.ID,
			&i.CreatedAt,
			&i.WalletAddress,
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
