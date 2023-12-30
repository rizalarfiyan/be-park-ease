-- name: GetAllLocation :many
SELECT * FROM location;

-- name: CountAllLocation :one
SELECT COUNT(*) FROM location;

-- name: GetLocationByCode :one
SELECT * FROM location where code = $1 LIMIT 1;

-- name: CreateLocation :exec
Insert into location (code, name, is_exit , created_by) values ($1, $2, $3, $4);

-- name: UpdateLocation :exec
UPDATE location SET name = $1, is_exit = $2, updated_by = $3, updated_at = CURRENT_TIMESTAMP WHERE code = $4;

-- name: DeleteLocation :exec
UPDATE location SET deleted_by = $1, deleted_at = CURRENT_TIMESTAMP WHERE code = $2;
