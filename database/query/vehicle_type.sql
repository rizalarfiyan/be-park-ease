-- name: GetAllVehicleType :many
SELECT * FROM vehicle_type;

-- name: CountAllVehicleType :one
SELECT COUNT(*) FROM vehicle_type;

-- name: GetVehicleTypeByCode :one
SELECT * FROM vehicle_type where code = $1 LIMIT 1;
