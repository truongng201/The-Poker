-- name: CreateUser :one
INSERT INTO users (
    user_id,
    username,
    email,
    hashed_password,
    image_url
) VALUES (
    $1,
    $2,
    $3,
    $4,
    $5
) RETURNING user_id, email;

-- name: CheckEmailExists :one
SELECT 
    id,
    user_id,
    email,
    is_verified
FROM users
WHERE email = $1;