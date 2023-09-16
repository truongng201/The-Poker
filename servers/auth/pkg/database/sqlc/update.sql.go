// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.21.0
// source: update.sql

package database

import (
	"context"
)

const resetPassword = `-- name: ResetPassword :exec
UPDATE users SET
    hashed_password = $2
WHERE user_id = $1
`

type ResetPasswordParams struct {
	UserID         string `json:"user_id"`
	HashedPassword string `json:"hashed_password"`
}

func (q *Queries) ResetPassword(ctx context.Context, arg ResetPasswordParams) error {
	_, err := q.db.Exec(ctx, resetPassword, arg.UserID, arg.HashedPassword)
	return err
}

const verifyEmail = `-- name: VerifyEmail :exec
UPDATE users SET
    is_verified = TRUE
WHERE email = $1 AND is_verified = FALSE
`

func (q *Queries) VerifyEmail(ctx context.Context, email string) error {
	_, err := q.db.Exec(ctx, verifyEmail, email)
	return err
}