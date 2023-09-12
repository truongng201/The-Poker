// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.21.0
// source: signout.sql

package database

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
)

const deleteRefreshToken = `-- name: DeleteRefreshToken :exec
DELETE FROM refresh_tokens
WHERE token = $1
`

func (q *Queries) DeleteRefreshToken(ctx context.Context, token string) error {
	_, err := q.db.Exec(ctx, deleteRefreshToken, token)
	return err
}

const getUserByRefreshToken = `-- name: GetUserByRefreshToken :one
SELECT
    users.id,
    refresh_tokens.id,
    users.user_id,
    users.email
FROM users
LEFT JOIN refresh_tokens
    ON users.id = refresh_tokens.user_id
WHERE refresh_tokens.token = $1 
    AND users.email = $2
`

type GetUserByRefreshTokenParams struct {
	Token string `json:"token"`
	Email string `json:"email"`
}

type GetUserByRefreshTokenRow struct {
	ID     int64       `json:"id"`
	ID_2   pgtype.Int8 `json:"id_2"`
	UserID string      `json:"user_id"`
	Email  string      `json:"email"`
}

func (q *Queries) GetUserByRefreshToken(ctx context.Context, arg GetUserByRefreshTokenParams) (GetUserByRefreshTokenRow, error) {
	row := q.db.QueryRow(ctx, getUserByRefreshToken, arg.Token, arg.Email)
	var i GetUserByRefreshTokenRow
	err := row.Scan(
		&i.ID,
		&i.ID_2,
		&i.UserID,
		&i.Email,
	)
	return i, err
}
