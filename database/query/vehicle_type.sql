-- name: GetAllVehicleType :many
SELECT * FROM vehicle_type;

-- name: CountAllVehicleType :one
SELECT COUNT(*) FROM vehicle_type;
