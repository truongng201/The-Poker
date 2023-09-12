-- name: VerifyEmail :exec
UPDATE users SET
    is_verified = TRUE
WHERE email = $1 AND is_verified = FALSE;
