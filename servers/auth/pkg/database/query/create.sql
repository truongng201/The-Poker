-- name: CreateRefreshToken :one
INSERT INTO refresh_tokens (
    user_id,
    token,
    ip_address,
    user_agent,
    device_type
) VALUES (
    $1,
    $2,
    $3,
    $4,
    $5
) RETURNING id, user_id, token;

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