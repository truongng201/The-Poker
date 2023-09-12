-- name: GetUserByRefreshToken :one
SELECT
    users.id,
    refresh_tokens.id,
    users.user_id,
    users.email
FROM users
LEFT JOIN refresh_tokens
    ON users.id = refresh_tokens.user_id
WHERE refresh_tokens.token = $1 
    AND users.email = $2;

-- name: DeleteRefreshToken :exec
DELETE FROM refresh_tokens
WHERE token = $1;