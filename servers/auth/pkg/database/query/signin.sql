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