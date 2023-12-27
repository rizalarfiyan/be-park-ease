-- name: GetAllLocation :many
SELECT * FROM location;

-- name: CountAllLocation :one
SELECT COUNT(*) FROM location;

-- name: GetLocationByCode :one
SELECT * FROM location where code = $1 LIMIT 1;

-- name: CreateLocation :exec
Insert into location (code, name, is_exit , created_by) values ($1, $2, $3, $4);

-- name: UpdateLocation :exec
UPDATE location SET code = $1, name = $2, is_exit = $3, updated_by = $4, updated_at = CURRENT_TIMESTAMP WHERE code = $5;

-- name: DeleteLocation :exec
UPDATE location SET deleted_by = $1, deleted_at = CURRENT_TIMESTAMP WHERE code = $2;
