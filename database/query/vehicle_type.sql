-- name: GetAllVehicleType :many
SELECT * FROM vehicle_type;

-- name: CountAllVehicleType :one
SELECT COUNT(*) FROM vehicle_type;

-- name: GetVehicleTypeByCode :one
SELECT * FROM vehicle_type where code = $1 LIMIT 1;

-- name: CreateVehicleType :exec
INSERT INTO vehicle_type (code, name, price, created_by) VALUES ($1, $2, $3, $4);

-- name: UpdateVehicleType :exec
UPDATE vehicle_type SET name = $1, price = $2, updated_by = $3, updated_at = CURRENT_TIMESTAMP WHERE code = $4;

-- name: DeleteVehicleType :exec
UPDATE vehicle_type SET deleted_by = $1, deleted_at = CURRENT_TIMESTAMP WHERE code = $2;
