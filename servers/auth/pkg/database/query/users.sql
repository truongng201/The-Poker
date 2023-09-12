-- name: GetUsers :many
SELECT 
    id,
    user_id,
    email,
    username,
    image_url
FROM users
ORDER BY created_at DESC
LIMIT $1
OFFSET $2;

-- name: GetUserById :one
SELECT 
    id,
    user_id,
    email,
    username,
    image_url
FROM users
WHERE id = $1;

-- name: GetUserByEmail :one
SELECT 
    id,
    user_id,
    email,
    username,
    image_url
FROM users
WHERE email = $1;

-- name: CreateUser :one
INSERT INTO users (
    user_id,
    email,
    username,
    hashed_password,
    image_url
) VALUES (
    $1,
    $2,
    $3,
    $4,
    $5
) RETURNING user_id, email;

-- name: UpdateUser :one
UPDATE users SET
    email = $1,
    username = $2,
    hashed_password = $3,
    image_url = $4
WHERE email = $5
RETURNING user_id, email;

-- name: DeleteUser :exec
DELETE FROM users
WHERE id = $1;