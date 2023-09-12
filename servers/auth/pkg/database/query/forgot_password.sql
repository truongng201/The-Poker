-- name: FindUserByEmail :one
SELECT
    user_id,
    email,
    is_verified
FROM users
WHERE email = $1;