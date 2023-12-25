-- name: GetAllVehicleType :many
SELECT * FROM vehicle_type;

-- name: CountAllVehicleType :one
SELECT COUNT(*) FROM vehicle_type;

-- name: GetVehicleTypeByCode :one
SELECT * FROM vehicle_type where code = $1 LIMIT 1;

-- name: CreateVehicleType :exec
INSERT INTO vehicle_type (code, name, price, created_by) VALUES ($1, $2, $3, $4);
