-- name: ResetPassword :exec
UPDATE users SET
    hashed_password = $2
WHERE user_id = $1;

-- name: VerifyEmail :exec
UPDATE users SET
    is_verified = TRUE
WHERE email = $1 AND is_verified = FALSE;