-- name: CreateUser :one
INSERT INTO users (
    user_id,
    username,
    email,
    hashed_password
) VALUES (
    $1,
    $2,
    $3,
    $4
) RETURNING user_id, email;

-- name: FindUserByEmail :one
SELECT 
    id,
    user_id,
    email,
    is_verified
FROM users
WHERE email = $1;