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
) RETURNING id, user_id;