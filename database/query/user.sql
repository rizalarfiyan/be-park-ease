-- name: GetUserByUsername :one
SELECT * FROM users WHERE username = $1 LIMIT 1;

-- name: GetUserByToken :one
SELECT * FROM users WHERE token = $1 LIMIT 1;

-- name: UpdateUserToken :exec
UPDATE users
SET token = $1, expired_at = $2, updated_at = CURRENT_TIMESTAMP
WHERE id = $3;
