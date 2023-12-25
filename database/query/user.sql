-- name: GetUserByUsername :one
SELECT * FROM users WHERE username = $1 LIMIT 1;

-- name: GetUserByToken :one
SELECT * FROM users WHERE token = $1 LIMIT 1;

-- name: UpdateUserToken :exec
UPDATE users
SET token = $1, expired_at = $2, updated_at = CURRENT_TIMESTAMP
WHERE id = $3;

-- name: GetAllUser :many
select * from users;

-- name: CountAllUser :one
SELECT COUNT(*) FROM users;

-- name: GetUserById :one
SELECT * FROM users WHERE id = $1 LIMIT 1;

-- name: CreateUser :exec
insert into users (name, username, password, role, status)
values ($1, $2, $3, $4, $5);

-- name: UpdateUser :exec
UPDATE users
SET name = $1, username = $2, password = $3, role = $4, status = $5, updated_at = CURRENT_TIMESTAMP
WHERE id = $6;

-- name: UpdatePassword :exec
UPDATE users
SET password = $1, updated_at = CURRENT_TIMESTAMP
WHERE id = $2;
