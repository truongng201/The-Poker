-- name: ResetPassword :exec
UPDATE users SET
    hashed_password = $2
WHERE user_id = $1;