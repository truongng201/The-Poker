-- name: HealthCheck :one
SELECT 1 AS "one";

-- name: GetUserByEmail :one
SELECT 
    id,
    user_id,
    email,
    username,
    image_url,
    is_verified,
    hashed_password
FROM users
WHERE email = $1;

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

