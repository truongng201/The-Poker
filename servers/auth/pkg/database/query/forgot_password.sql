-- name: FindUserByEmail :one
SELECT
    user_id,
    email,
    is_verified,
    username
FROM users
WHERE email = $1;